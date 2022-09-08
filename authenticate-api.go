package kable

import (
	"log"

	"github.com/Kable-io/kable-go/internal/auth"
)

type AuthenticateApi struct {
	api     *auth.APIClient
	options *KableOptions
}

func (a *AuthenticateApi) Authenticate() error {
	req := a.api.AuthenticateApi.Authenticate(*a.options.Context)
	req = req.KableClientId(a.options.KableClientId)
	req = req.KableClientSecret(a.options.KableClientSecret)
	req = req.XClientId(a.options.KableClientId)

	_, err := a.api.AuthenticateApi.AuthenticateExecute(req)
	if err != nil {
		log.Fatal(err)
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
