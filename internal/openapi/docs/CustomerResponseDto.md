# CustomerResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account | 
**CustomerId** | **string** | A Kable-defined identifier for the customer. | 
**ClientId** | **string** | A unique identifier for the customer, defined by you. | 
**CompanyName** | **string** | The name of the customer. | 
**Currency** | **string** | Three-letter ISO currency code | 
**Status** | **string** | The status of the customer. | 
**NextInvoiceDate** | Pointer to **time.Time** | The date at which the customer&#39;s next invoice will occur. | [optional] 
**BillingContact** | Pointer to [**CustomerResponseDtoBillingContact**](CustomerResponseDtoBillingContact.md) |  | [optional] 
**Plans** | [**[]CustomerPlanResponseDto**](CustomerPlanResponseDto.md) | The customer&#39;s plans. | 
**Bundles** | [**[]CustomerBundleResponseDto**](CustomerBundleResponseDto.md) | The customer&#39;s bundles. | 
**Entitlements** | [**[]ResourceEntitlementResponseDto**](ResourceEntitlementResponseDto.md) | The customer&#39;s entitlements. | 
**Created** | **time.Time** | Date at which entity was first created | 
**Modified** | **time.Time** | Date at which entity was last modified | 

## Methods

### NewCustomerResponseDto

`func NewCustomerResponseDto(kableClientId string, customerId string, clientId string, companyName string, currency string, status string, plans []CustomerPlanResponseDto, bundles []CustomerBundleResponseDto, entitlements []ResourceEntitlementResponseDto, created time.Time, modified time.Time, ) *CustomerResponseDto`

NewCustomerResponseDto instantiates a new CustomerResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerResponseDtoWithDefaults

`func NewCustomerResponseDtoWithDefaults() *CustomerResponseDto`

NewCustomerResponseDtoWithDefaults instantiates a new CustomerResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKableClientId

`func (o *CustomerResponseDto) GetKableClientId() string`

GetKableClientId returns the KableClientId field if non-nil, zero value otherwise.

### GetKableClientIdOk

`func (o *CustomerResponseDto) GetKableClientIdOk() (*string, bool)`

GetKableClientIdOk returns a tuple with the KableClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKableClientId

`func (o *CustomerResponseDto) SetKableClientId(v string)`

SetKableClientId sets KableClientId field to given value.


### GetCustomerId

`func (o *CustomerResponseDto) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerResponseDto) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerResponseDto) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetClientId

`func (o *CustomerResponseDto) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CustomerResponseDto) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CustomerResponseDto) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetCompanyName

`func (o *CustomerResponseDto) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CustomerResponseDto) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CustomerResponseDto) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetCurrency

`func (o *CustomerResponseDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomerResponseDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomerResponseDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetStatus

`func (o *CustomerResponseDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerResponseDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerResponseDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNextInvoiceDate

`func (o *CustomerResponseDto) GetNextInvoiceDate() time.Time`

GetNextInvoiceDate returns the NextInvoiceDate field if non-nil, zero value otherwise.

### GetNextInvoiceDateOk

`func (o *CustomerResponseDto) GetNextInvoiceDateOk() (*time.Time, bool)`

GetNextInvoiceDateOk returns a tuple with the NextInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoiceDate

`func (o *CustomerResponseDto) SetNextInvoiceDate(v time.Time)`

SetNextInvoiceDate sets NextInvoiceDate field to given value.

### HasNextInvoiceDate

`func (o *CustomerResponseDto) HasNextInvoiceDate() bool`

HasNextInvoiceDate returns a boolean if a field has been set.

### GetBillingContact

`func (o *CustomerResponseDto) GetBillingContact() CustomerResponseDtoBillingContact`

GetBillingContact returns the BillingContact field if non-nil, zero value otherwise.

### GetBillingContactOk

`func (o *CustomerResponseDto) GetBillingContactOk() (*CustomerResponseDtoBillingContact, bool)`

GetBillingContactOk returns a tuple with the BillingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingContact

`func (o *CustomerResponseDto) SetBillingContact(v CustomerResponseDtoBillingContact)`

SetBillingContact sets BillingContact field to given value.

### HasBillingContact

`func (o *CustomerResponseDto) HasBillingContact() bool`

HasBillingContact returns a boolean if a field has been set.

### GetPlans

`func (o *CustomerResponseDto) GetPlans() []CustomerPlanResponseDto`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *CustomerResponseDto) GetPlansOk() (*[]CustomerPlanResponseDto, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *CustomerResponseDto) SetPlans(v []CustomerPlanResponseDto)`

SetPlans sets Plans field to given value.


### GetBundles

`func (o *CustomerResponseDto) GetBundles() []CustomerBundleResponseDto`

GetBundles returns the Bundles field if non-nil, zero value otherwise.

### GetBundlesOk

`func (o *CustomerResponseDto) GetBundlesOk() (*[]CustomerBundleResponseDto, bool)`

GetBundlesOk returns a tuple with the Bundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundles

`func (o *CustomerResponseDto) SetBundles(v []CustomerBundleResponseDto)`

SetBundles sets Bundles field to given value.


### GetEntitlements

`func (o *CustomerResponseDto) GetEntitlements() []ResourceEntitlementResponseDto`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *CustomerResponseDto) GetEntitlementsOk() (*[]ResourceEntitlementResponseDto, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *CustomerResponseDto) SetEntitlements(v []ResourceEntitlementResponseDto)`

SetEntitlements sets Entitlements field to given value.


### GetCreated

`func (o *CustomerResponseDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomerResponseDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomerResponseDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *CustomerResponseDto) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CustomerResponseDto) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CustomerResponseDto) SetModified(v time.Time)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


