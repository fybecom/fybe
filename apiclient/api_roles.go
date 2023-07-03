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

// RolesApiService RolesApi service
type RolesApiService service

type ApiCreateRoleRequest struct {
	ctx _context.Context
	ApiService *RolesApiService
	xRequestId *string
	createRoleRequest *CreateRoleRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiCreateRoleRequest) XRequestId(xRequestId string) ApiCreateRoleRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiCreateRoleRequest) CreateRoleRequest(createRoleRequest CreateRoleRequest) ApiCreateRoleRequest {
	r.createRoleRequest = &createRoleRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiCreateRoleRequest) XTraceId(xTraceId string) ApiCreateRoleRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiCreateRoleRequest) Execute() (CreateRoleResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateRoleExecute(r)
}

/*
CreateRole Create a role

To create a new role, you can refer to the GET api-permissions endpoint to obtain a list of available API endpoints (apiName) and their corresponding actions. To specify the resources for the role, you can enter tag IDs, which need to be assigned to a resource using the tag management API for them to take effect.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateRoleRequest
*/
func (a *RolesApiService) CreateRole(ctx _context.Context) ApiCreateRoleRequest {
	return ApiCreateRoleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateRoleResponse
func (a *RolesApiService) CreateRoleExecute(r ApiCreateRoleRequest) (CreateRoleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesApiService.CreateRole")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/roles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}
	if r.createRoleRequest == nil {
		return localVarReturnValue, nil, reportError("createRoleRequest is required and must be specified")
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
	localVarPostBody = r.createRoleRequest
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

type ApiDeleteRoleRequest struct {
	ctx _context.Context
	ApiService *RolesApiService
	xRequestId *string
	roleId int64
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiDeleteRoleRequest) XRequestId(xRequestId string) ApiDeleteRoleRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiDeleteRoleRequest) XTraceId(xTraceId string) ApiDeleteRoleRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiDeleteRoleRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteRoleExecute(r)
}

/*
DeleteRole Delete a role

If a role is assigned to a user, it cannot be deleted. To proceed with the deletion, you must first remove the role from the respective user(s).

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleId Role ID
 @return ApiDeleteRoleRequest
*/
func (a *RolesApiService) DeleteRole(ctx _context.Context, roleId int64) ApiDeleteRoleRequest {
	return ApiDeleteRoleRequest{
		ApiService: a,
		ctx: ctx,
		roleId: roleId,
	}
}

// Execute executes the request
func (a *RolesApiService) DeleteRoleExecute(r ApiDeleteRoleRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesApiService.DeleteRole")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/roles/{roleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", _neturl.PathEscape(parameterToString(r.roleId, "")), -1)

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

type ApiRetrieveApiPermissionsListRequest struct {
	ctx _context.Context
	ApiService *RolesApiService
	xRequestId *string
	xTraceId *string
	page *int64
	size *int64
	orderBy *[]string
	apiName *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveApiPermissionsListRequest) XRequestId(xRequestId string) ApiRetrieveApiPermissionsListRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveApiPermissionsListRequest) XTraceId(xTraceId string) ApiRetrieveApiPermissionsListRequest {
	r.xTraceId = &xTraceId
	return r
}
// Number of page to be fetched.
func (r ApiRetrieveApiPermissionsListRequest) Page(page int64) ApiRetrieveApiPermissionsListRequest {
	r.page = &page
	return r
}
// Number of elements per page.
func (r ApiRetrieveApiPermissionsListRequest) Size(size int64) ApiRetrieveApiPermissionsListRequest {
	r.size = &size
	return r
}
// Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;.
func (r ApiRetrieveApiPermissionsListRequest) OrderBy(orderBy []string) ApiRetrieveApiPermissionsListRequest {
	r.orderBy = &orderBy
	return r
}
// Api endpoint name
func (r ApiRetrieveApiPermissionsListRequest) ApiName(apiName string) ApiRetrieveApiPermissionsListRequest {
	r.apiName = &apiName
	return r
}

func (r ApiRetrieveApiPermissionsListRequest) Execute() (ListApiPermissionResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveApiPermissionsListExecute(r)
}

