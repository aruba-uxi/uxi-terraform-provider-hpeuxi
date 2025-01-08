# AgentGroupAssignmentsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AgentGroupAssignmentsGetItem**](AgentGroupAssignmentsGetItem.md) | The list of resources. | 
**Count** | **int32** | The number of resources returned in the response. | 
**Next** | **NullableString** | The next cursor for pagination. | 

## Methods

### NewAgentGroupAssignmentsGetResponse

`func NewAgentGroupAssignmentsGetResponse(items []AgentGroupAssignmentsGetItem, count int32, next NullableString, ) *AgentGroupAssignmentsGetResponse`

NewAgentGroupAssignmentsGetResponse instantiates a new AgentGroupAssignmentsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentGroupAssignmentsGetResponseWithDefaults

`func NewAgentGroupAssignmentsGetResponseWithDefaults() *AgentGroupAssignmentsGetResponse`

NewAgentGroupAssignmentsGetResponseWithDefaults instantiates a new AgentGroupAssignmentsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AgentGroupAssignmentsGetResponse) GetItems() []AgentGroupAssignmentsGetItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AgentGroupAssignmentsGetResponse) GetItemsOk() (*[]AgentGroupAssignmentsGetItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AgentGroupAssignmentsGetResponse) SetItems(v []AgentGroupAssignmentsGetItem)`

SetItems sets Items field to given value.


### GetCount

`func (o *AgentGroupAssignmentsGetResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AgentGroupAssignmentsGetResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AgentGroupAssignmentsGetResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *AgentGroupAssignmentsGetResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *AgentGroupAssignmentsGetResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *AgentGroupAssignmentsGetResponse) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *AgentGroupAssignmentsGetResponse) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *AgentGroupAssignmentsGetResponse) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


