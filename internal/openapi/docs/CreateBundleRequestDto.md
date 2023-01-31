# CreateBundleRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | **string** |  | 
**Name** | **string** |  | 
**ExternalId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**PlanIds** | Pointer to **[]string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CostPerCredit** | Pointer to **float32** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateBundleRequestDto

`func NewCreateBundleRequestDto(bundleId string, name string, ) *CreateBundleRequestDto`

NewCreateBundleRequestDto instantiates a new CreateBundleRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBundleRequestDtoWithDefaults

`func NewCreateBundleRequestDtoWithDefaults() *CreateBundleRequestDto`

NewCreateBundleRequestDtoWithDefaults instantiates a new CreateBundleRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *CreateBundleRequestDto) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *CreateBundleRequestDto) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *CreateBundleRequestDto) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.


### GetName

`func (o *CreateBundleRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBundleRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBundleRequestDto) SetName(v string)`

SetName sets Name field to given value.


### GetExternalId

`func (o *CreateBundleRequestDto) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateBundleRequestDto) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateBundleRequestDto) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateBundleRequestDto) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDescription

`func (o *CreateBundleRequestDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateBundleRequestDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateBundleRequestDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateBundleRequestDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNickname

`func (o *CreateBundleRequestDto) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *CreateBundleRequestDto) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *CreateBundleRequestDto) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *CreateBundleRequestDto) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPlanIds

`func (o *CreateBundleRequestDto) GetPlanIds() []string`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *CreateBundleRequestDto) GetPlanIdsOk() (*[]string, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *CreateBundleRequestDto) SetPlanIds(v []string)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *CreateBundleRequestDto) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetCurrency

`func (o *CreateBundleRequestDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateBundleRequestDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateBundleRequestDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateBundleRequestDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCostPerCredit

`func (o *CreateBundleRequestDto) GetCostPerCredit() float32`

GetCostPerCredit returns the CostPerCredit field if non-nil, zero value otherwise.

### GetCostPerCreditOk

`func (o *CreateBundleRequestDto) GetCostPerCreditOk() (*float32, bool)`

GetCostPerCreditOk returns a tuple with the CostPerCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerCredit

`func (o *CreateBundleRequestDto) SetCostPerCredit(v float32)`

SetCostPerCredit sets CostPerCredit field to given value.

### HasCostPerCredit

`func (o *CreateBundleRequestDto) HasCostPerCredit() bool`

HasCostPerCredit returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateBundleRequestDto) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateBundleRequestDto) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateBundleRequestDto) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateBundleRequestDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


