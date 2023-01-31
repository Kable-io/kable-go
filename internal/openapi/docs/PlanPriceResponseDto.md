# PlanPriceResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Three-letter ISO currency code | 
**Tiers** | [**[]PriceTierResponseDto**](PriceTierResponseDto.md) | The tiered pricing of the plan | 

## Methods

### NewPlanPriceResponseDto

`func NewPlanPriceResponseDto(currency string, tiers []PriceTierResponseDto, ) *PlanPriceResponseDto`

NewPlanPriceResponseDto instantiates a new PlanPriceResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanPriceResponseDtoWithDefaults

`func NewPlanPriceResponseDtoWithDefaults() *PlanPriceResponseDto`

NewPlanPriceResponseDtoWithDefaults instantiates a new PlanPriceResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *PlanPriceResponseDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PlanPriceResponseDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PlanPriceResponseDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTiers

`func (o *PlanPriceResponseDto) GetTiers() []PriceTierResponseDto`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *PlanPriceResponseDto) GetTiersOk() (*[]PriceTierResponseDto, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *PlanPriceResponseDto) SetTiers(v []PriceTierResponseDto)`

SetTiers sets Tiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


