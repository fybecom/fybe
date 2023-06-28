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

// UsersApiService UsersApi service
type UsersApiService service

type ApiCreateUserRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	xRequestId *string
	createUserRequest *CreateUserRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiCreateUserRequest) XRequestId(xRequestId string) ApiCreateUserRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiCreateUserRequest) CreateUserRequest(createUserRequest CreateUserRequest) ApiCreateUserRequest {
	r.createUserRequest = &createUserRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiCreateUserRequest) XTraceId(xTraceId string) ApiCreateUserRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiCreateUserRequest) Execute() (CreateUserResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateUserExecute(r)
}

/*
CreateUser Create User.

Create a new user by specifying the required attributes, such as name, email, enabled status, Two-factor authentication (2FA) via TOTP, admin status (which grants access to all endpoints and resources), accessAllResources, and roles. It's not possible to set a password or secrets for the user during the creation process, as the user must specify their own secrets for security reasons.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateUserRequest
*/
func (a *UsersApiService) CreateUser(ctx _context.Context) ApiCreateUserRequest {
	return ApiCreateUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateUserResponse
func (a *UsersApiService) CreateUserExecute(r ApiCreateUserRequest) (CreateUserResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.CreateUser")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}
	if r.createUserRequest == nil {
		return localVarReturnValue, nil, reportError("createUserRequest is required and must be specified")
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
	localVarPostBody = r.createUserRequest
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

type ApiDeleteUserRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	xRequestId *string
	userId string
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiDeleteUserRequest) XRequestId(xRequestId string) ApiDeleteUserRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiDeleteUserRequest) XTraceId(xTraceId string) ApiDeleteUserRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiDeleteUserRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteUserExecute(r)
}

/*
DeleteUser Delete User.

Deleting a user will revoke their access to all endpoints and resources permanently. To temporarily suspend a user's access, you can update their enabled attribute.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId User ID.
 @return ApiDeleteUserRequest
*/
func (a *UsersApiService) DeleteUser(ctx _context.Context, userId string) ApiDeleteUserRequest {
	return ApiDeleteUserRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
func (a *UsersApiService) DeleteUserExecute(r ApiDeleteUserRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.DeleteUser")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)

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

type ApiGenerateClientSecretRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	xRequestId *string
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiGenerateClientSecretRequest) XRequestId(xRequestId string) ApiGenerateClientSecretRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiGenerateClientSecretRequest) XTraceId(xTraceId string) ApiGenerateClientSecretRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiGenerateClientSecretRequest) Execute() (GenerateClientSecretResponse, *_nethttp.Response, error) {
	return r.ApiService.GenerateClientSecretExecute(r)
}

/*
GenerateClientSecret Regenerate Client Secret.

Regenerate Client Secret.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGenerateClientSecretRequest
*/
func (a *UsersApiService) GenerateClientSecret(ctx _context.Context) ApiGenerateClientSecretRequest {
	return ApiGenerateClientSecretRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GenerateClientSecretResponse
func (a *UsersApiService) GenerateClientSecretExecute(r ApiGenerateClientSecretRequest) (GenerateClientSecretResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GenerateClientSecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.GenerateClientSecret")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/client/secret"

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

type ApiGetObjectStorageCredentialsRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	xRequestId *string
	userId string
	objectStorageId string
	credentialId int64
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiGetObjectStorageCredentialsRequest) XRequestId(xRequestId string) ApiGetObjectStorageCredentialsRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiGetObjectStorageCredentialsRequest) XTraceId(xTraceId string) ApiGetObjectStorageCredentialsRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiGetObjectStorageCredentialsRequest) Execute() (FindCredentialResponse, *_nethttp.Response, error) {
	return r.ApiService.GetObjectStorageCredentialsExecute(r)
}

/*
GetObjectStorageCredentials Get S3 compatible object storage credentials.

Retrieving the S3 compatible object storage credentials for accessing it through S3 compatible tools, such as the aws cli.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId User ID.
 @param objectStorageId Object Storage ID
 @param credentialId Credential ID.
 @return ApiGetObjectStorageCredentialsRequest
*/
func (a *UsersApiService) GetObjectStorageCredentials(ctx _context.Context, userId string, objectStorageId string, credentialId int64) ApiGetObjectStorageCredentialsRequest {
	return ApiGetObjectStorageCredentialsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		objectStorageId: objectStorageId,
		credentialId: credentialId,
	}
}

