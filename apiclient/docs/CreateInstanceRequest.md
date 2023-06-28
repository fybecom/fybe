# CreateInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | Pointer to **string** | Specify the ImageId to set up the compute instance. The default value is Ubuntu 22.04 (LTS). | [optional] [default to "a46a0297-7f23-41a5-b978-112e55019048"]
**ProductId** | **string** | Specify the product to set up the compute instance. The default value is V1 | [default to "V1"]
**Region** | Pointer to **string** | Specify the region in which the compute instance should be located. Default is us-east-1 | [optional] [default to "us-east-1"]
**SshKeys** | Pointer to **[]int64** | This parameter represents an array of secretIds corresponding to public SSH keys that allow logging in as the defaultUser with administrator/root privileges on Linux/BSD systems. For more information, please refer to the Secrets Management API. | [optional] 
**RootPassword** | Pointer to **int64** | The secretId field in this parameter refers to the password for the defaultUser with administrator/root privileges. Use SSH for Linux/BSD and RDP for Windows. For further details, please consult the Secrets Management API. | [optional] 
**UserData** | Pointer to **string** | You can customize the Cloud-Init configuration during the start of a compute instance. [Cloud-Init](https://cloud-init.io/)  | [optional] 
**License** | Pointer to **string** | To enhance your selected product, you may require an additional license, primarily for software licenses (excluding Windows). | [optional] 
**Period** | **int64** | The contract period is 1 month | [default to 1]
**DisplayName** | Pointer to **string** | Custom name of the compute instance. The name that appears on the screen to represent the compute instance. | [optional] 
**DefaultUser** | Pointer to **string** | The default username created for login with administrative privileges during (re-)installation depends on the operating system. For Linux/BSD, allowed values are admin (to be used with sudo for administrative privileges like root) or root. For Windows, allowed values are admin (with administrative privileges like administrator) or administrator. | [optional] [default to "admin"]
**AddOns** | Pointer to [**CreateInstanceAddons**](CreateInstanceAddons.md) | To add corresponding addons to the instance, specify their attributes in the addons object. | [optional] 

## Methods

### NewCreateInstanceRequest

`func NewCreateInstanceRequest(productId string, period int64, ) *CreateInstanceRequest`

NewCreateInstanceRequest instantiates a new CreateInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceRequestWithDefaults

`func NewCreateInstanceRequestWithDefaults() *CreateInstanceRequest`

NewCreateInstanceRequestWithDefaults instantiates a new CreateInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *CreateInstanceRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateInstanceRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateInstanceRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *CreateInstanceRequest) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetProductId

`func (o *CreateInstanceRequest) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreateInstanceRequest) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreateInstanceRequest) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetRegion

`func (o *CreateInstanceRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateInstanceRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateInstanceRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateInstanceRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSshKeys

`func (o *CreateInstanceRequest) GetSshKeys() []int64`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *CreateInstanceRequest) GetSshKeysOk() (*[]int64, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *CreateInstanceRequest) SetSshKeys(v []int64)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *CreateInstanceRequest) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetRootPassword

`func (o *CreateInstanceRequest) GetRootPassword() int64`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *CreateInstanceRequest) GetRootPasswordOk() (*int64, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *CreateInstanceRequest) SetRootPassword(v int64)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *CreateInstanceRequest) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetUserData

`func (o *CreateInstanceRequest) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreateInstanceRequest) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CreateInstanceRequest) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *CreateInstanceRequest) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetLicense

`func (o *CreateInstanceRequest) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *CreateInstanceRequest) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *CreateInstanceRequest) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *CreateInstanceRequest) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetPeriod

`func (o *CreateInstanceRequest) GetPeriod() int64`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CreateInstanceRequest) GetPeriodOk() (*int64, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CreateInstanceRequest) SetPeriod(v int64)`

SetPeriod sets Period field to given value.


### GetDisplayName

`func (o *CreateInstanceRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateInstanceRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateInstanceRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateInstanceRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDefaultUser

`func (o *CreateInstanceRequest) GetDefaultUser() string`

GetDefaultUser returns the DefaultUser field if non-nil, zero value otherwise.

### GetDefaultUserOk

`func (o *CreateInstanceRequest) GetDefaultUserOk() (*string, bool)`

GetDefaultUserOk returns a tuple with the DefaultUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUser

`func (o *CreateInstanceRequest) SetDefaultUser(v string)`

SetDefaultUser sets DefaultUser field to given value.

### HasDefaultUser

`func (o *CreateInstanceRequest) HasDefaultUser() bool`

HasDefaultUser returns a boolean if a field has been set.

### GetAddOns

`func (o *CreateInstanceRequest) GetAddOns() CreateInstanceAddons`

GetAddOns returns the AddOns field if non-nil, zero value otherwise.

### GetAddOnsOk

`func (o *CreateInstanceRequest) GetAddOnsOk() (*CreateInstanceAddons, bool)`

GetAddOnsOk returns a tuple with the AddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOns

`func (o *CreateInstanceRequest) SetAddOns(v CreateInstanceAddons)`

SetAddOns sets AddOns field to given value.

### HasAddOns

`func (o *CreateInstanceRequest) HasAddOns() bool`

HasAddOns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


