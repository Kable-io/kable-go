# CreatePlanRequestDtoPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Three-letter currency code. | 
**Tiers** | **[]string** | Price tiers that define price for usage-based plans. Subscription plans will always have prices with a single tier. | 

## Methods

### NewCreatePlanRequestDtoPrice

`func NewCreatePlanRequestDtoPrice(currency string, tiers []string, ) *CreatePlanRequestDtoPrice`

NewCreatePlanRequestDtoPrice instantiates a new CreatePlanRequestDtoPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePlanRequestDtoPriceWithDefaults

`func NewCreatePlanRequestDtoPriceWithDefaults() *CreatePlanRequestDtoPrice`

NewCreatePlanRequestDtoPriceWithDefaults instantiates a new CreatePlanRequestDtoPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CreatePlanRequestDtoPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreatePlanRequestDtoPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreatePlanRequestDtoPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTiers

`func (o *CreatePlanRequestDtoPrice) GetTiers() []string`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *CreatePlanRequestDtoPrice) GetTiersOk() (*[]string, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *CreatePlanRequestDtoPrice) SetTiers(v []string)`

SetTiers sets Tiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


