# NetworkGroupAssignmentsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkGroupAssignments** | [**[]NetworkGroupAssignmentsItem**](NetworkGroupAssignmentsItem.md) |  | 
**Pagination** | [**PaginationDetails**](PaginationDetails.md) |  | 

## Methods

### NewNetworkGroupAssignmentsGetResponse

`func NewNetworkGroupAssignmentsGetResponse(networkGroupAssignments []NetworkGroupAssignmentsItem, pagination PaginationDetails, ) *NetworkGroupAssignmentsGetResponse`

NewNetworkGroupAssignmentsGetResponse instantiates a new NetworkGroupAssignmentsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkGroupAssignmentsGetResponseWithDefaults

`func NewNetworkGroupAssignmentsGetResponseWithDefaults() *NetworkGroupAssignmentsGetResponse`

NewNetworkGroupAssignmentsGetResponseWithDefaults instantiates a new NetworkGroupAssignmentsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkGroupAssignments

`func (o *NetworkGroupAssignmentsGetResponse) GetNetworkGroupAssignments() []NetworkGroupAssignmentsItem`

GetNetworkGroupAssignments returns the NetworkGroupAssignments field if non-nil, zero value otherwise.

### GetNetworkGroupAssignmentsOk

`func (o *NetworkGroupAssignmentsGetResponse) GetNetworkGroupAssignmentsOk() (*[]NetworkGroupAssignmentsItem, bool)`

GetNetworkGroupAssignmentsOk returns a tuple with the NetworkGroupAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGroupAssignments

`func (o *NetworkGroupAssignmentsGetResponse) SetNetworkGroupAssignments(v []NetworkGroupAssignmentsItem)`

SetNetworkGroupAssignments sets NetworkGroupAssignments field to given value.


### GetPagination

`func (o *NetworkGroupAssignmentsGetResponse) GetPagination() PaginationDetails`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *NetworkGroupAssignmentsGetResponse) GetPaginationOk() (*PaginationDetails, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *NetworkGroupAssignmentsGetResponse) SetPagination(v PaginationDetails)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


