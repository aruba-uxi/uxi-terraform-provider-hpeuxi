# \ConfigurationAPI

All URIs are relative to *https://api.capenetworks.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigurationAppV1SensorGroupAssignmentsGet**](ConfigurationAPI.md#GetConfigurationAppV1SensorGroupAssignmentsGet) | **Get** /configuration/app/v1/sensor-group-assignments | Get
[**GetLivezHealthLivezGet**](ConfigurationAPI.md#GetLivezHealthLivezGet) | **Get** /health/livez | Live health check
[**GetReadyzHealthReadyzGet**](ConfigurationAPI.md#GetReadyzHealthReadyzGet) | **Get** /health/readyz | Ready health check
[**GetStatusHealthStatusGet**](ConfigurationAPI.md#GetStatusHealthStatusGet) | **Get** /health/status | Service stats endpoint
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
	limit := int32(56) // int32 |  (optional)

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
 **limit** | **int32** |  | 

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


## GetLivezHealthLivezGet

> LivenessResponse GetLivezHealthLivezGet(ctx).Execute()

Live health check



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GetLivezHealthLivezGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetLivezHealthLivezGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivezHealthLivezGet`: LivenessResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetLivezHealthLivezGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivezHealthLivezGetRequest struct via the builder pattern


### Return type

[**LivenessResponse**](LivenessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReadyzHealthReadyzGet

> ReadinessResponse GetReadyzHealthReadyzGet(ctx).Execute()

Ready health check



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GetReadyzHealthReadyzGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetReadyzHealthReadyzGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReadyzHealthReadyzGet`: ReadinessResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetReadyzHealthReadyzGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReadyzHealthReadyzGetRequest struct via the builder pattern


### Return type

[**ReadinessResponse**](ReadinessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatusHealthStatusGet

> StatusResponse GetStatusHealthStatusGet(ctx).Execute()

Service stats endpoint



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GetStatusHealthStatusGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetStatusHealthStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatusHealthStatusGet`: StatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetStatusHealthStatusGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusHealthStatusGetRequest struct via the builder pattern


### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

No authorization required

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

