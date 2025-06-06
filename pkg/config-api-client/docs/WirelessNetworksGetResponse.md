# WirelessNetworksGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]WirelessNetworksGetItem**](WirelessNetworksGetItem.md) | The list of resources. | 
**Count** | **int32** | The number of resources returned in the response. | 
**Next** | **NullableString** |  | 

## Methods

### NewWirelessNetworksGetResponse

`func NewWirelessNetworksGetResponse(items []WirelessNetworksGetItem, count int32, next NullableString, ) *WirelessNetworksGetResponse`

NewWirelessNetworksGetResponse instantiates a new WirelessNetworksGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessNetworksGetResponseWithDefaults

`func NewWirelessNetworksGetResponseWithDefaults() *WirelessNetworksGetResponse`

NewWirelessNetworksGetResponseWithDefaults instantiates a new WirelessNetworksGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *WirelessNetworksGetResponse) GetItems() []WirelessNetworksGetItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WirelessNetworksGetResponse) GetItemsOk() (*[]WirelessNetworksGetItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WirelessNetworksGetResponse) SetItems(v []WirelessNetworksGetItem)`

SetItems sets Items field to given value.


### GetCount

`func (o *WirelessNetworksGetResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WirelessNetworksGetResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WirelessNetworksGetResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *WirelessNetworksGetResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *WirelessNetworksGetResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *WirelessNetworksGetResponse) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *WirelessNetworksGetResponse) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *WirelessNetworksGetResponse) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


