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

// AddCustomerPaymentMethod200Response struct for AddCustomerPaymentMethod200Response
type AddCustomerPaymentMethod200Response struct {
	Url string `json:"url"`
}

// NewAddCustomerPaymentMethod200Response instantiates a new AddCustomerPaymentMethod200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCustomerPaymentMethod200Response(url string) *AddCustomerPaymentMethod200Response {
	this := AddCustomerPaymentMethod200Response{}
	this.Url = url
	return &this
}

// NewAddCustomerPaymentMethod200ResponseWithDefaults instantiates a new AddCustomerPaymentMethod200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCustomerPaymentMethod200ResponseWithDefaults() *AddCustomerPaymentMethod200Response {
	this := AddCustomerPaymentMethod200Response{}
	return &this
}

// GetUrl returns the Url field value
func (o *AddCustomerPaymentMethod200Response) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AddCustomerPaymentMethod200Response) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AddCustomerPaymentMethod200Response) SetUrl(v string) {
	o.Url = v
}

func (o AddCustomerPaymentMethod200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableAddCustomerPaymentMethod200Response struct {
	value *AddCustomerPaymentMethod200Response
	isSet bool
}

func (v NullableAddCustomerPaymentMethod200Response) Get() *AddCustomerPaymentMethod200Response {
	return v.value
}

func (v *NullableAddCustomerPaymentMethod200Response) Set(val *AddCustomerPaymentMethod200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCustomerPaymentMethod200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCustomerPaymentMethod200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCustomerPaymentMethod200Response(val *AddCustomerPaymentMethod200Response) *NullableAddCustomerPaymentMethod200Response {
	return &NullableAddCustomerPaymentMethod200Response{value: val, isSet: true}
}

func (v NullableAddCustomerPaymentMethod200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCustomerPaymentMethod200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


