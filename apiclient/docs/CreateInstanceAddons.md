# CreateInstanceAddons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalIps** | Pointer to **map[string]interface{}** | Set this attribute if you want to upgrade your instance with the Additional IPs addon. Please provide an empty object for the time being as value. There will be more configuration possible in the future. | [optional] 

## Methods

### NewCreateInstanceAddons

`func NewCreateInstanceAddons() *CreateInstanceAddons`

NewCreateInstanceAddons instantiates a new CreateInstanceAddons object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceAddonsWithDefaults

`func NewCreateInstanceAddonsWithDefaults() *CreateInstanceAddons`

NewCreateInstanceAddonsWithDefaults instantiates a new CreateInstanceAddons object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalIps

`func (o *CreateInstanceAddons) GetAdditionalIps() map[string]interface{}`

GetAdditionalIps returns the AdditionalIps field if non-nil, zero value otherwise.

### GetAdditionalIpsOk

`func (o *CreateInstanceAddons) GetAdditionalIpsOk() (*map[string]interface{}, bool)`

GetAdditionalIpsOk returns a tuple with the AdditionalIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalIps

`func (o *CreateInstanceAddons) SetAdditionalIps(v map[string]interface{})`

SetAdditionalIps sets AdditionalIps field to given value.

### HasAdditionalIps

`func (o *CreateInstanceAddons) HasAdditionalIps() bool`

HasAdditionalIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


