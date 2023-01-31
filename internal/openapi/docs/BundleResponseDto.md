# BundleResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | **string** | A Kable-defined identifier for the bundle. | 
**Name** | **string** | A human-readable name for the bundle, visible on dashboards, invoices, and reports. | 
**ExternalId** | Pointer to **string** | An identifier for the bundle as defined by your API. | [optional] 
**Description** | Pointer to **string** | A human-readable description for the bundle, visible on dashboards, invoices, and reports. | [optional] 
**Nickname** | Pointer to **string** | An additional identifier for the bundle, defined by you, that is *not* visible to customers. If you have different *Monthly Active Users Bundle*s for different cohorts of customers, this is a useful place to differentiate those. | [optional] 
**Currency** | **string** | Three-letter ISO currency code | 
**Plans** | Pointer to [**[]PlanResponseDto**](PlanResponseDto.md) | The plans included in this bundle. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary key-value pairs to attach to the object that can be useful for controlling functionality inside your API. | [optional] 
**Created** | **time.Time** | Date at which entity was first created | 
**Modified** | **time.Time** | Date at which entity was last modified | 

## Methods

### NewBundleResponseDto

`func NewBundleResponseDto(bundleId string, name string, currency string, created time.Time, modified time.Time, ) *BundleResponseDto`

NewBundleResponseDto instantiates a new BundleResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleResponseDtoWithDefaults

`func NewBundleResponseDtoWithDefaults() *BundleResponseDto`

NewBundleResponseDtoWithDefaults instantiates a new BundleResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *BundleResponseDto) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *BundleResponseDto) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *BundleResponseDto) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.


### GetName

`func (o *BundleResponseDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BundleResponseDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BundleResponseDto) SetName(v string)`

SetName sets Name field to given value.


### GetExternalId

`func (o *BundleResponseDto) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BundleResponseDto) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BundleResponseDto) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *BundleResponseDto) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDescription

`func (o *BundleResponseDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BundleResponseDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BundleResponseDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BundleResponseDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNickname

`func (o *BundleResponseDto) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *BundleResponseDto) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *BundleResponseDto) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *BundleResponseDto) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetCurrency

`func (o *BundleResponseDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BundleResponseDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BundleResponseDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPlans

`func (o *BundleResponseDto) GetPlans() []PlanResponseDto`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *BundleResponseDto) GetPlansOk() (*[]PlanResponseDto, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *BundleResponseDto) SetPlans(v []PlanResponseDto)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *BundleResponseDto) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetMetadata

`func (o *BundleResponseDto) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BundleResponseDto) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BundleResponseDto) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BundleResponseDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreated

`func (o *BundleResponseDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BundleResponseDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BundleResponseDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *BundleResponseDto) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *BundleResponseDto) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *BundleResponseDto) SetModified(v time.Time)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


