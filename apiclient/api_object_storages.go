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

// ObjectStoragesApiService ObjectStoragesApi service
type ObjectStoragesApiService service

type ApiCancelObjectStorageRequest struct {
	ctx _context.Context
	ApiService *ObjectStoragesApiService
	xRequestId *string
	objectStorageId string
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiCancelObjectStorageRequest) XRequestId(xRequestId string) ApiCancelObjectStorageRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiCancelObjectStorageRequest) XTraceId(xTraceId string) ApiCancelObjectStorageRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiCancelObjectStorageRequest) Execute() (CancelObjectStorageResponse, *_nethttp.Response, error) {
	return r.ApiService.CancelObjectStorageExecute(r)
}

/*
CancelObjectStorage Cancels the selected object storage at the next possible date

Cancels the selected object storage at the next possible date (depending on contract period).

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectStorageId Object Storage Identifier.
 @return ApiCancelObjectStorageRequest
*/
func (a *ObjectStoragesApiService) CancelObjectStorage(ctx _context.Context, objectStorageId string) ApiCancelObjectStorageRequest {
	return ApiCancelObjectStorageRequest{
		ApiService: a,
		ctx: ctx,
		objectStorageId: objectStorageId,
	}
}

// Execute executes the request
//  @return CancelObjectStorageResponse
func (a *ObjectStoragesApiService) CancelObjectStorageExecute(r ApiCancelObjectStorageRequest) (CancelObjectStorageResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CancelObjectStorageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStoragesApiService.CancelObjectStorage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/object-storages/{objectStorageId}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"objectStorageId"+"}", _neturl.PathEscape(parameterToString(r.objectStorageId, "")), -1)

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

type ApiCreateObjectStorageRequest struct {
	ctx _context.Context
	ApiService *ObjectStoragesApiService
	xRequestId *string
	createObjectStorageRequest *CreateObjectStorageRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiCreateObjectStorageRequest) XRequestId(xRequestId string) ApiCreateObjectStorageRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiCreateObjectStorageRequest) CreateObjectStorageRequest(createObjectStorageRequest CreateObjectStorageRequest) ApiCreateObjectStorageRequest {
	r.createObjectStorageRequest = &createObjectStorageRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiCreateObjectStorageRequest) XTraceId(xTraceId string) ApiCreateObjectStorageRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiCreateObjectStorageRequest) Execute() (CreateObjectStorageResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateObjectStorageExecute(r)
}

