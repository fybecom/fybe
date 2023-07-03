/*
Fybe API Reference

# Introduction  The Fybe API facilitates resource management through HTTP requests. This documentation comprises a collection of HTTP endpoints that adhere to RESTful principles. Each endpoint is accompanied by descriptions, as well as request and response schemas.  ## OpenAPI specification (OAS)  Fybe's OpenAPI Specification can be [downloaded here](https://api.fybe.com/api-v1.yaml).  ## Getting Started  To utilize the Fybe API, you'll require following credentials that can be obtained from from the [Fybe Cockpit](https://cockpit.fybe.com/account/security): 1. ClientId 2. ClientSecret 3. API User (your email address to login to the [Fybe Cockpit](https://cockpit.fybe.com/account/security)) 4. API Password (this is a new password which you'll set or change in the [Fybe Cockpit](https://cockpit.fybe.com/account/security))  ### How to use the API?  As authentication [Bearer Tokens](https://oauth.net/2/bearer-tokens/) in form of [JWT](https://www.rfc-editor.org/rfc/rfc7519) are used. This approach follows the [OAuth 2.0](https://oauth.net/2/) specification.  #### Retrieve the Bearer Token  ```sh POST /auth/realms/Fybe/protocol/openid-connect/token HTTP/1.1 Host: airlock.fybe.com  grant_type=password &password=xxxxxx &redirect_uri=https://example-app.com/redirect &username=xxxxxx &client_id=xxxxxx &client_secret=xxxxxx ```  The actual values for `client_id`, `client_secret`, `username` and `password` can be retrieved from [Fybe Cockpit](https://cockpit.fybe.com/account/security)  ## Using the Command-Line Interface (CLI)  Fybe provides the CLI called `fybe` which can be downloaded from <https://github.com/fybecom/fybe>. Please follow the instructions in the `README.md` to perform the installation on your OS. `fybe` supports Windows, macOS and Linux operating systems.  Using `fybe` CLI invoking makes invoking the API much more comfortable. E.g. `fybe get instances` for retrieving the list of compute instances.  ## Requests  As stated below, the Fybe API accommodates HTTP requests. However, not all endpoints endorse every method. You can find a list of allowed methods within this documentation.  Method | Description ---    | --- GET    | When you need to obtain information about a resource, you should utilize the GET method. The data will be provided as a JSON object. It's essential to note that GET methods are read-only and don't impact any resources. POST   | To create a new object, send a POST method that encodes all necessary attributes in the request body as JSON. PATCH  | PATCH can be used to partially modify a resource, allowing specific attributes to be changed without updating the complete object. PUT    | If you need to update information about a resource, use the PUT method. PUT will overwrite existing values of the item, without considering the current state. DELETE | To remove a resource from your account, apply the DELETE method. If the resource is not found, the operation will generate a 4xx error along with a relevant message.  ## Responses  Typically, the Fybe API will respond to your requests by returning data in [JSON](https://www.json.org/) format, which can be easily processed using any programming language or tools.  Like for any HTTP requests, you'll receive an HTTP status code, which provides general information about the success or error of the request. The table below outlines some common HTTP status codes.  It's important to note that the description of the endpoints and methods doesn't provide an exhaustive list of all possible status codes. Only specific return codes and their corresponding response data are explicitly outlined.  Response Code | Description --- | --- 200 | The response will contain the information you requested. 201 | Your request has been acknowledged, and the resource has been created. 204 | Your request was successful, and no further information is provided in the response. 400 | Your request was not properly formed. 401 | You didn't provide valid authentication credentials. 402 | The request was declined as it necessitates an additional paid service. 403 | You are prohibited from performing the request. You'll need to pass a bearer token. 404 | No results were found for your request, or the resource you're trying to access does not exist. 409 | There is a conflict with resources, such as a violation of unique data constraints when attempting to create or modify resources. 429 | The rate limit has been reached. Please wait for some time before making further requests. 500 | We couldn't fulfill your request due to server-side issues. If this happens, please try again later or reach out to our support team.  Not every endpoint returns data. For example DELETE requests usually don't return any data. All others do return data. For easy handling the return values consists of metadata denoted with and underscore (\"_\") like `_links` or `_pagination`. The actual data is returned in a field called `data`. For convenience reasons this `data` field is always returned as an array even if it consists of only one single element.  DELETE requests usually don't return any data. Return values consists of metadata starting with an underscore (\"_\"), e.g. `_links` and `_pagination`. The requested data is to be found in the field `data`. The `data` field is always an array.   Please visit [Fybe](https://fybe.com) for more general information. 

API version: 1.0.0
Contact: support@fybe.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// TagAssignmentsApiService TagAssignmentsApi service
type TagAssignmentsApiService service

type ApiCreateAssignmentRequest struct {
	ctx _context.Context
	ApiService *TagAssignmentsApiService
	xRequestId *string
	tagId int64
	resourceType string
	resourceId string
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiCreateAssignmentRequest) XRequestId(xRequestId string) ApiCreateAssignmentRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiCreateAssignmentRequest) XTraceId(xTraceId string) ApiCreateAssignmentRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiCreateAssignmentRequest) Execute() (CreateAssignmentResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateAssignmentExecute(r)
}

/*
CreateAssignment Create a new tag assignment

By creating a new tag assignment, you can label a specified resource with a specific tag. This helps with organization or can be used to limit access to the resource.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tagId Tag ID.
 @param resourceType The resource type identifier specifies whether the resource pertains to an `instance`, `image`, or `object-storage`.
 @param resourceId Resource Identifier
 @return ApiCreateAssignmentRequest
*/
func (a *TagAssignmentsApiService) CreateAssignment(ctx _context.Context, tagId int64, resourceType string, resourceId string) ApiCreateAssignmentRequest {
	return ApiCreateAssignmentRequest{
		ApiService: a,
		ctx: ctx,
		tagId: tagId,
		resourceType: resourceType,
		resourceId: resourceId,
	}
}

