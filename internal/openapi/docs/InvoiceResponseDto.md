# InvoiceResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KableClientId** | **string** | Your Kable client ID, found in the dashboard of your Kable account | 
**InvoiceId** | **string** | A Kable-defined identifier for the customer. | 
**Customer** | [**InvoiceResponseDtoCustomer**](InvoiceResponseDtoCustomer.md) |  | 
**PostDate** | **time.Time** | The date the invoice was posted. | 
**Status** | **string** | The status of the invoice. | 
**Price** | [**InvoiceResponseDtoPrice**](InvoiceResponseDtoPrice.md) |  | 
**Items** | [**[]InvoiceLineItemResponseDto**](InvoiceLineItemResponseDto.md) | The invoice line items. | 
**Created** | **time.Time** | Date at which entity was first created | 
**Modified** | **time.Time** | Date at which entity was last modified | 

## Methods

### NewInvoiceResponseDto

`func NewInvoiceResponseDto(kableClientId string, invoiceId string, customer InvoiceResponseDtoCustomer, postDate time.Time, status string, price InvoiceResponseDtoPrice, items []InvoiceLineItemResponseDto, created time.Time, modified time.Time, ) *InvoiceResponseDto`

NewInvoiceResponseDto instantiates a new InvoiceResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceResponseDtoWithDefaults

`func NewInvoiceResponseDtoWithDefaults() *InvoiceResponseDto`

NewInvoiceResponseDtoWithDefaults instantiates a new InvoiceResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKableClientId

`func (o *InvoiceResponseDto) GetKableClientId() string`

GetKableClientId returns the KableClientId field if non-nil, zero value otherwise.

### GetKableClientIdOk

`func (o *InvoiceResponseDto) GetKableClientIdOk() (*string, bool)`

GetKableClientIdOk returns a tuple with the KableClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKableClientId

`func (o *InvoiceResponseDto) SetKableClientId(v string)`

SetKableClientId sets KableClientId field to given value.


### GetInvoiceId

`func (o *InvoiceResponseDto) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceResponseDto) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceResponseDto) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetCustomer

`func (o *InvoiceResponseDto) GetCustomer() InvoiceResponseDtoCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InvoiceResponseDto) GetCustomerOk() (*InvoiceResponseDtoCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InvoiceResponseDto) SetCustomer(v InvoiceResponseDtoCustomer)`

SetCustomer sets Customer field to given value.


### GetPostDate

`func (o *InvoiceResponseDto) GetPostDate() time.Time`

GetPostDate returns the PostDate field if non-nil, zero value otherwise.

### GetPostDateOk

`func (o *InvoiceResponseDto) GetPostDateOk() (*time.Time, bool)`

GetPostDateOk returns a tuple with the PostDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostDate

`func (o *InvoiceResponseDto) SetPostDate(v time.Time)`

SetPostDate sets PostDate field to given value.


### GetStatus

`func (o *InvoiceResponseDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceResponseDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceResponseDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPrice

`func (o *InvoiceResponseDto) GetPrice() InvoiceResponseDtoPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvoiceResponseDto) GetPriceOk() (*InvoiceResponseDtoPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvoiceResponseDto) SetPrice(v InvoiceResponseDtoPrice)`

SetPrice sets Price field to given value.


### GetItems

`func (o *InvoiceResponseDto) GetItems() []InvoiceLineItemResponseDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InvoiceResponseDto) GetItemsOk() (*[]InvoiceLineItemResponseDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InvoiceResponseDto) SetItems(v []InvoiceLineItemResponseDto)`

SetItems sets Items field to given value.


### GetCreated

`func (o *InvoiceResponseDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InvoiceResponseDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InvoiceResponseDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *InvoiceResponseDto) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *InvoiceResponseDto) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *InvoiceResponseDto) SetModified(v time.Time)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


