package kable

import (
	"context"
	"fmt"
	"github.com/Kable-io/kable-go/internal/openapi"
	"github.com/Kable-io/kable-go/internal/sdk"
	"github.com/Kable-io/kable-go/internal/version"
	"log"
	"net/http"
	"net/url"
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
		Context    *context.Context
	}

	Kable struct {
		eventsApi       *EventsApi
		authenticateApi *AuthenticateApi
	}
)

func New(options *KableOptions) *Kable {
	apiConf := openapi.NewConfiguration()
	sdkConf := sdk.NewConfiguration()

	if options != nil {
		if options.BaseUrl != nil {
			apiConf.Scheme = options.BaseUrl.Scheme
			apiConf.Host = options.BaseUrl.Host
			sdkConf.Scheme = options.BaseUrl.Scheme
			sdkConf.Host = options.BaseUrl.Host
		} else {
			log.Fatal("Must specify baseUrl")
		}
	} else {
		log.Fatal("Must specify kable options")
	}

	apiConf.AddDefaultHeader("Content-Type", "application/json")
	sdkConf.AddDefaultHeader("Content-Type", "application/json")

	apiConf.AddDefaultHeader("KABLE-ENVIRONMENT", "TEST")
	sdkConf.AddDefaultHeader("KABLE-ENVIRONMENT", "TEST")

	apiConf.UserAgent = fmt.Sprintf("kable-libs/%s/go", version.Version)
	sdkConf.UserAgent = fmt.Sprintf("kable-libs/%s/go", version.Version)

	// TODO : populate conf here.

	apiClient := openapi.NewAPIClient(apiConf)
	sdkClient := sdk.NewAPIClient(sdkConf)

	authApi := NewAuthenticateApi(sdkClient, options)
	err := authApi.TestAuthentication()

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully authenticated...")
	}

	return &Kable{
		eventsApi:       NewEventsApi(apiClient, options),
		authenticateApi: NewAuthenticateApi(sdkClient, options),
	}
}

func (k *Kable) Record(events ...Event) {
	k.eventsApi.Record(events...)
}
