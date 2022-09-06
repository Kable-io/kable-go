# UsageMetricResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **time.Time** | Start of the period over which to query. Timestamps must be formatted as RFC 3339 strings like &#x60;2022-01-09T09:32:01Z&#x60; | 
**EndDate** | **time.Time** | End of the period over which to query. Timestamps must be formatted as RFC 3339 strings like &#x60;2022-01-09T09:32:01Z&#x60; | 
**Dimension** | Pointer to [**Dimension**](Dimension.md) |  | [optional] 
**ClientId** | Pointer to **string** | When provided, query will consider only this customer. When &#x60;null&#x60;, query will consider all of your customers. | [optional] 
**Intervals** | [**[]UsageMetricInterval**](UsageMetricInterval.md) |  | 

## Methods

### NewUsageMetricResponse

`func NewUsageMetricResponse(startDate time.Time, endDate time.Time, intervals []UsageMetricInterval, ) *UsageMetricResponse`

NewUsageMetricResponse instantiates a new UsageMetricResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMetricResponseWithDefaults

`func NewUsageMetricResponseWithDefaults() *UsageMetricResponse`

NewUsageMetricResponseWithDefaults instantiates a new UsageMetricResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *UsageMetricResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UsageMetricResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UsageMetricResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *UsageMetricResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UsageMetricResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UsageMetricResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetDimension

`func (o *UsageMetricResponse) GetDimension() Dimension`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *UsageMetricResponse) GetDimensionOk() (*Dimension, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *UsageMetricResponse) SetDimension(v Dimension)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *UsageMetricResponse) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetClientId

`func (o *UsageMetricResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *UsageMetricResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *UsageMetricResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *UsageMetricResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetIntervals

`func (o *UsageMetricResponse) GetIntervals() []UsageMetricInterval`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *UsageMetricResponse) GetIntervalsOk() (*[]UsageMetricInterval, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *UsageMetricResponse) SetIntervals(v []UsageMetricInterval)`

SetIntervals sets Intervals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


