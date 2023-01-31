# PlanResponseDtoPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Three-letter ISO currency code | 
**Tiers** | [**[]PriceTierResponseDto**](PriceTierResponseDto.md) | The tiered pricing of the plan | 

## Methods

### NewPlanResponseDtoPrice

`func NewPlanResponseDtoPrice(currency string, tiers []PriceTierResponseDto, ) *PlanResponseDtoPrice`

NewPlanResponseDtoPrice instantiates a new PlanResponseDtoPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanResponseDtoPriceWithDefaults

`func NewPlanResponseDtoPriceWithDefaults() *PlanResponseDtoPrice`

NewPlanResponseDtoPriceWithDefaults instantiates a new PlanResponseDtoPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *PlanResponseDtoPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PlanResponseDtoPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PlanResponseDtoPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTiers

`func (o *PlanResponseDtoPrice) GetTiers() []PriceTierResponseDto`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *PlanResponseDtoPrice) GetTiersOk() (*[]PriceTierResponseDto, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *PlanResponseDtoPrice) SetTiers(v []PriceTierResponseDto)`

SetTiers sets Tiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


