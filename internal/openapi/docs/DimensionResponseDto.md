# DimensionResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DimensionId** | **string** | A Kable-defined identifier for the dimension. | 
**Key** | **string** | An identifier for the dimension defined by you. This key corresponds to the key in the JSON payload of a usage event, and can be passed into the &#x60;record&#x60; library method. | 
**Name** | **string** | A human-readable name for the dimension, visible on dashboards, invoices, and reports. | 

## Methods

### NewDimensionResponseDto

`func NewDimensionResponseDto(dimensionId string, key string, name string, ) *DimensionResponseDto`

NewDimensionResponseDto instantiates a new DimensionResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDimensionResponseDtoWithDefaults

`func NewDimensionResponseDtoWithDefaults() *DimensionResponseDto`

NewDimensionResponseDtoWithDefaults instantiates a new DimensionResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensionId

`func (o *DimensionResponseDto) GetDimensionId() string`

GetDimensionId returns the DimensionId field if non-nil, zero value otherwise.

### GetDimensionIdOk

`func (o *DimensionResponseDto) GetDimensionIdOk() (*string, bool)`

GetDimensionIdOk returns a tuple with the DimensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionId

`func (o *DimensionResponseDto) SetDimensionId(v string)`

SetDimensionId sets DimensionId field to given value.


### GetKey

`func (o *DimensionResponseDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DimensionResponseDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DimensionResponseDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *DimensionResponseDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DimensionResponseDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DimensionResponseDto) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