// Execute executes the request
//  @return FindCredentialResponse
func (a *UsersApiService) GetObjectStorageCredentialsExecute(r ApiGetObjectStorageCredentialsRequest) (FindCredentialResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FindCredentialResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.GetObjectStorageCredentials")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{userId}/object-storages/{objectStorageId}/credentials/{credentialId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"objectStorageId"+"}", _neturl.PathEscape(parameterToString(r.objectStorageId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"credentialId"+"}", _neturl.PathEscape(parameterToString(r.credentialId, "")), -1)

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

type ApiListObjectStorageCredentialsRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	xRequestId *string
	userId string
	xTraceId *string
	page *int64
	size *int64
	orderBy *[]string
	objectStorageId *string
	regionName *string
	displayName *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiListObjectStorageCredentialsRequest) XRequestId(xRequestId string) ApiListObjectStorageCredentialsRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiListObjectStorageCredentialsRequest) XTraceId(xTraceId string) ApiListObjectStorageCredentialsRequest {
	r.xTraceId = &xTraceId
	return r
}
// Number of page to be fetched.
func (r ApiListObjectStorageCredentialsRequest) Page(page int64) ApiListObjectStorageCredentialsRequest {
	r.page = &page
	return r
}
// Number of elements per page.
func (r ApiListObjectStorageCredentialsRequest) Size(size int64) ApiListObjectStorageCredentialsRequest {
	r.size = &size
	return r
}
// Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;.
func (r ApiListObjectStorageCredentialsRequest) OrderBy(orderBy []string) ApiListObjectStorageCredentialsRequest {
	r.orderBy = &orderBy
	return r
}
// Object Storage ID
func (r ApiListObjectStorageCredentialsRequest) ObjectStorageId(objectStorageId string) ApiListObjectStorageCredentialsRequest {
	r.objectStorageId = &objectStorageId
	return r
}
// You can filter object storage by specific regions. The following regions are available: Asia (Singapore), European Union (Germany), United States (Central)
func (r ApiListObjectStorageCredentialsRequest) RegionName(regionName string) ApiListObjectStorageCredentialsRequest {
	r.regionName = &regionName
	return r
}
// Filter the Objest STorage by the displayName property
func (r ApiListObjectStorageCredentialsRequest) DisplayName(displayName string) ApiListObjectStorageCredentialsRequest {
	r.displayName = &displayName
	return r
}

func (r ApiListObjectStorageCredentialsRequest) Execute() (ListCredentialResponse, *_nethttp.Response, error) {
	return r.ApiService.ListObjectStorageCredentialsExecute(r)
}

/*
ListObjectStorageCredentials Retrieve a list of S3-compatible object storage credentials for a user.

Retrieve a list of S3-compatible object storage credentials that can be used to access it via S3-compatible tools such as the aws CLI.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId User ID.
 @return ApiListObjectStorageCredentialsRequest
*/
func (a *UsersApiService) ListObjectStorageCredentials(ctx _context.Context, userId string) ApiListObjectStorageCredentialsRequest {
	return ApiListObjectStorageCredentialsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return ListCredentialResponse
func (a *UsersApiService) ListObjectStorageCredentialsExecute(r ApiListObjectStorageCredentialsRequest) (ListCredentialResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListCredentialResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.ListObjectStorageCredentials")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{userId}/object-storages/credentials"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)

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
	if r.objectStorageId != nil {
		localVarQueryParams.Add("objectStorageId", parameterToString(*r.objectStorageId, ""))
	}
	if r.regionName != nil {
		localVarQueryParams.Add("regionName", parameterToString(*r.regionName, ""))
	}
	if r.displayName != nil {
		localVarQueryParams.Add("displayName", parameterToString(*r.displayName, ""))
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

type ApiRegenerateObjectStorageCredentialsRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	xRequestId *string
	userId string
	objectStorageId string
	credentialId int64
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRegenerateObjectStorageCredentialsRequest) XRequestId(xRequestId string) ApiRegenerateObjectStorageCredentialsRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRegenerateObjectStorageCredentialsRequest) XTraceId(xTraceId string) ApiRegenerateObjectStorageCredentialsRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiRegenerateObjectStorageCredentialsRequest) Execute() (FindCredentialResponse, *_nethttp.Response, error) {
	return r.ApiService.RegenerateObjectStorageCredentialsExecute(r)
}

