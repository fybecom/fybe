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

// ObjectStorageResponse struct for ObjectStorageResponse
type ObjectStorageResponse struct {
	// Tenant ID of your customer.
	TenantId string `json:"tenantId"`
	// Customer number
	CustomerId string `json:"customerId"`
	// Identifier of your object storage.
	ObjectStorageId string `json:"objectStorageId"`
	// Creation date for object storage.
	CreatedDate time.Time `json:"createdDate"`
	// Cancellation date of the object storage.
	CancelDate string `json:"cancelDate"`
	// Settings of Autoscaling.
	AutoScaling AutoScalingTypeResponse `json:"autoScaling"`
	// Location of the Data center of your object storage.
	DataCenter string `json:"dataCenter"`
	// Size of purchased of object storage in TerraByte.
	TotalPurchasedSpaceTB float64 `json:"totalPurchasedSpaceTB"`
	// URL to connect to your S3 compatible Fybe Object Storage
	S3Url string `json:"s3Url"`
	// Your S3 tenantID that is required for public sharing.
	S3TenantId string `json:"s3TenantId"`
	// Status of the object storage.
	Status string `json:"status"`
	// Location of your object storage.
	Region string `json:"region"`
	// Specified Display name for the object storage.
	DisplayName string `json:"displayName"`
}

// NewObjectStorageResponse instantiates a new ObjectStorageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStorageResponse(tenantId string, customerId string, objectStorageId string, createdDate time.Time, cancelDate string, autoScaling AutoScalingTypeResponse, dataCenter string, totalPurchasedSpaceTB float64, s3Url string, s3TenantId string, status string, region string, displayName string) *ObjectStorageResponse {
	this := ObjectStorageResponse{}
	this.TenantId = tenantId
	this.CustomerId = customerId
	this.ObjectStorageId = objectStorageId
	this.CreatedDate = createdDate
	this.CancelDate = cancelDate
	this.AutoScaling = autoScaling
	this.DataCenter = dataCenter
	this.TotalPurchasedSpaceTB = totalPurchasedSpaceTB
	this.S3Url = s3Url
	this.S3TenantId = s3TenantId
	this.Status = status
	this.Region = region
	this.DisplayName = displayName
	return &this
}

// NewObjectStorageResponseWithDefaults instantiates a new ObjectStorageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStorageResponseWithDefaults() *ObjectStorageResponse {
	this := ObjectStorageResponse{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *ObjectStorageResponse) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponse) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ObjectStorageResponse) SetTenantId(v string) {
	o.TenantId = v
}

// GetCustomerId returns the CustomerId field value
func (o *ObjectStorageResponse) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *ObjectStorageResponse) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetObjectStorageId returns the ObjectStorageId field value
func (o *ObjectStorageResponse) GetObjectStorageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectStorageId
}

// GetObjectStorageIdOk returns a tuple with the ObjectStorageId field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponse) GetObjectStorageIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectStorageId, true
}

// SetObjectStorageId sets field value
func (o *ObjectStorageResponse) SetObjectStorageId(v string) {
	o.ObjectStorageId = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *ObjectStorageResponse) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponse) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *ObjectStorageResponse) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetCancelDate returns the CancelDate field value
func (o *ObjectStorageResponse) GetCancelDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CancelDate
}

// GetCancelDateOk returns a tuple with the CancelDate field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponse) GetCancelDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CancelDate, true
}

// SetCancelDate sets field value
func (o *ObjectStorageResponse) SetCancelDate(v string) {
	o.CancelDate = v
}

// GetAutoScaling returns the AutoScaling field value
func (o *ObjectStorageResponse) GetAutoScaling() AutoScalingTypeResponse {
	if o == nil {
		var ret AutoScalingTypeResponse
		return ret
	}

	return o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponse) GetAutoScalingOk() (*AutoScalingTypeResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AutoScaling, true
}

// SetAutoScaling sets field value
func (o *ObjectStorageResponse) SetAutoScaling(v AutoScalingTypeResponse) {
	o.AutoScaling = v
}

// GetDataCenter returns the DataCenter field value
func (o *ObjectStorageResponse) GetDataCenter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataCenter
}

// GetDataCenterOk returns a tuple with the DataCenter field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponse) GetDataCenterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DataCenter, true
}

// SetDataCenter sets field value
func (o *ObjectStorageResponse) SetDataCenter(v string) {
	o.DataCenter = v
}

// GetTotalPurchasedSpaceTB returns the TotalPurchasedSpaceTB field value
func (o *ObjectStorageResponse) GetTotalPurchasedSpaceTB() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.TotalPurchasedSpaceTB
}

// GetTotalPurchasedSpaceTBOk returns a tuple with the TotalPurchasedSpaceTB field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponse) GetTotalPurchasedSpaceTBOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalPurchasedSpaceTB, true
}

// SetTotalPurchasedSpaceTB sets field value
func (o *ObjectStorageResponse) SetTotalPurchasedSpaceTB(v float64) {
	o.TotalPurchasedSpaceTB = v
}

// GetS3Url returns the S3Url field value
func (o *ObjectStorageResponse) GetS3Url() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.S3Url
}

// GetS3UrlOk returns a tuple with the S3Url field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponse) GetS3UrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.S3Url, true
}

// SetS3Url sets field value
func (o *ObjectStorageResponse) SetS3Url(v string) {
	o.S3Url = v
}

// GetS3TenantId returns the S3TenantId field value
func (o *ObjectStorageResponse) GetS3TenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.S3TenantId
}

// GetS3TenantIdOk returns a tuple with the S3TenantId field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponse) GetS3TenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.S3TenantId, true
}

// SetS3TenantId sets field value
func (o *ObjectStorageResponse) SetS3TenantId(v string) {
	o.S3TenantId = v
}

// GetStatus returns the Status field value
func (o *ObjectStorageResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponse) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ObjectStorageResponse) SetStatus(v string) {
	o.Status = v
}

// GetRegion returns the Region field value
func (o *ObjectStorageResponse) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponse) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *ObjectStorageResponse) SetRegion(v string) {
	o.Region = v
}

// GetDisplayName returns the DisplayName field value
func (o *ObjectStorageResponse) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ObjectStorageResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ObjectStorageResponse) SetDisplayName(v string) {
	o.DisplayName = v
}

func (o ObjectStorageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["customerId"] = o.CustomerId
	}
	if true {
		toSerialize["objectStorageId"] = o.ObjectStorageId
	}
	if true {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if true {
		toSerialize["cancelDate"] = o.CancelDate
	}
	if true {
		toSerialize["autoScaling"] = o.AutoScaling
	}
	if true {
		toSerialize["dataCenter"] = o.DataCenter
	}
	if true {
		toSerialize["totalPurchasedSpaceTB"] = o.TotalPurchasedSpaceTB
	}
	if true {
		toSerialize["s3Url"] = o.S3Url
	}
	if true {
		toSerialize["s3TenantId"] = o.S3TenantId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	return json.Marshal(toSerialize)
}

type NullableObjectStorageResponse struct {
	value *ObjectStorageResponse
	isSet bool
}

func (v NullableObjectStorageResponse) Get() *ObjectStorageResponse {
	return v.value
}

func (v *NullableObjectStorageResponse) Set(val *ObjectStorageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectStorageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectStorageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectStorageResponse(val *ObjectStorageResponse) *NullableObjectStorageResponse {
	return &NullableObjectStorageResponse{value: val, isSet: true}
}

func (v NullableObjectStorageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectStorageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


