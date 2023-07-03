# \TagAssignmentsAuditsApi

All URIs are relative to *https://api.fybe.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveAssignmentsAuditsList**](TagAssignmentsAuditsApi.md#RetrieveAssignmentsAuditsList) | **Get** /v1/tags/assignments/audits | List a history about the assignments (audit)



## RetrieveAssignmentsAuditsList

> ListAssignmentAuditsResponse RetrieveAssignmentsAuditsList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).TagId(tagId).ResourceId(resourceId).RequestId(requestId).ChangedBy(changedBy).StartDate(startDate).EndDate(endDate).Execute()

List a history about the assignments (audit)



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
    tagId := int64(98765) // int64 | Tag ID. (optional)
    resourceId := "Z9Y8X7" // string | Resource ID. (optional)
    requestId := "4d989930-1327-44ce-8865-25ea5de9081b" // string | The Identifier of the API call request that made the change. (optional)
    changedBy := "d7ff5581-6f05-47d8-a6f0-2e8c1a8108bf" // string | Identifier of the User that made the change. (optional)
    startDate := time.Now() // string | Start date of search range. (optional)
    endDate := time.Now() // string | End date of search range. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagAssignmentsAuditsApi.RetrieveAssignmentsAuditsList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).TagId(tagId).ResourceId(resourceId).RequestId(requestId).ChangedBy(changedBy).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagAssignmentsAuditsApi.RetrieveAssignmentsAuditsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAssignmentsAuditsList`: ListAssignmentAuditsResponse
    fmt.Fprintf(os.Stdout, "Response from `TagAssignmentsAuditsApi.RetrieveAssignmentsAuditsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAssignmentsAuditsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **tagId** | **int64** | Tag ID. | 
 **resourceId** | **string** | Resource ID. | 
 **requestId** | **string** | The Identifier of the API call request that made the change. | 
 **changedBy** | **string** | Identifier of the User that made the change. | 
 **startDate** | **string** | Start date of search range. | 
 **endDate** | **string** | End date of search range. | 

### Return type

[**ListAssignmentAuditsResponse**](ListAssignmentAuditsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

