# SensorPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The updated sensor name | [optional] 
**AddressNote** | Pointer to **string** | The updated address note for the sensor | [optional] 
**Notes** | Pointer to **string** | Additional notes for the sensor | [optional] 
**PcapMode** | Pointer to [**SensorPcapMode**](SensorPcapMode.md) | The updated pcap mode for the sensor | [optional] 

## Methods

### NewSensorPatchRequest

`func NewSensorPatchRequest() *SensorPatchRequest`

NewSensorPatchRequest instantiates a new SensorPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorPatchRequestWithDefaults

`func NewSensorPatchRequestWithDefaults() *SensorPatchRequest`

NewSensorPatchRequestWithDefaults instantiates a new SensorPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SensorPatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SensorPatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SensorPatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SensorPatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddressNote

`func (o *SensorPatchRequest) GetAddressNote() string`

GetAddressNote returns the AddressNote field if non-nil, zero value otherwise.

### GetAddressNoteOk

`func (o *SensorPatchRequest) GetAddressNoteOk() (*string, bool)`

GetAddressNoteOk returns a tuple with the AddressNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressNote

`func (o *SensorPatchRequest) SetAddressNote(v string)`

SetAddressNote sets AddressNote field to given value.

### HasAddressNote

`func (o *SensorPatchRequest) HasAddressNote() bool`

HasAddressNote returns a boolean if a field has been set.

### GetNotes

`func (o *SensorPatchRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SensorPatchRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SensorPatchRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SensorPatchRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPcapMode

`func (o *SensorPatchRequest) GetPcapMode() SensorPcapMode`

GetPcapMode returns the PcapMode field if non-nil, zero value otherwise.

### GetPcapModeOk

`func (o *SensorPatchRequest) GetPcapModeOk() (*SensorPcapMode, bool)`

GetPcapModeOk returns a tuple with the PcapMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcapMode

`func (o *SensorPatchRequest) SetPcapMode(v SensorPcapMode)`

SetPcapMode sets PcapMode field to given value.

### HasPcapMode

`func (o *SensorPatchRequest) HasPcapMode() bool`

HasPcapMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


