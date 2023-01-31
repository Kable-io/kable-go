# CustomerEntitlementResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the entitlement | 
**EntitlementId** | **string** | A Kable-defined identifier for the entitlement. | 
**ExternalId** | **string** | An identifier for the entitlement as defined by your API. | 
**Datatype** | **string** | The type of data supported by the entitlement. | 
**Description** | **string** | A description of the entitlement. | 
**Created** | **time.Time** | Date at which entity was first created | 
**Modified** | **time.Time** | Date at which entity was last modified | 
**Value** | **map[string]interface{}** | The value of the entitlement. | 

## Methods

### NewCustomerEntitlementResponseDto

`func NewCustomerEntitlementResponseDto(name string, entitlementId string, externalId string, datatype string, description string, created time.Time, modified time.Time, value map[string]interface{}, ) *CustomerEntitlementResponseDto`

NewCustomerEntitlementResponseDto instantiates a new CustomerEntitlementResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerEntitlementResponseDtoWithDefaults

`func NewCustomerEntitlementResponseDtoWithDefaults() *CustomerEntitlementResponseDto`

NewCustomerEntitlementResponseDtoWithDefaults instantiates a new CustomerEntitlementResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomerEntitlementResponseDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerEntitlementResponseDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerEntitlementResponseDto) SetName(v string)`

SetName sets Name field to given value.


### GetEntitlementId

`func (o *CustomerEntitlementResponseDto) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *CustomerEntitlementResponseDto) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *CustomerEntitlementResponseDto) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetExternalId

`func (o *CustomerEntitlementResponseDto) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CustomerEntitlementResponseDto) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CustomerEntitlementResponseDto) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetDatatype

`func (o *CustomerEntitlementResponseDto) GetDatatype() string`

GetDatatype returns the Datatype field if non-nil, zero value otherwise.

### GetDatatypeOk

`func (o *CustomerEntitlementResponseDto) GetDatatypeOk() (*string, bool)`

GetDatatypeOk returns a tuple with the Datatype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatype

`func (o *CustomerEntitlementResponseDto) SetDatatype(v string)`

SetDatatype sets Datatype field to given value.


### GetDescription

`func (o *CustomerEntitlementResponseDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomerEntitlementResponseDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomerEntitlementResponseDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreated

`func (o *CustomerEntitlementResponseDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomerEntitlementResponseDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomerEntitlementResponseDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *CustomerEntitlementResponseDto) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CustomerEntitlementResponseDto) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CustomerEntitlementResponseDto) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetValue

`func (o *CustomerEntitlementResponseDto) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomerEntitlementResponseDto) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomerEntitlementResponseDto) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


