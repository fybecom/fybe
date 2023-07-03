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

// InstancesApiService InstancesApi service
type InstancesApiService service

type ApiCancelInstanceRequest struct {
	ctx _context.Context
	ApiService *InstancesApiService
	xRequestId *string
	instanceId int64
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiCancelInstanceRequest) XRequestId(xRequestId string) ApiCancelInstanceRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiCancelInstanceRequest) XTraceId(xTraceId string) ApiCancelInstanceRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiCancelInstanceRequest) Execute() (CancelInstanceResponse, *_nethttp.Response, error) {
	return r.ApiService.CancelInstanceExecute(r)
}

/*
CancelInstance Cancel specific instance by id

You are at liberty to terminate a previously established instance at any point.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param instanceId The ID of the instance
 @return ApiCancelInstanceRequest
*/
func (a *InstancesApiService) CancelInstance(ctx _context.Context, instanceId int64) ApiCancelInstanceRequest {
	return ApiCancelInstanceRequest{
		ApiService: a,
		ctx: ctx,
		instanceId: instanceId,
	}
}

// Execute executes the request
//  @return CancelInstanceResponse
func (a *InstancesApiService) CancelInstanceExecute(r ApiCancelInstanceRequest) (CancelInstanceResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CancelInstanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstancesApiService.CancelInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/compute/instances/{instanceId}/cancel"
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

type ApiCreateInstanceRequest struct {
	ctx _context.Context
	ApiService *InstancesApiService
	xRequestId *string
	createInstanceRequest *CreateInstanceRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiCreateInstanceRequest) XRequestId(xRequestId string) ApiCreateInstanceRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiCreateInstanceRequest) CreateInstanceRequest(createInstanceRequest CreateInstanceRequest) ApiCreateInstanceRequest {
	r.createInstanceRequest = &createInstanceRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiCreateInstanceRequest) XTraceId(xTraceId string) ApiCreateInstanceRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiCreateInstanceRequest) Execute() (CreateInstanceResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateInstanceExecute(r)
}

