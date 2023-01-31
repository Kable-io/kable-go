# \InvoicesApi

All URIs are relative to *https://live.kable.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadInvoice**](InvoicesApi.md#DownloadInvoice) | **Get** /api/v1/invoices/{invoiceId}/download | download invoice
[**GetInvoice**](InvoicesApi.md#GetInvoice) | **Get** /api/v1/invoices/{invoiceId} | get invoice
[**GetInvoices**](InvoicesApi.md#GetInvoices) | **Get** /api/v1/invoices | get all invoices



## DownloadInvoice

> DownloadInvoice(ctx, invoiceId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Format(format).Execute()

download invoice



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
    invoiceId := "inv_346daaa242674b179ecc44f43b5b9db4" // string | The invoice ID
    format := "pdf" // string | The downloaded file format

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesApi.DownloadInvoice(context.Background(), invoiceId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.DownloadInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | The invoice ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 

 **format** | **string** | The downloaded file format | 

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


## GetInvoice

> InvoiceResponseDto GetInvoice(ctx, invoiceId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()

get invoice



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
    invoiceId := "inv_346daaa242674b179ecc44f43b5b9db4" // string | The invoice ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesApi.GetInvoice(context.Background(), invoiceId).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.GetInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoice`: InvoiceResponseDto
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.GetInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | The invoice ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 


### Return type

[**InvoiceResponseDto**](InvoiceResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoices

> []InvoiceResponseDto GetInvoices(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CustomerId(customerId).ClientId(clientId).Execute()

get all invoices



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
    customerId := "cus_8276e1ac8ed84d21c64f31ae0082fe" // string | A Kable-defined identifier for the customer. (optional)
    clientId := "yourcompanyuser_1234567890" // string | A unique identifier for the customer, defined by you. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvoicesApi.GetInvoices(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CustomerId(customerId).ClientId(clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.GetInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoices`: []InvoiceResponseDto
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.GetInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 
 **customerId** | **string** | A Kable-defined identifier for the customer. | 
 **clientId** | **string** | A unique identifier for the customer, defined by you. | 

### Return type

[**[]InvoiceResponseDto**](InvoiceResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

