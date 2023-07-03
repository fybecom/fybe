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

// PrivateNetworkResponse struct for PrivateNetworkResponse
type PrivateNetworkResponse struct {
	// Your customer tenant id
	TenantId string `json:"tenantId"`
	// Your customer number
	CustomerId string `json:"customerId"`
	// ID of the Virtual Private Cloud
	PrivateNetworkId int64 `json:"privateNetworkId"`
	// Location of the data center of the Virtual Private Cloud
	DataCenter string `json:"dataCenter"`
	// Region Slug of your Virtual Private Cloud location
	Region string `json:"region"`
	// Region of the Virtual Private Cloud.
	RegionName string `json:"regionName"`
	// Virtual Private Cloud Name
	Name string `json:"name"`
	// Detailed Description of the Virtual Private Cloud
	Description string `json:"description"`
	// The Range of the Classless Inter-Domain Routing of the Virtual Private Cloud
	Cidr string `json:"cidr"`
	// Number of IPs of the Virtual Private Cloud
	AvailableIps int64 `json:"availableIps"`
	// Creation date of the Virtual Private Cloud
	CreatedDate time.Time `json:"createdDate"`
	Instances []Instances `json:"instances"`
}

// NewPrivateNetworkResponse instantiates a new PrivateNetworkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateNetworkResponse(tenantId string, customerId string, privateNetworkId int64, dataCenter string, region string, regionName string, name string, description string, cidr string, availableIps int64, createdDate time.Time, instances []Instances) *PrivateNetworkResponse {
	this := PrivateNetworkResponse{}
	this.TenantId = tenantId
	this.CustomerId = customerId
	this.PrivateNetworkId = privateNetworkId
	this.DataCenter = dataCenter
	this.Region = region
	this.RegionName = regionName
	this.Name = name
	this.Description = description
	this.Cidr = cidr
	this.AvailableIps = availableIps
	this.CreatedDate = createdDate
	this.Instances = instances
	return &this
}

// NewPrivateNetworkResponseWithDefaults instantiates a new PrivateNetworkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateNetworkResponseWithDefaults() *PrivateNetworkResponse {
	this := PrivateNetworkResponse{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *PrivateNetworkResponse) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkResponse) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *PrivateNetworkResponse) SetTenantId(v string) {
	o.TenantId = v
}

// GetCustomerId returns the CustomerId field value
func (o *PrivateNetworkResponse) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *PrivateNetworkResponse) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetPrivateNetworkId returns the PrivateNetworkId field value
func (o *PrivateNetworkResponse) GetPrivateNetworkId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PrivateNetworkId
}

// GetPrivateNetworkIdOk returns a tuple with the PrivateNetworkId field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkResponse) GetPrivateNetworkIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PrivateNetworkId, true
}

// SetPrivateNetworkId sets field value
func (o *PrivateNetworkResponse) SetPrivateNetworkId(v int64) {
	o.PrivateNetworkId = v
}

// GetDataCenter returns the DataCenter field value
func (o *PrivateNetworkResponse) GetDataCenter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataCenter
}

// GetDataCenterOk returns a tuple with the DataCenter field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkResponse) GetDataCenterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DataCenter, true
}

// SetDataCenter sets field value
func (o *PrivateNetworkResponse) SetDataCenter(v string) {
	o.DataCenter = v
}

// GetRegion returns the Region field value
func (o *PrivateNetworkResponse) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkResponse) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *PrivateNetworkResponse) SetRegion(v string) {
	o.Region = v
}

// GetRegionName returns the RegionName field value
func (o *PrivateNetworkResponse) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkResponse) GetRegionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *PrivateNetworkResponse) SetRegionName(v string) {
	o.RegionName = v
}

// GetName returns the Name field value
func (o *PrivateNetworkResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkResponse) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PrivateNetworkResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *PrivateNetworkResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkResponse) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PrivateNetworkResponse) SetDescription(v string) {
	o.Description = v
}

// GetCidr returns the Cidr field value
func (o *PrivateNetworkResponse) GetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkResponse) GetCidrOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *PrivateNetworkResponse) SetCidr(v string) {
	o.Cidr = v
}

// GetAvailableIps returns the AvailableIps field value
func (o *PrivateNetworkResponse) GetAvailableIps() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AvailableIps
}

// GetAvailableIpsOk returns a tuple with the AvailableIps field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkResponse) GetAvailableIpsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AvailableIps, true
}

// SetAvailableIps sets field value
func (o *PrivateNetworkResponse) SetAvailableIps(v int64) {
	o.AvailableIps = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *PrivateNetworkResponse) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkResponse) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *PrivateNetworkResponse) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetInstances returns the Instances field value
func (o *PrivateNetworkResponse) GetInstances() []Instances {
	if o == nil {
		var ret []Instances
		return ret
	}

	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkResponse) GetInstancesOk() (*[]Instances, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Instances, true
}

// SetInstances sets field value
func (o *PrivateNetworkResponse) SetInstances(v []Instances) {
	o.Instances = v
}

func (o PrivateNetworkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["customerId"] = o.CustomerId
	}
	if true {
		toSerialize["privateNetworkId"] = o.PrivateNetworkId
	}
	if true {
		toSerialize["dataCenter"] = o.DataCenter
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["regionName"] = o.RegionName
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["cidr"] = o.Cidr
	}
	if true {
		toSerialize["availableIps"] = o.AvailableIps
	}
	if true {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if true {
		toSerialize["instances"] = o.Instances
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateNetworkResponse struct {
	value *PrivateNetworkResponse
	isSet bool
}

func (v NullablePrivateNetworkResponse) Get() *PrivateNetworkResponse {
	return v.value
}

func (v *NullablePrivateNetworkResponse) Set(val *PrivateNetworkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateNetworkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateNetworkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateNetworkResponse(val *PrivateNetworkResponse) *NullablePrivateNetworkResponse {
	return &NullablePrivateNetworkResponse{value: val, isSet: true}
}

func (v NullablePrivateNetworkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateNetworkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


