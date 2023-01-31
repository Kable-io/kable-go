# CustomerPlanResponseDto

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
**AddedDate** | **time.Time** |  | 
**EndedDate** | Pointer to **time.Time** |  | [optional] 
**PeriodStartDate** | Pointer to **time.Time** |  | [optional] 
**PeriodEndDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCustomerPlanResponseDto

`func NewCustomerPlanResponseDto(planId string, name string, type_ string, dimensionId string, interval string, price PlanResponseDtoPrice, entitlements []ResourceEntitlementResponseDto, created time.Time, modified time.Time, addedDate time.Time, ) *CustomerPlanResponseDto`

NewCustomerPlanResponseDto instantiates a new CustomerPlanResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPlanResponseDtoWithDefaults

`func NewCustomerPlanResponseDtoWithDefaults() *CustomerPlanResponseDto`

NewCustomerPlanResponseDtoWithDefaults instantiates a new CustomerPlanResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *CustomerPlanResponseDto) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CustomerPlanResponseDto) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CustomerPlanResponseDto) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetName

`func (o *CustomerPlanResponseDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerPlanResponseDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerPlanResponseDto) SetName(v string)`

SetName sets Name field to given value.


### GetExternalId

`func (o *CustomerPlanResponseDto) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CustomerPlanResponseDto) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CustomerPlanResponseDto) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CustomerPlanResponseDto) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDescription

`func (o *CustomerPlanResponseDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomerPlanResponseDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomerPlanResponseDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomerPlanResponseDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNickname

`func (o *CustomerPlanResponseDto) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *CustomerPlanResponseDto) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *CustomerPlanResponseDto) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *CustomerPlanResponseDto) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetType

`func (o *CustomerPlanResponseDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerPlanResponseDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerPlanResponseDto) SetType(v string)`

SetType sets Type field to given value.


### GetDimensionId

`func (o *CustomerPlanResponseDto) GetDimensionId() string`

GetDimensionId returns the DimensionId field if non-nil, zero value otherwise.

### GetDimensionIdOk

`func (o *CustomerPlanResponseDto) GetDimensionIdOk() (*string, bool)`

GetDimensionIdOk returns a tuple with the DimensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionId

`func (o *CustomerPlanResponseDto) SetDimensionId(v string)`

SetDimensionId sets DimensionId field to given value.


### GetAggregation

`func (o *CustomerPlanResponseDto) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *CustomerPlanResponseDto) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *CustomerPlanResponseDto) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *CustomerPlanResponseDto) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetInterval

`func (o *CustomerPlanResponseDto) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CustomerPlanResponseDto) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CustomerPlanResponseDto) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetMetadata

`func (o *CustomerPlanResponseDto) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerPlanResponseDto) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerPlanResponseDto) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerPlanResponseDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPrice

`func (o *CustomerPlanResponseDto) GetPrice() PlanResponseDtoPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CustomerPlanResponseDto) GetPriceOk() (*PlanResponseDtoPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CustomerPlanResponseDto) SetPrice(v PlanResponseDtoPrice)`

SetPrice sets Price field to given value.


### GetEntitlements

`func (o *CustomerPlanResponseDto) GetEntitlements() []ResourceEntitlementResponseDto`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *CustomerPlanResponseDto) GetEntitlementsOk() (*[]ResourceEntitlementResponseDto, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *CustomerPlanResponseDto) SetEntitlements(v []ResourceEntitlementResponseDto)`

SetEntitlements sets Entitlements field to given value.


### GetCreated

`func (o *CustomerPlanResponseDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomerPlanResponseDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomerPlanResponseDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *CustomerPlanResponseDto) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CustomerPlanResponseDto) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CustomerPlanResponseDto) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetAddedDate

`func (o *CustomerPlanResponseDto) GetAddedDate() time.Time`

GetAddedDate returns the AddedDate field if non-nil, zero value otherwise.

### GetAddedDateOk

`func (o *CustomerPlanResponseDto) GetAddedDateOk() (*time.Time, bool)`

GetAddedDateOk returns a tuple with the AddedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedDate

`func (o *CustomerPlanResponseDto) SetAddedDate(v time.Time)`

SetAddedDate sets AddedDate field to given value.


### GetEndedDate

`func (o *CustomerPlanResponseDto) GetEndedDate() time.Time`

GetEndedDate returns the EndedDate field if non-nil, zero value otherwise.

### GetEndedDateOk

`func (o *CustomerPlanResponseDto) GetEndedDateOk() (*time.Time, bool)`

GetEndedDateOk returns a tuple with the EndedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedDate

`func (o *CustomerPlanResponseDto) SetEndedDate(v time.Time)`

SetEndedDate sets EndedDate field to given value.

### HasEndedDate

`func (o *CustomerPlanResponseDto) HasEndedDate() bool`

HasEndedDate returns a boolean if a field has been set.

### GetPeriodStartDate

`func (o *CustomerPlanResponseDto) GetPeriodStartDate() time.Time`

GetPeriodStartDate returns the PeriodStartDate field if non-nil, zero value otherwise.

### GetPeriodStartDateOk

`func (o *CustomerPlanResponseDto) GetPeriodStartDateOk() (*time.Time, bool)`

GetPeriodStartDateOk returns a tuple with the PeriodStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStartDate

`func (o *CustomerPlanResponseDto) SetPeriodStartDate(v time.Time)`

SetPeriodStartDate sets PeriodStartDate field to given value.

### HasPeriodStartDate

`func (o *CustomerPlanResponseDto) HasPeriodStartDate() bool`

HasPeriodStartDate returns a boolean if a field has been set.

### GetPeriodEndDate

`func (o *CustomerPlanResponseDto) GetPeriodEndDate() time.Time`

GetPeriodEndDate returns the PeriodEndDate field if non-nil, zero value otherwise.

### GetPeriodEndDateOk

`func (o *CustomerPlanResponseDto) GetPeriodEndDateOk() (*time.Time, bool)`

GetPeriodEndDateOk returns a tuple with the PeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEndDate

`func (o *CustomerPlanResponseDto) SetPeriodEndDate(v time.Time)`

SetPeriodEndDate sets PeriodEndDate field to given value.

### HasPeriodEndDate

`func (o *CustomerPlanResponseDto) HasPeriodEndDate() bool`

HasPeriodEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


