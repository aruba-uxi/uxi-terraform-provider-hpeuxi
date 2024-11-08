# ServiceTestGroupAssignmentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ServiceTestGroupAssignmentsItem**](ServiceTestGroupAssignmentsItem.md) |  | 
**Count** | **int32** |  | 
**Next** | **NullableString** |  | 

## Methods

### NewServiceTestGroupAssignmentsResponse

`func NewServiceTestGroupAssignmentsResponse(items []ServiceTestGroupAssignmentsItem, count int32, next NullableString, ) *ServiceTestGroupAssignmentsResponse`

NewServiceTestGroupAssignmentsResponse instantiates a new ServiceTestGroupAssignmentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTestGroupAssignmentsResponseWithDefaults

`func NewServiceTestGroupAssignmentsResponseWithDefaults() *ServiceTestGroupAssignmentsResponse`

NewServiceTestGroupAssignmentsResponseWithDefaults instantiates a new ServiceTestGroupAssignmentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ServiceTestGroupAssignmentsResponse) GetItems() []ServiceTestGroupAssignmentsItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ServiceTestGroupAssignmentsResponse) GetItemsOk() (*[]ServiceTestGroupAssignmentsItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ServiceTestGroupAssignmentsResponse) SetItems(v []ServiceTestGroupAssignmentsItem)`

SetItems sets Items field to given value.


### GetCount

`func (o *ServiceTestGroupAssignmentsResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ServiceTestGroupAssignmentsResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ServiceTestGroupAssignmentsResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *ServiceTestGroupAssignmentsResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ServiceTestGroupAssignmentsResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ServiceTestGroupAssignmentsResponse) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *ServiceTestGroupAssignmentsResponse) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *ServiceTestGroupAssignmentsResponse) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


