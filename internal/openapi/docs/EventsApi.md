# \EventsApi

All URIs are relative to *https://live.kable.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEvent**](EventsApi.md#CreateEvent) | **Post** /api/v1/events/create | record events



## CreateEvent

> CreateEvent(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).EventDto(eventDto).Execute()

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
    kableClientId := "kableClientId_example" // string | Your Kable client ID, found in the dashboard of your Kable account.
    kableClientSecret := "kableClientSecret_example" // string | Your `LIVE` or `TEST` secret key.
    eventDto := []openapiclient.EventDto{*openapiclient.NewEventDto("yourcompanyuser_1234567890", time.Now())} // []EventDto | An event or list of usage events to record.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.CreateEvent(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).EventDto(eventDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.CreateEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 
 **eventDto** | [**[]EventDto**](EventDto.md) | An event or list of usage events to record. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

