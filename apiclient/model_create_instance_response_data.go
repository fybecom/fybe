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
	"time"
)

// CreateInstanceResponseData struct for CreateInstanceResponseData
type CreateInstanceResponseData struct {
	// Your customer tenant id
	TenantId string `json:"tenantId"`
	// Your customer number
	CustomerId string `json:"customerId"`
	// Identifier of the instance
	InstanceId int64 `json:"instanceId"`
	// The date on which the instance was created.
	CreatedDate time.Time `json:"createdDate"`
	// Identifier of the image
	ImageId string `json:"imageId"`
	// Product identifier
	ProductId string `json:"productId"`
	// The region where the compute instance should be located.
	Region string `json:"region"`
	AddOns []AddOnResponse `json:"addOns"`
	// The category of operating system (OS).
	OsType string `json:"osType"`
	Status InstanceStatus `json:"status"`
	// A collection of `secretId` values representing public SSH keys that can be used to log in as the `defaultUser` with administrator or root privileges. This is applicable for Linux/BSD systems, and additional information can be found in the Secrets Management API.
	SshKeys []int64 `json:"sshKeys"`
}

// NewCreateInstanceResponseData instantiates a new CreateInstanceResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInstanceResponseData(tenantId string, customerId string, instanceId int64, createdDate time.Time, imageId string, productId string, region string, addOns []AddOnResponse, osType string, status InstanceStatus, sshKeys []int64) *CreateInstanceResponseData {
	this := CreateInstanceResponseData{}
	this.TenantId = tenantId
	this.CustomerId = customerId
	this.InstanceId = instanceId
	this.CreatedDate = createdDate
	this.ImageId = imageId
	this.ProductId = productId
	this.Region = region
	this.AddOns = addOns
	this.OsType = osType
	this.Status = status
	this.SshKeys = sshKeys
	return &this
}

// NewCreateInstanceResponseDataWithDefaults instantiates a new CreateInstanceResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInstanceResponseDataWithDefaults() *CreateInstanceResponseData {
	this := CreateInstanceResponseData{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *CreateInstanceResponseData) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *CreateInstanceResponseData) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *CreateInstanceResponseData) SetTenantId(v string) {
	o.TenantId = v
}

// GetCustomerId returns the CustomerId field value
func (o *CreateInstanceResponseData) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CreateInstanceResponseData) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CreateInstanceResponseData) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetInstanceId returns the InstanceId field value
func (o *CreateInstanceResponseData) GetInstanceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *CreateInstanceResponseData) GetInstanceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *CreateInstanceResponseData) SetInstanceId(v int64) {
	o.InstanceId = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *CreateInstanceResponseData) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *CreateInstanceResponseData) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *CreateInstanceResponseData) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetImageId returns the ImageId field value
func (o *CreateInstanceResponseData) GetImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value
// and a boolean to check if the value has been set.
func (o *CreateInstanceResponseData) GetImageIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImageId, true
}

// SetImageId sets field value
func (o *CreateInstanceResponseData) SetImageId(v string) {
	o.ImageId = v
}

// GetProductId returns the ProductId field value
func (o *CreateInstanceResponseData) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *CreateInstanceResponseData) GetProductIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *CreateInstanceResponseData) SetProductId(v string) {
	o.ProductId = v
}

// GetRegion returns the Region field value
func (o *CreateInstanceResponseData) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CreateInstanceResponseData) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *CreateInstanceResponseData) SetRegion(v string) {
	o.Region = v
}

// GetAddOns returns the AddOns field value
func (o *CreateInstanceResponseData) GetAddOns() []AddOnResponse {
	if o == nil {
		var ret []AddOnResponse
		return ret
	}

	return o.AddOns
}

// GetAddOnsOk returns a tuple with the AddOns field value
// and a boolean to check if the value has been set.
func (o *CreateInstanceResponseData) GetAddOnsOk() (*[]AddOnResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AddOns, true
}

// SetAddOns sets field value
func (o *CreateInstanceResponseData) SetAddOns(v []AddOnResponse) {
	o.AddOns = v
}

// GetOsType returns the OsType field value
func (o *CreateInstanceResponseData) GetOsType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value
// and a boolean to check if the value has been set.
func (o *CreateInstanceResponseData) GetOsTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OsType, true
}

// SetOsType sets field value
func (o *CreateInstanceResponseData) SetOsType(v string) {
	o.OsType = v
}

// GetStatus returns the Status field value
func (o *CreateInstanceResponseData) GetStatus() InstanceStatus {
	if o == nil {
		var ret InstanceStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateInstanceResponseData) GetStatusOk() (*InstanceStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateInstanceResponseData) SetStatus(v InstanceStatus) {
	o.Status = v
}

// GetSshKeys returns the SshKeys field value
func (o *CreateInstanceResponseData) GetSshKeys() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value
// and a boolean to check if the value has been set.
func (o *CreateInstanceResponseData) GetSshKeysOk() (*[]int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SshKeys, true
}

// SetSshKeys sets field value
func (o *CreateInstanceResponseData) SetSshKeys(v []int64) {
	o.SshKeys = v
}

func (o CreateInstanceResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["customerId"] = o.CustomerId
	}
	if true {
		toSerialize["instanceId"] = o.InstanceId
	}
	if true {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if true {
		toSerialize["imageId"] = o.ImageId
	}
	if true {
		toSerialize["productId"] = o.ProductId
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["addOns"] = o.AddOns
	}
	if true {
		toSerialize["osType"] = o.OsType
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["sshKeys"] = o.SshKeys
	}
	return json.Marshal(toSerialize)
}

type NullableCreateInstanceResponseData struct {
	value *CreateInstanceResponseData
	isSet bool
}

func (v NullableCreateInstanceResponseData) Get() *CreateInstanceResponseData {
	return v.value
}

func (v *NullableCreateInstanceResponseData) Set(val *CreateInstanceResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInstanceResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInstanceResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInstanceResponseData(val *CreateInstanceResponseData) *NullableCreateInstanceResponseData {
	return &NullableCreateInstanceResponseData{value: val, isSet: true}
}

func (v NullableCreateInstanceResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInstanceResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


