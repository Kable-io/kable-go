# \UsageApi

All URIs are relative to *https://live.kable.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsage**](UsageApi.md#GetUsage) | **Post** /usage/get | get usage metrics



## GetUsage

> UsageMetricResponse GetUsage(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).GetUsageRequest(getUsageRequest).Execute()

get usage metrics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    kableClientId := "kci_3c90e9ac92c64f31ae8ed84d21e18740" // string | Your client ID, found in the dashboard of your Kable account.
    kableClientSecret := "sk_test.jI92Cbu0.XeHWdYM1VTLy4oLtGMw8wrmpt5q9d04n" // string | Your `LIVE` or `TEST` secret key, depending on the environment in which the key is being created.
    getUsageRequest := *openapiclient.NewGetUsageRequest("COUNT_DISTINCT", time.Now(), time.Now(), "DAY") // GetUsageRequest | Parameters of the usage metrics query to execute. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetUsage(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).GetUsageRequest(getUsageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsage`: UsageMetricResponse
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key, depending on the environment in which the key is being created. | 
 **getUsageRequest** | [**GetUsageRequest**](GetUsageRequest.md) | Parameters of the usage metrics query to execute. | 

### Return type

[**UsageMetricResponse**](UsageMetricResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

