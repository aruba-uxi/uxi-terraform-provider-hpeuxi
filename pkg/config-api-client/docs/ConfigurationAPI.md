# \ConfigurationAPI

All URIs are relative to *https://api.capenetworks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigurationAppV1SensorGroupAssignmentsGet**](ConfigurationAPI.md#GetConfigurationAppV1SensorGroupAssignmentsGet) | **Get** /configuration/app/v1/sensor-group-assignments | Get
[**GetConfigurationAppV1WiredNetworksGet**](ConfigurationAPI.md#GetConfigurationAppV1WiredNetworksGet) | **Get** /configuration/app/v1/wired-networks | Get
[**GetConfigurationAppV1WirelessNetworksGet**](ConfigurationAPI.md#GetConfigurationAppV1WirelessNetworksGet) | **Get** /configuration/app/v1/wireless-networks | Get
[**GroupsGetConfigurationAppV1GroupsGet**](ConfigurationAPI.md#GroupsGetConfigurationAppV1GroupsGet) | **Get** /configuration/app/v1/groups | Groups Get
[**GroupsPostConfigurationAppV1GroupsPost**](ConfigurationAPI.md#GroupsPostConfigurationAppV1GroupsPost) | **Post** /configuration/app/v1/groups | Groups Post



## GetConfigurationAppV1SensorGroupAssignmentsGet

> SensorGroupAssignmentsResponse GetConfigurationAppV1SensorGroupAssignmentsGet(ctx).Uid(uid).Cursor(cursor).Limit(limit).Execute()

Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
)

func main() {
	uid := "uid_example" // string |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GetConfigurationAppV1SensorGroupAssignmentsGet(context.Background()).Uid(uid).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetConfigurationAppV1SensorGroupAssignmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationAppV1SensorGroupAssignmentsGet`: SensorGroupAssignmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetConfigurationAppV1SensorGroupAssignmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationAppV1SensorGroupAssignmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**SensorGroupAssignmentsResponse**](SensorGroupAssignmentsResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationAppV1WiredNetworksGet

> WiredNetworksResponse GetConfigurationAppV1WiredNetworksGet(ctx).Uid(uid).Cursor(cursor).Limit(limit).Execute()

Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
)

func main() {
	uid := "uid_example" // string |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GetConfigurationAppV1WiredNetworksGet(context.Background()).Uid(uid).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetConfigurationAppV1WiredNetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationAppV1WiredNetworksGet`: WiredNetworksResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetConfigurationAppV1WiredNetworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationAppV1WiredNetworksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**WiredNetworksResponse**](WiredNetworksResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationAppV1WirelessNetworksGet

> WirelessNetworksResponse GetConfigurationAppV1WirelessNetworksGet(ctx).Uid(uid).Cursor(cursor).Limit(limit).Execute()

Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
)

func main() {
	uid := "uid_example" // string |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GetConfigurationAppV1WirelessNetworksGet(context.Background()).Uid(uid).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetConfigurationAppV1WirelessNetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationAppV1WirelessNetworksGet`: WirelessNetworksResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetConfigurationAppV1WirelessNetworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationAppV1WirelessNetworksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**WirelessNetworksResponse**](WirelessNetworksResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGetConfigurationAppV1GroupsGet

> GroupsGetResponse GroupsGetConfigurationAppV1GroupsGet(ctx).Uid(uid).Cursor(cursor).Limit(limit).Execute()

Groups Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
)

func main() {
	uid := "uid_example" // string |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GroupsGetConfigurationAppV1GroupsGet(context.Background()).Uid(uid).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GroupsGetConfigurationAppV1GroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGetConfigurationAppV1GroupsGet`: GroupsGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GroupsGetConfigurationAppV1GroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetConfigurationAppV1GroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**GroupsGetResponse**](GroupsGetResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsPostConfigurationAppV1GroupsPost

> GroupsPostResponse GroupsPostConfigurationAppV1GroupsPost(ctx).GroupsPostRequest(groupsPostRequest).Execute()

Groups Post



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
)

func main() {
	groupsPostRequest := *openapiclient.NewGroupsPostRequest("ParentUid_example", "Name_example") // GroupsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GroupsPostConfigurationAppV1GroupsPost(context.Background()).GroupsPostRequest(groupsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GroupsPostConfigurationAppV1GroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsPostConfigurationAppV1GroupsPost`: GroupsPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GroupsPostConfigurationAppV1GroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsPostConfigurationAppV1GroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupsPostRequest** | [**GroupsPostRequest**](GroupsPostRequest.md) |  | 

### Return type

[**GroupsPostResponse**](GroupsPostResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

