# AgentsGetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Serial** | **string** |  | 
**Name** | **string** |  | 
**ModelNumber** | **NullableString** |  | 
**WifiMacAddress** | **NullableString** |  | 
**EthernetMacAddress** | **NullableString** |  | 
**Notes** | **NullableString** |  | 
**PcapMode** | [**NullableAgentPcapMode**](AgentPcapMode.md) |  | 
**Type** | **string** |  | 

## Methods

### NewAgentsGetItem

`func NewAgentsGetItem(id string, serial string, name string, modelNumber NullableString, wifiMacAddress NullableString, ethernetMacAddress NullableString, notes NullableString, pcapMode NullableAgentPcapMode, type_ string, ) *AgentsGetItem`

NewAgentsGetItem instantiates a new AgentsGetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsGetItemWithDefaults

`func NewAgentsGetItemWithDefaults() *AgentsGetItem`

NewAgentsGetItemWithDefaults instantiates a new AgentsGetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentsGetItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentsGetItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentsGetItem) SetId(v string)`

SetId sets Id field to given value.


### GetSerial

`func (o *AgentsGetItem) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *AgentsGetItem) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *AgentsGetItem) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetName

`func (o *AgentsGetItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentsGetItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentsGetItem) SetName(v string)`

SetName sets Name field to given value.


### GetModelNumber

`func (o *AgentsGetItem) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *AgentsGetItem) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *AgentsGetItem) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.


### SetModelNumberNil

`func (o *AgentsGetItem) SetModelNumberNil(b bool)`

 SetModelNumberNil sets the value for ModelNumber to be an explicit nil

### UnsetModelNumber
`func (o *AgentsGetItem) UnsetModelNumber()`

UnsetModelNumber ensures that no value is present for ModelNumber, not even an explicit nil
### GetWifiMacAddress

`func (o *AgentsGetItem) GetWifiMacAddress() string`

GetWifiMacAddress returns the WifiMacAddress field if non-nil, zero value otherwise.

### GetWifiMacAddressOk

`func (o *AgentsGetItem) GetWifiMacAddressOk() (*string, bool)`

GetWifiMacAddressOk returns a tuple with the WifiMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacAddress

`func (o *AgentsGetItem) SetWifiMacAddress(v string)`

SetWifiMacAddress sets WifiMacAddress field to given value.


### SetWifiMacAddressNil

`func (o *AgentsGetItem) SetWifiMacAddressNil(b bool)`

 SetWifiMacAddressNil sets the value for WifiMacAddress to be an explicit nil

### UnsetWifiMacAddress
`func (o *AgentsGetItem) UnsetWifiMacAddress()`

UnsetWifiMacAddress ensures that no value is present for WifiMacAddress, not even an explicit nil
### GetEthernetMacAddress

`func (o *AgentsGetItem) GetEthernetMacAddress() string`

GetEthernetMacAddress returns the EthernetMacAddress field if non-nil, zero value otherwise.

### GetEthernetMacAddressOk

`func (o *AgentsGetItem) GetEthernetMacAddressOk() (*string, bool)`

GetEthernetMacAddressOk returns a tuple with the EthernetMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetMacAddress

`func (o *AgentsGetItem) SetEthernetMacAddress(v string)`

SetEthernetMacAddress sets EthernetMacAddress field to given value.


### SetEthernetMacAddressNil

`func (o *AgentsGetItem) SetEthernetMacAddressNil(b bool)`

 SetEthernetMacAddressNil sets the value for EthernetMacAddress to be an explicit nil

### UnsetEthernetMacAddress
`func (o *AgentsGetItem) UnsetEthernetMacAddress()`

UnsetEthernetMacAddress ensures that no value is present for EthernetMacAddress, not even an explicit nil
### GetNotes

`func (o *AgentsGetItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AgentsGetItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AgentsGetItem) SetNotes(v string)`

SetNotes sets Notes field to given value.


### SetNotesNil

`func (o *AgentsGetItem) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *AgentsGetItem) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetPcapMode

`func (o *AgentsGetItem) GetPcapMode() AgentPcapMode`

GetPcapMode returns the PcapMode field if non-nil, zero value otherwise.

### GetPcapModeOk

`func (o *AgentsGetItem) GetPcapModeOk() (*AgentPcapMode, bool)`

GetPcapModeOk returns a tuple with the PcapMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcapMode

`func (o *AgentsGetItem) SetPcapMode(v AgentPcapMode)`

SetPcapMode sets PcapMode field to given value.


### SetPcapModeNil

`func (o *AgentsGetItem) SetPcapModeNil(b bool)`

 SetPcapModeNil sets the value for PcapMode to be an explicit nil

### UnsetPcapMode
`func (o *AgentsGetItem) UnsetPcapMode()`

UnsetPcapMode ensures that no value is present for PcapMode, not even an explicit nil
### GetType

`func (o *AgentsGetItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentsGetItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentsGetItem) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


