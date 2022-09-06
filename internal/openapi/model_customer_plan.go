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

// CustomerPlan struct for CustomerPlan
type CustomerPlan struct {
	// A Kable-defined identifier for the plan.
	PlanId string `json:"planId"`
	// An identifier of the dimension along which usage is aggregated in this plan.
	DimensionId string `json:"dimensionId"`
	// A human-readable name for the plan, visible on dashboards, invoices, and reports.
	Name string `json:"name"`
	// A human-readable description for the plan, visible on dashboards, invoices, and reports.
	Description *string `json:"description,omitempty"`
	// An additional identifier for the plan, defined by you, that is *not* visible to customers. If you have different *Monthly Active Users Plan*s for different cohorts of customers, this is a useful place to differentiate those.
	Nickname *string `json:"nickname,omitempty"`
	// Date the plan was added to the customer.
	AddedDate string `json:"addedDate"`
	// Date the plan ended or will end for trials or pilot programs.
	EndedDate *string `json:"endedDate,omitempty"`
	// Date the current invoice cycle begins.
	PeriodStartDate *string `json:"periodStartDate,omitempty"`
	// Date the current invoice cycle ends.
	PeriodEndDate *string `json:"periodEndDate,omitempty"`
	// An optional identifier of any \"next\" plan after the termination of a free trials or pilot programs.
	NextPlanId *string `json:"nextPlanId,omitempty"`
}

// NewCustomerPlan instantiates a new CustomerPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPlan(planId string, dimensionId string, name string, addedDate string) *CustomerPlan {
	this := CustomerPlan{}
	this.PlanId = planId
	this.DimensionId = dimensionId
	this.Name = name
	this.AddedDate = addedDate
	return &this
}

// NewCustomerPlanWithDefaults instantiates a new CustomerPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPlanWithDefaults() *CustomerPlan {
	this := CustomerPlan{}
	return &this
}

// GetPlanId returns the PlanId field value
func (o *CustomerPlan) GetPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *CustomerPlan) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *CustomerPlan) SetPlanId(v string) {
	o.PlanId = v
}

// GetDimensionId returns the DimensionId field value
func (o *CustomerPlan) GetDimensionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DimensionId
}

// GetDimensionIdOk returns a tuple with the DimensionId field value
// and a boolean to check if the value has been set.
func (o *CustomerPlan) GetDimensionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DimensionId, true
}

// SetDimensionId sets field value
func (o *CustomerPlan) SetDimensionId(v string) {
	o.DimensionId = v
}

// GetName returns the Name field value
func (o *CustomerPlan) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomerPlan) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomerPlan) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomerPlan) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPlan) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomerPlan) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomerPlan) SetDescription(v string) {
	o.Description = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *CustomerPlan) GetNickname() string {
	if o == nil || o.Nickname == nil {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPlan) GetNicknameOk() (*string, bool) {
	if o == nil || o.Nickname == nil {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *CustomerPlan) HasNickname() bool {
	if o != nil && o.Nickname != nil {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *CustomerPlan) SetNickname(v string) {
	o.Nickname = &v
}

// GetAddedDate returns the AddedDate field value
func (o *CustomerPlan) GetAddedDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddedDate
}

// GetAddedDateOk returns a tuple with the AddedDate field value
// and a boolean to check if the value has been set.
func (o *CustomerPlan) GetAddedDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddedDate, true
}

// SetAddedDate sets field value
func (o *CustomerPlan) SetAddedDate(v string) {
	o.AddedDate = v
}

// GetEndedDate returns the EndedDate field value if set, zero value otherwise.
func (o *CustomerPlan) GetEndedDate() string {
	if o == nil || o.EndedDate == nil {
		var ret string
		return ret
	}
	return *o.EndedDate
}

// GetEndedDateOk returns a tuple with the EndedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPlan) GetEndedDateOk() (*string, bool) {
	if o == nil || o.EndedDate == nil {
		return nil, false
	}
	return o.EndedDate, true
}

// HasEndedDate returns a boolean if a field has been set.
func (o *CustomerPlan) HasEndedDate() bool {
	if o != nil && o.EndedDate != nil {
		return true
	}

	return false
}

