package kable

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Kable-io/kable-go/v2/internal/openapi"
	"github.com/ehsaniara/gointerlock"
	"github.com/google/uuid"
)

type Event struct {
	ClientId      string                 `json:"clientId"`
	Data          map[string]interface{} `json:"data,omitempty"`
	TransactionId string                 `json:"transactionId,omitempty"`
	Timestamp     *time.Time             `json:"timestamp,omitempty"`
}

type EventsApi struct {
	api     *openapi.APIClient
	options *KableOptions
	ctx     context.Context
	queue   []Event
	lock    sync.Mutex
}

func (e *EventsApi) sendEvents(events []Event) error {
	req := e.api.EventsApi.CreateEvents(context.Background())

	var openapiEvents []openapi.Event
	for _, event := range events {
		var timestamp = time.Now()
		if event.Timestamp != nil {
			timestamp = *event.Timestamp
		}

		if event.TransactionId == "" {
			transactionId := uuid.New().String()
			event.TransactionId = transactionId
		}

		openapiEvents = append(openapiEvents, openapi.Event{
			ClientId:      event.ClientId,
			Timestamp:     timestamp,
			TransactionId: &event.TransactionId,
			Data:          event.Data,
		})
	}

	var countToSend = len(openapiEvents)

	if e.options.Debug {
		log.Printf("[KABLE] Sending %d events", countToSend)
	}

	req = req.Event(openapiEvents)
	req = req.KableClientId(e.options.KableClientId)
	req = req.KableClientSecret(e.options.KableClientSecret)

	_, err := req.Execute()
	if err != nil {
		log.Printf("[KABLE] Failed to send %d events to Kable: %s", countToSend, err)
		for _, event := range openapiEvents {
			jsonEvent, err := json.Marshal(event)
			if err != nil {
				log.Printf("[KABLE] Kable Event (Error): %+v", event)
			} else {
				log.Printf("[KABLE] Kable Event (Error): %s", jsonEvent)
			}
		}
		return err
	}

	log.Printf("[KABLE] Successfully sent %d events to Kable server", countToSend)

	return nil
}

func (e *EventsApi) flush() {
	e.lock.Lock()
	defer e.lock.Unlock()

	if e.options.Debug {
		log.Printf("[KABLE] Flushing Kable event queue...")
	}

	if len(e.queue) <= 0 {
		if e.options.Debug {
			log.Printf("[KABLE] ...no Kable events to flush...")
		}
		return
	}

	err := e.sendEvents(e.queue)
	if err != nil {
		return // return without clearing the queue
	}

	// clear all events from the queue, as they must have been sent successfully
	e.queue = e.queue[:0]
}

func (e *EventsApi) scheduleFlushQueue() {
	var job = gointerlock.GoInterval{
		Interval: 10 * time.Second,
		Arg:      e.flush,
	}

	err := job.Run(e.ctx)
	if err != nil {
		log.Printf("[KABLE] Error running scheduled flush queue: %s", err)
	}
}

func (e *EventsApi) handleShutdown(stop context.CancelFunc) {
	defer stop()
	<-e.ctx.Done()
	log.Println("[KABLE] Shutdown signal received, flushing event queue")
	e.flush()

	// TODO: is this going to send the right shutdown signal?
	log.Fatal("[KABLE] Shutting down", e.ctx.Err())
}

func NewEventsApi(apiClient *openapi.APIClient, options *KableOptions) *EventsApi {
	eventsApi := &EventsApi{
		api:     apiClient,
		options: options,
	}

	// if the queue size is less than 2, then recording any number of events will result
	// in calling flush, so there will be no need to initiate the queue, as it won't be used.
	// also, as flush won't be needed, then there's no need
	if options.MaxQueueSize > 1 {
		// initiate the queue with the max queue size specified, so that reporting any
		// number of events below that number will not result in allocating a new queue.
		eventsApi.queue = make([]Event, 0, options.MaxQueueSize)
	}

	// only start the scheduled goroutines if the queue was initialized
	if eventsApi.queue != nil {
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		eventsApi.ctx = ctx

		go eventsApi.scheduleFlushQueue()
		go eventsApi.handleShutdown(stop)
	}

	return eventsApi
}

func (e *EventsApi) EnqueueEvent(events ...Event) {
	// if queue size is less than 2, then no need to add the events to the queue, so send
	// them directly and return(no need to clear the queue, as it must be already empty).
	if e.options.MaxQueueSize < 2 {
		err := e.sendEvents(events)
		if err != nil {
			return // return without retry - the events will be logged in the sendEvents functions
		}
		return
	}

	// acquire the lock while modifying the queue, to prevent concurrent modifications
	// to the queue from other EnqueueEvent calls, and wait for any ongoing flush call.
	e.lock.Lock()
	e.queue = append(e.queue, events...)
	queueLen := len(e.queue)
	e.lock.Unlock()

	if queueLen >= e.options.MaxQueueSize {
		e.flush()
	}
}
