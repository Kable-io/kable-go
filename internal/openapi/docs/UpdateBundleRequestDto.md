# UpdateBundleRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**PlanIds** | Pointer to **[]string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CostPerCredit** | Pointer to **float32** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateBundleRequestDto

`func NewUpdateBundleRequestDto() *UpdateBundleRequestDto`

NewUpdateBundleRequestDto instantiates a new UpdateBundleRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBundleRequestDtoWithDefaults

`func NewUpdateBundleRequestDtoWithDefaults() *UpdateBundleRequestDto`

NewUpdateBundleRequestDtoWithDefaults instantiates a new UpdateBundleRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *UpdateBundleRequestDto) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *UpdateBundleRequestDto) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *UpdateBundleRequestDto) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *UpdateBundleRequestDto) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetName

`func (o *UpdateBundleRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBundleRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBundleRequestDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateBundleRequestDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateBundleRequestDto) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateBundleRequestDto) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateBundleRequestDto) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateBundleRequestDto) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateBundleRequestDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateBundleRequestDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateBundleRequestDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateBundleRequestDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNickname

`func (o *UpdateBundleRequestDto) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *UpdateBundleRequestDto) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *UpdateBundleRequestDto) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *UpdateBundleRequestDto) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPlanIds

`func (o *UpdateBundleRequestDto) GetPlanIds() []string`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *UpdateBundleRequestDto) GetPlanIdsOk() (*[]string, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *UpdateBundleRequestDto) SetPlanIds(v []string)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *UpdateBundleRequestDto) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdateBundleRequestDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateBundleRequestDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateBundleRequestDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdateBundleRequestDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCostPerCredit

`func (o *UpdateBundleRequestDto) GetCostPerCredit() float32`

GetCostPerCredit returns the CostPerCredit field if non-nil, zero value otherwise.

### GetCostPerCreditOk

`func (o *UpdateBundleRequestDto) GetCostPerCreditOk() (*float32, bool)`

GetCostPerCreditOk returns a tuple with the CostPerCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerCredit

`func (o *UpdateBundleRequestDto) SetCostPerCredit(v float32)`

SetCostPerCredit sets CostPerCredit field to given value.

### HasCostPerCredit

`func (o *UpdateBundleRequestDto) HasCostPerCredit() bool`

HasCostPerCredit returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateBundleRequestDto) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateBundleRequestDto) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateBundleRequestDto) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateBundleRequestDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


