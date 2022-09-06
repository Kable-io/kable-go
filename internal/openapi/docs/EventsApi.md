# \EventsApi

All URIs are relative to *https://live.kable.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEvents**](EventsApi.md#CreateEvents) | **Post** /events/create | record events



## CreateEvents

> CreateEvents(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Event(event).Execute()

record events



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
    event := []openapiclient.Event{*openapiclient.NewEvent("ClientId_example", time.Now())} // []Event | Events containing information about things that happened in your API (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.CreateEvents(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).Event(event).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.CreateEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key, depending on the environment in which the key is being created. | 
 **event** | [**[]Event**](Event.md) | Events containing information about things that happened in your API | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

