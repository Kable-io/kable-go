# UsageMetricResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **time.Time** | Start of the period over which usage was evaluated. | 
**EndDate** | **time.Time** | End of the period over which usage was evaluated. | 
**Dimension** | [**UsageMetricIntervalResponseDtoDimension**](UsageMetricIntervalResponseDtoDimension.md) |  | 
**ClientId** | Pointer to **string** | The customer for which usage was evaluated. | [optional] 
**Total** | **float32** | The usage metric value for the given dimension and interval. | 
**Intervals** | [**[]UsageMetricIntervalResponseDto**](UsageMetricIntervalResponseDto.md) | Time-series intervals over which usage was evaluated. | 

## Methods

### NewUsageMetricResponseDto

`func NewUsageMetricResponseDto(startDate time.Time, endDate time.Time, dimension UsageMetricIntervalResponseDtoDimension, total float32, intervals []UsageMetricIntervalResponseDto, ) *UsageMetricResponseDto`

NewUsageMetricResponseDto instantiates a new UsageMetricResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMetricResponseDtoWithDefaults

`func NewUsageMetricResponseDtoWithDefaults() *UsageMetricResponseDto`

NewUsageMetricResponseDtoWithDefaults instantiates a new UsageMetricResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *UsageMetricResponseDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UsageMetricResponseDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UsageMetricResponseDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *UsageMetricResponseDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UsageMetricResponseDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UsageMetricResponseDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetDimension

`func (o *UsageMetricResponseDto) GetDimension() UsageMetricIntervalResponseDtoDimension`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *UsageMetricResponseDto) GetDimensionOk() (*UsageMetricIntervalResponseDtoDimension, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *UsageMetricResponseDto) SetDimension(v UsageMetricIntervalResponseDtoDimension)`

SetDimension sets Dimension field to given value.


### GetClientId

`func (o *UsageMetricResponseDto) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *UsageMetricResponseDto) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *UsageMetricResponseDto) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *UsageMetricResponseDto) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetTotal

`func (o *UsageMetricResponseDto) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UsageMetricResponseDto) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UsageMetricResponseDto) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetIntervals

`func (o *UsageMetricResponseDto) GetIntervals() []UsageMetricIntervalResponseDto`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *UsageMetricResponseDto) GetIntervalsOk() (*[]UsageMetricIntervalResponseDto, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *UsageMetricResponseDto) SetIntervals(v []UsageMetricIntervalResponseDto)`

SetIntervals sets Intervals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


