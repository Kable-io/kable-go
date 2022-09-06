# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Your customer&#39;s client ID (as defined by you during customer creation). This is how Kable attributes events to a given customer. | 
**Timestamp** | **time.Time** | When the event occurred. Timestamps must be formatted as RFC 3339 strings like &#x60;2022-01-09T09:32:01Z&#x60;. | 
**Data** | Pointer to **map[string]interface{}** | &#x60;data&#x60; is a JSON-formatted object containing key-value pairs of information to be tracked by Kable. The keys provided in the &#x60;data&#x60; JSON correspond to the Dimensions you&#39;ve defined on Kable.  When using a Kable library, you must **always** include a &#x60;clientId&#x60; in the &#x60;record&#x60; payload so that Kable can accurately attribute events to your customers.  | [optional] 

## Methods

### NewEvent

`func NewEvent(clientId string, timestamp time.Time, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *Event) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Event) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Event) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetTimestamp

`func (o *Event) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Event) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Event) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetData

`func (o *Event) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Event) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Event) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *Event) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


