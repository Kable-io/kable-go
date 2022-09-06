package kable

import (
	"context"
	"github.com/Kable-io/kable-go/internal/openapi"
	"net/http"
)

type (
	Event openapi.Event
)

type EventsApi struct {
	api     *openapi.APIClient
	options *KableOptions
}

func (e *EventsApi) Record(events ...Event) (*http.Response, error) {
	req := e.api.EventsApi.CreateEvents(context.Background())
	var openapiEvents []openapi.Event
	for _, event := range events {
		openapiEvents = append(openapiEvents, openapi.Event(event))
	}

	req = req.Event(openapiEvents)
	req = req.KableClientId(e.options.KableClientId)
	req = req.KableClientSecret(e.options.KableClientSecret)

	res, err := req.Execute()
	if err != nil {
		return nil, wrapError(err, res)
	}

	return res, nil
}
