/*
Kable API

The Kable API allows developers to manage customers, plans, and usage data for their API.  ## API Host and Environments  Kable is accessible in Live and Test environments for authentication and metering of client API requests. You will have separate API keys to access each environment.  You should only use Kable's Live environment for your own production data. All other configured environments should use Kable's Test environment.  ## API Protocols and Headers  All requests to the Kable API are made over HTTPS TLS v1.2+ to ensure security. Calls made over HTTP will fail. Any requests without proper authentication will also fail.  The Kable API uses standard JSON for requests and responses. Be sure to set both the `Content-Type` and `Accept` headers on each request to `application/json`.  Each Kable API response includes a `requestId` as the `X-REQUEST-ID` response header. The `requestId` is included on most responses regardless whether the API request succeeded or failed. You can use this `requestId` to help with debugging or when contacting support regarding a specific API call.  ## API Versioning  All Kable endpoints are versioned. After the host, each API can be found at `/api/vX/...` where X is the API version.  We strive to ensure that changes to the Kable API are backward compatible. Sometimes, though, we must break from older design paradigms to make the product better. When this happens, a new version of the API is released.  The current version of the Kable API is `v1`.  

API version: 1.0.0
Contact: contact@kable.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ErrorMessage500 struct for ErrorMessage500
type ErrorMessage500 struct {
	// Human-readable error message
	Message string `json:"message"`
}

// NewErrorMessage500 instantiates a new ErrorMessage500 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorMessage500(message string) *ErrorMessage500 {
	this := ErrorMessage500{}
	this.Message = message
	return &this
}

// NewErrorMessage500WithDefaults instantiates a new ErrorMessage500 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorMessage500WithDefaults() *ErrorMessage500 {
	this := ErrorMessage500{}
	return &this
}

// GetMessage returns the Message field value
func (o *ErrorMessage500) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorMessage500) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorMessage500) SetMessage(v string) {
	o.Message = v
}

func (o ErrorMessage500) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableErrorMessage500 struct {
	value *ErrorMessage500
	isSet bool
}

func (v NullableErrorMessage500) Get() *ErrorMessage500 {
	return v.value
}

func (v *NullableErrorMessage500) Set(val *ErrorMessage500) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorMessage500) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorMessage500) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorMessage500(val *ErrorMessage500) *NullableErrorMessage500 {
	return &NullableErrorMessage500{value: val, isSet: true}
}

func (v NullableErrorMessage500) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorMessage500) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


