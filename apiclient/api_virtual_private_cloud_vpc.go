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

// VirtualPrivateCloudVPCApiService VirtualPrivateCloudVPCApi service
type VirtualPrivateCloudVPCApiService service

type ApiAssignInstancePrivateNetworkRequest struct {
	ctx _context.Context
	ApiService *VirtualPrivateCloudVPCApiService
	xRequestId *string
	privateNetworkId int64
	instanceId int64
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiAssignInstancePrivateNetworkRequest) XRequestId(xRequestId string) ApiAssignInstancePrivateNetworkRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiAssignInstancePrivateNetworkRequest) XTraceId(xTraceId string) ApiAssignInstancePrivateNetworkRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiAssignInstancePrivateNetworkRequest) Execute() (AssignInstancePrivateNetworkResponse, *_nethttp.Response, error) {
	return r.ApiService.AssignInstancePrivateNetworkExecute(r)
}

/*
AssignInstancePrivateNetwork Add instance to a Virtual Private Cloud

Add a particular instance to a Virtual Private Cloud

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param privateNetworkId Virtual Private Cloud Identifier
 @param instanceId Compute Instance identifier
 @return ApiAssignInstancePrivateNetworkRequest
*/
func (a *VirtualPrivateCloudVPCApiService) AssignInstancePrivateNetwork(ctx _context.Context, privateNetworkId int64, instanceId int64) ApiAssignInstancePrivateNetworkRequest {
	return ApiAssignInstancePrivateNetworkRequest{
		ApiService: a,
		ctx: ctx,
		privateNetworkId: privateNetworkId,
		instanceId: instanceId,
	}
}

// Execute executes the request
//  @return AssignInstancePrivateNetworkResponse
func (a *VirtualPrivateCloudVPCApiService) AssignInstancePrivateNetworkExecute(r ApiAssignInstancePrivateNetworkRequest) (AssignInstancePrivateNetworkResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AssignInstancePrivateNetworkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VirtualPrivateCloudVPCApiService.AssignInstancePrivateNetwork")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/private-networks/{privateNetworkId}/instances/{instanceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"privateNetworkId"+"}", _neturl.PathEscape(parameterToString(r.privateNetworkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", _neturl.PathEscape(parameterToString(r.instanceId, "")), -1)

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

type ApiCreatePrivateNetworkRequest struct {
	ctx _context.Context
	ApiService *VirtualPrivateCloudVPCApiService
	xRequestId *string
	createPrivateNetworkRequest *CreatePrivateNetworkRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiCreatePrivateNetworkRequest) XRequestId(xRequestId string) ApiCreatePrivateNetworkRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiCreatePrivateNetworkRequest) CreatePrivateNetworkRequest(createPrivateNetworkRequest CreatePrivateNetworkRequest) ApiCreatePrivateNetworkRequest {
	r.createPrivateNetworkRequest = &createPrivateNetworkRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiCreatePrivateNetworkRequest) XTraceId(xTraceId string) ApiCreatePrivateNetworkRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiCreatePrivateNetworkRequest) Execute() (CreatePrivateNetworkResponse, *_nethttp.Response, error) {
	return r.ApiService.CreatePrivateNetworkExecute(r)
}

/*
CreatePrivateNetwork Create a new Virtual Private Cloud

Create a new Virtual Private Cloud. 
 The Virtual Private Cloud API enables you to establish and oversee Virtual Private Clouds (VPC) for your Cloud instances. Through VPCs, instances can communicate directly with each other using private IPs, without traffic leaving the Data Center. This function is available as a paid add-on for each instance that belongs to a Virtual Private Cloud and allows instances to be part of multiple Virtual Private Clouds

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreatePrivateNetworkRequest
*/
func (a *VirtualPrivateCloudVPCApiService) CreatePrivateNetwork(ctx _context.Context) ApiCreatePrivateNetworkRequest {
	return ApiCreatePrivateNetworkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreatePrivateNetworkResponse
func (a *VirtualPrivateCloudVPCApiService) CreatePrivateNetworkExecute(r ApiCreatePrivateNetworkRequest) (CreatePrivateNetworkResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreatePrivateNetworkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VirtualPrivateCloudVPCApiService.CreatePrivateNetwork")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/private-networks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}
	if r.createPrivateNetworkRequest == nil {
		return localVarReturnValue, nil, reportError("createPrivateNetworkRequest is required and must be specified")
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
	localVarPostBody = r.createPrivateNetworkRequest
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

type ApiDeletePrivateNetworkRequest struct {
	ctx _context.Context
	ApiService *VirtualPrivateCloudVPCApiService
	xRequestId *string
	privateNetworkId int64
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiDeletePrivateNetworkRequest) XRequestId(xRequestId string) ApiDeletePrivateNetworkRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiDeletePrivateNetworkRequest) XTraceId(xTraceId string) ApiDeletePrivateNetworkRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiDeletePrivateNetworkRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeletePrivateNetworkExecute(r)
}

/*
DeletePrivateNetwork Delete the virtual private cloud with the given id

Delete the Virtual Private Cloud with the provided id and automatically remove all instances assigned to it

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param privateNetworkId Virtual Private Cloud ID
 @return ApiDeletePrivateNetworkRequest
*/
func (a *VirtualPrivateCloudVPCApiService) DeletePrivateNetwork(ctx _context.Context, privateNetworkId int64) ApiDeletePrivateNetworkRequest {
	return ApiDeletePrivateNetworkRequest{
		ApiService: a,
		ctx: ctx,
		privateNetworkId: privateNetworkId,
	}
}

// Execute executes the request
func (a *VirtualPrivateCloudVPCApiService) DeletePrivateNetworkExecute(r ApiDeletePrivateNetworkRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VirtualPrivateCloudVPCApiService.DeletePrivateNetwork")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/private-networks/{privateNetworkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"privateNetworkId"+"}", _neturl.PathEscape(parameterToString(r.privateNetworkId, "")), -1)

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

type ApiPatchPrivateNetworkRequest struct {
	ctx _context.Context
	ApiService *VirtualPrivateCloudVPCApiService
	xRequestId *string
	privateNetworkId int64
	patchPrivateNetworkRequest *PatchPrivateNetworkRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiPatchPrivateNetworkRequest) XRequestId(xRequestId string) ApiPatchPrivateNetworkRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiPatchPrivateNetworkRequest) PatchPrivateNetworkRequest(patchPrivateNetworkRequest PatchPrivateNetworkRequest) ApiPatchPrivateNetworkRequest {
	r.patchPrivateNetworkRequest = &patchPrivateNetworkRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiPatchPrivateNetworkRequest) XTraceId(xTraceId string) ApiPatchPrivateNetworkRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiPatchPrivateNetworkRequest) Execute() (PatchPrivateNetworkResponse, *_nethttp.Response, error) {
	return r.ApiService.PatchPrivateNetworkExecute(r)
}

