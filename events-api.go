package kable

import (
	"context"
	"encoding/json"
	"github.com/Kable-io/kable-go/internal/openapi"
	"github.com/ehsaniara/gointerlock"
	"io"
	"log"
	"time"
)

type (
	Event openapi.Event
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

		log.Println("Flush Response : ", res.Message)
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

func NewEventsApi(apiClient *openapi.APIClient, options *KableOptions) *EventsApi {
	eventsApi := &EventsApi{
		api:     apiClient,
		options: options,
		ctx:     context.Background(),
		queue:   []Event{},
	}

	go eventsApi.scheduleFlushQueue()

	return eventsApi
}

func (e *EventsApi) flush() (*RecordEventOut, error) {
	req := e.api.EventsApi.CreateEvents(context.Background())
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

	ret := &RecordEventOut{}
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bytes, ret)
	if err != nil {
		log.Fatal(err)
	}

	return ret, nil
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
