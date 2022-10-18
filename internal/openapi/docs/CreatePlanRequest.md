# CreatePlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-readable name for the plan, visible on dashboards, invoices, and reports. | 
**ExternalId** | Pointer to **string** | An identifier for the plan as defined by your API. | [optional] 
**Description** | Pointer to **string** | A human-readable description for the plan, visible on dashboards, invoices, and reports. | [optional] 
**Nickname** | Pointer to **string** | An additional identifier for the plan, defined by you, that is *not* visible to customers. If you have different *Monthly Active Users Plan*s for different cohorts of customers, this is a useful place to differentiate those. | [optional] 
**Type** | **string** | The type of plan. | 
**Interval** | **string** | The billing interval for the plan. | 
**DimensionId** | Pointer to **string** | An identifier of the dimension along which usage is aggregated in this plan. | [optional] 
**Aggregation** | Pointer to **string** | The aggregation along which usage metrics are calculated in this plan, relevant only for usage plans. | [optional] 
**Price** | [**Price**](Price.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary key-value pairs to attach to the object that can be useful for controlling functionality inside your API. | [optional] 

## Methods

### NewCreatePlanRequest

`func NewCreatePlanRequest(name string, type_ string, interval string, price Price, ) *CreatePlanRequest`

NewCreatePlanRequest instantiates a new CreatePlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePlanRequestWithDefaults

`func NewCreatePlanRequestWithDefaults() *CreatePlanRequest`

NewCreatePlanRequestWithDefaults instantiates a new CreatePlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatePlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePlanRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExternalId

`func (o *CreatePlanRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreatePlanRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreatePlanRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreatePlanRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDescription

`func (o *CreatePlanRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePlanRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePlanRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePlanRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNickname

`func (o *CreatePlanRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *CreatePlanRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *CreatePlanRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *CreatePlanRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetType

`func (o *CreatePlanRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePlanRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePlanRequest) SetType(v string)`

SetType sets Type field to given value.


### GetInterval

`func (o *CreatePlanRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CreatePlanRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CreatePlanRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetDimensionId

`func (o *CreatePlanRequest) GetDimensionId() string`

GetDimensionId returns the DimensionId field if non-nil, zero value otherwise.

### GetDimensionIdOk

`func (o *CreatePlanRequest) GetDimensionIdOk() (*string, bool)`

GetDimensionIdOk returns a tuple with the DimensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionId

`func (o *CreatePlanRequest) SetDimensionId(v string)`

SetDimensionId sets DimensionId field to given value.

### HasDimensionId

`func (o *CreatePlanRequest) HasDimensionId() bool`

HasDimensionId returns a boolean if a field has been set.

### GetAggregation

`func (o *CreatePlanRequest) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *CreatePlanRequest) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *CreatePlanRequest) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *CreatePlanRequest) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetPrice

`func (o *CreatePlanRequest) GetPrice() Price`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreatePlanRequest) GetPriceOk() (*Price, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreatePlanRequest) SetPrice(v Price)`

SetPrice sets Price field to given value.


### GetMetadata

`func (o *CreatePlanRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreatePlanRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreatePlanRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreatePlanRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