/*
PatchPrivateNetwork Update a Virtual Private Cloud by its id

Update your Virtual Private Cloud by id.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param privateNetworkId Virtual Private Cloud ID
 @return ApiPatchPrivateNetworkRequest
*/
func (a *VirtualPrivateCloudVPCApiService) PatchPrivateNetwork(ctx _context.Context, privateNetworkId int64) ApiPatchPrivateNetworkRequest {
	return ApiPatchPrivateNetworkRequest{
		ApiService: a,
		ctx: ctx,
		privateNetworkId: privateNetworkId,
	}
}

// Execute executes the request
//  @return PatchPrivateNetworkResponse
func (a *VirtualPrivateCloudVPCApiService) PatchPrivateNetworkExecute(r ApiPatchPrivateNetworkRequest) (PatchPrivateNetworkResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PatchPrivateNetworkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VirtualPrivateCloudVPCApiService.PatchPrivateNetwork")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/private-networks/{privateNetworkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"privateNetworkId"+"}", _neturl.PathEscape(parameterToString(r.privateNetworkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}
	if r.patchPrivateNetworkRequest == nil {
		return localVarReturnValue, nil, reportError("patchPrivateNetworkRequest is required and must be specified")
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
	localVarPostBody = r.patchPrivateNetworkRequest
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

type ApiRetrievePrivateNetworkRequest struct {
	ctx _context.Context
	ApiService *VirtualPrivateCloudVPCApiService
	xRequestId *string
	privateNetworkId int64
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrievePrivateNetworkRequest) XRequestId(xRequestId string) ApiRetrievePrivateNetworkRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrievePrivateNetworkRequest) XTraceId(xTraceId string) ApiRetrievePrivateNetworkRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiRetrievePrivateNetworkRequest) Execute() (FindPrivateNetworkResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrievePrivateNetworkExecute(r)
}

/*
RetrievePrivateNetwork Get specific Virtual Private Cloud by its id

Get the attribute values for a specific Virtual Private Cloud in your account.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param privateNetworkId Virtual Private Cloud ID
 @return ApiRetrievePrivateNetworkRequest
*/
func (a *VirtualPrivateCloudVPCApiService) RetrievePrivateNetwork(ctx _context.Context, privateNetworkId int64) ApiRetrievePrivateNetworkRequest {
	return ApiRetrievePrivateNetworkRequest{
		ApiService: a,
		ctx: ctx,
		privateNetworkId: privateNetworkId,
	}
}

// Execute executes the request
//  @return FindPrivateNetworkResponse
func (a *VirtualPrivateCloudVPCApiService) RetrievePrivateNetworkExecute(r ApiRetrievePrivateNetworkRequest) (FindPrivateNetworkResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FindPrivateNetworkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VirtualPrivateCloudVPCApiService.RetrievePrivateNetwork")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/private-networks/{privateNetworkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"privateNetworkId"+"}", _neturl.PathEscape(parameterToString(r.privateNetworkId, "")), -1)

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

