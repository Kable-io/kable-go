# InvoicePriceResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Three-letter ISO currency code | 
**Amount** | **float32** | The total price amount. | 

## Methods

### NewInvoicePriceResponseDto

`func NewInvoicePriceResponseDto(currency string, amount float32, ) *InvoicePriceResponseDto`

NewInvoicePriceResponseDto instantiates a new InvoicePriceResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicePriceResponseDtoWithDefaults

`func NewInvoicePriceResponseDtoWithDefaults() *InvoicePriceResponseDto`

NewInvoicePriceResponseDtoWithDefaults instantiates a new InvoicePriceResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *InvoicePriceResponseDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoicePriceResponseDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoicePriceResponseDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAmount

`func (o *InvoicePriceResponseDto) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoicePriceResponseDto) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoicePriceResponseDto) SetAmount(v float32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


