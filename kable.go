package kable

import (
	"encoding/json"
	"fmt"
	"github.com/Kable-io/kable-go/internal/auth"
	"github.com/Kable-io/kable-go/internal/openapi"
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
		BaseUrl           string
		HTTPClient        *http.Client
	}

	Kable struct {
		eventsApi       *EventsApi
		authenticateApi *AuthenticateApi
		options         *KableOptions
	}
)

func New(options *KableOptions) *Kable {
	apiConf := openapi.NewConfiguration()
	authConf := auth.NewConfiguration()

	if options != nil {
		if options.Debug {
			jsonOptions, err := json.Marshal(options)
			if err != nil {
				log.Println("Something is wrong with the options provided to Kable. Cannot parse to JSON.")
			}
			log.Println("Instantiating Kable with options : ", string(jsonOptions))
		}
		baseUrl, err := url.Parse(options.BaseUrl)
		if err != nil {
			log.Fatal("baseUrl is not a valid url.", err)
		}
		apiConf.Scheme = baseUrl.Scheme
		apiConf.Host = baseUrl.Host
		authConf.Scheme = baseUrl.Scheme
		authConf.Host = baseUrl.Host

	} else {
		log.Fatal("Must specify kable options")
	}

	apiConf.AddDefaultHeader("Content-Type", "application/json")
	authConf.AddDefaultHeader("Content-Type", "application/json")

	apiConf.AddDefaultHeader("KABLE-ENVIRONMENT", "TEST")
	authConf.AddDefaultHeader("KABLE-ENVIRONMENT", "TEST")

	apiConf.UserAgent = fmt.Sprintf("kable-libs/%s/go", version.Version)
	authConf.UserAgent = fmt.Sprintf("kable-libs/%s/go", version.Version)

	// TODO : populate conf here.

	apiClient := openapi.NewAPIClient(apiConf)
	authClient := auth.NewAPIClient(authConf)

	authApi := NewAuthenticateApi(authClient, options)
	err := authApi.Authenticate()

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully authenticated...")
	}

	return &Kable{
		eventsApi:       NewEventsApi(apiClient, options),
		authenticateApi: NewAuthenticateApi(authClient, options),
		options:         options,
	}
}

func (k *Kable) Record(events ...Event) {
	if k.options.Debug {
		log.Printf("Received %d event(s) to record.", len(events))
	}
	k.eventsApi.Record(events...)
}
