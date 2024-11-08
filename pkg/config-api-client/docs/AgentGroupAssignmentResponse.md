# AgentGroupAssignmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Group** | [**Group**](Group.md) |  | 
**Agent** | [**Agent**](Agent.md) |  | 
**Type** | **string** |  | 

## Methods

### NewAgentGroupAssignmentResponse

`func NewAgentGroupAssignmentResponse(id string, group Group, agent Agent, type_ string, ) *AgentGroupAssignmentResponse`

NewAgentGroupAssignmentResponse instantiates a new AgentGroupAssignmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentGroupAssignmentResponseWithDefaults

`func NewAgentGroupAssignmentResponseWithDefaults() *AgentGroupAssignmentResponse`

NewAgentGroupAssignmentResponseWithDefaults instantiates a new AgentGroupAssignmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentGroupAssignmentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentGroupAssignmentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentGroupAssignmentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *AgentGroupAssignmentResponse) GetGroup() Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AgentGroupAssignmentResponse) GetGroupOk() (*Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AgentGroupAssignmentResponse) SetGroup(v Group)`

SetGroup sets Group field to given value.


### GetAgent

`func (o *AgentGroupAssignmentResponse) GetAgent() Agent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentGroupAssignmentResponse) GetAgentOk() (*Agent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentGroupAssignmentResponse) SetAgent(v Agent)`

SetAgent sets Agent field to given value.


### GetType

`func (o *AgentGroupAssignmentResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentGroupAssignmentResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentGroupAssignmentResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


