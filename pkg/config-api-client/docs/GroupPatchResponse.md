# GroupPatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the group | 
**Name** | **string** | The name of the group | 
**Path** | **string** | The path of the group | 
**Parent** | [**GroupPatchParent**](GroupPatchParent.md) | The parent group | 
**Type** | **string** | The type of the resource. | 

## Methods

### NewGroupPatchResponse

`func NewGroupPatchResponse(id string, name string, path string, parent GroupPatchParent, type_ string, ) *GroupPatchResponse`

NewGroupPatchResponse instantiates a new GroupPatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPatchResponseWithDefaults

`func NewGroupPatchResponseWithDefaults() *GroupPatchResponse`

NewGroupPatchResponseWithDefaults instantiates a new GroupPatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupPatchResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupPatchResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupPatchResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GroupPatchResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupPatchResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupPatchResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *GroupPatchResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GroupPatchResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GroupPatchResponse) SetPath(v string)`

SetPath sets Path field to given value.


### GetParent

`func (o *GroupPatchResponse) GetParent() GroupPatchParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GroupPatchResponse) GetParentOk() (*GroupPatchParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GroupPatchResponse) SetParent(v GroupPatchParent)`

SetParent sets Parent field to given value.


### GetType

`func (o *GroupPatchResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupPatchResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupPatchResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


