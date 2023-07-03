# CreateSecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The secret&#39;s name that will serve as the password safeguard. | 
**Value** | **string** | The value of the secret that requires preservation, such as a password, must conform to a specific pattern. The pattern necessitates at least one uppercase and lowercase letter, along with either one number and two special characters &#x60;!@#$^&amp;*?_~&#x60; or a minimum of three numbers and one special character &#x60;!@#$^&amp;*?_~&#x60;. The pattern for the secret is represented by the subsequent regular expression:\&quot;: &#x60;^((?&#x3D;.*?[A-Z]{1,})(?&#x3D;.*?[a-z]{1,}))(((?&#x3D;(?:[^d]*d){1})(?&#x3D;([^^&amp;*?_~]*[!@#$^&amp;*?_~]){2,}))|((?&#x3D;(?:[^d]*d){3})(?&#x3D;.*?[!@#$^&amp;*?_~]+))).{8,}$&#x60; | 
**Type** | **string** | Type of the secret (&#x60;password&#x60; or &#x60;ssh&#x60;). | 

## Methods

### NewCreateSecretRequest

`func NewCreateSecretRequest(name string, value string, type_ string, ) *CreateSecretRequest`

NewCreateSecretRequest instantiates a new CreateSecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecretRequestWithDefaults

`func NewCreateSecretRequestWithDefaults() *CreateSecretRequest`

NewCreateSecretRequestWithDefaults instantiates a new CreateSecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSecretRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSecretRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSecretRequest) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *CreateSecretRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateSecretRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateSecretRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *CreateSecretRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateSecretRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateSecretRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


