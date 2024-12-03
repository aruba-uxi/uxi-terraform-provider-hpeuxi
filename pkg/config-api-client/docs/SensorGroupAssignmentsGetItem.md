# SensorGroupAssignmentsGetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Group** | [**SensorGroupAssignmentsGetGroup**](SensorGroupAssignmentsGetGroup.md) |  | 
**Sensor** | [**SensorGroupAssignmentsGetSensor**](SensorGroupAssignmentsGetSensor.md) |  | 
**Type** | **string** |  | 

## Methods

### NewSensorGroupAssignmentsGetItem

`func NewSensorGroupAssignmentsGetItem(id string, group SensorGroupAssignmentsGetGroup, sensor SensorGroupAssignmentsGetSensor, type_ string, ) *SensorGroupAssignmentsGetItem`

NewSensorGroupAssignmentsGetItem instantiates a new SensorGroupAssignmentsGetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorGroupAssignmentsGetItemWithDefaults

`func NewSensorGroupAssignmentsGetItemWithDefaults() *SensorGroupAssignmentsGetItem`

NewSensorGroupAssignmentsGetItemWithDefaults instantiates a new SensorGroupAssignmentsGetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SensorGroupAssignmentsGetItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SensorGroupAssignmentsGetItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SensorGroupAssignmentsGetItem) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *SensorGroupAssignmentsGetItem) GetGroup() SensorGroupAssignmentsGetGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SensorGroupAssignmentsGetItem) GetGroupOk() (*SensorGroupAssignmentsGetGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SensorGroupAssignmentsGetItem) SetGroup(v SensorGroupAssignmentsGetGroup)`

SetGroup sets Group field to given value.


### GetSensor

`func (o *SensorGroupAssignmentsGetItem) GetSensor() SensorGroupAssignmentsGetSensor`

GetSensor returns the Sensor field if non-nil, zero value otherwise.

### GetSensorOk

`func (o *SensorGroupAssignmentsGetItem) GetSensorOk() (*SensorGroupAssignmentsGetSensor, bool)`

GetSensorOk returns a tuple with the Sensor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensor

`func (o *SensorGroupAssignmentsGetItem) SetSensor(v SensorGroupAssignmentsGetSensor)`

SetSensor sets Sensor field to given value.


### GetType

`func (o *SensorGroupAssignmentsGetItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SensorGroupAssignmentsGetItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SensorGroupAssignmentsGetItem) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


