# SensorPatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the sensor | 
**Serial** | **string** | The serial number of the sensor | 
**Name** | **string** | The name of the sensor | 
**GroupName** | **NullableString** |  | 
**GroupPath** | **NullableString** |  | 
**ModelNumber** | **string** | The model number of the sensor | 
**WifiMacAddress** | **NullableString** |  | 
**EthernetMacAddress** | **NullableString** |  | 
**AddressNote** | **NullableString** |  | 
**Longitude** | **NullableFloat32** |  | 
**Latitude** | **NullableFloat32** |  | 
**Notes** | **NullableString** |  | 
**PcapMode** | [**NullableSensorPcapMode**](SensorPcapMode.md) |  | 
**Type** | **string** | The type of the resource. | 

## Methods

### NewSensorPatchResponse

`func NewSensorPatchResponse(id string, serial string, name string, groupName NullableString, groupPath NullableString, modelNumber string, wifiMacAddress NullableString, ethernetMacAddress NullableString, addressNote NullableString, longitude NullableFloat32, latitude NullableFloat32, notes NullableString, pcapMode NullableSensorPcapMode, type_ string, ) *SensorPatchResponse`

NewSensorPatchResponse instantiates a new SensorPatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorPatchResponseWithDefaults

`func NewSensorPatchResponseWithDefaults() *SensorPatchResponse`

NewSensorPatchResponseWithDefaults instantiates a new SensorPatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SensorPatchResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SensorPatchResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SensorPatchResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSerial

`func (o *SensorPatchResponse) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *SensorPatchResponse) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *SensorPatchResponse) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetName

`func (o *SensorPatchResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SensorPatchResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SensorPatchResponse) SetName(v string)`

SetName sets Name field to given value.


### GetGroupName

`func (o *SensorPatchResponse) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *SensorPatchResponse) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *SensorPatchResponse) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### SetGroupNameNil

`func (o *SensorPatchResponse) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *SensorPatchResponse) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetGroupPath

`func (o *SensorPatchResponse) GetGroupPath() string`

GetGroupPath returns the GroupPath field if non-nil, zero value otherwise.

### GetGroupPathOk

`func (o *SensorPatchResponse) GetGroupPathOk() (*string, bool)`

GetGroupPathOk returns a tuple with the GroupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPath

`func (o *SensorPatchResponse) SetGroupPath(v string)`

SetGroupPath sets GroupPath field to given value.


### SetGroupPathNil

`func (o *SensorPatchResponse) SetGroupPathNil(b bool)`

 SetGroupPathNil sets the value for GroupPath to be an explicit nil

### UnsetGroupPath
`func (o *SensorPatchResponse) UnsetGroupPath()`

UnsetGroupPath ensures that no value is present for GroupPath, not even an explicit nil
### GetModelNumber

`func (o *SensorPatchResponse) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *SensorPatchResponse) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *SensorPatchResponse) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.


### GetWifiMacAddress

`func (o *SensorPatchResponse) GetWifiMacAddress() string`

GetWifiMacAddress returns the WifiMacAddress field if non-nil, zero value otherwise.

### GetWifiMacAddressOk

`func (o *SensorPatchResponse) GetWifiMacAddressOk() (*string, bool)`

GetWifiMacAddressOk returns a tuple with the WifiMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacAddress

`func (o *SensorPatchResponse) SetWifiMacAddress(v string)`

SetWifiMacAddress sets WifiMacAddress field to given value.


### SetWifiMacAddressNil

`func (o *SensorPatchResponse) SetWifiMacAddressNil(b bool)`

 SetWifiMacAddressNil sets the value for WifiMacAddress to be an explicit nil

### UnsetWifiMacAddress
`func (o *SensorPatchResponse) UnsetWifiMacAddress()`

UnsetWifiMacAddress ensures that no value is present for WifiMacAddress, not even an explicit nil
### GetEthernetMacAddress

`func (o *SensorPatchResponse) GetEthernetMacAddress() string`

GetEthernetMacAddress returns the EthernetMacAddress field if non-nil, zero value otherwise.

### GetEthernetMacAddressOk

`func (o *SensorPatchResponse) GetEthernetMacAddressOk() (*string, bool)`

GetEthernetMacAddressOk returns a tuple with the EthernetMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetMacAddress

`func (o *SensorPatchResponse) SetEthernetMacAddress(v string)`

SetEthernetMacAddress sets EthernetMacAddress field to given value.


### SetEthernetMacAddressNil

`func (o *SensorPatchResponse) SetEthernetMacAddressNil(b bool)`

 SetEthernetMacAddressNil sets the value for EthernetMacAddress to be an explicit nil

### UnsetEthernetMacAddress
`func (o *SensorPatchResponse) UnsetEthernetMacAddress()`

UnsetEthernetMacAddress ensures that no value is present for EthernetMacAddress, not even an explicit nil
### GetAddressNote

`func (o *SensorPatchResponse) GetAddressNote() string`

GetAddressNote returns the AddressNote field if non-nil, zero value otherwise.

### GetAddressNoteOk

`func (o *SensorPatchResponse) GetAddressNoteOk() (*string, bool)`

GetAddressNoteOk returns a tuple with the AddressNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressNote

`func (o *SensorPatchResponse) SetAddressNote(v string)`

SetAddressNote sets AddressNote field to given value.


### SetAddressNoteNil

`func (o *SensorPatchResponse) SetAddressNoteNil(b bool)`

 SetAddressNoteNil sets the value for AddressNote to be an explicit nil

### UnsetAddressNote
`func (o *SensorPatchResponse) UnsetAddressNote()`

UnsetAddressNote ensures that no value is present for AddressNote, not even an explicit nil
### GetLongitude

`func (o *SensorPatchResponse) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SensorPatchResponse) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SensorPatchResponse) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### SetLongitudeNil

`func (o *SensorPatchResponse) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *SensorPatchResponse) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetLatitude

`func (o *SensorPatchResponse) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SensorPatchResponse) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SensorPatchResponse) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### SetLatitudeNil

`func (o *SensorPatchResponse) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *SensorPatchResponse) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetNotes

`func (o *SensorPatchResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SensorPatchResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SensorPatchResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.


### SetNotesNil

`func (o *SensorPatchResponse) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *SensorPatchResponse) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetPcapMode

`func (o *SensorPatchResponse) GetPcapMode() SensorPcapMode`

GetPcapMode returns the PcapMode field if non-nil, zero value otherwise.

### GetPcapModeOk

`func (o *SensorPatchResponse) GetPcapModeOk() (*SensorPcapMode, bool)`

GetPcapModeOk returns a tuple with the PcapMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcapMode

`func (o *SensorPatchResponse) SetPcapMode(v SensorPcapMode)`

SetPcapMode sets PcapMode field to given value.


### SetPcapModeNil

`func (o *SensorPatchResponse) SetPcapModeNil(b bool)`

 SetPcapModeNil sets the value for PcapMode to be an explicit nil

### UnsetPcapMode
`func (o *SensorPatchResponse) UnsetPcapMode()`

UnsetPcapMode ensures that no value is present for PcapMode, not even an explicit nil
### GetType

`func (o *SensorPatchResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SensorPatchResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SensorPatchResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


