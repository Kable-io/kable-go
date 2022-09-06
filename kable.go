package kable

import (
	"fmt"
	"github.com/Kable-io/kable-go/internal/openapi"
	"net/http"
	"net/url"
	"time"
)

type (
	KableOptions struct {
		KableClientId     string
		KableClientSecret string
		Debug             bool
		MaxQueueSize      int
		//RecordAuthentication bool

		BaseUrl    *url.URL
		HTTPClient *http.Client
	}

	Kable struct {
		EventsApi *EventsApi
		queue     []Event
	}
)

var defaultHTTPClient = &http.Client{
	Timeout: 60 * time.Second,
}

func New(options *KableOptions) *Kable {
	conf := openapi.NewConfiguration()
	conf.Scheme = "https"
	conf.Host = "test.kable.io"
	conf.HTTPClient = defaultHTTPClient
	if options != nil {
		conf.Debug = options.Debug
		if options.BaseUrl != nil {
			conf.Scheme = options.BaseUrl.Scheme
			conf.Host = options.BaseUrl.Host
		}
	}

	// TODO : nil checks here
	//conf.AddDefaultHeader("Kable-Client-Id", options.KableClientId)
	//conf.AddDefaultHeader("Kable-Client-Secret", options.KableClientSecret)
	conf.AddDefaultHeader("Content-Type", "application/json")
	conf.AddDefaultHeader("Kable-Environment", "TEST")
	// TODO : Version introduction
	conf.UserAgent = fmt.Sprintf("kable-libs/%s/go", "0.0.1")

	apiClient := openapi.NewAPIClient(conf)

	return &Kable{
		EventsApi: &EventsApi{
			api:     apiClient,
			options: options,
		},
	}
}
