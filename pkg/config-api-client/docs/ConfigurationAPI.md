# \ConfigurationAPI

All URIs are relative to *https://api.capenetworks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentDelete**](ConfigurationAPI.md#AgentDelete) | **Delete** /networking-uxi/v1alpha1/agents/{id} | Agent Delete
[**AgentGroupAssignmentDelete**](ConfigurationAPI.md#AgentGroupAssignmentDelete) | **Delete** /networking-uxi/v1alpha1/agent-group-assignments/{id} | Agent Group Assignment Delete
[**AgentGroupAssignmentPost**](ConfigurationAPI.md#AgentGroupAssignmentPost) | **Post** /networking-uxi/v1alpha1/agent-group-assignments | Agent Group Assignment Post
[**AgentGroupAssignmentsGet**](ConfigurationAPI.md#AgentGroupAssignmentsGet) | **Get** /networking-uxi/v1alpha1/agent-group-assignments | Agent Group Assignments Get
[**AgentPatch**](ConfigurationAPI.md#AgentPatch) | **Patch** /networking-uxi/v1alpha1/agents/{id} | Agent Patch
[**AgentsGet**](ConfigurationAPI.md#AgentsGet) | **Get** /networking-uxi/v1alpha1/agents | Agents Get
[**GroupDelete**](ConfigurationAPI.md#GroupDelete) | **Delete** /networking-uxi/v1alpha1/groups/{id} | Group Delete
[**GroupPatch**](ConfigurationAPI.md#GroupPatch) | **Patch** /networking-uxi/v1alpha1/groups/{id} | Group Patch
[**GroupPost**](ConfigurationAPI.md#GroupPost) | **Post** /networking-uxi/v1alpha1/groups | Group Post
[**GroupsGet**](ConfigurationAPI.md#GroupsGet) | **Get** /networking-uxi/v1alpha1/groups | Groups Get
[**NetworkGroupAssignmentDelete**](ConfigurationAPI.md#NetworkGroupAssignmentDelete) | **Delete** /networking-uxi/v1alpha1/network-group-assignments/{id} | Network Group Assignment Delete
[**NetworkGroupAssignmentPost**](ConfigurationAPI.md#NetworkGroupAssignmentPost) | **Post** /networking-uxi/v1alpha1/network-group-assignments | Network Group Assignment Post
[**NetworkGroupAssignmentsGet**](ConfigurationAPI.md#NetworkGroupAssignmentsGet) | **Get** /networking-uxi/v1alpha1/network-group-assignments | Network Group Assignments Get
[**SensorGroupAssignmentDelete**](ConfigurationAPI.md#SensorGroupAssignmentDelete) | **Delete** /networking-uxi/v1alpha1/sensor-group-assignments/{id} | Sensor Group Assignment Delete
[**SensorGroupAssignmentPost**](ConfigurationAPI.md#SensorGroupAssignmentPost) | **Post** /networking-uxi/v1alpha1/sensor-group-assignments | Sensor Group Assignment Post
[**SensorGroupAssignmentsGet**](ConfigurationAPI.md#SensorGroupAssignmentsGet) | **Get** /networking-uxi/v1alpha1/sensor-group-assignments | Sensor Group Assignments Get
[**SensorPatch**](ConfigurationAPI.md#SensorPatch) | **Patch** /networking-uxi/v1alpha1/sensors/{id} | Sensor Patch
[**SensorsGet**](ConfigurationAPI.md#SensorsGet) | **Get** /networking-uxi/v1alpha1/sensors | Sensors Get
[**ServiceTestGroupAssignmentDelete**](ConfigurationAPI.md#ServiceTestGroupAssignmentDelete) | **Delete** /networking-uxi/v1alpha1/service-test-group-assignments/{id} | Service Test Group Assignment Delete
[**ServiceTestGroupAssignmentPost**](ConfigurationAPI.md#ServiceTestGroupAssignmentPost) | **Post** /networking-uxi/v1alpha1/service-test-group-assignments | Service Test Group Assignment Post
[**ServiceTestGroupAssignmentsGet**](ConfigurationAPI.md#ServiceTestGroupAssignmentsGet) | **Get** /networking-uxi/v1alpha1/service-test-group-assignments | Service Test Group Assignments Get
[**ServiceTestsGet**](ConfigurationAPI.md#ServiceTestsGet) | **Get** /networking-uxi/v1alpha1/service-tests | Service Tests Get
[**WiredNetworksGet**](ConfigurationAPI.md#WiredNetworksGet) | **Get** /networking-uxi/v1alpha1/wired-networks | Wired Networks Get
[**WirelessNetworksGet**](ConfigurationAPI.md#WirelessNetworksGet) | **Get** /networking-uxi/v1alpha1/wireless-networks | Wireless Networks Get



## AgentDelete

> AgentDelete(ctx, id).Execute()

Agent Delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the agent

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationAPI.AgentDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.AgentDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the agent | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGroupAssignmentDelete

> AgentGroupAssignmentDelete(ctx, id).Execute()

Agent Group Assignment Delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationAPI.AgentGroupAssignmentDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.AgentGroupAssignmentDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentGroupAssignmentDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGroupAssignmentPost

> AgentGroupAssignmentPostResponse AgentGroupAssignmentPost(ctx).AgentGroupAssignmentPostRequest(agentGroupAssignmentPostRequest).Execute()

Agent Group Assignment Post



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	agentGroupAssignmentPostRequest := *openapiclient.NewAgentGroupAssignmentPostRequest("GroupId_example", "AgentId_example") // AgentGroupAssignmentPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.AgentGroupAssignmentPost(context.Background()).AgentGroupAssignmentPostRequest(agentGroupAssignmentPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.AgentGroupAssignmentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentGroupAssignmentPost`: AgentGroupAssignmentPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.AgentGroupAssignmentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentGroupAssignmentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentGroupAssignmentPostRequest** | [**AgentGroupAssignmentPostRequest**](AgentGroupAssignmentPostRequest.md) |  | 

### Return type

[**AgentGroupAssignmentPostResponse**](AgentGroupAssignmentPostResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGroupAssignmentsGet

> AgentGroupAssignmentsGetResponse AgentGroupAssignmentsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Agent Group Assignments Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The ID of the resource. (optional)
	next := "next_example" // string | The next cursor for pagination. (optional)
	limit := int32(56) // int32 | The maximum number of items returned in the response. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.AgentGroupAssignmentsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.AgentGroupAssignmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentGroupAssignmentsGet`: AgentGroupAssignmentsGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.AgentGroupAssignmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentGroupAssignmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The ID of the resource. | 
 **next** | **string** | The next cursor for pagination. | 
 **limit** | **int32** | The maximum number of items returned in the response. | [default to 50]

### Return type

[**AgentGroupAssignmentsGetResponse**](AgentGroupAssignmentsGetResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentPatch

> AgentPatchResponse AgentPatch(ctx, id).AgentPatchRequest(agentPatchRequest).Execute()

Agent Patch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the agent
	agentPatchRequest := *openapiclient.NewAgentPatchRequest() // AgentPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.AgentPatch(context.Background(), id).AgentPatchRequest(agentPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.AgentPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentPatch`: AgentPatchResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.AgentPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the agent | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentPatchRequest** | [**AgentPatchRequest**](AgentPatchRequest.md) |  | 

### Return type

[**AgentPatchResponse**](AgentPatchResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsGet

> AgentsGetResponse AgentsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Agents Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The ID of the resource. (optional)
	next := "next_example" // string | The next cursor for pagination. (optional)
	limit := int32(56) // int32 | The maximum number of items returned in the response. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.AgentsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.AgentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsGet`: AgentsGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.AgentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The ID of the resource. | 
 **next** | **string** | The next cursor for pagination. | 
 **limit** | **int32** | The maximum number of items returned in the response. | [default to 50]

### Return type

[**AgentsGetResponse**](AgentsGetResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupDelete

> GroupDelete(ctx, id).Execute()

Group Delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationAPI.GroupDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GroupDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPatch

> GroupPatchResponse GroupPatch(ctx, id).GroupPatchRequest(groupPatchRequest).Execute()

Group Patch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the group
	groupPatchRequest := *openapiclient.NewGroupPatchRequest() // GroupPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GroupPatch(context.Background(), id).GroupPatchRequest(groupPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GroupPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupPatch`: GroupPatchResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GroupPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupPatchRequest** | [**GroupPatchRequest**](GroupPatchRequest.md) |  | 

### Return type

[**GroupPatchResponse**](GroupPatchResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupPost

> GroupPostResponse GroupPost(ctx).GroupPostRequest(groupPostRequest).Execute()

Group Post



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	groupPostRequest := *openapiclient.NewGroupPostRequest("Name_example") // GroupPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GroupPost(context.Background()).GroupPostRequest(groupPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupPost`: GroupPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GroupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupPostRequest** | [**GroupPostRequest**](GroupPostRequest.md) |  | 

### Return type

[**GroupPostResponse**](GroupPostResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The ID of the resource. (optional)
	next := "next_example" // string | The next cursor for pagination. (optional)
	limit := int32(56) // int32 | The maximum number of items returned in the response. (optional) (default to 50)

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
 **id** | **string** | The ID of the resource. | 
 **next** | **string** | The next cursor for pagination. | 
 **limit** | **int32** | The maximum number of items returned in the response. | [default to 50]

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


## NetworkGroupAssignmentDelete

> NetworkGroupAssignmentDelete(ctx, id).Execute()

Network Group Assignment Delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the network group assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationAPI.NetworkGroupAssignmentDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.NetworkGroupAssignmentDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the network group assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkGroupAssignmentDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkGroupAssignmentPost

> NetworkGroupAssignmentPostResponse NetworkGroupAssignmentPost(ctx).NetworkGroupAssignmentPostRequest(networkGroupAssignmentPostRequest).Execute()

Network Group Assignment Post



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	networkGroupAssignmentPostRequest := *openapiclient.NewNetworkGroupAssignmentPostRequest("GroupId_example", "NetworkId_example") // NetworkGroupAssignmentPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.NetworkGroupAssignmentPost(context.Background()).NetworkGroupAssignmentPostRequest(networkGroupAssignmentPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.NetworkGroupAssignmentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkGroupAssignmentPost`: NetworkGroupAssignmentPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.NetworkGroupAssignmentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkGroupAssignmentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkGroupAssignmentPostRequest** | [**NetworkGroupAssignmentPostRequest**](NetworkGroupAssignmentPostRequest.md) |  | 

### Return type

[**NetworkGroupAssignmentPostResponse**](NetworkGroupAssignmentPostResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkGroupAssignmentsGet

> NetworkGroupAssignmentsGetResponse NetworkGroupAssignmentsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Network Group Assignments Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The ID of the resource. (optional)
	next := "next_example" // string | The next cursor for pagination. (optional)
	limit := int32(56) // int32 | The maximum number of items returned in the response. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.NetworkGroupAssignmentsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.NetworkGroupAssignmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkGroupAssignmentsGet`: NetworkGroupAssignmentsGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.NetworkGroupAssignmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkGroupAssignmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The ID of the resource. | 
 **next** | **string** | The next cursor for pagination. | 
 **limit** | **int32** | The maximum number of items returned in the response. | [default to 50]

### Return type

[**NetworkGroupAssignmentsGetResponse**](NetworkGroupAssignmentsGetResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensorGroupAssignmentDelete

> SensorGroupAssignmentDelete(ctx, id).Execute()

Sensor Group Assignment Delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the sensor group assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationAPI.SensorGroupAssignmentDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.SensorGroupAssignmentDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the sensor group assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensorGroupAssignmentDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensorGroupAssignmentPost

> SensorGroupAssignmentPostResponse SensorGroupAssignmentPost(ctx).SensorGroupAssignmentPostRequest(sensorGroupAssignmentPostRequest).Execute()

Sensor Group Assignment Post



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	sensorGroupAssignmentPostRequest := *openapiclient.NewSensorGroupAssignmentPostRequest("GroupId_example", "SensorId_example") // SensorGroupAssignmentPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.SensorGroupAssignmentPost(context.Background()).SensorGroupAssignmentPostRequest(sensorGroupAssignmentPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.SensorGroupAssignmentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SensorGroupAssignmentPost`: SensorGroupAssignmentPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.SensorGroupAssignmentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSensorGroupAssignmentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sensorGroupAssignmentPostRequest** | [**SensorGroupAssignmentPostRequest**](SensorGroupAssignmentPostRequest.md) |  | 

### Return type

[**SensorGroupAssignmentPostResponse**](SensorGroupAssignmentPostResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensorGroupAssignmentsGet

> SensorGroupAssignmentsGetResponse SensorGroupAssignmentsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Sensor Group Assignments Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The ID of the resource. (optional)
	next := "next_example" // string | The next cursor for pagination. (optional)
	limit := int32(56) // int32 | The maximum number of items returned in the response. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.SensorGroupAssignmentsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.SensorGroupAssignmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SensorGroupAssignmentsGet`: SensorGroupAssignmentsGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.SensorGroupAssignmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSensorGroupAssignmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The ID of the resource. | 
 **next** | **string** | The next cursor for pagination. | 
 **limit** | **int32** | The maximum number of items returned in the response. | [default to 50]

### Return type

[**SensorGroupAssignmentsGetResponse**](SensorGroupAssignmentsGetResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensorPatch

> SensorPatchResponse SensorPatch(ctx, id).SensorPatchRequest(sensorPatchRequest).Execute()

Sensor Patch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the sensor
	sensorPatchRequest := *openapiclient.NewSensorPatchRequest() // SensorPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.SensorPatch(context.Background(), id).SensorPatchRequest(sensorPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.SensorPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SensorPatch`: SensorPatchResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.SensorPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the sensor | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensorPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sensorPatchRequest** | [**SensorPatchRequest**](SensorPatchRequest.md) |  | 

### Return type

[**SensorPatchResponse**](SensorPatchResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensorsGet

> SensorsGetResponse SensorsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Sensors Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The ID of the resource. (optional)
	next := "next_example" // string | The next cursor for pagination. (optional)
	limit := int32(56) // int32 | The maximum number of items returned in the response. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.SensorsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.SensorsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SensorsGet`: SensorsGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.SensorsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSensorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The ID of the resource. | 
 **next** | **string** | The next cursor for pagination. | 
 **limit** | **int32** | The maximum number of items returned in the response. | [default to 50]

### Return type

[**SensorsGetResponse**](SensorsGetResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceTestGroupAssignmentDelete

> ServiceTestGroupAssignmentDelete(ctx, id).Execute()

Service Test Group Assignment Delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The unique identifier of the service test group assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationAPI.ServiceTestGroupAssignmentDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.ServiceTestGroupAssignmentDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the service test group assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceTestGroupAssignmentDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceTestGroupAssignmentPost

> ServiceTestGroupAssignmentPostResponse ServiceTestGroupAssignmentPost(ctx).ServiceTestGroupAssignmentPostRequest(serviceTestGroupAssignmentPostRequest).Execute()

Service Test Group Assignment Post



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	serviceTestGroupAssignmentPostRequest := *openapiclient.NewServiceTestGroupAssignmentPostRequest("GroupId_example", "ServiceTestId_example") // ServiceTestGroupAssignmentPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.ServiceTestGroupAssignmentPost(context.Background()).ServiceTestGroupAssignmentPostRequest(serviceTestGroupAssignmentPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.ServiceTestGroupAssignmentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceTestGroupAssignmentPost`: ServiceTestGroupAssignmentPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.ServiceTestGroupAssignmentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceTestGroupAssignmentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceTestGroupAssignmentPostRequest** | [**ServiceTestGroupAssignmentPostRequest**](ServiceTestGroupAssignmentPostRequest.md) |  | 

### Return type

[**ServiceTestGroupAssignmentPostResponse**](ServiceTestGroupAssignmentPostResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceTestGroupAssignmentsGet

> ServiceTestGroupAssignmentsGetResponse ServiceTestGroupAssignmentsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Service Test Group Assignments Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The ID of the resource. (optional)
	next := "next_example" // string | The next cursor for pagination. (optional)
	limit := int32(56) // int32 | The maximum number of items returned in the response. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.ServiceTestGroupAssignmentsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.ServiceTestGroupAssignmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceTestGroupAssignmentsGet`: ServiceTestGroupAssignmentsGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.ServiceTestGroupAssignmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceTestGroupAssignmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The ID of the resource. | 
 **next** | **string** | The next cursor for pagination. | 
 **limit** | **int32** | The maximum number of items returned in the response. | [default to 50]

### Return type

[**ServiceTestGroupAssignmentsGetResponse**](ServiceTestGroupAssignmentsGetResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceTestsGet

> ServiceTestsGetResponse ServiceTestsGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Service Tests Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The ID of the resource. (optional)
	next := "next_example" // string | The next cursor for pagination. (optional)
	limit := int32(56) // int32 | The maximum number of items returned in the response. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.ServiceTestsGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.ServiceTestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceTestsGet`: ServiceTestsGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.ServiceTestsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceTestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The ID of the resource. | 
 **next** | **string** | The next cursor for pagination. | 
 **limit** | **int32** | The maximum number of items returned in the response. | [default to 50]

### Return type

[**ServiceTestsGetResponse**](ServiceTestsGetResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WiredNetworksGet

> WiredNetworksGetResponse WiredNetworksGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Wired Networks Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The ID of the resource. (optional)
	next := "next_example" // string | The next cursor for pagination. (optional)
	limit := int32(56) // int32 | The maximum number of items returned in the response. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.WiredNetworksGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.WiredNetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WiredNetworksGet`: WiredNetworksGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.WiredNetworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWiredNetworksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The ID of the resource. | 
 **next** | **string** | The next cursor for pagination. | 
 **limit** | **int32** | The maximum number of items returned in the response. | [default to 50]

### Return type

[**WiredNetworksGetResponse**](WiredNetworksGetResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessNetworksGet

> WirelessNetworksGetResponse WirelessNetworksGet(ctx).Id(id).Next(next).Limit(limit).Execute()

Wireless Networks Get



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func main() {
	id := "id_example" // string | The ID of the resource. (optional)
	next := "next_example" // string | The next cursor for pagination. (optional)
	limit := int32(56) // int32 | The maximum number of items returned in the response. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.WirelessNetworksGet(context.Background()).Id(id).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.WirelessNetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessNetworksGet`: WirelessNetworksGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.WirelessNetworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessNetworksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The ID of the resource. | 
 **next** | **string** | The next cursor for pagination. | 
 **limit** | **int32** | The maximum number of items returned in the response. | [default to 50]

### Return type

[**WirelessNetworksGetResponse**](WirelessNetworksGetResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

