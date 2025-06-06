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

// checks if the ServiceTestGroupAssignmentPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTestGroupAssignmentPostRequest{}

// ServiceTestGroupAssignmentPostRequest struct for ServiceTestGroupAssignmentPostRequest
type ServiceTestGroupAssignmentPostRequest struct {
	// The unique identifier of the group
	GroupId string `json:"groupId"`
	// The unique identifier of the service test
	ServiceTestId string `json:"serviceTestId"`
}

type _ServiceTestGroupAssignmentPostRequest ServiceTestGroupAssignmentPostRequest

// NewServiceTestGroupAssignmentPostRequest instantiates a new ServiceTestGroupAssignmentPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTestGroupAssignmentPostRequest(
	groupId string,
	serviceTestId string,
) *ServiceTestGroupAssignmentPostRequest {
	this := ServiceTestGroupAssignmentPostRequest{}
	this.GroupId = groupId
	this.ServiceTestId = serviceTestId
	return &this
}

// NewServiceTestGroupAssignmentPostRequestWithDefaults instantiates a new ServiceTestGroupAssignmentPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTestGroupAssignmentPostRequestWithDefaults() *ServiceTestGroupAssignmentPostRequest {
	this := ServiceTestGroupAssignmentPostRequest{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *ServiceTestGroupAssignmentPostRequest) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ServiceTestGroupAssignmentPostRequest) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ServiceTestGroupAssignmentPostRequest) SetGroupId(v string) {
	o.GroupId = v
}

// GetServiceTestId returns the ServiceTestId field value
func (o *ServiceTestGroupAssignmentPostRequest) GetServiceTestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceTestId
}

// GetServiceTestIdOk returns a tuple with the ServiceTestId field value
// and a boolean to check if the value has been set.
func (o *ServiceTestGroupAssignmentPostRequest) GetServiceTestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceTestId, true
}

// SetServiceTestId sets field value
func (o *ServiceTestGroupAssignmentPostRequest) SetServiceTestId(v string) {
	o.ServiceTestId = v
}

func (o ServiceTestGroupAssignmentPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTestGroupAssignmentPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupId"] = o.GroupId
	toSerialize["serviceTestId"] = o.ServiceTestId
	return toSerialize, nil
}

func (o *ServiceTestGroupAssignmentPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groupId",
		"serviceTestId",
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

	varServiceTestGroupAssignmentPostRequest := _ServiceTestGroupAssignmentPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceTestGroupAssignmentPostRequest)

	if err != nil {
		return err
	}

	*o = ServiceTestGroupAssignmentPostRequest(varServiceTestGroupAssignmentPostRequest)

	return err
}

type NullableServiceTestGroupAssignmentPostRequest struct {
	value *ServiceTestGroupAssignmentPostRequest
	isSet bool
}

func (v NullableServiceTestGroupAssignmentPostRequest) Get() *ServiceTestGroupAssignmentPostRequest {
	return v.value
}

func (v *NullableServiceTestGroupAssignmentPostRequest) Set(
	val *ServiceTestGroupAssignmentPostRequest,
) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTestGroupAssignmentPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTestGroupAssignmentPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTestGroupAssignmentPostRequest(
	val *ServiceTestGroupAssignmentPostRequest,
) *NullableServiceTestGroupAssignmentPostRequest {
	return &NullableServiceTestGroupAssignmentPostRequest{value: val, isSet: true}
}

func (v NullableServiceTestGroupAssignmentPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTestGroupAssignmentPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
