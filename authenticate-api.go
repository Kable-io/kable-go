package kable

import (
	"context"
	"github.com/Kable-io/kable-go/internal/sdk"
	"log"
)

type AuthenticateApi struct {
	api     *sdk.APIClient
	options *KableOptions
	ctx     context.Context
}

func (a *AuthenticateApi) TestAuthentication() error {
	req := a.api.AuthenticateApi.Authenticate(a.ctx)
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

func NewAuthenticateApi(apiClient *sdk.APIClient, options *KableOptions) *AuthenticateApi {
	return &AuthenticateApi{
		api:     apiClient,
		options: options,
		ctx:     context.Background(),
	}
}