type ApiRetrievePrivateNetworkListRequest struct {
	ctx _context.Context
	ApiService *VirtualPrivateCloudVPCApiService
	xRequestId *string
	xTraceId *string
	page *int64
	size *int64
	orderBy *[]string
	name *string
	instanceIds *string
	region *string
	dataCenter *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrievePrivateNetworkListRequest) XRequestId(xRequestId string) ApiRetrievePrivateNetworkListRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrievePrivateNetworkListRequest) XTraceId(xTraceId string) ApiRetrievePrivateNetworkListRequest {
	r.xTraceId = &xTraceId
	return r
}
// Number of page to be fetched.
func (r ApiRetrievePrivateNetworkListRequest) Page(page int64) ApiRetrievePrivateNetworkListRequest {
	r.page = &page
	return r
}
// Number of elements per page.
func (r ApiRetrievePrivateNetworkListRequest) Size(size int64) ApiRetrievePrivateNetworkListRequest {
	r.size = &size
	return r
}
// Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;.
func (r ApiRetrievePrivateNetworkListRequest) OrderBy(orderBy []string) ApiRetrievePrivateNetworkListRequest {
	r.orderBy = &orderBy
	return r
}
// Title of the Virtual Private Cloud
func (r ApiRetrievePrivateNetworkListRequest) Name(name string) ApiRetrievePrivateNetworkListRequest {
	r.name = &name
	return r
}
// Compute instance identifiers, separated by comma
func (r ApiRetrievePrivateNetworkListRequest) InstanceIds(instanceIds string) ApiRetrievePrivateNetworkListRequest {
	r.instanceIds = &instanceIds
	return r
}
// The region slug indicating the location of your VPC.
func (r ApiRetrievePrivateNetworkListRequest) Region(region string) ApiRetrievePrivateNetworkListRequest {
	r.region = &region
	return r
}
// The physical facility housing your VPC.
func (r ApiRetrievePrivateNetworkListRequest) DataCenter(dataCenter string) ApiRetrievePrivateNetworkListRequest {
	r.dataCenter = &dataCenter
	return r
}

func (r ApiRetrievePrivateNetworkListRequest) Execute() (ListPrivateNetworkResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrievePrivateNetworkListExecute(r)
}

/*
RetrievePrivateNetworkList List all Virtual Private Clouds

Shows all Virtual Private Clouds in your account using filters.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRetrievePrivateNetworkListRequest
*/
func (a *VirtualPrivateCloudVPCApiService) RetrievePrivateNetworkList(ctx _context.Context) ApiRetrievePrivateNetworkListRequest {
	return ApiRetrievePrivateNetworkListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListPrivateNetworkResponse
func (a *VirtualPrivateCloudVPCApiService) RetrievePrivateNetworkListExecute(r ApiRetrievePrivateNetworkListRequest) (ListPrivateNetworkResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListPrivateNetworkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VirtualPrivateCloudVPCApiService.RetrievePrivateNetworkList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/private-networks"

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
	if r.instanceIds != nil {
		localVarQueryParams.Add("instanceIds", parameterToString(*r.instanceIds, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.dataCenter != nil {
		localVarQueryParams.Add("dataCenter", parameterToString(*r.dataCenter, ""))
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

type ApiUnassignInstancePrivateNetworkRequest struct {
	ctx _context.Context
	ApiService *VirtualPrivateCloudVPCApiService
	xRequestId *string
	privateNetworkId int64
	instanceId int64
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiUnassignInstancePrivateNetworkRequest) XRequestId(xRequestId string) ApiUnassignInstancePrivateNetworkRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiUnassignInstancePrivateNetworkRequest) XTraceId(xTraceId string) ApiUnassignInstancePrivateNetworkRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiUnassignInstancePrivateNetworkRequest) Execute() (UnassignInstancePrivateNetworkResponse, *_nethttp.Response, error) {
	return r.ApiService.UnassignInstancePrivateNetworkExecute(r)
}

/*
UnassignInstancePrivateNetwork Remove particular instance from a Virtual Private Cloud

Take out a particular instance from a Virtual Private Cloud.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param privateNetworkId Virtual Private Cloud Identifier
 @param instanceId Compute Instance Identifier
 @return ApiUnassignInstancePrivateNetworkRequest
*/
func (a *VirtualPrivateCloudVPCApiService) UnassignInstancePrivateNetwork(ctx _context.Context, privateNetworkId int64, instanceId int64) ApiUnassignInstancePrivateNetworkRequest {
	return ApiUnassignInstancePrivateNetworkRequest{
		ApiService: a,
		ctx: ctx,
		privateNetworkId: privateNetworkId,
		instanceId: instanceId,
	}
}

// Execute executes the request
//  @return UnassignInstancePrivateNetworkResponse
func (a *VirtualPrivateCloudVPCApiService) UnassignInstancePrivateNetworkExecute(r ApiUnassignInstancePrivateNetworkRequest) (UnassignInstancePrivateNetworkResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UnassignInstancePrivateNetworkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VirtualPrivateCloudVPCApiService.UnassignInstancePrivateNetwork")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/private-networks/{privateNetworkId}/instances/{instanceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"privateNetworkId"+"}", _neturl.PathEscape(parameterToString(r.privateNetworkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", _neturl.PathEscape(parameterToString(r.instanceId, "")), -1)

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
