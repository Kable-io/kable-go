package kable

import (
	"context"
	"log"

	"github.com/Kable-io/kable-go/v2/internal/auth"
)

type AuthenticateApi struct {
	api     *auth.APIClient
	options *KableOptions
}

func (a *AuthenticateApi) Authenticate() error {
	req := a.api.AuthenticateApi.Authenticate(context.Background())

	if a.options.KableClientId == "" {
		log.Fatal("[KABLE] Failed to initialize Kable: KableClientId not provided")
	}

	if a.options.KableClientSecret == "" {
		log.Fatal("[KABLE] Failed to initialize Kable: KableClientSecret not provided")
	}

	req = req.KableClientId(a.options.KableClientId)
	req = req.KableClientSecret(a.options.KableClientSecret)
	req = req.XClientId(a.options.KableClientId)

	if a.options.Debug {
		log.Printf("[KABLE] Authenticating with Kable at %s", a.api.GetConfig().Host)
	}

	_, err := a.api.AuthenticateApi.AuthenticateExecute(req)
	if err != nil {
		log.Fatal("[KABLE] Failed to initialize Kable: ", err)
		return err
	}

	return nil
}

func NewAuthenticateApi(apiClient *auth.APIClient, options *KableOptions) *AuthenticateApi {
	return &AuthenticateApi{
		api:     apiClient,
		options: options,
	}
}
