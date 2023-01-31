# \EntitlementsApi

All URIs are relative to *https://live.kable.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEntitlement**](EntitlementsApi.md#CreateEntitlement) | **Post** /api/v1/entitlements/create | create entitlement
[**DeleteEntitlement**](EntitlementsApi.md#DeleteEntitlement) | **Post** /api/v1/entitlements/{entitlementId}/delete | delete entitlement
[**GetEntitlement**](EntitlementsApi.md#GetEntitlement) | **Get** /api/v1/entitlements/{entitlementId} | get entitlement
[**GetEntitlements**](EntitlementsApi.md#GetEntitlements) | **Get** /api/v1/entitlements | get all entitlements
[**UpdateEntitlement**](EntitlementsApi.md#UpdateEntitlement) | **Post** /api/v1/entitlements/{entitlementId}/update | update entitlement



## CreateEntitlement

> EntitlementResponseDto CreateEntitlement(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateEntitlementDto(createEntitlementDto).Execute()

create entitlement



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
    createEntitlementDto := *openapiclient.NewCreateEntitlementDto("Premium Support", "yourcompany_entitlement789") // CreateEntitlementDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.CreateEntitlement(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateEntitlementDto(createEntitlementDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.CreateEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEntitlement`: EntitlementResponseDto
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.CreateEntitlement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 
 **createEntitlementDto** | [**CreateEntitlementDto**](CreateEntitlementDto.md) |  | 

### Return type

[**EntitlementResponseDto**](EntitlementResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEntitlement

> DeleteEntitlement(ctx, entitlementId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

delete entitlement



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
    entitlementId := "ent_a44e148584d0463482c961b7f62f825d" // string | The entitlement ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.DeleteEntitlement(context.Background(), entitlementId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.DeleteEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entitlementId** | **string** | The entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntitlementRequest struct via the builder pattern


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


## GetEntitlement

> EntitlementResponseDto GetEntitlement(ctx, entitlementId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get entitlement



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
    entitlementId := "ent_a44e148584d0463482c961b7f62f825d" // string | The entitlement ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.GetEntitlement(context.Background(), entitlementId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.GetEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntitlement`: EntitlementResponseDto
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.GetEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entitlementId** | **string** | The entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 


### Return type

[**EntitlementResponseDto**](EntitlementResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlements

> []EntitlementResponseDto GetEntitlements(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get all entitlements



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
    resp, r, err := apiClient.EntitlementsApi.GetEntitlements(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.GetEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntitlements`: []EntitlementResponseDto
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.GetEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

### Return type

[**[]EntitlementResponseDto**](EntitlementResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntitlement

> EntitlementResponseDto UpdateEntitlement(ctx, entitlementId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UpdateEntitlementDto(updateEntitlementDto).Execute()

update entitlement



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
    entitlementId := "ent_a44e148584d0463482c961b7f62f825d" // string | The entitlement ID
    updateEntitlementDto := *openapiclient.NewUpdateEntitlementDto() // UpdateEntitlementDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.UpdateEntitlement(context.Background(), entitlementId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UpdateEntitlementDto(updateEntitlementDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.UpdateEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntitlement`: EntitlementResponseDto
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.UpdateEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entitlementId** | **string** | The entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

 **updateEntitlementDto** | [**UpdateEntitlementDto**](UpdateEntitlementDto.md) |  | 

### Return type

[**EntitlementResponseDto**](EntitlementResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

