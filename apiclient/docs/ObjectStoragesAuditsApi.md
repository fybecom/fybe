# \ObjectStoragesAuditsApi

All URIs are relative to *https://api.fybe.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveObjectStorageAuditsList**](ObjectStoragesAuditsApi.md#RetrieveObjectStorageAuditsList) | **Get** /v1/object-storages/audits | List history about your object storages



## RetrieveObjectStorageAuditsList

> ListObjectStorageAuditResponse RetrieveObjectStorageAuditsList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).ObjectStorageId(objectStorageId).RequestId(requestId).ChangedBy(changedBy).StartDate(startDate).EndDate(endDate).Execute()

List history about your object storages



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
    objectStorageId := "a8df4fab-6e70-4b52-ae84-b7f88087cd6f" // string | Object Storage Identifier. (optional)
    requestId := "a0f699a8-c451-4fce-99bc-9ba43f9c81af" // string | The ID of the API call request that made the change. (optional)
    changedBy := "9c5dc0dd-033f-43e6-8cff-31784d5c49a8" // string | ID of the user that made the change. (optional)
    startDate := time.Now() // string | Begin of search range. (optional)
    endDate := time.Now() // string | End date of search range. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStoragesAuditsApi.RetrieveObjectStorageAuditsList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).ObjectStorageId(objectStorageId).RequestId(requestId).ChangedBy(changedBy).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoragesAuditsApi.RetrieveObjectStorageAuditsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveObjectStorageAuditsList`: ListObjectStorageAuditResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoragesAuditsApi.RetrieveObjectStorageAuditsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveObjectStorageAuditsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **objectStorageId** | **string** | Object Storage Identifier. | 
 **requestId** | **string** | The ID of the API call request that made the change. | 
 **changedBy** | **string** | ID of the user that made the change. | 
 **startDate** | **string** | Begin of search range. | 
 **endDate** | **string** | End date of search range. | 

### Return type

[**ListObjectStorageAuditResponse**](ListObjectStorageAuditResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

