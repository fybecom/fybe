/*
Fybe API Reference

# Introduction  The Fybe API facilitates resource management through HTTP requests. This documentation comprises a collection of HTTP endpoints that adhere to RESTful principles. Each endpoint is accompanied by descriptions, as well as request and response schemas.  ## OpenAPI specification (OAS)  Fybe's OpenAPI Specification can be [downloaded here](https://api.fybe.com/api-v1.yaml).  ## Getting Started  To utilize the Fybe API, you'll require following credentials that can be obtained from from the [Fybe Cockpit](https://cockpit.fybe.com/account/security): 1. ClientId 2. ClientSecret 3. API User (your email address to login to the [Fybe Cockpit](https://cockpit.fybe.com/account/security)) 4. API Password (this is a new password which you'll set or change in the [Fybe Cockpit](https://cockpit.fybe.com/account/security))  ### How to use the API?  As authentication [Bearer Tokens](https://oauth.net/2/bearer-tokens/) in form of [JWT](https://www.rfc-editor.org/rfc/rfc7519) are used. This approach follows the [OAuth 2.0](https://oauth.net/2/) specification.  #### Retrieve the Bearer Token  ```sh POST /auth/realms/Fybe/protocol/openid-connect/token HTTP/1.1 Host: airlock.fybe.com  grant_type=password &password=xxxxxx &redirect_uri=https://example-app.com/redirect &username=xxxxxx &client_id=xxxxxx &client_secret=xxxxxx ```  The actual values for `client_id`, `client_secret`, `username` and `password` can be retrieved from [Fybe Cockpit](https://cockpit.fybe.com/account/security)  ## Using the Command-Line Interface (CLI)  Fybe provides the CLI called `fybe` which can be downloaded from <https://github.com/fybecom/fybe>. Please follow the instructions in the `README.md` to perform the installation on your OS. `fybe` supports Windows, macOS and Linux operating systems.  Using `fybe` CLI invoking makes invoking the API much more comfortable. E.g. `fybe get instances` for retrieving the list of compute instances.  ## Requests  As stated below, the Fybe API accommodates HTTP requests. However, not all endpoints endorse every method. You can find a list of allowed methods within this documentation.  Method | Description ---    | --- GET    | When you need to obtain information about a resource, you should utilize the GET method. The data will be provided as a JSON object. It's essential to note that GET methods are read-only and don't impact any resources. POST   | To create a new object, send a POST method that encodes all necessary attributes in the request body as JSON. PATCH  | PATCH can be used to partially modify a resource, allowing specific attributes to be changed without updating the complete object. PUT    | If you need to update information about a resource, use the PUT method. PUT will overwrite existing values of the item, without considering the current state. DELETE | To remove a resource from your account, apply the DELETE method. If the resource is not found, the operation will generate a 4xx error along with a relevant message.  ## Responses  Typically, the Fybe API will respond to your requests by returning data in [JSON](https://www.json.org/) format, which can be easily processed using any programming language or tools.  Like for any HTTP requests, you'll receive an HTTP status code, which provides general information about the success or error of the request. The table below outlines some common HTTP status codes.  It's important to note that the description of the endpoints and methods doesn't provide an exhaustive list of all possible status codes. Only specific return codes and their corresponding response data are explicitly outlined.  Response Code | Description --- | --- 200 | The response will contain the information you requested. 201 | Your request has been acknowledged, and the resource has been created. 204 | Your request was successful, and no further information is provided in the response. 400 | Your request was not properly formed. 401 | You didn't provide valid authentication credentials. 402 | The request was declined as it necessitates an additional paid service. 403 | You are prohibited from performing the request. You'll need to pass a bearer token. 404 | No results were found for your request, or the resource you're trying to access does not exist. 409 | There is a conflict with resources, such as a violation of unique data constraints when attempting to create or modify resources. 429 | The rate limit has been reached. Please wait for some time before making further requests. 500 | We couldn't fulfill your request due to server-side issues. If this happens, please try again later or reach out to our support team.  Not every endpoint returns data. For example DELETE requests usually don't return any data. All others do return data. For easy handling the return values consists of metadata denoted with and underscore (\"_\") like `_links` or `_pagination`. The actual data is returned in a field called `data`. For convenience reasons this `data` field is always returned as an array even if it consists of only one single element.  DELETE requests usually don't return any data. Return values consists of metadata starting with an underscore (\"_\"), e.g. `_links` and `_pagination`. The requested data is to be found in the field `data`. The `data` field is always an array.   Please visit [Fybe](https://fybe.com) for more general information. 

API version: 1.0.0
Contact: support@fybe.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// RoleResponse struct for RoleResponse
type RoleResponse struct {
	// Role ID
	RoleId int64 `json:"roleId"`
	// Tenant ID
	TenantId string `json:"tenantId"`
	// Customer ID
	CustomerId string `json:"customerId"`
	// Role Name
	Name string `json:"name"`
	// Admin flag
	Admin bool `json:"admin"`
	// Access All Resources flag
	AccessAllResources bool `json:"accessAllResources"`
	// The role type can be either default or custom. Default roles cannot be modified or deleted.
	Type string `json:"type"`
	Permissions *[]PermissionResponse `json:"permissions,omitempty"`
}

// NewRoleResponse instantiates a new RoleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleResponse(roleId int64, tenantId string, customerId string, name string, admin bool, accessAllResources bool, type_ string) *RoleResponse {
	this := RoleResponse{}
	this.RoleId = roleId
	this.TenantId = tenantId
	this.CustomerId = customerId
	this.Name = name
	this.Admin = admin
	this.AccessAllResources = accessAllResources
	this.Type = type_
	return &this
}

// NewRoleResponseWithDefaults instantiates a new RoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleResponseWithDefaults() *RoleResponse {
	this := RoleResponse{}
	return &this
}

// GetRoleId returns the RoleId field value
func (o *RoleResponse) GetRoleId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *RoleResponse) GetRoleIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *RoleResponse) SetRoleId(v int64) {
	o.RoleId = v
}

// GetTenantId returns the TenantId field value
func (o *RoleResponse) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *RoleResponse) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *RoleResponse) SetTenantId(v string) {
	o.TenantId = v
}

// GetCustomerId returns the CustomerId field value
func (o *RoleResponse) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *RoleResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *RoleResponse) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetName returns the Name field value
func (o *RoleResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoleResponse) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoleResponse) SetName(v string) {
	o.Name = v
}

// GetAdmin returns the Admin field value
func (o *RoleResponse) GetAdmin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Admin
}

// GetAdminOk returns a tuple with the Admin field value
// and a boolean to check if the value has been set.
func (o *RoleResponse) GetAdminOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Admin, true
}

// SetAdmin sets field value
func (o *RoleResponse) SetAdmin(v bool) {
	o.Admin = v
}

// GetAccessAllResources returns the AccessAllResources field value
func (o *RoleResponse) GetAccessAllResources() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AccessAllResources
}

// GetAccessAllResourcesOk returns a tuple with the AccessAllResources field value
// and a boolean to check if the value has been set.
func (o *RoleResponse) GetAccessAllResourcesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessAllResources, true
}

// SetAccessAllResources sets field value
func (o *RoleResponse) SetAccessAllResources(v bool) {
	o.AccessAllResources = v
}

// GetType returns the Type field value
func (o *RoleResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoleResponse) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RoleResponse) SetType(v string) {
	o.Type = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *RoleResponse) GetPermissions() []PermissionResponse {
	if o == nil || o.Permissions == nil {
		var ret []PermissionResponse
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleResponse) GetPermissionsOk() (*[]PermissionResponse, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *RoleResponse) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []PermissionResponse and assigns it to the Permissions field.
func (o *RoleResponse) SetPermissions(v []PermissionResponse) {
	o.Permissions = &v
}

func (o RoleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["roleId"] = o.RoleId
	}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["customerId"] = o.CustomerId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["admin"] = o.Admin
	}
	if true {
		toSerialize["accessAllResources"] = o.AccessAllResources
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableRoleResponse struct {
	value *RoleResponse
	isSet bool
}

func (v NullableRoleResponse) Get() *RoleResponse {
	return v.value
}

func (v *NullableRoleResponse) Set(val *RoleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleResponse(val *RoleResponse) *NullableRoleResponse {
	return &NullableRoleResponse{value: val, isSet: true}
}

func (v NullableRoleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


