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

// checks if the GroupPatchParent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupPatchParent{}

// GroupPatchParent struct for GroupPatchParent
type GroupPatchParent struct {
	// The unique identifier of the parent group
	Id string `json:"id"`
}

type _GroupPatchParent GroupPatchParent

// NewGroupPatchParent instantiates a new GroupPatchParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupPatchParent(id string) *GroupPatchParent {
	this := GroupPatchParent{}
	this.Id = id
	return &this
}

// NewGroupPatchParentWithDefaults instantiates a new GroupPatchParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupPatchParentWithDefaults() *GroupPatchParent {
	this := GroupPatchParent{}
	return &this
}

// GetId returns the Id field value
func (o *GroupPatchParent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupPatchParent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupPatchParent) SetId(v string) {
	o.Id = v
}

func (o GroupPatchParent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupPatchParent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *GroupPatchParent) UnmarshalJSON(data []byte) (err error) {
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

	varGroupPatchParent := _GroupPatchParent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupPatchParent)

	if err != nil {
		return err
	}

	*o = GroupPatchParent(varGroupPatchParent)

	return err
}

type NullableGroupPatchParent struct {
	value *GroupPatchParent
	isSet bool
}

func (v NullableGroupPatchParent) Get() *GroupPatchParent {
	return v.value
}

func (v *NullableGroupPatchParent) Set(val *GroupPatchParent) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupPatchParent) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupPatchParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupPatchParent(val *GroupPatchParent) *NullableGroupPatchParent {
	return &NullableGroupPatchParent{value: val, isSet: true}
}

func (v NullableGroupPatchParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupPatchParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
