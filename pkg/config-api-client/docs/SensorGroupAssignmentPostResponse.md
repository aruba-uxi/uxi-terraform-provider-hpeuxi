# SensorGroupAssignmentPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Group** | [**SensorGroupAssignmentPostGroup**](SensorGroupAssignmentPostGroup.md) |  | 
**Sensor** | [**SensorGroupAssignmentPostSensor**](SensorGroupAssignmentPostSensor.md) |  | 
**Type** | **string** |  | 

## Methods

### NewSensorGroupAssignmentPostResponse

`func NewSensorGroupAssignmentPostResponse(id string, group SensorGroupAssignmentPostGroup, sensor SensorGroupAssignmentPostSensor, type_ string, ) *SensorGroupAssignmentPostResponse`

NewSensorGroupAssignmentPostResponse instantiates a new SensorGroupAssignmentPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorGroupAssignmentPostResponseWithDefaults

`func NewSensorGroupAssignmentPostResponseWithDefaults() *SensorGroupAssignmentPostResponse`

NewSensorGroupAssignmentPostResponseWithDefaults instantiates a new SensorGroupAssignmentPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SensorGroupAssignmentPostResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SensorGroupAssignmentPostResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SensorGroupAssignmentPostResponse) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *SensorGroupAssignmentPostResponse) GetGroup() SensorGroupAssignmentPostGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SensorGroupAssignmentPostResponse) GetGroupOk() (*SensorGroupAssignmentPostGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SensorGroupAssignmentPostResponse) SetGroup(v SensorGroupAssignmentPostGroup)`

SetGroup sets Group field to given value.


### GetSensor

`func (o *SensorGroupAssignmentPostResponse) GetSensor() SensorGroupAssignmentPostSensor`

GetSensor returns the Sensor field if non-nil, zero value otherwise.

### GetSensorOk

`func (o *SensorGroupAssignmentPostResponse) GetSensorOk() (*SensorGroupAssignmentPostSensor, bool)`

GetSensorOk returns a tuple with the Sensor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensor

`func (o *SensorGroupAssignmentPostResponse) SetSensor(v SensorGroupAssignmentPostSensor)`

SetSensor sets Sensor field to given value.


### GetType

`func (o *SensorGroupAssignmentPostResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SensorGroupAssignmentPostResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SensorGroupAssignmentPostResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


