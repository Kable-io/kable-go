# PlanEntitlementResponseDto

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

### NewPlanEntitlementResponseDto

`func NewPlanEntitlementResponseDto(name string, entitlementId string, externalId string, datatype string, description string, created time.Time, modified time.Time, value map[string]interface{}, ) *PlanEntitlementResponseDto`

NewPlanEntitlementResponseDto instantiates a new PlanEntitlementResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanEntitlementResponseDtoWithDefaults

`func NewPlanEntitlementResponseDtoWithDefaults() *PlanEntitlementResponseDto`

NewPlanEntitlementResponseDtoWithDefaults instantiates a new PlanEntitlementResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlanEntitlementResponseDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanEntitlementResponseDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanEntitlementResponseDto) SetName(v string)`

SetName sets Name field to given value.


### GetEntitlementId

`func (o *PlanEntitlementResponseDto) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *PlanEntitlementResponseDto) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *PlanEntitlementResponseDto) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetExternalId

`func (o *PlanEntitlementResponseDto) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PlanEntitlementResponseDto) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PlanEntitlementResponseDto) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetDatatype

`func (o *PlanEntitlementResponseDto) GetDatatype() string`

GetDatatype returns the Datatype field if non-nil, zero value otherwise.

### GetDatatypeOk

`func (o *PlanEntitlementResponseDto) GetDatatypeOk() (*string, bool)`

GetDatatypeOk returns a tuple with the Datatype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatype

`func (o *PlanEntitlementResponseDto) SetDatatype(v string)`

SetDatatype sets Datatype field to given value.


### GetDescription

`func (o *PlanEntitlementResponseDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanEntitlementResponseDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanEntitlementResponseDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreated

`func (o *PlanEntitlementResponseDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PlanEntitlementResponseDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PlanEntitlementResponseDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *PlanEntitlementResponseDto) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *PlanEntitlementResponseDto) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *PlanEntitlementResponseDto) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetValue

`func (o *PlanEntitlementResponseDto) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PlanEntitlementResponseDto) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PlanEntitlementResponseDto) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


