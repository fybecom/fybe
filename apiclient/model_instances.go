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

// Instances struct for Instances
type Instances struct {
	// ID of the Instance
	InstanceId int64 `json:"instanceId"`
	// Display Name of the Instance
	DisplayName string `json:"displayName"`
	// Name of the Instance
	Name string `json:"name"`
	// Id of the Product
	ProductId string `json:"productId"`
	PrivateIpConfig PrivateIpConfig `json:"privateIpConfig"`
	IpConfig IpConfig `json:"ipConfig"`
	// Instance Status in Virtual Private Cloud
	Status string `json:"status"`
	// Error Message.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// NewInstances instantiates a new Instances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstances(instanceId int64, displayName string, name string, productId string, privateIpConfig PrivateIpConfig, ipConfig IpConfig, status string) *Instances {
	this := Instances{}
	this.InstanceId = instanceId
	this.DisplayName = displayName
	this.Name = name
	this.ProductId = productId
	this.PrivateIpConfig = privateIpConfig
	this.IpConfig = ipConfig
	this.Status = status
	return &this
}

// NewInstancesWithDefaults instantiates a new Instances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstancesWithDefaults() *Instances {
	this := Instances{}
	return &this
}

// GetInstanceId returns the InstanceId field value
func (o *Instances) GetInstanceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *Instances) GetInstanceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *Instances) SetInstanceId(v int64) {
	o.InstanceId = v
}

// GetDisplayName returns the DisplayName field value
func (o *Instances) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *Instances) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *Instances) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetName returns the Name field value
func (o *Instances) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Instances) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Instances) SetName(v string) {
	o.Name = v
}

// GetProductId returns the ProductId field value
func (o *Instances) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *Instances) GetProductIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *Instances) SetProductId(v string) {
	o.ProductId = v
}

// GetPrivateIpConfig returns the PrivateIpConfig field value
func (o *Instances) GetPrivateIpConfig() PrivateIpConfig {
	if o == nil {
		var ret PrivateIpConfig
		return ret
	}

	return o.PrivateIpConfig
}

// GetPrivateIpConfigOk returns a tuple with the PrivateIpConfig field value
// and a boolean to check if the value has been set.
func (o *Instances) GetPrivateIpConfigOk() (*PrivateIpConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PrivateIpConfig, true
}

// SetPrivateIpConfig sets field value
func (o *Instances) SetPrivateIpConfig(v PrivateIpConfig) {
	o.PrivateIpConfig = v
}

// GetIpConfig returns the IpConfig field value
func (o *Instances) GetIpConfig() IpConfig {
	if o == nil {
		var ret IpConfig
		return ret
	}

	return o.IpConfig
}

// GetIpConfigOk returns a tuple with the IpConfig field value
// and a boolean to check if the value has been set.
func (o *Instances) GetIpConfigOk() (*IpConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IpConfig, true
}

// SetIpConfig sets field value
func (o *Instances) SetIpConfig(v IpConfig) {
	o.IpConfig = v
}

// GetStatus returns the Status field value
func (o *Instances) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Instances) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Instances) SetStatus(v string) {
	o.Status = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *Instances) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instances) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *Instances) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *Instances) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o Instances) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["instanceId"] = o.InstanceId
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["productId"] = o.ProductId
	}
	if true {
		toSerialize["privateIpConfig"] = o.PrivateIpConfig
	}
	if true {
		toSerialize["ipConfig"] = o.IpConfig
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	return json.Marshal(toSerialize)
}

type NullableInstances struct {
	value *Instances
	isSet bool
}

func (v NullableInstances) Get() *Instances {
	return v.value
}

func (v *NullableInstances) Set(val *Instances) {
	v.value = val
	v.isSet = true
}

func (v NullableInstances) IsSet() bool {
	return v.isSet
}

func (v *NullableInstances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstances(val *Instances) *NullableInstances {
	return &NullableInstances{value: val, isSet: true}
}

func (v NullableInstances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


