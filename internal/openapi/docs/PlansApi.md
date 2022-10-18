# \PlansApi

All URIs are relative to *https://live.kable.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePlan**](PlansApi.md#CreatePlan) | **Post** /plans/create | create a plan
[**DeletePlan**](PlansApi.md#DeletePlan) | **Post** /plans/{planId}/delete | delete a plan
[**GetAllPlans**](PlansApi.md#GetAllPlans) | **Get** /plans | get all plans
[**GetPlan**](PlansApi.md#GetPlan) | **Get** /plans/{planId} | get plan
[**UpdatePlan**](PlansApi.md#UpdatePlan) | **Post** /plans/{planId}/update | update a plan



## CreatePlan

> Plan CreatePlan(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreatePlanRequest(createPlanRequest).Execute()

create a plan



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
    kableClientSecret := "sk_test.jI92Cbu0.XeHWdYM1VTLy4oLtGMw8wrmpt5q9d04n" // string | Your `LIVE` or `TEST` secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. Each customer will have separate keys for `LIVE` and `TEST` environments of your API.
    createPlanRequest := *openapiclient.NewCreatePlanRequest("Requests Usage Plan", "USAGE", "MONTH", *openapiclient.NewPrice("USD", []openapiclient.PriceTier{*openapiclient.NewPriceTier("USD", float32(999.99))})) // CreatePlanRequest | Information about the plan to create. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.CreatePlan(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreatePlanRequest(createPlanRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.CreatePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlan`: Plan
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.CreatePlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. Each customer will have separate keys for &#x60;LIVE&#x60; and &#x60;TEST&#x60; environments of your API. | 
 **createPlanRequest** | [**CreatePlanRequest**](CreatePlanRequest.md) | Information about the plan to create. | 

### Return type

[**Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlan

> DeletePlan(ctx, planId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

delete a plan



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
    kableClientSecret := "sk_test.jI92Cbu0.XeHWdYM1VTLy4oLtGMw8wrmpt5q9d04n" // string | Your `LIVE` or `TEST` secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. Each customer will have separate keys for `LIVE` and `TEST` environments of your API.
    planId := "pln_a23e148584d0463482c961b7f62f824c" // string | The identifier for the plan. You can pass in *either* the `planId` (as defined by Kable) or the `externalId` (as defined by you).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.DeletePlan(context.Background(), planId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.DeletePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | The identifier for the plan. You can pass in *either* the &#x60;planId&#x60; (as defined by Kable) or the &#x60;externalId&#x60; (as defined by you). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. Each customer will have separate keys for &#x60;LIVE&#x60; and &#x60;TEST&#x60; environments of your API. | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPlans

> []Plan GetAllPlans(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get all plans



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
    resp, r, err := apiClient.PlansApi.GetAllPlans(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.GetAllPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPlans`: []Plan
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.GetAllPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to fetch customers. | 

### Return type

[**[]Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlan

> Plan GetPlan(ctx, planId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get plan



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
    planId := "pln_a23e148584d0463482c961b7f62f824c" // string | The identifier for the plan. You can pass in *either* the `planId` (as defined by Kable) or the `externalId` (as defined by you).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.GetPlan(context.Background(), planId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.GetPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlan`: Plan
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.GetPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | The identifier for the plan. You can pass in *either* the &#x60;planId&#x60; (as defined by Kable) or the &#x60;externalId&#x60; (as defined by you). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to fetch customers. | 


### Return type

[**Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlan

> Plan UpdatePlan(ctx, planId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UpdatePlanRequest(updatePlanRequest).Execute()

update a plan



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
    kableClientSecret := "sk_test.jI92Cbu0.XeHWdYM1VTLy4oLtGMw8wrmpt5q9d04n" // string | Your `LIVE` or `TEST` secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. Each customer will have separate keys for `LIVE` and `TEST` environments of your API.
    planId := "pln_a23e148584d0463482c961b7f62f824c" // string | The identifier for the plan. You can pass in *either* the `planId` (as defined by Kable) or the `externalId` (as defined by you).
    updatePlanRequest := *openapiclient.NewUpdatePlanRequest() // UpdatePlanRequest | Information about the plan to update. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.UpdatePlan(context.Background(), planId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UpdatePlanRequest(updatePlanRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.UpdatePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePlan`: Plan
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.UpdatePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | The identifier for the plan. You can pass in *either* the &#x60;planId&#x60; (as defined by Kable) or the &#x60;externalId&#x60; (as defined by you). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. Each customer will have separate keys for &#x60;LIVE&#x60; and &#x60;TEST&#x60; environments of your API. | 

 **updatePlanRequest** | [**UpdatePlanRequest**](UpdatePlanRequest.md) | Information about the plan to update. | 

### Return type

[**Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

