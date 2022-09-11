package kable

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Kable-io/kable-go/internal/auth"
	"github.com/Kable-io/kable-go/internal/openapi"
	"github.com/Kable-io/kable-go/internal/version"
)

type (
	KableOptions struct {
		KableClientId     string
		KableClientSecret string
		Debug             bool
		MaxQueueSize      int
		HTTPClient        *http.Client
	}

	Kable struct {
		eventsApi *EventsApi
		options   *KableOptions
	}
)

func New(options *KableOptions) *Kable {
	log.Println("[KABLE] Initializing Kable")

	if options == nil {
		log.Fatal("[KABLE] Invalid Kable configuration, options not found")
	}

	if options.MaxQueueSize > 1000 {
		options.MaxQueueSize = 1000
	}

	if options.MaxQueueSize < 1 {
		options.MaxQueueSize = 1
	}

	log.Printf("[KABLE] Starting Kable with MaxQueueSize=%d, Debug=%t", options.MaxQueueSize, options.Debug)

	apiConf := openapi.NewConfiguration()
	authConf := auth.NewConfiguration()

	var isLive bool = strings.HasPrefix(options.KableClientSecret, "sk_live")
	if isLive {
		apiConf.Servers = apiConf.Servers[0:1]
		authConf.Servers = authConf.Servers[0:1]
	} else {
		apiConf.Servers = apiConf.Servers[1:2]
		authConf.Servers = authConf.Servers[1:2]
	}

	apiConf.AddDefaultHeader("Content-Type", "application/json")
	authConf.AddDefaultHeader("Content-Type", "application/json")

	apiConf.UserAgent = fmt.Sprintf("kable-libs/%s/go", version.Version)
	authConf.UserAgent = fmt.Sprintf("kable-libs/%s/go", version.Version)

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
		eventsApi: NewEventsApi(apiClient, options),
		options:   options,
	}
}

func (k *Kable) Record(events ...Event) {
	if k.options.Debug {
		log.Printf("[KABLE] Received %d event(s) to record", len(events))
	}
	k.eventsApi.EnqueueEvent(events...)
}
