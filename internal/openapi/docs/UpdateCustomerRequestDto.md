# UpdateCustomerRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | Pointer to **string** | The name of the customer, visible on dashboards, invoices, and reports. | [optional] 
**Currency** | Pointer to **string** | The currency with which this customer pays for plans. | [optional] 
**StripeCustomerId** | Pointer to **string** | The customer&#39;s Stripe &#x60;customer_id&#x60;, if any, for automatic invoice processing through Stripe. | [optional] 
**BillingContact** | Pointer to [**CustomerResponseDtoBillingContact**](CustomerResponseDtoBillingContact.md) |  | [optional] 

## Methods

### NewUpdateCustomerRequestDto

`func NewUpdateCustomerRequestDto() *UpdateCustomerRequestDto`

NewUpdateCustomerRequestDto instantiates a new UpdateCustomerRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomerRequestDtoWithDefaults

`func NewUpdateCustomerRequestDtoWithDefaults() *UpdateCustomerRequestDto`

NewUpdateCustomerRequestDtoWithDefaults instantiates a new UpdateCustomerRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *UpdateCustomerRequestDto) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UpdateCustomerRequestDto) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UpdateCustomerRequestDto) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UpdateCustomerRequestDto) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdateCustomerRequestDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateCustomerRequestDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateCustomerRequestDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdateCustomerRequestDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetStripeCustomerId

`func (o *UpdateCustomerRequestDto) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *UpdateCustomerRequestDto) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *UpdateCustomerRequestDto) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *UpdateCustomerRequestDto) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.

### GetBillingContact

`func (o *UpdateCustomerRequestDto) GetBillingContact() CustomerResponseDtoBillingContact`

GetBillingContact returns the BillingContact field if non-nil, zero value otherwise.

### GetBillingContactOk

`func (o *UpdateCustomerRequestDto) GetBillingContactOk() (*CustomerResponseDtoBillingContact, bool)`

GetBillingContactOk returns a tuple with the BillingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingContact

`func (o *UpdateCustomerRequestDto) SetBillingContact(v CustomerResponseDtoBillingContact)`

SetBillingContact sets BillingContact field to given value.

### HasBillingContact

`func (o *UpdateCustomerRequestDto) HasBillingContact() bool`

HasBillingContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


