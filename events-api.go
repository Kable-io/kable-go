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

	"github.com/Kable-io/kable-go/internal/openapi"
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

	req := e.api.EventsApi.CreateEvents(context.Background())

	var openapiEvents []openapi.Event
	for _, event := range e.queue {
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
		log.Printf("[KABLE] Flushing %d events", countToSend)
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
				log.Printf("[KABLE] Kable Event (Error): %T", event)
			} else {
				log.Printf("[KABLE] Kable Event (Error): %s", jsonEvent)
			}
		}
		return
	}

	log.Printf("[KABLE] Successfully sent %d events to Kable server", countToSend)

	// Remove events from queue that have been sent
	e.queue = e.queue[countToSend:]
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
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	eventsApi := &EventsApi{
		api:     apiClient,
		options: options,
		ctx:     ctx,
		queue:   []Event{},
	}

	go eventsApi.scheduleFlushQueue()
	go eventsApi.handleShutdown(stop)

	// Listening to the OS Signals

	return eventsApi
}

func (e *EventsApi) EnqueueEvent(events ...Event) {
	e.queue = append(e.queue, events...)
	if len(e.queue) >= e.options.MaxQueueSize {
		e.flush()
	}
}
