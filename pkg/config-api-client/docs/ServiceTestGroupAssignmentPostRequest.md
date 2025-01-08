# ServiceTestGroupAssignmentPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The unique identifier of the group | 
**ServiceTestId** | **string** | The unique identifier of the service test | 

## Methods

### NewServiceTestGroupAssignmentPostRequest

`func NewServiceTestGroupAssignmentPostRequest(groupId string, serviceTestId string, ) *ServiceTestGroupAssignmentPostRequest`

NewServiceTestGroupAssignmentPostRequest instantiates a new ServiceTestGroupAssignmentPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTestGroupAssignmentPostRequestWithDefaults

`func NewServiceTestGroupAssignmentPostRequestWithDefaults() *ServiceTestGroupAssignmentPostRequest`

NewServiceTestGroupAssignmentPostRequestWithDefaults instantiates a new ServiceTestGroupAssignmentPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ServiceTestGroupAssignmentPostRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ServiceTestGroupAssignmentPostRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ServiceTestGroupAssignmentPostRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetServiceTestId

`func (o *ServiceTestGroupAssignmentPostRequest) GetServiceTestId() string`

GetServiceTestId returns the ServiceTestId field if non-nil, zero value otherwise.

### GetServiceTestIdOk

`func (o *ServiceTestGroupAssignmentPostRequest) GetServiceTestIdOk() (*string, bool)`

GetServiceTestIdOk returns a tuple with the ServiceTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTestId

`func (o *ServiceTestGroupAssignmentPostRequest) SetServiceTestId(v string)`

SetServiceTestId sets ServiceTestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


