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


// DimensionsApiService DimensionsApi service
type DimensionsApiService service

type ApiCreateDimensionRequest struct {
	ctx context.Context
	ApiService *DimensionsApiService
	kableClientId *string
	kableClientSecret *string
	createDimensionRequestDto *CreateDimensionRequestDto
}

// Your Kable client ID, found in the dashboard of your Kable account.
func (r ApiCreateDimensionRequest) KableClientId(kableClientId string) ApiCreateDimensionRequest {
	r.kableClientId = &kableClientId
	return r
}

// Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key.
func (r ApiCreateDimensionRequest) KableClientSecret(kableClientSecret string) ApiCreateDimensionRequest {
	r.kableClientSecret = &kableClientSecret
	return r
}

func (r ApiCreateDimensionRequest) CreateDimensionRequestDto(createDimensionRequestDto CreateDimensionRequestDto) ApiCreateDimensionRequest {
	r.createDimensionRequestDto = &createDimensionRequestDto
	return r
}

func (r ApiCreateDimensionRequest) Execute() (*DimensionResponseDto, *http.Response, error) {
	return r.ApiService.CreateDimensionExecute(r)
}

/*
CreateDimension create dimension

Create a dimension.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDimensionRequest
*/
func (a *DimensionsApiService) CreateDimension(ctx context.Context) ApiCreateDimensionRequest {
	return ApiCreateDimensionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DimensionResponseDto
func (a *DimensionsApiService) CreateDimensionExecute(r ApiCreateDimensionRequest) (*DimensionResponseDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DimensionResponseDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DimensionsApiService.CreateDimension")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/dimensions/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.kableClientId == nil {
		return localVarReturnValue, nil, reportError("kableClientId is required and must be specified")
	}
	if r.kableClientSecret == nil {
		return localVarReturnValue, nil, reportError("kableClientSecret is required and must be specified")
	}
	if r.createDimensionRequestDto == nil {
		return localVarReturnValue, nil, reportError("createDimensionRequestDto is required and must be specified")
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
	localVarPostBody = r.createDimensionRequestDto
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

type ApiDeleteDimensionRequest struct {
	ctx context.Context
	ApiService *DimensionsApiService
	kableClientId *string
	kableClientSecret *string
	dimensionId string
}

// Your Kable client ID, found in the dashboard of your Kable account.
func (r ApiDeleteDimensionRequest) KableClientId(kableClientId string) ApiDeleteDimensionRequest {
	r.kableClientId = &kableClientId
	return r
}

// Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key.
func (r ApiDeleteDimensionRequest) KableClientSecret(kableClientSecret string) ApiDeleteDimensionRequest {
	r.kableClientSecret = &kableClientSecret
	return r
}

func (r ApiDeleteDimensionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDimensionExecute(r)
}

/*
DeleteDimension delete dimension

Delete a dimension.

 You cannot delete a dimension if it is in use by a plan.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dimensionId The dimension ID
 @return ApiDeleteDimensionRequest
*/
func (a *DimensionsApiService) DeleteDimension(ctx context.Context, dimensionId string) ApiDeleteDimensionRequest {
	return ApiDeleteDimensionRequest{
		ApiService: a,
		ctx: ctx,
		dimensionId: dimensionId,
	}
}

// Execute executes the request
func (a *DimensionsApiService) DeleteDimensionExecute(r ApiDeleteDimensionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DimensionsApiService.DeleteDimension")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/dimensions/{dimensionId}/delete"
	localVarPath = strings.Replace(localVarPath, "{"+"dimensionId"+"}", url.PathEscape(parameterToString(r.dimensionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.kableClientId == nil {
		return nil, reportError("kableClientId is required and must be specified")
	}
	if r.kableClientSecret == nil {
		return nil, reportError("kableClientSecret is required and must be specified")
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

type ApiGetDimensionRequest struct {
	ctx context.Context
	ApiService *DimensionsApiService
	kableClientId *string
	kableClientSecret *string
	dimensionId string
}

// Your Kable client ID, found in the dashboard of your Kable account.
func (r ApiGetDimensionRequest) KableClientId(kableClientId string) ApiGetDimensionRequest {
	r.kableClientId = &kableClientId
	return r
}

// Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key.
func (r ApiGetDimensionRequest) KableClientSecret(kableClientSecret string) ApiGetDimensionRequest {
	r.kableClientSecret = &kableClientSecret
	return r
}

func (r ApiGetDimensionRequest) Execute() (*DimensionResponseDto, *http.Response, error) {
	return r.ApiService.GetDimensionExecute(r)
}

/*
GetDimension get dimension

Retrieve a dimension.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dimensionId The dimension ID
 @return ApiGetDimensionRequest
*/
func (a *DimensionsApiService) GetDimension(ctx context.Context, dimensionId string) ApiGetDimensionRequest {
	return ApiGetDimensionRequest{
		ApiService: a,
		ctx: ctx,
		dimensionId: dimensionId,
	}
}

// Execute executes the request
//  @return DimensionResponseDto
func (a *DimensionsApiService) GetDimensionExecute(r ApiGetDimensionRequest) (*DimensionResponseDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DimensionResponseDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DimensionsApiService.GetDimension")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/dimensions/{dimensionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"dimensionId"+"}", url.PathEscape(parameterToString(r.dimensionId, "")), -1)

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

type ApiGetDimensionsRequest struct {
	ctx context.Context
	ApiService *DimensionsApiService
	kableClientId *string
	kableClientSecret *string
}

// Your Kable client ID, found in the dashboard of your Kable account.
func (r ApiGetDimensionsRequest) KableClientId(kableClientId string) ApiGetDimensionsRequest {
	r.kableClientId = &kableClientId
	return r
}

// Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key.
func (r ApiGetDimensionsRequest) KableClientSecret(kableClientSecret string) ApiGetDimensionsRequest {
	r.kableClientSecret = &kableClientSecret
	return r
}

func (r ApiGetDimensionsRequest) Execute() ([]DimensionResponseDto, *http.Response, error) {
	return r.ApiService.GetDimensionsExecute(r)
}

/*
GetDimensions get all dimensions

Retrieve all dimensions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDimensionsRequest
*/
func (a *DimensionsApiService) GetDimensions(ctx context.Context) ApiGetDimensionsRequest {
	return ApiGetDimensionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []DimensionResponseDto
func (a *DimensionsApiService) GetDimensionsExecute(r ApiGetDimensionsRequest) ([]DimensionResponseDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []DimensionResponseDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DimensionsApiService.GetDimensions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/dimensions"

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

type ApiUpdateDimensionRequest struct {
	ctx context.Context
	ApiService *DimensionsApiService
	kableClientId *string
	kableClientSecret *string
	dimensionId string
	updateDimensionRequestDto *UpdateDimensionRequestDto
}

// Your Kable client ID, found in the dashboard of your Kable account.
func (r ApiUpdateDimensionRequest) KableClientId(kableClientId string) ApiUpdateDimensionRequest {
	r.kableClientId = &kableClientId
	return r
}

// Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key.
func (r ApiUpdateDimensionRequest) KableClientSecret(kableClientSecret string) ApiUpdateDimensionRequest {
	r.kableClientSecret = &kableClientSecret
	return r
}

func (r ApiUpdateDimensionRequest) UpdateDimensionRequestDto(updateDimensionRequestDto UpdateDimensionRequestDto) ApiUpdateDimensionRequest {
	r.updateDimensionRequestDto = &updateDimensionRequestDto
	return r
}

func (r ApiUpdateDimensionRequest) Execute() (*DimensionResponseDto, *http.Response, error) {
	return r.ApiService.UpdateDimensionExecute(r)
}

/*
UpdateDimension update dimension

Update a dimension.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dimensionId The dimension ID
 @return ApiUpdateDimensionRequest
*/
func (a *DimensionsApiService) UpdateDimension(ctx context.Context, dimensionId string) ApiUpdateDimensionRequest {
	return ApiUpdateDimensionRequest{
		ApiService: a,
		ctx: ctx,
		dimensionId: dimensionId,
	}
}

// Execute executes the request
//  @return DimensionResponseDto
func (a *DimensionsApiService) UpdateDimensionExecute(r ApiUpdateDimensionRequest) (*DimensionResponseDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DimensionResponseDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DimensionsApiService.UpdateDimension")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/dimensions/{dimensionId}/update"
	localVarPath = strings.Replace(localVarPath, "{"+"dimensionId"+"}", url.PathEscape(parameterToString(r.dimensionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.kableClientId == nil {
		return localVarReturnValue, nil, reportError("kableClientId is required and must be specified")
	}
	if r.kableClientSecret == nil {
		return localVarReturnValue, nil, reportError("kableClientSecret is required and must be specified")
	}
	if r.updateDimensionRequestDto == nil {
		return localVarReturnValue, nil, reportError("updateDimensionRequestDto is required and must be specified")
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
	localVarPostBody = r.updateDimensionRequestDto
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
