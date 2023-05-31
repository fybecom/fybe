# \InstanceActionsAuditsApi

All URIs are relative to *https://api.fybe.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveInstancesActionsAuditsList**](InstanceActionsAuditsApi.md#RetrieveInstancesActionsAuditsList) | **Get** /v1/compute/instances/actions/audits | List history about your actions (audit) triggered via the API



## RetrieveInstancesActionsAuditsList

> ListInstancesActionsAuditResponse RetrieveInstancesActionsAuditsList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).InstanceId(instanceId).RequestId(requestId).ChangedBy(changedBy).StartDate(startDate).EndDate(endDate).Execute()

List history about your actions (audit) triggered via the API



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
    instanceId := int64(12345) // int64 | The ID of the compute instancesActions. (optional)
    requestId := "D5FD9FAF-58C0-4406-8F46-F449B8E4FEC3" // string | The requestId of the API call which led to the change. (optional)
    changedBy := "23cbb6d6-cb11-4330-bdff-7bb791df2e23" // string | changedBy of the user which led to the change. (optional)
    startDate := time.Now() // string | Start of search time range. (optional)
    endDate := time.Now() // string | End of search time range. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstanceActionsAuditsApi.RetrieveInstancesActionsAuditsList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).InstanceId(instanceId).RequestId(requestId).ChangedBy(changedBy).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceActionsAuditsApi.RetrieveInstancesActionsAuditsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveInstancesActionsAuditsList`: ListInstancesActionsAuditResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceActionsAuditsApi.RetrieveInstancesActionsAuditsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveInstancesActionsAuditsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **instanceId** | **int64** | The ID of the compute instancesActions. | 
 **requestId** | **string** | The requestId of the API call which led to the change. | 
 **changedBy** | **string** | changedBy of the user which led to the change. | 
 **startDate** | **string** | Start of search time range. | 
 **endDate** | **string** | End of search time range. | 

### Return type

[**ListInstancesActionsAuditResponse**](ListInstancesActionsAuditResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

