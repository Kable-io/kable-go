# CreateDimensionRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key by which Kable identifies the dimension in your usage events. | 
**Name** | **string** | A human-readable name for the dimension which will be visible in dashboards and on invoices. Dimension names should be singular, not plural (ie \&quot;User,\&quot; not \&quot;Users\&quot;) | 

## Methods

### NewCreateDimensionRequestDto

`func NewCreateDimensionRequestDto(key string, name string, ) *CreateDimensionRequestDto`

NewCreateDimensionRequestDto instantiates a new CreateDimensionRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDimensionRequestDtoWithDefaults

`func NewCreateDimensionRequestDtoWithDefaults() *CreateDimensionRequestDto`

NewCreateDimensionRequestDtoWithDefaults instantiates a new CreateDimensionRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CreateDimensionRequestDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateDimensionRequestDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateDimensionRequestDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CreateDimensionRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDimensionRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDimensionRequestDto) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


