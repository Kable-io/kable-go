# BillingContactRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | Contact&#39;s first name. | [optional] 
**LastName** | Pointer to **string** | Contact&#39;s last name. | [optional] 
**Email** | Pointer to **string** | Contact&#39;s email address. | [optional] 
**Phone** | Pointer to **string** | Contact&#39;s phone number. | [optional] 
**AddressLine1** | Pointer to **string** | First line of contact&#39;s mailing address. | [optional] 
**AddressLine2** | Pointer to **string** | Second line of contact&#39;s mailing address. | [optional] 
**City** | Pointer to **string** | City of contact&#39;s mailing address. | [optional] 
**State** | Pointer to **string** | State of contact&#39;s mailing address. | [optional] 
**PostalCode** | Pointer to **string** | Postal code of contact&#39;s mailing address. | [optional] 
**Country** | Pointer to **string** | Country of contact&#39;s mailing address. | [optional] 

## Methods

### NewBillingContactRequestDto

`func NewBillingContactRequestDto() *BillingContactRequestDto`

NewBillingContactRequestDto instantiates a new BillingContactRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingContactRequestDtoWithDefaults

`func NewBillingContactRequestDtoWithDefaults() *BillingContactRequestDto`

NewBillingContactRequestDtoWithDefaults instantiates a new BillingContactRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *BillingContactRequestDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BillingContactRequestDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BillingContactRequestDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *BillingContactRequestDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *BillingContactRequestDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BillingContactRequestDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BillingContactRequestDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *BillingContactRequestDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *BillingContactRequestDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BillingContactRequestDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BillingContactRequestDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BillingContactRequestDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *BillingContactRequestDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BillingContactRequestDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BillingContactRequestDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *BillingContactRequestDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAddressLine1

`func (o *BillingContactRequestDto) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *BillingContactRequestDto) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *BillingContactRequestDto) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *BillingContactRequestDto) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *BillingContactRequestDto) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *BillingContactRequestDto) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *BillingContactRequestDto) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *BillingContactRequestDto) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *BillingContactRequestDto) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BillingContactRequestDto) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BillingContactRequestDto) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *BillingContactRequestDto) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *BillingContactRequestDto) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BillingContactRequestDto) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BillingContactRequestDto) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BillingContactRequestDto) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *BillingContactRequestDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *BillingContactRequestDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *BillingContactRequestDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *BillingContactRequestDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *BillingContactRequestDto) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BillingContactRequestDto) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BillingContactRequestDto) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BillingContactRequestDto) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