/*
RetrieveApiPermissionsList API permissions list

Enumerate all the API permissions that are available. This comprehensive list can be used as a point of reference when defining roles. It is important to note that not all actions are supported by each endpoint, as their capabilities may differ.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRetrieveApiPermissionsListRequest
*/
func (a *RolesApiService) RetrieveApiPermissionsList(ctx _context.Context) ApiRetrieveApiPermissionsListRequest {
	return ApiRetrieveApiPermissionsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListApiPermissionResponse
func (a *RolesApiService) RetrieveApiPermissionsListExecute(r ApiRetrieveApiPermissionsListRequest) (ListApiPermissionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListApiPermissionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesApiService.RetrieveApiPermissionsList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/roles/api-permissions"

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
	if r.apiName != nil {
		localVarQueryParams.Add("apiName", parameterToString(*r.apiName, ""))
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

type ApiRetrieveRoleRequest struct {
	ctx _context.Context
	ApiService *RolesApiService
	xRequestId *string
	roleId int64
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveRoleRequest) XRequestId(xRequestId string) ApiRetrieveRoleRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveRoleRequest) XTraceId(xTraceId string) ApiRetrieveRoleRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiRetrieveRoleRequest) Execute() (FindRoleResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveRoleExecute(r)
}

/*
RetrieveRole Get a role

Get a role with its corresponding attributes.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleId Role ID
 @return ApiRetrieveRoleRequest
*/
func (a *RolesApiService) RetrieveRole(ctx _context.Context, roleId int64) ApiRetrieveRoleRequest {
	return ApiRetrieveRoleRequest{
		ApiService: a,
		ctx: ctx,
		roleId: roleId,
	}
}

// Execute executes the request
//  @return FindRoleResponse
func (a *RolesApiService) RetrieveRoleExecute(r ApiRetrieveRoleRequest) (FindRoleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FindRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesApiService.RetrieveRole")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/roles/{roleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", _neturl.PathEscape(parameterToString(r.roleId, "")), -1)

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

type ApiRetrieveRoleListRequest struct {
	ctx _context.Context
	ApiService *RolesApiService
	xRequestId *string
	xTraceId *string
	page *int64
	size *int64
	orderBy *[]string
	name *string
	apiName *string
	tagName *string
	type_ *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveRoleListRequest) XRequestId(xRequestId string) ApiRetrieveRoleListRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveRoleListRequest) XTraceId(xTraceId string) ApiRetrieveRoleListRequest {
	r.xTraceId = &xTraceId
	return r
}
// Number of page to be fetched.
func (r ApiRetrieveRoleListRequest) Page(page int64) ApiRetrieveRoleListRequest {
	r.page = &page
	return r
}
// Number of elements per page.
func (r ApiRetrieveRoleListRequest) Size(size int64) ApiRetrieveRoleListRequest {
	r.size = &size
	return r
}
// Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;.
func (r ApiRetrieveRoleListRequest) OrderBy(orderBy []string) ApiRetrieveRoleListRequest {
	r.orderBy = &orderBy
	return r
}
// Role name
func (r ApiRetrieveRoleListRequest) Name(name string) ApiRetrieveRoleListRequest {
	r.name = &name
	return r
}
// Api endpoint name
func (r ApiRetrieveRoleListRequest) ApiName(apiName string) ApiRetrieveRoleListRequest {
	r.apiName = &apiName
	return r
}
// Tag name
func (r ApiRetrieveRoleListRequest) TagName(tagName string) ApiRetrieveRoleListRequest {
	r.tagName = &tagName
	return r
}
// This indicates the tag type, which can be either default or custom.
func (r ApiRetrieveRoleListRequest) Type_(type_ string) ApiRetrieveRoleListRequest {
	r.type_ = &type_
	return r
}

func (r ApiRetrieveRoleListRequest) Execute() (ListRoleResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveRoleListExecute(r)
}

/*
RetrieveRoleList Roles list

Enumerate and apply filters to your roles, which grant permissions to access API endpoints and computing resources.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRetrieveRoleListRequest
*/
func (a *RolesApiService) RetrieveRoleList(ctx _context.Context) ApiRetrieveRoleListRequest {
	return ApiRetrieveRoleListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListRoleResponse
func (a *RolesApiService) RetrieveRoleListExecute(r ApiRetrieveRoleListRequest) (ListRoleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesApiService.RetrieveRoleList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/roles"

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
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.apiName != nil {
		localVarQueryParams.Add("apiName", parameterToString(*r.apiName, ""))
	}
	if r.tagName != nil {
		localVarQueryParams.Add("tagName", parameterToString(*r.tagName, ""))
	}
	if r.type_ != nil {
		localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
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

type ApiUpdateRoleRequest struct {
	ctx _context.Context
	ApiService *RolesApiService
	xRequestId *string
	roleId int64
	updateRoleRequest *UpdateRoleRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiUpdateRoleRequest) XRequestId(xRequestId string) ApiUpdateRoleRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiUpdateRoleRequest) UpdateRoleRequest(updateRoleRequest UpdateRoleRequest) ApiUpdateRoleRequest {
	r.updateRoleRequest = &updateRoleRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiUpdateRoleRequest) XTraceId(xTraceId string) ApiUpdateRoleRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiUpdateRoleRequest) Execute() (UpdateRoleResponse, *_nethttp.Response, error) {
	return r.ApiService.UpdateRoleExecute(r)
}

/*
UpdateRole Update a role

Modify the attributes of your role. Note that the attributes are optional and will retain their original values if not set.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleId Role ID
 @return ApiUpdateRoleRequest
*/
func (a *RolesApiService) UpdateRole(ctx _context.Context, roleId int64) ApiUpdateRoleRequest {
	return ApiUpdateRoleRequest{
		ApiService: a,
		ctx: ctx,
		roleId: roleId,
	}
}

// Execute executes the request
//  @return UpdateRoleResponse
func (a *RolesApiService) UpdateRoleExecute(r ApiUpdateRoleRequest) (UpdateRoleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UpdateRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RolesApiService.UpdateRole")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/roles/{roleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", _neturl.PathEscape(parameterToString(r.roleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}
	if r.updateRoleRequest == nil {
		return localVarReturnValue, nil, reportError("updateRoleRequest is required and must be specified")
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
	localVarPostBody = r.updateRoleRequest
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
