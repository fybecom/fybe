/*
Fybe API Reference

# Introduction  The Fybe API facilitates resource management through HTTP requests. This documentation comprises a collection of HTTP endpoints that adhere to RESTful principles. Each endpoint is accompanied by descriptions, as well as request and response schemas.  ## OpenAPI specification (OAS)  Fybe's OpenAPI Specification can be [downloaded here](https://api.fybe.com/api-v1.yaml).  ## Getting Started  To utilize the Fybe API, you'll require following credentials that can be obtained from from the [Fybe Cockpit](https://cockpit.fybe.com/account/security): 1. ClientId 2. ClientSecret 3. API User (your email address to login to the [Fybe Cockpit](https://cockpit.fybe.com/account/security)) 4. API Password (this is a new password which you'll set or change in the [Fybe Cockpit](https://cockpit.fybe.com/account/security))  ### How to use the API?  As authentication [Bearer Tokens](https://oauth.net/2/bearer-tokens/) in form of [JWT](https://www.rfc-editor.org/rfc/rfc7519) are used. This approach follows the [OAuth 2.0](https://oauth.net/2/) specification.  #### Retrieve the Bearer Token  ```sh POST /auth/realms/Fybe/protocol/openid-connect/token HTTP/1.1 Host: airlock.fybe.com  grant_type=password &password=xxxxxx &redirect_uri=https://example-app.com/redirect &username=xxxxxx &client_id=xxxxxx &client_secret=xxxxxx ```  The actual values for `client_id`, `client_secret`, `username` and `password` can be retrieved from [Fybe Cockpit](https://cockpit.fybe.com/account/security)  ## Using the Command-Line Interface (CLI)  Fybe provides the CLI called `fybe` which can be downloaded from <https://github.com/fybecom/fybe>. Please follow the instructions in the `README.md` to perform the installation on your OS. `fybe` supports Windows, macOS and Linux operating systems.  Using `fybe` CLI invoking makes invoking the API much more comfortable. E.g. `fybe get instances` for retrieving the list of compute instances.  ## Requests  As stated below, the Fybe API accommodates HTTP requests. However, not all endpoints endorse every method. You can find a list of allowed methods within this documentation.  Method | Description ---    | --- GET    | When you need to obtain information about a resource, you should utilize the GET method. The data will be provided as a JSON object. It's essential to note that GET methods are read-only and don't impact any resources. POST   | To create a new object, send a POST method that encodes all necessary attributes in the request body as JSON. PATCH  | PATCH can be used to partially modify a resource, allowing specific attributes to be changed without updating the complete object. PUT    | If you need to update information about a resource, use the PUT method. PUT will overwrite existing values of the item, without considering the current state. DELETE | To remove a resource from your account, apply the DELETE method. If the resource is not found, the operation will generate a 4xx error along with a relevant message.  ## Responses  Typically, the Fybe API will respond to your requests by returning data in [JSON](https://www.json.org/) format, which can be easily processed using any programming language or tools.  Like for any HTTP requests, you'll receive an HTTP status code, which provides general information about the success or error of the request. The table below outlines some common HTTP status codes.  It's important to note that the description of the endpoints and methods doesn't provide an exhaustive list of all possible status codes. Only specific return codes and their corresponding response data are explicitly outlined.  Response Code | Description --- | --- 200 | The response will contain the information you requested. 201 | Your request has been acknowledged, and the resource has been created. 204 | Your request was successful, and no further information is provided in the response. 400 | Your request was not properly formed. 401 | You didn't provide valid authentication credentials. 402 | The request was declined as it necessitates an additional paid service. 403 | You are prohibited from performing the request. You'll need to pass a bearer token. 404 | No results were found for your request, or the resource you're trying to access does not exist. 409 | There is a conflict with resources, such as a violation of unique data constraints when attempting to create or modify resources. 429 | The rate limit has been reached. Please wait for some time before making further requests. 500 | We couldn't fulfill your request due to server-side issues. If this happens, please try again later or reach out to our support team.  Not every endpoint returns data. For example DELETE requests usually don't return any data. All others do return data. For easy handling the return values consists of metadata denoted with and underscore (\"_\") like `_links` or `_pagination`. The actual data is returned in a field called `data`. For convenience reasons this `data` field is always returned as an array even if it consists of only one single element.  DELETE requests usually don't return any data. Return values consists of metadata starting with an underscore (\"_\"), e.g. `_links` and `_pagination`. The requested data is to be found in the field `data`. The `data` field is always an array.   Please visit [Fybe](https://fybe.com) for more general information. 

API version: 1.0.0
Contact: support@fybe.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ListInstancesResponseData struct for ListInstancesResponseData
type ListInstancesResponseData struct {
	// Your customer tenant id
	TenantId string `json:"tenantId"`
	// Customer ID
	CustomerId string `json:"customerId"`
	AdditionalIps []AdditionalIp `json:"additionalIps"`
	// Name of the compute instance
	Name string `json:"name"`
	// The customer name by which the compute instance is represented
	DisplayName string `json:"displayName"`
	// Identifier of the compute instance
	InstanceId int64 `json:"instanceId"`
	// The data center where your Private Network is located
	DataCenter string `json:"dataCenter"`
	// Instance region where the compute instance should be located.
	Region string `json:"region"`
	// The name of the region where the instance is located.
	RegionName string `json:"regionName"`
	// Product ID
	ProductId string `json:"productId"`
	// Image's id
	ImageId string `json:"imageId"`
	IpConfig IpConfig1 `json:"ipConfig"`
	// MAC Address
	MacAddress string `json:"macAddress"`
	// Image RAM size in MB
	RamMb float32 `json:"ramMb"`
	// CPU core count
	CpuCores int64 `json:"cpuCores"`
	// Type of operating system (OS)
	OsType string `json:"osType"`
	// Image Disk size in MB
	DiskMb float32 `json:"diskMb"`
	// Array of `secretId`s of public SSH keys for logging into as `defaultUser` with administrator/root privileges. Applies to Linux/BSD systems. Please refer to Secrets Management API.
	SshKeys []int64 `json:"sshKeys"`
	// The date on which the compute instance was created
	CreatedDate time.Time `json:"createdDate"`
	// The date when the compute instance is to be terminated
	CancelDate string `json:"cancelDate"`
	Status InstanceStatus `json:"status"`
	// ID of host system
	VHostId int64 `json:"vHostId"`
	AddOns []AddOnResponse `json:"addOns"`
	// Message in case of an error.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// The category to which the instance belongs based on its ProductId.
	ProductType string `json:"productType"`
	// Default user name created for login during (re-)installation with administrative privileges. Allowed values for Linux/BSD are `admin` (use sudo to apply administrative privileges like root) or `root`. Allowed values for Windows are `admin` (has administrative privileges like administrator) or `administrator`.
	DefaultUser *string `json:"defaultUser,omitempty"`
}

// NewListInstancesResponseData instantiates a new ListInstancesResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInstancesResponseData(tenantId string, customerId string, additionalIps []AdditionalIp, name string, displayName string, instanceId int64, dataCenter string, region string, regionName string, productId string, imageId string, ipConfig IpConfig1, macAddress string, ramMb float32, cpuCores int64, osType string, diskMb float32, sshKeys []int64, createdDate time.Time, cancelDate string, status InstanceStatus, vHostId int64, addOns []AddOnResponse, productType string) *ListInstancesResponseData {
	this := ListInstancesResponseData{}
	this.TenantId = tenantId
	this.CustomerId = customerId
	this.AdditionalIps = additionalIps
	this.Name = name
	this.DisplayName = displayName
	this.InstanceId = instanceId
	this.DataCenter = dataCenter
	this.Region = region
	this.RegionName = regionName
	this.ProductId = productId
	this.ImageId = imageId
	this.IpConfig = ipConfig
	this.MacAddress = macAddress
	this.RamMb = ramMb
	this.CpuCores = cpuCores
	this.OsType = osType
	this.DiskMb = diskMb
	this.SshKeys = sshKeys
	this.CreatedDate = createdDate
	this.CancelDate = cancelDate
	this.Status = status
	this.VHostId = vHostId
	this.AddOns = addOns
	this.ProductType = productType
	return &this
}

// NewListInstancesResponseDataWithDefaults instantiates a new ListInstancesResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInstancesResponseDataWithDefaults() *ListInstancesResponseData {
	this := ListInstancesResponseData{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *ListInstancesResponseData) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ListInstancesResponseData) SetTenantId(v string) {
	o.TenantId = v
}

// GetCustomerId returns the CustomerId field value
func (o *ListInstancesResponseData) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *ListInstancesResponseData) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetAdditionalIps returns the AdditionalIps field value
func (o *ListInstancesResponseData) GetAdditionalIps() []AdditionalIp {
	if o == nil {
		var ret []AdditionalIp
		return ret
	}

	return o.AdditionalIps
}

// GetAdditionalIpsOk returns a tuple with the AdditionalIps field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetAdditionalIpsOk() (*[]AdditionalIp, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AdditionalIps, true
}

// SetAdditionalIps sets field value
func (o *ListInstancesResponseData) SetAdditionalIps(v []AdditionalIp) {
	o.AdditionalIps = v
}

// GetName returns the Name field value
func (o *ListInstancesResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListInstancesResponseData) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *ListInstancesResponseData) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ListInstancesResponseData) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetInstanceId returns the InstanceId field value
func (o *ListInstancesResponseData) GetInstanceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetInstanceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *ListInstancesResponseData) SetInstanceId(v int64) {
	o.InstanceId = v
}

// GetDataCenter returns the DataCenter field value
func (o *ListInstancesResponseData) GetDataCenter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataCenter
}

// GetDataCenterOk returns a tuple with the DataCenter field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetDataCenterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DataCenter, true
}

// SetDataCenter sets field value
func (o *ListInstancesResponseData) SetDataCenter(v string) {
	o.DataCenter = v
}

// GetRegion returns the Region field value
func (o *ListInstancesResponseData) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *ListInstancesResponseData) SetRegion(v string) {
	o.Region = v
}

// GetRegionName returns the RegionName field value
func (o *ListInstancesResponseData) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetRegionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *ListInstancesResponseData) SetRegionName(v string) {
	o.RegionName = v
}

// GetProductId returns the ProductId field value
func (o *ListInstancesResponseData) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetProductIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *ListInstancesResponseData) SetProductId(v string) {
	o.ProductId = v
}

// GetImageId returns the ImageId field value
func (o *ListInstancesResponseData) GetImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetImageIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImageId, true
}

// SetImageId sets field value
func (o *ListInstancesResponseData) SetImageId(v string) {
	o.ImageId = v
}

// GetIpConfig returns the IpConfig field value
func (o *ListInstancesResponseData) GetIpConfig() IpConfig1 {
	if o == nil {
		var ret IpConfig1
		return ret
	}

	return o.IpConfig
}

// GetIpConfigOk returns a tuple with the IpConfig field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetIpConfigOk() (*IpConfig1, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IpConfig, true
}

// SetIpConfig sets field value
func (o *ListInstancesResponseData) SetIpConfig(v IpConfig1) {
	o.IpConfig = v
}

// GetMacAddress returns the MacAddress field value
func (o *ListInstancesResponseData) GetMacAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetMacAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MacAddress, true
}

// SetMacAddress sets field value
func (o *ListInstancesResponseData) SetMacAddress(v string) {
	o.MacAddress = v
}

// GetRamMb returns the RamMb field value
func (o *ListInstancesResponseData) GetRamMb() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RamMb
}

// GetRamMbOk returns a tuple with the RamMb field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetRamMbOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RamMb, true
}

// SetRamMb sets field value
func (o *ListInstancesResponseData) SetRamMb(v float32) {
	o.RamMb = v
}

// GetCpuCores returns the CpuCores field value
func (o *ListInstancesResponseData) GetCpuCores() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CpuCores
}

// GetCpuCoresOk returns a tuple with the CpuCores field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetCpuCoresOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CpuCores, true
}

// SetCpuCores sets field value
func (o *ListInstancesResponseData) SetCpuCores(v int64) {
	o.CpuCores = v
}

// GetOsType returns the OsType field value
func (o *ListInstancesResponseData) GetOsType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetOsTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OsType, true
}

// SetOsType sets field value
func (o *ListInstancesResponseData) SetOsType(v string) {
	o.OsType = v
}

// GetDiskMb returns the DiskMb field value
func (o *ListInstancesResponseData) GetDiskMb() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DiskMb
}

// GetDiskMbOk returns a tuple with the DiskMb field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetDiskMbOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DiskMb, true
}

// SetDiskMb sets field value
func (o *ListInstancesResponseData) SetDiskMb(v float32) {
	o.DiskMb = v
}

// GetSshKeys returns the SshKeys field value
func (o *ListInstancesResponseData) GetSshKeys() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetSshKeysOk() (*[]int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SshKeys, true
}

// SetSshKeys sets field value
func (o *ListInstancesResponseData) SetSshKeys(v []int64) {
	o.SshKeys = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *ListInstancesResponseData) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *ListInstancesResponseData) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetCancelDate returns the CancelDate field value
func (o *ListInstancesResponseData) GetCancelDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CancelDate
}

// GetCancelDateOk returns a tuple with the CancelDate field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetCancelDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CancelDate, true
}

// SetCancelDate sets field value
func (o *ListInstancesResponseData) SetCancelDate(v string) {
	o.CancelDate = v
}

// GetStatus returns the Status field value
func (o *ListInstancesResponseData) GetStatus() InstanceStatus {
	if o == nil {
		var ret InstanceStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetStatusOk() (*InstanceStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ListInstancesResponseData) SetStatus(v InstanceStatus) {
	o.Status = v
}

// GetVHostId returns the VHostId field value
func (o *ListInstancesResponseData) GetVHostId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VHostId
}

// GetVHostIdOk returns a tuple with the VHostId field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetVHostIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VHostId, true
}

// SetVHostId sets field value
func (o *ListInstancesResponseData) SetVHostId(v int64) {
	o.VHostId = v
}

// GetAddOns returns the AddOns field value
func (o *ListInstancesResponseData) GetAddOns() []AddOnResponse {
	if o == nil {
		var ret []AddOnResponse
		return ret
	}

	return o.AddOns
}

// GetAddOnsOk returns a tuple with the AddOns field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetAddOnsOk() (*[]AddOnResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AddOns, true
}

// SetAddOns sets field value
func (o *ListInstancesResponseData) SetAddOns(v []AddOnResponse) {
	o.AddOns = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ListInstancesResponseData) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ListInstancesResponseData) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ListInstancesResponseData) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetProductType returns the ProductType field value
func (o *ListInstancesResponseData) GetProductType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetProductTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *ListInstancesResponseData) SetProductType(v string) {
	o.ProductType = v
}

// GetDefaultUser returns the DefaultUser field value if set, zero value otherwise.
func (o *ListInstancesResponseData) GetDefaultUser() string {
	if o == nil || o.DefaultUser == nil {
		var ret string
		return ret
	}
	return *o.DefaultUser
}

// GetDefaultUserOk returns a tuple with the DefaultUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetDefaultUserOk() (*string, bool) {
	if o == nil || o.DefaultUser == nil {
		return nil, false
	}
	return o.DefaultUser, true
}

// HasDefaultUser returns a boolean if a field has been set.
func (o *ListInstancesResponseData) HasDefaultUser() bool {
	if o != nil && o.DefaultUser != nil {
		return true
	}

	return false
}

// SetDefaultUser gets a reference to the given string and assigns it to the DefaultUser field.
func (o *ListInstancesResponseData) SetDefaultUser(v string) {
	o.DefaultUser = &v
}

func (o ListInstancesResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["customerId"] = o.CustomerId
	}
	if true {
		toSerialize["additionalIps"] = o.AdditionalIps
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["instanceId"] = o.InstanceId
	}
	if true {
		toSerialize["dataCenter"] = o.DataCenter
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["regionName"] = o.RegionName
	}
	if true {
		toSerialize["productId"] = o.ProductId
	}
	if true {
		toSerialize["imageId"] = o.ImageId
	}
	if true {
		toSerialize["ipConfig"] = o.IpConfig
	}
	if true {
		toSerialize["macAddress"] = o.MacAddress
	}
	if true {
		toSerialize["ramMb"] = o.RamMb
	}
	if true {
		toSerialize["cpuCores"] = o.CpuCores
	}
	if true {
		toSerialize["osType"] = o.OsType
	}
	if true {
		toSerialize["diskMb"] = o.DiskMb
	}
	if true {
		toSerialize["sshKeys"] = o.SshKeys
	}
	if true {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if true {
		toSerialize["cancelDate"] = o.CancelDate
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["vHostId"] = o.VHostId
	}
	if true {
		toSerialize["addOns"] = o.AddOns
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if true {
		toSerialize["productType"] = o.ProductType
	}
	if o.DefaultUser != nil {
		toSerialize["defaultUser"] = o.DefaultUser
	}
	return json.Marshal(toSerialize)
}

type NullableListInstancesResponseData struct {
	value *ListInstancesResponseData
	isSet bool
}

func (v NullableListInstancesResponseData) Get() *ListInstancesResponseData {
	return v.value
}

func (v *NullableListInstancesResponseData) Set(val *ListInstancesResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableListInstancesResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableListInstancesResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInstancesResponseData(val *ListInstancesResponseData) *NullableListInstancesResponseData {
	return &NullableListInstancesResponseData{value: val, isSet: true}
}

func (v NullableListInstancesResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInstancesResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