/*
CreateInstance Create a new compute instance

Create a compute instance for your account using the specified parameters.         <table>           <tr><th>Product ID</th><th>Product Name</th></tr>           <tr><td>V1</td><td>S-0.5 NVMe</td></tr>           <tr><td>V2</td><td>S-1 NVMe</td></tr>           <tr><td>V3</td><td>S-2 NVMe</td></tr>           <tr><td>V4</td><td>S-4 NVMe</td></tr>           <tr><td>V5</td><td>S-8 NVMe</td></tr>           <tr><td>V6</td><td>S-16 NVMe</td></tr>           <tr><td>V7</td><td>S-30 NVMe</td></tr>           <tr><td>V9</td><td>S-0.5 SSD</td></tr>           <tr><td>V10</td><td>S-1 SSD</td></tr>           <tr><td>V11</td><td>S-2 SSD</td></tr>           <tr><td>V12</td><td>S-4 SSD</td></tr>           <tr><td>V13</td><td>S-8 SSD</td></tr>           <tr><td>V14</td><td>S-16 SSD</td></tr>           <tr><td>V15</td><td>S-30 SSD</td></tr>           <tr><td>V16</td><td>GP-4</td></tr>           <tr><td>V17</td><td>GP-8</td></tr>           <tr><td>V18</td><td>GP-16</td></tr>           <tr><td>V19</td><td>GP-32</td></tr>           <tr><td>V20</td><td>GP-64</td></tr>           <tr><td>V21</td><td>CO-2</td></tr>           <tr><td>V22</td><td>CO-4</td></tr>           <tr><td>V23</td><td>CO-8</td></tr>           <tr><td>V24</td><td>CO-16</td></tr>           <tr><td>V25</td><td>CO-32</td></tr>           <tr><td>V26</td><td>CO-48</td></tr>           <tr><td>V27</td><td>MO-8</td></tr>           <tr><td>V28</td><td>MO-16</td></tr>           <tr><td>V29</td><td>MO-32</td></tr>           <tr><td>V30</td><td>MO-64</td></tr>           <tr><td>V31</td><td>MO-128</td></tr>           <tr><td>V32</td><td>SO-8 NVMe</td></tr>           <tr><td>V33</td><td>SO-16 NVMe</td></tr>           <tr><td>V34</td><td>SO-32 NVMe</td></tr>           <tr><td>V35</td><td>SO-64 NVMe</td></tr>           <tr><td>V36</td><td>SO-128 NVMe</td></tr>           <tr><td>V37</td><td>SO-8 SSD</td></tr>           <tr><td>V38</td><td>SO-16 SSD</td></tr>           <tr><td>V39</td><td>SO-32 SSD</td></tr>           <tr><td>V40</td><td>SO-64 SSD</td></tr>           <tr><td>V41</td><td>SO-128 SSD</td></tr>           </table>

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateInstanceRequest
*/
func (a *InstancesApiService) CreateInstance(ctx _context.Context) ApiCreateInstanceRequest {
	return ApiCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateInstanceResponse
func (a *InstancesApiService) CreateInstanceExecute(r ApiCreateInstanceRequest) (CreateInstanceResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateInstanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstancesApiService.CreateInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/compute/instances"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}
	if r.createInstanceRequest == nil {
		return localVarReturnValue, nil, reportError("createInstanceRequest is required and must be specified")
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
	localVarPostBody = r.createInstanceRequest
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

type ApiPatchInstanceRequest struct {
	ctx _context.Context
	ApiService *InstancesApiService
	xRequestId *string
	instanceId int64
	patchInstanceRequest *PatchInstanceRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiPatchInstanceRequest) XRequestId(xRequestId string) ApiPatchInstanceRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiPatchInstanceRequest) PatchInstanceRequest(patchInstanceRequest PatchInstanceRequest) ApiPatchInstanceRequest {
	r.patchInstanceRequest = &patchInstanceRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiPatchInstanceRequest) XTraceId(xTraceId string) ApiPatchInstanceRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiPatchInstanceRequest) Execute() (PatchInstanceResponse, *_nethttp.Response, error) {
	return r.ApiService.PatchInstanceExecute(r)
}

/*
PatchInstance Update specific instance

Alter a specific instance using its instanceId.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param instanceId The ID of the instance
 @return ApiPatchInstanceRequest
*/
func (a *InstancesApiService) PatchInstance(ctx _context.Context, instanceId int64) ApiPatchInstanceRequest {
	return ApiPatchInstanceRequest{
		ApiService: a,
		ctx: ctx,
		instanceId: instanceId,
	}
}

// Execute executes the request
//  @return PatchInstanceResponse
func (a *InstancesApiService) PatchInstanceExecute(r ApiPatchInstanceRequest) (PatchInstanceResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PatchInstanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstancesApiService.PatchInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/compute/instances/{instanceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", _neturl.PathEscape(parameterToString(r.instanceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}
	if r.patchInstanceRequest == nil {
		return localVarReturnValue, nil, reportError("patchInstanceRequest is required and must be specified")
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
	localVarPostBody = r.patchInstanceRequest
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

type ApiReinstallInstanceRequest struct {
	ctx _context.Context
	ApiService *InstancesApiService
	xRequestId *string
	instanceId int64
	reinstallInstanceRequest *ReinstallInstanceRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiReinstallInstanceRequest) XRequestId(xRequestId string) ApiReinstallInstanceRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiReinstallInstanceRequest) ReinstallInstanceRequest(reinstallInstanceRequest ReinstallInstanceRequest) ApiReinstallInstanceRequest {
	r.reinstallInstanceRequest = &reinstallInstanceRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiReinstallInstanceRequest) XTraceId(xTraceId string) ApiReinstallInstanceRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiReinstallInstanceRequest) Execute() (ReinstallInstanceResponse, *_nethttp.Response, error) {
	return r.ApiService.ReinstallInstanceExecute(r)
}

/*
ReinstallInstance Reinstall specific instance

It is possible to reconfigure a particular instance by installing a new image, and optionally incorporating ssh keys, a root password, or cloud-init.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param instanceId The ID of the instance
 @return ApiReinstallInstanceRequest
*/
func (a *InstancesApiService) ReinstallInstance(ctx _context.Context, instanceId int64) ApiReinstallInstanceRequest {
	return ApiReinstallInstanceRequest{
		ApiService: a,
		ctx: ctx,
		instanceId: instanceId,
	}
}

// Execute executes the request
//  @return ReinstallInstanceResponse
func (a *InstancesApiService) ReinstallInstanceExecute(r ApiReinstallInstanceRequest) (ReinstallInstanceResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReinstallInstanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstancesApiService.ReinstallInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/compute/instances/{instanceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"instanceId"+"}", _neturl.PathEscape(parameterToString(r.instanceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}
	if r.reinstallInstanceRequest == nil {
		return localVarReturnValue, nil, reportError("reinstallInstanceRequest is required and must be specified")
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
	localVarPostBody = r.reinstallInstanceRequest
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

type ApiRetrieveInstanceRequest struct {
	ctx _context.Context
	ApiService *InstancesApiService
	xRequestId *string
	instanceId int64
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveInstanceRequest) XRequestId(xRequestId string) ApiRetrieveInstanceRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveInstanceRequest) XTraceId(xTraceId string) ApiRetrieveInstanceRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiRetrieveInstanceRequest) Execute() (FindInstanceResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveInstanceExecute(r)
}

/*
RetrieveInstance Get specific instance by id

Obtain the values of attributes for a particular instance associated with your account.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param instanceId The ID of the instance
 @return ApiRetrieveInstanceRequest
*/
func (a *InstancesApiService) RetrieveInstance(ctx _context.Context, instanceId int64) ApiRetrieveInstanceRequest {
	return ApiRetrieveInstanceRequest{
		ApiService: a,
		ctx: ctx,
		instanceId: instanceId,
	}
}

// Execute executes the request
//  @return FindInstanceResponse
func (a *InstancesApiService) RetrieveInstanceExecute(r ApiRetrieveInstanceRequest) (FindInstanceResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FindInstanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstancesApiService.RetrieveInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/compute/instances/{instanceId}"
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

type ApiRetrieveInstancesListRequest struct {
	ctx _context.Context
	ApiService *InstancesApiService
	xRequestId *string
	xTraceId *string
	page *int64
	size *int64
	orderBy *[]string
	name *string
	displayName *string
	dataCenter *string
	region *string
	instanceIds *string
	status *string
	addOnIds *string
	productTypes *string
	ipConfig *bool
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveInstancesListRequest) XRequestId(xRequestId string) ApiRetrieveInstancesListRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveInstancesListRequest) XTraceId(xTraceId string) ApiRetrieveInstancesListRequest {
	r.xTraceId = &xTraceId
	return r
}
// Number of page to be fetched.
func (r ApiRetrieveInstancesListRequest) Page(page int64) ApiRetrieveInstancesListRequest {
	r.page = &page
	return r
}
// Number of elements per page.
func (r ApiRetrieveInstancesListRequest) Size(size int64) ApiRetrieveInstancesListRequest {
	r.size = &size
	return r
}
// Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;.
func (r ApiRetrieveInstancesListRequest) OrderBy(orderBy []string) ApiRetrieveInstancesListRequest {
	r.orderBy = &orderBy
	return r
}
// The name of the instance
func (r ApiRetrieveInstancesListRequest) Name(name string) ApiRetrieveInstancesListRequest {
	r.name = &name
	return r
}
// The display name of the instance
func (r ApiRetrieveInstancesListRequest) DisplayName(displayName string) ApiRetrieveInstancesListRequest {
	r.displayName = &displayName
	return r
}
// The data center of the instance
func (r ApiRetrieveInstancesListRequest) DataCenter(dataCenter string) ApiRetrieveInstancesListRequest {
	r.dataCenter = &dataCenter
	return r
}
// The locality of the instance
func (r ApiRetrieveInstancesListRequest) Region(region string) ApiRetrieveInstancesListRequest {
	r.region = &region
	return r
}
// Identifiers of instances separated by commas
func (r ApiRetrieveInstancesListRequest) InstanceIds(instanceIds string) ApiRetrieveInstancesListRequest {
	r.instanceIds = &instanceIds
	return r
}
// The state of the instance
func (r ApiRetrieveInstancesListRequest) Status(status string) ApiRetrieveInstancesListRequest {
	r.status = &status
	return r
}
// Addons IDs of the instance
func (r ApiRetrieveInstancesListRequest) AddOnIds(addOnIds string) ApiRetrieveInstancesListRequest {
	r.addOnIds = &addOnIds
	return r
}
// Categories of instances, separated by commas, based on the ProductId
func (r ApiRetrieveInstancesListRequest) ProductTypes(productTypes string) ApiRetrieveInstancesListRequest {
	r.productTypes = &productTypes
	return r
}
// Filter instances by ip config. True will return those with the ip config set
func (r ApiRetrieveInstancesListRequest) IpConfig(ipConfig bool) ApiRetrieveInstancesListRequest {
	r.ipConfig = &ipConfig
	return r
}

func (r ApiRetrieveInstancesListRequest) Execute() (ListInstancesResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveInstancesListExecute(r)
}

/*
RetrieveInstancesList List of instances

Retrieve all instances in your account and apply filters to refine the results

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRetrieveInstancesListRequest
*/
func (a *InstancesApiService) RetrieveInstancesList(ctx _context.Context) ApiRetrieveInstancesListRequest {
	return ApiRetrieveInstancesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListInstancesResponse
func (a *InstancesApiService) RetrieveInstancesListExecute(r ApiRetrieveInstancesListRequest) (ListInstancesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListInstancesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstancesApiService.RetrieveInstancesList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/compute/instances"

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
	if r.displayName != nil {
		localVarQueryParams.Add("displayName", parameterToString(*r.displayName, ""))
	}
	if r.dataCenter != nil {
		localVarQueryParams.Add("dataCenter", parameterToString(*r.dataCenter, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.instanceIds != nil {
		localVarQueryParams.Add("instanceIds", parameterToString(*r.instanceIds, ""))
	}
	if r.status != nil {
		localVarQueryParams.Add("status", parameterToString(*r.status, ""))
	}
	if r.addOnIds != nil {
		localVarQueryParams.Add("addOnIds", parameterToString(*r.addOnIds, ""))
	}
	if r.productTypes != nil {
		localVarQueryParams.Add("productTypes", parameterToString(*r.productTypes, ""))
	}
	if r.ipConfig != nil {
		localVarQueryParams.Add("ipConfig", parameterToString(*r.ipConfig, ""))
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
