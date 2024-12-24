# AgentGroupAssignmentsGetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the agent group assignment | 
**Group** | [**AgentGroupAssignmentsGetGroup**](AgentGroupAssignmentsGetGroup.md) |  | 
**Agent** | [**AgentGroupAssignmentsGetAgent**](AgentGroupAssignmentsGetAgent.md) |  | 
**Type** | **string** | The type of the resource. | 

## Methods

### NewAgentGroupAssignmentsGetItem

`func NewAgentGroupAssignmentsGetItem(id string, group AgentGroupAssignmentsGetGroup, agent AgentGroupAssignmentsGetAgent, type_ string, ) *AgentGroupAssignmentsGetItem`

NewAgentGroupAssignmentsGetItem instantiates a new AgentGroupAssignmentsGetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentGroupAssignmentsGetItemWithDefaults

`func NewAgentGroupAssignmentsGetItemWithDefaults() *AgentGroupAssignmentsGetItem`

NewAgentGroupAssignmentsGetItemWithDefaults instantiates a new AgentGroupAssignmentsGetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentGroupAssignmentsGetItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentGroupAssignmentsGetItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentGroupAssignmentsGetItem) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *AgentGroupAssignmentsGetItem) GetGroup() AgentGroupAssignmentsGetGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AgentGroupAssignmentsGetItem) GetGroupOk() (*AgentGroupAssignmentsGetGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AgentGroupAssignmentsGetItem) SetGroup(v AgentGroupAssignmentsGetGroup)`

SetGroup sets Group field to given value.


### GetAgent

`func (o *AgentGroupAssignmentsGetItem) GetAgent() AgentGroupAssignmentsGetAgent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentGroupAssignmentsGetItem) GetAgentOk() (*AgentGroupAssignmentsGetAgent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentGroupAssignmentsGetItem) SetAgent(v AgentGroupAssignmentsGetAgent)`

SetAgent sets Agent field to given value.


### GetType

`func (o *AgentGroupAssignmentsGetItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentGroupAssignmentsGetItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentGroupAssignmentsGetItem) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


