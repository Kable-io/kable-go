# CreditGrantResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditGrantId** | **string** | A Kable-defined identifier for the credit grant. | 
**CustomerId** | **string** | The customer for whom the credit was granted. | 
**GrantAmount** | **float32** | The number of credits granted. | 
**BalanceAmount** | **float32** | The number of credits available. | 
**PaidAmount** | Pointer to **float32** | The amount paid for the grant. | [optional] 
**Currency** | **string** | Three-letter ISO currency code | 
**Description** | Pointer to **string** | A human-readable description that explains why the credit grant was created. | [optional] 
**Expiration** | Pointer to **time.Time** | The expiration date of the credit grant. | [optional] 

## Methods

### NewCreditGrantResponseDto

`func NewCreditGrantResponseDto(creditGrantId string, customerId string, grantAmount float32, balanceAmount float32, currency string, ) *CreditGrantResponseDto`

NewCreditGrantResponseDto instantiates a new CreditGrantResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditGrantResponseDtoWithDefaults

`func NewCreditGrantResponseDtoWithDefaults() *CreditGrantResponseDto`

NewCreditGrantResponseDtoWithDefaults instantiates a new CreditGrantResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditGrantId

`func (o *CreditGrantResponseDto) GetCreditGrantId() string`

GetCreditGrantId returns the CreditGrantId field if non-nil, zero value otherwise.

### GetCreditGrantIdOk

`func (o *CreditGrantResponseDto) GetCreditGrantIdOk() (*string, bool)`

GetCreditGrantIdOk returns a tuple with the CreditGrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditGrantId

`func (o *CreditGrantResponseDto) SetCreditGrantId(v string)`

SetCreditGrantId sets CreditGrantId field to given value.


### GetCustomerId

`func (o *CreditGrantResponseDto) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreditGrantResponseDto) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreditGrantResponseDto) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetGrantAmount

`func (o *CreditGrantResponseDto) GetGrantAmount() float32`

GetGrantAmount returns the GrantAmount field if non-nil, zero value otherwise.

### GetGrantAmountOk

`func (o *CreditGrantResponseDto) GetGrantAmountOk() (*float32, bool)`

GetGrantAmountOk returns a tuple with the GrantAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantAmount

`func (o *CreditGrantResponseDto) SetGrantAmount(v float32)`

SetGrantAmount sets GrantAmount field to given value.


### GetBalanceAmount

`func (o *CreditGrantResponseDto) GetBalanceAmount() float32`

GetBalanceAmount returns the BalanceAmount field if non-nil, zero value otherwise.

### GetBalanceAmountOk

`func (o *CreditGrantResponseDto) GetBalanceAmountOk() (*float32, bool)`

GetBalanceAmountOk returns a tuple with the BalanceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAmount

`func (o *CreditGrantResponseDto) SetBalanceAmount(v float32)`

SetBalanceAmount sets BalanceAmount field to given value.


### GetPaidAmount

`func (o *CreditGrantResponseDto) GetPaidAmount() float32`

GetPaidAmount returns the PaidAmount field if non-nil, zero value otherwise.

### GetPaidAmountOk

`func (o *CreditGrantResponseDto) GetPaidAmountOk() (*float32, bool)`

GetPaidAmountOk returns a tuple with the PaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAmount

`func (o *CreditGrantResponseDto) SetPaidAmount(v float32)`

SetPaidAmount sets PaidAmount field to given value.

### HasPaidAmount

`func (o *CreditGrantResponseDto) HasPaidAmount() bool`

HasPaidAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *CreditGrantResponseDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreditGrantResponseDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreditGrantResponseDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *CreditGrantResponseDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreditGrantResponseDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreditGrantResponseDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreditGrantResponseDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiration

`func (o *CreditGrantResponseDto) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *CreditGrantResponseDto) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *CreditGrantResponseDto) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *CreditGrantResponseDto) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


