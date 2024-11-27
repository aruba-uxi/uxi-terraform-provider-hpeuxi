# WirelessNetworksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]WirelessNetworksItem**](WirelessNetworksItem.md) |  | 
**Count** | **int32** |  | 
**Next** | **NullableString** |  | 

## Methods

### NewWirelessNetworksResponse

`func NewWirelessNetworksResponse(items []WirelessNetworksItem, count int32, next NullableString, ) *WirelessNetworksResponse`

NewWirelessNetworksResponse instantiates a new WirelessNetworksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessNetworksResponseWithDefaults

`func NewWirelessNetworksResponseWithDefaults() *WirelessNetworksResponse`

NewWirelessNetworksResponseWithDefaults instantiates a new WirelessNetworksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *WirelessNetworksResponse) GetItems() []WirelessNetworksItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WirelessNetworksResponse) GetItemsOk() (*[]WirelessNetworksItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WirelessNetworksResponse) SetItems(v []WirelessNetworksItem)`

SetItems sets Items field to given value.


### GetCount

`func (o *WirelessNetworksResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WirelessNetworksResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WirelessNetworksResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *WirelessNetworksResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *WirelessNetworksResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *WirelessNetworksResponse) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *WirelessNetworksResponse) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *WirelessNetworksResponse) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