/*
RegenerateObjectStorageCredentials Regenerate secret key for Object Storage.

This endpoint regenerates the secret key for a specified user, which can be used to access S3-compatible object storage solutions.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId User ID.
 @param objectStorageId Object Storage ID
 @param credentialId Credential ID.
 @return ApiRegenerateObjectStorageCredentialsRequest
*/
func (a *UsersApiService) RegenerateObjectStorageCredentials(ctx _context.Context, userId string, objectStorageId string, credentialId int64) ApiRegenerateObjectStorageCredentialsRequest {
	return ApiRegenerateObjectStorageCredentialsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		objectStorageId: objectStorageId,
		credentialId: credentialId,
	}
}

// Execute executes the request
//  @return FindCredentialResponse
func (a *UsersApiService) RegenerateObjectStorageCredentialsExecute(r ApiRegenerateObjectStorageCredentialsRequest) (FindCredentialResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FindCredentialResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.RegenerateObjectStorageCredentials")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{userId}/object-storages/{objectStorageId}/credentials/{credentialId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"objectStorageId"+"}", _neturl.PathEscape(parameterToString(r.objectStorageId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"credentialId"+"}", _neturl.PathEscape(parameterToString(r.credentialId, "")), -1)

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

type ApiResendEmailVerificationRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	xRequestId *string
	userId string
	xTraceId *string
	redirectUrl *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiResendEmailVerificationRequest) XRequestId(xRequestId string) ApiResendEmailVerificationRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiResendEmailVerificationRequest) XTraceId(xTraceId string) ApiResendEmailVerificationRequest {
	r.xTraceId = &xTraceId
	return r
}
// The URL where users will be redirected after resetting their password.
func (r ApiResendEmailVerificationRequest) RedirectUrl(redirectUrl string) ApiResendEmailVerificationRequest {
	r.redirectUrl = &redirectUrl
	return r
}

func (r ApiResendEmailVerificationRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ResendEmailVerificationExecute(r)
}

/*
ResendEmailVerification Get verification email

Get verification email

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId User ID.
 @return ApiResendEmailVerificationRequest
*/
func (a *UsersApiService) ResendEmailVerification(ctx _context.Context, userId string) ApiResendEmailVerificationRequest {
	return ApiResendEmailVerificationRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
func (a *UsersApiService) ResendEmailVerificationExecute(r ApiResendEmailVerificationRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.ResendEmailVerification")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{userId}/resend-email-verification"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return nil, reportError("xRequestId is required and must be specified")
	}

	if r.redirectUrl != nil {
		localVarQueryParams.Add("redirectUrl", parameterToString(*r.redirectUrl, ""))
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

type ApiResetPasswordRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	xRequestId *string
	userId string
	xTraceId *string
	redirectUrl *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiResetPasswordRequest) XRequestId(xRequestId string) ApiResetPasswordRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiResetPasswordRequest) XTraceId(xTraceId string) ApiResetPasswordRequest {
	r.xTraceId = &xTraceId
	return r
}
// The URL where users will be redirected after resetting their password.
func (r ApiResetPasswordRequest) RedirectUrl(redirectUrl string) ApiResetPasswordRequest {
	r.redirectUrl = &redirectUrl
	return r
}

func (r ApiResetPasswordRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ResetPasswordExecute(r)
}

/*
ResetPassword Get email to reset password.

Get email to reset password for a specific user

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId User ID.
 @return ApiResetPasswordRequest
*/
func (a *UsersApiService) ResetPassword(ctx _context.Context, userId string) ApiResetPasswordRequest {
	return ApiResetPasswordRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
func (a *UsersApiService) ResetPasswordExecute(r ApiResetPasswordRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.ResetPassword")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{userId}/reset-password"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return nil, reportError("xRequestId is required and must be specified")
	}

	if r.redirectUrl != nil {
		localVarQueryParams.Add("redirectUrl", parameterToString(*r.redirectUrl, ""))
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

type ApiRetrieveUserRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	xRequestId *string
	userId string
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveUserRequest) XRequestId(xRequestId string) ApiRetrieveUserRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveUserRequest) XTraceId(xTraceId string) ApiRetrieveUserRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiRetrieveUserRequest) Execute() (FindUserResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveUserExecute(r)
}

/*
RetrieveUser Get specific User.

Get specific User and attributes.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId User ID.
 @return ApiRetrieveUserRequest
*/
func (a *UsersApiService) RetrieveUser(ctx _context.Context, userId string) ApiRetrieveUserRequest {
	return ApiRetrieveUserRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return FindUserResponse
func (a *UsersApiService) RetrieveUserExecute(r ApiRetrieveUserRequest) (FindUserResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FindUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.RetrieveUser")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)

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

type ApiRetrieveUserClientRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	xRequestId *string
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveUserClientRequest) XRequestId(xRequestId string) ApiRetrieveUserClientRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveUserClientRequest) XTraceId(xTraceId string) ApiRetrieveUserClientRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiRetrieveUserClientRequest) Execute() (FindClientResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveUserClientExecute(r)
}

