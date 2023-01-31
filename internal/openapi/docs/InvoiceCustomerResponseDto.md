# InvoiceCustomerResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | A Kable-defined identifier for the customer. | 
**ClientId** | **string** | A unique identifier for the customer, defined by you. | 
**CompanyName** | **string** | The name of the customer. | 

## Methods

### NewInvoiceCustomerResponseDto

`func NewInvoiceCustomerResponseDto(customerId string, clientId string, companyName string, ) *InvoiceCustomerResponseDto`

NewInvoiceCustomerResponseDto instantiates a new InvoiceCustomerResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceCustomerResponseDtoWithDefaults

`func NewInvoiceCustomerResponseDtoWithDefaults() *InvoiceCustomerResponseDto`

NewInvoiceCustomerResponseDtoWithDefaults instantiates a new InvoiceCustomerResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *InvoiceCustomerResponseDto) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InvoiceCustomerResponseDto) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InvoiceCustomerResponseDto) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetClientId

`func (o *InvoiceCustomerResponseDto) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InvoiceCustomerResponseDto) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InvoiceCustomerResponseDto) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetCompanyName

`func (o *InvoiceCustomerResponseDto) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InvoiceCustomerResponseDto) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InvoiceCustomerResponseDto) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


