# \AuthenticationApi

All URIs are relative to *https://live.kable.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DoAuthentication**](AuthenticationApi.md#DoAuthentication) | **Post** /api/v1/authenticate | authenticate API request



## DoAuthentication

> DoAuthentication(ctx).KableClientId(kableClientId).XClientId(xClientId).XApiKey(xApiKey).Execute()

authenticate API request



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
    xClientId := "xClientId_example" // string | The client ID of the customer to authenticate.
    xApiKey := "xApiKey_example" // string | An API key to authenticate for the provided X-Client-Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.DoAuthentication(context.Background()).KableClientId(kableClientId).XClientId(xClientId).XApiKey(xApiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.DoAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDoAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **xClientId** | **string** | The client ID of the customer to authenticate. | 
 **xApiKey** | **string** | An API key to authenticate for the provided X-Client-Id. | 

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

