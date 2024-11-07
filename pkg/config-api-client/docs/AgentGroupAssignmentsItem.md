# AgentGroupAssignmentsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Group** | [**Group**](Group.md) |  | 
**Agent** | [**Agent**](Agent.md) |  | 
**Type** | **string** |  | 

## Methods

### NewAgentGroupAssignmentsItem

`func NewAgentGroupAssignmentsItem(id string, group Group, agent Agent, type_ string, ) *AgentGroupAssignmentsItem`

NewAgentGroupAssignmentsItem instantiates a new AgentGroupAssignmentsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentGroupAssignmentsItemWithDefaults

`func NewAgentGroupAssignmentsItemWithDefaults() *AgentGroupAssignmentsItem`

NewAgentGroupAssignmentsItemWithDefaults instantiates a new AgentGroupAssignmentsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentGroupAssignmentsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentGroupAssignmentsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentGroupAssignmentsItem) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *AgentGroupAssignmentsItem) GetGroup() Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AgentGroupAssignmentsItem) GetGroupOk() (*Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AgentGroupAssignmentsItem) SetGroup(v Group)`

SetGroup sets Group field to given value.


### GetAgent

`func (o *AgentGroupAssignmentsItem) GetAgent() Agent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentGroupAssignmentsItem) GetAgentOk() (*Agent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentGroupAssignmentsItem) SetAgent(v Agent)`

SetAgent sets Agent field to given value.


### GetType

`func (o *AgentGroupAssignmentsItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentGroupAssignmentsItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentGroupAssignmentsItem) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


