# UpdatePlanRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | A human-readable name for the plan, visible on dashboards, invoices, and reports. | [optional] 
**ExternalId** | Pointer to **string** | An identifier for the plan as defined by your API. | [optional] 
**Description** | Pointer to **string** | A human-readable description for the plan, visible on dashboards, invoices, and reports. | [optional] 
**Nickname** | Pointer to **string** | An additional identifier for the plan, defined by you, that is *not* visible to customers. If you have different *Monthly Active Users Plan*s for different cohorts of customers, this is a useful place to differentiate those. | [optional] 
**Type** | Pointer to **string** | The type of plan. | [optional] 
**DimensionId** | Pointer to **string** | An identifier of the dimension along which usage is aggregated in this plan. | [optional] 
**Aggregation** | Pointer to **string** | The aggregation along which usage metrics are calculated in this plan, relevant only for usage plans. | [optional] 
**Interval** | Pointer to **string** | The billing interval for the plan. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary key-value pairs to attach to the object that can be useful for controlling functionality inside your API. | [optional] 
**Price** | Pointer to [**CreatePlanRequestDtoPrice**](CreatePlanRequestDtoPrice.md) |  | [optional] 

## Methods

### NewUpdatePlanRequestDto

`func NewUpdatePlanRequestDto() *UpdatePlanRequestDto`

NewUpdatePlanRequestDto instantiates a new UpdatePlanRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePlanRequestDtoWithDefaults

`func NewUpdatePlanRequestDtoWithDefaults() *UpdatePlanRequestDto`

NewUpdatePlanRequestDtoWithDefaults instantiates a new UpdatePlanRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *UpdatePlanRequestDto) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UpdatePlanRequestDto) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UpdatePlanRequestDto) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UpdatePlanRequestDto) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetName

`func (o *UpdatePlanRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePlanRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePlanRequestDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePlanRequestDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdatePlanRequestDto) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdatePlanRequestDto) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdatePlanRequestDto) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdatePlanRequestDto) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePlanRequestDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePlanRequestDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePlanRequestDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePlanRequestDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNickname

`func (o *UpdatePlanRequestDto) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *UpdatePlanRequestDto) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *UpdatePlanRequestDto) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *UpdatePlanRequestDto) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetType

`func (o *UpdatePlanRequestDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdatePlanRequestDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdatePlanRequestDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdatePlanRequestDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDimensionId

`func (o *UpdatePlanRequestDto) GetDimensionId() string`

GetDimensionId returns the DimensionId field if non-nil, zero value otherwise.

### GetDimensionIdOk

`func (o *UpdatePlanRequestDto) GetDimensionIdOk() (*string, bool)`

GetDimensionIdOk returns a tuple with the DimensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionId

`func (o *UpdatePlanRequestDto) SetDimensionId(v string)`

SetDimensionId sets DimensionId field to given value.

### HasDimensionId

`func (o *UpdatePlanRequestDto) HasDimensionId() bool`

HasDimensionId returns a boolean if a field has been set.

### GetAggregation

`func (o *UpdatePlanRequestDto) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *UpdatePlanRequestDto) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *UpdatePlanRequestDto) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *UpdatePlanRequestDto) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetInterval

`func (o *UpdatePlanRequestDto) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UpdatePlanRequestDto) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UpdatePlanRequestDto) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *UpdatePlanRequestDto) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdatePlanRequestDto) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdatePlanRequestDto) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdatePlanRequestDto) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdatePlanRequestDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPrice

`func (o *UpdatePlanRequestDto) GetPrice() CreatePlanRequestDtoPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UpdatePlanRequestDto) GetPriceOk() (*CreatePlanRequestDtoPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UpdatePlanRequestDto) SetPrice(v CreatePlanRequestDtoPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UpdatePlanRequestDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


