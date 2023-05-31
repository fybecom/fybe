# TagResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagId** | **float32** | Tag identifier | 
**TagName** | **string** | The name of tag | 

## Methods

### NewTagResponse

`func NewTagResponse(tagId float32, tagName string, ) *TagResponse`

NewTagResponse instantiates a new TagResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagResponseWithDefaults

`func NewTagResponseWithDefaults() *TagResponse`

NewTagResponseWithDefaults instantiates a new TagResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagId

`func (o *TagResponse) GetTagId() float32`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *TagResponse) GetTagIdOk() (*float32, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *TagResponse) SetTagId(v float32)`

SetTagId sets TagId field to given value.


### GetTagName

`func (o *TagResponse) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *TagResponse) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *TagResponse) SetTagName(v string)`

SetTagName sets TagName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


