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

// InstancesResetPasswordActionsRequest struct for InstancesResetPasswordActionsRequest
type InstancesResetPasswordActionsRequest struct {
	// Array of `secretId`s of public SSH keys for logging into as `defaultUser` with administrator/root privileges. Applies to Linux/BSD systems. Please refer to Secrets Management API.
	SshKeys *[]int64 `json:"sshKeys,omitempty"`
	// `secretId` of the password for the `defaultUser` with administrator/root privileges. For Linux/BSD please use SSH, for Windows RDP. Please refer to Secrets Management API.
	RootPassword *int64 `json:"rootPassword,omitempty"`
	// [Cloud-Init](https://cloud-init.io/) Config in order to customize during start of compute instance.
	UserData *string `json:"userData,omitempty"`
}

// NewInstancesResetPasswordActionsRequest instantiates a new InstancesResetPasswordActionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstancesResetPasswordActionsRequest() *InstancesResetPasswordActionsRequest {
	this := InstancesResetPasswordActionsRequest{}
	return &this
}

// NewInstancesResetPasswordActionsRequestWithDefaults instantiates a new InstancesResetPasswordActionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstancesResetPasswordActionsRequestWithDefaults() *InstancesResetPasswordActionsRequest {
	this := InstancesResetPasswordActionsRequest{}
	return &this
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *InstancesResetPasswordActionsRequest) GetSshKeys() []int64 {
	if o == nil || o.SshKeys == nil {
		var ret []int64
		return ret
	}
	return *o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesResetPasswordActionsRequest) GetSshKeysOk() (*[]int64, bool) {
	if o == nil || o.SshKeys == nil {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *InstancesResetPasswordActionsRequest) HasSshKeys() bool {
	if o != nil && o.SshKeys != nil {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []int64 and assigns it to the SshKeys field.
func (o *InstancesResetPasswordActionsRequest) SetSshKeys(v []int64) {
	o.SshKeys = &v
}

// GetRootPassword returns the RootPassword field value if set, zero value otherwise.
func (o *InstancesResetPasswordActionsRequest) GetRootPassword() int64 {
	if o == nil || o.RootPassword == nil {
		var ret int64
		return ret
	}
	return *o.RootPassword
}

// GetRootPasswordOk returns a tuple with the RootPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesResetPasswordActionsRequest) GetRootPasswordOk() (*int64, bool) {
	if o == nil || o.RootPassword == nil {
		return nil, false
	}
	return o.RootPassword, true
}

// HasRootPassword returns a boolean if a field has been set.
func (o *InstancesResetPasswordActionsRequest) HasRootPassword() bool {
	if o != nil && o.RootPassword != nil {
		return true
	}

	return false
}

// SetRootPassword gets a reference to the given int64 and assigns it to the RootPassword field.
func (o *InstancesResetPasswordActionsRequest) SetRootPassword(v int64) {
	o.RootPassword = &v
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *InstancesResetPasswordActionsRequest) GetUserData() string {
	if o == nil || o.UserData == nil {
		var ret string
		return ret
	}
	return *o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesResetPasswordActionsRequest) GetUserDataOk() (*string, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *InstancesResetPasswordActionsRequest) HasUserData() bool {
	if o != nil && o.UserData != nil {
		return true
	}

	return false
}

// SetUserData gets a reference to the given string and assigns it to the UserData field.
func (o *InstancesResetPasswordActionsRequest) SetUserData(v string) {
	o.UserData = &v
}

func (o InstancesResetPasswordActionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SshKeys != nil {
		toSerialize["sshKeys"] = o.SshKeys
	}
	if o.RootPassword != nil {
		toSerialize["rootPassword"] = o.RootPassword
	}
	if o.UserData != nil {
		toSerialize["userData"] = o.UserData
	}
	return json.Marshal(toSerialize)
}

type NullableInstancesResetPasswordActionsRequest struct {
	value *InstancesResetPasswordActionsRequest
	isSet bool
}

func (v NullableInstancesResetPasswordActionsRequest) Get() *InstancesResetPasswordActionsRequest {
	return v.value
}

func (v *NullableInstancesResetPasswordActionsRequest) Set(val *InstancesResetPasswordActionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInstancesResetPasswordActionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInstancesResetPasswordActionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstancesResetPasswordActionsRequest(val *InstancesResetPasswordActionsRequest) *NullableInstancesResetPasswordActionsRequest {
	return &NullableInstancesResetPasswordActionsRequest{value: val, isSet: true}
}

func (v NullableInstancesResetPasswordActionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstancesResetPasswordActionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


