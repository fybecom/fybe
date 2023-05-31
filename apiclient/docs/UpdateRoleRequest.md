# UpdateRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | This refers to the title of the role. Please note that the role name cannot exceed 255 characters. | 
**Admin** | **bool** | If a user is an admin, they will have permissions to access all API endpoints and resources. By enabling this option, all role definitions and accessAllResources will be overridden. | 
**AccessAllResources** | **bool** | Grant access to all resources. This will override all assigned resources associated with a particular role. | 
**Permissions** | Pointer to [**[]PermissionRequest**](PermissionRequest.md) |  | [optional] 

## Methods

### NewUpdateRoleRequest

`func NewUpdateRoleRequest(name string, admin bool, accessAllResources bool, ) *UpdateRoleRequest`

NewUpdateRoleRequest instantiates a new UpdateRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleRequestWithDefaults

`func NewUpdateRoleRequestWithDefaults() *UpdateRoleRequest`

NewUpdateRoleRequestWithDefaults instantiates a new UpdateRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRoleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAdmin

`func (o *UpdateRoleRequest) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *UpdateRoleRequest) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *UpdateRoleRequest) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.


### GetAccessAllResources

`func (o *UpdateRoleRequest) GetAccessAllResources() bool`

GetAccessAllResources returns the AccessAllResources field if non-nil, zero value otherwise.

### GetAccessAllResourcesOk

`func (o *UpdateRoleRequest) GetAccessAllResourcesOk() (*bool, bool)`

GetAccessAllResourcesOk returns a tuple with the AccessAllResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAllResources

`func (o *UpdateRoleRequest) SetAccessAllResources(v bool)`

SetAccessAllResources sets AccessAllResources field to given value.


### GetPermissions

`func (o *UpdateRoleRequest) GetPermissions() []PermissionRequest`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateRoleRequest) GetPermissionsOk() (*[]PermissionRequest, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateRoleRequest) SetPermissions(v []PermissionRequest)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UpdateRoleRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


