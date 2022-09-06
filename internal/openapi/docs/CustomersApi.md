# \CustomersApi

All URIs are relative to *https://live.kable.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCustomerPaymentMethod**](CustomersApi.md#AddCustomerPaymentMethod) | **Post** /customers/{customerId}/payment_methods/add | add a customer payment method
[**AddCustomerPlans**](CustomersApi.md#AddCustomerPlans) | **Post** /customers/{customerId}/plans/add | add a plan(s) to a customer
[**CancelCustomer**](CustomersApi.md#CancelCustomer) | **Post** /customers/{customerId}/cancel | cancel a customer
[**CreateCustomer**](CustomersApi.md#CreateCustomer) | **Post** /customers/create | create a customer
[**GetAllCustomers**](CustomersApi.md#GetAllCustomers) | **Get** /customers | get all customers
[**GetCustomer**](CustomersApi.md#GetCustomer) | **Get** /customers/{customerId} | get customer
[**RemoveCustomerPlans**](CustomersApi.md#RemoveCustomerPlans) | **Post** /customers/{customerId}/plans/remove | remove a plan(s) from a customer
[**UpdateCustomer**](CustomersApi.md#UpdateCustomer) | **Post** /customers/{customerId}/update | update a customer



## AddCustomerPaymentMethod

> AddCustomerPaymentMethod200Response AddCustomerPaymentMethod(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).AddCustomerPaymentMethodRequest(addCustomerPaymentMethodRequest).Execute()

add a customer payment method



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
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | The identifier for the customer. You can pass in *either* the `customerId` (as defined by Kable) or the `clientId` (as defined by you).
    addCustomerPaymentMethodRequest := *openapiclient.NewAddCustomerPaymentMethodRequest("https://yourcompany.com/signup/complete?success=true", "https://yourcompany.com/signup/failure?success=false") // AddCustomerPaymentMethodRequest | URLs to which to redirect after payment method details are stored in Stripe (and completion of Stripe Checkout Session). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.AddCustomerPaymentMethod(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).AddCustomerPaymentMethodRequest(addCustomerPaymentMethodRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.AddCustomerPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCustomerPaymentMethod`: AddCustomerPaymentMethod200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.AddCustomerPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The identifier for the customer. You can pass in *either* the &#x60;customerId&#x60; (as defined by Kable) or the &#x60;clientId&#x60; (as defined by you). | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCustomerPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. Each customer will have separate keys for &#x60;LIVE&#x60; and &#x60;TEST&#x60; environments of your API. | 

 **addCustomerPaymentMethodRequest** | [**AddCustomerPaymentMethodRequest**](AddCustomerPaymentMethodRequest.md) | URLs to which to redirect after payment method details are stored in Stripe (and completion of Stripe Checkout Session). | 

### Return type

[**AddCustomerPaymentMethod200Response**](AddCustomerPaymentMethod200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddCustomerPlans

> Customer AddCustomerPlans(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).AddCustomerPlansRequest(addCustomerPlansRequest).Execute()

add a plan(s) to a customer



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
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | The identifier for the customer. You can pass in *either* the `customerId` (as defined by Kable) or the `clientId` (as defined by you).
    addCustomerPlansRequest := *openapiclient.NewAddCustomerPlansRequest() // AddCustomerPlansRequest | Plan ID(s) to add. You can provide either a singular `planId` or a list of `planIds`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.AddCustomerPlans(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).AddCustomerPlansRequest(addCustomerPlansRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.AddCustomerPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCustomerPlans`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.AddCustomerPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The identifier for the customer. You can pass in *either* the &#x60;customerId&#x60; (as defined by Kable) or the &#x60;clientId&#x60; (as defined by you). | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCustomerPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. Each customer will have separate keys for &#x60;LIVE&#x60; and &#x60;TEST&#x60; environments of your API. | 

 **addCustomerPlansRequest** | [**AddCustomerPlansRequest**](AddCustomerPlansRequest.md) | Plan ID(s) to add. You can provide either a singular &#x60;planId&#x60; or a list of &#x60;planIds&#x60;. | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelCustomer

> Customer CancelCustomer(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

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
    kableClientId := "kci_3c90e9ac92c64f31ae8ed84d21e18740" // string | Your client ID, found in the dashboard of your Kable account.
    kableClientSecret := "sk_test.jI92Cbu0.XeHWdYM1VTLy4oLtGMw8wrmpt5q9d04n" // string | Your `LIVE` or `TEST` secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. Each customer will have separate keys for `LIVE` and `TEST` environments of your API.
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | The identifier for the customer. You can pass in *either* the `customerId` (as defined by Kable) or the `clientId` (as defined by you).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CancelCustomer(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CancelCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelCustomer`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CancelCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The identifier for the customer. You can pass in *either* the &#x60;customerId&#x60; (as defined by Kable) or the &#x60;clientId&#x60; (as defined by you). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. Each customer will have separate keys for &#x60;LIVE&#x60; and &#x60;TEST&#x60; environments of your API. | 


### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomer

> CreateCustomer200Response CreateCustomer(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Keys(keys).Stripe(stripe).CreateCustomerRequest(createCustomerRequest).Execute()

create a customer



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
    kableClientSecret := "sk_test.jI92Cbu0.XeHWdYM1VTLy4oLtGMw8wrmpt5q9d04n" // string | Your `LIVE` or `TEST` secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. You will be able to record usage data for each customer separately in `LIVE` and `TEST` environments.
    keys := true // bool | When true, Kable will create API keys for this customer. (This is only necessary if you use Kable for authentication.) (optional)
    stripe := true // bool | When true, Kable will also create and attach a Stripe customer.  If your account is not connected to Stripe, the request will fail. If you provide a `stripeCustomerId` in the request body, this query parameter will be ignored.  (optional)
    createCustomerRequest := *openapiclient.NewCreateCustomerRequest("yourcompanyuser_1234567890", "Awesome Corp") // CreateCustomerRequest | Information about the customer you are creating.  Client IDs should correspond to identifiers in your own application. If you are migrating from another authentication system, this is often the same `clientId` that the customer was using before. Each client ID must be unique, as Kable uses this field to uniquely identify customers. Company names can be edited later in the dashboard, but client IDs cannot be changed.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CreateCustomer(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Keys(keys).Stripe(stripe).CreateCustomerRequest(createCustomerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomer`: CreateCustomer200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. You will be able to record usage data for each customer separately in &#x60;LIVE&#x60; and &#x60;TEST&#x60; environments. | 
 **keys** | **bool** | When true, Kable will create API keys for this customer. (This is only necessary if you use Kable for authentication.) | 
 **stripe** | **bool** | When true, Kable will also create and attach a Stripe customer.  If your account is not connected to Stripe, the request will fail. If you provide a &#x60;stripeCustomerId&#x60; in the request body, this query parameter will be ignored.  | 
 **createCustomerRequest** | [**CreateCustomerRequest**](CreateCustomerRequest.md) | Information about the customer you are creating.  Client IDs should correspond to identifiers in your own application. If you are migrating from another authentication system, this is often the same &#x60;clientId&#x60; that the customer was using before. Each client ID must be unique, as Kable uses this field to uniquely identify customers. Company names can be edited later in the dashboard, but client IDs cannot be changed.  | 

### Return type

[**CreateCustomer200Response**](CreateCustomer200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllCustomers

> []Customer GetAllCustomers(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

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
    kableClientId := "kci_3c90e9ac92c64f31ae8ed84d21e18740" // string | Your client ID, found in the dashboard of your Kable account.
    kableClientSecret := "sk_test.jI92Cbu0.XeHWdYM1VTLy4oLtGMw8wrmpt5q9d04n" // string | Your `LIVE` or `TEST` secret key. Customers exist across all environments, so it does not matter which environment you use to fetch customers.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.GetAllCustomers(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetAllCustomers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllCustomers`: []Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetAllCustomers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to fetch customers. | 

### Return type

[**[]Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomer

> Customer GetCustomer(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

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
    kableClientId := "kci_3c90e9ac92c64f31ae8ed84d21e18740" // string | Your client ID, found in the dashboard of your Kable account.
    kableClientSecret := "sk_test.jI92Cbu0.XeHWdYM1VTLy4oLtGMw8wrmpt5q9d04n" // string | Your `LIVE` or `TEST` secret key. Customers exist across all environments, so it does not matter which environment you use to fetch customers.
    customerId := "yourcompanyuser_1234567890" // string | The identifier for the customer. You can pass in *either* the `customerId` (as defined by Kable) or the `clientId` (as defined by you).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.GetCustomer(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomer`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The identifier for the customer. You can pass in *either* the &#x60;customerId&#x60; (as defined by Kable) or the &#x60;clientId&#x60; (as defined by you). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to fetch customers. | 


### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCustomerPlans

> Customer RemoveCustomerPlans(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).AddCustomerPlansRequest(addCustomerPlansRequest).Execute()

remove a plan(s) from a customer



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
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | The identifier for the customer. You can pass in *either* the `customerId` (as defined by Kable) or the `clientId` (as defined by you).
    addCustomerPlansRequest := *openapiclient.NewAddCustomerPlansRequest() // AddCustomerPlansRequest | Plan ID(s) to remove. You can provide either a singular `planId` or a list of `planIds`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.RemoveCustomerPlans(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).AddCustomerPlansRequest(addCustomerPlansRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.RemoveCustomerPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveCustomerPlans`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.RemoveCustomerPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The identifier for the customer. You can pass in *either* the &#x60;customerId&#x60; (as defined by Kable) or the &#x60;clientId&#x60; (as defined by you). | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCustomerPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. Each customer will have separate keys for &#x60;LIVE&#x60; and &#x60;TEST&#x60; environments of your API. | 

 **addCustomerPlansRequest** | [**AddCustomerPlansRequest**](AddCustomerPlansRequest.md) | Plan ID(s) to remove. You can provide either a singular &#x60;planId&#x60; or a list of &#x60;planIds&#x60;. | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomer

> Customer UpdateCustomer(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UpdateCustomerRequest(updateCustomerRequest).Execute()

update a customer



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
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | The identifier for the customer. You can pass in *either* the `customerId` (as defined by Kable) or the `clientId` (as defined by you).
    updateCustomerRequest := *openapiclient.NewUpdateCustomerRequest() // UpdateCustomerRequest | Information to update about the customer. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.UpdateCustomer(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UpdateCustomerRequest(updateCustomerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.UpdateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomer`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.UpdateCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The identifier for the customer. You can pass in *either* the &#x60;customerId&#x60; (as defined by Kable) or the &#x60;clientId&#x60; (as defined by you). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. Each customer will have separate keys for &#x60;LIVE&#x60; and &#x60;TEST&#x60; environments of your API. | 

 **updateCustomerRequest** | [**UpdateCustomerRequest**](UpdateCustomerRequest.md) | Information to update about the customer. | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

