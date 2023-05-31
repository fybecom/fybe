# ImageAuditResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The audit entry identifier refers to the unique identifier associated with a particular audit record | 
**Action** | **string** | Category of the action. | 
**Timestamp** | **time.Time** | Point of time when the change occured | 
**TenantId** | **string** | Tenant ID of the customer | 
**CustomerId** | **string** | Identifier of the customer | 
**ChangedBy** | **string** | The identifier of the user who made the modification | 
**Username** | **string** | The name of the user who did the modification. | 
**RequestId** | **string** | The unique identifier of the API request that resulted in the modification. | 
**TraceId** | **string** | The trace ID of the API request that resulted in the modification. | 
**ImageId** | **string** | Image identifier | 
**Changes** | Pointer to **map[string]interface{}** | List of latest changes. | [optional] 

## Methods

### NewImageAuditResponseData

`func NewImageAuditResponseData(id int64, action string, timestamp time.Time, tenantId string, customerId string, changedBy string, username string, requestId string, traceId string, imageId string, ) *ImageAuditResponseData`

NewImageAuditResponseData instantiates a new ImageAuditResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageAuditResponseDataWithDefaults

`func NewImageAuditResponseDataWithDefaults() *ImageAuditResponseData`

NewImageAuditResponseDataWithDefaults instantiates a new ImageAuditResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageAuditResponseData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageAuditResponseData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageAuditResponseData) SetId(v int64)`

SetId sets Id field to given value.


### GetAction

`func (o *ImageAuditResponseData) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ImageAuditResponseData) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ImageAuditResponseData) SetAction(v string)`

SetAction sets Action field to given value.


### GetTimestamp

`func (o *ImageAuditResponseData) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ImageAuditResponseData) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ImageAuditResponseData) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTenantId

`func (o *ImageAuditResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ImageAuditResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ImageAuditResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *ImageAuditResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ImageAuditResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ImageAuditResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetChangedBy

`func (o *ImageAuditResponseData) GetChangedBy() string`

GetChangedBy returns the ChangedBy field if non-nil, zero value otherwise.

### GetChangedByOk

`func (o *ImageAuditResponseData) GetChangedByOk() (*string, bool)`

GetChangedByOk returns a tuple with the ChangedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedBy

`func (o *ImageAuditResponseData) SetChangedBy(v string)`

SetChangedBy sets ChangedBy field to given value.


### GetUsername

`func (o *ImageAuditResponseData) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ImageAuditResponseData) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ImageAuditResponseData) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetRequestId

`func (o *ImageAuditResponseData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ImageAuditResponseData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ImageAuditResponseData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetTraceId

`func (o *ImageAuditResponseData) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ImageAuditResponseData) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ImageAuditResponseData) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetImageId

`func (o *ImageAuditResponseData) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ImageAuditResponseData) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ImageAuditResponseData) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetChanges

`func (o *ImageAuditResponseData) GetChanges() map[string]interface{}`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *ImageAuditResponseData) GetChangesOk() (*map[string]interface{}, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *ImageAuditResponseData) SetChanges(v map[string]interface{})`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *ImageAuditResponseData) HasChanges() bool`

HasChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


