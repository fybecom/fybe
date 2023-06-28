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

// IpV4 struct for IpV4
type IpV4 struct {
	// IP Address
	Ip string `json:"ip"`
	// Netmask Classless Inter-Domain Routing
	NetmaskCidr int32 `json:"netmaskCidr"`
	// Gateway
	Gateway string `json:"gateway"`
}

// NewIpV4 instantiates a new IpV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpV4(ip string, netmaskCidr int32, gateway string) *IpV4 {
	this := IpV4{}
	this.Ip = ip
	this.NetmaskCidr = netmaskCidr
	this.Gateway = gateway
	return &this
}

// NewIpV4WithDefaults instantiates a new IpV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpV4WithDefaults() *IpV4 {
	this := IpV4{}
	return &this
}

// GetIp returns the Ip field value
func (o *IpV4) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *IpV4) GetIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *IpV4) SetIp(v string) {
	o.Ip = v
}

// GetNetmaskCidr returns the NetmaskCidr field value
func (o *IpV4) GetNetmaskCidr() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NetmaskCidr
}

// GetNetmaskCidrOk returns a tuple with the NetmaskCidr field value
// and a boolean to check if the value has been set.
func (o *IpV4) GetNetmaskCidrOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NetmaskCidr, true
}

// SetNetmaskCidr sets field value
func (o *IpV4) SetNetmaskCidr(v int32) {
	o.NetmaskCidr = v
}

// GetGateway returns the Gateway field value
func (o *IpV4) GetGateway() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value
// and a boolean to check if the value has been set.
func (o *IpV4) GetGatewayOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Gateway, true
}

// SetGateway sets field value
func (o *IpV4) SetGateway(v string) {
	o.Gateway = v
}

func (o IpV4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ip"] = o.Ip
	}
	if true {
		toSerialize["netmaskCidr"] = o.NetmaskCidr
	}
	if true {
		toSerialize["gateway"] = o.Gateway
	}
	return json.Marshal(toSerialize)
}

type NullableIpV4 struct {
	value *IpV4
	isSet bool
}

func (v NullableIpV4) Get() *IpV4 {
	return v.value
}

func (v *NullableIpV4) Set(val *IpV4) {
	v.value = val
	v.isSet = true
}

func (v NullableIpV4) IsSet() bool {
	return v.isSet
}

func (v *NullableIpV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpV4(val *IpV4) *NullableIpV4 {
	return &NullableIpV4{value: val, isSet: true}
}

func (v NullableIpV4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


