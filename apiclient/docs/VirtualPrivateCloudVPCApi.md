# \VirtualPrivateCloudVPCApi

All URIs are relative to *https://api.fybe.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignInstancePrivateNetwork**](VirtualPrivateCloudVPCApi.md#AssignInstancePrivateNetwork) | **Post** /v1/private-networks/{privateNetworkId}/instances/{instanceId} | Add instance to a Virtual Private Cloud
[**CreatePrivateNetwork**](VirtualPrivateCloudVPCApi.md#CreatePrivateNetwork) | **Post** /v1/private-networks | Create a new Virtual Private Cloud
[**DeletePrivateNetwork**](VirtualPrivateCloudVPCApi.md#DeletePrivateNetwork) | **Delete** /v1/private-networks/{privateNetworkId} | Delete the virtual private cloud with the given id
[**PatchPrivateNetwork**](VirtualPrivateCloudVPCApi.md#PatchPrivateNetwork) | **Patch** /v1/private-networks/{privateNetworkId} | Update a Virtual Private Cloud by its id
[**RetrievePrivateNetwork**](VirtualPrivateCloudVPCApi.md#RetrievePrivateNetwork) | **Get** /v1/private-networks/{privateNetworkId} | Get specific Virtual Private Cloud by its id
[**RetrievePrivateNetworkList**](VirtualPrivateCloudVPCApi.md#RetrievePrivateNetworkList) | **Get** /v1/private-networks | List all Virtual Private Clouds
[**UnassignInstancePrivateNetwork**](VirtualPrivateCloudVPCApi.md#UnassignInstancePrivateNetwork) | **Delete** /v1/private-networks/{privateNetworkId}/instances/{instanceId} | Remove particular instance from a Virtual Private Cloud



## AssignInstancePrivateNetwork

> AssignInstancePrivateNetworkResponse AssignInstancePrivateNetwork(ctx, privateNetworkId, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Add instance to a Virtual Private Cloud



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
    privateNetworkId := int64(97531) // int64 | Virtual Private Cloud Identifier
    instanceId := int64(555) // int64 | Compute Instance identifier
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualPrivateCloudVPCApi.AssignInstancePrivateNetwork(context.Background(), privateNetworkId, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateCloudVPCApi.AssignInstancePrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignInstancePrivateNetwork`: AssignInstancePrivateNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateCloudVPCApi.AssignInstancePrivateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNetworkId** | **int64** | Virtual Private Cloud Identifier | 
**instanceId** | **int64** | Compute Instance identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignInstancePrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**AssignInstancePrivateNetworkResponse**](AssignInstancePrivateNetworkResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivateNetwork

> CreatePrivateNetworkResponse CreatePrivateNetwork(ctx).XRequestId(xRequestId).CreatePrivateNetworkRequest(createPrivateNetworkRequest).XTraceId(xTraceId).Execute()

Create a new Virtual Private Cloud



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
    createPrivateNetworkRequest := *openapiclient.NewCreatePrivateNetworkRequest("My VPC") // CreatePrivateNetworkRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualPrivateCloudVPCApi.CreatePrivateNetwork(context.Background()).XRequestId(xRequestId).CreatePrivateNetworkRequest(createPrivateNetworkRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateCloudVPCApi.CreatePrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePrivateNetwork`: CreatePrivateNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateCloudVPCApi.CreatePrivateNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **createPrivateNetworkRequest** | [**CreatePrivateNetworkRequest**](CreatePrivateNetworkRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**CreatePrivateNetworkResponse**](CreatePrivateNetworkResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrivateNetwork

> DeletePrivateNetwork(ctx, privateNetworkId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Delete the virtual private cloud with the given id



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
    privateNetworkId := int64(97531) // int64 | Virtual Private Cloud ID
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualPrivateCloudVPCApi.DeletePrivateNetwork(context.Background(), privateNetworkId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateCloudVPCApi.DeletePrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNetworkId** | **int64** | Virtual Private Cloud ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPrivateNetwork

> PatchPrivateNetworkResponse PatchPrivateNetwork(ctx, privateNetworkId).XRequestId(xRequestId).PatchPrivateNetworkRequest(patchPrivateNetworkRequest).XTraceId(xTraceId).Execute()

Update a Virtual Private Cloud by its id



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
    privateNetworkId := int64(97531) // int64 | Virtual Private Cloud ID
    patchPrivateNetworkRequest := *openapiclient.NewPatchPrivateNetworkRequest() // PatchPrivateNetworkRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualPrivateCloudVPCApi.PatchPrivateNetwork(context.Background(), privateNetworkId).XRequestId(xRequestId).PatchPrivateNetworkRequest(patchPrivateNetworkRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateCloudVPCApi.PatchPrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPrivateNetwork`: PatchPrivateNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateCloudVPCApi.PatchPrivateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNetworkId** | **int64** | Virtual Private Cloud ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **patchPrivateNetworkRequest** | [**PatchPrivateNetworkRequest**](PatchPrivateNetworkRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**PatchPrivateNetworkResponse**](PatchPrivateNetworkResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePrivateNetwork

> FindPrivateNetworkResponse RetrievePrivateNetwork(ctx, privateNetworkId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get specific Virtual Private Cloud by its id



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
    privateNetworkId := int64(97531) // int64 | Virtual Private Cloud ID
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualPrivateCloudVPCApi.RetrievePrivateNetwork(context.Background(), privateNetworkId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateCloudVPCApi.RetrievePrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePrivateNetwork`: FindPrivateNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateCloudVPCApi.RetrievePrivateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNetworkId** | **int64** | Virtual Private Cloud ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**FindPrivateNetworkResponse**](FindPrivateNetworkResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePrivateNetworkList

> ListPrivateNetworkResponse RetrievePrivateNetworkList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Name(name).InstanceIds(instanceIds).Region(region).DataCenter(dataCenter).Execute()

List all Virtual Private Clouds



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
    name := "My VPC" // string | Title of the Virtual Private Cloud (optional)
    instanceIds := "1, 2, 3" // string | Compute instance identifiers, separated by comma (optional)
    region := "US-east" // string | The region slug indicating the location of your VPC. (optional)
    dataCenter := "US East" // string | The physical facility housing your VPC. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualPrivateCloudVPCApi.RetrievePrivateNetworkList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Name(name).InstanceIds(instanceIds).Region(region).DataCenter(dataCenter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateCloudVPCApi.RetrievePrivateNetworkList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePrivateNetworkList`: ListPrivateNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateCloudVPCApi.RetrievePrivateNetworkList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePrivateNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **name** | **string** | Title of the Virtual Private Cloud | 
 **instanceIds** | **string** | Compute instance identifiers, separated by comma | 
 **region** | **string** | The region slug indicating the location of your VPC. | 
 **dataCenter** | **string** | The physical facility housing your VPC. | 

### Return type

[**ListPrivateNetworkResponse**](ListPrivateNetworkResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignInstancePrivateNetwork

> UnassignInstancePrivateNetworkResponse UnassignInstancePrivateNetwork(ctx, privateNetworkId, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Remove particular instance from a Virtual Private Cloud



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
    privateNetworkId := int64(555) // int64 | Virtual Private Cloud Identifier
    instanceId := int64(555) // int64 | Compute Instance Identifier
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirtualPrivateCloudVPCApi.UnassignInstancePrivateNetwork(context.Background(), privateNetworkId, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateCloudVPCApi.UnassignInstancePrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnassignInstancePrivateNetwork`: UnassignInstancePrivateNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateCloudVPCApi.UnassignInstancePrivateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNetworkId** | **int64** | Virtual Private Cloud Identifier | 
**instanceId** | **int64** | Compute Instance Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignInstancePrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**UnassignInstancePrivateNetworkResponse**](UnassignInstancePrivateNetworkResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

