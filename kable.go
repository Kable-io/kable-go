package kable

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/Kable-io/kable-go/internal/auth"
	"github.com/Kable-io/kable-go/internal/openapi"
	"github.com/Kable-io/kable-go/internal/version"
)

type (
	KableOptions struct {
		KableClientId     string
		KableClientSecret string
		BaseUrl           string
		Debug             bool
		MaxQueueSize      int
		HTTPClient        *http.Client
	}

	Kable struct {
		eventsApi       *EventsApi
		authenticateApi *AuthenticateApi
		options         *KableOptions
	}
)

func (options *KableOptions) String() string {
	return fmt.Sprintf("KableClientId=%s, KableClientSecret=%s, Debug=%t, MaxQueueSize=%d, BaseUrl=%s", options.KableClientId, options.KableClientSecret, options.Debug, options.MaxQueueSize, options.BaseUrl)
}

func New(options *KableOptions) *Kable {
	log.Println("[KABLE] Initializing Kable")

	apiConf := openapi.NewConfiguration()
	authConf := auth.NewConfiguration()

	if options == nil {
		log.Fatal("[KABLE] Invalid Kable configuration, options not found")
	} else {
		if options.Debug {
			log.Println("[KABLE] Kable configuration:", options)
		}
		baseUrl, err := url.Parse(options.BaseUrl)
		if err != nil {
			log.Fatal("[KABLE] Failed to initialize Kable: BaseUrl is invalid", err)
		}
		apiConf.Scheme = baseUrl.Scheme
		apiConf.Host = baseUrl.Host
		authConf.Scheme = baseUrl.Scheme
		authConf.Host = baseUrl.Host
	}

	apiConf.AddDefaultHeader("Content-Type", "application/json")
	authConf.AddDefaultHeader("Content-Type", "application/json")

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
		log.Println("[KABLE] Kable initialized successfully")
	}

	return &Kable{
		eventsApi:       NewEventsApi(apiClient, options),
		authenticateApi: NewAuthenticateApi(authClient, options),
		options:         options,
	}
}

func (k *Kable) Record(events ...Event) {
	if k.options.Debug {
		log.Printf("[KABLE] Received %d event(s) to record", len(events))
	}
	k.eventsApi.EnqueueEvent(events...)
}
