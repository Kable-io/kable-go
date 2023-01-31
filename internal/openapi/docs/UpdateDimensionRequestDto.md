# UpdateDimensionRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-readable name for the dimension which will be visible in dashboards and on invoices. Dimension names should be singular, not plural (ie \&quot;User,\&quot; not \&quot;Users\&quot;). | 

## Methods

### NewUpdateDimensionRequestDto

`func NewUpdateDimensionRequestDto(name string, ) *UpdateDimensionRequestDto`

NewUpdateDimensionRequestDto instantiates a new UpdateDimensionRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDimensionRequestDtoWithDefaults

`func NewUpdateDimensionRequestDtoWithDefaults() *UpdateDimensionRequestDto`

NewUpdateDimensionRequestDtoWithDefaults instantiates a new UpdateDimensionRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDimensionRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDimensionRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDimensionRequestDto) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


