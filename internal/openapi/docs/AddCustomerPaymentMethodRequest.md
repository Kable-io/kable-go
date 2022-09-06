# AddCustomerPaymentMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessUrl** | **string** |  | 
**CancelUrl** | **string** |  | 

## Methods

### NewAddCustomerPaymentMethodRequest

`func NewAddCustomerPaymentMethodRequest(successUrl string, cancelUrl string, ) *AddCustomerPaymentMethodRequest`

NewAddCustomerPaymentMethodRequest instantiates a new AddCustomerPaymentMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCustomerPaymentMethodRequestWithDefaults

`func NewAddCustomerPaymentMethodRequestWithDefaults() *AddCustomerPaymentMethodRequest`

NewAddCustomerPaymentMethodRequestWithDefaults instantiates a new AddCustomerPaymentMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessUrl

`func (o *AddCustomerPaymentMethodRequest) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *AddCustomerPaymentMethodRequest) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *AddCustomerPaymentMethodRequest) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.


### GetCancelUrl

`func (o *AddCustomerPaymentMethodRequest) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *AddCustomerPaymentMethodRequest) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *AddCustomerPaymentMethodRequest) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


