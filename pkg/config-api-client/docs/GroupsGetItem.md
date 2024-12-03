# GroupsGetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Parent** | [**NullableGroupsGetParent**](GroupsGetParent.md) |  | 
**Path** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewGroupsGetItem

`func NewGroupsGetItem(id string, name string, parent NullableGroupsGetParent, path string, type_ string, ) *GroupsGetItem`

NewGroupsGetItem instantiates a new GroupsGetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsGetItemWithDefaults

`func NewGroupsGetItemWithDefaults() *GroupsGetItem`

NewGroupsGetItemWithDefaults instantiates a new GroupsGetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupsGetItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupsGetItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupsGetItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GroupsGetItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupsGetItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupsGetItem) SetName(v string)`

SetName sets Name field to given value.


### GetParent

`func (o *GroupsGetItem) GetParent() GroupsGetParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GroupsGetItem) GetParentOk() (*GroupsGetParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GroupsGetItem) SetParent(v GroupsGetParent)`

SetParent sets Parent field to given value.


### SetParentNil

`func (o *GroupsGetItem) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *GroupsGetItem) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetPath

`func (o *GroupsGetItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GroupsGetItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GroupsGetItem) SetPath(v string)`

SetPath sets Path field to given value.


### GetType

`func (o *GroupsGetItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupsGetItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupsGetItem) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


