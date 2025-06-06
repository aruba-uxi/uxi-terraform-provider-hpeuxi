# AgentGroupAssignmentPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the agent group assignment | 
**Group** | [**AgentGroupAssignmentPostGroup**](AgentGroupAssignmentPostGroup.md) | The group component of the agent group assignment | 
**Agent** | [**AgentGroupAssignmentPostAgent**](AgentGroupAssignmentPostAgent.md) | The agent component of the agent group assignment | 
**Type** | **string** | The type of the resource. | 

## Methods

### NewAgentGroupAssignmentPostResponse

`func NewAgentGroupAssignmentPostResponse(id string, group AgentGroupAssignmentPostGroup, agent AgentGroupAssignmentPostAgent, type_ string, ) *AgentGroupAssignmentPostResponse`

NewAgentGroupAssignmentPostResponse instantiates a new AgentGroupAssignmentPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentGroupAssignmentPostResponseWithDefaults

`func NewAgentGroupAssignmentPostResponseWithDefaults() *AgentGroupAssignmentPostResponse`

NewAgentGroupAssignmentPostResponseWithDefaults instantiates a new AgentGroupAssignmentPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentGroupAssignmentPostResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentGroupAssignmentPostResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentGroupAssignmentPostResponse) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *AgentGroupAssignmentPostResponse) GetGroup() AgentGroupAssignmentPostGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AgentGroupAssignmentPostResponse) GetGroupOk() (*AgentGroupAssignmentPostGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AgentGroupAssignmentPostResponse) SetGroup(v AgentGroupAssignmentPostGroup)`

SetGroup sets Group field to given value.


### GetAgent

`func (o *AgentGroupAssignmentPostResponse) GetAgent() AgentGroupAssignmentPostAgent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentGroupAssignmentPostResponse) GetAgentOk() (*AgentGroupAssignmentPostAgent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentGroupAssignmentPostResponse) SetAgent(v AgentGroupAssignmentPostAgent)`

SetAgent sets Agent field to given value.


### GetType

`func (o *AgentGroupAssignmentPostResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentGroupAssignmentPostResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentGroupAssignmentPostResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