/*
RetrieveUserClient Retrieve Client.

Retrieve Client.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRetrieveUserClientRequest
*/
func (a *UsersApiService) RetrieveUserClient(ctx _context.Context) ApiRetrieveUserClientRequest {
	return ApiRetrieveUserClientRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FindClientResponse
func (a *UsersApiService) RetrieveUserClientExecute(r ApiRetrieveUserClientRequest) (FindClientResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FindClientResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.RetrieveUserClient")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/client"

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

type ApiRetrieveUserListRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	xRequestId *string
	xTraceId *string
	page *int64
	size *int64
	orderBy *[]string
	email *string
	enabled *bool
	owner *bool
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveUserListRequest) XRequestId(xRequestId string) ApiRetrieveUserListRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveUserListRequest) XTraceId(xTraceId string) ApiRetrieveUserListRequest {
	r.xTraceId = &xTraceId
	return r
}
// Number of page to be fetched.
func (r ApiRetrieveUserListRequest) Page(page int64) ApiRetrieveUserListRequest {
	r.page = &page
	return r
}
// Number of elements per page.
func (r ApiRetrieveUserListRequest) Size(size int64) ApiRetrieveUserListRequest {
	r.size = &size
	return r
}
// Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;.
func (r ApiRetrieveUserListRequest) OrderBy(orderBy []string) ApiRetrieveUserListRequest {
	r.orderBy = &orderBy
	return r
}
// Filter user emails by matching a substring.
func (r ApiRetrieveUserListRequest) Email(email string) ApiRetrieveUserListRequest {
	r.email = &email
	return r
}
// Filter users based on their enabled status.
func (r ApiRetrieveUserListRequest) Enabled(enabled bool) ApiRetrieveUserListRequest {
	r.enabled = &enabled
	return r
}
// Filter users based on whether they are owners or not.
func (r ApiRetrieveUserListRequest) Owner(owner bool) ApiRetrieveUserListRequest {
	r.owner = &owner
	return r
}

func (r ApiRetrieveUserListRequest) Execute() (ListUserResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveUserListExecute(r)
}

/*
RetrieveUserList User List

List and filter users.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRetrieveUserListRequest
*/
func (a *UsersApiService) RetrieveUserList(ctx _context.Context) ApiRetrieveUserListRequest {
	return ApiRetrieveUserListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListUserResponse
func (a *UsersApiService) RetrieveUserListExecute(r ApiRetrieveUserListRequest) (ListUserResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.RetrieveUserList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users"

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
	if r.email != nil {
		localVarQueryParams.Add("email", parameterToString(*r.email, ""))
	}
	if r.enabled != nil {
		localVarQueryParams.Add("enabled", parameterToString(*r.enabled, ""))
	}
	if r.owner != nil {
		localVarQueryParams.Add("owner", parameterToString(*r.owner, ""))
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

type ApiUpdateUserRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	xRequestId *string
	userId string
	updateUserRequest *UpdateUserRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiUpdateUserRequest) XRequestId(xRequestId string) ApiUpdateUserRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiUpdateUserRequest) UpdateUserRequest(updateUserRequest UpdateUserRequest) ApiUpdateUserRequest {
	r.updateUserRequest = &updateUserRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiUpdateUserRequest) XTraceId(xTraceId string) ApiUpdateUserRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiUpdateUserRequest) Execute() (UpdateUserResponse, *_nethttp.Response, error) {
	return r.ApiService.UpdateUserExecute(r)
}

/*
UpdateUser Update User.

Update the specified attributes of a user. Only include the attributes that need to be changed. If an attribute is not included, its original value will be retained.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId User ID.
 @return ApiUpdateUserRequest
*/
func (a *UsersApiService) UpdateUser(ctx _context.Context, userId string) ApiUpdateUserRequest {
	return ApiUpdateUserRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return UpdateUserResponse
func (a *UsersApiService) UpdateUserExecute(r ApiUpdateUserRequest) (UpdateUserResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UpdateUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.UpdateUser")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}
	if r.updateUserRequest == nil {
		return localVarReturnValue, nil, reportError("updateUserRequest is required and must be specified")
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
	localVarPostBody = r.updateUserRequest
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
