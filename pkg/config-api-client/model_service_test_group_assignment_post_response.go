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

// checks if the ServiceTestGroupAssignmentPostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTestGroupAssignmentPostResponse{}

// ServiceTestGroupAssignmentPostResponse struct for ServiceTestGroupAssignmentPostResponse
type ServiceTestGroupAssignmentPostResponse struct {
	// The unique identifier of the service test group assignment
	Id string `json:"id"`
	// The group component of the service test group assignment
	Group ServiceTestGroupAssignmentPostGroup `json:"group"`
	// The service test component of the service test group assignment
	ServiceTest ServiceTestGroupAssignmentPostServiceTest `json:"serviceTest"`
	// The type of the resource.
	Type string `json:"type"`
}

type _ServiceTestGroupAssignmentPostResponse ServiceTestGroupAssignmentPostResponse

// NewServiceTestGroupAssignmentPostResponse instantiates a new ServiceTestGroupAssignmentPostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTestGroupAssignmentPostResponse(
	id string,
	group ServiceTestGroupAssignmentPostGroup,
	serviceTest ServiceTestGroupAssignmentPostServiceTest,
	type_ string,
) *ServiceTestGroupAssignmentPostResponse {
	this := ServiceTestGroupAssignmentPostResponse{}
	this.Id = id
	this.Group = group
	this.ServiceTest = serviceTest
	this.Type = type_
	return &this
}

// NewServiceTestGroupAssignmentPostResponseWithDefaults instantiates a new ServiceTestGroupAssignmentPostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTestGroupAssignmentPostResponseWithDefaults() *ServiceTestGroupAssignmentPostResponse {
	this := ServiceTestGroupAssignmentPostResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ServiceTestGroupAssignmentPostResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceTestGroupAssignmentPostResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceTestGroupAssignmentPostResponse) SetId(v string) {
	o.Id = v
}

// GetGroup returns the Group field value
func (o *ServiceTestGroupAssignmentPostResponse) GetGroup() ServiceTestGroupAssignmentPostGroup {
	if o == nil {
		var ret ServiceTestGroupAssignmentPostGroup
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *ServiceTestGroupAssignmentPostResponse) GetGroupOk() (*ServiceTestGroupAssignmentPostGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *ServiceTestGroupAssignmentPostResponse) SetGroup(v ServiceTestGroupAssignmentPostGroup) {
	o.Group = v
}

// GetServiceTest returns the ServiceTest field value
func (o *ServiceTestGroupAssignmentPostResponse) GetServiceTest() ServiceTestGroupAssignmentPostServiceTest {
	if o == nil {
		var ret ServiceTestGroupAssignmentPostServiceTest
		return ret
	}

	return o.ServiceTest
}

// GetServiceTestOk returns a tuple with the ServiceTest field value
// and a boolean to check if the value has been set.
func (o *ServiceTestGroupAssignmentPostResponse) GetServiceTestOk() (*ServiceTestGroupAssignmentPostServiceTest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceTest, true
}

// SetServiceTest sets field value
func (o *ServiceTestGroupAssignmentPostResponse) SetServiceTest(
	v ServiceTestGroupAssignmentPostServiceTest,
) {
	o.ServiceTest = v
}

// GetType returns the Type field value
func (o *ServiceTestGroupAssignmentPostResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceTestGroupAssignmentPostResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceTestGroupAssignmentPostResponse) SetType(v string) {
	o.Type = v
}

func (o ServiceTestGroupAssignmentPostResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTestGroupAssignmentPostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["group"] = o.Group
	toSerialize["serviceTest"] = o.ServiceTest
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ServiceTestGroupAssignmentPostResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"group",
		"serviceTest",
		"type",
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

	varServiceTestGroupAssignmentPostResponse := _ServiceTestGroupAssignmentPostResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceTestGroupAssignmentPostResponse)

	if err != nil {
		return err
	}

	*o = ServiceTestGroupAssignmentPostResponse(varServiceTestGroupAssignmentPostResponse)

	return err
}

type NullableServiceTestGroupAssignmentPostResponse struct {
	value *ServiceTestGroupAssignmentPostResponse
	isSet bool
}

func (v NullableServiceTestGroupAssignmentPostResponse) Get() *ServiceTestGroupAssignmentPostResponse {
	return v.value
}

func (v *NullableServiceTestGroupAssignmentPostResponse) Set(
	val *ServiceTestGroupAssignmentPostResponse,
) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTestGroupAssignmentPostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTestGroupAssignmentPostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTestGroupAssignmentPostResponse(
	val *ServiceTestGroupAssignmentPostResponse,
) *NullableServiceTestGroupAssignmentPostResponse {
	return &NullableServiceTestGroupAssignmentPostResponse{value: val, isSet: true}
}

func (v NullableServiceTestGroupAssignmentPostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTestGroupAssignmentPostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
