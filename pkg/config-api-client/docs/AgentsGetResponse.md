# AgentsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AgentsGetItem**](AgentsGetItem.md) | The list of resources. | 
**Count** | **int32** | The number of resources returned in the response. | 
**Next** | **NullableString** |  | 

## Methods

### NewAgentsGetResponse

`func NewAgentsGetResponse(items []AgentsGetItem, count int32, next NullableString, ) *AgentsGetResponse`

NewAgentsGetResponse instantiates a new AgentsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsGetResponseWithDefaults

`func NewAgentsGetResponseWithDefaults() *AgentsGetResponse`

NewAgentsGetResponseWithDefaults instantiates a new AgentsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AgentsGetResponse) GetItems() []AgentsGetItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AgentsGetResponse) GetItemsOk() (*[]AgentsGetItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AgentsGetResponse) SetItems(v []AgentsGetItem)`

SetItems sets Items field to given value.


### GetCount

`func (o *AgentsGetResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AgentsGetResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AgentsGetResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *AgentsGetResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *AgentsGetResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *AgentsGetResponse) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *AgentsGetResponse) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *AgentsGetResponse) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


