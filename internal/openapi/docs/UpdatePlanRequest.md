# UpdatePlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A human-readable name for the plan, visible on dashboards, invoices, and reports. | [optional] 
**ExternalId** | Pointer to **string** | An identifier for the plan as defined by your API. | [optional] 
**Description** | Pointer to **string** | A human-readable description for the plan, visible on dashboards, invoices, and reports. | [optional] 
**Nickname** | Pointer to **string** | An additional identifier for the plan, defined by you, that is *not* visible to customers. If you have different *Monthly Active Users Plan*s for different cohorts of customers, this is a useful place to differentiate those. | [optional] 
**Type** | Pointer to **string** | The type of plan. | [optional] 
**Interval** | Pointer to **string** | The billing interval for the plan. | [optional] 
**DimensionId** | Pointer to **string** | An identifier of the dimension along which usage is aggregated in this plan. | [optional] 
**Aggregation** | Pointer to **string** | The aggregation along which usage metrics are calculated in this plan, relevant only for usage plans. | [optional] 
**Price** | Pointer to [**Price**](Price.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary key-value pairs to attach to the object that can be useful for controlling functionality inside your API. | [optional] 

## Methods

### NewUpdatePlanRequest

`func NewUpdatePlanRequest() *UpdatePlanRequest`

NewUpdatePlanRequest instantiates a new UpdatePlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePlanRequestWithDefaults

`func NewUpdatePlanRequestWithDefaults() *UpdatePlanRequest`

NewUpdatePlanRequestWithDefaults instantiates a new UpdatePlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePlanRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePlanRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdatePlanRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdatePlanRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdatePlanRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdatePlanRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePlanRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePlanRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePlanRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePlanRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNickname

`func (o *UpdatePlanRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *UpdatePlanRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *UpdatePlanRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *UpdatePlanRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetType

`func (o *UpdatePlanRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdatePlanRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdatePlanRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdatePlanRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInterval

`func (o *UpdatePlanRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UpdatePlanRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UpdatePlanRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *UpdatePlanRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetDimensionId

`func (o *UpdatePlanRequest) GetDimensionId() string`

GetDimensionId returns the DimensionId field if non-nil, zero value otherwise.

### GetDimensionIdOk

`func (o *UpdatePlanRequest) GetDimensionIdOk() (*string, bool)`

GetDimensionIdOk returns a tuple with the DimensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionId

`func (o *UpdatePlanRequest) SetDimensionId(v string)`

SetDimensionId sets DimensionId field to given value.

### HasDimensionId

`func (o *UpdatePlanRequest) HasDimensionId() bool`

HasDimensionId returns a boolean if a field has been set.

### GetAggregation

`func (o *UpdatePlanRequest) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *UpdatePlanRequest) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *UpdatePlanRequest) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *UpdatePlanRequest) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetPrice

`func (o *UpdatePlanRequest) GetPrice() Price`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UpdatePlanRequest) GetPriceOk() (*Price, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UpdatePlanRequest) SetPrice(v Price)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UpdatePlanRequest) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdatePlanRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdatePlanRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdatePlanRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdatePlanRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


