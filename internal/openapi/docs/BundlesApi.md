# \BundlesApi

All URIs are relative to *https://live.kable.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBundle**](BundlesApi.md#CreateBundle) | **Post** /api/v1/bundles/create | create bundle
[**DeleteBundle**](BundlesApi.md#DeleteBundle) | **Post** /api/v1/bundles/{bundleId}/delete | delete bundle
[**GetBundle**](BundlesApi.md#GetBundle) | **Get** /api/v1/bundles/{bundleId} | get bundle
[**GetBundles**](BundlesApi.md#GetBundles) | **Get** /api/v1/bundles | get all bundles
[**UpdateBundle**](BundlesApi.md#UpdateBundle) | **Post** /api/v1/bundles/{bundleId}/update | update bundle



## CreateBundle

> []BundleResponseDto CreateBundle(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateBundleRequestDto(createBundleRequestDto).Execute()

create bundle



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
    createBundleRequestDto := *openapiclient.NewCreateBundleRequestDto("BundleId_example", "Name_example") // CreateBundleRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundlesApi.CreateBundle(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateBundleRequestDto(createBundleRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesApi.CreateBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBundle`: []BundleResponseDto
    fmt.Fprintf(os.Stdout, "Response from `BundlesApi.CreateBundle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 
 **createBundleRequestDto** | [**CreateBundleRequestDto**](CreateBundleRequestDto.md) |  | 

### Return type

[**[]BundleResponseDto**](BundleResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBundle

> DeleteBundle(ctx, bundleId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

delete bundle



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
    bundleId := "bun_a23e148584d0463482c961b7f62f824c" // string | The bundle ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundlesApi.DeleteBundle(context.Background(), bundleId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesApi.DeleteBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The bundle ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBundleRequest struct via the builder pattern


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


## GetBundle

> BundleResponseDto GetBundle(ctx, bundleId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get bundle



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
    bundleId := "bun_a23e148584d0463482c961b7f62f824c" // string | The bundle ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundlesApi.GetBundle(context.Background(), bundleId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesApi.GetBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBundle`: BundleResponseDto
    fmt.Fprintf(os.Stdout, "Response from `BundlesApi.GetBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The bundle ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 


### Return type

[**BundleResponseDto**](BundleResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBundles

> []BundleResponseDto GetBundles(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get all bundles



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
    resp, r, err := apiClient.BundlesApi.GetBundles(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesApi.GetBundles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBundles`: []BundleResponseDto
    fmt.Fprintf(os.Stdout, "Response from `BundlesApi.GetBundles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

### Return type

[**[]BundleResponseDto**](BundleResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBundle

> BundleResponseDto UpdateBundle(ctx, bundleId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UpdateBundleRequestDto(updateBundleRequestDto).Execute()

update bundle



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
    bundleId := "bun_a23e148584d0463482c961b7f62f824c" // string | The bundle ID
    updateBundleRequestDto := *openapiclient.NewUpdateBundleRequestDto() // UpdateBundleRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundlesApi.UpdateBundle(context.Background(), bundleId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UpdateBundleRequestDto(updateBundleRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesApi.UpdateBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBundle`: BundleResponseDto
    fmt.Fprintf(os.Stdout, "Response from `BundlesApi.UpdateBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The bundle ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

 **updateBundleRequestDto** | [**UpdateBundleRequestDto**](UpdateBundleRequestDto.md) |  | 

### Return type

[**BundleResponseDto**](BundleResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

