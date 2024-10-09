# GroupsPatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Path** | **string** |  | 
**Parent** | [**Parent**](Parent.md) |  | 
**Type** | Pointer to **string** |  | [optional] [default to "uxi/group"]

## Methods

### NewGroupsPatchResponse

`func NewGroupsPatchResponse(id string, name string, path string, parent Parent, ) *GroupsPatchResponse`

NewGroupsPatchResponse instantiates a new GroupsPatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsPatchResponseWithDefaults

`func NewGroupsPatchResponseWithDefaults() *GroupsPatchResponse`

NewGroupsPatchResponseWithDefaults instantiates a new GroupsPatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupsPatchResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupsPatchResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupsPatchResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GroupsPatchResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupsPatchResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupsPatchResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *GroupsPatchResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GroupsPatchResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GroupsPatchResponse) SetPath(v string)`

SetPath sets Path field to given value.


### GetParent

`func (o *GroupsPatchResponse) GetParent() Parent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GroupsPatchResponse) GetParentOk() (*Parent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GroupsPatchResponse) SetParent(v Parent)`

SetParent sets Parent field to given value.


### GetType

`func (o *GroupsPatchResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupsPatchResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupsPatchResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GroupsPatchResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


