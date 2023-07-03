# UpgradeObjectStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScaling** | Pointer to [**UpgradeAutoScalingType**](UpgradeAutoScalingType.md) | If autoscaling is enabled, this is the new monthly limit for the size of object storage. | [optional] 
**TotalPurchasedSpaceTB** | Pointer to **float64** | There is a new overall limit on the amount of object storage that can be used. If this limit is greater than the previous one, any additional storage space used will be charged. Please note that it is not possible to downgrade the limit once it has been increased. | [optional] 

## Methods

### NewUpgradeObjectStorageRequest

`func NewUpgradeObjectStorageRequest() *UpgradeObjectStorageRequest`

NewUpgradeObjectStorageRequest instantiates a new UpgradeObjectStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeObjectStorageRequestWithDefaults

`func NewUpgradeObjectStorageRequestWithDefaults() *UpgradeObjectStorageRequest`

NewUpgradeObjectStorageRequestWithDefaults instantiates a new UpgradeObjectStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScaling

`func (o *UpgradeObjectStorageRequest) GetAutoScaling() UpgradeAutoScalingType`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *UpgradeObjectStorageRequest) GetAutoScalingOk() (*UpgradeAutoScalingType, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *UpgradeObjectStorageRequest) SetAutoScaling(v UpgradeAutoScalingType)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *UpgradeObjectStorageRequest) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetTotalPurchasedSpaceTB

`func (o *UpgradeObjectStorageRequest) GetTotalPurchasedSpaceTB() float64`

GetTotalPurchasedSpaceTB returns the TotalPurchasedSpaceTB field if non-nil, zero value otherwise.

### GetTotalPurchasedSpaceTBOk

`func (o *UpgradeObjectStorageRequest) GetTotalPurchasedSpaceTBOk() (*float64, bool)`

GetTotalPurchasedSpaceTBOk returns a tuple with the TotalPurchasedSpaceTB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPurchasedSpaceTB

`func (o *UpgradeObjectStorageRequest) SetTotalPurchasedSpaceTB(v float64)`

SetTotalPurchasedSpaceTB sets TotalPurchasedSpaceTB field to given value.

### HasTotalPurchasedSpaceTB

`func (o *UpgradeObjectStorageRequest) HasTotalPurchasedSpaceTB() bool`

HasTotalPurchasedSpaceTB returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


