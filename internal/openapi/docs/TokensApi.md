# \TokensApi

All URIs are relative to *https://live.kable.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](TokensApi.md#CreateToken) | **Post** /tokens/create | create a token



## CreateToken

> CreateToken200Response CreateToken(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateTokenRequest(createTokenRequest).Execute()

create a token



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
    kableClientSecret := "sk_test.jI92Cbu0.XeHWdYM1VTLy4oLtGMw8wrmpt5q9d04n" // string | Your `LIVE` or `TEST` secret key, depending on the environment in which the key is being created.
    createTokenRequest := *openapiclient.NewCreateTokenRequest("yourcompanyuser_1234567890") // CreateTokenRequest | The customer `clientId` for which to grant data access permissions to the token (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.CreateToken(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateTokenRequest(createTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.CreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToken`: CreateToken200Response
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.CreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key, depending on the environment in which the key is being created. | 
 **createTokenRequest** | [**CreateTokenRequest**](CreateTokenRequest.md) | The customer &#x60;clientId&#x60; for which to grant data access permissions to the token | 

### Return type

[**CreateToken200Response**](CreateToken200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

