# ApiPermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiName** | **string** | This refers to the API endpoint. To obtain a list of available API endpoints, please refer to the GET api-permissions endpoint. | 
**Actions** | **[]string** | This specifies the permitted action for the API endpoint. Generally, CREATE corresponds to POST endpoints, READ to GET endpoints, UPDATE to PATCH/PUT endpoints, and DELETE to DELETE endpoints. | 

## Methods

### NewApiPermissionsResponse

`func NewApiPermissionsResponse(apiName string, actions []string, ) *ApiPermissionsResponse`

NewApiPermissionsResponse instantiates a new ApiPermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPermissionsResponseWithDefaults

`func NewApiPermissionsResponseWithDefaults() *ApiPermissionsResponse`

NewApiPermissionsResponseWithDefaults instantiates a new ApiPermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiName

`func (o *ApiPermissionsResponse) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *ApiPermissionsResponse) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *ApiPermissionsResponse) SetApiName(v string)`

SetApiName sets ApiName field to given value.


### GetActions

`func (o *ApiPermissionsResponse) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiPermissionsResponse) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiPermissionsResponse) SetActions(v []string)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


