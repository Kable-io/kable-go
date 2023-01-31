# EventDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | A unique identifier for the customer, defined by you. | 
**Timestamp** | **time.Time** | When the event occurred. Timestamps must be formatted as RFC-3339 strings like &#x60;2022-01-09T09:32:01Z&#x60;. | 
**Data** | Pointer to **map[string]interface{}** | &#x60;data&#x60; is a JSON-formatted object containing key-value pairs of information to be tracked by Kable. The keys provided in the &#x60;data&#x60; JSON correspond to the Dimensions you&#39;ve defined on Kable.  When using a Kable library, you must **always** include a &#x60;clientId&#x60; in the &#x60;record&#x60; payload so that Kable can accurately attribute events to your customers. | [optional] 
**TransactionId** | Pointer to **string** | A unique identifier for the event, used as an idempotency key for event deduplication. | [optional] 

## Methods

### NewEventDto

`func NewEventDto(clientId string, timestamp time.Time, ) *EventDto`

NewEventDto instantiates a new EventDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDtoWithDefaults

`func NewEventDtoWithDefaults() *EventDto`

NewEventDtoWithDefaults instantiates a new EventDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *EventDto) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *EventDto) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *EventDto) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetTimestamp

`func (o *EventDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetData

`func (o *EventDto) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventDto) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventDto) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *EventDto) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTransactionId

`func (o *EventDto) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *EventDto) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *EventDto) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *EventDto) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


