# PlanPriceRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Three-letter currency code. | 
**Tiers** | **[]string** | Price tiers that define price for usage-based plans. Subscription plans will always have prices with a single tier. | 

## Methods

### NewPlanPriceRequestDto

`func NewPlanPriceRequestDto(currency string, tiers []string, ) *PlanPriceRequestDto`

NewPlanPriceRequestDto instantiates a new PlanPriceRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanPriceRequestDtoWithDefaults

`func NewPlanPriceRequestDtoWithDefaults() *PlanPriceRequestDto`

NewPlanPriceRequestDtoWithDefaults instantiates a new PlanPriceRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *PlanPriceRequestDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PlanPriceRequestDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PlanPriceRequestDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTiers

`func (o *PlanPriceRequestDto) GetTiers() []string`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *PlanPriceRequestDto) GetTiersOk() (*[]string, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *PlanPriceRequestDto) SetTiers(v []string)`

SetTiers sets Tiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


