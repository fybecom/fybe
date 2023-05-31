# CreatePrivateNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | Virtual Private Cloud location. Default is &#x60;US&#x60; | [optional] [default to "US"]
**Name** | **string** |  | 
**Description** | Pointer to **string** | Thorough explanation of the Virtual Private Cloud. | [optional] 

## Methods

### NewCreatePrivateNetworkRequest

`func NewCreatePrivateNetworkRequest(name string, ) *CreatePrivateNetworkRequest`

NewCreatePrivateNetworkRequest instantiates a new CreatePrivateNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePrivateNetworkRequestWithDefaults

`func NewCreatePrivateNetworkRequestWithDefaults() *CreatePrivateNetworkRequest`

NewCreatePrivateNetworkRequestWithDefaults instantiates a new CreatePrivateNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *CreatePrivateNetworkRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreatePrivateNetworkRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreatePrivateNetworkRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreatePrivateNetworkRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetName

`func (o *CreatePrivateNetworkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePrivateNetworkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePrivateNetworkRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreatePrivateNetworkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePrivateNetworkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePrivateNetworkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePrivateNetworkRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


