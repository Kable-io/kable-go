# \PlansApi

All URIs are relative to *https://live.kable.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPlanEntitlement**](PlansApi.md#AddPlanEntitlement) | **Post** /api/v1/plans/{planId}/entitlements/add | add plan entitlement
[**CreatePlan**](PlansApi.md#CreatePlan) | **Post** /api/v1/plans/create | create plan
[**DeletePlan**](PlansApi.md#DeletePlan) | **Post** /api/v1/plans/{planId}/delete | delete plan
[**GetAllPlanEntitlements**](PlansApi.md#GetAllPlanEntitlements) | **Get** /api/v1/plans/{planId}/entitlements | get all plan entitlements
[**GetPlan**](PlansApi.md#GetPlan) | **Get** /api/v1/plans/{planId} | get plan
[**GetPlanEntitlement**](PlansApi.md#GetPlanEntitlement) | **Get** /api/v1/plans/{planId}/entitlements/{entitlementId} | get plan entitlement
[**GetPlans**](PlansApi.md#GetPlans) | **Get** /api/v1/plans | get all plans
[**RemovePlanEntitlement**](PlansApi.md#RemovePlanEntitlement) | **Post** /api/v1/plans/{planId}/entitlements/remove | remove plan entitlement
[**UpdatePlan**](PlansApi.md#UpdatePlan) | **Post** /api/v1/plans/{planId}/update | update plan
[**UpdatePlanEntitlement**](PlansApi.md#UpdatePlanEntitlement) | **Post** /api/v1/plans/{planId}/entitlements/{entitlementId}/update | update plan entitlement



## AddPlanEntitlement

> PlanResponseDto AddPlanEntitlement(ctx, planId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).AddPlanEntitlementRequestDto(addPlanEntitlementRequestDto).Execute()

add plan entitlement



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
    planId := "pln_a23e148584d0463482c961b7f62f824c" // string | The plan ID
    addPlanEntitlementRequestDto := *openapiclient.NewAddPlanEntitlementRequestDto() // AddPlanEntitlementRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.AddPlanEntitlement(context.Background(), planId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).AddPlanEntitlementRequestDto(addPlanEntitlementRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.AddPlanEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPlanEntitlement`: PlanResponseDto
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.AddPlanEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | The plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPlanEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

 **addPlanEntitlementRequestDto** | [**AddPlanEntitlementRequestDto**](AddPlanEntitlementRequestDto.md) |  | 

### Return type

[**PlanResponseDto**](PlanResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlan

> PlanResponseDto CreatePlan(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreatePlanRequestDto(createPlanRequestDto).Execute()

create plan



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
    createPlanRequestDto := *openapiclient.NewCreatePlanRequestDto("PlanId_example", "Requests Usage Plan", "USAGE", *openapiclient.NewCreatePlanRequestDtoPrice("USD", []string{"Tiers_example"})) // CreatePlanRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.CreatePlan(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreatePlanRequestDto(createPlanRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.CreatePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlan`: PlanResponseDto
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.CreatePlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 
 **createPlanRequestDto** | [**CreatePlanRequestDto**](CreatePlanRequestDto.md) |  | 

### Return type

[**PlanResponseDto**](PlanResponseDto.md)

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

delete plan



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
    planId := "pln_a23e148584d0463482c961b7f62f824c" // string | The plan ID

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
**planId** | **string** | The plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanRequest struct via the builder pattern


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


## GetAllPlanEntitlements

> []PlanEntitlementResponseDto GetAllPlanEntitlements(ctx, planId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get all plan entitlements



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
    planId := "pln_a23e148584d0463482c961b7f62f824c" // string | The plan ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.GetAllPlanEntitlements(context.Background(), planId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.GetAllPlanEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPlanEntitlements`: []PlanEntitlementResponseDto
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.GetAllPlanEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | The plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPlanEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 


### Return type

[**[]PlanEntitlementResponseDto**](PlanEntitlementResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlan

> PlanResponseDto GetPlan(ctx, planId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

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
    kableClientId := "kableClientId_example" // string | Your Kable client ID, found in the dashboard of your Kable account.
    kableClientSecret := "kableClientSecret_example" // string | Your `LIVE` or `TEST` secret key.
    planId := "pln_a23e148584d0463482c961b7f62f824c" // string | The plan ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.GetPlan(context.Background(), planId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.GetPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlan`: PlanResponseDto
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.GetPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | The plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 


### Return type

[**PlanResponseDto**](PlanResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanEntitlement

> PlanEntitlementResponseDto GetPlanEntitlement(ctx, planId, entitlementId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get plan entitlement



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
    planId := "pln_a23e148584d0463482c961b7f62f824c" // string | The plan ID
    entitlementId := "ent_a44e148584d0463482c961b7f62f825d" // string | The entitlement ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.GetPlanEntitlement(context.Background(), planId, entitlementId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.GetPlanEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlanEntitlement`: PlanEntitlementResponseDto
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.GetPlanEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | The plan ID | 
**entitlementId** | **string** | The entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 



### Return type

[**PlanEntitlementResponseDto**](PlanEntitlementResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlans

> []PlanResponseDto GetPlans(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

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
    kableClientId := "kableClientId_example" // string | Your Kable client ID, found in the dashboard of your Kable account.
    kableClientSecret := "kableClientSecret_example" // string | Your `LIVE` or `TEST` secret key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.GetPlans(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.GetPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlans`: []PlanResponseDto
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.GetPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

### Return type

[**[]PlanResponseDto**](PlanResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePlanEntitlement

> PlanResponseDto RemovePlanEntitlement(ctx, planId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).RemovePlanEntitlementRequestDto(removePlanEntitlementRequestDto).Execute()

remove plan entitlement



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
    planId := "pln_a23e148584d0463482c961b7f62f824c" // string | The plan ID
    removePlanEntitlementRequestDto := *openapiclient.NewRemovePlanEntitlementRequestDto() // RemovePlanEntitlementRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.RemovePlanEntitlement(context.Background(), planId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).RemovePlanEntitlementRequestDto(removePlanEntitlementRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.RemovePlanEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemovePlanEntitlement`: PlanResponseDto
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.RemovePlanEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | The plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePlanEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

 **removePlanEntitlementRequestDto** | [**RemovePlanEntitlementRequestDto**](RemovePlanEntitlementRequestDto.md) |  | 

### Return type

[**PlanResponseDto**](PlanResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlan

> PlanResponseDto UpdatePlan(ctx, planId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UpdatePlanRequestDto(updatePlanRequestDto).Execute()

update plan



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
    planId := "pln_a23e148584d0463482c961b7f62f824c" // string | The plan ID
    updatePlanRequestDto := *openapiclient.NewUpdatePlanRequestDto() // UpdatePlanRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.UpdatePlan(context.Background(), planId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UpdatePlanRequestDto(updatePlanRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.UpdatePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePlan`: PlanResponseDto
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.UpdatePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | The plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

 **updatePlanRequestDto** | [**UpdatePlanRequestDto**](UpdatePlanRequestDto.md) |  | 

### Return type

[**PlanResponseDto**](PlanResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlanEntitlement

> PlanEntitlementResponseDto UpdatePlanEntitlement(ctx, planId, entitlementId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UpdatePlanEntitlementRequestDto(updatePlanEntitlementRequestDto).Execute()

update plan entitlement



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
    planId := "pln_a23e148584d0463482c961b7f62f824c" // string | The plan ID
    entitlementId := "ent_a44e148584d0463482c961b7f62f825d" // string | The entitlement ID
    updatePlanEntitlementRequestDto := *openapiclient.NewUpdatePlanEntitlementRequestDto(map[string]interface{}(true)) // UpdatePlanEntitlementRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.UpdatePlanEntitlement(context.Background(), planId, entitlementId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UpdatePlanEntitlementRequestDto(updatePlanEntitlementRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.UpdatePlanEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePlanEntitlement`: PlanEntitlementResponseDto
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.UpdatePlanEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | The plan ID | 
**entitlementId** | **string** | The entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlanEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 


 **updatePlanEntitlementRequestDto** | [**UpdatePlanEntitlementRequestDto**](UpdatePlanEntitlementRequestDto.md) |  | 

### Return type

[**PlanEntitlementResponseDto**](PlanEntitlementResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

