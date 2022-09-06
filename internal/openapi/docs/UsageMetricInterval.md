# UsageMetricInterval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntervalStart** | **time.Time** | Start of the interval | 
**IntervalEnd** | **time.Time** | End of the interval | 
**Dimension** | Pointer to [**Dimension**](Dimension.md) |  | [optional] 
**ClientId** | Pointer to **string** | When provided, query will consider only this customer. When &#x60;null&#x60;, query will consider all of your customers. | [optional] 
**Total** | **float32** | The total usage for the requested interval | 

## Methods

### NewUsageMetricInterval

`func NewUsageMetricInterval(intervalStart time.Time, intervalEnd time.Time, total float32, ) *UsageMetricInterval`

NewUsageMetricInterval instantiates a new UsageMetricInterval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMetricIntervalWithDefaults

`func NewUsageMetricIntervalWithDefaults() *UsageMetricInterval`

NewUsageMetricIntervalWithDefaults instantiates a new UsageMetricInterval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntervalStart

`func (o *UsageMetricInterval) GetIntervalStart() time.Time`

GetIntervalStart returns the IntervalStart field if non-nil, zero value otherwise.

### GetIntervalStartOk

`func (o *UsageMetricInterval) GetIntervalStartOk() (*time.Time, bool)`

GetIntervalStartOk returns a tuple with the IntervalStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalStart

`func (o *UsageMetricInterval) SetIntervalStart(v time.Time)`

SetIntervalStart sets IntervalStart field to given value.


### GetIntervalEnd

`func (o *UsageMetricInterval) GetIntervalEnd() time.Time`

GetIntervalEnd returns the IntervalEnd field if non-nil, zero value otherwise.

### GetIntervalEndOk

`func (o *UsageMetricInterval) GetIntervalEndOk() (*time.Time, bool)`

GetIntervalEndOk returns a tuple with the IntervalEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalEnd

`func (o *UsageMetricInterval) SetIntervalEnd(v time.Time)`

SetIntervalEnd sets IntervalEnd field to given value.


### GetDimension

`func (o *UsageMetricInterval) GetDimension() Dimension`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *UsageMetricInterval) GetDimensionOk() (*Dimension, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *UsageMetricInterval) SetDimension(v Dimension)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *UsageMetricInterval) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetClientId

`func (o *UsageMetricInterval) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *UsageMetricInterval) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *UsageMetricInterval) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *UsageMetricInterval) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetTotal

`func (o *UsageMetricInterval) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UsageMetricInterval) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UsageMetricInterval) SetTotal(v float32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


