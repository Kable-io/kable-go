package kable

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Kable-io/kable-go/internal/openapi"
	"github.com/ehsaniara/gointerlock"
)

type (
	Event       openapi.Event
	ApiResponse interface{}
)

type EventsApi struct {
	api     *openapi.APIClient
	options *KableOptions
	ctx     context.Context
	queue   []Event
	lock    sync.Mutex
}

type RecordEventOut struct {
	Message string `json:"message"`
}

func (e *EventsApi) handleFlush() {
	e.lock.Lock()
	defer e.lock.Unlock()

	if len(e.queue) > 0 {
		if e.options.Debug {
			log.Printf("Flushing queue...")
		}
		res, toFlush, err := e.flush()
		if err != nil {
			log.Printf("Error flushing queue: %s", err)
			jsonOutput, err := json.Marshal(toFlush)
			if err != nil {
				log.Printf("Unable to convert events to JSON: %s", err)
			}
			log.Printf("[Kable] : %s", jsonOutput)
		} else if e.options.Debug {
			log.Printf("Flushed queue. Response : %s", res.Body)
		}
	}
}

func (e *EventsApi) scheduleFlushQueue() {
	var job = gointerlock.GoInterval{
		Interval: 10 * time.Second,
		Arg:      e.handleFlush,
	}

	err := job.Run(e.ctx)
	if err != nil {
		log.Printf("Error running scheduled flush queue: %s", err)
	}
}

func (e *EventsApi) handleShutdown(stop context.CancelFunc) {
	defer stop()
	<-e.ctx.Done()
	log.Println("Shutdown signal received. Flushing event queue.")
	e.handleFlush()
	log.Fatal("Shutting down.", e.ctx.Err())
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

func (e *EventsApi) flush() (*http.Response, []Event, error) {
	req := e.api.EventsApi.CreateEvents(context.Background())
	var toFlush []Event
	if len(e.queue) > e.options.MaxQueueSize {
		temp := e.queue[0:e.options.MaxQueueSize]
		toFlush = temp
	} else {
		toFlush = e.queue
	}

	e.queue = e.queue[len(toFlush):len(e.queue)]

	var openapiEvents []openapi.Event
	for _, event := range toFlush {
		openapiEvents = append(openapiEvents, openapi.Event(event))
	}

	if e.options.Debug {
		log.Printf("Attempting to flush %d events.", len(openapiEvents))
	}

	req = req.Event(openapiEvents)
	req = req.KableClientId(e.options.KableClientId)
	req = req.KableClientSecret(e.options.KableClientSecret)

	res, err := req.Execute()
	if err != nil {
		return nil, toFlush, wrapError(err, res)
	}

	// Set the queue to empty.
	e.queue = []Event{}

	return res, toFlush, nil
}

func (e *EventsApi) Record(events ...Event) {
	e.queue = append(e.queue, events...)
	if len(e.queue) >= e.options.MaxQueueSize {
		e.handleFlush()
	}
}
