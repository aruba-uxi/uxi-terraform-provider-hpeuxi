# AgentPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the agent | [optional] 
**Notes** | Pointer to **string** | The notes of the agent | [optional] 
**PcapMode** | Pointer to [**AgentPcapMode**](AgentPcapMode.md) |  | [optional] 

## Methods

### NewAgentPatchRequest

`func NewAgentPatchRequest() *AgentPatchRequest`

NewAgentPatchRequest instantiates a new AgentPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentPatchRequestWithDefaults

`func NewAgentPatchRequestWithDefaults() *AgentPatchRequest`

NewAgentPatchRequestWithDefaults instantiates a new AgentPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AgentPatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentPatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentPatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentPatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *AgentPatchRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AgentPatchRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AgentPatchRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AgentPatchRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPcapMode

`func (o *AgentPatchRequest) GetPcapMode() AgentPcapMode`

GetPcapMode returns the PcapMode field if non-nil, zero value otherwise.

### GetPcapModeOk

`func (o *AgentPatchRequest) GetPcapModeOk() (*AgentPcapMode, bool)`

GetPcapModeOk returns a tuple with the PcapMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcapMode

`func (o *AgentPatchRequest) SetPcapMode(v AgentPcapMode)`

SetPcapMode sets PcapMode field to given value.

### HasPcapMode

`func (o *AgentPatchRequest) HasPcapMode() bool`

HasPcapMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


