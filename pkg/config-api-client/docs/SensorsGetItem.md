# SensorsGetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the sensor | 
**Serial** | **string** | The serial number of the sensor | 
**Name** | **string** | The name of the sensor | 
**ModelNumber** | **string** | The model number of the sensor | 
**WifiMacAddress** | **NullableString** | The WiFi MAC address of the sensor | 
**EthernetMacAddress** | **NullableString** | The Ethernet MAC address of the sensor | 
**AddressNote** | **NullableString** | The address note of the sensor | 
**Longitude** | **NullableFloat32** | The longitude of the sensor | 
**Latitude** | **NullableFloat32** | The latitude of the sensor | 
**Notes** | **NullableString** | The notes of the sensor | 
**PcapMode** | [**NullableSensorPcapMode**](SensorPcapMode.md) |  | 
**Type** | **string** | The type of the resource. | 

## Methods

### NewSensorsGetItem

`func NewSensorsGetItem(id string, serial string, name string, modelNumber string, wifiMacAddress NullableString, ethernetMacAddress NullableString, addressNote NullableString, longitude NullableFloat32, latitude NullableFloat32, notes NullableString, pcapMode NullableSensorPcapMode, type_ string, ) *SensorsGetItem`

NewSensorsGetItem instantiates a new SensorsGetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorsGetItemWithDefaults

`func NewSensorsGetItemWithDefaults() *SensorsGetItem`

NewSensorsGetItemWithDefaults instantiates a new SensorsGetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SensorsGetItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SensorsGetItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SensorsGetItem) SetId(v string)`

SetId sets Id field to given value.


### GetSerial

`func (o *SensorsGetItem) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *SensorsGetItem) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *SensorsGetItem) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetName

`func (o *SensorsGetItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SensorsGetItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SensorsGetItem) SetName(v string)`

SetName sets Name field to given value.


### GetModelNumber

`func (o *SensorsGetItem) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *SensorsGetItem) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *SensorsGetItem) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.


### GetWifiMacAddress

`func (o *SensorsGetItem) GetWifiMacAddress() string`

GetWifiMacAddress returns the WifiMacAddress field if non-nil, zero value otherwise.

### GetWifiMacAddressOk

`func (o *SensorsGetItem) GetWifiMacAddressOk() (*string, bool)`

GetWifiMacAddressOk returns a tuple with the WifiMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacAddress

`func (o *SensorsGetItem) SetWifiMacAddress(v string)`

SetWifiMacAddress sets WifiMacAddress field to given value.


### SetWifiMacAddressNil

`func (o *SensorsGetItem) SetWifiMacAddressNil(b bool)`

 SetWifiMacAddressNil sets the value for WifiMacAddress to be an explicit nil

### UnsetWifiMacAddress
`func (o *SensorsGetItem) UnsetWifiMacAddress()`

UnsetWifiMacAddress ensures that no value is present for WifiMacAddress, not even an explicit nil
### GetEthernetMacAddress

`func (o *SensorsGetItem) GetEthernetMacAddress() string`

GetEthernetMacAddress returns the EthernetMacAddress field if non-nil, zero value otherwise.

### GetEthernetMacAddressOk

`func (o *SensorsGetItem) GetEthernetMacAddressOk() (*string, bool)`

GetEthernetMacAddressOk returns a tuple with the EthernetMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthernetMacAddress

`func (o *SensorsGetItem) SetEthernetMacAddress(v string)`

SetEthernetMacAddress sets EthernetMacAddress field to given value.


### SetEthernetMacAddressNil

`func (o *SensorsGetItem) SetEthernetMacAddressNil(b bool)`

 SetEthernetMacAddressNil sets the value for EthernetMacAddress to be an explicit nil

### UnsetEthernetMacAddress
`func (o *SensorsGetItem) UnsetEthernetMacAddress()`

UnsetEthernetMacAddress ensures that no value is present for EthernetMacAddress, not even an explicit nil
### GetAddressNote

`func (o *SensorsGetItem) GetAddressNote() string`

GetAddressNote returns the AddressNote field if non-nil, zero value otherwise.

### GetAddressNoteOk

`func (o *SensorsGetItem) GetAddressNoteOk() (*string, bool)`

GetAddressNoteOk returns a tuple with the AddressNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressNote

`func (o *SensorsGetItem) SetAddressNote(v string)`

SetAddressNote sets AddressNote field to given value.


### SetAddressNoteNil

`func (o *SensorsGetItem) SetAddressNoteNil(b bool)`

 SetAddressNoteNil sets the value for AddressNote to be an explicit nil

### UnsetAddressNote
`func (o *SensorsGetItem) UnsetAddressNote()`

UnsetAddressNote ensures that no value is present for AddressNote, not even an explicit nil
### GetLongitude

`func (o *SensorsGetItem) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SensorsGetItem) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SensorsGetItem) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### SetLongitudeNil

`func (o *SensorsGetItem) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *SensorsGetItem) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetLatitude

`func (o *SensorsGetItem) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SensorsGetItem) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SensorsGetItem) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### SetLatitudeNil

`func (o *SensorsGetItem) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *SensorsGetItem) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetNotes

`func (o *SensorsGetItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SensorsGetItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SensorsGetItem) SetNotes(v string)`

SetNotes sets Notes field to given value.


### SetNotesNil

`func (o *SensorsGetItem) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *SensorsGetItem) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetPcapMode

`func (o *SensorsGetItem) GetPcapMode() SensorPcapMode`

GetPcapMode returns the PcapMode field if non-nil, zero value otherwise.

### GetPcapModeOk

`func (o *SensorsGetItem) GetPcapModeOk() (*SensorPcapMode, bool)`

GetPcapModeOk returns a tuple with the PcapMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcapMode

`func (o *SensorsGetItem) SetPcapMode(v SensorPcapMode)`

SetPcapMode sets PcapMode field to given value.


### SetPcapModeNil

`func (o *SensorsGetItem) SetPcapModeNil(b bool)`

 SetPcapModeNil sets the value for PcapMode to be an explicit nil

### UnsetPcapMode
`func (o *SensorsGetItem) UnsetPcapMode()`

UnsetPcapMode ensures that no value is present for PcapMode, not even an explicit nil
### GetType

`func (o *SensorsGetItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SensorsGetItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SensorsGetItem) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


