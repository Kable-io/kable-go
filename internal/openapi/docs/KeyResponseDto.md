# KeyResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | A unique identifier for the customer, defined by you. | 
**Key** | **string** | The API key value | 
**RevokedAt** | Pointer to **time.Time** | The time at which the key was revoked | [optional] 
**Created** | **time.Time** | Date at which entity was first created | 

## Methods

### NewKeyResponseDto

`func NewKeyResponseDto(clientId string, key string, created time.Time, ) *KeyResponseDto`

NewKeyResponseDto instantiates a new KeyResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyResponseDtoWithDefaults

`func NewKeyResponseDtoWithDefaults() *KeyResponseDto`

NewKeyResponseDtoWithDefaults instantiates a new KeyResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *KeyResponseDto) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *KeyResponseDto) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *KeyResponseDto) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetKey

`func (o *KeyResponseDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KeyResponseDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KeyResponseDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetRevokedAt

`func (o *KeyResponseDto) GetRevokedAt() time.Time`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *KeyResponseDto) GetRevokedAtOk() (*time.Time, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *KeyResponseDto) SetRevokedAt(v time.Time)`

SetRevokedAt sets RevokedAt field to given value.

### HasRevokedAt

`func (o *KeyResponseDto) HasRevokedAt() bool`

HasRevokedAt returns a boolean if a field has been set.

### GetCreated

`func (o *KeyResponseDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *KeyResponseDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *KeyResponseDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


