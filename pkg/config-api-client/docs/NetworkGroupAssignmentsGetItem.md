# NetworkGroupAssignmentsGetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Group** | [**NetworkGroupAssignmentsGetGroup**](NetworkGroupAssignmentsGetGroup.md) |  | 
**Network** | [**NetworkGroupAssignmentsGetNetwork**](NetworkGroupAssignmentsGetNetwork.md) |  | 
**Type** | **string** |  | 

## Methods

### NewNetworkGroupAssignmentsGetItem

`func NewNetworkGroupAssignmentsGetItem(id string, group NetworkGroupAssignmentsGetGroup, network NetworkGroupAssignmentsGetNetwork, type_ string, ) *NetworkGroupAssignmentsGetItem`

NewNetworkGroupAssignmentsGetItem instantiates a new NetworkGroupAssignmentsGetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkGroupAssignmentsGetItemWithDefaults

`func NewNetworkGroupAssignmentsGetItemWithDefaults() *NetworkGroupAssignmentsGetItem`

NewNetworkGroupAssignmentsGetItemWithDefaults instantiates a new NetworkGroupAssignmentsGetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkGroupAssignmentsGetItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkGroupAssignmentsGetItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkGroupAssignmentsGetItem) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *NetworkGroupAssignmentsGetItem) GetGroup() NetworkGroupAssignmentsGetGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *NetworkGroupAssignmentsGetItem) GetGroupOk() (*NetworkGroupAssignmentsGetGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *NetworkGroupAssignmentsGetItem) SetGroup(v NetworkGroupAssignmentsGetGroup)`

SetGroup sets Group field to given value.


### GetNetwork

`func (o *NetworkGroupAssignmentsGetItem) GetNetwork() NetworkGroupAssignmentsGetNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkGroupAssignmentsGetItem) GetNetworkOk() (*NetworkGroupAssignmentsGetNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkGroupAssignmentsGetItem) SetNetwork(v NetworkGroupAssignmentsGetNetwork)`

SetNetwork sets Network field to given value.


### GetType

`func (o *NetworkGroupAssignmentsGetItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkGroupAssignmentsGetItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkGroupAssignmentsGetItem) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


