# CustomerResponseWithKeysDto

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
**Keys** | Pointer to [**[]KeyResponseDto**](KeyResponseDto.md) | The customer&#39;s API keys (relevant only if you use Kable as an authentication provider) | [optional] 

## Methods

### NewCustomerResponseWithKeysDto

`func NewCustomerResponseWithKeysDto(kableClientId string, customerId string, clientId string, companyName string, currency string, status string, plans []CustomerPlanResponseDto, bundles []CustomerBundleResponseDto, entitlements []ResourceEntitlementResponseDto, created time.Time, modified time.Time, ) *CustomerResponseWithKeysDto`

NewCustomerResponseWithKeysDto instantiates a new CustomerResponseWithKeysDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerResponseWithKeysDtoWithDefaults

`func NewCustomerResponseWithKeysDtoWithDefaults() *CustomerResponseWithKeysDto`

NewCustomerResponseWithKeysDtoWithDefaults instantiates a new CustomerResponseWithKeysDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKableClientId

`func (o *CustomerResponseWithKeysDto) GetKableClientId() string`

GetKableClientId returns the KableClientId field if non-nil, zero value otherwise.

### GetKableClientIdOk

`func (o *CustomerResponseWithKeysDto) GetKableClientIdOk() (*string, bool)`

GetKableClientIdOk returns a tuple with the KableClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKableClientId

`func (o *CustomerResponseWithKeysDto) SetKableClientId(v string)`

SetKableClientId sets KableClientId field to given value.


### GetCustomerId

`func (o *CustomerResponseWithKeysDto) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerResponseWithKeysDto) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerResponseWithKeysDto) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetClientId

`func (o *CustomerResponseWithKeysDto) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CustomerResponseWithKeysDto) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CustomerResponseWithKeysDto) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetCompanyName

`func (o *CustomerResponseWithKeysDto) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CustomerResponseWithKeysDto) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CustomerResponseWithKeysDto) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetCurrency

`func (o *CustomerResponseWithKeysDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomerResponseWithKeysDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomerResponseWithKeysDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetStatus

`func (o *CustomerResponseWithKeysDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerResponseWithKeysDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerResponseWithKeysDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNextInvoiceDate

`func (o *CustomerResponseWithKeysDto) GetNextInvoiceDate() time.Time`

GetNextInvoiceDate returns the NextInvoiceDate field if non-nil, zero value otherwise.

### GetNextInvoiceDateOk

`func (o *CustomerResponseWithKeysDto) GetNextInvoiceDateOk() (*time.Time, bool)`

GetNextInvoiceDateOk returns a tuple with the NextInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoiceDate

`func (o *CustomerResponseWithKeysDto) SetNextInvoiceDate(v time.Time)`

SetNextInvoiceDate sets NextInvoiceDate field to given value.

### HasNextInvoiceDate

`func (o *CustomerResponseWithKeysDto) HasNextInvoiceDate() bool`

HasNextInvoiceDate returns a boolean if a field has been set.

### GetBillingContact

`func (o *CustomerResponseWithKeysDto) GetBillingContact() CustomerResponseDtoBillingContact`

GetBillingContact returns the BillingContact field if non-nil, zero value otherwise.

### GetBillingContactOk

`func (o *CustomerResponseWithKeysDto) GetBillingContactOk() (*CustomerResponseDtoBillingContact, bool)`

GetBillingContactOk returns a tuple with the BillingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingContact

`func (o *CustomerResponseWithKeysDto) SetBillingContact(v CustomerResponseDtoBillingContact)`

SetBillingContact sets BillingContact field to given value.

### HasBillingContact

`func (o *CustomerResponseWithKeysDto) HasBillingContact() bool`

HasBillingContact returns a boolean if a field has been set.

### GetPlans

`func (o *CustomerResponseWithKeysDto) GetPlans() []CustomerPlanResponseDto`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *CustomerResponseWithKeysDto) GetPlansOk() (*[]CustomerPlanResponseDto, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *CustomerResponseWithKeysDto) SetPlans(v []CustomerPlanResponseDto)`

SetPlans sets Plans field to given value.


### GetBundles

`func (o *CustomerResponseWithKeysDto) GetBundles() []CustomerBundleResponseDto`

GetBundles returns the Bundles field if non-nil, zero value otherwise.

### GetBundlesOk

`func (o *CustomerResponseWithKeysDto) GetBundlesOk() (*[]CustomerBundleResponseDto, bool)`

GetBundlesOk returns a tuple with the Bundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundles

`func (o *CustomerResponseWithKeysDto) SetBundles(v []CustomerBundleResponseDto)`

SetBundles sets Bundles field to given value.


### GetEntitlements

`func (o *CustomerResponseWithKeysDto) GetEntitlements() []ResourceEntitlementResponseDto`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *CustomerResponseWithKeysDto) GetEntitlementsOk() (*[]ResourceEntitlementResponseDto, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *CustomerResponseWithKeysDto) SetEntitlements(v []ResourceEntitlementResponseDto)`

SetEntitlements sets Entitlements field to given value.


### GetCreated

`func (o *CustomerResponseWithKeysDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomerResponseWithKeysDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomerResponseWithKeysDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *CustomerResponseWithKeysDto) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CustomerResponseWithKeysDto) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CustomerResponseWithKeysDto) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetKeys

`func (o *CustomerResponseWithKeysDto) GetKeys() []KeyResponseDto`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *CustomerResponseWithKeysDto) GetKeysOk() (*[]KeyResponseDto, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *CustomerResponseWithKeysDto) SetKeys(v []KeyResponseDto)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *CustomerResponseWithKeysDto) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


