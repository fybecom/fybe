# CreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | The user&#39;s name can consist of letters, numbers, colons, dashes, and underscores. It is limited to 255 characters. | [optional] 
**LastName** | Pointer to **string** | The user&#39;s last name. The last name may consist of letters, numbers, colons, dashes, and underscores. The maximum length of the last name is 255 characters. | [optional] 
**Email** | **string** | The user&#39;s email address to which activation and forgot password links will be sent. The email field has a limit of 255 characters. | 
**Enabled** | **bool** | If the user is disabled, they won&#39;t be able to log in and access the services anymore. | 
**Totp** | **bool** | Toggle the option to use time-based one-time passwords (TOTP) for two-factor authentication (2FA). | 
**Locale** | **string** | The user&#39;s locale indicates their preferred language and region settings. Accepted values include en-US, en. | 
**Roles** | Pointer to **[]int64** | The user&#39;s roles are represented as a list of roleIds. | [optional] 

## Methods

### NewCreateUserRequest

`func NewCreateUserRequest(email string, enabled bool, totp bool, locale string, ) *CreateUserRequest`

NewCreateUserRequest instantiates a new CreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserRequestWithDefaults

`func NewCreateUserRequestWithDefaults() *CreateUserRequest`

NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *CreateUserRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CreateUserRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CreateUserRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CreateUserRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CreateUserRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CreateUserRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CreateUserRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CreateUserRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *CreateUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEnabled

`func (o *CreateUserRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateUserRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateUserRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetTotp

`func (o *CreateUserRequest) GetTotp() bool`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *CreateUserRequest) GetTotpOk() (*bool, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *CreateUserRequest) SetTotp(v bool)`

SetTotp sets Totp field to given value.


### GetLocale

`func (o *CreateUserRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CreateUserRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CreateUserRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetRoles

`func (o *CreateUserRequest) GetRoles() []int64`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateUserRequest) GetRolesOk() (*[]int64, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateUserRequest) SetRoles(v []int64)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateUserRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


