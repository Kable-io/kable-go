# Key

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **string** | The environment to which this key corresponds. | 
**ClientId** | Pointer to **string** | The identifier (defined by you) for the customer to which this key belongs. | [optional] 
**Key** | **string** | The secret value of the key. | 

## Methods

### NewKey

`func NewKey(environment string, key string, ) *Key`

NewKey instantiates a new Key object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyWithDefaults

`func NewKeyWithDefaults() *Key`

NewKeyWithDefaults instantiates a new Key object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *Key) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Key) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Key) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetClientId

`func (o *Key) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Key) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Key) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Key) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetKey

`func (o *Key) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Key) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Key) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


