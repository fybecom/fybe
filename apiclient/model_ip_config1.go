/*
Fybe API Reference

# Introduction  The Fybe API facilitates resource management through HTTP requests. This documentation comprises a collection of HTTP endpoints that adhere to RESTful principles. Each endpoint is accompanied by descriptions, as well as request and response schemas.  ## OpenAPI specification (OAS)  Fybe's OpenAPI Specification can be [downloaded here](https://api.fybe.com/api-v1.yaml).  ## Getting Started  To utilize the Fybe API, you'll require following credentials that can be obtained from from the [Fybe Cockpit](https://cockpit.fybe.com/api/details): 1. ClientId 2. ClientSecret 3. API User (your email address to login to the [Fybe Cockpit](https://cockpit.fybe.com/api/details)) 4. API Password (this is a new password which you'll set or change in the [Fybe Cockpit](https://cockpit.fybe.com/api/details))  ### How to use the API?  As authentication [Bearer Tokens](https://oauth.net/2/bearer-tokens/) in form of [JWT](https://www.rfc-editor.org/rfc/rfc7519) are used. This approach follows the [OAuth 2.0](https://oauth.net/2/) specification.  #### Retrieve the Bearer Token  ```sh POST /auth/realms/Fybe/protocol/openid-connect/token HTTP/1.1 Host: auth.fybe.com  grant_type=password &password=xxxxxx &redirect_uri=https://example-app.com/redirect &username=xxxxxx &client_id=xxxxxx &client_secret=xxxxxx ```  The actual values for `client_id`, `client_secret`, `username` and `password` can be retrieved from [Fybe Cockpit](https://cockpit.fybe.com/api/details))  ## Requests  As stated below, the Fybe API accommodates HTTP requests. However, not all endpoints endorse every method. You can find a list of allowed methods within this documentation.  Method | Description ---    | --- GET    | When you need to obtain information about a resource, you should utilize the GET method. The data will be provided as a JSON object. It's essential to note that GET methods are read-only and don't impact any resources. POST   | To create a new object, send a POST method that encodes all necessary attributes in the request body as JSON. PATCH  | PATCH can be used to partially modify a resource, allowing specific attributes to be changed without updating the complete object. PUT    | If you need to update information about a resource, use the PUT method. PUT will overwrite existing values of the item, without considering the current state. DELETE | To remove a resource from your account, apply the DELETE method. If the resource is not found, the operation will generate a 4xx error along with a relevant message.  ## Responses  Typically, the Fybe API will respond to your requests by returning data in [JSON](https://www.json.org/) format, which can be easily processed using any programming language or tools.  Like for any HTTP requests, you'll receive an HTTP status code, which provides general information about the success or error of the request. The table below outlines some common HTTP status codes.  It's important to note that the description of the endpoints and methods doesn't provide an exhaustive list of all possible status codes. Only specific return codes and their corresponding response data are explicitly outlined.  Response Code | Description --- | --- 200 | The response will contain the information you requested. 201 | Your request has been acknowledged, and the resource has been created. 204 | Your request was successful, and no further information is provided in the response. 400 | Your request was not properly formed. 401 | You didn't provide valid authentication credentials. 402 | The request was declined as it necessitates an additional paid service. 403 | You are prohibited from performing the request. You'll need to pass a bearer token. 404 | No results were found for your request, or the resource you're trying to access does not exist. 409 | There is a conflict with resources, such as a violation of unique data constraints when attempting to create or modify resources. 429 | The rate limit has been reached. Please wait for some time before making further requests. 500 | We couldn't fulfill your request due to server-side issues. If this happens, please try again later or reach out to our support team.  Not every endpoint returns data. For example DELETE requests usually don't return any data. All others do return data. For easy handling the return values consists of metadata denoted with and underscore (\"_\") like `_links` or `_pagination`. The actual data is returned in a field called `data`. For convenience reasons this `data` field is always returned as an array even if it consists of only one single element.  DELETE requests usually don't return any data. Return values consists of metadata starting with an underscore (\"_\"), e.g. `_links` and `_pagination`. The requested data is to be found in the field `data`. The `data` field is always an array.   Please visit [Fybe](https://fybe.com) for more general information. 

API version: 1.0.0
Contact: support@fybe.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// IpConfig1 struct for IpConfig1
type IpConfig1 struct {
	V4 []IpV41 `json:"v4"`
	V6 []IpV61 `json:"v6"`
}

// NewIpConfig1 instantiates a new IpConfig1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpConfig1(v4 []IpV41, v6 []IpV61) *IpConfig1 {
	this := IpConfig1{}
	this.V4 = v4
	this.V6 = v6
	return &this
}

// NewIpConfig1WithDefaults instantiates a new IpConfig1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpConfig1WithDefaults() *IpConfig1 {
	this := IpConfig1{}
	return &this
}

// GetV4 returns the V4 field value
func (o *IpConfig1) GetV4() []IpV41 {
	if o == nil {
		var ret []IpV41
		return ret
	}

	return o.V4
}

// GetV4Ok returns a tuple with the V4 field value
// and a boolean to check if the value has been set.
func (o *IpConfig1) GetV4Ok() (*[]IpV41, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.V4, true
}

// SetV4 sets field value
func (o *IpConfig1) SetV4(v []IpV41) {
	o.V4 = v
}

// GetV6 returns the V6 field value
func (o *IpConfig1) GetV6() []IpV61 {
	if o == nil {
		var ret []IpV61
		return ret
	}

	return o.V6
}

// GetV6Ok returns a tuple with the V6 field value
// and a boolean to check if the value has been set.
func (o *IpConfig1) GetV6Ok() (*[]IpV61, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.V6, true
}

// SetV6 sets field value
func (o *IpConfig1) SetV6(v []IpV61) {
	o.V6 = v
}

func (o IpConfig1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["v4"] = o.V4
	}
	if true {
		toSerialize["v6"] = o.V6
	}
	return json.Marshal(toSerialize)
}

type NullableIpConfig1 struct {
	value *IpConfig1
	isSet bool
}

func (v NullableIpConfig1) Get() *IpConfig1 {
	return v.value
}

func (v *NullableIpConfig1) Set(val *IpConfig1) {
	v.value = val
	v.isSet = true
}

func (v NullableIpConfig1) IsSet() bool {
	return v.isSet
}

func (v *NullableIpConfig1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpConfig1(val *IpConfig1) *NullableIpConfig1 {
	return &NullableIpConfig1{value: val, isSet: true}
}

func (v NullableIpConfig1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpConfig1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


