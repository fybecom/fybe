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
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// ObjectStoragesAuditsApiService ObjectStoragesAuditsApi service
type ObjectStoragesAuditsApiService service

type ApiRetrieveObjectStorageAuditsListRequest struct {
	ctx _context.Context
	ApiService *ObjectStoragesAuditsApiService
	xRequestId *string
	xTraceId *string
	page *int64
	size *int64
	orderBy *[]string
	objectStorageId *string
	requestId *string
	changedBy *string
	startDate *string
	endDate *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveObjectStorageAuditsListRequest) XRequestId(xRequestId string) ApiRetrieveObjectStorageAuditsListRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveObjectStorageAuditsListRequest) XTraceId(xTraceId string) ApiRetrieveObjectStorageAuditsListRequest {
	r.xTraceId = &xTraceId
	return r
}
// Number of page to be fetched.
func (r ApiRetrieveObjectStorageAuditsListRequest) Page(page int64) ApiRetrieveObjectStorageAuditsListRequest {
	r.page = &page
	return r
}
// Number of elements per page.
func (r ApiRetrieveObjectStorageAuditsListRequest) Size(size int64) ApiRetrieveObjectStorageAuditsListRequest {
	r.size = &size
	return r
}
// Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;.
func (r ApiRetrieveObjectStorageAuditsListRequest) OrderBy(orderBy []string) ApiRetrieveObjectStorageAuditsListRequest {
	r.orderBy = &orderBy
	return r
}
// Object Storage Identifier.
func (r ApiRetrieveObjectStorageAuditsListRequest) ObjectStorageId(objectStorageId string) ApiRetrieveObjectStorageAuditsListRequest {
	r.objectStorageId = &objectStorageId
	return r
}
// The ID of the API call request that made the change.
func (r ApiRetrieveObjectStorageAuditsListRequest) RequestId(requestId string) ApiRetrieveObjectStorageAuditsListRequest {
	r.requestId = &requestId
	return r
}
// ID of the user that made the change.
func (r ApiRetrieveObjectStorageAuditsListRequest) ChangedBy(changedBy string) ApiRetrieveObjectStorageAuditsListRequest {
	r.changedBy = &changedBy
	return r
}
// Begin of search range.
func (r ApiRetrieveObjectStorageAuditsListRequest) StartDate(startDate string) ApiRetrieveObjectStorageAuditsListRequest {
	r.startDate = &startDate
	return r
}
// End date of search range.
func (r ApiRetrieveObjectStorageAuditsListRequest) EndDate(endDate string) ApiRetrieveObjectStorageAuditsListRequest {
	r.endDate = &endDate
	return r
}

func (r ApiRetrieveObjectStorageAuditsListRequest) Execute() (ListObjectStorageAuditResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveObjectStorageAuditsListExecute(r)
}

/*
RetrieveObjectStorageAuditsList List history about your object storages

Displays history about your object storages using filters.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRetrieveObjectStorageAuditsListRequest
*/
func (a *ObjectStoragesAuditsApiService) RetrieveObjectStorageAuditsList(ctx _context.Context) ApiRetrieveObjectStorageAuditsListRequest {
	return ApiRetrieveObjectStorageAuditsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListObjectStorageAuditResponse
func (a *ObjectStoragesAuditsApiService) RetrieveObjectStorageAuditsListExecute(r ApiRetrieveObjectStorageAuditsListRequest) (ListObjectStorageAuditResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListObjectStorageAuditResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectStoragesAuditsApiService.RetrieveObjectStorageAuditsList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/object-storages/audits"

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
	if r.requestId != nil {
		localVarQueryParams.Add("requestId", parameterToString(*r.requestId, ""))
	}
	if r.changedBy != nil {
		localVarQueryParams.Add("changedBy", parameterToString(*r.changedBy, ""))
	}
	if r.startDate != nil {
		localVarQueryParams.Add("startDate", parameterToString(*r.startDate, ""))
	}
	if r.endDate != nil {
		localVarQueryParams.Add("endDate", parameterToString(*r.endDate, ""))
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
