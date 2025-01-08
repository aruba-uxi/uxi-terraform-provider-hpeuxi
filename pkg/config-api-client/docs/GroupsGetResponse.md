# GroupsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]GroupsGetItem**](GroupsGetItem.md) | The list of resources. | 
**Count** | **int32** | The number of resources returned in the response. | 
**Next** | **NullableString** | The next cursor for pagination. | 

## Methods

### NewGroupsGetResponse

`func NewGroupsGetResponse(items []GroupsGetItem, count int32, next NullableString, ) *GroupsGetResponse`

NewGroupsGetResponse instantiates a new GroupsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsGetResponseWithDefaults

`func NewGroupsGetResponseWithDefaults() *GroupsGetResponse`

NewGroupsGetResponseWithDefaults instantiates a new GroupsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GroupsGetResponse) GetItems() []GroupsGetItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GroupsGetResponse) GetItemsOk() (*[]GroupsGetItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GroupsGetResponse) SetItems(v []GroupsGetItem)`

SetItems sets Items field to given value.


### GetCount

`func (o *GroupsGetResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GroupsGetResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GroupsGetResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *GroupsGetResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GroupsGetResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GroupsGetResponse) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *GroupsGetResponse) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *GroupsGetResponse) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


