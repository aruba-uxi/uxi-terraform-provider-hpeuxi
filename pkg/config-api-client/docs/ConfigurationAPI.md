# \ConfigurationAPI

All URIs are relative to *https://api.capenetworks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUxiV1alpha1NetworkGroupAssignmentsGet**](ConfigurationAPI.md#GetUxiV1alpha1NetworkGroupAssignmentsGet) | **Get** /uxi/v1alpha1/network-group-assignments | Get
[**GetUxiV1alpha1SensorGroupAssignmentsGet**](ConfigurationAPI.md#GetUxiV1alpha1SensorGroupAssignmentsGet) | **Get** /uxi/v1alpha1/sensor-group-assignments | Get
[**GetUxiV1alpha1WiredNetworksGet**](ConfigurationAPI.md#GetUxiV1alpha1WiredNetworksGet) | **Get** /uxi/v1alpha1/wired-networks | Get
[**GetUxiV1alpha1WirelessNetworksGet**](ConfigurationAPI.md#GetUxiV1alpha1WirelessNetworksGet) | **Get** /uxi/v1alpha1/wireless-networks | Get
[**GroupsGetUxiV1alpha1GroupsGet**](ConfigurationAPI.md#GroupsGetUxiV1alpha1GroupsGet) | **Get** /uxi/v1alpha1/groups | Groups Get
[**GroupsPostUxiV1alpha1GroupsPost**](ConfigurationAPI.md#GroupsPostUxiV1alpha1GroupsPost) | **Post** /uxi/v1alpha1/groups | Groups Post
[**PostUxiV1alpha1SensorGroupAssignmentsPost**](ConfigurationAPI.md#PostUxiV1alpha1SensorGroupAssignmentsPost) | **Post** /uxi/v1alpha1/sensor-group-assignments | Post



## GetUxiV1alpha1NetworkGroupAssignmentsGet

> NetworkGroupAssignmentsResponse GetUxiV1alpha1NetworkGroupAssignmentsGet(ctx).Uid(uid).Cursor(cursor).Limit(limit).Execute()

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
	resp, r, err := apiClient.ConfigurationAPI.GetUxiV1alpha1NetworkGroupAssignmentsGet(context.Background()).Uid(uid).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetUxiV1alpha1NetworkGroupAssignmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUxiV1alpha1NetworkGroupAssignmentsGet`: NetworkGroupAssignmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetUxiV1alpha1NetworkGroupAssignmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUxiV1alpha1NetworkGroupAssignmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**NetworkGroupAssignmentsResponse**](NetworkGroupAssignmentsResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUxiV1alpha1SensorGroupAssignmentsGet

> SensorGroupAssignmentsResponse GetUxiV1alpha1SensorGroupAssignmentsGet(ctx).Uid(uid).Cursor(cursor).Limit(limit).Execute()

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
	resp, r, err := apiClient.ConfigurationAPI.GetUxiV1alpha1SensorGroupAssignmentsGet(context.Background()).Uid(uid).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetUxiV1alpha1SensorGroupAssignmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUxiV1alpha1SensorGroupAssignmentsGet`: SensorGroupAssignmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetUxiV1alpha1SensorGroupAssignmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUxiV1alpha1SensorGroupAssignmentsGetRequest struct via the builder pattern


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


## GetUxiV1alpha1WiredNetworksGet

> WiredNetworksResponse GetUxiV1alpha1WiredNetworksGet(ctx).Uid(uid).Cursor(cursor).Limit(limit).Execute()

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
	resp, r, err := apiClient.ConfigurationAPI.GetUxiV1alpha1WiredNetworksGet(context.Background()).Uid(uid).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetUxiV1alpha1WiredNetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUxiV1alpha1WiredNetworksGet`: WiredNetworksResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetUxiV1alpha1WiredNetworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUxiV1alpha1WiredNetworksGetRequest struct via the builder pattern


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


## GetUxiV1alpha1WirelessNetworksGet

> WirelessNetworksResponse GetUxiV1alpha1WirelessNetworksGet(ctx).Uid(uid).Cursor(cursor).Limit(limit).Execute()

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
	resp, r, err := apiClient.ConfigurationAPI.GetUxiV1alpha1WirelessNetworksGet(context.Background()).Uid(uid).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetUxiV1alpha1WirelessNetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUxiV1alpha1WirelessNetworksGet`: WirelessNetworksResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetUxiV1alpha1WirelessNetworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUxiV1alpha1WirelessNetworksGetRequest struct via the builder pattern


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


## GroupsGetUxiV1alpha1GroupsGet

> GroupsGetResponse GroupsGetUxiV1alpha1GroupsGet(ctx).Uid(uid).Cursor(cursor).Limit(limit).Execute()

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
	resp, r, err := apiClient.ConfigurationAPI.GroupsGetUxiV1alpha1GroupsGet(context.Background()).Uid(uid).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GroupsGetUxiV1alpha1GroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGetUxiV1alpha1GroupsGet`: GroupsGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GroupsGetUxiV1alpha1GroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetUxiV1alpha1GroupsGetRequest struct via the builder pattern


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


## GroupsPostUxiV1alpha1GroupsPost

> GroupsPostResponse GroupsPostUxiV1alpha1GroupsPost(ctx).GroupsPostRequest(groupsPostRequest).Execute()

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
	groupsPostRequest := *openapiclient.NewGroupsPostRequest("ParentId_example", "Name_example") // GroupsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GroupsPostUxiV1alpha1GroupsPost(context.Background()).GroupsPostRequest(groupsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GroupsPostUxiV1alpha1GroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsPostUxiV1alpha1GroupsPost`: GroupsPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GroupsPostUxiV1alpha1GroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsPostUxiV1alpha1GroupsPostRequest struct via the builder pattern


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


## PostUxiV1alpha1SensorGroupAssignmentsPost

> SensorGroupAssignmentResponse PostUxiV1alpha1SensorGroupAssignmentsPost(ctx).SensorGroupAssignmentsPostRequest(sensorGroupAssignmentsPostRequest).Execute()

Post



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
	sensorGroupAssignmentsPostRequest := *openapiclient.NewSensorGroupAssignmentsPostRequest("GroupId_example", "SensorId_example") // SensorGroupAssignmentsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.PostUxiV1alpha1SensorGroupAssignmentsPost(context.Background()).SensorGroupAssignmentsPostRequest(sensorGroupAssignmentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.PostUxiV1alpha1SensorGroupAssignmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUxiV1alpha1SensorGroupAssignmentsPost`: SensorGroupAssignmentResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.PostUxiV1alpha1SensorGroupAssignmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUxiV1alpha1SensorGroupAssignmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sensorGroupAssignmentsPostRequest** | [**SensorGroupAssignmentsPostRequest**](SensorGroupAssignmentsPostRequest.md) |  | 

### Return type

[**SensorGroupAssignmentResponse**](SensorGroupAssignmentResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

