# UpdateCustomerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | Pointer to **string** |  | [optional] 
**StripeCustomerId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateCustomerRequest

`func NewUpdateCustomerRequest() *UpdateCustomerRequest`

NewUpdateCustomerRequest instantiates a new UpdateCustomerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomerRequestWithDefaults

`func NewUpdateCustomerRequestWithDefaults() *UpdateCustomerRequest`

NewUpdateCustomerRequestWithDefaults instantiates a new UpdateCustomerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *UpdateCustomerRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UpdateCustomerRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UpdateCustomerRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UpdateCustomerRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetStripeCustomerId

`func (o *UpdateCustomerRequest) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *UpdateCustomerRequest) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *UpdateCustomerRequest) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *UpdateCustomerRequest) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


