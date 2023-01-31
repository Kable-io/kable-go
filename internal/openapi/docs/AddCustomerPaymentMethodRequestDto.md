# AddCustomerPaymentMethodRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessUrl** | **string** | The url to which to redirect after successful payment method creation | 
**CancelUrl** | **string** | The url to which to redirect after failed payment method creation | 

## Methods

### NewAddCustomerPaymentMethodRequestDto

`func NewAddCustomerPaymentMethodRequestDto(successUrl string, cancelUrl string, ) *AddCustomerPaymentMethodRequestDto`

NewAddCustomerPaymentMethodRequestDto instantiates a new AddCustomerPaymentMethodRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCustomerPaymentMethodRequestDtoWithDefaults

`func NewAddCustomerPaymentMethodRequestDtoWithDefaults() *AddCustomerPaymentMethodRequestDto`

NewAddCustomerPaymentMethodRequestDtoWithDefaults instantiates a new AddCustomerPaymentMethodRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessUrl

`func (o *AddCustomerPaymentMethodRequestDto) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *AddCustomerPaymentMethodRequestDto) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *AddCustomerPaymentMethodRequestDto) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.


### GetCancelUrl

`func (o *AddCustomerPaymentMethodRequestDto) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *AddCustomerPaymentMethodRequestDto) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *AddCustomerPaymentMethodRequestDto) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


