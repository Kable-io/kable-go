# AddPlanEntitlementRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlements** | Pointer to [**[]PlanEntitlementRequestDto**](PlanEntitlementRequestDto.md) | The plan entitlements to add | [optional] 

## Methods

### NewAddPlanEntitlementRequestDto

`func NewAddPlanEntitlementRequestDto() *AddPlanEntitlementRequestDto`

NewAddPlanEntitlementRequestDto instantiates a new AddPlanEntitlementRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPlanEntitlementRequestDtoWithDefaults

`func NewAddPlanEntitlementRequestDtoWithDefaults() *AddPlanEntitlementRequestDto`

NewAddPlanEntitlementRequestDtoWithDefaults instantiates a new AddPlanEntitlementRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlements

`func (o *AddPlanEntitlementRequestDto) GetEntitlements() []PlanEntitlementRequestDto`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *AddPlanEntitlementRequestDto) GetEntitlementsOk() (*[]PlanEntitlementRequestDto, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *AddPlanEntitlementRequestDto) SetEntitlements(v []PlanEntitlementRequestDto)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *AddPlanEntitlementRequestDto) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


