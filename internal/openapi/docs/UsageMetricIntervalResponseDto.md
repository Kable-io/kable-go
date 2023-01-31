# UsageMetricIntervalResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntervalStart** | **time.Time** | Start of the interval over which usage was evaluated. | 
**IntervalEnd** | **time.Time** | End of the interval over which usage was evaluated. | 
**Dimension** | [**UsageMetricIntervalResponseDtoDimension**](UsageMetricIntervalResponseDtoDimension.md) |  | 
**ClientId** | Pointer to **string** | The customer for which usage was evaluated. | [optional] 
**Total** | **float32** | The usage metric value for the given dimension and interval. | 

## Methods

### NewUsageMetricIntervalResponseDto

`func NewUsageMetricIntervalResponseDto(intervalStart time.Time, intervalEnd time.Time, dimension UsageMetricIntervalResponseDtoDimension, total float32, ) *UsageMetricIntervalResponseDto`

NewUsageMetricIntervalResponseDto instantiates a new UsageMetricIntervalResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMetricIntervalResponseDtoWithDefaults

`func NewUsageMetricIntervalResponseDtoWithDefaults() *UsageMetricIntervalResponseDto`

NewUsageMetricIntervalResponseDtoWithDefaults instantiates a new UsageMetricIntervalResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntervalStart

`func (o *UsageMetricIntervalResponseDto) GetIntervalStart() time.Time`

GetIntervalStart returns the IntervalStart field if non-nil, zero value otherwise.

### GetIntervalStartOk

`func (o *UsageMetricIntervalResponseDto) GetIntervalStartOk() (*time.Time, bool)`

GetIntervalStartOk returns a tuple with the IntervalStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalStart

`func (o *UsageMetricIntervalResponseDto) SetIntervalStart(v time.Time)`

SetIntervalStart sets IntervalStart field to given value.


### GetIntervalEnd

`func (o *UsageMetricIntervalResponseDto) GetIntervalEnd() time.Time`

GetIntervalEnd returns the IntervalEnd field if non-nil, zero value otherwise.

### GetIntervalEndOk

`func (o *UsageMetricIntervalResponseDto) GetIntervalEndOk() (*time.Time, bool)`

GetIntervalEndOk returns a tuple with the IntervalEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalEnd

`func (o *UsageMetricIntervalResponseDto) SetIntervalEnd(v time.Time)`

SetIntervalEnd sets IntervalEnd field to given value.


### GetDimension

`func (o *UsageMetricIntervalResponseDto) GetDimension() UsageMetricIntervalResponseDtoDimension`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *UsageMetricIntervalResponseDto) GetDimensionOk() (*UsageMetricIntervalResponseDtoDimension, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *UsageMetricIntervalResponseDto) SetDimension(v UsageMetricIntervalResponseDtoDimension)`

SetDimension sets Dimension field to given value.


### GetClientId

`func (o *UsageMetricIntervalResponseDto) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *UsageMetricIntervalResponseDto) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *UsageMetricIntervalResponseDto) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *UsageMetricIntervalResponseDto) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetTotal

`func (o *UsageMetricIntervalResponseDto) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UsageMetricIntervalResponseDto) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UsageMetricIntervalResponseDto) SetTotal(v float32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


