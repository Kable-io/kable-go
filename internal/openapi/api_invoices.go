/*
Kable API

The Kable API allows developers to manage customers, plans, and usage data for their API.  ## API Host and Environments  Kable is accessible in Live and Test environments for authentication and metering of client API requests. You will have separate API keys to access each environment.  You should only use Kable's Live environment for your own production data. All other configured environments should use Kable's Test environment.  ## API Protocols and Headers  All requests to the Kable API are made over HTTPS TLS v1.2+ to ensure security. Calls made over HTTP will fail. Any requests without proper authentication will also fail.  The Kable API uses standard JSON for requests and responses. Be sure to set both the `Content-Type` and `Accept` headers on each request to `application/json`.  Each Kable API response includes a `requestId` as the `X-REQUEST-ID` response header. The `requestId` is included on most responses regardless whether the API request succeeded or failed. You can use this `requestId` to help with debugging or when contacting support regarding a specific API call.  ## API Versioning  All Kable endpoints are versioned. After the host, each API can be found at `/api/vX/...` where X is the API version.  We strive to ensure that changes to the Kable API are backward compatible. Sometimes, though, we must break from older design paradigms to make the product better. When this happens, a new version of the API is released.  The current version of the Kable API is `v1`.  

API version: 1.0.0
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


// InvoicesApiService InvoicesApi service
type InvoicesApiService service

type ApiDownloadInvoiceRequest struct {
	ctx context.Context
	ApiService *InvoicesApiService
	kableClientId *string
	kableClientSecret *string
	invoiceId string
	format *string
}

// Your Kable client ID, found in the dashboard of your Kable account.
func (r ApiDownloadInvoiceRequest) KableClientId(kableClientId string) ApiDownloadInvoiceRequest {
	r.kableClientId = &kableClientId
	return r
}

// Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key.
func (r ApiDownloadInvoiceRequest) KableClientSecret(kableClientSecret string) ApiDownloadInvoiceRequest {
	r.kableClientSecret = &kableClientSecret
	return r
}

// The downloaded file format
func (r ApiDownloadInvoiceRequest) Format(format string) ApiDownloadInvoiceRequest {
	r.format = &format
	return r
}

func (r ApiDownloadInvoiceRequest) Execute() (*http.Response, error) {
	return r.ApiService.DownloadInvoiceExecute(r)
}

/*
DownloadInvoice download invoice

Download an invoice

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invoiceId The invoice ID
 @return ApiDownloadInvoiceRequest
*/
func (a *InvoicesApiService) DownloadInvoice(ctx context.Context, invoiceId string) ApiDownloadInvoiceRequest {
	return ApiDownloadInvoiceRequest{
		ApiService: a,
		ctx: ctx,
		invoiceId: invoiceId,
	}
}

// Execute executes the request
func (a *InvoicesApiService) DownloadInvoiceExecute(r ApiDownloadInvoiceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.DownloadInvoice")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/invoices/{invoiceId}/download"
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(parameterToString(r.invoiceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.kableClientId == nil {
		return nil, reportError("kableClientId is required and must be specified")
	}
	if r.kableClientSecret == nil {
		return nil, reportError("kableClientSecret is required and must be specified")
	}
	if r.format == nil {
		return nil, reportError("format is required and must be specified")
	}

	localVarQueryParams.Add("format", parameterToString(*r.format, ""))
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorMessage401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorMessage500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetInvoiceRequest struct {
	ctx context.Context
	ApiService *InvoicesApiService
	kableClientId *string
	kableClientSecret *string
	invoiceId string
}

// Your Kable client ID, found in the dashboard of your Kable account.
func (r ApiGetInvoiceRequest) KableClientId(kableClientId string) ApiGetInvoiceRequest {
	r.kableClientId = &kableClientId
	return r
}

// Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key.
func (r ApiGetInvoiceRequest) KableClientSecret(kableClientSecret string) ApiGetInvoiceRequest {
	r.kableClientSecret = &kableClientSecret
	return r
}

func (r ApiGetInvoiceRequest) Execute() (*InvoiceResponseDto, *http.Response, error) {
	return r.ApiService.GetInvoiceExecute(r)
}

/*
GetInvoice get invoice

Retrieve an invoice

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invoiceId The invoice ID
 @return ApiGetInvoiceRequest
*/
func (a *InvoicesApiService) GetInvoice(ctx context.Context, invoiceId string) ApiGetInvoiceRequest {
	return ApiGetInvoiceRequest{
		ApiService: a,
		ctx: ctx,
		invoiceId: invoiceId,
	}
}

// Execute executes the request
//  @return InvoiceResponseDto
func (a *InvoicesApiService) GetInvoiceExecute(r ApiGetInvoiceRequest) (*InvoiceResponseDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceResponseDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.GetInvoice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/invoices/{invoiceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(parameterToString(r.invoiceId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorMessage401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorMessage500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiGetInvoicesRequest struct {
	ctx context.Context
	ApiService *InvoicesApiService
	kableClientId *string
	kableClientSecret *string
	customerId *string
	clientId *string
}

// Your Kable client ID, found in the dashboard of your Kable account.
func (r ApiGetInvoicesRequest) KableClientId(kableClientId string) ApiGetInvoicesRequest {
	r.kableClientId = &kableClientId
	return r
}

// Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key.
func (r ApiGetInvoicesRequest) KableClientSecret(kableClientSecret string) ApiGetInvoicesRequest {
	r.kableClientSecret = &kableClientSecret
	return r
}

// A Kable-defined identifier for the customer.
func (r ApiGetInvoicesRequest) CustomerId(customerId string) ApiGetInvoicesRequest {
	r.customerId = &customerId
	return r
}

// A unique identifier for the customer, defined by you.
func (r ApiGetInvoicesRequest) ClientId(clientId string) ApiGetInvoicesRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetInvoicesRequest) Execute() ([]InvoiceResponseDto, *http.Response, error) {
	return r.ApiService.GetInvoicesExecute(r)
}

/*
GetInvoices get all invoices

Retrieve all invoices.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetInvoicesRequest
*/
func (a *InvoicesApiService) GetInvoices(ctx context.Context) ApiGetInvoicesRequest {
	return ApiGetInvoicesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []InvoiceResponseDto
func (a *InvoicesApiService) GetInvoicesExecute(r ApiGetInvoicesRequest) ([]InvoiceResponseDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InvoiceResponseDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.GetInvoices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/invoices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.kableClientId == nil {
		return localVarReturnValue, nil, reportError("kableClientId is required and must be specified")
	}
	if r.kableClientSecret == nil {
		return localVarReturnValue, nil, reportError("kableClientSecret is required and must be specified")
	}

	if r.customerId != nil {
		localVarQueryParams.Add("customerId", parameterToString(*r.customerId, ""))
	}
	if r.clientId != nil {
		localVarQueryParams.Add("clientId", parameterToString(*r.clientId, ""))
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorMessage401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorMessage500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
