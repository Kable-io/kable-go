# \DimensionsApi

All URIs are relative to *https://live.kable.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllDimensions**](DimensionsApi.md#GetAllDimensions) | **Get** /dimensions | get all dimensions
[**GetDimension**](DimensionsApi.md#GetDimension) | **Get** /dimensions/{dimensionId} | get dimension



## GetAllDimensions

> []Dimension GetAllDimensions(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

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
    kableClientId := "kci_3c90e9ac92c64f31ae8ed84d21e18740" // string | Your client ID, found in the dashboard of your Kable account.
    kableClientSecret := "sk_test.jI92Cbu0.XeHWdYM1VTLy4oLtGMw8wrmpt5q9d04n" // string | Your `LIVE` or `TEST` secret key. Customers exist across all environments, so it does not matter which environment you use to fetch customers.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DimensionsApi.GetAllDimensions(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DimensionsApi.GetAllDimensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllDimensions`: []Dimension
    fmt.Fprintf(os.Stdout, "Response from `DimensionsApi.GetAllDimensions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDimensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to fetch customers. | 

### Return type

[**[]Dimension**](Dimension.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDimension

> Dimension GetDimension(ctx, dimensionId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

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
    kableClientId := "kci_3c90e9ac92c64f31ae8ed84d21e18740" // string | Your client ID, found in the dashboard of your Kable account.
    kableClientSecret := "sk_test.jI92Cbu0.XeHWdYM1VTLy4oLtGMw8wrmpt5q9d04n" // string | Your `LIVE` or `TEST` secret key. Customers exist across all environments, so it does not matter which environment you use to fetch customers.
    dimensionId := "userId" // string | The identifier for the dimension. You can pass in *either* the `dimensionId` (as defined by Kable) or the `key` (as defined by you).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DimensionsApi.GetDimension(context.Background(), dimensionId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DimensionsApi.GetDimension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDimension`: Dimension
    fmt.Fprintf(os.Stdout, "Response from `DimensionsApi.GetDimension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dimensionId** | **string** | The identifier for the dimension. You can pass in *either* the &#x60;dimensionId&#x60; (as defined by Kable) or the &#x60;key&#x60; (as defined by you). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDimensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to fetch customers. | 


### Return type

[**Dimension**](Dimension.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

