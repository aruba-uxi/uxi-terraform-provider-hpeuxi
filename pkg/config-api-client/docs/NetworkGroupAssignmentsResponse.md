# NetworkGroupAssignmentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkGroupAssignments** | [**[]NetworkGroupAssignmentsItem**](NetworkGroupAssignmentsItem.md) |  | 
**Pagination** | [**PaginationDetails**](PaginationDetails.md) |  | 

## Methods

### NewNetworkGroupAssignmentsResponse

`func NewNetworkGroupAssignmentsResponse(networkGroupAssignments []NetworkGroupAssignmentsItem, pagination PaginationDetails, ) *NetworkGroupAssignmentsResponse`

NewNetworkGroupAssignmentsResponse instantiates a new NetworkGroupAssignmentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkGroupAssignmentsResponseWithDefaults

`func NewNetworkGroupAssignmentsResponseWithDefaults() *NetworkGroupAssignmentsResponse`

NewNetworkGroupAssignmentsResponseWithDefaults instantiates a new NetworkGroupAssignmentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkGroupAssignments

`func (o *NetworkGroupAssignmentsResponse) GetNetworkGroupAssignments() []NetworkGroupAssignmentsItem`

GetNetworkGroupAssignments returns the NetworkGroupAssignments field if non-nil, zero value otherwise.

### GetNetworkGroupAssignmentsOk

`func (o *NetworkGroupAssignmentsResponse) GetNetworkGroupAssignmentsOk() (*[]NetworkGroupAssignmentsItem, bool)`

GetNetworkGroupAssignmentsOk returns a tuple with the NetworkGroupAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGroupAssignments

`func (o *NetworkGroupAssignmentsResponse) SetNetworkGroupAssignments(v []NetworkGroupAssignmentsItem)`

SetNetworkGroupAssignments sets NetworkGroupAssignments field to given value.


### GetPagination

`func (o *NetworkGroupAssignmentsResponse) GetPagination() PaginationDetails`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *NetworkGroupAssignmentsResponse) GetPaginationOk() (*PaginationDetails, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *NetworkGroupAssignmentsResponse) SetPagination(v PaginationDetails)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


