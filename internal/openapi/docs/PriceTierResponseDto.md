# PriceTierResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The tier price type | 
**Amount** | **float32** | The tier price amount | 
**Start** | Pointer to **float32** | The tier start amount | [optional] 
**Finish** | Pointer to **float32** | The tier finish amount (infinity indicated by -1) | [optional] 
**PackageSize** | Pointer to **float32** | The tier package size (only applicable to package price tiers) | [optional] 

## Methods

### NewPriceTierResponseDto

`func NewPriceTierResponseDto(type_ string, amount float32, ) *PriceTierResponseDto`

NewPriceTierResponseDto instantiates a new PriceTierResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceTierResponseDtoWithDefaults

`func NewPriceTierResponseDtoWithDefaults() *PriceTierResponseDto`

NewPriceTierResponseDtoWithDefaults instantiates a new PriceTierResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceTierResponseDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceTierResponseDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceTierResponseDto) SetType(v string)`

SetType sets Type field to given value.


### GetAmount

`func (o *PriceTierResponseDto) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PriceTierResponseDto) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PriceTierResponseDto) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetStart

`func (o *PriceTierResponseDto) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PriceTierResponseDto) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PriceTierResponseDto) SetStart(v float32)`

SetStart sets Start field to given value.

### HasStart

`func (o *PriceTierResponseDto) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetFinish

`func (o *PriceTierResponseDto) GetFinish() float32`

GetFinish returns the Finish field if non-nil, zero value otherwise.

### GetFinishOk

`func (o *PriceTierResponseDto) GetFinishOk() (*float32, bool)`

GetFinishOk returns a tuple with the Finish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinish

`func (o *PriceTierResponseDto) SetFinish(v float32)`

SetFinish sets Finish field to given value.

### HasFinish

`func (o *PriceTierResponseDto) HasFinish() bool`

HasFinish returns a boolean if a field has been set.

### GetPackageSize

`func (o *PriceTierResponseDto) GetPackageSize() float32`

GetPackageSize returns the PackageSize field if non-nil, zero value otherwise.

### GetPackageSizeOk

`func (o *PriceTierResponseDto) GetPackageSizeOk() (*float32, bool)`

GetPackageSizeOk returns a tuple with the PackageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSize

`func (o *PriceTierResponseDto) SetPackageSize(v float32)`

SetPackageSize sets PackageSize field to given value.

### HasPackageSize

`func (o *PriceTierResponseDto) HasPackageSize() bool`

HasPackageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


