# ServiceTestsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ServiceTestsListItem**](ServiceTestsListItem.md) |  | 
**Count** | **int32** |  | 
**Next** | **NullableString** |  | 

## Methods

### NewServiceTestsListResponse

`func NewServiceTestsListResponse(items []ServiceTestsListItem, count int32, next NullableString, ) *ServiceTestsListResponse`

NewServiceTestsListResponse instantiates a new ServiceTestsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTestsListResponseWithDefaults

`func NewServiceTestsListResponseWithDefaults() *ServiceTestsListResponse`

NewServiceTestsListResponseWithDefaults instantiates a new ServiceTestsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ServiceTestsListResponse) GetItems() []ServiceTestsListItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ServiceTestsListResponse) GetItemsOk() (*[]ServiceTestsListItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ServiceTestsListResponse) SetItems(v []ServiceTestsListItem)`

SetItems sets Items field to given value.


### GetCount

`func (o *ServiceTestsListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ServiceTestsListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ServiceTestsListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *ServiceTestsListResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ServiceTestsListResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ServiceTestsListResponse) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *ServiceTestsListResponse) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *ServiceTestsListResponse) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


