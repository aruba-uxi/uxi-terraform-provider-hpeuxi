# SensorsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]SensorsGetItem**](SensorsGetItem.md) |  | 
**Count** | **int32** |  | 
**Next** | **NullableString** |  | 

## Methods

### NewSensorsGetResponse

`func NewSensorsGetResponse(items []SensorsGetItem, count int32, next NullableString, ) *SensorsGetResponse`

NewSensorsGetResponse instantiates a new SensorsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorsGetResponseWithDefaults

`func NewSensorsGetResponseWithDefaults() *SensorsGetResponse`

NewSensorsGetResponseWithDefaults instantiates a new SensorsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SensorsGetResponse) GetItems() []SensorsGetItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SensorsGetResponse) GetItemsOk() (*[]SensorsGetItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SensorsGetResponse) SetItems(v []SensorsGetItem)`

SetItems sets Items field to given value.


### GetCount

`func (o *SensorsGetResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SensorsGetResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SensorsGetResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *SensorsGetResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *SensorsGetResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *SensorsGetResponse) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *SensorsGetResponse) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *SensorsGetResponse) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


