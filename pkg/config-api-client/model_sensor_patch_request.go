/*
Copyright 2025 Hewlett Packard Enterprise Development LP.
*/

/*
HPE Aruba Networking UXI Configuration

This document outlines the API contracts for HPE Aruba Networking UXI.

API version: 6.7.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"encoding/json"
)

// checks if the SensorPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SensorPatchRequest{}

// SensorPatchRequest Request body for patching a sensor.
type SensorPatchRequest struct {
	// The updated sensor name
	Name *string `json:"name,omitempty"`
	// The updated address note for the sensor
	AddressNote *string `json:"addressNote,omitempty"`
	// Additional notes for the sensor
	Notes *string `json:"notes,omitempty"`
	// The updated pcap mode for the sensor
	PcapMode *SensorPcapMode `json:"pcapMode,omitempty"`
}

// NewSensorPatchRequest instantiates a new SensorPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensorPatchRequest() *SensorPatchRequest {
	this := SensorPatchRequest{}
	return &this
}

// NewSensorPatchRequestWithDefaults instantiates a new SensorPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensorPatchRequestWithDefaults() *SensorPatchRequest {
	this := SensorPatchRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SensorPatchRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensorPatchRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SensorPatchRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SensorPatchRequest) SetName(v string) {
	o.Name = &v
}

// GetAddressNote returns the AddressNote field value if set, zero value otherwise.
func (o *SensorPatchRequest) GetAddressNote() string {
	if o == nil || IsNil(o.AddressNote) {
		var ret string
		return ret
	}
	return *o.AddressNote
}

// GetAddressNoteOk returns a tuple with the AddressNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensorPatchRequest) GetAddressNoteOk() (*string, bool) {
	if o == nil || IsNil(o.AddressNote) {
		return nil, false
	}
	return o.AddressNote, true
}

// HasAddressNote returns a boolean if a field has been set.
func (o *SensorPatchRequest) HasAddressNote() bool {
	if o != nil && !IsNil(o.AddressNote) {
		return true
	}

	return false
}

// SetAddressNote gets a reference to the given string and assigns it to the AddressNote field.
func (o *SensorPatchRequest) SetAddressNote(v string) {
	o.AddressNote = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *SensorPatchRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensorPatchRequest) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *SensorPatchRequest) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *SensorPatchRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetPcapMode returns the PcapMode field value if set, zero value otherwise.
func (o *SensorPatchRequest) GetPcapMode() SensorPcapMode {
	if o == nil || IsNil(o.PcapMode) {
		var ret SensorPcapMode
		return ret
	}
	return *o.PcapMode
}

// GetPcapModeOk returns a tuple with the PcapMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensorPatchRequest) GetPcapModeOk() (*SensorPcapMode, bool) {
	if o == nil || IsNil(o.PcapMode) {
		return nil, false
	}
	return o.PcapMode, true
}

// HasPcapMode returns a boolean if a field has been set.
func (o *SensorPatchRequest) HasPcapMode() bool {
	if o != nil && !IsNil(o.PcapMode) {
		return true
	}

	return false
}

// SetPcapMode gets a reference to the given SensorPcapMode and assigns it to the PcapMode field.
func (o *SensorPatchRequest) SetPcapMode(v SensorPcapMode) {
	o.PcapMode = &v
}

func (o SensorPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SensorPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AddressNote) {
		toSerialize["addressNote"] = o.AddressNote
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.PcapMode) {
		toSerialize["pcapMode"] = o.PcapMode
	}
	return toSerialize, nil
}

type NullableSensorPatchRequest struct {
	value *SensorPatchRequest
	isSet bool
}

func (v NullableSensorPatchRequest) Get() *SensorPatchRequest {
	return v.value
}

func (v *NullableSensorPatchRequest) Set(val *SensorPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSensorPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSensorPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensorPatchRequest(val *SensorPatchRequest) *NullableSensorPatchRequest {
	return &NullableSensorPatchRequest{value: val, isSet: true}
}

func (v NullableSensorPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensorPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
