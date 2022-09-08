package kable

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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
}

type RecordEventOut struct {
	Message string `json:"message"`
}

func (e *EventsApi) handleFlush() {
	if len(e.queue) > 0 {
		res, err := e.flush()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Flush Response : ", res.Body)
	}
}

func (e *EventsApi) scheduleFlushQueue() {
	var job = gointerlock.GoInterval{
		Interval: 10 * time.Second,
		Arg:      e.handleFlush,
	}

	err := job.Run(e.ctx)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func (e *EventsApi) handleShutdown(stop context.CancelFunc) {
	defer stop()
	for {
		select {
		case <-e.ctx.Done():
			log.Println("Shutdown signal received. Flushing event queue.")
			e.handleFlush()
			return
		}
	}

}

func NewEventsApi(apiClient *openapi.APIClient, options *KableOptions) *EventsApi {
	providedCtx := *options.Context
	extendedCtx, stop := signal.NotifyContext(providedCtx, os.Interrupt, syscall.SIGTERM)

	eventsApi := &EventsApi{
		api:     apiClient,
		options: options,
		ctx:     extendedCtx,
		queue:   []Event{},
	}

	go eventsApi.scheduleFlushQueue()
	go eventsApi.handleShutdown(stop)

	// Listening to the OS Signals

	return eventsApi
}

func (e *EventsApi) flush() (*http.Response, error) {
	req := e.api.EventsApi.CreateEvents(*e.options.Context)
	var openapiEvents []openapi.Event
	for _, event := range e.queue {
		openapiEvents = append(openapiEvents, openapi.Event(event))
	}

	req = req.Event(openapiEvents)
	req = req.KableClientId(e.options.KableClientId)
	req = req.KableClientSecret(e.options.KableClientSecret)

	res, err := req.Execute()
	if err != nil {
		return nil, wrapError(err, res)
	}

	// Set the queue to empty.
	e.queue = []Event{}

	return res, nil
}

func (e *EventsApi) Record(events ...Event) {
	e.queue = append(e.queue, events...)
	if len(e.queue) >= e.options.MaxQueueSize {
		res, err := e.flush()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Response : ", res)
	}
}
