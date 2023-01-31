# InvoiceLineItemTierResponseDtoUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | **string** | The aggregation type | 
**DimensionId** | **string** | The identifier of the dimension along which usage is aggregated in this line item. | 
**Total** | **float32** | The total usage consumption amount. | 

## Methods

### NewInvoiceLineItemTierResponseDtoUsage

`func NewInvoiceLineItemTierResponseDtoUsage(aggregation string, dimensionId string, total float32, ) *InvoiceLineItemTierResponseDtoUsage`

NewInvoiceLineItemTierResponseDtoUsage instantiates a new InvoiceLineItemTierResponseDtoUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineItemTierResponseDtoUsageWithDefaults

`func NewInvoiceLineItemTierResponseDtoUsageWithDefaults() *InvoiceLineItemTierResponseDtoUsage`

NewInvoiceLineItemTierResponseDtoUsageWithDefaults instantiates a new InvoiceLineItemTierResponseDtoUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *InvoiceLineItemTierResponseDtoUsage) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *InvoiceLineItemTierResponseDtoUsage) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *InvoiceLineItemTierResponseDtoUsage) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.


### GetDimensionId

`func (o *InvoiceLineItemTierResponseDtoUsage) GetDimensionId() string`

GetDimensionId returns the DimensionId field if non-nil, zero value otherwise.

### GetDimensionIdOk

`func (o *InvoiceLineItemTierResponseDtoUsage) GetDimensionIdOk() (*string, bool)`

GetDimensionIdOk returns a tuple with the DimensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionId

`func (o *InvoiceLineItemTierResponseDtoUsage) SetDimensionId(v string)`

SetDimensionId sets DimensionId field to given value.


### GetTotal

`func (o *InvoiceLineItemTierResponseDtoUsage) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InvoiceLineItemTierResponseDtoUsage) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InvoiceLineItemTierResponseDtoUsage) SetTotal(v float32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


