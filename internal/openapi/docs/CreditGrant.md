# CreditGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditGrantId** | **string** | The identifier of the credit grant created on Kable. | 
**GrantAmount** | **float32** | The number of credits created in the grant. | 
**BalanceAmount** | **float32** | The available, non-consumed balance remaining in the grant. | 
**PaidAmount** | Pointer to **float32** | The amount paid (in the specified &#x60;currency&#x60;), if any, for the credit grant. | [optional] 
**Currency** | **string** | Three-letter currency code. | 

## Methods

### NewCreditGrant

`func NewCreditGrant(creditGrantId string, grantAmount float32, balanceAmount float32, currency string, ) *CreditGrant`

NewCreditGrant instantiates a new CreditGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditGrantWithDefaults

`func NewCreditGrantWithDefaults() *CreditGrant`

NewCreditGrantWithDefaults instantiates a new CreditGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditGrantId

`func (o *CreditGrant) GetCreditGrantId() string`

GetCreditGrantId returns the CreditGrantId field if non-nil, zero value otherwise.

### GetCreditGrantIdOk

`func (o *CreditGrant) GetCreditGrantIdOk() (*string, bool)`

GetCreditGrantIdOk returns a tuple with the CreditGrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditGrantId

`func (o *CreditGrant) SetCreditGrantId(v string)`

SetCreditGrantId sets CreditGrantId field to given value.


### GetGrantAmount

`func (o *CreditGrant) GetGrantAmount() float32`

GetGrantAmount returns the GrantAmount field if non-nil, zero value otherwise.

### GetGrantAmountOk

`func (o *CreditGrant) GetGrantAmountOk() (*float32, bool)`

GetGrantAmountOk returns a tuple with the GrantAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantAmount

`func (o *CreditGrant) SetGrantAmount(v float32)`

SetGrantAmount sets GrantAmount field to given value.


### GetBalanceAmount

`func (o *CreditGrant) GetBalanceAmount() float32`

GetBalanceAmount returns the BalanceAmount field if non-nil, zero value otherwise.

### GetBalanceAmountOk

`func (o *CreditGrant) GetBalanceAmountOk() (*float32, bool)`

GetBalanceAmountOk returns a tuple with the BalanceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAmount

`func (o *CreditGrant) SetBalanceAmount(v float32)`

SetBalanceAmount sets BalanceAmount field to given value.


### GetPaidAmount

`func (o *CreditGrant) GetPaidAmount() float32`

GetPaidAmount returns the PaidAmount field if non-nil, zero value otherwise.

### GetPaidAmountOk

`func (o *CreditGrant) GetPaidAmountOk() (*float32, bool)`

GetPaidAmountOk returns a tuple with the PaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAmount

`func (o *CreditGrant) SetPaidAmount(v float32)`

SetPaidAmount sets PaidAmount field to given value.

### HasPaidAmount

`func (o *CreditGrant) HasPaidAmount() bool`

HasPaidAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *CreditGrant) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreditGrant) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreditGrant) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


