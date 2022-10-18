# \CreditGrantsApi

All URIs are relative to *https://live.kable.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCreditGrant**](CreditGrantsApi.md#CreateCreditGrant) | **Post** /customers/{customerId}/credits/create | create a new credit grant for a customer
[**GetCustomerCreditBalance**](CreditGrantsApi.md#GetCustomerCreditBalance) | **Get** /customers/{customerId}/credits/balance | get a customer&#39;s credit balance



## CreateCreditGrant

> CreditGrant CreateCreditGrant(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateCreditGrantRequest(createCreditGrantRequest).Execute()

create a new credit grant for a customer



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
    createCreditGrantRequest := *openapiclient.NewCreateCreditGrantRequest(float32(100)) // CreateCreditGrantRequest | Information about the credit grant to create. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreditGrantsApi.CreateCreditGrant(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateCreditGrantRequest(createCreditGrantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditGrantsApi.CreateCreditGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreditGrant`: CreditGrant
    fmt.Fprintf(os.Stdout, "Response from `CreditGrantsApi.CreateCreditGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The identifier for the customer. You can pass in *either* the &#x60;customerId&#x60; (as defined by Kable) or the &#x60;clientId&#x60; (as defined by you). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreditGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. Each customer will have separate keys for &#x60;LIVE&#x60; and &#x60;TEST&#x60; environments of your API. | 

 **createCreditGrantRequest** | [**CreateCreditGrantRequest**](CreateCreditGrantRequest.md) | Information about the credit grant to create. | 

### Return type

[**CreditGrant**](CreditGrant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerCreditBalance

> GetCustomerCreditBalance200Response GetCustomerCreditBalance(ctx, customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get a customer's credit balance



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
    resp, r, err := apiClient.CreditGrantsApi.GetCustomerCreditBalance(context.Background(), customerId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreditGrantsApi.GetCustomerCreditBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerCreditBalance`: GetCustomerCreditBalance200Response
    fmt.Fprintf(os.Stdout, "Response from `CreditGrantsApi.GetCustomerCreditBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The identifier for the customer. You can pass in *either* the &#x60;customerId&#x60; (as defined by Kable) or the &#x60;clientId&#x60; (as defined by you). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerCreditBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. Customers exist across all environments, so it does not matter which environment you use to create customers. Each customer will have separate keys for &#x60;LIVE&#x60; and &#x60;TEST&#x60; environments of your API. | 


### Return type

[**GetCustomerCreditBalance200Response**](GetCustomerCreditBalance200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

