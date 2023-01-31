# EntitlementResponseDto

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

## Methods

### NewEntitlementResponseDto

`func NewEntitlementResponseDto(name string, entitlementId string, externalId string, datatype string, description string, created time.Time, modified time.Time, ) *EntitlementResponseDto`

NewEntitlementResponseDto instantiates a new EntitlementResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementResponseDtoWithDefaults

`func NewEntitlementResponseDtoWithDefaults() *EntitlementResponseDto`

NewEntitlementResponseDtoWithDefaults instantiates a new EntitlementResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EntitlementResponseDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementResponseDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementResponseDto) SetName(v string)`

SetName sets Name field to given value.


### GetEntitlementId

`func (o *EntitlementResponseDto) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *EntitlementResponseDto) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *EntitlementResponseDto) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetExternalId

`func (o *EntitlementResponseDto) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EntitlementResponseDto) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EntitlementResponseDto) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetDatatype

`func (o *EntitlementResponseDto) GetDatatype() string`

GetDatatype returns the Datatype field if non-nil, zero value otherwise.

### GetDatatypeOk

`func (o *EntitlementResponseDto) GetDatatypeOk() (*string, bool)`

GetDatatypeOk returns a tuple with the Datatype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatype

`func (o *EntitlementResponseDto) SetDatatype(v string)`

SetDatatype sets Datatype field to given value.


### GetDescription

`func (o *EntitlementResponseDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementResponseDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementResponseDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreated

`func (o *EntitlementResponseDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EntitlementResponseDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EntitlementResponseDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *EntitlementResponseDto) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *EntitlementResponseDto) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *EntitlementResponseDto) SetModified(v time.Time)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


