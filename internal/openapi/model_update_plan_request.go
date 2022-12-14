/*
Kable API

The Kable API allows developers to manage **customers** and **API keys** and record **events** about their API.   ## API Host and Environments  Kable is accessible in Live and Test environments for authentication and metering of client API requests. You will have separate API keys to access each environment.  You should only use Kable's Live environment for your own production data. All other configured environments should use Kable's Test environment.   ## API Protocols and Headers  All requests to the Kable API are made over HTTPS TLS v1.2+ to ensure security. Calls made over HTTP will fail. Any requests without proper authentication will also fail.  The Kable API uses standard JSON for requests and responses. Be sure to set both the `Content-Type` and `Accept` headers on each request to application/json.  Each Kable API response includes a `requestId` as the `X-REQUEST-ID` response header. The `requestId` is included regardless whether the API request succeeded or failed. You can use this requestId to help with debugging or when contacting support regarding a specific API call.   ## API Keys  There are two types of API keys on Kable.  ### Kable Keys Kable Keys are the keys you, the Kable customer, use to interact with Kable. These keys help us ensure that only you are interacting with Kable on your behalf. You can find your keys on the Company page of the dashboard after you sign up.  Kable Keys should be included in every request to the Kable API. You must provide your client ID as the `KABLE-CLIENT-ID` header and your secret key as the `KABLE-CLIENT-SECRET` header on each request to Kable. If you are using a language-specific Kable library, you will initialize the SDK using these keys.  ### Customer Keys Customer Keys are the keys your customers use to interact with your API. Customer Keys are authenticated by Kable when a customer makes a request to your API if you use Kable's authentication services. Customers must provide their client ID (defined as `clientId` when you create the customer) as the `X-CLIENT-ID` header and their secret key as the `X-API-KEY` header on each request to your API that Kable is to authenticate.   ## API Versioning  All Kable endpoints are versioned. After the host, each API can be found at `/api/vX/...` where X is the API version.  We strive to ensure that changes to the Kable API are backward compatible. Sometimes, though, we must break from older design paradigms to make the product better. When this happens, a new version of the API is released.  The current version of Kable is **v1**. 

API version: 1.2.1
Contact: contact@kable.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UpdatePlanRequest struct for UpdatePlanRequest
type UpdatePlanRequest struct {
	// A human-readable name for the plan, visible on dashboards, invoices, and reports.
	Name *string `json:"name,omitempty"`
	// An identifier for the plan as defined by your API.
	ExternalId *string `json:"externalId,omitempty"`
	// A human-readable description for the plan, visible on dashboards, invoices, and reports.
	Description *string `json:"description,omitempty"`
	// An additional identifier for the plan, defined by you, that is *not* visible to customers. If you have different *Monthly Active Users Plan*s for different cohorts of customers, this is a useful place to differentiate those.
	Nickname *string `json:"nickname,omitempty"`
	// The type of plan.
	Type *string `json:"type,omitempty"`
	// The billing interval for the plan.
	Interval *string `json:"interval,omitempty"`
	// An identifier of the dimension along which usage is aggregated in this plan.
	DimensionId *string `json:"dimensionId,omitempty"`
	// The aggregation along which usage metrics are calculated in this plan, relevant only for usage plans.
	Aggregation *string `json:"aggregation,omitempty"`
	Price *Price `json:"price,omitempty"`
	// Arbitrary key-value pairs to attach to the object that can be useful for controlling functionality inside your API.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewUpdatePlanRequest instantiates a new UpdatePlanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePlanRequest() *UpdatePlanRequest {
	this := UpdatePlanRequest{}
	return &this
}

// NewUpdatePlanRequestWithDefaults instantiates a new UpdatePlanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePlanRequestWithDefaults() *UpdatePlanRequest {
	this := UpdatePlanRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdatePlanRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlanRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdatePlanRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdatePlanRequest) SetName(v string) {
	o.Name = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *UpdatePlanRequest) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlanRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *UpdatePlanRequest) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *UpdatePlanRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdatePlanRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlanRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdatePlanRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdatePlanRequest) SetDescription(v string) {
	o.Description = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *UpdatePlanRequest) GetNickname() string {
	if o == nil || o.Nickname == nil {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlanRequest) GetNicknameOk() (*string, bool) {
	if o == nil || o.Nickname == nil {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *UpdatePlanRequest) HasNickname() bool {
	if o != nil && o.Nickname != nil {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *UpdatePlanRequest) SetNickname(v string) {
	o.Nickname = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdatePlanRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlanRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdatePlanRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdatePlanRequest) SetType(v string) {
	o.Type = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *UpdatePlanRequest) GetInterval() string {
	if o == nil || o.Interval == nil {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlanRequest) GetIntervalOk() (*string, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *UpdatePlanRequest) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *UpdatePlanRequest) SetInterval(v string) {
	o.Interval = &v
}

// GetDimensionId returns the DimensionId field value if set, zero value otherwise.
func (o *UpdatePlanRequest) GetDimensionId() string {
	if o == nil || o.DimensionId == nil {
		var ret string
		return ret
	}
	return *o.DimensionId
}

// GetDimensionIdOk returns a tuple with the DimensionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlanRequest) GetDimensionIdOk() (*string, bool) {
	if o == nil || o.DimensionId == nil {
		return nil, false
	}
	return o.DimensionId, true
}

// HasDimensionId returns a boolean if a field has been set.
func (o *UpdatePlanRequest) HasDimensionId() bool {
	if o != nil && o.DimensionId != nil {
		return true
	}

	return false
}

// SetDimensionId gets a reference to the given string and assigns it to the DimensionId field.
func (o *UpdatePlanRequest) SetDimensionId(v string) {
	o.DimensionId = &v
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *UpdatePlanRequest) GetAggregation() string {
	if o == nil || o.Aggregation == nil {
		var ret string
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlanRequest) GetAggregationOk() (*string, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *UpdatePlanRequest) HasAggregation() bool {
	if o != nil && o.Aggregation != nil {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given string and assigns it to the Aggregation field.
func (o *UpdatePlanRequest) SetAggregation(v string) {
	o.Aggregation = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *UpdatePlanRequest) GetPrice() Price {
	if o == nil || o.Price == nil {
		var ret Price
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlanRequest) GetPriceOk() (*Price, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *UpdatePlanRequest) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given Price and assigns it to the Price field.
func (o *UpdatePlanRequest) SetPrice(v Price) {
	o.Price = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdatePlanRequest) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePlanRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdatePlanRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UpdatePlanRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o UpdatePlanRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Nickname != nil {
		toSerialize["nickname"] = o.Nickname
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.DimensionId != nil {
		toSerialize["dimensionId"] = o.DimensionId
	}
	if o.Aggregation != nil {
		toSerialize["aggregation"] = o.Aggregation
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatePlanRequest struct {
	value *UpdatePlanRequest
	isSet bool
}

func (v NullableUpdatePlanRequest) Get() *UpdatePlanRequest {
	return v.value
}

func (v *NullableUpdatePlanRequest) Set(val *UpdatePlanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePlanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePlanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePlanRequest(val *UpdatePlanRequest) *NullableUpdatePlanRequest {
	return &NullableUpdatePlanRequest{value: val, isSet: true}
}

func (v NullableUpdatePlanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePlanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


