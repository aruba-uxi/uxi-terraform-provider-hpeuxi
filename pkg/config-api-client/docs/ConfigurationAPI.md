# \ConfigurationAPI

All URIs are relative to *https://api.capenetworks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNetworkGroupAssignmentUxiV1alpha1NetworkGroupAssignmentsIdDelete**](ConfigurationAPI.md#DeleteNetworkGroupAssignmentUxiV1alpha1NetworkGroupAssignmentsIdDelete) | **Delete** /uxi/v1alpha1/network-group-assignments/{id} | Delete Network Group Assignment
[**DeleteSensorGroupAssignmentUxiV1alpha1SensorGroupAssignmentsIdDelete**](ConfigurationAPI.md#DeleteSensorGroupAssignmentUxiV1alpha1SensorGroupAssignmentsIdDelete) | **Delete** /uxi/v1alpha1/sensor-group-assignments/{id} | Delete Sensor Group Assignment
[**GetUxiV1alpha1NetworkGroupAssignmentsGet**](ConfigurationAPI.md#GetUxiV1alpha1NetworkGroupAssignmentsGet) | **Get** /uxi/v1alpha1/network-group-assignments | Get
[**GetUxiV1alpha1SensorGroupAssignmentsGet**](ConfigurationAPI.md#GetUxiV1alpha1SensorGroupAssignmentsGet) | **Get** /uxi/v1alpha1/sensor-group-assignments | Get
[**GetUxiV1alpha1WiredNetworksGet**](ConfigurationAPI.md#GetUxiV1alpha1WiredNetworksGet) | **Get** /uxi/v1alpha1/wired-networks | Get
[**GetUxiV1alpha1WirelessNetworksGet**](ConfigurationAPI.md#GetUxiV1alpha1WirelessNetworksGet) | **Get** /uxi/v1alpha1/wireless-networks | Get
[**GroupsGetUxiV1alpha1GroupsGet**](ConfigurationAPI.md#GroupsGetUxiV1alpha1GroupsGet) | **Get** /uxi/v1alpha1/groups | Groups Get
[**GroupsPatchUxiV1alpha1GroupsGroupUidPatch**](ConfigurationAPI.md#GroupsPatchUxiV1alpha1GroupsGroupUidPatch) | **Patch** /uxi/v1alpha1/groups/{group_uid} | Groups Patch
[**GroupsPostUxiV1alpha1GroupsPost**](ConfigurationAPI.md#GroupsPostUxiV1alpha1GroupsPost) | **Post** /uxi/v1alpha1/groups | Groups Post
[**PostUxiV1alpha1NetworkGroupAssignmentsPost**](ConfigurationAPI.md#PostUxiV1alpha1NetworkGroupAssignmentsPost) | **Post** /uxi/v1alpha1/network-group-assignments | Post
[**PostUxiV1alpha1SensorGroupAssignmentsPost**](ConfigurationAPI.md#PostUxiV1alpha1SensorGroupAssignmentsPost) | **Post** /uxi/v1alpha1/sensor-group-assignments | Post
[**PostUxiV1alpha1ServiceTestGroupAssignmentsPost**](ConfigurationAPI.md#PostUxiV1alpha1ServiceTestGroupAssignmentsPost) | **Post** /uxi/v1alpha1/service-test-group-assignments | Post



## DeleteNetworkGroupAssignmentUxiV1alpha1NetworkGroupAssignmentsIdDelete

> interface{} DeleteNetworkGroupAssignmentUxiV1alpha1NetworkGroupAssignmentsIdDelete(ctx, id).Execute()

Delete Network Group Assignment



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.DeleteNetworkGroupAssignmentUxiV1alpha1NetworkGroupAssignmentsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.DeleteNetworkGroupAssignmentUxiV1alpha1NetworkGroupAssignmentsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkGroupAssignmentUxiV1alpha1NetworkGroupAssignmentsIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.DeleteNetworkGroupAssignmentUxiV1alpha1NetworkGroupAssignmentsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkGroupAssignmentUxiV1alpha1NetworkGroupAssignmentsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSensorGroupAssignmentUxiV1alpha1SensorGroupAssignmentsIdDelete

> interface{} DeleteSensorGroupAssignmentUxiV1alpha1SensorGroupAssignmentsIdDelete(ctx, id).Execute()

Delete Sensor Group Assignment



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.DeleteSensorGroupAssignmentUxiV1alpha1SensorGroupAssignmentsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.DeleteSensorGroupAssignmentUxiV1alpha1SensorGroupAssignmentsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSensorGroupAssignmentUxiV1alpha1SensorGroupAssignmentsIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.DeleteSensorGroupAssignmentUxiV1alpha1SensorGroupAssignmentsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSensorGroupAssignmentUxiV1alpha1SensorGroupAssignmentsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## GroupsPatchUxiV1alpha1GroupsGroupUidPatch

> GroupsPatchResponse GroupsPatchUxiV1alpha1GroupsGroupUidPatch(ctx, groupUid).GroupsPatchRequest(groupsPatchRequest).Execute()

Groups Patch



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
	groupUid := "groupUid_example" // string | 
	groupsPatchRequest := *openapiclient.NewGroupsPatchRequest("Name_example") // GroupsPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GroupsPatchUxiV1alpha1GroupsGroupUidPatch(context.Background(), groupUid).GroupsPatchRequest(groupsPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GroupsPatchUxiV1alpha1GroupsGroupUidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsPatchUxiV1alpha1GroupsGroupUidPatch`: GroupsPatchResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GroupsPatchUxiV1alpha1GroupsGroupUidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsPatchUxiV1alpha1GroupsGroupUidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupsPatchRequest** | [**GroupsPatchRequest**](GroupsPatchRequest.md) |  | 

### Return type

[**GroupsPatchResponse**](GroupsPatchResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
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
	groupsPostRequest := *openapiclient.NewGroupsPostRequest("Name_example") // GroupsPostRequest | 

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


## PostUxiV1alpha1NetworkGroupAssignmentsPost

> NetworkGroupAssignmentsPostResponse PostUxiV1alpha1NetworkGroupAssignmentsPost(ctx).NetworkGroupAssignmentsPostRequest(networkGroupAssignmentsPostRequest).Execute()

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
	networkGroupAssignmentsPostRequest := *openapiclient.NewNetworkGroupAssignmentsPostRequest("GroupId_example", "NetworkId_example") // NetworkGroupAssignmentsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.PostUxiV1alpha1NetworkGroupAssignmentsPost(context.Background()).NetworkGroupAssignmentsPostRequest(networkGroupAssignmentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.PostUxiV1alpha1NetworkGroupAssignmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUxiV1alpha1NetworkGroupAssignmentsPost`: NetworkGroupAssignmentsPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.PostUxiV1alpha1NetworkGroupAssignmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUxiV1alpha1NetworkGroupAssignmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkGroupAssignmentsPostRequest** | [**NetworkGroupAssignmentsPostRequest**](NetworkGroupAssignmentsPostRequest.md) |  | 

### Return type

[**NetworkGroupAssignmentsPostResponse**](NetworkGroupAssignmentsPostResponse.md)

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


## PostUxiV1alpha1ServiceTestGroupAssignmentsPost

> ServiceTestGroupAssignmentsPostResponse PostUxiV1alpha1ServiceTestGroupAssignmentsPost(ctx).ServiceTestGroupAssignmentsPostRequest(serviceTestGroupAssignmentsPostRequest).Execute()

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
	serviceTestGroupAssignmentsPostRequest := *openapiclient.NewServiceTestGroupAssignmentsPostRequest("GroupId_example", "ServiceTestId_example") // ServiceTestGroupAssignmentsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.PostUxiV1alpha1ServiceTestGroupAssignmentsPost(context.Background()).ServiceTestGroupAssignmentsPostRequest(serviceTestGroupAssignmentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.PostUxiV1alpha1ServiceTestGroupAssignmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUxiV1alpha1ServiceTestGroupAssignmentsPost`: ServiceTestGroupAssignmentsPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.PostUxiV1alpha1ServiceTestGroupAssignmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUxiV1alpha1ServiceTestGroupAssignmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceTestGroupAssignmentsPostRequest** | [**ServiceTestGroupAssignmentsPostRequest**](ServiceTestGroupAssignmentsPostRequest.md) |  | 

### Return type

[**ServiceTestGroupAssignmentsPostResponse**](ServiceTestGroupAssignmentsPostResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

