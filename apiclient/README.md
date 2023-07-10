# Go API client for openapi

# Introduction

The Fybe API facilitates resource management through HTTP requests.
This documentation comprises a collection of HTTP endpoints that adhere to RESTful principles.
Each endpoint is accompanied by descriptions, as well as request and response schemas.

## OpenAPI specification (OAS)

Fybe's OpenAPI Specification can be [downloaded here](https://api.fybe.com/api-v1.yaml).

## Getting Started

To utilize the Fybe API, you'll require following credentials that can be obtained from from the [Fybe Cockpit](https://cockpit.fybe.com/account/security):
1. ClientId
2. ClientSecret
3. API User (your email address to login to the [Fybe Cockpit](https://cockpit.fybe.com/account/security))
4. API Password (this is a new password which you'll set or change in the [Fybe Cockpit](https://cockpit.fybe.com/account/security))

### How to use the API?

As authentication [Bearer Tokens](https://oauth.net/2/bearer-tokens/) in form of [JWT](https://www.rfc-editor.org/rfc/rfc7519) are used. This approach follows the [OAuth 2.0](https://oauth.net/2/) specification.

#### Retrieve the Bearer Token

```sh
POST /auth/realms/Fybe/protocol/openid-connect/token HTTP/1.1
Host: airlock.fybe.com

grant_type=password
&password=xxxxxx
&redirect_uri=https://example-app.com/redirect
&username=xxxxxx
&client_id=xxxxxx
&client_secret=xxxxxx
```

The actual values for `client_id`, `client_secret`, `username` and `password` can be retrieved from [Fybe Cockpit](https://cockpit.fybe.com/account/security)

## Using the Command-Line Interface (CLI)

Fybe provides the CLI called `fybe` which can be downloaded from <https://github.com/fybecom/fybe>. Please follow the instructions in the `README.md` to perform the installation on your OS. `fybe` supports Windows, macOS and Linux operating systems.

Using `fybe` CLI invoking makes invoking the API much more comfortable. E.g. `fybe get instances` for retrieving the list of compute instances.

## Requests

As stated below, the Fybe API accommodates HTTP requests. However, not all endpoints endorse every method. You can find a list of allowed methods within this documentation.

Method | Description
---    | ---
GET    | When you need to obtain information about a resource, you should utilize the GET method. The data will be provided as a JSON object. It's essential to note that GET methods are read-only and don't impact any resources.
POST   | To create a new object, send a POST method that encodes all necessary attributes in the request body as JSON.
PATCH  | PATCH can be used to partially modify a resource, allowing specific attributes to be changed without updating the complete object.
PUT    | If you need to update information about a resource, use the PUT method. PUT will overwrite existing values of the item, without considering the current state.
DELETE | To remove a resource from your account, apply the DELETE method. If the resource is not found, the operation will generate a 4xx error along with a relevant message.

## Responses

Typically, the Fybe API will respond to your requests by returning data in [JSON](https://www.json.org/) format, which can be easily processed using any programming language or tools.

Like for any HTTP requests, you'll receive an HTTP status code, which provides general information about the success or error of the request. The table below outlines some common HTTP status codes.

It's important to note that the description of the endpoints and methods doesn't provide an exhaustive list of all possible status codes. Only specific return codes and their corresponding response data are explicitly outlined.

Response Code | Description
--- | ---
200 | The response will contain the information you requested.
201 | Your request has been acknowledged, and the resource has been created.
204 | Your request was successful, and no further information is provided in the response.
400 | Your request was not properly formed.
401 | You didn't provide valid authentication credentials.
402 | The request was declined as it necessitates an additional paid service.
403 | You are prohibited from performing the request. You'll need to pass a bearer token.
404 | No results were found for your request, or the resource you're trying to access does not exist.
409 | There is a conflict with resources, such as a violation of unique data constraints when attempting to create or modify resources.
429 | The rate limit has been reached. Please wait for some time before making further requests.
500 | We couldn't fulfill your request due to server-side issues. If this happens, please try again later or reach out to our support team.

Not every endpoint returns data. For example DELETE requests usually don't return any data. All others do return data. For easy handling the return values consists of metadata denoted with and underscore (\"_\") like `_links` or `_pagination`. The actual data is returned in a field called `data`. For convenience reasons this `data` field is always returned as an array even if it consists of only one single element.

DELETE requests usually don't return any data. Return values consists of metadata starting with an underscore (\"_\"), e.g. `_links` and `_pagination`. The requested data is to be found in the field `data`. The `data` field is always an array.


Please visit [Fybe](https://fybe.com) for more general information.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://fybe.com](https://fybe.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.fybe.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ImagesApi* | [**CreateCustomImage**](docs/ImagesApi.md#createcustomimage) | **Post** /v1/compute/images | Upload a custom image
*ImagesApi* | [**DeleteImage**](docs/ImagesApi.md#deleteimage) | **Delete** /v1/compute/images/{imageId} | Delete custom image by id
*ImagesApi* | [**RetrieveCustomImagesStats**](docs/ImagesApi.md#retrievecustomimagesstats) | **Get** /v1/compute/images/stats | Retrieve statistics related to the custom images of the customer
*ImagesApi* | [**RetrieveImage**](docs/ImagesApi.md#retrieveimage) | **Get** /v1/compute/images/{imageId} | Obtain information regarding a particular image based on its id.
*ImagesApi* | [**RetrieveImageList**](docs/ImagesApi.md#retrieveimagelist) | **Get** /v1/compute/images | Provide a list of both standard and custom images that are available
*ImagesApi* | [**UpdateImage**](docs/ImagesApi.md#updateimage) | **Patch** /v1/compute/images/{imageId} | Modify the name of a custom image by its id
*ImagesAuditsApi* | [**RetrieveImageAuditsList**](docs/ImagesAuditsApi.md#retrieveimageauditslist) | **Get** /v1/compute/images/audits | Retrieve a list of your custom images history
*InstanceActionsApi* | [**Rescue**](docs/InstanceActionsApi.md#rescue) | **Post** /v1/compute/instances/{instanceId}/actions/rescue | Rescue a compute instance / resource identified by its id
*InstanceActionsApi* | [**ResetPasswordAction**](docs/InstanceActionsApi.md#resetpasswordaction) | **Post** /v1/compute/instances/{instanceId}/actions/resetPassword | Reset password for a compute instance / resource referenced by an id
*InstanceActionsApi* | [**Restart**](docs/InstanceActionsApi.md#restart) | **Post** /v1/compute/instances/{instanceId}/actions/restart | Retrieve a list of your custom images history
*InstanceActionsApi* | [**Shutdown**](docs/InstanceActionsApi.md#shutdown) | **Post** /v1/compute/instances/{instanceId}/actions/shutdown | Shutdown compute instance / resource by its id
*InstanceActionsApi* | [**Start**](docs/InstanceActionsApi.md#start) | **Post** /v1/compute/instances/{instanceId}/actions/start | Start a compute instance / resource identified by its id
*InstanceActionsApi* | [**Stop**](docs/InstanceActionsApi.md#stop) | **Post** /v1/compute/instances/{instanceId}/actions/stop | Stop compute instance / resource by its id
*InstanceActionsAuditsApi* | [**RetrieveInstancesActionsAuditsList**](docs/InstanceActionsAuditsApi.md#retrieveinstancesactionsauditslist) | **Get** /v1/compute/instances/actions/audits | List history about your actions (audit) triggered via the API
*InstancesApi* | [**CancelInstance**](docs/InstancesApi.md#cancelinstance) | **Post** /v1/compute/instances/{instanceId}/cancel | Cancel specific instance by id
*InstancesApi* | [**CreateInstance**](docs/InstancesApi.md#createinstance) | **Post** /v1/compute/instances | Create a new compute instance
*InstancesApi* | [**PatchInstance**](docs/InstancesApi.md#patchinstance) | **Patch** /v1/compute/instances/{instanceId} | Update specific instance
*InstancesApi* | [**ReinstallInstance**](docs/InstancesApi.md#reinstallinstance) | **Put** /v1/compute/instances/{instanceId} | Reinstall specific instance
*InstancesApi* | [**RetrieveInstance**](docs/InstancesApi.md#retrieveinstance) | **Get** /v1/compute/instances/{instanceId} | Get specific instance by id
*InstancesApi* | [**RetrieveInstancesList**](docs/InstancesApi.md#retrieveinstanceslist) | **Get** /v1/compute/instances | List of instances
*InstancesAuditsApi* | [**RetrieveInstancesAuditsList**](docs/InstancesAuditsApi.md#retrieveinstancesauditslist) | **Get** /v1/compute/instances/audits | Retrieve a list of your custom images history
*ObjectStoragesApi* | [**CancelObjectStorage**](docs/ObjectStoragesApi.md#cancelobjectstorage) | **Patch** /v1/object-storages/{objectStorageId}/cancel | Cancels the selected object storage at the next possible date
*ObjectStoragesApi* | [**CreateObjectStorage**](docs/ObjectStoragesApi.md#createobjectstorage) | **Post** /v1/object-storages | Create a new object storage
*ObjectStoragesApi* | [**RetrieveDataCenterList**](docs/ObjectStoragesApi.md#retrievedatacenterlist) | **Get** /v1/data-centers | List all data centers
*ObjectStoragesApi* | [**RetrieveObjectStorage**](docs/ObjectStoragesApi.md#retrieveobjectstorage) | **Get** /v1/object-storages/{objectStorageId} | Get a specific object storage using its id
*ObjectStoragesApi* | [**RetrieveObjectStorageList**](docs/ObjectStoragesApi.md#retrieveobjectstoragelist) | **Get** /v1/object-storages | List all your object storages
*ObjectStoragesApi* | [**RetrieveObjectStoragesStats**](docs/ObjectStoragesApi.md#retrieveobjectstoragesstats) | **Get** /v1/object-storages/{objectStorageId}/stats | Gives statistics of selected object storage
*ObjectStoragesApi* | [**UpdateObjectStorage**](docs/ObjectStoragesApi.md#updateobjectstorage) | **Patch** /v1/object-storages/{objectStorageId} | Updates the display name of object storage
*ObjectStoragesApi* | [**UpgradeObjectStorage**](docs/ObjectStoragesApi.md#upgradeobjectstorage) | **Post** /v1/object-storages/{objectStorageId}/resize | Upgrade object storage size / update autoscaling settings.
*ObjectStoragesAuditsApi* | [**RetrieveObjectStorageAuditsList**](docs/ObjectStoragesAuditsApi.md#retrieveobjectstorageauditslist) | **Get** /v1/object-storages/audits | List history about your object storages
*RolesApi* | [**CreateRole**](docs/RolesApi.md#createrole) | **Post** /v1/roles | Create a role
*RolesApi* | [**DeleteRole**](docs/RolesApi.md#deleterole) | **Delete** /v1/roles/{roleId} | Delete a role
*RolesApi* | [**RetrieveApiPermissionsList**](docs/RolesApi.md#retrieveapipermissionslist) | **Get** /v1/roles/api-permissions | API permissions list
*RolesApi* | [**RetrieveRole**](docs/RolesApi.md#retrieverole) | **Get** /v1/roles/{roleId} | Get a role
*RolesApi* | [**RetrieveRoleList**](docs/RolesApi.md#retrieverolelist) | **Get** /v1/roles | Roles list
*RolesApi* | [**UpdateRole**](docs/RolesApi.md#updaterole) | **Put** /v1/roles/{roleId} | Update a role
*RolesAuditsApi* | [**RetrieveRoleAuditsList**](docs/RolesAuditsApi.md#retrieveroleauditslist) | **Get** /v1/roles/audits | Retrieve the audit history of your roles.
*SecretsApi* | [**CreateSecret**](docs/SecretsApi.md#createsecret) | **Post** /v1/secrets | Create a new secret
*SecretsApi* | [**DeleteSecret**](docs/SecretsApi.md#deletesecret) | **Delete** /v1/secrets/{secretId} | Delete existing secret by its identifier
*SecretsApi* | [**RetrieveSecret**](docs/SecretsApi.md#retrievesecret) | **Get** /v1/secrets/{secretId} | Get specific secret by id
*SecretsApi* | [**RetrieveSecretList**](docs/SecretsApi.md#retrievesecretlist) | **Get** /v1/secrets | Show secrets
*SecretsApi* | [**UpdateSecret**](docs/SecretsApi.md#updatesecret) | **Patch** /v1/secrets/{secretId} | Update specific secret by its identifier
*SecretsAuditsApi* | [**RetrieveSecretAuditsList**](docs/SecretsAuditsApi.md#retrievesecretauditslist) | **Get** /v1/secrets/audits | Display history about your secrets (audit)
*TagAssignmentsApi* | [**CreateAssignment**](docs/TagAssignmentsApi.md#createassignment) | **Post** /v1/tags/{tagId}/assignments/{resourceType}/{resourceId} | Create a new tag assignment
*TagAssignmentsApi* | [**DeleteAssignment**](docs/TagAssignmentsApi.md#deleteassignment) | **Delete** /v1/tags/{tagId}/assignments/{resourceType}/{resourceId} | Delete tag assignment
*TagAssignmentsApi* | [**RetrieveAssignment**](docs/TagAssignmentsApi.md#retrieveassignment) | **Get** /v1/tags/{tagId}/assignments/{resourceType}/{resourceId} | Get particular assignment for the tag
*TagAssignmentsApi* | [**RetrieveAssignmentList**](docs/TagAssignmentsApi.md#retrieveassignmentlist) | **Get** /v1/tags/{tagId}/assignments | List tag assignments
*TagAssignmentsAuditsApi* | [**RetrieveAssignmentsAuditsList**](docs/TagAssignmentsAuditsApi.md#retrieveassignmentsauditslist) | **Get** /v1/tags/assignments/audits | List a history about the assignments (audit)
*TagsApi* | [**CreateTag**](docs/TagsApi.md#createtag) | **Post** /v1/tags | Create a new tag
*TagsApi* | [**DeleteTag**](docs/TagsApi.md#deletetag) | **Delete** /v1/tags/{tagId} | Delete tag attributes using its id
*TagsApi* | [**RetrieveTag**](docs/TagsApi.md#retrievetag) | **Get** /v1/tags/{tagId} | Get tag attributes by its id
*TagsApi* | [**RetrieveTagList**](docs/TagsApi.md#retrievetaglist) | **Get** /v1/tags | List tags
*TagsApi* | [**UpdateTag**](docs/TagsApi.md#updatetag) | **Patch** /v1/tags/{tagId} | Update tag attributes using its id
*TagsAuditsApi* | [**RetrieveTagAuditsList**](docs/TagsAuditsApi.md#retrievetagauditslist) | **Get** /v1/tags/audits | List a history about the assignments (audit)
*UsersApi* | [**CreateUser**](docs/UsersApi.md#createuser) | **Post** /v1/users | Create User.
*UsersApi* | [**DeleteUser**](docs/UsersApi.md#deleteuser) | **Delete** /v1/users/{userId} | Delete User.
*UsersApi* | [**GenerateClientSecret**](docs/UsersApi.md#generateclientsecret) | **Put** /v1/users/client/secret | Regenerate Client Secret.
*UsersApi* | [**GetObjectStorageCredentials**](docs/UsersApi.md#getobjectstoragecredentials) | **Get** /v1/users/{userId}/object-storages/{objectStorageId}/credentials/{credentialId} | Get S3 compatible object storage credentials.
*UsersApi* | [**ListObjectStorageCredentials**](docs/UsersApi.md#listobjectstoragecredentials) | **Get** /v1/users/{userId}/object-storages/credentials | Retrieve a list of S3-compatible object storage credentials for a user.
*UsersApi* | [**RegenerateObjectStorageCredentials**](docs/UsersApi.md#regenerateobjectstoragecredentials) | **Patch** /v1/users/{userId}/object-storages/{objectStorageId}/credentials/{credentialId} | Regenerate secret key for Object Storage.
*UsersApi* | [**ResendEmailVerification**](docs/UsersApi.md#resendemailverification) | **Post** /v1/users/{userId}/resend-email-verification | Get verification email
*UsersApi* | [**ResetPassword**](docs/UsersApi.md#resetpassword) | **Post** /v1/users/{userId}/reset-password | Get email to reset password.
*UsersApi* | [**RetrieveUser**](docs/UsersApi.md#retrieveuser) | **Get** /v1/users/{userId} | Get specific User.
*UsersApi* | [**RetrieveUserClient**](docs/UsersApi.md#retrieveuserclient) | **Get** /v1/users/client | Retrieve Client.
*UsersApi* | [**RetrieveUserList**](docs/UsersApi.md#retrieveuserlist) | **Get** /v1/users | User List
*UsersApi* | [**UpdateUser**](docs/UsersApi.md#updateuser) | **Patch** /v1/users/{userId} | Update User.
*UsersAuditsApi* | [**RetrieveUserAuditsList**](docs/UsersAuditsApi.md#retrieveuserauditslist) | **Get** /v1/users/audits | Retrieve the audit history of your users.
*VirtualPrivateCloudAuditsApi* | [**RetrievePrivateNetworkAuditsList**](docs/VirtualPrivateCloudAuditsApi.md#retrieveprivatenetworkauditslist) | **Get** /v1/private-networks/audits | List history of your VPCs (audit)
*VirtualPrivateCloudVPCApi* | [**AssignInstancePrivateNetwork**](docs/VirtualPrivateCloudVPCApi.md#assigninstanceprivatenetwork) | **Post** /v1/private-networks/{privateNetworkId}/instances/{instanceId} | Add instance to a Virtual Private Cloud
*VirtualPrivateCloudVPCApi* | [**CreatePrivateNetwork**](docs/VirtualPrivateCloudVPCApi.md#createprivatenetwork) | **Post** /v1/private-networks | Create a new Virtual Private Cloud
*VirtualPrivateCloudVPCApi* | [**DeletePrivateNetwork**](docs/VirtualPrivateCloudVPCApi.md#deleteprivatenetwork) | **Delete** /v1/private-networks/{privateNetworkId} | Delete the virtual private cloud with the given id
*VirtualPrivateCloudVPCApi* | [**PatchPrivateNetwork**](docs/VirtualPrivateCloudVPCApi.md#patchprivatenetwork) | **Patch** /v1/private-networks/{privateNetworkId} | Update a Virtual Private Cloud by its id
*VirtualPrivateCloudVPCApi* | [**RetrievePrivateNetwork**](docs/VirtualPrivateCloudVPCApi.md#retrieveprivatenetwork) | **Get** /v1/private-networks/{privateNetworkId} | Get specific Virtual Private Cloud by its id
*VirtualPrivateCloudVPCApi* | [**RetrievePrivateNetworkList**](docs/VirtualPrivateCloudVPCApi.md#retrieveprivatenetworklist) | **Get** /v1/private-networks | List all Virtual Private Clouds
*VirtualPrivateCloudVPCApi* | [**UnassignInstancePrivateNetwork**](docs/VirtualPrivateCloudVPCApi.md#unassigninstanceprivatenetwork) | **Delete** /v1/private-networks/{privateNetworkId}/instances/{instanceId} | Remove particular instance from a Virtual Private Cloud


## Documentation For Models

 - [AddOnRequest](docs/AddOnRequest.md)
 - [AddOnResponse](docs/AddOnResponse.md)
 - [AdditionalIp](docs/AdditionalIp.md)
 - [ApiPermissionsResponse](docs/ApiPermissionsResponse.md)
 - [AssignInstancePrivateNetworkResponse](docs/AssignInstancePrivateNetworkResponse.md)
 - [AssignmentAuditResponse](docs/AssignmentAuditResponse.md)
 - [AssignmentResponse](docs/AssignmentResponse.md)
 - [AutoScalingTypeRequest](docs/AutoScalingTypeRequest.md)
 - [AutoScalingTypeResponse](docs/AutoScalingTypeResponse.md)
 - [CancelInstanceResponse](docs/CancelInstanceResponse.md)
 - [CancelInstanceResponseData](docs/CancelInstanceResponseData.md)
 - [CancelObjectStorageResponse](docs/CancelObjectStorageResponse.md)
 - [CancelObjectStorageResponseData](docs/CancelObjectStorageResponseData.md)
 - [ClientResponse](docs/ClientResponse.md)
 - [ClientSecretResponse](docs/ClientSecretResponse.md)
 - [CreateAssignmentResponse](docs/CreateAssignmentResponse.md)
 - [CreateCustomImageFailResponse](docs/CreateCustomImageFailResponse.md)
 - [CreateCustomImageRequest](docs/CreateCustomImageRequest.md)
 - [CreateCustomImageResponse](docs/CreateCustomImageResponse.md)
 - [CreateCustomImageResponseData](docs/CreateCustomImageResponseData.md)
 - [CreateInstanceAddons](docs/CreateInstanceAddons.md)
 - [CreateInstanceRequest](docs/CreateInstanceRequest.md)
 - [CreateInstanceResponse](docs/CreateInstanceResponse.md)
 - [CreateInstanceResponseData](docs/CreateInstanceResponseData.md)
 - [CreateObjectStorageRequest](docs/CreateObjectStorageRequest.md)
 - [CreateObjectStorageResponse](docs/CreateObjectStorageResponse.md)
 - [CreateObjectStorageResponseData](docs/CreateObjectStorageResponseData.md)
 - [CreatePrivateNetworkRequest](docs/CreatePrivateNetworkRequest.md)
 - [CreatePrivateNetworkResponse](docs/CreatePrivateNetworkResponse.md)
 - [CreateRoleRequest](docs/CreateRoleRequest.md)
 - [CreateRoleResponse](docs/CreateRoleResponse.md)
 - [CreateRoleResponseData](docs/CreateRoleResponseData.md)
 - [CreateSecretRequest](docs/CreateSecretRequest.md)
 - [CreateSecretResponse](docs/CreateSecretResponse.md)
 - [CreateSnapshotRequest](docs/CreateSnapshotRequest.md)
 - [CreateSnapshotResponse](docs/CreateSnapshotResponse.md)
 - [CreateTagRequest](docs/CreateTagRequest.md)
 - [CreateTagResponse](docs/CreateTagResponse.md)
 - [CreateTagResponseData](docs/CreateTagResponseData.md)
 - [CreateTicketRequest](docs/CreateTicketRequest.md)
 - [CreateTicketResponse](docs/CreateTicketResponse.md)
 - [CreateTicketResponseData](docs/CreateTicketResponseData.md)
 - [CreateUserRequest](docs/CreateUserRequest.md)
 - [CreateUserResponse](docs/CreateUserResponse.md)
 - [CreateUserResponseData](docs/CreateUserResponseData.md)
 - [CredentialData](docs/CredentialData.md)
 - [CustomImagesStatsResponse](docs/CustomImagesStatsResponse.md)
 - [CustomImagesStatsResponseData](docs/CustomImagesStatsResponseData.md)
 - [DataCenterResponse](docs/DataCenterResponse.md)
 - [ExtraStorageRequest](docs/ExtraStorageRequest.md)
 - [FindAssignmentResponse](docs/FindAssignmentResponse.md)
 - [FindClientResponse](docs/FindClientResponse.md)
 - [FindCredentialResponse](docs/FindCredentialResponse.md)
 - [FindImageResponse](docs/FindImageResponse.md)
 - [FindInstanceResponse](docs/FindInstanceResponse.md)
 - [FindObjectStorageResponse](docs/FindObjectStorageResponse.md)
 - [FindPrivateNetworkResponse](docs/FindPrivateNetworkResponse.md)
 - [FindRoleResponse](docs/FindRoleResponse.md)
 - [FindSecretResponse](docs/FindSecretResponse.md)
 - [FindSnapshotResponse](docs/FindSnapshotResponse.md)
 - [FindTagResponse](docs/FindTagResponse.md)
 - [FindUserIsPasswordSetResponse](docs/FindUserIsPasswordSetResponse.md)
 - [FindUserResponse](docs/FindUserResponse.md)
 - [FindVncResponse](docs/FindVncResponse.md)
 - [FirewallingUpgradeRequest](docs/FirewallingUpgradeRequest.md)
 - [GenerateClientSecretResponse](docs/GenerateClientSecretResponse.md)
 - [ImageAuditResponse](docs/ImageAuditResponse.md)
 - [ImageAuditResponseData](docs/ImageAuditResponseData.md)
 - [ImageResponse](docs/ImageResponse.md)
 - [InstanceAssignmentSelfLinks](docs/InstanceAssignmentSelfLinks.md)
 - [InstanceRescueActionResponse](docs/InstanceRescueActionResponse.md)
 - [InstanceRescueActionResponseData](docs/InstanceRescueActionResponseData.md)
 - [InstanceResetPasswordActionResponse](docs/InstanceResetPasswordActionResponse.md)
 - [InstanceResetPasswordActionResponseData](docs/InstanceResetPasswordActionResponseData.md)
 - [InstanceResponse](docs/InstanceResponse.md)
 - [InstanceRestartActionResponse](docs/InstanceRestartActionResponse.md)
 - [InstanceRestartActionResponseData](docs/InstanceRestartActionResponseData.md)
 - [InstanceShutdownActionResponse](docs/InstanceShutdownActionResponse.md)
 - [InstanceShutdownActionResponseData](docs/InstanceShutdownActionResponseData.md)
 - [InstanceStartActionResponse](docs/InstanceStartActionResponse.md)
 - [InstanceStartActionResponseData](docs/InstanceStartActionResponseData.md)
 - [InstanceStatus](docs/InstanceStatus.md)
 - [InstanceStopActionResponse](docs/InstanceStopActionResponse.md)
 - [InstanceStopActionResponseData](docs/InstanceStopActionResponseData.md)
 - [Instances](docs/Instances.md)
 - [InstancesActionsAuditResponse](docs/InstancesActionsAuditResponse.md)
 - [InstancesActionsRescueRequest](docs/InstancesActionsRescueRequest.md)
 - [InstancesAuditResponse](docs/InstancesAuditResponse.md)
 - [InstancesResetPasswordActionsRequest](docs/InstancesResetPasswordActionsRequest.md)
 - [IpConfig](docs/IpConfig.md)
 - [IpConfig1](docs/IpConfig1.md)
 - [IpV4](docs/IpV4.md)
 - [IpV41](docs/IpV41.md)
 - [IpV6](docs/IpV6.md)
 - [IpV61](docs/IpV61.md)
 - [Links](docs/Links.md)
 - [ListApiPermissionResponse](docs/ListApiPermissionResponse.md)
 - [ListAssignmentAuditsResponse](docs/ListAssignmentAuditsResponse.md)
 - [ListAssignmentResponse](docs/ListAssignmentResponse.md)
 - [ListCredentialResponse](docs/ListCredentialResponse.md)
 - [ListDataCenterResponse](docs/ListDataCenterResponse.md)
 - [ListImageResponse](docs/ListImageResponse.md)
 - [ListImageResponseData](docs/ListImageResponseData.md)
 - [ListInstancesActionsAuditResponse](docs/ListInstancesActionsAuditResponse.md)
 - [ListInstancesAuditResponse](docs/ListInstancesAuditResponse.md)
 - [ListInstancesResponse](docs/ListInstancesResponse.md)
 - [ListInstancesResponseData](docs/ListInstancesResponseData.md)
 - [ListObjectStorageAuditResponse](docs/ListObjectStorageAuditResponse.md)
 - [ListObjectStorageResponse](docs/ListObjectStorageResponse.md)
 - [ListPrivateNetworkAuditResponse](docs/ListPrivateNetworkAuditResponse.md)
 - [ListPrivateNetworkResponse](docs/ListPrivateNetworkResponse.md)
 - [ListPrivateNetworkResponseData](docs/ListPrivateNetworkResponseData.md)
 - [ListRoleAuditResponse](docs/ListRoleAuditResponse.md)
 - [ListRoleResponse](docs/ListRoleResponse.md)
 - [ListSecretAuditResponse](docs/ListSecretAuditResponse.md)
 - [ListSecretResponse](docs/ListSecretResponse.md)
 - [ListSnapshotResponse](docs/ListSnapshotResponse.md)
 - [ListSnapshotsAuditResponse](docs/ListSnapshotsAuditResponse.md)
 - [ListTagAuditsResponse](docs/ListTagAuditsResponse.md)
 - [ListTagResponse](docs/ListTagResponse.md)
 - [ListUserAuditResponse](docs/ListUserAuditResponse.md)
 - [ListUserResponse](docs/ListUserResponse.md)
 - [ObjectStorageAuditResponse](docs/ObjectStorageAuditResponse.md)
 - [ObjectStorageResponse](docs/ObjectStorageResponse.md)
 - [ObjectStoragesStatsResponse](docs/ObjectStoragesStatsResponse.md)
 - [ObjectStoragesStatsResponseData](docs/ObjectStoragesStatsResponseData.md)
 - [PaginationMeta](docs/PaginationMeta.md)
 - [PatchInstanceRequest](docs/PatchInstanceRequest.md)
 - [PatchInstanceResponse](docs/PatchInstanceResponse.md)
 - [PatchInstanceResponseData](docs/PatchInstanceResponseData.md)
 - [PatchObjectStorageRequest](docs/PatchObjectStorageRequest.md)
 - [PatchPrivateNetworkRequest](docs/PatchPrivateNetworkRequest.md)
 - [PatchPrivateNetworkResponse](docs/PatchPrivateNetworkResponse.md)
 - [PatchVncRequest](docs/PatchVncRequest.md)
 - [PermissionRequest](docs/PermissionRequest.md)
 - [PermissionResponse](docs/PermissionResponse.md)
 - [PrivateIpConfig](docs/PrivateIpConfig.md)
 - [PrivateNetworkAuditResponse](docs/PrivateNetworkAuditResponse.md)
 - [PrivateNetworkResponse](docs/PrivateNetworkResponse.md)
 - [ReinstallInstanceRequest](docs/ReinstallInstanceRequest.md)
 - [ReinstallInstanceResponse](docs/ReinstallInstanceResponse.md)
 - [ReinstallInstanceResponseData](docs/ReinstallInstanceResponseData.md)
 - [ResourcePermissionsResponse](docs/ResourcePermissionsResponse.md)
 - [RoleAuditResponse](docs/RoleAuditResponse.md)
 - [RoleResponse](docs/RoleResponse.md)
 - [RollbackSnapshotResponse](docs/RollbackSnapshotResponse.md)
 - [SecretAuditResponse](docs/SecretAuditResponse.md)
 - [SecretResponse](docs/SecretResponse.md)
 - [SelfLinks](docs/SelfLinks.md)
 - [SnapshotResponse](docs/SnapshotResponse.md)
 - [SnapshotsAuditResponse](docs/SnapshotsAuditResponse.md)
 - [TagAssignmentSelfLinks](docs/TagAssignmentSelfLinks.md)
 - [TagAuditResponse](docs/TagAuditResponse.md)
 - [TagResponse](docs/TagResponse.md)
 - [TagResponse1](docs/TagResponse1.md)
 - [UnassignInstancePrivateNetworkResponse](docs/UnassignInstancePrivateNetworkResponse.md)
 - [UpdateCustomImageRequest](docs/UpdateCustomImageRequest.md)
 - [UpdateCustomImageResponse](docs/UpdateCustomImageResponse.md)
 - [UpdateCustomImageResponseData](docs/UpdateCustomImageResponseData.md)
 - [UpdateRoleRequest](docs/UpdateRoleRequest.md)
 - [UpdateRoleResponse](docs/UpdateRoleResponse.md)
 - [UpdateSecretRequest](docs/UpdateSecretRequest.md)
 - [UpdateSecretResponse](docs/UpdateSecretResponse.md)
 - [UpdateSnapshotRequest](docs/UpdateSnapshotRequest.md)
 - [UpdateSnapshotResponse](docs/UpdateSnapshotResponse.md)
 - [UpdateTagRequest](docs/UpdateTagRequest.md)
 - [UpdateTagResponse](docs/UpdateTagResponse.md)
 - [UpdateUserRequest](docs/UpdateUserRequest.md)
 - [UpdateUserResponse](docs/UpdateUserResponse.md)
 - [UpgradeAutoScalingType](docs/UpgradeAutoScalingType.md)
 - [UpgradeInstanceRequest](docs/UpgradeInstanceRequest.md)
 - [UpgradeObjectStorageRequest](docs/UpgradeObjectStorageRequest.md)
 - [UpgradeObjectStorageResponse](docs/UpgradeObjectStorageResponse.md)
 - [UpgradeObjectStorageResponseData](docs/UpgradeObjectStorageResponseData.md)
 - [UserAuditResponse](docs/UserAuditResponse.md)
 - [UserIsPasswordSetResponse](docs/UserIsPasswordSetResponse.md)
 - [UserResponse](docs/UserResponse.md)
 - [VncResponse](docs/VncResponse.md)


## Documentation For Authorization



### bearer

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARERTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@fybe.com

