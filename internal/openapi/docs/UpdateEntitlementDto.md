# UpdateEntitlementDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the entitlement | [optional] 
**ExternalId** | Pointer to **string** | An identifier for the entitlement as defined by your API. | [optional] 
**Datatype** | Pointer to **string** | The type of data supported by the entitlement. | [optional] [default to "BOOLEAN"]
**Description** | Pointer to **string** | A description of the entitlement. | [optional] 

## Methods

### NewUpdateEntitlementDto

`func NewUpdateEntitlementDto() *UpdateEntitlementDto`

NewUpdateEntitlementDto instantiates a new UpdateEntitlementDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEntitlementDtoWithDefaults

`func NewUpdateEntitlementDtoWithDefaults() *UpdateEntitlementDto`

NewUpdateEntitlementDtoWithDefaults instantiates a new UpdateEntitlementDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateEntitlementDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateEntitlementDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateEntitlementDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateEntitlementDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateEntitlementDto) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateEntitlementDto) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateEntitlementDto) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateEntitlementDto) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDatatype

`func (o *UpdateEntitlementDto) GetDatatype() string`

GetDatatype returns the Datatype field if non-nil, zero value otherwise.

### GetDatatypeOk

`func (o *UpdateEntitlementDto) GetDatatypeOk() (*string, bool)`

GetDatatypeOk returns a tuple with the Datatype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatype

`func (o *UpdateEntitlementDto) SetDatatype(v string)`

SetDatatype sets Datatype field to given value.

### HasDatatype

`func (o *UpdateEntitlementDto) HasDatatype() bool`

HasDatatype returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateEntitlementDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateEntitlementDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateEntitlementDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateEntitlementDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


