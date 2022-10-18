# CreateCreditGrantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantAmount** | **float32** | The number of credits to create in the grant. | 
**Description** | Pointer to **string** | A human-readable description that explains why the credit grant is being created. | [optional] 
**ChargeAmount** | Pointer to **float32** | An optional amount to charge the customer (in the customer&#39;s configured &#x60;currency&#x60;) before issuing a credit grant.  When the &#x60;chargeAmount&#x60; field is supplied, Kable will attempt to charge the customer immediately through Stripe before issuing the grant. Therefore, when specifying a &#x60;chargeAmount&#x60;, the customer must have a Stripe integration defined and have a default payment card on-file.  | [optional] 

## Methods

### NewCreateCreditGrantRequest

`func NewCreateCreditGrantRequest(grantAmount float32, ) *CreateCreditGrantRequest`

NewCreateCreditGrantRequest instantiates a new CreateCreditGrantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCreditGrantRequestWithDefaults

`func NewCreateCreditGrantRequestWithDefaults() *CreateCreditGrantRequest`

NewCreateCreditGrantRequestWithDefaults instantiates a new CreateCreditGrantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantAmount

`func (o *CreateCreditGrantRequest) GetGrantAmount() float32`

GetGrantAmount returns the GrantAmount field if non-nil, zero value otherwise.

### GetGrantAmountOk

`func (o *CreateCreditGrantRequest) GetGrantAmountOk() (*float32, bool)`

GetGrantAmountOk returns a tuple with the GrantAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantAmount

`func (o *CreateCreditGrantRequest) SetGrantAmount(v float32)`

SetGrantAmount sets GrantAmount field to given value.


### GetDescription

`func (o *CreateCreditGrantRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCreditGrantRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCreditGrantRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCreditGrantRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetChargeAmount

`func (o *CreateCreditGrantRequest) GetChargeAmount() float32`

GetChargeAmount returns the ChargeAmount field if non-nil, zero value otherwise.

### GetChargeAmountOk

`func (o *CreateCreditGrantRequest) GetChargeAmountOk() (*float32, bool)`

GetChargeAmountOk returns a tuple with the ChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmount

`func (o *CreateCreditGrantRequest) SetChargeAmount(v float32)`

SetChargeAmount sets ChargeAmount field to given value.

### HasChargeAmount

`func (o *CreateCreditGrantRequest) HasChargeAmount() bool`

HasChargeAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


