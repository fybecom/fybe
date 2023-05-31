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

// TagAssignmentSelfLinks struct for TagAssignmentSelfLinks
type TagAssignmentSelfLinks struct {
	// Link to current resource.
	Self string `json:"self"`
	// Link to the related tag.
	Tag string `json:"tag"`
	// Link to the assigned resource
	Resource string `json:"_resource"`
}

// NewTagAssignmentSelfLinks instantiates a new TagAssignmentSelfLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagAssignmentSelfLinks(self string, tag string, resource string) *TagAssignmentSelfLinks {
	this := TagAssignmentSelfLinks{}
	this.Self = self
	this.Tag = tag
	this.Resource = resource
	return &this
}

// NewTagAssignmentSelfLinksWithDefaults instantiates a new TagAssignmentSelfLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagAssignmentSelfLinksWithDefaults() *TagAssignmentSelfLinks {
	this := TagAssignmentSelfLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *TagAssignmentSelfLinks) GetSelf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *TagAssignmentSelfLinks) GetSelfOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *TagAssignmentSelfLinks) SetSelf(v string) {
	o.Self = v
}

// GetTag returns the Tag field value
func (o *TagAssignmentSelfLinks) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *TagAssignmentSelfLinks) GetTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *TagAssignmentSelfLinks) SetTag(v string) {
	o.Tag = v
}

// GetResource returns the Resource field value
func (o *TagAssignmentSelfLinks) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *TagAssignmentSelfLinks) GetResourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *TagAssignmentSelfLinks) SetResource(v string) {
	o.Resource = v
}

func (o TagAssignmentSelfLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["self"] = o.Self
	}
	if true {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["_resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableTagAssignmentSelfLinks struct {
	value *TagAssignmentSelfLinks
	isSet bool
}

func (v NullableTagAssignmentSelfLinks) Get() *TagAssignmentSelfLinks {
	return v.value
}

func (v *NullableTagAssignmentSelfLinks) Set(val *TagAssignmentSelfLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableTagAssignmentSelfLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableTagAssignmentSelfLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagAssignmentSelfLinks(val *TagAssignmentSelfLinks) *NullableTagAssignmentSelfLinks {
	return &NullableTagAssignmentSelfLinks{value: val, isSet: true}
}

func (v NullableTagAssignmentSelfLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagAssignmentSelfLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


