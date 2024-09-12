# SensorGroupAssignmentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SensorGroupAssignments** | [**[]SensorGroupAssignmentsItem**](SensorGroupAssignmentsItem.md) |  | 
**Pagination** | [**PaginationDetails**](PaginationDetails.md) |  | 

## Methods

### NewSensorGroupAssignmentsResponse

`func NewSensorGroupAssignmentsResponse(sensorGroupAssignments []SensorGroupAssignmentsItem, pagination PaginationDetails, ) *SensorGroupAssignmentsResponse`

NewSensorGroupAssignmentsResponse instantiates a new SensorGroupAssignmentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorGroupAssignmentsResponseWithDefaults

`func NewSensorGroupAssignmentsResponseWithDefaults() *SensorGroupAssignmentsResponse`

NewSensorGroupAssignmentsResponseWithDefaults instantiates a new SensorGroupAssignmentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSensorGroupAssignments

`func (o *SensorGroupAssignmentsResponse) GetSensorGroupAssignments() []SensorGroupAssignmentsItem`

GetSensorGroupAssignments returns the SensorGroupAssignments field if non-nil, zero value otherwise.

### GetSensorGroupAssignmentsOk

`func (o *SensorGroupAssignmentsResponse) GetSensorGroupAssignmentsOk() (*[]SensorGroupAssignmentsItem, bool)`

GetSensorGroupAssignmentsOk returns a tuple with the SensorGroupAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorGroupAssignments

`func (o *SensorGroupAssignmentsResponse) SetSensorGroupAssignments(v []SensorGroupAssignmentsItem)`

SetSensorGroupAssignments sets SensorGroupAssignments field to given value.


### GetPagination

`func (o *SensorGroupAssignmentsResponse) GetPagination() PaginationDetails`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SensorGroupAssignmentsResponse) GetPaginationOk() (*PaginationDetails, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SensorGroupAssignmentsResponse) SetPagination(v PaginationDetails)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


