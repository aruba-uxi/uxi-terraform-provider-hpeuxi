# NetworkGroupAssignmentsPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Group** | [**Group**](Group.md) |  | 
**Network** | [**Network**](Network.md) |  | 
**Type** | Pointer to **string** |  | [optional] [default to "uxi/network-group-assignment"]

## Methods

### NewNetworkGroupAssignmentsPostResponse

`func NewNetworkGroupAssignmentsPostResponse(id string, group Group, network Network, ) *NetworkGroupAssignmentsPostResponse`

NewNetworkGroupAssignmentsPostResponse instantiates a new NetworkGroupAssignmentsPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkGroupAssignmentsPostResponseWithDefaults

`func NewNetworkGroupAssignmentsPostResponseWithDefaults() *NetworkGroupAssignmentsPostResponse`

NewNetworkGroupAssignmentsPostResponseWithDefaults instantiates a new NetworkGroupAssignmentsPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkGroupAssignmentsPostResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkGroupAssignmentsPostResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkGroupAssignmentsPostResponse) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *NetworkGroupAssignmentsPostResponse) GetGroup() Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *NetworkGroupAssignmentsPostResponse) GetGroupOk() (*Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *NetworkGroupAssignmentsPostResponse) SetGroup(v Group)`

SetGroup sets Group field to given value.


### GetNetwork

`func (o *NetworkGroupAssignmentsPostResponse) GetNetwork() Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkGroupAssignmentsPostResponse) GetNetworkOk() (*Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkGroupAssignmentsPostResponse) SetNetwork(v Network)`

SetNetwork sets Network field to given value.


### GetType

`func (o *NetworkGroupAssignmentsPostResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkGroupAssignmentsPostResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkGroupAssignmentsPostResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkGroupAssignmentsPostResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


