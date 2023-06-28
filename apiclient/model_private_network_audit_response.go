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

// PrivateNetworkAuditResponse struct for PrivateNetworkAuditResponse
type PrivateNetworkAuditResponse struct {
	// Audit Entry ID.
	Id int64 `json:"id"`
	// Virtual Private Cloud Identifier
	PrivateNetworkId float32 `json:"privateNetworkId"`
	// Action Specification.
	Action string `json:"action"`
	// Timestamp of corresponding change.
	Timestamp time.Time `json:"timestamp"`
	// Tenant ID of the customer
	TenantId string `json:"tenantId"`
	// Number of the customer
	CustomerId string `json:"customerId"`
	// ID of the User
	ChangedBy string `json:"changedBy"`
	// Name of the user that made the change.
	Username string `json:"username"`
	// The ID of the API call request that triggered the change.
	RequestId string `json:"requestId"`
	// The ID of the trace of the API call that triggered the change.
	TraceId string `json:"traceId"`
	// Display of all changes.
	Changes *map[string]interface{} `json:"changes,omitempty"`
}

// NewPrivateNetworkAuditResponse instantiates a new PrivateNetworkAuditResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateNetworkAuditResponse(id int64, privateNetworkId float32, action string, timestamp time.Time, tenantId string, customerId string, changedBy string, username string, requestId string, traceId string) *PrivateNetworkAuditResponse {
	this := PrivateNetworkAuditResponse{}
	this.Id = id
	this.PrivateNetworkId = privateNetworkId
	this.Action = action
	this.Timestamp = timestamp
	this.TenantId = tenantId
	this.CustomerId = customerId
	this.ChangedBy = changedBy
	this.Username = username
	this.RequestId = requestId
	this.TraceId = traceId
	return &this
}

// NewPrivateNetworkAuditResponseWithDefaults instantiates a new PrivateNetworkAuditResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateNetworkAuditResponseWithDefaults() *PrivateNetworkAuditResponse {
	this := PrivateNetworkAuditResponse{}
	return &this
}

// GetId returns the Id field value
func (o *PrivateNetworkAuditResponse) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkAuditResponse) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrivateNetworkAuditResponse) SetId(v int64) {
	o.Id = v
}

// GetPrivateNetworkId returns the PrivateNetworkId field value
func (o *PrivateNetworkAuditResponse) GetPrivateNetworkId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PrivateNetworkId
}

// GetPrivateNetworkIdOk returns a tuple with the PrivateNetworkId field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkAuditResponse) GetPrivateNetworkIdOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PrivateNetworkId, true
}

// SetPrivateNetworkId sets field value
func (o *PrivateNetworkAuditResponse) SetPrivateNetworkId(v float32) {
	o.PrivateNetworkId = v
}

// GetAction returns the Action field value
func (o *PrivateNetworkAuditResponse) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkAuditResponse) GetActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *PrivateNetworkAuditResponse) SetAction(v string) {
	o.Action = v
}

// GetTimestamp returns the Timestamp field value
func (o *PrivateNetworkAuditResponse) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkAuditResponse) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *PrivateNetworkAuditResponse) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetTenantId returns the TenantId field value
func (o *PrivateNetworkAuditResponse) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkAuditResponse) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *PrivateNetworkAuditResponse) SetTenantId(v string) {
	o.TenantId = v
}

// GetCustomerId returns the CustomerId field value
func (o *PrivateNetworkAuditResponse) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkAuditResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *PrivateNetworkAuditResponse) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetChangedBy returns the ChangedBy field value
func (o *PrivateNetworkAuditResponse) GetChangedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChangedBy
}

// GetChangedByOk returns a tuple with the ChangedBy field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkAuditResponse) GetChangedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ChangedBy, true
}

// SetChangedBy sets field value
func (o *PrivateNetworkAuditResponse) SetChangedBy(v string) {
	o.ChangedBy = v
}

// GetUsername returns the Username field value
func (o *PrivateNetworkAuditResponse) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkAuditResponse) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *PrivateNetworkAuditResponse) SetUsername(v string) {
	o.Username = v
}

// GetRequestId returns the RequestId field value
func (o *PrivateNetworkAuditResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkAuditResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PrivateNetworkAuditResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetTraceId returns the TraceId field value
func (o *PrivateNetworkAuditResponse) GetTraceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkAuditResponse) GetTraceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TraceId, true
}

// SetTraceId sets field value
func (o *PrivateNetworkAuditResponse) SetTraceId(v string) {
	o.TraceId = v
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *PrivateNetworkAuditResponse) GetChanges() map[string]interface{} {
	if o == nil || o.Changes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetworkAuditResponse) GetChangesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Changes == nil {
		return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *PrivateNetworkAuditResponse) HasChanges() bool {
	if o != nil && o.Changes != nil {
		return true
	}

	return false
}

// SetChanges gets a reference to the given map[string]interface{} and assigns it to the Changes field.
func (o *PrivateNetworkAuditResponse) SetChanges(v map[string]interface{}) {
	o.Changes = &v
}

func (o PrivateNetworkAuditResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["privateNetworkId"] = o.PrivateNetworkId
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["customerId"] = o.CustomerId
	}
	if true {
		toSerialize["changedBy"] = o.ChangedBy
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["requestId"] = o.RequestId
	}
	if true {
		toSerialize["traceId"] = o.TraceId
	}
	if o.Changes != nil {
		toSerialize["changes"] = o.Changes
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateNetworkAuditResponse struct {
	value *PrivateNetworkAuditResponse
	isSet bool
}

func (v NullablePrivateNetworkAuditResponse) Get() *PrivateNetworkAuditResponse {
	return v.value
}

func (v *NullablePrivateNetworkAuditResponse) Set(val *PrivateNetworkAuditResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateNetworkAuditResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateNetworkAuditResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateNetworkAuditResponse(val *PrivateNetworkAuditResponse) *NullablePrivateNetworkAuditResponse {
	return &NullablePrivateNetworkAuditResponse{value: val, isSet: true}
}

func (v NullablePrivateNetworkAuditResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateNetworkAuditResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


