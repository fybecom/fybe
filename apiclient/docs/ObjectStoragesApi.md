# \ObjectStoragesApi

All URIs are relative to *https://api.fybe.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelObjectStorage**](ObjectStoragesApi.md#CancelObjectStorage) | **Patch** /v1/object-storages/{objectStorageId}/cancel | Cancels the selected object storage at the next possible date
[**CreateObjectStorage**](ObjectStoragesApi.md#CreateObjectStorage) | **Post** /v1/object-storages | Create a new object storage
[**RetrieveDataCenterList**](ObjectStoragesApi.md#RetrieveDataCenterList) | **Get** /v1/data-centers | List all data centers
[**RetrieveObjectStorage**](ObjectStoragesApi.md#RetrieveObjectStorage) | **Get** /v1/object-storages/{objectStorageId} | Get a specific object storage using its id
[**RetrieveObjectStorageList**](ObjectStoragesApi.md#RetrieveObjectStorageList) | **Get** /v1/object-storages | List all your object storages
[**RetrieveObjectStoragesStats**](ObjectStoragesApi.md#RetrieveObjectStoragesStats) | **Get** /v1/object-storages/{objectStorageId}/stats | Gives statistics of selected object storage
[**UpdateObjectStorage**](ObjectStoragesApi.md#UpdateObjectStorage) | **Patch** /v1/object-storages/{objectStorageId} | Updates the display name of object storage
[**UpgradeObjectStorage**](ObjectStoragesApi.md#UpgradeObjectStorage) | **Post** /v1/object-storages/{objectStorageId}/resize | Upgrade object storage size / update autoscaling settings.



## CancelObjectStorage

> CancelObjectStorageResponse CancelObjectStorage(ctx, objectStorageId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Cancels the selected object storage at the next possible date



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    objectStorageId := "a8df4fab-6e70-4b52-ae84-b7f88087cd6f" // string | Object Storage Identifier.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStoragesApi.CancelObjectStorage(context.Background(), objectStorageId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoragesApi.CancelObjectStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelObjectStorage`: CancelObjectStorageResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoragesApi.CancelObjectStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | Object Storage Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelObjectStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**CancelObjectStorageResponse**](CancelObjectStorageResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateObjectStorage

> CreateObjectStorageResponse CreateObjectStorage(ctx).XRequestId(xRequestId).CreateObjectStorageRequest(createObjectStorageRequest).XTraceId(xTraceId).Execute()

Create a new object storage



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    createObjectStorageRequest := *openapiclient.NewCreateObjectStorageRequest("us-central-1", float64(12)) // CreateObjectStorageRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStoragesApi.CreateObjectStorage(context.Background()).XRequestId(xRequestId).CreateObjectStorageRequest(createObjectStorageRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoragesApi.CreateObjectStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateObjectStorage`: CreateObjectStorageResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoragesApi.CreateObjectStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **createObjectStorageRequest** | [**CreateObjectStorageRequest**](CreateObjectStorageRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**CreateObjectStorageResponse**](CreateObjectStorageResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDataCenterList

> ListDataCenterResponse RetrieveDataCenterList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Slug(slug).Name(name).RegionName(regionName).RegionSlug(regionSlug).Execute()

List all data centers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    page := int64(1) // int64 | Number of page to be fetched. (optional)
    size := int64(10) // int64 | Number of elements per page. (optional)
    orderBy := []string{"Inner_example"} // []string | Specify fields and ordering (ASC for ascending, DESC for descending) in following format `field:ASC|DESC`. (optional)
    slug := "SIN" // string | Filter that can be used as match for data centers. (optional)
    name := "United States 1" // string | Available region filter for object storage. (optional)
    regionName := "United States (East)" // string | Available region name filter for object storage. (optional)
    regionSlug := "US" // string | Available region slug filter for object storage. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStoragesApi.RetrieveDataCenterList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Slug(slug).Name(name).RegionName(regionName).RegionSlug(regionSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoragesApi.RetrieveDataCenterList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDataCenterList`: ListDataCenterResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoragesApi.RetrieveDataCenterList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDataCenterListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **slug** | **string** | Filter that can be used as match for data centers. | 
 **name** | **string** | Available region filter for object storage. | 
 **regionName** | **string** | Available region name filter for object storage. | 
 **regionSlug** | **string** | Available region slug filter for object storage. | 

### Return type

[**ListDataCenterResponse**](ListDataCenterResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveObjectStorage

> FindObjectStorageResponse RetrieveObjectStorage(ctx, objectStorageId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get a specific object storage using its id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    objectStorageId := "a8df4fab-6e70-4b52-ae84-b7f88087cd6f" // string | Object Storage Identifier.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStoragesApi.RetrieveObjectStorage(context.Background(), objectStorageId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoragesApi.RetrieveObjectStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveObjectStorage`: FindObjectStorageResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoragesApi.RetrieveObjectStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | Object Storage Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveObjectStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**FindObjectStorageResponse**](FindObjectStorageResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveObjectStorageList

> ListObjectStorageResponse RetrieveObjectStorageList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).DataCenterName(dataCenterName).S3TenantId(s3TenantId).Region(region).Execute()

List all your object storages



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    page := int64(1) // int64 | Number of page to be fetched. (optional)
    size := int64(10) // int64 | Number of elements per page. (optional)
    orderBy := []string{"Inner_example"} // []string | Specify fields and ordering (ASC for ascending, DESC for descending) in following format `field:ASC|DESC`. (optional)
    dataCenterName := "United States (US East)" // string | Available filter for selected object storage locations. (optional)
    s3TenantId := "5ab4f7i9865v71d1zug22d1230dcav77" // string | Available filter for object storage S3 tenantId. (optional)
    region := "SIN" // string | Available region filter for Object Storage (Choice between US-central, EU and SIN) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStoragesApi.RetrieveObjectStorageList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).DataCenterName(dataCenterName).S3TenantId(s3TenantId).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoragesApi.RetrieveObjectStorageList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveObjectStorageList`: ListObjectStorageResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoragesApi.RetrieveObjectStorageList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveObjectStorageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **dataCenterName** | **string** | Available filter for selected object storage locations. | 
 **s3TenantId** | **string** | Available filter for object storage S3 tenantId. | 
 **region** | **string** | Available region filter for Object Storage (Choice between US-central, EU and SIN) | 

### Return type

[**ListObjectStorageResponse**](ListObjectStorageResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveObjectStoragesStats

> ObjectStoragesStatsResponse RetrieveObjectStoragesStats(ctx, objectStorageId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Gives statistics of selected object storage



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    objectStorageId := "a8df4fab-6e70-4b52-ae84-b7f88087cd6f" // string | Object Storage Identifier.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStoragesApi.RetrieveObjectStoragesStats(context.Background(), objectStorageId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoragesApi.RetrieveObjectStoragesStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveObjectStoragesStats`: ObjectStoragesStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoragesApi.RetrieveObjectStoragesStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | Object Storage Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveObjectStoragesStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**ObjectStoragesStatsResponse**](ObjectStoragesStatsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateObjectStorage

> CancelObjectStorageResponse UpdateObjectStorage(ctx, objectStorageId).XRequestId(xRequestId).PatchObjectStorageRequest(patchObjectStorageRequest).XTraceId(xTraceId).Execute()

Updates the display name of object storage



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    objectStorageId := "a8df4fab-6e70-4b52-ae84-b7f88087cd6f" // string | Object Storage Identifier.
    patchObjectStorageRequest := *openapiclient.NewPatchObjectStorageRequest("myObjectStorage US 999") // PatchObjectStorageRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStoragesApi.UpdateObjectStorage(context.Background(), objectStorageId).XRequestId(xRequestId).PatchObjectStorageRequest(patchObjectStorageRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoragesApi.UpdateObjectStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateObjectStorage`: CancelObjectStorageResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoragesApi.UpdateObjectStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | Object Storage Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **patchObjectStorageRequest** | [**PatchObjectStorageRequest**](PatchObjectStorageRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**CancelObjectStorageResponse**](CancelObjectStorageResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeObjectStorage

> UpgradeObjectStorageResponse UpgradeObjectStorage(ctx, objectStorageId).XRequestId(xRequestId).UpgradeObjectStorageRequest(upgradeObjectStorageRequest).XTraceId(xTraceId).Execute()

Upgrade object storage size / update autoscaling settings.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    objectStorageId := "a8df4fab-6e70-4b52-ae84-b7f88087cd6f" // string | Object Storage Identifier.
    upgradeObjectStorageRequest := *openapiclient.NewUpgradeObjectStorageRequest() // UpgradeObjectStorageRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectStoragesApi.UpgradeObjectStorage(context.Background(), objectStorageId).XRequestId(xRequestId).UpgradeObjectStorageRequest(upgradeObjectStorageRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectStoragesApi.UpgradeObjectStorage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeObjectStorage`: UpgradeObjectStorageResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectStoragesApi.UpgradeObjectStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectStorageId** | **string** | Object Storage Identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeObjectStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **upgradeObjectStorageRequest** | [**UpgradeObjectStorageRequest**](UpgradeObjectStorageRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**UpgradeObjectStorageResponse**](UpgradeObjectStorageResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

