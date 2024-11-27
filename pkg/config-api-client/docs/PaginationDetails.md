# PaginationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Next** | **NullableString** |  | 
**Previous** | **NullableString** |  | 
**First** | **NullableString** |  | 
**Last** | **NullableString** |  | 

## Methods

### NewPaginationDetails

`func NewPaginationDetails(limit int32, next NullableString, previous NullableString, first NullableString, last NullableString, ) *PaginationDetails`

NewPaginationDetails instantiates a new PaginationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationDetailsWithDefaults

`func NewPaginationDetailsWithDefaults() *PaginationDetails`

NewPaginationDetailsWithDefaults instantiates a new PaginationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *PaginationDetails) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaginationDetails) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaginationDetails) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *PaginationDetails) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginationDetails) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginationDetails) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *PaginationDetails) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginationDetails) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginationDetails) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginationDetails) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginationDetails) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### SetPreviousNil

`func (o *PaginationDetails) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginationDetails) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetFirst

`func (o *PaginationDetails) GetFirst() string`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PaginationDetails) GetFirstOk() (*string, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PaginationDetails) SetFirst(v string)`

SetFirst sets First field to given value.


### SetFirstNil

`func (o *PaginationDetails) SetFirstNil(b bool)`

 SetFirstNil sets the value for First to be an explicit nil

### UnsetFirst
`func (o *PaginationDetails) UnsetFirst()`

UnsetFirst ensures that no value is present for First, not even an explicit nil
### GetLast

`func (o *PaginationDetails) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PaginationDetails) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PaginationDetails) SetLast(v string)`

SetLast sets Last field to given value.


### SetLastNil

`func (o *PaginationDetails) SetLastNil(b bool)`

 SetLastNil sets the value for Last to be an explicit nil

### UnsetLast
`func (o *PaginationDetails) UnsetLast()`

UnsetLast ensures that no value is present for Last, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


