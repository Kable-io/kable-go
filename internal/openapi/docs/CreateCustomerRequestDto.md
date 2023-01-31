# CreateCustomerRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | A unique identifier for the customer, defined by you. This usually corresponds to your own internal unique identifier for customers, and is the identifier by which Kable attributes customer usage. | 
**CompanyName** | **string** | The name of the customer, visible on dashboards, invoices, and reports. | 
**Currency** | **string** | Three-letter ISO currency code | 
**StripeCustomerId** | Pointer to **string** | The customer&#39;s Stripe &#x60;customer_id&#x60;, if one already exists, for automatic invoice processing through Stripe. | [optional] 
**PlanIds** | Pointer to **[]string** | A collection of plans to which to subscribe the customer. | [optional] 
**BundleIds** | Pointer to **[]string** | A collection of bundles to which to subscribe the customer. | [optional] 
**BillingContact** | Pointer to [**CustomerResponseDtoBillingContact**](CustomerResponseDtoBillingContact.md) |  | [optional] 

## Methods

### NewCreateCustomerRequestDto

`func NewCreateCustomerRequestDto(clientId string, companyName string, currency string, ) *CreateCustomerRequestDto`

NewCreateCustomerRequestDto instantiates a new CreateCustomerRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomerRequestDtoWithDefaults

`func NewCreateCustomerRequestDtoWithDefaults() *CreateCustomerRequestDto`

NewCreateCustomerRequestDtoWithDefaults instantiates a new CreateCustomerRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *CreateCustomerRequestDto) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateCustomerRequestDto) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateCustomerRequestDto) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetCompanyName

`func (o *CreateCustomerRequestDto) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CreateCustomerRequestDto) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CreateCustomerRequestDto) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetCurrency

`func (o *CreateCustomerRequestDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateCustomerRequestDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateCustomerRequestDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetStripeCustomerId

`func (o *CreateCustomerRequestDto) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *CreateCustomerRequestDto) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *CreateCustomerRequestDto) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *CreateCustomerRequestDto) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.

### GetPlanIds

`func (o *CreateCustomerRequestDto) GetPlanIds() []string`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *CreateCustomerRequestDto) GetPlanIdsOk() (*[]string, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *CreateCustomerRequestDto) SetPlanIds(v []string)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *CreateCustomerRequestDto) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetBundleIds

`func (o *CreateCustomerRequestDto) GetBundleIds() []string`

GetBundleIds returns the BundleIds field if non-nil, zero value otherwise.

### GetBundleIdsOk

`func (o *CreateCustomerRequestDto) GetBundleIdsOk() (*[]string, bool)`

GetBundleIdsOk returns a tuple with the BundleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleIds

`func (o *CreateCustomerRequestDto) SetBundleIds(v []string)`

SetBundleIds sets BundleIds field to given value.

### HasBundleIds

`func (o *CreateCustomerRequestDto) HasBundleIds() bool`

HasBundleIds returns a boolean if a field has been set.

### GetBillingContact

`func (o *CreateCustomerRequestDto) GetBillingContact() CustomerResponseDtoBillingContact`

GetBillingContact returns the BillingContact field if non-nil, zero value otherwise.

### GetBillingContactOk

`func (o *CreateCustomerRequestDto) GetBillingContactOk() (*CustomerResponseDtoBillingContact, bool)`

GetBillingContactOk returns a tuple with the BillingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingContact

`func (o *CreateCustomerRequestDto) SetBillingContact(v CustomerResponseDtoBillingContact)`

SetBillingContact sets BillingContact field to given value.

### HasBillingContact

`func (o *CreateCustomerRequestDto) HasBillingContact() bool`

HasBillingContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


