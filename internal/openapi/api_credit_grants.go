/*
Kable API

The Kable API allows developers to manage **customers** and **API keys** and record **events** about their API.   ## API Host and Environments  Kable is accessible in Live and Test environments for authentication and metering of client API requests. You will have separate API keys to access each environment.  You should only use Kable's Live environment for your own production data. All other configured environments should use Kable's Test environment.   ## API Protocols and Headers  All requests to the Kable API are made over HTTPS TLS v1.2+ to ensure security. Calls made over HTTP will fail. Any requests without proper authentication will also fail.  The Kable API uses standard JSON for requests and responses. Be sure to set both the `Content-Type` and `Accept` headers on each request to application/json.  Each Kable API response includes a `requestId` as the `X-REQUEST-ID` response header. The `requestId` is included regardless whether the API request succeeded or failed. You can use this requestId to help with debugging or when contacting support regarding a specific API call.   ## API Keys  There are two types of API keys on Kable.  ### Kable Keys Kable Keys are the keys you, the Kable customer, use to interact with Kable. These keys help us ensure that only you are interacting with Kable on your behalf. You can find your keys on the Company page of the dashboard after you sign up.  Kable Keys should be included in every request to the Kable API. You must provide your client ID as the `KABLE-CLIENT-ID` header and your secret key as the `KABLE-CLIENT-SECRET` header on each request to Kable. If you are using a language-specific Kable library, you will initialize the SDK using these keys.  ### Customer Keys Customer Keys are the keys your customers use to interact with your API. Customer Keys are authenticated by Kable when a customer makes a request to your API if you use Kable's authentication services. Customers must provide their client ID (defined as `clientId` when you create the customer) as the `X-CLIENT-ID` header and their secret key as the `X-API-KEY` header on each request to your API that Kable is to authenticate.   ## API Versioning  All Kable endpoints are versioned. After the host, each API can be found at `/api/vX/...` where X is the API version.  We strive to ensure that changes to the Kable API are backward compatible. Sometimes, though, we must break from older design paradigms to make the product better. When this happens, a new version of the API is released.  The current version of Kable is **v1**. 

API version: 1.2.1
Contact: contact@kable.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// CreditGrantsApiService CreditGrantsApi service
type CreditGrantsApiService service

type ApiCreateCreditGrantRequest struct {
	ctx context.Context
	ApiService *CreditGrantsApiService
	kableClientId *string
	kableClientSecret *string
	customerId string
	createCreditGrantRequest *CreateCreditGrantRequest
}

// Your client ID, found in the dashboard of your Kable account.
func (r ApiCreateCreditGrantRequest) KableClientId(kableClientId string) ApiCreateCreditGrantRequest {
	r.kableClientId = &kableClientId
	return r
}

// Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. Each customer will have separate keys for &#x60;LIVE&#x60; and &#x60;TEST&#x60; environments of your API.
func (r ApiCreateCreditGrantRequest) KableClientSecret(kableClientSecret string) ApiCreateCreditGrantRequest {
	r.kableClientSecret = &kableClientSecret
	return r
}

// Information about the credit grant to create.
func (r ApiCreateCreditGrantRequest) CreateCreditGrantRequest(createCreditGrantRequest CreateCreditGrantRequest) ApiCreateCreditGrantRequest {
	r.createCreditGrantRequest = &createCreditGrantRequest
	return r
}

func (r ApiCreateCreditGrantRequest) Execute() (*CreditGrant, *http.Response, error) {
	return r.ApiService.CreateCreditGrantExecute(r)
}

/*
CreateCreditGrant create a new credit grant for a customer

Add a credit grant to a customer.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerId The identifier for the customer. You can pass in *either* the `customerId` (as defined by Kable) or the `clientId` (as defined by you).
 @return ApiCreateCreditGrantRequest
*/
func (a *CreditGrantsApiService) CreateCreditGrant(ctx context.Context, customerId string) ApiCreateCreditGrantRequest {
	return ApiCreateCreditGrantRequest{
		ApiService: a,
		ctx: ctx,
		customerId: customerId,
	}
}

// Execute executes the request
//  @return CreditGrant
func (a *CreditGrantsApiService) CreateCreditGrantExecute(r ApiCreateCreditGrantRequest) (*CreditGrant, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreditGrant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditGrantsApiService.CreateCreditGrant")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customerId}/credits/create"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", url.PathEscape(parameterToString(r.customerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.kableClientId == nil {
		return localVarReturnValue, nil, reportError("kableClientId is required and must be specified")
	}
	if r.kableClientSecret == nil {
		return localVarReturnValue, nil, reportError("kableClientSecret is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Kable-Client-Id"] = parameterToString(*r.kableClientId, "")
	localVarHeaderParams["Kable-Client-Secret"] = parameterToString(*r.kableClientSecret, "")
	// body params
	localVarPostBody = r.createCreditGrantRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCustomerCreditBalanceRequest struct {
	ctx context.Context
	ApiService *CreditGrantsApiService
	kableClientId *string
	kableClientSecret *string
	customerId string
}

// Your client ID, found in the dashboard of your Kable account.
func (r ApiGetCustomerCreditBalanceRequest) KableClientId(kableClientId string) ApiGetCustomerCreditBalanceRequest {
	r.kableClientId = &kableClientId
	return r
}

// Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. Each customer will have separate keys for &#x60;LIVE&#x60; and &#x60;TEST&#x60; environments of your API.
func (r ApiGetCustomerCreditBalanceRequest) KableClientSecret(kableClientSecret string) ApiGetCustomerCreditBalanceRequest {
	r.kableClientSecret = &kableClientSecret
	return r
}

func (r ApiGetCustomerCreditBalanceRequest) Execute() (*GetCustomerCreditBalance200Response, *http.Response, error) {
	return r.ApiService.GetCustomerCreditBalanceExecute(r)
}

/*
GetCustomerCreditBalance get a customer's credit balance

Get a customer's available credit balance.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerId The identifier for the customer. You can pass in *either* the `customerId` (as defined by Kable) or the `clientId` (as defined by you).
 @return ApiGetCustomerCreditBalanceRequest
*/
func (a *CreditGrantsApiService) GetCustomerCreditBalance(ctx context.Context, customerId string) ApiGetCustomerCreditBalanceRequest {
	return ApiGetCustomerCreditBalanceRequest{
		ApiService: a,
		ctx: ctx,
		customerId: customerId,
	}
}

// Execute executes the request
//  @return GetCustomerCreditBalance200Response
func (a *CreditGrantsApiService) GetCustomerCreditBalanceExecute(r ApiGetCustomerCreditBalanceRequest) (*GetCustomerCreditBalance200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetCustomerCreditBalance200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditGrantsApiService.GetCustomerCreditBalance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customerId}/credits/balance"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", url.PathEscape(parameterToString(r.customerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.kableClientId == nil {
		return localVarReturnValue, nil, reportError("kableClientId is required and must be specified")
	}
	if r.kableClientSecret == nil {
		return localVarReturnValue, nil, reportError("kableClientSecret is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Kable-Client-Id"] = parameterToString(*r.kableClientId, "")
	localVarHeaderParams["Kable-Client-Secret"] = parameterToString(*r.kableClientSecret, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
