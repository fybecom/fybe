# \RolesAuditsApi

All URIs are relative to *https://api.fybe.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveRoleAuditsList**](RolesAuditsApi.md#RetrieveRoleAuditsList) | **Get** /v1/roles/audits | Retrieve the audit history of your roles.



## RetrieveRoleAuditsList

> ListRoleAuditResponse RetrieveRoleAuditsList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).RoleId(roleId).RequestId(requestId).ChangedBy(changedBy).StartDate(startDate).EndDate(endDate).Execute()

Retrieve the audit history of your roles.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    page := int64(1) // int64 | Number of page to be fetched. (optional)
    size := int64(10) // int64 | Number of elements per page. (optional)
    orderBy := []string{"Inner_example"} // []string | Specify fields and ordering (ASC for ascending, DESC for descending) in following format `field:ASC|DESC`. (optional)
    roleId := int64(2000) // int64 | Role ID. (optional)
    requestId := "d31a0953-5ba8-4ecd-ae93-dd56008af7b6" // string | Request ID. (optional)
    changedBy := "ca38cffc-5113-4049-9daa-a5c25e26f300" // string | The `changedBy` field indicates the user who made the change. (optional)
    startDate := time.Now() // string | Start time of the search range. (optional)
    endDate := time.Now() // string | The end time of the time range being searched. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RolesAuditsApi.RetrieveRoleAuditsList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).RoleId(roleId).RequestId(requestId).ChangedBy(changedBy).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesAuditsApi.RetrieveRoleAuditsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveRoleAuditsList`: ListRoleAuditResponse
    fmt.Fprintf(os.Stdout, "Response from `RolesAuditsApi.RetrieveRoleAuditsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRoleAuditsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **roleId** | **int64** | Role ID. | 
 **requestId** | **string** | Request ID. | 
 **changedBy** | **string** | The &#x60;changedBy&#x60; field indicates the user who made the change. | 
 **startDate** | **string** | Start time of the search range. | 
 **endDate** | **string** | The end time of the time range being searched. | 

### Return type

[**ListRoleAuditResponse**](ListRoleAuditResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

