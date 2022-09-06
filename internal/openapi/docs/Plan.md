# Plan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** | A Kable-defined identifier for the plan. | 
**DimensionId** | **string** | An identifier of the dimension along which usage is aggregated in this plan. | 
**Name** | **string** | A human-readable name for the plan, visible on dashboards, invoices, and reports. | 
**Description** | Pointer to **string** | A human-readable description for the plan, visible on dashboards, invoices, and reports. | [optional] 
**Nickname** | Pointer to **string** | An additional identifier for the plan, defined by you, that is *not* visible to customers. If you have different *Monthly Active Users Plan*s for different cohorts of customers, this is a useful place to differentiate those. | [optional] 
**NextPlanId** | Pointer to **string** | An optional identifier of any \&quot;next\&quot; plan after the termination of a free trials or pilot programs. | [optional] 

## Methods

### NewPlan

`func NewPlan(planId string, dimensionId string, name string, ) *Plan`

NewPlan instantiates a new Plan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWithDefaults

`func NewPlanWithDefaults() *Plan`

NewPlanWithDefaults instantiates a new Plan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *Plan) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *Plan) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *Plan) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetDimensionId

`func (o *Plan) GetDimensionId() string`

GetDimensionId returns the DimensionId field if non-nil, zero value otherwise.

### GetDimensionIdOk

`func (o *Plan) GetDimensionIdOk() (*string, bool)`

GetDimensionIdOk returns a tuple with the DimensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionId

`func (o *Plan) SetDimensionId(v string)`

SetDimensionId sets DimensionId field to given value.


### GetName

`func (o *Plan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Plan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Plan) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Plan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Plan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Plan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Plan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNickname

`func (o *Plan) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *Plan) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *Plan) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *Plan) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetNextPlanId

`func (o *Plan) GetNextPlanId() string`

GetNextPlanId returns the NextPlanId field if non-nil, zero value otherwise.

### GetNextPlanIdOk

`func (o *Plan) GetNextPlanIdOk() (*string, bool)`

GetNextPlanIdOk returns a tuple with the NextPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPlanId

`func (o *Plan) SetNextPlanId(v string)`

SetNextPlanId sets NextPlanId field to given value.

### HasNextPlanId

`func (o *Plan) HasNextPlanId() bool`

HasNextPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


