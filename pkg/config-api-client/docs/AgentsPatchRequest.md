# AgentsPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**PcapMode** | Pointer to [**PcapMode**](PcapMode.md) |  | [optional] 

## Methods

### NewAgentsPatchRequest

`func NewAgentsPatchRequest() *AgentsPatchRequest`

NewAgentsPatchRequest instantiates a new AgentsPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsPatchRequestWithDefaults

`func NewAgentsPatchRequestWithDefaults() *AgentsPatchRequest`

NewAgentsPatchRequestWithDefaults instantiates a new AgentsPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AgentsPatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentsPatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentsPatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentsPatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *AgentsPatchRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AgentsPatchRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AgentsPatchRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AgentsPatchRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPcapMode

`func (o *AgentsPatchRequest) GetPcapMode() PcapMode`

GetPcapMode returns the PcapMode field if non-nil, zero value otherwise.

### GetPcapModeOk

`func (o *AgentsPatchRequest) GetPcapModeOk() (*PcapMode, bool)`

GetPcapModeOk returns a tuple with the PcapMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcapMode

`func (o *AgentsPatchRequest) SetPcapMode(v PcapMode)`

SetPcapMode sets PcapMode field to given value.

### HasPcapMode

`func (o *AgentsPatchRequest) HasPcapMode() bool`

HasPcapMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


