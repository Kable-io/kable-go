# \PlansApi

All URIs are relative to *https://live.kable.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllPlans**](PlansApi.md#GetAllPlans) | **Get** /plans | get all plans
[**GetPlan**](PlansApi.md#GetPlan) | **Get** /plans/{planId} | get plan



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
    planId := "userId" // string | The identifier for the plan.

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
**planId** | **string** | The identifier for the plan. | 

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

