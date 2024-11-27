# AgentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AgentItem**](AgentItem.md) |  | 
**Count** | **int32** |  | 
**Next** | **NullableString** |  | 

## Methods

### NewAgentsResponse

`func NewAgentsResponse(items []AgentItem, count int32, next NullableString, ) *AgentsResponse`

NewAgentsResponse instantiates a new AgentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsResponseWithDefaults

`func NewAgentsResponseWithDefaults() *AgentsResponse`

NewAgentsResponseWithDefaults instantiates a new AgentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AgentsResponse) GetItems() []AgentItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AgentsResponse) GetItemsOk() (*[]AgentItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AgentsResponse) SetItems(v []AgentItem)`

SetItems sets Items field to given value.


### GetCount

`func (o *AgentsResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AgentsResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AgentsResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *AgentsResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *AgentsResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *AgentsResponse) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *AgentsResponse) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *AgentsResponse) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


