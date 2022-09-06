# \KeysApi

All URIs are relative to *https://live.kable.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKey**](KeysApi.md#CreateKey) | **Post** /keys/create | create a customer API key
[**RevokeKey**](KeysApi.md#RevokeKey) | **Post** /keys/revoke | revoke a customer API key(s)



## CreateKey

> CreateKey200Response CreateKey(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateKeyRequest(createKeyRequest).Execute()

create a customer API key



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
    createKeyRequest := *openapiclient.NewCreateKeyRequest("yourcompanyuser_1234567890", "TEST") // CreateKeyRequest | The customer and environment for which to create key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeysApi.CreateKey(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).CreateKeyRequest(createKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeysApi.CreateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKey`: CreateKey200Response
    fmt.Fprintf(os.Stdout, "Response from `KeysApi.CreateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key, depending on the environment in which the key is being created. | 
 **createKeyRequest** | [**CreateKeyRequest**](CreateKeyRequest.md) | The customer and environment for which to create key | 

### Return type

[**CreateKey200Response**](CreateKey200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeKey

> RevokeKey(ctx).KableClientId(kableClientId).KableClientSecret(kableClientSecret).RevokeKeyRequest(revokeKeyRequest).Execute()

revoke a customer API key(s)



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
    revokeKeyRequest := *openapiclient.NewRevokeKeyRequest("yourcompanyuser_1234567890", "TEST") // RevokeKeyRequest | The customer and environment for which to revoke key(s).  By default, all keys for a given customer and environment are revoked. To revoke an individual key, include the optional `key` parameter to revoke a single, specified key.  By default, revoked keys are useable for 10 minutes following revocation (to give developers time to update configurations). To revoke a key(s) immediately, use the optional `immediately` parameter.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeysApi.RevokeKey(context.Background()).KableClientId(kableClientId).KableClientSecret(kableClientSecret).RevokeKeyRequest(revokeKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeysApi.RevokeKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kableClientId** | **string** | Your client ID, found in the dashboard of your Kable account. | 
 **kableClientSecret** | **string** | Your &#x60;LIVE&#x60; or &#x60;TEST&#x60; secret key, depending on the environment in which the key is being created. | 
 **revokeKeyRequest** | [**RevokeKeyRequest**](RevokeKeyRequest.md) | The customer and environment for which to revoke key(s).  By default, all keys for a given customer and environment are revoked. To revoke an individual key, include the optional &#x60;key&#x60; parameter to revoke a single, specified key.  By default, revoked keys are useable for 10 minutes following revocation (to give developers time to update configurations). To revoke a key(s) immediately, use the optional &#x60;immediately&#x60; parameter.  | 

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

