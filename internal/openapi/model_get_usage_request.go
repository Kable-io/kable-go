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
	"time"
)

// GetUsageRequest struct for GetUsageRequest
type GetUsageRequest struct {
	// Valid values are `COUNT`, `COUNT_DISTINCT`, and `SUM`.
	Aggregation string `json:"aggregation"`
	// Note that one of either `dimensionId` or `dimensionKey` is required, though either may be provided.
	DimensionId *string `json:"dimensionId,omitempty"`
	// Note that one of either `dimensionId` or `dimensionKey` is required, though either may be provided.
	DimensionKey *string `json:"dimensionKey,omitempty"`
	// Start of the period over which to query. Timestamps must be formatted as RFC 3339 strings like `2022-01-09T09:32:01Z`
	StartDate time.Time `json:"startDate"`
	// End of the period over which to query. Timestamps must be formatted as RFC 3339 strings like `2022-01-09T09:32:01Z`
	EndDate time.Time `json:"endDate"`
	// Valid values are `HOUR`, `DAY`, and `MONTH`.
	Interval string `json:"interval"`
	// When provided, query will consider only this customer. When `null`, query will consider all of your customers.
	ClientId *string `json:"clientId,omitempty"`
}

// NewGetUsageRequest instantiates a new GetUsageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUsageRequest(aggregation string, startDate time.Time, endDate time.Time, interval string) *GetUsageRequest {
	this := GetUsageRequest{}
	this.Aggregation = aggregation
	this.StartDate = startDate
	this.EndDate = endDate
	this.Interval = interval
	return &this
}

// NewGetUsageRequestWithDefaults instantiates a new GetUsageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUsageRequestWithDefaults() *GetUsageRequest {
	this := GetUsageRequest{}
	return &this
}

// GetAggregation returns the Aggregation field value
func (o *GetUsageRequest) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *GetUsageRequest) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value
func (o *GetUsageRequest) SetAggregation(v string) {
	o.Aggregation = v
}

// GetDimensionId returns the DimensionId field value if set, zero value otherwise.
func (o *GetUsageRequest) GetDimensionId() string {
	if o == nil || o.DimensionId == nil {
		var ret string
		return ret
	}
	return *o.DimensionId
}

// GetDimensionIdOk returns a tuple with the DimensionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsageRequest) GetDimensionIdOk() (*string, bool) {
	if o == nil || o.DimensionId == nil {
		return nil, false
	}
	return o.DimensionId, true
}

// HasDimensionId returns a boolean if a field has been set.
func (o *GetUsageRequest) HasDimensionId() bool {
	if o != nil && o.DimensionId != nil {
		return true
	}

	return false
}

// SetDimensionId gets a reference to the given string and assigns it to the DimensionId field.
func (o *GetUsageRequest) SetDimensionId(v string) {
	o.DimensionId = &v
}

// GetDimensionKey returns the DimensionKey field value if set, zero value otherwise.
func (o *GetUsageRequest) GetDimensionKey() string {
	if o == nil || o.DimensionKey == nil {
		var ret string
		return ret
	}
	return *o.DimensionKey
}

// GetDimensionKeyOk returns a tuple with the DimensionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsageRequest) GetDimensionKeyOk() (*string, bool) {
	if o == nil || o.DimensionKey == nil {
		return nil, false
	}
	return o.DimensionKey, true
}

// HasDimensionKey returns a boolean if a field has been set.
func (o *GetUsageRequest) HasDimensionKey() bool {
	if o != nil && o.DimensionKey != nil {
		return true
	}

	return false
}

// SetDimensionKey gets a reference to the given string and assigns it to the DimensionKey field.
func (o *GetUsageRequest) SetDimensionKey(v string) {
	o.DimensionKey = &v
}

// GetStartDate returns the StartDate field value
func (o *GetUsageRequest) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *GetUsageRequest) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *GetUsageRequest) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *GetUsageRequest) GetEndDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *GetUsageRequest) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *GetUsageRequest) SetEndDate(v time.Time) {
	o.EndDate = v
}

// GetInterval returns the Interval field value
func (o *GetUsageRequest) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *GetUsageRequest) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *GetUsageRequest) SetInterval(v string) {
	o.Interval = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *GetUsageRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsageRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *GetUsageRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *GetUsageRequest) SetClientId(v string) {
	o.ClientId = &v
}

func (o GetUsageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["aggregation"] = o.Aggregation
	}
	if o.DimensionId != nil {
		toSerialize["dimensionId"] = o.DimensionId
	}
	if o.DimensionKey != nil {
		toSerialize["dimensionKey"] = o.DimensionKey
	}
	if true {
		toSerialize["startDate"] = o.StartDate
	}
	if true {
		toSerialize["endDate"] = o.EndDate
	}
	if true {
		toSerialize["interval"] = o.Interval
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	return json.Marshal(toSerialize)
}

type NullableGetUsageRequest struct {
	value *GetUsageRequest
	isSet bool
}

func (v NullableGetUsageRequest) Get() *GetUsageRequest {
	return v.value
}

func (v *NullableGetUsageRequest) Set(val *GetUsageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUsageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUsageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUsageRequest(val *GetUsageRequest) *NullableGetUsageRequest {
	return &NullableGetUsageRequest{value: val, isSet: true}
}

func (v NullableGetUsageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUsageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

