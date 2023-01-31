# \UsageApi

All URIs are relative to *https://live.kable.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsage**](UsageApi.md#GetUsage) | **Post** /api/v1/usage/get | get usage metric



## GetUsage

> UsageMetricResponseDto GetUsage(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UsageMetricRequestDto(usageMetricRequestDto).Execute()

get usage metric



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
    kableClientId := "kableClientId_example" // string | Your Kable client ID, found in the dashboard of your Kable account.
    kableClientSecret := "kableClientSecret_example" // string | Your `LIVE` or `TEST` secret key.
    usageMetricRequestDto := *openapiclient.NewUsageMetricRequestDto("COUNT", time.Now(), time.Now(), "DAY") // UsageMetricRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetUsage(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).UsageMetricRequestDto(usageMetricRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsage`: UsageMetricResponseDto
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 
 **usageMetricRequestDto** | [**UsageMetricRequestDto**](UsageMetricRequestDto.md) |  | 

### Return type

[**UsageMetricResponseDto**](UsageMetricResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