// SetEndedDate gets a reference to the given string and assigns it to the EndedDate field.
func (o *CustomerPlan) SetEndedDate(v string) {
	o.EndedDate = &v
}

// GetPeriodStartDate returns the PeriodStartDate field value if set, zero value otherwise.
func (o *CustomerPlan) GetPeriodStartDate() string {
	if o == nil || o.PeriodStartDate == nil {
		var ret string
		return ret
	}
	return *o.PeriodStartDate
}

// GetPeriodStartDateOk returns a tuple with the PeriodStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPlan) GetPeriodStartDateOk() (*string, bool) {
	if o == nil || o.PeriodStartDate == nil {
		return nil, false
	}
	return o.PeriodStartDate, true
}

// HasPeriodStartDate returns a boolean if a field has been set.
func (o *CustomerPlan) HasPeriodStartDate() bool {
	if o != nil && o.PeriodStartDate != nil {
		return true
	}

	return false
}

// SetPeriodStartDate gets a reference to the given string and assigns it to the PeriodStartDate field.
func (o *CustomerPlan) SetPeriodStartDate(v string) {
	o.PeriodStartDate = &v
}

// GetPeriodEndDate returns the PeriodEndDate field value if set, zero value otherwise.
func (o *CustomerPlan) GetPeriodEndDate() string {
	if o == nil || o.PeriodEndDate == nil {
		var ret string
		return ret
	}
	return *o.PeriodEndDate
}

// GetPeriodEndDateOk returns a tuple with the PeriodEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPlan) GetPeriodEndDateOk() (*string, bool) {
	if o == nil || o.PeriodEndDate == nil {
		return nil, false
	}
	return o.PeriodEndDate, true
}

// HasPeriodEndDate returns a boolean if a field has been set.
func (o *CustomerPlan) HasPeriodEndDate() bool {
	if o != nil && o.PeriodEndDate != nil {
		return true
	}

	return false
}

// SetPeriodEndDate gets a reference to the given string and assigns it to the PeriodEndDate field.
func (o *CustomerPlan) SetPeriodEndDate(v string) {
	o.PeriodEndDate = &v
}

// GetNextPlanId returns the NextPlanId field value if set, zero value otherwise.
func (o *CustomerPlan) GetNextPlanId() string {
	if o == nil || o.NextPlanId == nil {
		var ret string
		return ret
	}
	return *o.NextPlanId
}

// GetNextPlanIdOk returns a tuple with the NextPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPlan) GetNextPlanIdOk() (*string, bool) {
	if o == nil || o.NextPlanId == nil {
		return nil, false
	}
	return o.NextPlanId, true
}

// HasNextPlanId returns a boolean if a field has been set.
func (o *CustomerPlan) HasNextPlanId() bool {
	if o != nil && o.NextPlanId != nil {
		return true
	}

	return false
}

// SetNextPlanId gets a reference to the given string and assigns it to the NextPlanId field.
func (o *CustomerPlan) SetNextPlanId(v string) {
	o.NextPlanId = &v
}

func (o CustomerPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["planId"] = o.PlanId
	}
	if true {
		toSerialize["dimensionId"] = o.DimensionId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Nickname != nil {
		toSerialize["nickname"] = o.Nickname
	}
	if true {
		toSerialize["addedDate"] = o.AddedDate
	}
	if o.EndedDate != nil {
		toSerialize["endedDate"] = o.EndedDate
	}
	if o.PeriodStartDate != nil {
		toSerialize["periodStartDate"] = o.PeriodStartDate
	}
	if o.PeriodEndDate != nil {
		toSerialize["periodEndDate"] = o.PeriodEndDate
	}
	if o.NextPlanId != nil {
		toSerialize["nextPlanId"] = o.NextPlanId
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerPlan struct {
	value *CustomerPlan
	isSet bool
}

func (v NullableCustomerPlan) Get() *CustomerPlan {
	return v.value
}

func (v *NullableCustomerPlan) Set(val *CustomerPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPlan(val *CustomerPlan) *NullableCustomerPlan {
	return &NullableCustomerPlan{value: val, isSet: true}
}

func (v NullableCustomerPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


