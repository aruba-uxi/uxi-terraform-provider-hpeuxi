# NetworkGroupAssignmentsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Group** | [**Group**](Group.md) |  | 
**Network** | [**Network**](Network.md) |  | 
**Type** | **string** |  | 

## Methods

### NewNetworkGroupAssignmentsItem

`func NewNetworkGroupAssignmentsItem(id string, group Group, network Network, type_ string, ) *NetworkGroupAssignmentsItem`

NewNetworkGroupAssignmentsItem instantiates a new NetworkGroupAssignmentsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkGroupAssignmentsItemWithDefaults

`func NewNetworkGroupAssignmentsItemWithDefaults() *NetworkGroupAssignmentsItem`

NewNetworkGroupAssignmentsItemWithDefaults instantiates a new NetworkGroupAssignmentsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkGroupAssignmentsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkGroupAssignmentsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkGroupAssignmentsItem) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *NetworkGroupAssignmentsItem) GetGroup() Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *NetworkGroupAssignmentsItem) GetGroupOk() (*Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *NetworkGroupAssignmentsItem) SetGroup(v Group)`

SetGroup sets Group field to given value.


### GetNetwork

`func (o *NetworkGroupAssignmentsItem) GetNetwork() Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkGroupAssignmentsItem) GetNetworkOk() (*Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkGroupAssignmentsItem) SetNetwork(v Network)`

SetNetwork sets Network field to given value.


### GetType

`func (o *NetworkGroupAssignmentsItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkGroupAssignmentsItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkGroupAssignmentsItem) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


