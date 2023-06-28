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

// UserResponse struct for UserResponse
type UserResponse struct {
	// Your tenant id
	TenantId string `json:"tenantId"`
	// Your customer number
	CustomerId string `json:"customerId"`
	// User ID.
	UserId string `json:"userId"`
	// The given name of the user. Names may contain letters, numbers, colons, dashes, and underscores. There is a limit of 255 characters per user.
	FirstName string `json:"firstName"`
	// The user's last name. The name may consist of letters, numbers, colons, dashes, and underscores, and is limited to 255 characters.
	LastName string `json:"lastName"`
	// The user's email address to which activation and forgot password links will be sent. The email field has a limit of 255 characters.
	Email string `json:"email"`
	// Verification status of email.
	EmailVerified bool `json:"emailVerified"`
	// If the user is disabled, they will not be able to log in or access services.
	Enabled bool `json:"enabled"`
	// Activate or deactivate two-factor authentication (2FA) using time-based OTP.
	Totp bool `json:"totp"`
	// Activate or deactivate two-factor authentication (2FA) using time-based OTP.
	Locale string `json:"locale"`
	// This field contains the list of roleIds assigned to the user, representing their roles.
	Roles []RoleResponse `json:"roles"`
	// If the user is an owner, they will have permissions to all API endpoints and resources. Enabling this will override all role definitions and `accessAllResources`.
	Owner bool `json:"owner"`
}

// NewUserResponse instantiates a new UserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponse(tenantId string, customerId string, userId string, firstName string, lastName string, email string, emailVerified bool, enabled bool, totp bool, locale string, roles []RoleResponse, owner bool) *UserResponse {
	this := UserResponse{}
	this.TenantId = tenantId
	this.CustomerId = customerId
	this.UserId = userId
	this.FirstName = firstName
	this.LastName = lastName
	this.Email = email
	this.EmailVerified = emailVerified
	this.Enabled = enabled
	this.Totp = totp
	this.Locale = locale
	this.Roles = roles
	this.Owner = owner
	return &this
}

// NewUserResponseWithDefaults instantiates a new UserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseWithDefaults() *UserResponse {
	this := UserResponse{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *UserResponse) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *UserResponse) SetTenantId(v string) {
	o.TenantId = v
}

// GetCustomerId returns the CustomerId field value
func (o *UserResponse) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *UserResponse) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetUserId returns the UserId field value
func (o *UserResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserResponse) SetUserId(v string) {
	o.UserId = v
}

// GetFirstName returns the FirstName field value
func (o *UserResponse) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *UserResponse) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *UserResponse) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UserResponse) SetLastName(v string) {
	o.LastName = v
}

// GetEmail returns the Email field value
func (o *UserResponse) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserResponse) SetEmail(v string) {
	o.Email = v
}

// GetEmailVerified returns the EmailVerified field value
func (o *UserResponse) GetEmailVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmailVerified, true
}

// SetEmailVerified sets field value
func (o *UserResponse) SetEmailVerified(v bool) {
	o.EmailVerified = v
}

// GetEnabled returns the Enabled field value
func (o *UserResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *UserResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetTotp returns the Totp field value
func (o *UserResponse) GetTotp() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Totp
}

// GetTotpOk returns a tuple with the Totp field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetTotpOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Totp, true
}

// SetTotp sets field value
func (o *UserResponse) SetTotp(v bool) {
	o.Totp = v
}

// GetLocale returns the Locale field value
func (o *UserResponse) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetLocaleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *UserResponse) SetLocale(v string) {
	o.Locale = v
}

// GetRoles returns the Roles field value
func (o *UserResponse) GetRoles() []RoleResponse {
	if o == nil {
		var ret []RoleResponse
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetRolesOk() (*[]RoleResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *UserResponse) SetRoles(v []RoleResponse) {
	o.Roles = v
}

// GetOwner returns the Owner field value
func (o *UserResponse) GetOwner() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetOwnerOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *UserResponse) SetOwner(v bool) {
	o.Owner = v
}

func (o UserResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["customerId"] = o.CustomerId
	}
	if true {
		toSerialize["userId"] = o.UserId
	}
	if true {
		toSerialize["firstName"] = o.FirstName
	}
	if true {
		toSerialize["lastName"] = o.LastName
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["emailVerified"] = o.EmailVerified
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["totp"] = o.Totp
	}
	if true {
		toSerialize["locale"] = o.Locale
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	if true {
		toSerialize["owner"] = o.Owner
	}
	return json.Marshal(toSerialize)
}

type NullableUserResponse struct {
	value *UserResponse
	isSet bool
}

func (v NullableUserResponse) Get() *UserResponse {
	return v.value
}

func (v *NullableUserResponse) Set(val *UserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponse(val *UserResponse) *NullableUserResponse {
	return &NullableUserResponse{value: val, isSet: true}
}

func (v NullableUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


