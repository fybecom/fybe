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

// CustomImagesStatsResponseData struct for CustomImagesStatsResponseData
type CustomImagesStatsResponseData struct {
	// Your customer tenant id
	TenantId string `json:"tenantId"`
	// Your customer number
	CustomerId string `json:"customerId"`
	// How many custom images currently exist
	CurrentImagesCount float32 `json:"currentImagesCount"`
	// The amount of space that is currently free on the disk, measured in megabytes.
	TotalSizeMb float32 `json:"totalSizeMb"`
	// The amount of space that is currently occupied on the disk, measured in megabytes.
	UsedSizeMb float32 `json:"usedSizeMb"`
	// The amount of space that is currently available on the disk, measured in megabytes.
	FreeSizeMb float32 `json:"freeSizeMb"`
}

// NewCustomImagesStatsResponseData instantiates a new CustomImagesStatsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomImagesStatsResponseData(tenantId string, customerId string, currentImagesCount float32, totalSizeMb float32, usedSizeMb float32, freeSizeMb float32) *CustomImagesStatsResponseData {
	this := CustomImagesStatsResponseData{}
	this.TenantId = tenantId
	this.CustomerId = customerId
	this.CurrentImagesCount = currentImagesCount
	this.TotalSizeMb = totalSizeMb
	this.UsedSizeMb = usedSizeMb
	this.FreeSizeMb = freeSizeMb
	return &this
}

// NewCustomImagesStatsResponseDataWithDefaults instantiates a new CustomImagesStatsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomImagesStatsResponseDataWithDefaults() *CustomImagesStatsResponseData {
	this := CustomImagesStatsResponseData{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *CustomImagesStatsResponseData) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *CustomImagesStatsResponseData) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *CustomImagesStatsResponseData) SetTenantId(v string) {
	o.TenantId = v
}

// GetCustomerId returns the CustomerId field value
func (o *CustomImagesStatsResponseData) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CustomImagesStatsResponseData) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CustomImagesStatsResponseData) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetCurrentImagesCount returns the CurrentImagesCount field value
func (o *CustomImagesStatsResponseData) GetCurrentImagesCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CurrentImagesCount
}

// GetCurrentImagesCountOk returns a tuple with the CurrentImagesCount field value
// and a boolean to check if the value has been set.
func (o *CustomImagesStatsResponseData) GetCurrentImagesCountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CurrentImagesCount, true
}

// SetCurrentImagesCount sets field value
func (o *CustomImagesStatsResponseData) SetCurrentImagesCount(v float32) {
	o.CurrentImagesCount = v
}

// GetTotalSizeMb returns the TotalSizeMb field value
func (o *CustomImagesStatsResponseData) GetTotalSizeMb() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalSizeMb
}

// GetTotalSizeMbOk returns a tuple with the TotalSizeMb field value
// and a boolean to check if the value has been set.
func (o *CustomImagesStatsResponseData) GetTotalSizeMbOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalSizeMb, true
}

// SetTotalSizeMb sets field value
func (o *CustomImagesStatsResponseData) SetTotalSizeMb(v float32) {
	o.TotalSizeMb = v
}

// GetUsedSizeMb returns the UsedSizeMb field value
func (o *CustomImagesStatsResponseData) GetUsedSizeMb() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UsedSizeMb
}

// GetUsedSizeMbOk returns a tuple with the UsedSizeMb field value
// and a boolean to check if the value has been set.
func (o *CustomImagesStatsResponseData) GetUsedSizeMbOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UsedSizeMb, true
}

// SetUsedSizeMb sets field value
func (o *CustomImagesStatsResponseData) SetUsedSizeMb(v float32) {
	o.UsedSizeMb = v
}

// GetFreeSizeMb returns the FreeSizeMb field value
func (o *CustomImagesStatsResponseData) GetFreeSizeMb() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.FreeSizeMb
}

// GetFreeSizeMbOk returns a tuple with the FreeSizeMb field value
// and a boolean to check if the value has been set.
func (o *CustomImagesStatsResponseData) GetFreeSizeMbOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FreeSizeMb, true
}

// SetFreeSizeMb sets field value
func (o *CustomImagesStatsResponseData) SetFreeSizeMb(v float32) {
	o.FreeSizeMb = v
}

func (o CustomImagesStatsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["customerId"] = o.CustomerId
	}
	if true {
		toSerialize["currentImagesCount"] = o.CurrentImagesCount
	}
	if true {
		toSerialize["totalSizeMb"] = o.TotalSizeMb
	}
	if true {
		toSerialize["usedSizeMb"] = o.UsedSizeMb
	}
	if true {
		toSerialize["freeSizeMb"] = o.FreeSizeMb
	}
	return json.Marshal(toSerialize)
}

type NullableCustomImagesStatsResponseData struct {
	value *CustomImagesStatsResponseData
	isSet bool
}

func (v NullableCustomImagesStatsResponseData) Get() *CustomImagesStatsResponseData {
	return v.value
}

func (v *NullableCustomImagesStatsResponseData) Set(val *CustomImagesStatsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomImagesStatsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomImagesStatsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomImagesStatsResponseData(val *CustomImagesStatsResponseData) *NullableCustomImagesStatsResponseData {
	return &NullableCustomImagesStatsResponseData{value: val, isSet: true}
}

func (v NullableCustomImagesStatsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomImagesStatsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


