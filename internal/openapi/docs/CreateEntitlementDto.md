# CreateEntitlementDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the entitlement | 
**ExternalId** | **string** | An identifier for the entitlement as defined by your API. | 
**Datatype** | Pointer to **string** | The type of data supported by the entitlement. | [optional] [default to "BOOLEAN"]
**Description** | Pointer to **string** | A description of the entitlement. | [optional] 

## Methods

### NewCreateEntitlementDto

`func NewCreateEntitlementDto(name string, externalId string, ) *CreateEntitlementDto`

NewCreateEntitlementDto instantiates a new CreateEntitlementDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEntitlementDtoWithDefaults

`func NewCreateEntitlementDtoWithDefaults() *CreateEntitlementDto`

NewCreateEntitlementDtoWithDefaults instantiates a new CreateEntitlementDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEntitlementDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEntitlementDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEntitlementDto) SetName(v string)`

SetName sets Name field to given value.


### GetExternalId

`func (o *CreateEntitlementDto) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateEntitlementDto) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateEntitlementDto) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetDatatype

`func (o *CreateEntitlementDto) GetDatatype() string`

GetDatatype returns the Datatype field if non-nil, zero value otherwise.

### GetDatatypeOk

`func (o *CreateEntitlementDto) GetDatatypeOk() (*string, bool)`

GetDatatypeOk returns a tuple with the Datatype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatype

`func (o *CreateEntitlementDto) SetDatatype(v string)`

SetDatatype sets Datatype field to given value.

### HasDatatype

`func (o *CreateEntitlementDto) HasDatatype() bool`

HasDatatype returns a boolean if a field has been set.

### GetDescription

`func (o *CreateEntitlementDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateEntitlementDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateEntitlementDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateEntitlementDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


