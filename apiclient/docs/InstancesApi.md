# \InstancesApi

All URIs are relative to *https://api.fybe.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelInstance**](InstancesApi.md#CancelInstance) | **Post** /v1/compute/instances/{instanceId}/cancel | Cancel specific instance by id
[**CreateInstance**](InstancesApi.md#CreateInstance) | **Post** /v1/compute/instances | Create a new compute instance
[**PatchInstance**](InstancesApi.md#PatchInstance) | **Patch** /v1/compute/instances/{instanceId} | Update specific instance
[**ReinstallInstance**](InstancesApi.md#ReinstallInstance) | **Put** /v1/compute/instances/{instanceId} | Reinstall specific instance
[**RetrieveInstance**](InstancesApi.md#RetrieveInstance) | **Get** /v1/compute/instances/{instanceId} | Get specific instance by id
[**RetrieveInstancesList**](InstancesApi.md#RetrieveInstancesList) | **Get** /v1/compute/instances | List of instances



## CancelInstance

> CancelInstanceResponse CancelInstance(ctx, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Cancel specific instance by id



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
    instanceId := int64(56789) // int64 | The ID of the instance
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.CancelInstance(context.Background(), instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.CancelInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelInstance`: CancelInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.CancelInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **int64** | The ID of the instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**CancelInstanceResponse**](CancelInstanceResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInstance

> CreateInstanceResponse CreateInstance(ctx).XRequestId(xRequestId).CreateInstanceRequest(createInstanceRequest).XTraceId(xTraceId).Execute()

Create a new compute instance



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
    createInstanceRequest := *openapiclient.NewCreateInstanceRequest("V3", int64(1)) // CreateInstanceRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.CreateInstance(context.Background()).XRequestId(xRequestId).CreateInstanceRequest(createInstanceRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.CreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInstance`: CreateInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.CreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **createInstanceRequest** | [**CreateInstanceRequest**](CreateInstanceRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**CreateInstanceResponse**](CreateInstanceResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchInstance

> PatchInstanceResponse PatchInstance(ctx, instanceId).XRequestId(xRequestId).PatchInstanceRequest(patchInstanceRequest).XTraceId(xTraceId).Execute()

Update specific instance



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
    instanceId := int64(56789) // int64 | The ID of the instance
    patchInstanceRequest := *openapiclient.NewPatchInstanceRequest() // PatchInstanceRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.PatchInstance(context.Background(), instanceId).XRequestId(xRequestId).PatchInstanceRequest(patchInstanceRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.PatchInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchInstance`: PatchInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.PatchInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **int64** | The ID of the instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **patchInstanceRequest** | [**PatchInstanceRequest**](PatchInstanceRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**PatchInstanceResponse**](PatchInstanceResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReinstallInstance

> ReinstallInstanceResponse ReinstallInstance(ctx, instanceId).XRequestId(xRequestId).ReinstallInstanceRequest(reinstallInstanceRequest).XTraceId(xTraceId).Execute()

Reinstall specific instance



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
    instanceId := int64(56789) // int64 | The ID of the instance
    reinstallInstanceRequest := *openapiclient.NewReinstallInstanceRequest("3f184ab8-a600-4e7c-8c9b-3413e21a3752") // ReinstallInstanceRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.ReinstallInstance(context.Background(), instanceId).XRequestId(xRequestId).ReinstallInstanceRequest(reinstallInstanceRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.ReinstallInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReinstallInstance`: ReinstallInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.ReinstallInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **int64** | The ID of the instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiReinstallInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **reinstallInstanceRequest** | [**ReinstallInstanceRequest**](ReinstallInstanceRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**ReinstallInstanceResponse**](ReinstallInstanceResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveInstance

> FindInstanceResponse RetrieveInstance(ctx, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get specific instance by id



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
    instanceId := int64(56789) // int64 | The ID of the instance
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.RetrieveInstance(context.Background(), instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.RetrieveInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveInstance`: FindInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.RetrieveInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **int64** | The ID of the instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**FindInstanceResponse**](FindInstanceResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveInstancesList

> ListInstancesResponse RetrieveInstancesList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Name(name).DisplayName(displayName).DataCenter(dataCenter).Region(region).InstanceIds(instanceIds).Status(status).AddOnIds(addOnIds).ProductTypes(productTypes).IpConfig(ipConfig).Execute()

List of instances



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
    name := "vm98765" // string | The name of the instance (optional)
    displayName := "myTestInstance" // string | The display name of the instance (optional)
    dataCenter := "European Union (Germany) 1" // string | The data center of the instance (optional)
    region := "EU" // string | The locality of the instance (optional)
    instanceIds := "1000, 1001, 1002" // string | Identifiers of instances separated by commas (optional)
    status := "installing, running" // string | The state of the instance (optional)
    addOnIds := "1044,827" // string | Addons IDs of the instance (optional)
    productTypes := "ssd, nvme" // string | Categories of instances, separated by commas, based on the ProductId (optional)
    ipConfig := true // bool | Filter instances by ip config. True will return those with the ip config set (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.RetrieveInstancesList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Name(name).DisplayName(displayName).DataCenter(dataCenter).Region(region).InstanceIds(instanceIds).Status(status).AddOnIds(addOnIds).ProductTypes(productTypes).IpConfig(ipConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.RetrieveInstancesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveInstancesList`: ListInstancesResponse
    fmt.Fprintf(os.Stdout, "Response from `InstancesApi.RetrieveInstancesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveInstancesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **name** | **string** | The name of the instance | 
 **displayName** | **string** | The display name of the instance | 
 **dataCenter** | **string** | The data center of the instance | 
 **region** | **string** | The locality of the instance | 
 **instanceIds** | **string** | Identifiers of instances separated by commas | 
 **status** | **string** | The state of the instance | 
 **addOnIds** | **string** | Addons IDs of the instance | 
 **productTypes** | **string** | Categories of instances, separated by commas, based on the ProductId | 
 **ipConfig** | **bool** | Filter instances by ip config. True will return those with the ip config set | 

### Return type

[**ListInstancesResponse**](ListInstancesResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

