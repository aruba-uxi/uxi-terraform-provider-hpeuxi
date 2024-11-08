# \ConfigurationAPI

All URIs are relative to *https://api.capenetworks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentGroupAssignmentsGet**](ConfigurationAPI.md#AgentGroupAssignmentsGet) | **Get** /networking-uxi/v1alpha1/agent-group-assignments | Agent Group Assignments Get
[**AgentGroupAssignmentsPost**](ConfigurationAPI.md#AgentGroupAssignmentsPost) | **Post** /networking-uxi/v1alpha1/agent-group-assignments | Agent Group Assignments Post
[**AgentsDelete**](ConfigurationAPI.md#AgentsDelete) | **Delete** /networking-uxi/v1alpha1/agents/{agent_uid} | Agents Delete
[**AgentsGet**](ConfigurationAPI.md#AgentsGet) | **Get** /networking-uxi/v1alpha1/agents | Agents Get
[**GroupsDelete**](ConfigurationAPI.md#GroupsDelete) | **Delete** /networking-uxi/v1alpha1/groups/{group_uid} | Groups Delete
[**GroupsGet**](ConfigurationAPI.md#GroupsGet) | **Get** /networking-uxi/v1alpha1/groups | Groups Get
[**GroupsPatch**](ConfigurationAPI.md#GroupsPatch) | **Patch** /networking-uxi/v1alpha1/groups/{group_uid} | Groups Patch
[**GroupsPost**](ConfigurationAPI.md#GroupsPost) | **Post** /networking-uxi/v1alpha1/groups | Groups Post
[**NetworkGroupAssignmentsDelete**](ConfigurationAPI.md#NetworkGroupAssignmentsDelete) | **Delete** /networking-uxi/v1alpha1/network-group-assignments/{id} | Network Group Assignments Delete
[**NetworkGroupAssignmentsGet**](ConfigurationAPI.md#NetworkGroupAssignmentsGet) | **Get** /networking-uxi/v1alpha1/network-group-assignments | Network Group Assignments Get
[**NetworkGroupAssignmentsPost**](ConfigurationAPI.md#NetworkGroupAssignmentsPost) | **Post** /networking-uxi/v1alpha1/network-group-assignments | Network Group Assignments Post
[**SensorGroupAssignmentsDelete**](ConfigurationAPI.md#SensorGroupAssignmentsDelete) | **Delete** /networking-uxi/v1alpha1/sensor-group-assignments/{id} | Sensor Group Assignments Delete
[**SensorGroupAssignmentsGet**](ConfigurationAPI.md#SensorGroupAssignmentsGet) | **Get** /networking-uxi/v1alpha1/sensor-group-assignments | Sensor Group Assignments Get
[**SensorGroupAssignmentsPost**](ConfigurationAPI.md#SensorGroupAssignmentsPost) | **Post** /networking-uxi/v1alpha1/sensor-group-assignments | Sensor Group Assignments Post
[**SensorsGet**](ConfigurationAPI.md#SensorsGet) | **Get** /networking-uxi/v1alpha1/sensors | Sensors Get
[**SensorsPatch**](ConfigurationAPI.md#SensorsPatch) | **Patch** /networking-uxi/v1alpha1/sensors/{sensor_uid} | Sensors Patch
[**ServiceTestGroupAssignmentsDelete**](ConfigurationAPI.md#ServiceTestGroupAssignmentsDelete) | **Delete** /networking-uxi/v1alpha1/service-test-group-assignments/{id} | Service Test Group Assignments Delete
[**ServiceTestGroupAssignmentsGet**](ConfigurationAPI.md#ServiceTestGroupAssignmentsGet) | **Get** /networking-uxi/v1alpha1/service-test-group-assignments | Service Test Group Assignments Get
[**ServiceTestGroupAssignmentsPost**](ConfigurationAPI.md#ServiceTestGroupAssignmentsPost) | **Post** /networking-uxi/v1alpha1/service-test-group-assignments | Service Test Group Assignments Post
[**ServiceTestsGet**](ConfigurationAPI.md#ServiceTestsGet) | **Get** /networking-uxi/v1alpha1/service-tests | Service Tests Get
[**WiredNetworksGet**](ConfigurationAPI.md#WiredNetworksGet) | **Get** /networking-uxi/v1alpha1/wired-networks | Wired Networks Get
[**WirelessNetworksGet**](ConfigurationAPI.md#WirelessNetworksGet) | **Get** /networking-uxi/v1alpha1/wireless-networks | Wireless Networks Get



## AgentGroupAssignmentsGet

> AgentGroupAssignmentsResponse AgentGroupAssignmentsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Agent Group Assignments Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.AgentGroupAssignmentsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.AgentGroupAssignmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentGroupAssignmentsGet`: AgentGroupAssignmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.AgentGroupAssignmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentGroupAssignmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **next** | **string** |  | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**AgentGroupAssignmentsResponse**](AgentGroupAssignmentsResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGroupAssignmentsPost

> AgentGroupAssignmentResponse AgentGroupAssignmentsPost(ctx).AgentGroupAssignmentsPostRequest(agentGroupAssignmentsPostRequest).Execute()

Agent Group Assignments Post



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	agentGroupAssignmentsPostRequest := *openapiclient.NewAgentGroupAssignmentsPostRequest("GroupId_example", "AgentId_example") // AgentGroupAssignmentsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.AgentGroupAssignmentsPost(context.Background()).AgentGroupAssignmentsPostRequest(agentGroupAssignmentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.AgentGroupAssignmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentGroupAssignmentsPost`: AgentGroupAssignmentResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.AgentGroupAssignmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentGroupAssignmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentGroupAssignmentsPostRequest** | [**AgentGroupAssignmentsPostRequest**](AgentGroupAssignmentsPostRequest.md) |  | 

### Return type

[**AgentGroupAssignmentResponse**](AgentGroupAssignmentResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsDelete

> interface{} AgentsDelete(ctx, agentUid).Execute()

Agents Delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	agentUid := "agentUid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.AgentsDelete(context.Background(), agentUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.AgentsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.AgentsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsDeleteRequest struct via the builder pattern


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


## AgentsGet

> AgentsResponse AgentsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Agents Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.AgentsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.AgentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsGet`: AgentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.AgentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **next** | **string** |  | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**AgentsResponse**](AgentsResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsDelete

> interface{} GroupsDelete(ctx, groupUid).Execute()

Groups Delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	groupUid := "groupUid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GroupsDelete(context.Background(), groupUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GroupsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GroupsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsDeleteRequest struct via the builder pattern


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


## GroupsGet

> GroupsGetResponse GroupsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Groups Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GroupsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGet`: GroupsGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetRequest struct via the builder pattern


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


## GroupsPatch

> GroupsPatchResponse GroupsPatch(ctx, groupUid).GroupsPatchRequest(groupsPatchRequest).Execute()

Groups Patch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	groupUid := "groupUid_example" // string | 
	groupsPatchRequest := *openapiclient.NewGroupsPatchRequest("Name_example") // GroupsPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GroupsPatch(context.Background(), groupUid).GroupsPatchRequest(groupsPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GroupsPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsPatch`: GroupsPatchResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GroupsPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupsPatchRequest** | [**GroupsPatchRequest**](GroupsPatchRequest.md) |  | 

### Return type

[**GroupsPatchResponse**](GroupsPatchResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsPost

> GroupsPostResponse GroupsPost(ctx).GroupsPostRequest(groupsPostRequest).Execute()

Groups Post



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	groupsPostRequest := *openapiclient.NewGroupsPostRequest("Name_example") // GroupsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GroupsPost(context.Background()).GroupsPostRequest(groupsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsPost`: GroupsPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsPostRequest struct via the builder pattern


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


## NetworkGroupAssignmentsDelete

> interface{} NetworkGroupAssignmentsDelete(ctx, id).Execute()

Network Group Assignments Delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.NetworkGroupAssignmentsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.NetworkGroupAssignmentsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkGroupAssignmentsDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.NetworkGroupAssignmentsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkGroupAssignmentsDeleteRequest struct via the builder pattern


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


## NetworkGroupAssignmentsGet

> NetworkGroupAssignmentsResponse NetworkGroupAssignmentsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Network Group Assignments Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.NetworkGroupAssignmentsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.NetworkGroupAssignmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkGroupAssignmentsGet`: NetworkGroupAssignmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.NetworkGroupAssignmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkGroupAssignmentsGetRequest struct via the builder pattern


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


## NetworkGroupAssignmentsPost

> NetworkGroupAssignmentsPostResponse NetworkGroupAssignmentsPost(ctx).NetworkGroupAssignmentsPostRequest(networkGroupAssignmentsPostRequest).Execute()

Network Group Assignments Post



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	networkGroupAssignmentsPostRequest := *openapiclient.NewNetworkGroupAssignmentsPostRequest("GroupId_example", "NetworkId_example") // NetworkGroupAssignmentsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.NetworkGroupAssignmentsPost(context.Background()).NetworkGroupAssignmentsPostRequest(networkGroupAssignmentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.NetworkGroupAssignmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkGroupAssignmentsPost`: NetworkGroupAssignmentsPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.NetworkGroupAssignmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkGroupAssignmentsPostRequest struct via the builder pattern


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


## SensorGroupAssignmentsDelete

> interface{} SensorGroupAssignmentsDelete(ctx, id).Execute()

Sensor Group Assignments Delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.SensorGroupAssignmentsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.SensorGroupAssignmentsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SensorGroupAssignmentsDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.SensorGroupAssignmentsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensorGroupAssignmentsDeleteRequest struct via the builder pattern


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


## SensorGroupAssignmentsGet

> SensorGroupAssignmentsResponse SensorGroupAssignmentsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Sensor Group Assignments Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.SensorGroupAssignmentsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.SensorGroupAssignmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SensorGroupAssignmentsGet`: SensorGroupAssignmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.SensorGroupAssignmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSensorGroupAssignmentsGetRequest struct via the builder pattern


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


## SensorGroupAssignmentsPost

> SensorGroupAssignmentResponse SensorGroupAssignmentsPost(ctx).SensorGroupAssignmentsPostRequest(sensorGroupAssignmentsPostRequest).Execute()

Sensor Group Assignments Post



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	sensorGroupAssignmentsPostRequest := *openapiclient.NewSensorGroupAssignmentsPostRequest("GroupId_example", "SensorId_example") // SensorGroupAssignmentsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.SensorGroupAssignmentsPost(context.Background()).SensorGroupAssignmentsPostRequest(sensorGroupAssignmentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.SensorGroupAssignmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SensorGroupAssignmentsPost`: SensorGroupAssignmentResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.SensorGroupAssignmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSensorGroupAssignmentsPostRequest struct via the builder pattern


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


## SensorsGet

> SensorsResponse SensorsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Sensors Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.SensorsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.SensorsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SensorsGet`: SensorsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.SensorsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSensorsGetRequest struct via the builder pattern


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


## SensorsPatch

> SensorsPatchResponse SensorsPatch(ctx, sensorUid).SensorsPatchRequest(sensorsPatchRequest).Execute()

Sensors Patch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	sensorUid := "sensorUid_example" // string | 
	sensorsPatchRequest := *openapiclient.NewSensorsPatchRequest() // SensorsPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.SensorsPatch(context.Background(), sensorUid).SensorsPatchRequest(sensorsPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.SensorsPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SensorsPatch`: SensorsPatchResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.SensorsPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sensorUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensorsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sensorsPatchRequest** | [**SensorsPatchRequest**](SensorsPatchRequest.md) |  | 

### Return type

[**SensorsPatchResponse**](SensorsPatchResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceTestGroupAssignmentsDelete

> interface{} ServiceTestGroupAssignmentsDelete(ctx, id).Execute()

Service Test Group Assignments Delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.ServiceTestGroupAssignmentsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.ServiceTestGroupAssignmentsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceTestGroupAssignmentsDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.ServiceTestGroupAssignmentsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceTestGroupAssignmentsDeleteRequest struct via the builder pattern


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


## ServiceTestGroupAssignmentsGet

> ServiceTestGroupAssignmentsResponse ServiceTestGroupAssignmentsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Service Test Group Assignments Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.ServiceTestGroupAssignmentsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.ServiceTestGroupAssignmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceTestGroupAssignmentsGet`: ServiceTestGroupAssignmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.ServiceTestGroupAssignmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceTestGroupAssignmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **next** | **string** |  | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**ServiceTestGroupAssignmentsResponse**](ServiceTestGroupAssignmentsResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceTestGroupAssignmentsPost

> ServiceTestGroupAssignmentsPostResponse ServiceTestGroupAssignmentsPost(ctx).ServiceTestGroupAssignmentsPostRequest(serviceTestGroupAssignmentsPostRequest).Execute()

Service Test Group Assignments Post



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	serviceTestGroupAssignmentsPostRequest := *openapiclient.NewServiceTestGroupAssignmentsPostRequest("GroupId_example", "ServiceTestId_example") // ServiceTestGroupAssignmentsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.ServiceTestGroupAssignmentsPost(context.Background()).ServiceTestGroupAssignmentsPostRequest(serviceTestGroupAssignmentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.ServiceTestGroupAssignmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceTestGroupAssignmentsPost`: ServiceTestGroupAssignmentsPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.ServiceTestGroupAssignmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceTestGroupAssignmentsPostRequest struct via the builder pattern


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


## ServiceTestsGet

> ServiceTestsListResponse ServiceTestsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Service Tests Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.ServiceTestsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.ServiceTestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceTestsGet`: ServiceTestsListResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.ServiceTestsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceTestsGetRequest struct via the builder pattern


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


## WiredNetworksGet

> WiredNetworksResponse WiredNetworksGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Wired Networks Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.WiredNetworksGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.WiredNetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WiredNetworksGet`: WiredNetworksResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.WiredNetworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWiredNetworksGetRequest struct via the builder pattern


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


## WirelessNetworksGet

> WirelessNetworksResponse WirelessNetworksGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Wireless Networks Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
)

func main() {
	id := "id_example" // string |  (optional)
	next := "next_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.WirelessNetworksGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.WirelessNetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessNetworksGet`: WirelessNetworksResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.WirelessNetworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessNetworksGetRequest struct via the builder pattern


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

