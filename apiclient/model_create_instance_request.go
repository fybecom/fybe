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

// CreateInstanceRequest struct for CreateInstanceRequest
type CreateInstanceRequest struct {
	// Specify the ImageId to set up the compute instance. The default value is Ubuntu 22.04 (LTS).
	ImageId *string `json:"imageId,omitempty"`
	// Specify the product to set up the compute instance. The default value is V1
	ProductId string `json:"productId"`
	// Specify the region in which the compute instance should be located. Default is us-east-1
	Region *string `json:"region,omitempty"`
	// This parameter represents an array of secretIds corresponding to public SSH keys that allow logging in as the defaultUser with administrator/root privileges on Linux/BSD systems. For more information, please refer to the Secrets Management API.
	SshKeys *[]int64 `json:"sshKeys,omitempty"`
	// The secretId field in this parameter refers to the password for the defaultUser with administrator/root privileges. Use SSH for Linux/BSD and RDP for Windows. For further details, please consult the Secrets Management API.
	RootPassword *int64 `json:"rootPassword,omitempty"`
	// You can customize the Cloud-Init configuration during the start of a compute instance. [Cloud-Init](https://cloud-init.io/) 
	UserData *string `json:"userData,omitempty"`
	// To enhance your selected product, you may require an additional license, primarily for software licenses (excluding Windows).
	License *string `json:"license,omitempty"`
	// The contract period is 1 month
	Period int64 `json:"period"`
	// Custom name of the compute instance. The name that appears on the screen to represent the compute instance.
	DisplayName *string `json:"displayName,omitempty"`
	// The default username created for login with administrative privileges during (re-)installation depends on the operating system. For Linux/BSD, allowed values are admin (to be used with sudo for administrative privileges like root) or root. For Windows, allowed values are admin (with administrative privileges like administrator) or administrator.
	DefaultUser *string `json:"defaultUser,omitempty"`
	// To add corresponding addons to the instance, specify their attributes in the addons object.
	AddOns *CreateInstanceAddons `json:"addOns,omitempty"`
}

// NewCreateInstanceRequest instantiates a new CreateInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInstanceRequest(productId string, period int64) *CreateInstanceRequest {
	this := CreateInstanceRequest{}
	var imageId string = "a46a0297-7f23-41a5-b978-112e55019048"
	this.ImageId = &imageId
	this.ProductId = productId
	var region string = "us-east-1"
	this.Region = &region
	this.Period = period
	var defaultUser string = "admin"
	this.DefaultUser = &defaultUser
	return &this
}

