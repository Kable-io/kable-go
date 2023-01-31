# InvoiceLineItemTierResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usage** | Pointer to [**InvoiceLineItemTierResponseDtoUsage**](InvoiceLineItemTierResponseDtoUsage.md) |  | [optional] 
**Price** | [**InvoiceLineItemTierResponseDtoPrice**](InvoiceLineItemTierResponseDtoPrice.md) |  | 

## Methods

### NewInvoiceLineItemTierResponseDto

`func NewInvoiceLineItemTierResponseDto(price InvoiceLineItemTierResponseDtoPrice, ) *InvoiceLineItemTierResponseDto`

NewInvoiceLineItemTierResponseDto instantiates a new InvoiceLineItemTierResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineItemTierResponseDtoWithDefaults

`func NewInvoiceLineItemTierResponseDtoWithDefaults() *InvoiceLineItemTierResponseDto`

NewInvoiceLineItemTierResponseDtoWithDefaults instantiates a new InvoiceLineItemTierResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsage

`func (o *InvoiceLineItemTierResponseDto) GetUsage() InvoiceLineItemTierResponseDtoUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *InvoiceLineItemTierResponseDto) GetUsageOk() (*InvoiceLineItemTierResponseDtoUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *InvoiceLineItemTierResponseDto) SetUsage(v InvoiceLineItemTierResponseDtoUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *InvoiceLineItemTierResponseDto) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetPrice

`func (o *InvoiceLineItemTierResponseDto) GetPrice() InvoiceLineItemTierResponseDtoPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvoiceLineItemTierResponseDto) GetPriceOk() (*InvoiceLineItemTierResponseDtoPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvoiceLineItemTierResponseDto) SetPrice(v InvoiceLineItemTierResponseDtoPrice)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


