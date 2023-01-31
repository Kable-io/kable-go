# RemoveCustomerPlansRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | Pointer to **string** | The identifier for a single plan. | [optional] 
**PlanIds** | Pointer to **[]string** | The identifiers for many plans. | [optional] 

## Methods

### NewRemoveCustomerPlansRequestDto

`func NewRemoveCustomerPlansRequestDto() *RemoveCustomerPlansRequestDto`

NewRemoveCustomerPlansRequestDto instantiates a new RemoveCustomerPlansRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveCustomerPlansRequestDtoWithDefaults

`func NewRemoveCustomerPlansRequestDtoWithDefaults() *RemoveCustomerPlansRequestDto`

NewRemoveCustomerPlansRequestDtoWithDefaults instantiates a new RemoveCustomerPlansRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *RemoveCustomerPlansRequestDto) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *RemoveCustomerPlansRequestDto) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *RemoveCustomerPlansRequestDto) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *RemoveCustomerPlansRequestDto) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanIds

`func (o *RemoveCustomerPlansRequestDto) GetPlanIds() []string`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *RemoveCustomerPlansRequestDto) GetPlanIdsOk() (*[]string, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *RemoveCustomerPlansRequestDto) SetPlanIds(v []string)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *RemoveCustomerPlansRequestDto) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


