# \DimensionsApi

All URIs are relative to *https://live.kable.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDimension**](DimensionsApi.md#CreateDimension) | **Post** /api/v1/dimensions/create | create dimension
[**DeleteDimension**](DimensionsApi.md#DeleteDimension) | **Post** /api/v1/dimensions/{dimensionId}/delete | delete dimension
[**GetDimension**](DimensionsApi.md#GetDimension) | **Get** /api/v1/dimensions/{dimensionId} | get dimension
[**GetDimensions**](DimensionsApi.md#GetDimensions) | **Get** /api/v1/dimensions | get all dimensions
[**UpdateDimension**](DimensionsApi.md#UpdateDimension) | **Post** /api/v1/dimensions/{dimensionId}/update | update dimension



## CreateDimension

> DimensionResponseDto CreateDimension(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateDimensionRequestDto(createDimensionRequestDto).Execute()

create dimension



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    kableClientId := "kableClientId_example" // string | Your Kable client ID, found in the dashboard of your Kable account.
    kableClientSecret := "kableClientSecret_example" // string | Your `LIVE` or `TEST` secret key.
    createDimensionRequestDto := *openapiclient.NewCreateDimensionRequestDto("userId", "User") // CreateDimensionRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DimensionsApi.CreateDimension(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateDimensionRequestDto(createDimensionRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DimensionsApi.CreateDimension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDimension`: DimensionResponseDto
    fmt.Fprintf(os.Stdout, "Response from `DimensionsApi.CreateDimension`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDimensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 
 **createDimensionRequestDto** | [**CreateDimensionRequestDto**](CreateDimensionRequestDto.md) |  | 

### Return type

[**DimensionResponseDto**](DimensionResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDimension

> DeleteDimension(ctx, dimensionId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

delete dimension



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    kableClientId := "kableClientId_example" // string | Your Kable client ID, found in the dashboard of your Kable account.
    kableClientSecret := "kableClientSecret_example" // string | Your `LIVE` or `TEST` secret key.
    dimensionId := "dim_c5f32b8cd8934356b5167a3b2c6c6314" // string | The dimension ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DimensionsApi.DeleteDimension(context.Background(), dimensionId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DimensionsApi.DeleteDimension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dimensionId** | **string** | The dimension ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDimensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDimension

> DimensionResponseDto GetDimension(ctx, dimensionId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get dimension



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    kableClientId := "kableClientId_example" // string | Your Kable client ID, found in the dashboard of your Kable account.
    kableClientSecret := "kableClientSecret_example" // string | Your `LIVE` or `TEST` secret key.
    dimensionId := "dim_c5f32b8cd8934356b5167a3b2c6c6314" // string | The dimension ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DimensionsApi.GetDimension(context.Background(), dimensionId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DimensionsApi.GetDimension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDimension`: DimensionResponseDto
    fmt.Fprintf(os.Stdout, "Response from `DimensionsApi.GetDimension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dimensionId** | **string** | The dimension ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDimensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 


### Return type

[**DimensionResponseDto**](DimensionResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDimensions

> []DimensionResponseDto GetDimensions(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get all dimensions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    kableClientId := "kableClientId_example" // string | Your Kable client ID, found in the dashboard of your Kable account.
    kableClientSecret := "kableClientSecret_example" // string | Your `LIVE` or `TEST` secret key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DimensionsApi.GetDimensions(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DimensionsApi.GetDimensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDimensions`: []DimensionResponseDto
    fmt.Fprintf(os.Stdout, "Response from `DimensionsApi.GetDimensions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDimensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

### Return type

[**[]DimensionResponseDto**](DimensionResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDimension

> DimensionResponseDto UpdateDimension(ctx, dimensionId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UpdateDimensionRequestDto(updateDimensionRequestDto).Execute()

update dimension



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    kableClientId := "kableClientId_example" // string | Your Kable client ID, found in the dashboard of your Kable account.
    kableClientSecret := "kableClientSecret_example" // string | Your `LIVE` or `TEST` secret key.
    dimensionId := "dim_c5f32b8cd8934356b5167a3b2c6c6314" // string | The dimension ID
    updateDimensionRequestDto := *openapiclient.NewUpdateDimensionRequestDto("User") // UpdateDimensionRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DimensionsApi.UpdateDimension(context.Background(), dimensionId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UpdateDimensionRequestDto(updateDimensionRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DimensionsApi.UpdateDimension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDimension`: DimensionResponseDto
    fmt.Fprintf(os.Stdout, "Response from `DimensionsApi.UpdateDimension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dimensionId** | **string** | The dimension ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDimensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

 **updateDimensionRequestDto** | [**UpdateDimensionRequestDto**](UpdateDimensionRequestDto.md) |  | 

### Return type

[**DimensionResponseDto**](DimensionResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

