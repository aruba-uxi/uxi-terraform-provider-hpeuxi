# SensorItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Serial** | **string** |  | 
**Name** | **string** |  | 
**ModelNumber** | **NullableString** |  | 
**WifiMacAddress** | **NullableString** |  | 
**EthernetMacAddress** | **NullableString** |  | 
**AddressNote** | **NullableString** |  | 
**Longitude** | **NullableFloat32** |  | 
**Latitude** | **NullableFloat32** |  | 
**Notes** | **NullableString** |  | 
**PcapMode** | **NullableString** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "uxi/sensor"]

## Methods

### NewSensorItem

`func NewSensorItem(id string, serial string, name string, modelNumber NullableString, wifiMacAddress NullableString, ethernetMacAddress NullableString, addressNote NullableString, longitude NullableFloat32, latitude NullableFloat32, notes NullableString, pcapMode NullableString, ) *SensorItem`

NewSensorItem instantiates a new SensorItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorItemWithDefaults

`func NewSensorItemWithDefaults() *SensorItem`

NewSensorItemWithDefaults instantiates a new SensorItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SensorItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SensorItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SensorItem) SetId(v string)`

SetId sets Id field to given value.


### GetSerial

`func (o *SensorItem) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *SensorItem) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *SensorItem) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetName

`func (o *SensorItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SensorItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SensorItem) SetName(v string)`

SetName sets Name field to given value.


### GetModelNumber

`func (o *SensorItem) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *SensorItem) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *SensorItem) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.


### SetModelNumberNil

`func (o *SensorItem) SetModelNumberNil(b bool)`

 SetModelNumberNil sets the value for ModelNumber to be an explicit nil

### UnsetModelNumber
`func (o *SensorItem) UnsetModelNumber()`

UnsetModelNumber ensures that no value is present for ModelNumber, not even an explicit nil
### GetWifiMacAddress

`func (o *SensorItem) GetWifiMacAddress() string`

GetWifiMacAddress returns the WifiMacAddress field if non-nil, zero value otherwise.

### GetWifiMacAddressOk

`func (o *SensorItem) GetWifiMacAddressOk() (*string, bool)`

GetWifiMacAddressOk returns a tuple with the WifiMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacAddress

`func (o *SensorItem) SetWifiMacAddress(v string)`

SetWifiMacAddress sets WifiMacAddress field to given value.


### SetWifiMacAddressNil

`func (o *SensorItem) SetWifiMacAddressNil(b bool)`

 SetWifiMacAddressNil sets the value for WifiMacAddress to be an explicit nil

### UnsetWifiMacAddress
`func (o *SensorItem) UnsetWifiMacAddress()`

UnsetWifiMacAddress ensures that no value is present for WifiMacAddress, not even an explicit nil
### GetEthernetMacAddress

`func (o *SensorItem) GetEthernetMacAddress() string`

GetEthernetMacAddress returns the EthernetMacAddress field if non-nil, zero value otherwise.

### GetEthernetMacAddressOk

`func (o *SensorItem) GetEthernetMacAddressOk() (*string, bool)`

GetEthernetMacAddressOk returns a tuple with the EthernetMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetMacAddress

`func (o *SensorItem) SetEthernetMacAddress(v string)`

SetEthernetMacAddress sets EthernetMacAddress field to given value.


### SetEthernetMacAddressNil

`func (o *SensorItem) SetEthernetMacAddressNil(b bool)`

 SetEthernetMacAddressNil sets the value for EthernetMacAddress to be an explicit nil

### UnsetEthernetMacAddress
`func (o *SensorItem) UnsetEthernetMacAddress()`

UnsetEthernetMacAddress ensures that no value is present for EthernetMacAddress, not even an explicit nil
### GetAddressNote

`func (o *SensorItem) GetAddressNote() string`

GetAddressNote returns the AddressNote field if non-nil, zero value otherwise.

### GetAddressNoteOk

`func (o *SensorItem) GetAddressNoteOk() (*string, bool)`

GetAddressNoteOk returns a tuple with the AddressNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressNote

`func (o *SensorItem) SetAddressNote(v string)`

SetAddressNote sets AddressNote field to given value.


### SetAddressNoteNil

`func (o *SensorItem) SetAddressNoteNil(b bool)`

 SetAddressNoteNil sets the value for AddressNote to be an explicit nil

### UnsetAddressNote
`func (o *SensorItem) UnsetAddressNote()`

UnsetAddressNote ensures that no value is present for AddressNote, not even an explicit nil
### GetLongitude

`func (o *SensorItem) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SensorItem) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SensorItem) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### SetLongitudeNil

`func (o *SensorItem) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *SensorItem) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetLatitude

`func (o *SensorItem) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SensorItem) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SensorItem) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### SetLatitudeNil

`func (o *SensorItem) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *SensorItem) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetNotes

`func (o *SensorItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SensorItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SensorItem) SetNotes(v string)`

SetNotes sets Notes field to given value.


### SetNotesNil

`func (o *SensorItem) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *SensorItem) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetPcapMode

`func (o *SensorItem) GetPcapMode() string`

GetPcapMode returns the PcapMode field if non-nil, zero value otherwise.

### GetPcapModeOk

`func (o *SensorItem) GetPcapModeOk() (*string, bool)`

GetPcapModeOk returns a tuple with the PcapMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcapMode

`func (o *SensorItem) SetPcapMode(v string)`

SetPcapMode sets PcapMode field to given value.


### SetPcapModeNil

`func (o *SensorItem) SetPcapModeNil(b bool)`

 SetPcapModeNil sets the value for PcapMode to be an explicit nil

### UnsetPcapMode
`func (o *SensorItem) UnsetPcapMode()`

UnsetPcapMode ensures that no value is present for PcapMode, not even an explicit nil
### GetType

`func (o *SensorItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SensorItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SensorItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SensorItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