// Execute executes the request
//  @return CreateAssignmentResponse
func (a *TagAssignmentsApiService) CreateAssignmentExecute(r ApiCreateAssignmentRequest) (CreateAssignmentResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateAssignmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagAssignmentsApiService.CreateAssignment")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tags/{tagId}/assignments/{resourceType}/{resourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"tagId"+"}", _neturl.PathEscape(parameterToString(r.tagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceType"+"}", _neturl.PathEscape(parameterToString(r.resourceType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceId"+"}", _neturl.PathEscape(parameterToString(r.resourceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	if r.xTraceId != nil {
		localVarHeaderParams["x-trace-id"] = parameterToString(*r.xTraceId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteAssignmentRequest struct {
	ctx _context.Context
	ApiService *TagAssignmentsApiService
	xRequestId *string
	tagId int64
	resourceType string
	resourceId string
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiDeleteAssignmentRequest) XRequestId(xRequestId string) ApiDeleteAssignmentRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiDeleteAssignmentRequest) XTraceId(xTraceId string) ApiDeleteAssignmentRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiDeleteAssignmentRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteAssignmentExecute(r)
}

/*
DeleteAssignment Delete tag assignment

Once the tag assignment is removed from the specified resource, any users affected by access restrictions associated with that tag will no longer have access to the resource.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tagId Tag ID.
 @param resourceType The resource type identifier specifies whether the resource pertains to an `instance`, `image`, or `object-storage`.
 @param resourceId Resource Identifier
 @return ApiDeleteAssignmentRequest
*/
func (a *TagAssignmentsApiService) DeleteAssignment(ctx _context.Context, tagId int64, resourceType string, resourceId string) ApiDeleteAssignmentRequest {
	return ApiDeleteAssignmentRequest{
		ApiService: a,
		ctx: ctx,
		tagId: tagId,
		resourceType: resourceType,
		resourceId: resourceId,
	}
}