/*
CreateObjectStorage Create a new object storage

To add more storage to your account, you can either create or purchase a new object storage. Additionally, you can expand the storage capacity by sending a POST request to /v1/object-storages/{objectStorageId}/resize.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateObjectStorageRequest
*/
func (a *ObjectStoragesApiService) CreateObjectStorage(ctx _context.Context) ApiCreateObjectStorageRequest {
	return ApiCreateObjectStorageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateObjectStorageResponse
func (a *ObjectStoragesApiService) CreateObjectStorageExecute(r ApiCreateObjectStorageRequest) (CreateObjectStorageResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateObjectStorageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStoragesApiService.CreateObjectStorage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/object-storages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}
	if r.createObjectStorageRequest == nil {
		return localVarReturnValue, nil, reportError("createObjectStorageRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.createObjectStorageRequest
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

type ApiRetrieveDataCenterListRequest struct {
	ctx _context.Context
	ApiService *ObjectStoragesApiService
	xRequestId *string
	xTraceId *string
	page *int64
	size *int64
	orderBy *[]string
	slug *string
	name *string
	regionName *string
	regionSlug *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveDataCenterListRequest) XRequestId(xRequestId string) ApiRetrieveDataCenterListRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveDataCenterListRequest) XTraceId(xTraceId string) ApiRetrieveDataCenterListRequest {
	r.xTraceId = &xTraceId
	return r
}
// Number of page to be fetched.
func (r ApiRetrieveDataCenterListRequest) Page(page int64) ApiRetrieveDataCenterListRequest {
	r.page = &page
	return r
}
// Number of elements per page.
func (r ApiRetrieveDataCenterListRequest) Size(size int64) ApiRetrieveDataCenterListRequest {
	r.size = &size
	return r
}
// Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;.
func (r ApiRetrieveDataCenterListRequest) OrderBy(orderBy []string) ApiRetrieveDataCenterListRequest {
	r.orderBy = &orderBy
	return r
}
// Filter that can be used as match for data centers.
func (r ApiRetrieveDataCenterListRequest) Slug(slug string) ApiRetrieveDataCenterListRequest {
	r.slug = &slug
	return r
}
// Available region filter for object storage.
func (r ApiRetrieveDataCenterListRequest) Name(name string) ApiRetrieveDataCenterListRequest {
	r.name = &name
	return r
}
// Available region name filter for object storage.
func (r ApiRetrieveDataCenterListRequest) RegionName(regionName string) ApiRetrieveDataCenterListRequest {
	r.regionName = &regionName
	return r
}
// Available region slug filter for object storage.
func (r ApiRetrieveDataCenterListRequest) RegionSlug(regionSlug string) ApiRetrieveDataCenterListRequest {
	r.regionSlug = &regionSlug
	return r
}

func (r ApiRetrieveDataCenterListRequest) Execute() (ListDataCenterResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveDataCenterListExecute(r)
}

/*
RetrieveDataCenterList List all data centers

Gives a list of all data centers incl. their corresponding region.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRetrieveDataCenterListRequest
*/
func (a *ObjectStoragesApiService) RetrieveDataCenterList(ctx _context.Context) ApiRetrieveDataCenterListRequest {
	return ApiRetrieveDataCenterListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListDataCenterResponse
func (a *ObjectStoragesApiService) RetrieveDataCenterListExecute(r ApiRetrieveDataCenterListRequest) (ListDataCenterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListDataCenterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStoragesApiService.RetrieveDataCenterList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/data-centers"

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
	if r.slug != nil {
		localVarQueryParams.Add("slug", parameterToString(*r.slug, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.regionName != nil {
		localVarQueryParams.Add("regionName", parameterToString(*r.regionName, ""))
	}
	if r.regionSlug != nil {
		localVarQueryParams.Add("regionSlug", parameterToString(*r.regionSlug, ""))
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

type ApiRetrieveObjectStorageRequest struct {
	ctx _context.Context
	ApiService *ObjectStoragesApiService
	xRequestId *string
	objectStorageId string
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveObjectStorageRequest) XRequestId(xRequestId string) ApiRetrieveObjectStorageRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveObjectStorageRequest) XTraceId(xTraceId string) ApiRetrieveObjectStorageRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiRetrieveObjectStorageRequest) Execute() (FindObjectStorageResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveObjectStorageExecute(r)
}

/*
RetrieveObjectStorage Get a specific object storage using its id

Gets data for one of your object storage.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectStorageId Object Storage Identifier.
 @return ApiRetrieveObjectStorageRequest
*/
func (a *ObjectStoragesApiService) RetrieveObjectStorage(ctx _context.Context, objectStorageId string) ApiRetrieveObjectStorageRequest {
	return ApiRetrieveObjectStorageRequest{
		ApiService: a,
		ctx: ctx,
		objectStorageId: objectStorageId,
	}
}

// Execute executes the request
//  @return FindObjectStorageResponse
func (a *ObjectStoragesApiService) RetrieveObjectStorageExecute(r ApiRetrieveObjectStorageRequest) (FindObjectStorageResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FindObjectStorageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStoragesApiService.RetrieveObjectStorage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/object-storages/{objectStorageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectStorageId"+"}", _neturl.PathEscape(parameterToString(r.objectStorageId, "")), -1)

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

type ApiRetrieveObjectStorageListRequest struct {
	ctx _context.Context
	ApiService *ObjectStoragesApiService
	xRequestId *string
	xTraceId *string
	page *int64
	size *int64
	orderBy *[]string
	dataCenterName *string
	s3TenantId *string
	region *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveObjectStorageListRequest) XRequestId(xRequestId string) ApiRetrieveObjectStorageListRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveObjectStorageListRequest) XTraceId(xTraceId string) ApiRetrieveObjectStorageListRequest {
	r.xTraceId = &xTraceId
	return r
}
// Number of page to be fetched.
func (r ApiRetrieveObjectStorageListRequest) Page(page int64) ApiRetrieveObjectStorageListRequest {
	r.page = &page
	return r
}
// Number of elements per page.
func (r ApiRetrieveObjectStorageListRequest) Size(size int64) ApiRetrieveObjectStorageListRequest {
	r.size = &size
	return r
}
// Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;.
func (r ApiRetrieveObjectStorageListRequest) OrderBy(orderBy []string) ApiRetrieveObjectStorageListRequest {
	r.orderBy = &orderBy
	return r
}
// Available filter for selected object storage locations.
func (r ApiRetrieveObjectStorageListRequest) DataCenterName(dataCenterName string) ApiRetrieveObjectStorageListRequest {
	r.dataCenterName = &dataCenterName
	return r
}
// Available filter for object storage S3 tenantId.
func (r ApiRetrieveObjectStorageListRequest) S3TenantId(s3TenantId string) ApiRetrieveObjectStorageListRequest {
	r.s3TenantId = &s3TenantId
	return r
}
// Available region filter for Object Storage (Choice between US-central, EU and SIN)
func (r ApiRetrieveObjectStorageListRequest) Region(region string) ApiRetrieveObjectStorageListRequest {
	r.region = &region
	return r
}

func (r ApiRetrieveObjectStorageListRequest) Execute() (ListObjectStorageResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveObjectStorageListExecute(r)
}

/*
RetrieveObjectStorageList List all your object storages

Lists all your object storages using filters.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRetrieveObjectStorageListRequest
*/
func (a *ObjectStoragesApiService) RetrieveObjectStorageList(ctx _context.Context) ApiRetrieveObjectStorageListRequest {
	return ApiRetrieveObjectStorageListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListObjectStorageResponse
func (a *ObjectStoragesApiService) RetrieveObjectStorageListExecute(r ApiRetrieveObjectStorageListRequest) (ListObjectStorageResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListObjectStorageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStoragesApiService.RetrieveObjectStorageList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/object-storages"

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
	if r.dataCenterName != nil {
		localVarQueryParams.Add("dataCenterName", parameterToString(*r.dataCenterName, ""))
	}
	if r.s3TenantId != nil {
		localVarQueryParams.Add("s3TenantId", parameterToString(*r.s3TenantId, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
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

type ApiRetrieveObjectStoragesStatsRequest struct {
	ctx _context.Context
	ApiService *ObjectStoragesApiService
	xRequestId *string
	objectStorageId string
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveObjectStoragesStatsRequest) XRequestId(xRequestId string) ApiRetrieveObjectStoragesStatsRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveObjectStoragesStatsRequest) XTraceId(xTraceId string) ApiRetrieveObjectStoragesStatsRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiRetrieveObjectStoragesStatsRequest) Execute() (ObjectStoragesStatsResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveObjectStoragesStatsExecute(r)
}

/*
RetrieveObjectStoragesStats Gives statistics of selected object storage

Provides statistics on the selected object storage, including the quantity of uploaded/created objects and the amount of storage space occupied. It is important to note that these statistics are updated on a regular basis and are not in real-time.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectStorageId Object Storage Identifier.
 @return ApiRetrieveObjectStoragesStatsRequest
*/
func (a *ObjectStoragesApiService) RetrieveObjectStoragesStats(ctx _context.Context, objectStorageId string) ApiRetrieveObjectStoragesStatsRequest {
	return ApiRetrieveObjectStoragesStatsRequest{
		ApiService: a,
		ctx: ctx,
		objectStorageId: objectStorageId,
	}
}

// Execute executes the request
//  @return ObjectStoragesStatsResponse
func (a *ObjectStoragesApiService) RetrieveObjectStoragesStatsExecute(r ApiRetrieveObjectStoragesStatsRequest) (ObjectStoragesStatsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ObjectStoragesStatsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStoragesApiService.RetrieveObjectStoragesStats")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/object-storages/{objectStorageId}/stats"
	localVarPath = strings.Replace(localVarPath, "{"+"objectStorageId"+"}", _neturl.PathEscape(parameterToString(r.objectStorageId, "")), -1)

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

type ApiUpdateObjectStorageRequest struct {
	ctx _context.Context
	ApiService *ObjectStoragesApiService
	xRequestId *string
	objectStorageId string
	patchObjectStorageRequest *PatchObjectStorageRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiUpdateObjectStorageRequest) XRequestId(xRequestId string) ApiUpdateObjectStorageRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiUpdateObjectStorageRequest) PatchObjectStorageRequest(patchObjectStorageRequest PatchObjectStorageRequest) ApiUpdateObjectStorageRequest {
	r.patchObjectStorageRequest = &patchObjectStorageRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiUpdateObjectStorageRequest) XTraceId(xTraceId string) ApiUpdateObjectStorageRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiUpdateObjectStorageRequest) Execute() (CancelObjectStorageResponse, *_nethttp.Response, error) {
	return r.ApiService.UpdateObjectStorageExecute(r)
}

/*
UpdateObjectStorage Updates the display name of object storage

Updates the display name of object storage. Note: The display name must be unique.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectStorageId Object Storage Identifier.
 @return ApiUpdateObjectStorageRequest
*/
func (a *ObjectStoragesApiService) UpdateObjectStorage(ctx _context.Context, objectStorageId string) ApiUpdateObjectStorageRequest {
	return ApiUpdateObjectStorageRequest{
		ApiService: a,
		ctx: ctx,
		objectStorageId: objectStorageId,
	}
}

// Execute executes the request
//  @return CancelObjectStorageResponse
func (a *ObjectStoragesApiService) UpdateObjectStorageExecute(r ApiUpdateObjectStorageRequest) (CancelObjectStorageResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CancelObjectStorageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStoragesApiService.UpdateObjectStorage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/object-storages/{objectStorageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectStorageId"+"}", _neturl.PathEscape(parameterToString(r.objectStorageId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}
	if r.patchObjectStorageRequest == nil {
		return localVarReturnValue, nil, reportError("patchObjectStorageRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.patchObjectStorageRequest
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

type ApiUpgradeObjectStorageRequest struct {
	ctx _context.Context
	ApiService *ObjectStoragesApiService
	xRequestId *string
	objectStorageId string
	upgradeObjectStorageRequest *UpgradeObjectStorageRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiUpgradeObjectStorageRequest) XRequestId(xRequestId string) ApiUpgradeObjectStorageRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiUpgradeObjectStorageRequest) UpgradeObjectStorageRequest(upgradeObjectStorageRequest UpgradeObjectStorageRequest) ApiUpgradeObjectStorageRequest {
	r.upgradeObjectStorageRequest = &upgradeObjectStorageRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiUpgradeObjectStorageRequest) XTraceId(xTraceId string) ApiUpgradeObjectStorageRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiUpgradeObjectStorageRequest) Execute() (UpgradeObjectStorageResponse, *_nethttp.Response, error) {
	return r.ApiService.UpgradeObjectStorageExecute(r)
}

/*
UpgradeObjectStorage Upgrade object storage size / update autoscaling settings.

Increase the size of your object storage. You can also modify the autoscaling configuration for your storage. The autoscaling feature enables you to acquire additional storage capacity every month, up to the maximum limit, automatically.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectStorageId Object Storage Identifier.
 @return ApiUpgradeObjectStorageRequest
*/
func (a *ObjectStoragesApiService) UpgradeObjectStorage(ctx _context.Context, objectStorageId string) ApiUpgradeObjectStorageRequest {
	return ApiUpgradeObjectStorageRequest{
		ApiService: a,
		ctx: ctx,
		objectStorageId: objectStorageId,
	}
}

// Execute executes the request
//  @return UpgradeObjectStorageResponse
func (a *ObjectStoragesApiService) UpgradeObjectStorageExecute(r ApiUpgradeObjectStorageRequest) (UpgradeObjectStorageResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UpgradeObjectStorageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStoragesApiService.UpgradeObjectStorage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/object-storages/{objectStorageId}/resize"
	localVarPath = strings.Replace(localVarPath, "{"+"objectStorageId"+"}", _neturl.PathEscape(parameterToString(r.objectStorageId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}
	if r.upgradeObjectStorageRequest == nil {
		return localVarReturnValue, nil, reportError("upgradeObjectStorageRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.upgradeObjectStorageRequest
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
