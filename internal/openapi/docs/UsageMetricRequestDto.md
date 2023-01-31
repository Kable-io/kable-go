# UsageMetricRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | **string** | The aggregation by which to evaluate the metric. | 
**DimensionId** | Pointer to **string** | The dimension to measure.  Note that one of either &#x60;dimensionId&#x60; or &#x60;dimensionKey&#x60; must be provided. | [optional] 
**DimensionKey** | Pointer to **string** | The dimension to measure.  Note that one of either &#x60;dimensionId&#x60; or &#x60;dimensionKey&#x60; must be provided. | [optional] 
**StartDate** | **time.Time** | Start of the period over which to query. Timestamps must be formatted as RFC 3339 strings like &#x60;2022-01-09T09:32:01Z&#x60;. | 
**EndDate** | **time.Time** | End of the period over which to query. Timestamps must be formatted as RFC 3339 strings like &#x60;2022-01-09T09:32:01Z&#x60;. | 
**Interval** | **string** | Timeseries unit duration over which to return intervals. | 
**ClientId** | Pointer to **string** | When provided, query will consider only this customer. When &#x60;null&#x60;, query will consider all of your customers. | [optional] 

## Methods

### NewUsageMetricRequestDto

`func NewUsageMetricRequestDto(aggregation string, startDate time.Time, endDate time.Time, interval string, ) *UsageMetricRequestDto`

NewUsageMetricRequestDto instantiates a new UsageMetricRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMetricRequestDtoWithDefaults

`func NewUsageMetricRequestDtoWithDefaults() *UsageMetricRequestDto`

NewUsageMetricRequestDtoWithDefaults instantiates a new UsageMetricRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *UsageMetricRequestDto) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *UsageMetricRequestDto) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *UsageMetricRequestDto) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.


### GetDimensionId

`func (o *UsageMetricRequestDto) GetDimensionId() string`

GetDimensionId returns the DimensionId field if non-nil, zero value otherwise.

### GetDimensionIdOk

`func (o *UsageMetricRequestDto) GetDimensionIdOk() (*string, bool)`

GetDimensionIdOk returns a tuple with the DimensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionId

`func (o *UsageMetricRequestDto) SetDimensionId(v string)`

SetDimensionId sets DimensionId field to given value.

### HasDimensionId

`func (o *UsageMetricRequestDto) HasDimensionId() bool`

HasDimensionId returns a boolean if a field has been set.

### GetDimensionKey

`func (o *UsageMetricRequestDto) GetDimensionKey() string`

GetDimensionKey returns the DimensionKey field if non-nil, zero value otherwise.

### GetDimensionKeyOk

`func (o *UsageMetricRequestDto) GetDimensionKeyOk() (*string, bool)`

GetDimensionKeyOk returns a tuple with the DimensionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionKey

`func (o *UsageMetricRequestDto) SetDimensionKey(v string)`

SetDimensionKey sets DimensionKey field to given value.

### HasDimensionKey

`func (o *UsageMetricRequestDto) HasDimensionKey() bool`

HasDimensionKey returns a boolean if a field has been set.

### GetStartDate

`func (o *UsageMetricRequestDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UsageMetricRequestDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UsageMetricRequestDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *UsageMetricRequestDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UsageMetricRequestDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UsageMetricRequestDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetInterval

`func (o *UsageMetricRequestDto) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UsageMetricRequestDto) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UsageMetricRequestDto) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetClientId

`func (o *UsageMetricRequestDto) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *UsageMetricRequestDto) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *UsageMetricRequestDto) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *UsageMetricRequestDto) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


