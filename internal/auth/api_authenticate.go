/*
Kable SDK

Used for generation of SDK 

API version: 0.0.1
Contact: contact@kable.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auth

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// AuthenticateApiService AuthenticateApi service
type AuthenticateApiService service

type ApiAuthenticateRequest struct {
	ctx context.Context
	ApiService *AuthenticateApiService
	kableClientId *string
	kableClientSecret *string
	xClientId *string
}

// Your client ID, found in the dashboard of your Kable account.
func (r ApiAuthenticateRequest) KableClientId(kableClientId string) ApiAuthenticateRequest {
	r.kableClientId = &kableClientId
	return r
}

// Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to fetch customers.
func (r ApiAuthenticateRequest) KableClientSecret(kableClientSecret string) ApiAuthenticateRequest {
	r.kableClientSecret = &kableClientSecret
	return r
}

// Your client ID or the client ID of the customer you are fetching.
func (r ApiAuthenticateRequest) XClientId(xClientId string) ApiAuthenticateRequest {
	r.xClientId = &xClientId
	return r
}

func (r ApiAuthenticateRequest) Execute() (*http.Response, error) {
	return r.ApiService.AuthenticateExecute(r)
}

/*
Authenticate test authentication

Test authentication


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthenticateRequest
*/
func (a *AuthenticateApiService) Authenticate(ctx context.Context) ApiAuthenticateRequest {
	return ApiAuthenticateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AuthenticateApiService) AuthenticateExecute(r ApiAuthenticateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticateApiService.Authenticate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/authenticate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.kableClientId == nil {
		return nil, reportError("kableClientId is required and must be specified")
	}
	if r.kableClientSecret == nil {
		return nil, reportError("kableClientSecret is required and must be specified")
	}
	if r.xClientId == nil {
		return nil, reportError("xClientId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Kable-Client-Id"] = parameterToString(*r.kableClientId, "")
	localVarHeaderParams["Kable-Client-Secret"] = parameterToString(*r.kableClientSecret, "")
	localVarHeaderParams["X-Client-Id"] = parameterToString(*r.xClientId, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}