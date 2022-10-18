# PriceTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Price type for the tier. | 
**Amount** | **float32** | Cost for the tier, evaluated in the context of its associated type. | 
**Start** | Pointer to **float32** | The usage metric at which this tier pricing starts. | [optional] 
**Finish** | Pointer to **float32** | The usage metric at which this tier pricing ends. A value of -1 represents infinity. | [optional] 
**PackageSize** | Pointer to **float32** | The package size, applicable only for package pricing (ie $1 for every 10 requests). | [optional] 

## Methods

### NewPriceTier

`func NewPriceTier(type_ string, amount float32, ) *PriceTier`

NewPriceTier instantiates a new PriceTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceTierWithDefaults

`func NewPriceTierWithDefaults() *PriceTier`

NewPriceTierWithDefaults instantiates a new PriceTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceTier) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceTier) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceTier) SetType(v string)`

SetType sets Type field to given value.


### GetAmount

`func (o *PriceTier) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PriceTier) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PriceTier) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetStart

`func (o *PriceTier) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PriceTier) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PriceTier) SetStart(v float32)`

SetStart sets Start field to given value.

### HasStart

`func (o *PriceTier) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetFinish

`func (o *PriceTier) GetFinish() float32`

GetFinish returns the Finish field if non-nil, zero value otherwise.

### GetFinishOk

`func (o *PriceTier) GetFinishOk() (*float32, bool)`

GetFinishOk returns a tuple with the Finish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinish

`func (o *PriceTier) SetFinish(v float32)`

SetFinish sets Finish field to given value.

### HasFinish

`func (o *PriceTier) HasFinish() bool`

HasFinish returns a boolean if a field has been set.

### GetPackageSize

`func (o *PriceTier) GetPackageSize() float32`

GetPackageSize returns the PackageSize field if non-nil, zero value otherwise.

### GetPackageSizeOk

`func (o *PriceTier) GetPackageSizeOk() (*float32, bool)`

GetPackageSizeOk returns a tuple with the PackageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSize

`func (o *PriceTier) SetPackageSize(v float32)`

SetPackageSize sets PackageSize field to given value.

### HasPackageSize

`func (o *PriceTier) HasPackageSize() bool`

HasPackageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