// Execute executes the request
func (a *TagAssignmentsApiService) DeleteAssignmentExecute(r ApiDeleteAssignmentRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagAssignmentsApiService.DeleteAssignment")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tags/{tagId}/assignments/{resourceType}/{resourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"tagId"+"}", _neturl.PathEscape(parameterToString(r.tagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceType"+"}", _neturl.PathEscape(parameterToString(r.resourceType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceId"+"}", _neturl.PathEscape(parameterToString(r.resourceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return nil, reportError("xRequestId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	if r.xTraceId != nil {
		localVarHeaderParams["x-trace-id"] = parameterToString(*r.xTraceId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRetrieveAssignmentRequest struct {
	ctx _context.Context
	ApiService *TagAssignmentsApiService
	xRequestId *string
	tagId int64
	resourceType string
	resourceId string
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveAssignmentRequest) XRequestId(xRequestId string) ApiRetrieveAssignmentRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveAssignmentRequest) XTraceId(xTraceId string) ApiRetrieveAssignmentRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiRetrieveAssignmentRequest) Execute() (FindAssignmentResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveAssignmentExecute(r)
}

/*
RetrieveAssignment Get particular assignment for the tag

To get the attributes for a particular tag assignment in your account, you'll need to provide the resource type and resource ID.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tagId Tag ID.
 @param resourceType The resource type identifier specifies whether the resource pertains to an `instance`, `image`, or `object-storage`.
 @param resourceId Resource Identifier
 @return ApiRetrieveAssignmentRequest
*/
func (a *TagAssignmentsApiService) RetrieveAssignment(ctx _context.Context, tagId int64, resourceType string, resourceId string) ApiRetrieveAssignmentRequest {
	return ApiRetrieveAssignmentRequest{
		ApiService: a,
		ctx: ctx,
		tagId: tagId,
		resourceType: resourceType,
		resourceId: resourceId,
	}
}

// Execute executes the request
//  @return FindAssignmentResponse
func (a *TagAssignmentsApiService) RetrieveAssignmentExecute(r ApiRetrieveAssignmentRequest) (FindAssignmentResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FindAssignmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagAssignmentsApiService.RetrieveAssignment")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tags/{tagId}/assignments/{resourceType}/{resourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"tagId"+"}", _neturl.PathEscape(parameterToString(r.tagId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceType"+"}", _neturl.PathEscape(parameterToString(r.resourceType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceId"+"}", _neturl.PathEscape(parameterToString(r.resourceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	if r.xTraceId != nil {
		localVarHeaderParams["x-trace-id"] = parameterToString(*r.xTraceId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRetrieveAssignmentListRequest struct {
	ctx _context.Context
	ApiService *TagAssignmentsApiService
	xRequestId *string
	tagId int64
	xTraceId *string
	page *int64
	size *int64
	orderBy *[]string
	resourceType *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveAssignmentListRequest) XRequestId(xRequestId string) ApiRetrieveAssignmentListRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveAssignmentListRequest) XTraceId(xTraceId string) ApiRetrieveAssignmentListRequest {
	r.xTraceId = &xTraceId
	return r
}
// Number of page to be fetched.
func (r ApiRetrieveAssignmentListRequest) Page(page int64) ApiRetrieveAssignmentListRequest {
	r.page = &page
	return r
}
// Number of elements per page.
func (r ApiRetrieveAssignmentListRequest) Size(size int64) ApiRetrieveAssignmentListRequest {
	r.size = &size
	return r
}
// Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;.
func (r ApiRetrieveAssignmentListRequest) OrderBy(orderBy []string) ApiRetrieveAssignmentListRequest {
	r.orderBy = &orderBy
	return r
}
// The filter is a substring match used to determine whether the assignment resource type matches one of instance, image, or object-storage.
func (r ApiRetrieveAssignmentListRequest) ResourceType(resourceType string) ApiRetrieveAssignmentListRequest {
	r.resourceType = &resourceType
	return r
}

func (r ApiRetrieveAssignmentListRequest) Execute() (ListAssignmentResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveAssignmentListExecute(r)
}

/*
RetrieveAssignmentList List tag assignments

Displays all existing assignments for a tag in your account using filters

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tagId Tag ID.
 @return ApiRetrieveAssignmentListRequest
*/
func (a *TagAssignmentsApiService) RetrieveAssignmentList(ctx _context.Context, tagId int64) ApiRetrieveAssignmentListRequest {
	return ApiRetrieveAssignmentListRequest{
		ApiService: a,
		ctx: ctx,
		tagId: tagId,
	}
}

// Execute executes the request
//  @return ListAssignmentResponse
func (a *TagAssignmentsApiService) RetrieveAssignmentListExecute(r ApiRetrieveAssignmentListRequest) (ListAssignmentResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListAssignmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagAssignmentsApiService.RetrieveAssignmentList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tags/{tagId}/assignments"
	localVarPath = strings.Replace(localVarPath, "{"+"tagId"+"}", _neturl.PathEscape(parameterToString(r.tagId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.size != nil {
		localVarQueryParams.Add("size", parameterToString(*r.size, ""))
	}
	if r.orderBy != nil {
		t := *r.orderBy
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("orderBy", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("orderBy", parameterToString(t, "multi"))
		}
	}
	if r.resourceType != nil {
		localVarQueryParams.Add("resourceType", parameterToString(*r.resourceType, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	if r.xTraceId != nil {
		localVarHeaderParams["x-trace-id"] = parameterToString(*r.xTraceId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
