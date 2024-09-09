# GroupsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**PaginationDetails**](PaginationDetails.md) |  | 
**Groups** | [**[]GroupsGetItem**](GroupsGetItem.md) |  | 

## Methods

### NewGroupsGetResponse

`func NewGroupsGetResponse(pagination PaginationDetails, groups []GroupsGetItem, ) *GroupsGetResponse`

NewGroupsGetResponse instantiates a new GroupsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsGetResponseWithDefaults

`func NewGroupsGetResponseWithDefaults() *GroupsGetResponse`

NewGroupsGetResponseWithDefaults instantiates a new GroupsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *GroupsGetResponse) GetPagination() PaginationDetails`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GroupsGetResponse) GetPaginationOk() (*PaginationDetails, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GroupsGetResponse) SetPagination(v PaginationDetails)`

SetPagination sets Pagination field to given value.


### GetGroups

`func (o *GroupsGetResponse) GetGroups() []GroupsGetItem`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GroupsGetResponse) GetGroupsOk() (*[]GroupsGetItem, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GroupsGetResponse) SetGroups(v []GroupsGetItem)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


