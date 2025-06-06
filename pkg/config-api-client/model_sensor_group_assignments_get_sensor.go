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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SensorGroupAssignmentsGetSensor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SensorGroupAssignmentsGetSensor{}

// SensorGroupAssignmentsGetSensor struct for SensorGroupAssignmentsGetSensor
type SensorGroupAssignmentsGetSensor struct {
	// The unique identifier of the sensor
	Id string `json:"id"`
}

type _SensorGroupAssignmentsGetSensor SensorGroupAssignmentsGetSensor

// NewSensorGroupAssignmentsGetSensor instantiates a new SensorGroupAssignmentsGetSensor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensorGroupAssignmentsGetSensor(id string) *SensorGroupAssignmentsGetSensor {
	this := SensorGroupAssignmentsGetSensor{}
	this.Id = id
	return &this
}

// NewSensorGroupAssignmentsGetSensorWithDefaults instantiates a new SensorGroupAssignmentsGetSensor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensorGroupAssignmentsGetSensorWithDefaults() *SensorGroupAssignmentsGetSensor {
	this := SensorGroupAssignmentsGetSensor{}
	return &this
}

// GetId returns the Id field value
func (o *SensorGroupAssignmentsGetSensor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SensorGroupAssignmentsGetSensor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SensorGroupAssignmentsGetSensor) SetId(v string) {
	o.Id = v
}

func (o SensorGroupAssignmentsGetSensor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SensorGroupAssignmentsGetSensor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *SensorGroupAssignmentsGetSensor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSensorGroupAssignmentsGetSensor := _SensorGroupAssignmentsGetSensor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSensorGroupAssignmentsGetSensor)

	if err != nil {
		return err
	}

	*o = SensorGroupAssignmentsGetSensor(varSensorGroupAssignmentsGetSensor)

	return err
}

type NullableSensorGroupAssignmentsGetSensor struct {
	value *SensorGroupAssignmentsGetSensor
	isSet bool
}

func (v NullableSensorGroupAssignmentsGetSensor) Get() *SensorGroupAssignmentsGetSensor {
	return v.value
}

func (v *NullableSensorGroupAssignmentsGetSensor) Set(val *SensorGroupAssignmentsGetSensor) {
	v.value = val
	v.isSet = true
}

func (v NullableSensorGroupAssignmentsGetSensor) IsSet() bool {
	return v.isSet
}

func (v *NullableSensorGroupAssignmentsGetSensor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensorGroupAssignmentsGetSensor(
	val *SensorGroupAssignmentsGetSensor,
) *NullableSensorGroupAssignmentsGetSensor {
	return &NullableSensorGroupAssignmentsGetSensor{value: val, isSet: true}
}

func (v NullableSensorGroupAssignmentsGetSensor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensorGroupAssignmentsGetSensor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