// NewCreateInstanceRequestWithDefaults instantiates a new CreateInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInstanceRequestWithDefaults() *CreateInstanceRequest {
	this := CreateInstanceRequest{}
	var imageId string = "a46a0297-7f23-41a5-b978-112e55019048"
	this.ImageId = &imageId
	var productId string = "V1"
	this.ProductId = productId
	var region string = "us-east-1"
	this.Region = &region
	var period int64 = 1
	this.Period = period
	var defaultUser string = "admin"
	this.DefaultUser = &defaultUser
	return &this
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetImageId() string {
	if o == nil || o.ImageId == nil {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetImageIdOk() (*string, bool) {
	if o == nil || o.ImageId == nil {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasImageId() bool {
	if o != nil && o.ImageId != nil {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *CreateInstanceRequest) SetImageId(v string) {
	o.ImageId = &v
}

// GetProductId returns the ProductId field value
func (o *CreateInstanceRequest) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetProductIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *CreateInstanceRequest) SetProductId(v string) {
	o.ProductId = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CreateInstanceRequest) SetRegion(v string) {
	o.Region = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetSshKeys() []int64 {
	if o == nil || o.SshKeys == nil {
		var ret []int64
		return ret
	}
	return *o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetSshKeysOk() (*[]int64, bool) {
	if o == nil || o.SshKeys == nil {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasSshKeys() bool {
	if o != nil && o.SshKeys != nil {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []int64 and assigns it to the SshKeys field.
func (o *CreateInstanceRequest) SetSshKeys(v []int64) {
	o.SshKeys = &v
}

// GetRootPassword returns the RootPassword field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetRootPassword() int64 {
	if o == nil || o.RootPassword == nil {
		var ret int64
		return ret
	}
	return *o.RootPassword
}

// GetRootPasswordOk returns a tuple with the RootPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetRootPasswordOk() (*int64, bool) {
	if o == nil || o.RootPassword == nil {
		return nil, false
	}
	return o.RootPassword, true
}

// HasRootPassword returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasRootPassword() bool {
	if o != nil && o.RootPassword != nil {
		return true
	}

	return false
}

// SetRootPassword gets a reference to the given int64 and assigns it to the RootPassword field.
func (o *CreateInstanceRequest) SetRootPassword(v int64) {
	o.RootPassword = &v
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetUserData() string {
	if o == nil || o.UserData == nil {
		var ret string
		return ret
	}
	return *o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetUserDataOk() (*string, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasUserData() bool {
	if o != nil && o.UserData != nil {
		return true
	}

	return false
}

// SetUserData gets a reference to the given string and assigns it to the UserData field.
func (o *CreateInstanceRequest) SetUserData(v string) {
	o.UserData = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetLicense() string {
	if o == nil || o.License == nil {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetLicenseOk() (*string, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasLicense() bool {
	if o != nil && o.License != nil {
		return true
	}

	return false
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *CreateInstanceRequest) SetLicense(v string) {
	o.License = &v
}

// GetPeriod returns the Period field value
func (o *CreateInstanceRequest) GetPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetPeriodOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *CreateInstanceRequest) SetPeriod(v int64) {
	o.Period = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CreateInstanceRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDefaultUser returns the DefaultUser field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetDefaultUser() string {
	if o == nil || o.DefaultUser == nil {
		var ret string
		return ret
	}
	return *o.DefaultUser
}

// GetDefaultUserOk returns a tuple with the DefaultUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetDefaultUserOk() (*string, bool) {
	if o == nil || o.DefaultUser == nil {
		return nil, false
	}
	return o.DefaultUser, true
}

// HasDefaultUser returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasDefaultUser() bool {
	if o != nil && o.DefaultUser != nil {
		return true
	}

	return false
}

// SetDefaultUser gets a reference to the given string and assigns it to the DefaultUser field.
func (o *CreateInstanceRequest) SetDefaultUser(v string) {
	o.DefaultUser = &v
}

// GetAddOns returns the AddOns field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetAddOns() CreateInstanceAddons {
	if o == nil || o.AddOns == nil {
		var ret CreateInstanceAddons
		return ret
	}
	return *o.AddOns
}

// GetAddOnsOk returns a tuple with the AddOns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetAddOnsOk() (*CreateInstanceAddons, bool) {
	if o == nil || o.AddOns == nil {
		return nil, false
	}
	return o.AddOns, true
}

// HasAddOns returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasAddOns() bool {
	if o != nil && o.AddOns != nil {
		return true
	}

	return false
}

// SetAddOns gets a reference to the given CreateInstanceAddons and assigns it to the AddOns field.
func (o *CreateInstanceRequest) SetAddOns(v CreateInstanceAddons) {
	o.AddOns = &v
}

func (o CreateInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageId != nil {
		toSerialize["imageId"] = o.ImageId
	}
	if true {
		toSerialize["productId"] = o.ProductId
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.SshKeys != nil {
		toSerialize["sshKeys"] = o.SshKeys
	}
	if o.RootPassword != nil {
		toSerialize["rootPassword"] = o.RootPassword
	}
	if o.UserData != nil {
		toSerialize["userData"] = o.UserData
	}
	if o.License != nil {
		toSerialize["license"] = o.License
	}
	if true {
		toSerialize["period"] = o.Period
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.DefaultUser != nil {
		toSerialize["defaultUser"] = o.DefaultUser
	}
	if o.AddOns != nil {
		toSerialize["addOns"] = o.AddOns
	}
	return json.Marshal(toSerialize)
}

type NullableCreateInstanceRequest struct {
	value *CreateInstanceRequest
	isSet bool
}

func (v NullableCreateInstanceRequest) Get() *CreateInstanceRequest {
	return v.value
}

func (v *NullableCreateInstanceRequest) Set(val *CreateInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInstanceRequest(val *CreateInstanceRequest) *NullableCreateInstanceRequest {
	return &NullableCreateInstanceRequest{value: val, isSet: true}
}

func (v NullableCreateInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


