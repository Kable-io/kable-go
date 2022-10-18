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

// PriceTier struct for PriceTier
type PriceTier struct {
	// Price type for the tier.
	Type string `json:"type"`
	// Cost for the tier, evaluated in the context of its associated type.
	Amount float32 `json:"amount"`
	// The usage metric at which this tier pricing starts.
	Start *float32 `json:"start,omitempty"`
	// The usage metric at which this tier pricing ends. A value of -1 represents infinity.
	Finish *float32 `json:"finish,omitempty"`
	// The package size, applicable only for package pricing (ie $1 for every 10 requests).
	PackageSize *float32 `json:"packageSize,omitempty"`
}

// NewPriceTier instantiates a new PriceTier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceTier(type_ string, amount float32) *PriceTier {
	this := PriceTier{}
	this.Type = type_
	this.Amount = amount
	return &this
}

// NewPriceTierWithDefaults instantiates a new PriceTier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceTierWithDefaults() *PriceTier {
	this := PriceTier{}
	return &this
}

// GetType returns the Type field value
func (o *PriceTier) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PriceTier) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PriceTier) SetType(v string) {
	o.Type = v
}

// GetAmount returns the Amount field value
func (o *PriceTier) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PriceTier) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PriceTier) SetAmount(v float32) {
	o.Amount = v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *PriceTier) GetStart() float32 {
	if o == nil || o.Start == nil {
		var ret float32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceTier) GetStartOk() (*float32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *PriceTier) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given float32 and assigns it to the Start field.
func (o *PriceTier) SetStart(v float32) {
	o.Start = &v
}

// GetFinish returns the Finish field value if set, zero value otherwise.
func (o *PriceTier) GetFinish() float32 {
	if o == nil || o.Finish == nil {
		var ret float32
		return ret
	}
	return *o.Finish
}

// GetFinishOk returns a tuple with the Finish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceTier) GetFinishOk() (*float32, bool) {
	if o == nil || o.Finish == nil {
		return nil, false
	}
	return o.Finish, true
}

// HasFinish returns a boolean if a field has been set.
func (o *PriceTier) HasFinish() bool {
	if o != nil && o.Finish != nil {
		return true
	}

	return false
}

// SetFinish gets a reference to the given float32 and assigns it to the Finish field.
func (o *PriceTier) SetFinish(v float32) {
	o.Finish = &v
}

// GetPackageSize returns the PackageSize field value if set, zero value otherwise.
func (o *PriceTier) GetPackageSize() float32 {
	if o == nil || o.PackageSize == nil {
		var ret float32
		return ret
	}
	return *o.PackageSize
}

// GetPackageSizeOk returns a tuple with the PackageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceTier) GetPackageSizeOk() (*float32, bool) {
	if o == nil || o.PackageSize == nil {
		return nil, false
	}
	return o.PackageSize, true
}

// HasPackageSize returns a boolean if a field has been set.
func (o *PriceTier) HasPackageSize() bool {
	if o != nil && o.PackageSize != nil {
		return true
	}

	return false
}

// SetPackageSize gets a reference to the given float32 and assigns it to the PackageSize field.
func (o *PriceTier) SetPackageSize(v float32) {
	o.PackageSize = &v
}

func (o PriceTier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.Finish != nil {
		toSerialize["finish"] = o.Finish
	}
	if o.PackageSize != nil {
		toSerialize["packageSize"] = o.PackageSize
	}
	return json.Marshal(toSerialize)
}

type NullablePriceTier struct {
	value *PriceTier
	isSet bool
}

func (v NullablePriceTier) Get() *PriceTier {
	return v.value
}

func (v *NullablePriceTier) Set(val *PriceTier) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceTier) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceTier(val *PriceTier) *NullablePriceTier {
	return &NullablePriceTier{value: val, isSet: true}
}

func (v NullablePriceTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


