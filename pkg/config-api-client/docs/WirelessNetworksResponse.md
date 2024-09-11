# WirelessNetworksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WirelessNetworks** | [**[]WirelessNetwork**](WirelessNetwork.md) |  | 
**Pagination** | [**PaginationDetails**](PaginationDetails.md) |  | 

## Methods

### NewWirelessNetworksResponse

`func NewWirelessNetworksResponse(wirelessNetworks []WirelessNetwork, pagination PaginationDetails, ) *WirelessNetworksResponse`

NewWirelessNetworksResponse instantiates a new WirelessNetworksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessNetworksResponseWithDefaults

`func NewWirelessNetworksResponseWithDefaults() *WirelessNetworksResponse`

NewWirelessNetworksResponseWithDefaults instantiates a new WirelessNetworksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWirelessNetworks

`func (o *WirelessNetworksResponse) GetWirelessNetworks() []WirelessNetwork`

GetWirelessNetworks returns the WirelessNetworks field if non-nil, zero value otherwise.

### GetWirelessNetworksOk

`func (o *WirelessNetworksResponse) GetWirelessNetworksOk() (*[]WirelessNetwork, bool)`

GetWirelessNetworksOk returns a tuple with the WirelessNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessNetworks

`func (o *WirelessNetworksResponse) SetWirelessNetworks(v []WirelessNetwork)`

SetWirelessNetworks sets WirelessNetworks field to given value.


### GetPagination

`func (o *WirelessNetworksResponse) GetPagination() PaginationDetails`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *WirelessNetworksResponse) GetPaginationOk() (*PaginationDetails, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *WirelessNetworksResponse) SetPagination(v PaginationDetails)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


