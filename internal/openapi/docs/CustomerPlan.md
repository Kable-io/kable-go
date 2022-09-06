# CustomerPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** | A Kable-defined identifier for the plan. | 
**DimensionId** | **string** | An identifier of the dimension along which usage is aggregated in this plan. | 
**Name** | **string** | A human-readable name for the plan, visible on dashboards, invoices, and reports. | 
**Description** | Pointer to **string** | A human-readable description for the plan, visible on dashboards, invoices, and reports. | [optional] 
**Nickname** | Pointer to **string** | An additional identifier for the plan, defined by you, that is *not* visible to customers. If you have different *Monthly Active Users Plan*s for different cohorts of customers, this is a useful place to differentiate those. | [optional] 
**AddedDate** | **string** | Date the plan was added to the customer. | 
**EndedDate** | Pointer to **string** | Date the plan ended or will end for trials or pilot programs. | [optional] 
**PeriodStartDate** | Pointer to **string** | Date the current invoice cycle begins. | [optional] 
**PeriodEndDate** | Pointer to **string** | Date the current invoice cycle ends. | [optional] 
**NextPlanId** | Pointer to **string** | An optional identifier of any \&quot;next\&quot; plan after the termination of a free trials or pilot programs. | [optional] 

## Methods

### NewCustomerPlan

`func NewCustomerPlan(planId string, dimensionId string, name string, addedDate string, ) *CustomerPlan`

NewCustomerPlan instantiates a new CustomerPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPlanWithDefaults

`func NewCustomerPlanWithDefaults() *CustomerPlan`

NewCustomerPlanWithDefaults instantiates a new CustomerPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *CustomerPlan) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CustomerPlan) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CustomerPlan) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetDimensionId

`func (o *CustomerPlan) GetDimensionId() string`

GetDimensionId returns the DimensionId field if non-nil, zero value otherwise.

### GetDimensionIdOk

`func (o *CustomerPlan) GetDimensionIdOk() (*string, bool)`

GetDimensionIdOk returns a tuple with the DimensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionId

`func (o *CustomerPlan) SetDimensionId(v string)`

SetDimensionId sets DimensionId field to given value.


### GetName

`func (o *CustomerPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerPlan) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CustomerPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomerPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomerPlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomerPlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNickname

`func (o *CustomerPlan) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *CustomerPlan) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *CustomerPlan) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *CustomerPlan) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetAddedDate

`func (o *CustomerPlan) GetAddedDate() string`

GetAddedDate returns the AddedDate field if non-nil, zero value otherwise.

### GetAddedDateOk

`func (o *CustomerPlan) GetAddedDateOk() (*string, bool)`

GetAddedDateOk returns a tuple with the AddedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedDate

`func (o *CustomerPlan) SetAddedDate(v string)`

SetAddedDate sets AddedDate field to given value.


### GetEndedDate

`func (o *CustomerPlan) GetEndedDate() string`

GetEndedDate returns the EndedDate field if non-nil, zero value otherwise.

### GetEndedDateOk

`func (o *CustomerPlan) GetEndedDateOk() (*string, bool)`

GetEndedDateOk returns a tuple with the EndedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedDate

`func (o *CustomerPlan) SetEndedDate(v string)`

SetEndedDate sets EndedDate field to given value.

### HasEndedDate

`func (o *CustomerPlan) HasEndedDate() bool`

HasEndedDate returns a boolean if a field has been set.

### GetPeriodStartDate

`func (o *CustomerPlan) GetPeriodStartDate() string`

GetPeriodStartDate returns the PeriodStartDate field if non-nil, zero value otherwise.

### GetPeriodStartDateOk

`func (o *CustomerPlan) GetPeriodStartDateOk() (*string, bool)`

GetPeriodStartDateOk returns a tuple with the PeriodStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStartDate

`func (o *CustomerPlan) SetPeriodStartDate(v string)`

SetPeriodStartDate sets PeriodStartDate field to given value.

### HasPeriodStartDate

`func (o *CustomerPlan) HasPeriodStartDate() bool`

HasPeriodStartDate returns a boolean if a field has been set.

### GetPeriodEndDate

`func (o *CustomerPlan) GetPeriodEndDate() string`

GetPeriodEndDate returns the PeriodEndDate field if non-nil, zero value otherwise.

### GetPeriodEndDateOk

`func (o *CustomerPlan) GetPeriodEndDateOk() (*string, bool)`

GetPeriodEndDateOk returns a tuple with the PeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEndDate

`func (o *CustomerPlan) SetPeriodEndDate(v string)`

SetPeriodEndDate sets PeriodEndDate field to given value.

### HasPeriodEndDate

`func (o *CustomerPlan) HasPeriodEndDate() bool`

HasPeriodEndDate returns a boolean if a field has been set.

### GetNextPlanId

`func (o *CustomerPlan) GetNextPlanId() string`

GetNextPlanId returns the NextPlanId field if non-nil, zero value otherwise.

### GetNextPlanIdOk

`func (o *CustomerPlan) GetNextPlanIdOk() (*string, bool)`

GetNextPlanIdOk returns a tuple with the NextPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPlanId

`func (o *CustomerPlan) SetNextPlanId(v string)`

SetNextPlanId sets NextPlanId field to given value.

### HasNextPlanId

`func (o *CustomerPlan) HasNextPlanId() bool`

HasNextPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


