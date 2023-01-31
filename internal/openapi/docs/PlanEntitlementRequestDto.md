# PlanEntitlementRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementId** | **string** | The entitlement to assign. | 
**Value** | **map[string]interface{}** | The value of the entitlement. | 

## Methods

### NewPlanEntitlementRequestDto

`func NewPlanEntitlementRequestDto(entitlementId string, value map[string]interface{}, ) *PlanEntitlementRequestDto`

NewPlanEntitlementRequestDto instantiates a new PlanEntitlementRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanEntitlementRequestDtoWithDefaults

`func NewPlanEntitlementRequestDtoWithDefaults() *PlanEntitlementRequestDto`

NewPlanEntitlementRequestDtoWithDefaults instantiates a new PlanEntitlementRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementId

`func (o *PlanEntitlementRequestDto) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *PlanEntitlementRequestDto) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *PlanEntitlementRequestDto) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetValue

`func (o *PlanEntitlementRequestDto) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PlanEntitlementRequestDto) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PlanEntitlementRequestDto) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


