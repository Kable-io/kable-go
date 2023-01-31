# \CustomersApi

All URIs are relative to *https://live.kable.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelCustomer**](CustomersApi.md#CancelCustomer) | **Post** /api/v1/customers/{customerId}/cancel | cancel a customer
[**CreateCustomer**](CustomersApi.md#CreateCustomer) | **Post** /api/v1/customers/create | create customer
[**CreateCustomerBundle**](CustomersApi.md#CreateCustomerBundle) | **Post** /api/v1/customers/{customerId}/bundles/add | add customer bundle
[**CreateCustomerCreditGrant**](CustomersApi.md#CreateCustomerCreditGrant) | **Post** /api/v1/customers/{customerId}/credits/create | create customer credit grant
[**CreateCustomerPaymentMethod**](CustomersApi.md#CreateCustomerPaymentMethod) | **Post** /api/v1/customers/{customerId}/payment_methods/add | add customer payment method
[**CreateCustomerPlan**](CustomersApi.md#CreateCustomerPlan) | **Post** /api/v1/customers/{customerId}/plans/add | add customer plan(s)
[**DeleteCustomer**](CustomersApi.md#DeleteCustomer) | **Post** /api/v1/customers/{customerId}/delete | delete customer
[**GetCustomer**](CustomersApi.md#GetCustomer) | **Get** /api/v1/customers/{customerId} | get customer
[**GetCustomerCreditBalance**](CustomersApi.md#GetCustomerCreditBalance) | **Get** /api/v1/customers/{customerId}/credits/balance | get customer credit balance
[**GetCustomerEntitlement**](CustomersApi.md#GetCustomerEntitlement) | **Get** /api/v1/customers/{customerId}/entitlements/{entitlementId} | get customer entitlement
[**GetCustomerEntitlements**](CustomersApi.md#GetCustomerEntitlements) | **Get** /api/v1/customers/{customerId}/entitlements | get all customer entitlements
[**GetCustomers**](CustomersApi.md#GetCustomers) | **Get** /api/v1/customers | get all customers
[**RemoveCustomerBundle**](CustomersApi.md#RemoveCustomerBundle) | **Post** /api/v1/customers/{customerId}/bundles/remove | remove customer bundle
[**RemoveCustomerPlan**](CustomersApi.md#RemoveCustomerPlan) | **Post** /api/v1/customers/{customerId}/plans/remove | remove customer plan(s)
[**UpdateCustomer**](CustomersApi.md#UpdateCustomer) | **Post** /api/v1/customers/{customerId}/update | update customer



## CancelCustomer

> CustomerResponseDto CancelCustomer(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

cancel a customer



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
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | The customer ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CancelCustomer(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CancelCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelCustomer`: CustomerResponseDto
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CancelCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 


### Return type

[**CustomerResponseDto**](CustomerResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomer

> CustomerResponseWithKeysDto CreateCustomer(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateCustomerRequestDto(createCustomerRequestDto).Stripe(stripe).Keys(keys).Execute()

create customer



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
    createCustomerRequestDto := *openapiclient.NewCreateCustomerRequestDto("yourcompanyuser_1234567890", "Awesome Corp", "USD") // CreateCustomerRequestDto | 
    stripe := true // bool | When true, Kable will also create and attach a Stripe customer.  If your account is not connected to Stripe, the request will fail. If you provide a `stripeCustomerId` in the request body, this query parameter will be ignored. (optional)
    keys := true // bool | When true, Kable will create API keys for this customer. (This is only necessary if you use Kable for authentication.) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CreateCustomer(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateCustomerRequestDto(createCustomerRequestDto).Stripe(stripe).Keys(keys).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomer`: CustomerResponseWithKeysDto
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 
 **createCustomerRequestDto** | [**CreateCustomerRequestDto**](CreateCustomerRequestDto.md) |  | 
 **stripe** | **bool** | When true, Kable will also create and attach a Stripe customer.  If your account is not connected to Stripe, the request will fail. If you provide a &#x60;stripeCustomerId&#x60; in the request body, this query parameter will be ignored. | 
 **keys** | **bool** | When true, Kable will create API keys for this customer. (This is only necessary if you use Kable for authentication.) | 

### Return type

[**CustomerResponseWithKeysDto**](CustomerResponseWithKeysDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomerBundle

> CustomerResponseDto CreateCustomerBundle(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateCustomerBundleRequestDto(createCustomerBundleRequestDto).Execute()

add customer bundle



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
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | The customer ID
    createCustomerBundleRequestDto := *openapiclient.NewCreateCustomerBundleRequestDto("bun_a23e148584d0463482c961b7f62f824c") // CreateCustomerBundleRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CreateCustomerBundle(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateCustomerBundleRequestDto(createCustomerBundleRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateCustomerBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerBundle`: CustomerResponseDto
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateCustomerBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

 **createCustomerBundleRequestDto** | [**CreateCustomerBundleRequestDto**](CreateCustomerBundleRequestDto.md) |  | 

### Return type

[**CustomerResponseDto**](CustomerResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomerCreditGrant

> CreditGrantResponseDto CreateCustomerCreditGrant(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateCreditGrantRequestDto(createCreditGrantRequestDto).Execute()

create customer credit grant



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
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | The customer ID
    createCreditGrantRequestDto := *openapiclient.NewCreateCreditGrantRequestDto(float32(100)) // CreateCreditGrantRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CreateCustomerCreditGrant(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateCreditGrantRequestDto(createCreditGrantRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateCustomerCreditGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerCreditGrant`: CreditGrantResponseDto
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateCustomerCreditGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerCreditGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

 **createCreditGrantRequestDto** | [**CreateCreditGrantRequestDto**](CreateCreditGrantRequestDto.md) |  | 

### Return type

[**CreditGrantResponseDto**](CreditGrantResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomerPaymentMethod

> AddCustomerPaymentMethodResponseDto CreateCustomerPaymentMethod(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).AddCustomerPaymentMethodRequestDto(addCustomerPaymentMethodRequestDto).Execute()

add customer payment method



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
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | The customer ID
    addCustomerPaymentMethodRequestDto := *openapiclient.NewAddCustomerPaymentMethodRequestDto("https://yourcompany.com/signup/complete?success=true", "https://yourcompany.com/signup/failure?success=false") // AddCustomerPaymentMethodRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CreateCustomerPaymentMethod(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).AddCustomerPaymentMethodRequestDto(addCustomerPaymentMethodRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateCustomerPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerPaymentMethod`: AddCustomerPaymentMethodResponseDto
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateCustomerPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

 **addCustomerPaymentMethodRequestDto** | [**AddCustomerPaymentMethodRequestDto**](AddCustomerPaymentMethodRequestDto.md) |  | 

### Return type

[**AddCustomerPaymentMethodResponseDto**](AddCustomerPaymentMethodResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomerPlan

> CustomerResponseDto CreateCustomerPlan(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateCustomerPlansRequestDto(createCustomerPlansRequestDto).Execute()

add customer plan(s)



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
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | The customer ID
    createCustomerPlansRequestDto := *openapiclient.NewCreateCustomerPlansRequestDto() // CreateCustomerPlansRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CreateCustomerPlan(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateCustomerPlansRequestDto(createCustomerPlansRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateCustomerPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerPlan`: CustomerResponseDto
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateCustomerPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

 **createCustomerPlansRequestDto** | [**CreateCustomerPlansRequestDto**](CreateCustomerPlansRequestDto.md) |  | 

### Return type

[**CustomerResponseDto**](CustomerResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomer

> DeleteCustomer(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

delete customer



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
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | The customer ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.DeleteCustomer(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.DeleteCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerRequest struct via the builder pattern


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


## GetCustomer

> CustomerResponseDto GetCustomer(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get customer



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
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | The customer ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.GetCustomer(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomer`: CustomerResponseDto
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 


### Return type

[**CustomerResponseDto**](CustomerResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerCreditBalance

> CreditBalanceResponseDto GetCustomerCreditBalance(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get customer credit balance



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
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | The customer ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.GetCustomerCreditBalance(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetCustomerCreditBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerCreditBalance`: CreditBalanceResponseDto
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetCustomerCreditBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerCreditBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 


### Return type

[**CreditBalanceResponseDto**](CreditBalanceResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerEntitlement

> CustomerEntitlementResponseDto GetCustomerEntitlement(ctx, customerId, entitlementId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get customer entitlement



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
    customerId := "cus_2a3b4c5d6e7f8g9h0i1j2k3l4m5n6o7p" // string | The customer ID
    entitlementId := "ent_4d5e6f7g8h9i0j1k2l3m4n5o6p1a2b3c" // string | The entitlement ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.GetCustomerEntitlement(context.Background(), customerId, entitlementId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetCustomerEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerEntitlement`: CustomerEntitlementResponseDto
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetCustomerEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 
**entitlementId** | **string** | The entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 



### Return type

[**CustomerEntitlementResponseDto**](CustomerEntitlementResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerEntitlements

> []CustomerEntitlementResponseDto GetCustomerEntitlements(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get all customer entitlements



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
    customerId := "cus_2a3b4c5d6e7f8g9h0i1j2k3l4m5n6o7p" // string | The customer ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.GetCustomerEntitlements(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetCustomerEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerEntitlements`: []CustomerEntitlementResponseDto
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetCustomerEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 


### Return type

[**[]CustomerEntitlementResponseDto**](CustomerEntitlementResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomers

> []CustomerResponseDto GetCustomers(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get all customers



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
    resp, r, err := apiClient.CustomersApi.GetCustomers(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetCustomers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomers`: []CustomerResponseDto
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetCustomers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

### Return type

[**[]CustomerResponseDto**](CustomerResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCustomerBundle

> CustomerResponseDto RemoveCustomerBundle(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).RemoveCustomerBundleRequestDto(removeCustomerBundleRequestDto).Execute()

remove customer bundle



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
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | The customer ID
    removeCustomerBundleRequestDto := *openapiclient.NewRemoveCustomerBundleRequestDto("bun_a23e148584d0463482c961b7f62f824c") // RemoveCustomerBundleRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.RemoveCustomerBundle(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).RemoveCustomerBundleRequestDto(removeCustomerBundleRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.RemoveCustomerBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveCustomerBundle`: CustomerResponseDto
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.RemoveCustomerBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCustomerBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

 **removeCustomerBundleRequestDto** | [**RemoveCustomerBundleRequestDto**](RemoveCustomerBundleRequestDto.md) |  | 

### Return type

[**CustomerResponseDto**](CustomerResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCustomerPlan

> CustomerResponseDto RemoveCustomerPlan(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).RemoveCustomerPlansRequestDto(removeCustomerPlansRequestDto).Execute()

remove customer plan(s)



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
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | The customer ID
    removeCustomerPlansRequestDto := *openapiclient.NewRemoveCustomerPlansRequestDto() // RemoveCustomerPlansRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.RemoveCustomerPlan(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).RemoveCustomerPlansRequestDto(removeCustomerPlansRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.RemoveCustomerPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveCustomerPlan`: CustomerResponseDto
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.RemoveCustomerPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCustomerPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

 **removeCustomerPlansRequestDto** | [**RemoveCustomerPlansRequestDto**](RemoveCustomerPlansRequestDto.md) |  | 

### Return type

[**CustomerResponseDto**](CustomerResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomer

> CustomerResponseDto UpdateCustomer(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UpdateCustomerRequestDto(updateCustomerRequestDto).Execute()

update customer



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
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | The customer ID
    updateCustomerRequestDto := *openapiclient.NewUpdateCustomerRequestDto() // UpdateCustomerRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.UpdateCustomer(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UpdateCustomerRequestDto(updateCustomerRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.UpdateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomer`: CustomerResponseDto
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.UpdateCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

 **updateCustomerRequestDto** | [**UpdateCustomerRequestDto**](UpdateCustomerRequestDto.md) |  | 

### Return type

[**CustomerResponseDto**](CustomerResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

