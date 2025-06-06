# NetworkGroupAssignmentPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the network group assignment | 
**Group** | [**NetworkGroupAssignmentPostGroup**](NetworkGroupAssignmentPostGroup.md) | The group component of the network group assignment | 
**Network** | [**NetworkGroupAssignmentPostNetwork**](NetworkGroupAssignmentPostNetwork.md) | The network component of the network group assignment | 
**Type** | **string** | The type of the resource. | 

## Methods

### NewNetworkGroupAssignmentPostResponse

`func NewNetworkGroupAssignmentPostResponse(id string, group NetworkGroupAssignmentPostGroup, network NetworkGroupAssignmentPostNetwork, type_ string, ) *NetworkGroupAssignmentPostResponse`

NewNetworkGroupAssignmentPostResponse instantiates a new NetworkGroupAssignmentPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkGroupAssignmentPostResponseWithDefaults

`func NewNetworkGroupAssignmentPostResponseWithDefaults() *NetworkGroupAssignmentPostResponse`

NewNetworkGroupAssignmentPostResponseWithDefaults instantiates a new NetworkGroupAssignmentPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkGroupAssignmentPostResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkGroupAssignmentPostResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkGroupAssignmentPostResponse) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *NetworkGroupAssignmentPostResponse) GetGroup() NetworkGroupAssignmentPostGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *NetworkGroupAssignmentPostResponse) GetGroupOk() (*NetworkGroupAssignmentPostGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *NetworkGroupAssignmentPostResponse) SetGroup(v NetworkGroupAssignmentPostGroup)`

SetGroup sets Group field to given value.


### GetNetwork

`func (o *NetworkGroupAssignmentPostResponse) GetNetwork() NetworkGroupAssignmentPostNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkGroupAssignmentPostResponse) GetNetworkOk() (*NetworkGroupAssignmentPostNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkGroupAssignmentPostResponse) SetNetwork(v NetworkGroupAssignmentPostNetwork)`

SetNetwork sets Network field to given value.


### GetType

`func (o *NetworkGroupAssignmentPostResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkGroupAssignmentPostResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkGroupAssignmentPostResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


