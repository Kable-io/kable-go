# CreateCustomer200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KableClientId** | Pointer to **string** |  | [optional] 
**CustomerId** | **string** |  | 
**ClientId** | **string** |  | 
**CompanyName** | **string** |  | 
**Currency** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**NextInvoiceDate** | Pointer to **time.Time** |  | [optional] 
**StripeCustomerId** | Pointer to **string** |  | [optional] 
**Plans** | Pointer to [**[]CustomerPlan**](CustomerPlan.md) |  | [optional] 
**Keys** | Pointer to [**[]Key**](Key.md) |  | [optional] 

## Methods

### NewCreateCustomer200Response

`func NewCreateCustomer200Response(customerId string, clientId string, companyName string, ) *CreateCustomer200Response`

NewCreateCustomer200Response instantiates a new CreateCustomer200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomer200ResponseWithDefaults

`func NewCreateCustomer200ResponseWithDefaults() *CreateCustomer200Response`

NewCreateCustomer200ResponseWithDefaults instantiates a new CreateCustomer200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKableClientId

`func (o *CreateCustomer200Response) GetKableClientId() string`

GetKableClientId returns the KableClientId field if non-nil, zero value otherwise.

### GetKableClientIdOk

`func (o *CreateCustomer200Response) GetKableClientIdOk() (*string, bool)`

GetKableClientIdOk returns a tuple with the KableClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKableClientId

`func (o *CreateCustomer200Response) SetKableClientId(v string)`

SetKableClientId sets KableClientId field to given value.

### HasKableClientId

`func (o *CreateCustomer200Response) HasKableClientId() bool`

HasKableClientId returns a boolean if a field has been set.

### GetCustomerId

`func (o *CreateCustomer200Response) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateCustomer200Response) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateCustomer200Response) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetClientId

`func (o *CreateCustomer200Response) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateCustomer200Response) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateCustomer200Response) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetCompanyName

`func (o *CreateCustomer200Response) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CreateCustomer200Response) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CreateCustomer200Response) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetCurrency

`func (o *CreateCustomer200Response) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateCustomer200Response) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateCustomer200Response) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateCustomer200Response) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetStatus

`func (o *CreateCustomer200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateCustomer200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateCustomer200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateCustomer200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNextInvoiceDate

`func (o *CreateCustomer200Response) GetNextInvoiceDate() time.Time`

GetNextInvoiceDate returns the NextInvoiceDate field if non-nil, zero value otherwise.

### GetNextInvoiceDateOk

`func (o *CreateCustomer200Response) GetNextInvoiceDateOk() (*time.Time, bool)`

GetNextInvoiceDateOk returns a tuple with the NextInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoiceDate

`func (o *CreateCustomer200Response) SetNextInvoiceDate(v time.Time)`

SetNextInvoiceDate sets NextInvoiceDate field to given value.

### HasNextInvoiceDate

`func (o *CreateCustomer200Response) HasNextInvoiceDate() bool`

HasNextInvoiceDate returns a boolean if a field has been set.

### GetStripeCustomerId

`func (o *CreateCustomer200Response) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *CreateCustomer200Response) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *CreateCustomer200Response) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *CreateCustomer200Response) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.

### GetPlans

`func (o *CreateCustomer200Response) GetPlans() []CustomerPlan`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *CreateCustomer200Response) GetPlansOk() (*[]CustomerPlan, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *CreateCustomer200Response) SetPlans(v []CustomerPlan)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *CreateCustomer200Response) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetKeys

`func (o *CreateCustomer200Response) GetKeys() []Key`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *CreateCustomer200Response) GetKeysOk() (*[]Key, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *CreateCustomer200Response) SetKeys(v []Key)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *CreateCustomer200Response) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


