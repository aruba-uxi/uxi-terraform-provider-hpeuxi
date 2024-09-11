# WiredNetworksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WiredNetworks** | [**[]WiredNetwork**](WiredNetwork.md) |  | 
**Pagination** | [**PaginationDetails**](PaginationDetails.md) |  | 

## Methods

### NewWiredNetworksResponse

`func NewWiredNetworksResponse(wiredNetworks []WiredNetwork, pagination PaginationDetails, ) *WiredNetworksResponse`

NewWiredNetworksResponse instantiates a new WiredNetworksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWiredNetworksResponseWithDefaults

`func NewWiredNetworksResponseWithDefaults() *WiredNetworksResponse`

NewWiredNetworksResponseWithDefaults instantiates a new WiredNetworksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWiredNetworks

`func (o *WiredNetworksResponse) GetWiredNetworks() []WiredNetwork`

GetWiredNetworks returns the WiredNetworks field if non-nil, zero value otherwise.

### GetWiredNetworksOk

`func (o *WiredNetworksResponse) GetWiredNetworksOk() (*[]WiredNetwork, bool)`

GetWiredNetworksOk returns a tuple with the WiredNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredNetworks

`func (o *WiredNetworksResponse) SetWiredNetworks(v []WiredNetwork)`

SetWiredNetworks sets WiredNetworks field to given value.


### GetPagination

`func (o *WiredNetworksResponse) GetPagination() PaginationDetails`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *WiredNetworksResponse) GetPaginationOk() (*PaginationDetails, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *WiredNetworksResponse) SetPagination(v PaginationDetails)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


