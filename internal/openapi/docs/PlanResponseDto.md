# PlanResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** | A Kable-defined identifier for the plan. | 
**Name** | **string** | A human-readable name for the plan, visible on dashboards, invoices, and reports. | 
**ExternalId** | Pointer to **string** | An identifier for the plan as defined by your API. | [optional] 
**Description** | Pointer to **string** | A human-readable description for the plan, visible on dashboards, invoices, and reports. | [optional] 
**Nickname** | Pointer to **string** | An additional identifier for the plan, defined by you, that is *not* visible to customers. If you have different *Monthly Active Users Plan*s for different cohorts of customers, this is a useful place to differentiate those. | [optional] 
**Type** | **string** | The type of plan. | 
**DimensionId** | **string** | An identifier of the dimension along which usage is aggregated in this plan, relevant only for usage plans. | 
**Aggregation** | Pointer to **string** | The aggregation along which usage metrics are calculated in this plan, relevant only for usage plans. | [optional] 
**Interval** | **string** | The billing interval for the plan. | 
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary key-value pairs to attach to the object that can be useful for controlling functionality inside your API. | [optional] 
**Price** | [**PlanResponseDtoPrice**](PlanResponseDtoPrice.md) |  | 
**Entitlements** | [**[]ResourceEntitlementResponseDto**](ResourceEntitlementResponseDto.md) | The entitlements associated with this plan. | 
**Created** | **time.Time** | Date at which entity was first created | 
**Modified** | **time.Time** | Date at which entity was last modified | 

## Methods

### NewPlanResponseDto

`func NewPlanResponseDto(planId string, name string, type_ string, dimensionId string, interval string, price PlanResponseDtoPrice, entitlements []ResourceEntitlementResponseDto, created time.Time, modified time.Time, ) *PlanResponseDto`

NewPlanResponseDto instantiates a new PlanResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanResponseDtoWithDefaults

`func NewPlanResponseDtoWithDefaults() *PlanResponseDto`

NewPlanResponseDtoWithDefaults instantiates a new PlanResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *PlanResponseDto) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PlanResponseDto) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PlanResponseDto) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetName

`func (o *PlanResponseDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanResponseDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanResponseDto) SetName(v string)`

SetName sets Name field to given value.


### GetExternalId

`func (o *PlanResponseDto) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PlanResponseDto) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PlanResponseDto) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PlanResponseDto) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDescription

`func (o *PlanResponseDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanResponseDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanResponseDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlanResponseDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNickname

`func (o *PlanResponseDto) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *PlanResponseDto) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *PlanResponseDto) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *PlanResponseDto) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetType

`func (o *PlanResponseDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlanResponseDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlanResponseDto) SetType(v string)`

SetType sets Type field to given value.


### GetDimensionId

`func (o *PlanResponseDto) GetDimensionId() string`

GetDimensionId returns the DimensionId field if non-nil, zero value otherwise.

### GetDimensionIdOk

`func (o *PlanResponseDto) GetDimensionIdOk() (*string, bool)`

GetDimensionIdOk returns a tuple with the DimensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionId

`func (o *PlanResponseDto) SetDimensionId(v string)`

SetDimensionId sets DimensionId field to given value.


### GetAggregation

`func (o *PlanResponseDto) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *PlanResponseDto) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *PlanResponseDto) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *PlanResponseDto) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetInterval

`func (o *PlanResponseDto) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PlanResponseDto) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PlanResponseDto) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetMetadata

`func (o *PlanResponseDto) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PlanResponseDto) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PlanResponseDto) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PlanResponseDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPrice

`func (o *PlanResponseDto) GetPrice() PlanResponseDtoPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PlanResponseDto) GetPriceOk() (*PlanResponseDtoPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PlanResponseDto) SetPrice(v PlanResponseDtoPrice)`

SetPrice sets Price field to given value.


### GetEntitlements

`func (o *PlanResponseDto) GetEntitlements() []ResourceEntitlementResponseDto`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *PlanResponseDto) GetEntitlementsOk() (*[]ResourceEntitlementResponseDto, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *PlanResponseDto) SetEntitlements(v []ResourceEntitlementResponseDto)`

SetEntitlements sets Entitlements field to given value.


### GetCreated

`func (o *PlanResponseDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PlanResponseDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PlanResponseDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *PlanResponseDto) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *PlanResponseDto) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *PlanResponseDto) SetModified(v time.Time)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


