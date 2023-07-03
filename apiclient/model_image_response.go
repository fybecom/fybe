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

// ImageResponse struct for ImageResponse
type ImageResponse struct {
	// The unique id for the image
	ImageId string `json:"imageId"`
	// The tenant id
	TenantId string `json:"tenantId"`
	// The customer id
	CustomerId string `json:"customerId"`
	// The name of the image
	Name string `json:"name"`
	// The description for image
	Description string `json:"description"`
	// The URL source for the downloaded/provided image.
	Url string `json:"url"`
	// Size of the image in megabytes (MB)
	SizeMb float32 `json:"sizeMb"`
	// Size of the uploaded image in megabytes (MB)
	UploadedSizeMb float32 `json:"uploadedSizeMb"`
	// Type of operating system (OS)
	OsType string `json:"osType"`
	// The version number that distinguishes the contents of the image? It could be the version of the operating system, for example.
	Version string `json:"version"`
	// Format type of image
	Format string `json:"format"`
	// Current status of the image, for example, whether it is still downloading or not
	Status string `json:"status"`
	// Image download error message
	ErrorMessage string `json:"errorMessage"`
	// Flag for indicating if the image is a standard image (flagged as true) or a custom image (flagged as false)
	StandardImage bool `json:"standardImage"`
	// The date and time when the image was created
	CreationDate time.Time `json:"creationDate"`
	// The date and time when the image was last modified
	LastModifiedDate time.Time `json:"lastModifiedDate"`
}

// NewImageResponse instantiates a new ImageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageResponse(imageId string, tenantId string, customerId string, name string, description string, url string, sizeMb float32, uploadedSizeMb float32, osType string, version string, format string, status string, errorMessage string, standardImage bool, creationDate time.Time, lastModifiedDate time.Time) *ImageResponse {
	this := ImageResponse{}
	this.ImageId = imageId
	this.TenantId = tenantId
	this.CustomerId = customerId
	this.Name = name
	this.Description = description
	this.Url = url
	this.SizeMb = sizeMb
	this.UploadedSizeMb = uploadedSizeMb
	this.OsType = osType
	this.Version = version
	this.Format = format
	this.Status = status
	this.ErrorMessage = errorMessage
	this.StandardImage = standardImage
	this.CreationDate = creationDate
	this.LastModifiedDate = lastModifiedDate
	return &this
}

// NewImageResponseWithDefaults instantiates a new ImageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageResponseWithDefaults() *ImageResponse {
	this := ImageResponse{}
	return &this
}

// GetImageId returns the ImageId field value
func (o *ImageResponse) GetImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetImageIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImageId, true
}

// SetImageId sets field value
func (o *ImageResponse) SetImageId(v string) {
	o.ImageId = v
}

// GetTenantId returns the TenantId field value
func (o *ImageResponse) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ImageResponse) SetTenantId(v string) {
	o.TenantId = v
}

// GetCustomerId returns the CustomerId field value
func (o *ImageResponse) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *ImageResponse) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetName returns the Name field value
func (o *ImageResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ImageResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ImageResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ImageResponse) SetDescription(v string) {
	o.Description = v
}

// GetUrl returns the Url field value
func (o *ImageResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ImageResponse) SetUrl(v string) {
	o.Url = v
}

// GetSizeMb returns the SizeMb field value
func (o *ImageResponse) GetSizeMb() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SizeMb
}

// GetSizeMbOk returns a tuple with the SizeMb field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetSizeMbOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SizeMb, true
}

// SetSizeMb sets field value
func (o *ImageResponse) SetSizeMb(v float32) {
	o.SizeMb = v
}

// GetUploadedSizeMb returns the UploadedSizeMb field value
func (o *ImageResponse) GetUploadedSizeMb() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UploadedSizeMb
}

// GetUploadedSizeMbOk returns a tuple with the UploadedSizeMb field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetUploadedSizeMbOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UploadedSizeMb, true
}

// SetUploadedSizeMb sets field value
func (o *ImageResponse) SetUploadedSizeMb(v float32) {
	o.UploadedSizeMb = v
}

// GetOsType returns the OsType field value
func (o *ImageResponse) GetOsType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetOsTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OsType, true
}

// SetOsType sets field value
func (o *ImageResponse) SetOsType(v string) {
	o.OsType = v
}

// GetVersion returns the Version field value
func (o *ImageResponse) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ImageResponse) SetVersion(v string) {
	o.Version = v
}

// GetFormat returns the Format field value
func (o *ImageResponse) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetFormatOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *ImageResponse) SetFormat(v string) {
	o.Format = v
}

// GetStatus returns the Status field value
func (o *ImageResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ImageResponse) SetStatus(v string) {
	o.Status = v
}

// GetErrorMessage returns the ErrorMessage field value
func (o *ImageResponse) GetErrorMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorMessage, true
}

// SetErrorMessage sets field value
func (o *ImageResponse) SetErrorMessage(v string) {
	o.ErrorMessage = v
}

// GetStandardImage returns the StandardImage field value
func (o *ImageResponse) GetStandardImage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StandardImage
}

// GetStandardImageOk returns a tuple with the StandardImage field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetStandardImageOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StandardImage, true
}

// SetStandardImage sets field value
func (o *ImageResponse) SetStandardImage(v bool) {
	o.StandardImage = v
}

// GetCreationDate returns the CreationDate field value
func (o *ImageResponse) GetCreationDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetCreationDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreationDate, true
}

// SetCreationDate sets field value
func (o *ImageResponse) SetCreationDate(v time.Time) {
	o.CreationDate = v
}

// GetLastModifiedDate returns the LastModifiedDate field value
func (o *ImageResponse) GetLastModifiedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedDate
}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value
// and a boolean to check if the value has been set.
func (o *ImageResponse) GetLastModifiedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastModifiedDate, true
}

// SetLastModifiedDate sets field value
func (o *ImageResponse) SetLastModifiedDate(v time.Time) {
	o.LastModifiedDate = v
}

func (o ImageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["imageId"] = o.ImageId
	}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["customerId"] = o.CustomerId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["sizeMb"] = o.SizeMb
	}
	if true {
		toSerialize["uploadedSizeMb"] = o.UploadedSizeMb
	}
	if true {
		toSerialize["osType"] = o.OsType
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["format"] = o.Format
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if true {
		toSerialize["standardImage"] = o.StandardImage
	}
	if true {
		toSerialize["creationDate"] = o.CreationDate
	}
	if true {
		toSerialize["lastModifiedDate"] = o.LastModifiedDate
	}
	return json.Marshal(toSerialize)
}

type NullableImageResponse struct {
	value *ImageResponse
	isSet bool
}

func (v NullableImageResponse) Get() *ImageResponse {
	return v.value
}

func (v *NullableImageResponse) Set(val *ImageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableImageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableImageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageResponse(val *ImageResponse) *NullableImageResponse {
	return &NullableImageResponse{value: val, isSet: true}
}

func (v NullableImageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


