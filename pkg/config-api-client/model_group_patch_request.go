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

// checks if the GroupPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupPatchRequest{}

// GroupPatchRequest struct for GroupPatchRequest
type GroupPatchRequest struct {
	// The updated group name
	Name *string `json:"name,omitempty"`
}

// NewGroupPatchRequest instantiates a new GroupPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupPatchRequest() *GroupPatchRequest {
	this := GroupPatchRequest{}
	return &this
}

// NewGroupPatchRequestWithDefaults instantiates a new GroupPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupPatchRequestWithDefaults() *GroupPatchRequest {
	this := GroupPatchRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupPatchRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPatchRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupPatchRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupPatchRequest) SetName(v string) {
	o.Name = &v
}

func (o GroupPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGroupPatchRequest struct {
	value *GroupPatchRequest
	isSet bool
}

func (v NullableGroupPatchRequest) Get() *GroupPatchRequest {
	return v.value
}

func (v *NullableGroupPatchRequest) Set(val *GroupPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupPatchRequest(val *GroupPatchRequest) *NullableGroupPatchRequest {
	return &NullableGroupPatchRequest{value: val, isSet: true}
}

func (v NullableGroupPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
