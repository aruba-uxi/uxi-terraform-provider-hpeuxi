# AgentPatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the agent | 
**Serial** | **string** | The serial number of the agent | 
**Name** | **string** | The name of the agent | 
**GroupName** | **NullableString** |  | 
**GroupPath** | **NullableString** |  | 
**ModelNumber** | **NullableString** |  | 
**WifiMacAddress** | **NullableString** |  | 
**EthernetMacAddress** | **NullableString** |  | 
**Notes** | **NullableString** |  | 
**PcapMode** | [**NullableAgentPcapMode**](AgentPcapMode.md) |  | 
**Type** | **string** | The type of the resource. | 

## Methods

### NewAgentPatchResponse

`func NewAgentPatchResponse(id string, serial string, name string, groupName NullableString, groupPath NullableString, modelNumber NullableString, wifiMacAddress NullableString, ethernetMacAddress NullableString, notes NullableString, pcapMode NullableAgentPcapMode, type_ string, ) *AgentPatchResponse`

NewAgentPatchResponse instantiates a new AgentPatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentPatchResponseWithDefaults

`func NewAgentPatchResponseWithDefaults() *AgentPatchResponse`

NewAgentPatchResponseWithDefaults instantiates a new AgentPatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentPatchResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentPatchResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentPatchResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSerial

`func (o *AgentPatchResponse) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *AgentPatchResponse) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *AgentPatchResponse) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetName

`func (o *AgentPatchResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentPatchResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentPatchResponse) SetName(v string)`

SetName sets Name field to given value.


### GetGroupName

`func (o *AgentPatchResponse) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *AgentPatchResponse) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *AgentPatchResponse) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### SetGroupNameNil

`func (o *AgentPatchResponse) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *AgentPatchResponse) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetGroupPath

`func (o *AgentPatchResponse) GetGroupPath() string`

GetGroupPath returns the GroupPath field if non-nil, zero value otherwise.

### GetGroupPathOk

`func (o *AgentPatchResponse) GetGroupPathOk() (*string, bool)`

GetGroupPathOk returns a tuple with the GroupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPath

`func (o *AgentPatchResponse) SetGroupPath(v string)`

SetGroupPath sets GroupPath field to given value.


### SetGroupPathNil

`func (o *AgentPatchResponse) SetGroupPathNil(b bool)`

 SetGroupPathNil sets the value for GroupPath to be an explicit nil

### UnsetGroupPath
`func (o *AgentPatchResponse) UnsetGroupPath()`

UnsetGroupPath ensures that no value is present for GroupPath, not even an explicit nil
### GetModelNumber

`func (o *AgentPatchResponse) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *AgentPatchResponse) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *AgentPatchResponse) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.


### SetModelNumberNil

`func (o *AgentPatchResponse) SetModelNumberNil(b bool)`

 SetModelNumberNil sets the value for ModelNumber to be an explicit nil

### UnsetModelNumber
`func (o *AgentPatchResponse) UnsetModelNumber()`

UnsetModelNumber ensures that no value is present for ModelNumber, not even an explicit nil
### GetWifiMacAddress

`func (o *AgentPatchResponse) GetWifiMacAddress() string`

GetWifiMacAddress returns the WifiMacAddress field if non-nil, zero value otherwise.

### GetWifiMacAddressOk

`func (o *AgentPatchResponse) GetWifiMacAddressOk() (*string, bool)`

GetWifiMacAddressOk returns a tuple with the WifiMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacAddress

`func (o *AgentPatchResponse) SetWifiMacAddress(v string)`

SetWifiMacAddress sets WifiMacAddress field to given value.


### SetWifiMacAddressNil

`func (o *AgentPatchResponse) SetWifiMacAddressNil(b bool)`

 SetWifiMacAddressNil sets the value for WifiMacAddress to be an explicit nil

### UnsetWifiMacAddress
`func (o *AgentPatchResponse) UnsetWifiMacAddress()`

UnsetWifiMacAddress ensures that no value is present for WifiMacAddress, not even an explicit nil
### GetEthernetMacAddress

`func (o *AgentPatchResponse) GetEthernetMacAddress() string`

GetEthernetMacAddress returns the EthernetMacAddress field if non-nil, zero value otherwise.

### GetEthernetMacAddressOk

`func (o *AgentPatchResponse) GetEthernetMacAddressOk() (*string, bool)`

GetEthernetMacAddressOk returns a tuple with the EthernetMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetMacAddress

`func (o *AgentPatchResponse) SetEthernetMacAddress(v string)`

SetEthernetMacAddress sets EthernetMacAddress field to given value.


### SetEthernetMacAddressNil

`func (o *AgentPatchResponse) SetEthernetMacAddressNil(b bool)`

 SetEthernetMacAddressNil sets the value for EthernetMacAddress to be an explicit nil

### UnsetEthernetMacAddress
`func (o *AgentPatchResponse) UnsetEthernetMacAddress()`

UnsetEthernetMacAddress ensures that no value is present for EthernetMacAddress, not even an explicit nil
### GetNotes

`func (o *AgentPatchResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AgentPatchResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AgentPatchResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.


### SetNotesNil

`func (o *AgentPatchResponse) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *AgentPatchResponse) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetPcapMode

`func (o *AgentPatchResponse) GetPcapMode() AgentPcapMode`

GetPcapMode returns the PcapMode field if non-nil, zero value otherwise.

### GetPcapModeOk

`func (o *AgentPatchResponse) GetPcapModeOk() (*AgentPcapMode, bool)`

GetPcapModeOk returns a tuple with the PcapMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcapMode

`func (o *AgentPatchResponse) SetPcapMode(v AgentPcapMode)`

SetPcapMode sets PcapMode field to given value.


### SetPcapModeNil

`func (o *AgentPatchResponse) SetPcapModeNil(b bool)`

 SetPcapModeNil sets the value for PcapMode to be an explicit nil

### UnsetPcapMode
`func (o *AgentPatchResponse) UnsetPcapMode()`

UnsetPcapMode ensures that no value is present for PcapMode, not even an explicit nil
### GetType

`func (o *AgentPatchResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentPatchResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentPatchResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


