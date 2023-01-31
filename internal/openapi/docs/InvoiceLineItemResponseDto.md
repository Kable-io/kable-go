# InvoiceLineItemResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | Pointer to **string** | The identifier of the plan to which this line item corresponds. | [optional] 
**Price** | [**InvoiceLineItemTierResponseDtoPrice**](InvoiceLineItemTierResponseDtoPrice.md) |  | 
**Usage** | Pointer to [**InvoiceLineItemTierResponseDtoUsage**](InvoiceLineItemTierResponseDtoUsage.md) |  | [optional] 
**Items** | [**[]InvoiceLineItemTierResponseDto**](InvoiceLineItemTierResponseDto.md) | Sub-line items associated with this line item. | 
**StartDate** | **time.Time** | Coverage period start date for the line item. | 
**EndDate** | **time.Time** | Coverage period end date for the line item. | 

## Methods

### NewInvoiceLineItemResponseDto

`func NewInvoiceLineItemResponseDto(price InvoiceLineItemTierResponseDtoPrice, items []InvoiceLineItemTierResponseDto, startDate time.Time, endDate time.Time, ) *InvoiceLineItemResponseDto`

NewInvoiceLineItemResponseDto instantiates a new InvoiceLineItemResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineItemResponseDtoWithDefaults

`func NewInvoiceLineItemResponseDtoWithDefaults() *InvoiceLineItemResponseDto`

NewInvoiceLineItemResponseDtoWithDefaults instantiates a new InvoiceLineItemResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *InvoiceLineItemResponseDto) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *InvoiceLineItemResponseDto) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *InvoiceLineItemResponseDto) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *InvoiceLineItemResponseDto) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPrice

`func (o *InvoiceLineItemResponseDto) GetPrice() InvoiceLineItemTierResponseDtoPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvoiceLineItemResponseDto) GetPriceOk() (*InvoiceLineItemTierResponseDtoPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvoiceLineItemResponseDto) SetPrice(v InvoiceLineItemTierResponseDtoPrice)`

SetPrice sets Price field to given value.


### GetUsage

`func (o *InvoiceLineItemResponseDto) GetUsage() InvoiceLineItemTierResponseDtoUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *InvoiceLineItemResponseDto) GetUsageOk() (*InvoiceLineItemTierResponseDtoUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *InvoiceLineItemResponseDto) SetUsage(v InvoiceLineItemTierResponseDtoUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *InvoiceLineItemResponseDto) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetItems

`func (o *InvoiceLineItemResponseDto) GetItems() []InvoiceLineItemTierResponseDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InvoiceLineItemResponseDto) GetItemsOk() (*[]InvoiceLineItemTierResponseDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InvoiceLineItemResponseDto) SetItems(v []InvoiceLineItemTierResponseDto)`

SetItems sets Items field to given value.


### GetStartDate

`func (o *InvoiceLineItemResponseDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InvoiceLineItemResponseDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InvoiceLineItemResponseDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *InvoiceLineItemResponseDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InvoiceLineItemResponseDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InvoiceLineItemResponseDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


