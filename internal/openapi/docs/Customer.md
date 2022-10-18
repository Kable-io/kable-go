# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KableClientId** | Pointer to **string** | Your company&#39;s Kable client ID. | [optional] 
**CustomerId** | **string** | A Kable-defined identifier for the customer. | 
**ClientId** | **string** | A unique identifier for the customer, defined by you. This usually corresponds to your own internal unique identifier for customers, and is the identifier by which Kable attributes customer usage. | 
**CompanyName** | **string** | The name of the customer, visible on dashboards, invoices, and reports. | 
**Currency** | **string** | The three-letter currency code with which this customer pays for plans. | 
**Status** | **string** | The current status of the customer. | 
**NextInvoiceDate** | Pointer to **string** | The date of the customer&#39;s next invoice, if any. | [optional] 
**StripeCustomerId** | Pointer to **string** | The customer&#39;s Stripe &#x60;customer_id&#x60;, if any, for automatic invoice processing through Stripe. | [optional] 
**Plans** | Pointer to [**[]CustomerPlan**](CustomerPlan.md) | A collection of plans to which this customer is currently subscribed. | [optional] 

## Methods

### NewCustomer

`func NewCustomer(customerId string, clientId string, companyName string, currency string, status string, ) *Customer`

NewCustomer instantiates a new Customer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerWithDefaults

`func NewCustomerWithDefaults() *Customer`

NewCustomerWithDefaults instantiates a new Customer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKableClientId

`func (o *Customer) GetKableClientId() string`

GetKableClientId returns the KableClientId field if non-nil, zero value otherwise.

### GetKableClientIdOk

`func (o *Customer) GetKableClientIdOk() (*string, bool)`

GetKableClientIdOk returns a tuple with the KableClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKableClientId

`func (o *Customer) SetKableClientId(v string)`

SetKableClientId sets KableClientId field to given value.

### HasKableClientId

`func (o *Customer) HasKableClientId() bool`

HasKableClientId returns a boolean if a field has been set.

### GetCustomerId

`func (o *Customer) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Customer) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Customer) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetClientId

`func (o *Customer) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Customer) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Customer) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetCompanyName

`func (o *Customer) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Customer) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Customer) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetCurrency

`func (o *Customer) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Customer) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Customer) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetStatus

`func (o *Customer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Customer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Customer) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNextInvoiceDate

`func (o *Customer) GetNextInvoiceDate() string`

GetNextInvoiceDate returns the NextInvoiceDate field if non-nil, zero value otherwise.

### GetNextInvoiceDateOk

`func (o *Customer) GetNextInvoiceDateOk() (*string, bool)`

GetNextInvoiceDateOk returns a tuple with the NextInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoiceDate

`func (o *Customer) SetNextInvoiceDate(v string)`

SetNextInvoiceDate sets NextInvoiceDate field to given value.

### HasNextInvoiceDate

`func (o *Customer) HasNextInvoiceDate() bool`

HasNextInvoiceDate returns a boolean if a field has been set.

### GetStripeCustomerId

`func (o *Customer) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *Customer) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *Customer) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *Customer) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.

### GetPlans

`func (o *Customer) GetPlans() []CustomerPlan`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *Customer) GetPlansOk() (*[]CustomerPlan, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *Customer) SetPlans(v []CustomerPlan)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *Customer) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


