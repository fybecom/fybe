# \VirtualPrivateCloudAuditsApi

All URIs are relative to *https://api.fybe.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrievePrivateNetworkAuditsList**](VirtualPrivateCloudAuditsApi.md#RetrievePrivateNetworkAuditsList) | **Get** /v1/private-networks/audits | List history of your VPCs (audit)



## RetrievePrivateNetworkAuditsList

> ListPrivateNetworkAuditResponse RetrievePrivateNetworkAuditsList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).PrivateNetworkId(privateNetworkId).RequestId(requestId).ChangedBy(changedBy).StartDate(startDate).EndDate(endDate).Execute()

List history of your VPCs (audit)



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
    privateNetworkId := int64(98765) // int64 | Virtual Private Cloud identifier. (optional)
    requestId := "71022538-7619-4d0b-be3e-a5846248a087" // string | The Id of the API call that triggered the change. (optional)
    changedBy := "1e38da17-1361-4c19-8f7b-4fac72b80441" // string | Identifier for the user who made the change. (optional)
    startDate := time.Now() // string | Beginning of the search period. (optional)
    endDate := time.Now() // string | Ending of the search period. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualPrivateCloudAuditsApi.RetrievePrivateNetworkAuditsList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).PrivateNetworkId(privateNetworkId).RequestId(requestId).ChangedBy(changedBy).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateCloudAuditsApi.RetrievePrivateNetworkAuditsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePrivateNetworkAuditsList`: ListPrivateNetworkAuditResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateCloudAuditsApi.RetrievePrivateNetworkAuditsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePrivateNetworkAuditsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **privateNetworkId** | **int64** | Virtual Private Cloud identifier. | 
 **requestId** | **string** | The Id of the API call that triggered the change. | 
 **changedBy** | **string** | Identifier for the user who made the change. | 
 **startDate** | **string** | Beginning of the search period. | 
 **endDate** | **string** | Ending of the search period. | 

### Return type

[**ListPrivateNetworkAuditResponse**](ListPrivateNetworkAuditResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

