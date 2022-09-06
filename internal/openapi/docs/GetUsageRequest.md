# GetUsageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | **string** | Valid values are &#x60;COUNT&#x60;, &#x60;COUNT_DISTINCT&#x60;, and &#x60;SUM&#x60;. | 
**DimensionId** | Pointer to **string** | Note that one of either &#x60;dimensionId&#x60; or &#x60;dimensionKey&#x60; is required, though either may be provided. | [optional] 
**DimensionKey** | Pointer to **string** | Note that one of either &#x60;dimensionId&#x60; or &#x60;dimensionKey&#x60; is required, though either may be provided. | [optional] 
**StartDate** | **time.Time** | Start of the period over which to query. Timestamps must be formatted as RFC 3339 strings like &#x60;2022-01-09T09:32:01Z&#x60; | 
**EndDate** | **time.Time** | End of the period over which to query. Timestamps must be formatted as RFC 3339 strings like &#x60;2022-01-09T09:32:01Z&#x60; | 
**Interval** | **string** | Valid values are &#x60;HOUR&#x60;, &#x60;DAY&#x60;, and &#x60;MONTH&#x60;. | 
**ClientId** | Pointer to **string** | When provided, query will consider only this customer. When &#x60;null&#x60;, query will consider all of your customers. | [optional] 

## Methods

### NewGetUsageRequest

`func NewGetUsageRequest(aggregation string, startDate time.Time, endDate time.Time, interval string, ) *GetUsageRequest`

NewGetUsageRequest instantiates a new GetUsageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsageRequestWithDefaults

`func NewGetUsageRequestWithDefaults() *GetUsageRequest`

NewGetUsageRequestWithDefaults instantiates a new GetUsageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *GetUsageRequest) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *GetUsageRequest) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *GetUsageRequest) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.


### GetDimensionId

`func (o *GetUsageRequest) GetDimensionId() string`

GetDimensionId returns the DimensionId field if non-nil, zero value otherwise.

### GetDimensionIdOk

`func (o *GetUsageRequest) GetDimensionIdOk() (*string, bool)`

GetDimensionIdOk returns a tuple with the DimensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionId

`func (o *GetUsageRequest) SetDimensionId(v string)`

SetDimensionId sets DimensionId field to given value.

### HasDimensionId

`func (o *GetUsageRequest) HasDimensionId() bool`

HasDimensionId returns a boolean if a field has been set.

### GetDimensionKey

`func (o *GetUsageRequest) GetDimensionKey() string`

GetDimensionKey returns the DimensionKey field if non-nil, zero value otherwise.

### GetDimensionKeyOk

`func (o *GetUsageRequest) GetDimensionKeyOk() (*string, bool)`

GetDimensionKeyOk returns a tuple with the DimensionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionKey

`func (o *GetUsageRequest) SetDimensionKey(v string)`

SetDimensionKey sets DimensionKey field to given value.

### HasDimensionKey

`func (o *GetUsageRequest) HasDimensionKey() bool`

HasDimensionKey returns a boolean if a field has been set.

### GetStartDate

`func (o *GetUsageRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetUsageRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetUsageRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *GetUsageRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetUsageRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetUsageRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetInterval

`func (o *GetUsageRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetUsageRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetUsageRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetClientId

`func (o *GetUsageRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetUsageRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetUsageRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetUsageRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


