# \TokensApi

All URIs are relative to *https://live.kable.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](TokensApi.md#CreateToken) | **Post** /api/v1/tokens/create | create token



## CreateToken

> DashboardTokenResponseDto CreateToken(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateTokenDto(createTokenDto).Execute()

create token



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
    createTokenDto := *openapiclient.NewCreateTokenDto("yourcompanyuser_1234567890") // CreateTokenDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.CreateToken(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateTokenDto(createTokenDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.CreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToken`: DashboardTokenResponseDto
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.CreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key. | 
 **createTokenDto** | [**CreateTokenDto**](CreateTokenDto.md) |  | 

### Return type

[**DashboardTokenResponseDto**](DashboardTokenResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

