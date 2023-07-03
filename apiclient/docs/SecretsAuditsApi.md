# \SecretsAuditsApi

All URIs are relative to *https://api.fybe.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveSecretAuditsList**](SecretsAuditsApi.md#RetrieveSecretAuditsList) | **Get** /v1/secrets/audits | Display history about your secrets (audit)



## RetrieveSecretAuditsList

> ListSecretAuditResponse RetrieveSecretAuditsList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).SecretId(secretId).RequestId(requestId).ChangedBy(changedBy).StartDate(startDate).EndDate(endDate).Execute()

Display history about your secrets (audit)



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
    secretId := int64(444) // int64 | Secret identifier. (optional)
    requestId := "82690b35-2554-4690-a720-e8f826d20976" // string | The API call's requestId that triggered the modification. (optional)
    changedBy := "82690b35-2554-4690-a720-e8f826d20976" // string | Identifier of the user who made the change. (optional)
    startDate := time.Now() // string | The beginning of the time range for the search. (optional)
    endDate := time.Now() // string | The end date of the time range for the search. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecretsAuditsApi.RetrieveSecretAuditsList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).SecretId(secretId).RequestId(requestId).ChangedBy(changedBy).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsAuditsApi.RetrieveSecretAuditsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSecretAuditsList`: ListSecretAuditResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsAuditsApi.RetrieveSecretAuditsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSecretAuditsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **secretId** | **int64** | Secret identifier. | 
 **requestId** | **string** | The API call&#39;s requestId that triggered the modification. | 
 **changedBy** | **string** | Identifier of the user who made the change. | 
 **startDate** | **string** | The beginning of the time range for the search. | 
 **endDate** | **string** | The end date of the time range for the search. | 

### Return type

[**ListSecretAuditResponse**](ListSecretAuditResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

