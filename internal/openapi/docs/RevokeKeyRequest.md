# RevokeKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**Environment** | **string** |  | 
**Key** | Pointer to **string** |  | [optional] 
**Immediately** | Pointer to **bool** |  | [optional] 

## Methods

### NewRevokeKeyRequest

`func NewRevokeKeyRequest(clientId string, environment string, ) *RevokeKeyRequest`

NewRevokeKeyRequest instantiates a new RevokeKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevokeKeyRequestWithDefaults

`func NewRevokeKeyRequestWithDefaults() *RevokeKeyRequest`

NewRevokeKeyRequestWithDefaults instantiates a new RevokeKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *RevokeKeyRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *RevokeKeyRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *RevokeKeyRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetEnvironment

`func (o *RevokeKeyRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RevokeKeyRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RevokeKeyRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetKey

`func (o *RevokeKeyRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RevokeKeyRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RevokeKeyRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RevokeKeyRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetImmediately

`func (o *RevokeKeyRequest) GetImmediately() bool`

GetImmediately returns the Immediately field if non-nil, zero value otherwise.

### GetImmediatelyOk

`func (o *RevokeKeyRequest) GetImmediatelyOk() (*bool, bool)`

GetImmediatelyOk returns a tuple with the Immediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediately

`func (o *RevokeKeyRequest) SetImmediately(v bool)`

SetImmediately sets Immediately field to given value.

### HasImmediately

`func (o *RevokeKeyRequest) HasImmediately() bool`

HasImmediately returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


