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

// UpdateUserRequest struct for UpdateUserRequest
type UpdateUserRequest struct {
	// The user's name can consist of letters, numbers, colons, dashes, and underscores. It is limited to 255 characters.
	FirstName *string `json:"firstName,omitempty"`
	// The user's last name. The last name may consist of letters, numbers, colons, dashes, and underscores. The maximum length of the last name is 255 characters.
	LastName *string `json:"lastName,omitempty"`
	// The user's email address to which activation and forgot password links will be sent. The email field has a limit of 255 characters.
	Email *string `json:"email,omitempty"`
	// If the user is disabled, they won't be able to log in and access the services anymore.
	Enabled *bool `json:"enabled,omitempty"`
	// Toggle the option to use time-based one-time passwords (TOTP) for two-factor authentication (2FA).
	Totp *bool `json:"totp,omitempty"`
	// The user's locale indicates their preferred language and region settings. Accepted values include en-US, en.
	Locale *string `json:"locale,omitempty"`
	// The user's roles are represented as a list of roleIds.
	Roles *[]int64 `json:"roles,omitempty"`
}

// NewUpdateUserRequest instantiates a new UpdateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserRequest() *UpdateUserRequest {
	this := UpdateUserRequest{}
	return &this
}

// NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserRequestWithDefaults() *UpdateUserRequest {
	this := UpdateUserRequest{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UpdateUserRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UpdateUserRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateUserRequest) SetEmail(v string) {
	o.Email = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateUserRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTotp returns the Totp field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetTotp() bool {
	if o == nil || o.Totp == nil {
		var ret bool
		return ret
	}
	return *o.Totp
}

// GetTotpOk returns a tuple with the Totp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetTotpOk() (*bool, bool) {
	if o == nil || o.Totp == nil {
		return nil, false
	}
	return o.Totp, true
}

// HasTotp returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasTotp() bool {
	if o != nil && o.Totp != nil {
		return true
	}

	return false
}

// SetTotp gets a reference to the given bool and assigns it to the Totp field.
func (o *UpdateUserRequest) SetTotp(v bool) {
	o.Totp = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *UpdateUserRequest) SetLocale(v string) {
	o.Locale = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetRoles() []int64 {
	if o == nil || o.Roles == nil {
		var ret []int64
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetRolesOk() (*[]int64, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []int64 and assigns it to the Roles field.
func (o *UpdateUserRequest) SetRoles(v []int64) {
	o.Roles = &v
}

func (o UpdateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Totp != nil {
		toSerialize["totp"] = o.Totp
	}
	if o.Locale != nil {
		toSerialize["locale"] = o.Locale
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateUserRequest struct {
	value *UpdateUserRequest
	isSet bool
}

func (v NullableUpdateUserRequest) Get() *UpdateUserRequest {
	return v.value
}

func (v *NullableUpdateUserRequest) Set(val *UpdateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserRequest(val *UpdateUserRequest) *NullableUpdateUserRequest {
	return &NullableUpdateUserRequest{value: val, isSet: true}
}

func (v NullableUpdateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


