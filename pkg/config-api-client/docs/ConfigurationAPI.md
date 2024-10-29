# \ConfigurationAPI

All URIs are relative to *https://api.capenetworks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNetworkGroupAssignmentNetworkingUxiV1alpha1NetworkGroupAssignmentsIdDelete**](ConfigurationAPI.md#DeleteNetworkGroupAssignmentNetworkingUxiV1alpha1NetworkGroupAssignmentsIdDelete) | **Delete** /networking-uxi/v1alpha1/network-group-assignments/{id} | Delete Network Group Assignment
[**DeleteSensorGroupAssignmentNetworkingUxiV1alpha1SensorGroupAssignmentsIdDelete**](ConfigurationAPI.md#DeleteSensorGroupAssignmentNetworkingUxiV1alpha1SensorGroupAssignmentsIdDelete) | **Delete** /networking-uxi/v1alpha1/sensor-group-assignments/{id} | Delete Sensor Group Assignment
[**DeleteServiceTestGroupAssignmentNetworkingUxiV1alpha1ServiceTestGroupAssignmentsIdDelete**](ConfigurationAPI.md#DeleteServiceTestGroupAssignmentNetworkingUxiV1alpha1ServiceTestGroupAssignmentsIdDelete) | **Delete** /networking-uxi/v1alpha1/service-test-group-assignments/{id} | Delete Service Test Group Assignment
[**GetNetworkingUxiV1alpha1NetworkGroupAssignmentsGet**](ConfigurationAPI.md#GetNetworkingUxiV1alpha1NetworkGroupAssignmentsGet) | **Get** /networking-uxi/v1alpha1/network-group-assignments | Get
[**GetNetworkingUxiV1alpha1SensorGroupAssignmentsGet**](ConfigurationAPI.md#GetNetworkingUxiV1alpha1SensorGroupAssignmentsGet) | **Get** /networking-uxi/v1alpha1/sensor-group-assignments | Get
[**GetNetworkingUxiV1alpha1SensorsGet**](ConfigurationAPI.md#GetNetworkingUxiV1alpha1SensorsGet) | **Get** /networking-uxi/v1alpha1/sensors | Get
[**GetNetworkingUxiV1alpha1WiredNetworksGet**](ConfigurationAPI.md#GetNetworkingUxiV1alpha1WiredNetworksGet) | **Get** /networking-uxi/v1alpha1/wired-networks | Get
[**GetNetworkingUxiV1alpha1WirelessNetworksGet**](ConfigurationAPI.md#GetNetworkingUxiV1alpha1WirelessNetworksGet) | **Get** /networking-uxi/v1alpha1/wireless-networks | Get
[**GroupsDeleteNetworkingUxiV1alpha1GroupsGroupUidDelete**](ConfigurationAPI.md#GroupsDeleteNetworkingUxiV1alpha1GroupsGroupUidDelete) | **Delete** /networking-uxi/v1alpha1/groups/{group_uid} | Groups Delete
[**GroupsGetNetworkingUxiV1alpha1GroupsGet**](ConfigurationAPI.md#GroupsGetNetworkingUxiV1alpha1GroupsGet) | **Get** /networking-uxi/v1alpha1/groups | Groups Get
[**GroupsPatchNetworkingUxiV1alpha1GroupsGroupUidPatch**](ConfigurationAPI.md#GroupsPatchNetworkingUxiV1alpha1GroupsGroupUidPatch) | **Patch** /networking-uxi/v1alpha1/groups/{group_uid} | Groups Patch
[**GroupsPostNetworkingUxiV1alpha1GroupsPost**](ConfigurationAPI.md#GroupsPostNetworkingUxiV1alpha1GroupsPost) | **Post** /networking-uxi/v1alpha1/groups | Groups Post
[**PostNetworkingUxiV1alpha1NetworkGroupAssignmentsPost**](ConfigurationAPI.md#PostNetworkingUxiV1alpha1NetworkGroupAssignmentsPost) | **Post** /networking-uxi/v1alpha1/network-group-assignments | Post
[**PostNetworkingUxiV1alpha1SensorGroupAssignmentsPost**](ConfigurationAPI.md#PostNetworkingUxiV1alpha1SensorGroupAssignmentsPost) | **Post** /networking-uxi/v1alpha1/sensor-group-assignments | Post
[**PostNetworkingUxiV1alpha1ServiceTestGroupAssignmentsPost**](ConfigurationAPI.md#PostNetworkingUxiV1alpha1ServiceTestGroupAssignmentsPost) | **Post** /networking-uxi/v1alpha1/service-test-group-assignments | Post
[**ServiceTestsGetNetworkingUxiV1alpha1ServiceTestsGet**](ConfigurationAPI.md#ServiceTestsGetNetworkingUxiV1alpha1ServiceTestsGet) | **Get** /networking-uxi/v1alpha1/service-tests | Service Tests Get



## DeleteNetworkGroupAssignmentNetworkingUxiV1alpha1NetworkGroupAssignmentsIdDelete

> interface{} DeleteNetworkGroupAssignmentNetworkingUxiV1alpha1NetworkGroupAssignmentsIdDelete(ctx, id).Execute()

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
	resp, r, err := apiClient.ConfigurationAPI.DeleteNetworkGroupAssignmentNetworkingUxiV1alpha1NetworkGroupAssignmentsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.DeleteNetworkGroupAssignmentNetworkingUxiV1alpha1NetworkGroupAssignmentsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkGroupAssignmentNetworkingUxiV1alpha1NetworkGroupAssignmentsIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.DeleteNetworkGroupAssignmentNetworkingUxiV1alpha1NetworkGroupAssignmentsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkGroupAssignmentNetworkingUxiV1alpha1NetworkGroupAssignmentsIdDeleteRequest struct via the builder pattern


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


## DeleteSensorGroupAssignmentNetworkingUxiV1alpha1SensorGroupAssignmentsIdDelete

> interface{} DeleteSensorGroupAssignmentNetworkingUxiV1alpha1SensorGroupAssignmentsIdDelete(ctx, id).Execute()

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
	resp, r, err := apiClient.ConfigurationAPI.DeleteSensorGroupAssignmentNetworkingUxiV1alpha1SensorGroupAssignmentsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.DeleteSensorGroupAssignmentNetworkingUxiV1alpha1SensorGroupAssignmentsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSensorGroupAssignmentNetworkingUxiV1alpha1SensorGroupAssignmentsIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.DeleteSensorGroupAssignmentNetworkingUxiV1alpha1SensorGroupAssignmentsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSensorGroupAssignmentNetworkingUxiV1alpha1SensorGroupAssignmentsIdDeleteRequest struct via the builder pattern


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


## DeleteServiceTestGroupAssignmentNetworkingUxiV1alpha1ServiceTestGroupAssignmentsIdDelete

> interface{} DeleteServiceTestGroupAssignmentNetworkingUxiV1alpha1ServiceTestGroupAssignmentsIdDelete(ctx, id).Execute()

Delete Service Test Group Assignment



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
	resp, r, err := apiClient.ConfigurationAPI.DeleteServiceTestGroupAssignmentNetworkingUxiV1alpha1ServiceTestGroupAssignmentsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.DeleteServiceTestGroupAssignmentNetworkingUxiV1alpha1ServiceTestGroupAssignmentsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteServiceTestGroupAssignmentNetworkingUxiV1alpha1ServiceTestGroupAssignmentsIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.DeleteServiceTestGroupAssignmentNetworkingUxiV1alpha1ServiceTestGroupAssignmentsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceTestGroupAssignmentNetworkingUxiV1alpha1ServiceTestGroupAssignmentsIdDeleteRequest struct via the builder pattern


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


## GetNetworkingUxiV1alpha1NetworkGroupAssignmentsGet

> NetworkGroupAssignmentsResponse GetNetworkingUxiV1alpha1NetworkGroupAssignmentsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

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
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GetNetworkingUxiV1alpha1NetworkGroupAssignmentsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetNetworkingUxiV1alpha1NetworkGroupAssignmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkingUxiV1alpha1NetworkGroupAssignmentsGet`: NetworkGroupAssignmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetNetworkingUxiV1alpha1NetworkGroupAssignmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingUxiV1alpha1NetworkGroupAssignmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **next** | **string** |  | 
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


## GetNetworkingUxiV1alpha1SensorGroupAssignmentsGet

> SensorGroupAssignmentsResponse GetNetworkingUxiV1alpha1SensorGroupAssignmentsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

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
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GetNetworkingUxiV1alpha1SensorGroupAssignmentsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetNetworkingUxiV1alpha1SensorGroupAssignmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkingUxiV1alpha1SensorGroupAssignmentsGet`: SensorGroupAssignmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetNetworkingUxiV1alpha1SensorGroupAssignmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingUxiV1alpha1SensorGroupAssignmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **next** | **string** |  | 
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


## GetNetworkingUxiV1alpha1SensorsGet

> SensorsResponse GetNetworkingUxiV1alpha1SensorsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

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
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GetNetworkingUxiV1alpha1SensorsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetNetworkingUxiV1alpha1SensorsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkingUxiV1alpha1SensorsGet`: SensorsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetNetworkingUxiV1alpha1SensorsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingUxiV1alpha1SensorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **next** | **string** |  | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**SensorsResponse**](SensorsResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkingUxiV1alpha1WiredNetworksGet

> WiredNetworksResponse GetNetworkingUxiV1alpha1WiredNetworksGet(ctx).Id(id).Next(next).Limit(limit).Execute()

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
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GetNetworkingUxiV1alpha1WiredNetworksGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetNetworkingUxiV1alpha1WiredNetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkingUxiV1alpha1WiredNetworksGet`: WiredNetworksResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetNetworkingUxiV1alpha1WiredNetworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingUxiV1alpha1WiredNetworksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **next** | **string** |  | 
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


## GetNetworkingUxiV1alpha1WirelessNetworksGet

> WirelessNetworksResponse GetNetworkingUxiV1alpha1WirelessNetworksGet(ctx).Id(id).Next(next).Limit(limit).Execute()

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
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GetNetworkingUxiV1alpha1WirelessNetworksGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetNetworkingUxiV1alpha1WirelessNetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkingUxiV1alpha1WirelessNetworksGet`: WirelessNetworksResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetNetworkingUxiV1alpha1WirelessNetworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingUxiV1alpha1WirelessNetworksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **next** | **string** |  | 
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


## GroupsDeleteNetworkingUxiV1alpha1GroupsGroupUidDelete

> interface{} GroupsDeleteNetworkingUxiV1alpha1GroupsGroupUidDelete(ctx, groupUid).Execute()

Groups Delete



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GroupsDeleteNetworkingUxiV1alpha1GroupsGroupUidDelete(context.Background(), groupUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GroupsDeleteNetworkingUxiV1alpha1GroupsGroupUidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsDeleteNetworkingUxiV1alpha1GroupsGroupUidDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GroupsDeleteNetworkingUxiV1alpha1GroupsGroupUidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsDeleteNetworkingUxiV1alpha1GroupsGroupUidDeleteRequest struct via the builder pattern


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


## GroupsGetNetworkingUxiV1alpha1GroupsGet

> GroupsGetResponse GroupsGetNetworkingUxiV1alpha1GroupsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

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
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GroupsGetNetworkingUxiV1alpha1GroupsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GroupsGetNetworkingUxiV1alpha1GroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGetNetworkingUxiV1alpha1GroupsGet`: GroupsGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GroupsGetNetworkingUxiV1alpha1GroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetNetworkingUxiV1alpha1GroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **next** | **string** |  | 
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


## GroupsPatchNetworkingUxiV1alpha1GroupsGroupUidPatch

> GroupsPatchResponse GroupsPatchNetworkingUxiV1alpha1GroupsGroupUidPatch(ctx, groupUid).GroupsPatchRequest(groupsPatchRequest).Execute()

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
	resp, r, err := apiClient.ConfigurationAPI.GroupsPatchNetworkingUxiV1alpha1GroupsGroupUidPatch(context.Background(), groupUid).GroupsPatchRequest(groupsPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GroupsPatchNetworkingUxiV1alpha1GroupsGroupUidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsPatchNetworkingUxiV1alpha1GroupsGroupUidPatch`: GroupsPatchResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GroupsPatchNetworkingUxiV1alpha1GroupsGroupUidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsPatchNetworkingUxiV1alpha1GroupsGroupUidPatchRequest struct via the builder pattern


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


## GroupsPostNetworkingUxiV1alpha1GroupsPost

> GroupsPostResponse GroupsPostNetworkingUxiV1alpha1GroupsPost(ctx).GroupsPostRequest(groupsPostRequest).Execute()

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
	resp, r, err := apiClient.ConfigurationAPI.GroupsPostNetworkingUxiV1alpha1GroupsPost(context.Background()).GroupsPostRequest(groupsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GroupsPostNetworkingUxiV1alpha1GroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsPostNetworkingUxiV1alpha1GroupsPost`: GroupsPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GroupsPostNetworkingUxiV1alpha1GroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsPostNetworkingUxiV1alpha1GroupsPostRequest struct via the builder pattern


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


## PostNetworkingUxiV1alpha1NetworkGroupAssignmentsPost

> NetworkGroupAssignmentsPostResponse PostNetworkingUxiV1alpha1NetworkGroupAssignmentsPost(ctx).NetworkGroupAssignmentsPostRequest(networkGroupAssignmentsPostRequest).Execute()

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
	resp, r, err := apiClient.ConfigurationAPI.PostNetworkingUxiV1alpha1NetworkGroupAssignmentsPost(context.Background()).NetworkGroupAssignmentsPostRequest(networkGroupAssignmentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.PostNetworkingUxiV1alpha1NetworkGroupAssignmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNetworkingUxiV1alpha1NetworkGroupAssignmentsPost`: NetworkGroupAssignmentsPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.PostNetworkingUxiV1alpha1NetworkGroupAssignmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNetworkingUxiV1alpha1NetworkGroupAssignmentsPostRequest struct via the builder pattern


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


## PostNetworkingUxiV1alpha1SensorGroupAssignmentsPost

> SensorGroupAssignmentResponse PostNetworkingUxiV1alpha1SensorGroupAssignmentsPost(ctx).SensorGroupAssignmentsPostRequest(sensorGroupAssignmentsPostRequest).Execute()

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
	resp, r, err := apiClient.ConfigurationAPI.PostNetworkingUxiV1alpha1SensorGroupAssignmentsPost(context.Background()).SensorGroupAssignmentsPostRequest(sensorGroupAssignmentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.PostNetworkingUxiV1alpha1SensorGroupAssignmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNetworkingUxiV1alpha1SensorGroupAssignmentsPost`: SensorGroupAssignmentResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.PostNetworkingUxiV1alpha1SensorGroupAssignmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNetworkingUxiV1alpha1SensorGroupAssignmentsPostRequest struct via the builder pattern


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


## PostNetworkingUxiV1alpha1ServiceTestGroupAssignmentsPost

> ServiceTestGroupAssignmentsPostResponse PostNetworkingUxiV1alpha1ServiceTestGroupAssignmentsPost(ctx).ServiceTestGroupAssignmentsPostRequest(serviceTestGroupAssignmentsPostRequest).Execute()

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
	resp, r, err := apiClient.ConfigurationAPI.PostNetworkingUxiV1alpha1ServiceTestGroupAssignmentsPost(context.Background()).ServiceTestGroupAssignmentsPostRequest(serviceTestGroupAssignmentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.PostNetworkingUxiV1alpha1ServiceTestGroupAssignmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNetworkingUxiV1alpha1ServiceTestGroupAssignmentsPost`: ServiceTestGroupAssignmentsPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.PostNetworkingUxiV1alpha1ServiceTestGroupAssignmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNetworkingUxiV1alpha1ServiceTestGroupAssignmentsPostRequest struct via the builder pattern


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


## ServiceTestsGetNetworkingUxiV1alpha1ServiceTestsGet

> ServiceTestsListResponse ServiceTestsGetNetworkingUxiV1alpha1ServiceTestsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Service Tests Get



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
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.ServiceTestsGetNetworkingUxiV1alpha1ServiceTestsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.ServiceTestsGetNetworkingUxiV1alpha1ServiceTestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceTestsGetNetworkingUxiV1alpha1ServiceTestsGet`: ServiceTestsListResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.ServiceTestsGetNetworkingUxiV1alpha1ServiceTestsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceTestsGetNetworkingUxiV1alpha1ServiceTestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **next** | **string** |  | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**ServiceTestsListResponse**](ServiceTestsListResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

