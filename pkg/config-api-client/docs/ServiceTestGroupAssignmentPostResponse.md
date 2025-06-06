# ServiceTestGroupAssignmentPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the service test group assignment | 
**Group** | [**ServiceTestGroupAssignmentPostGroup**](ServiceTestGroupAssignmentPostGroup.md) | The group component of the service test group assignment | 
**ServiceTest** | [**ServiceTestGroupAssignmentPostServiceTest**](ServiceTestGroupAssignmentPostServiceTest.md) | The service test component of the service test group assignment | 
**Type** | **string** | The type of the resource. | 

## Methods

### NewServiceTestGroupAssignmentPostResponse

`func NewServiceTestGroupAssignmentPostResponse(id string, group ServiceTestGroupAssignmentPostGroup, serviceTest ServiceTestGroupAssignmentPostServiceTest, type_ string, ) *ServiceTestGroupAssignmentPostResponse`

NewServiceTestGroupAssignmentPostResponse instantiates a new ServiceTestGroupAssignmentPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTestGroupAssignmentPostResponseWithDefaults

`func NewServiceTestGroupAssignmentPostResponseWithDefaults() *ServiceTestGroupAssignmentPostResponse`

NewServiceTestGroupAssignmentPostResponseWithDefaults instantiates a new ServiceTestGroupAssignmentPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceTestGroupAssignmentPostResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceTestGroupAssignmentPostResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceTestGroupAssignmentPostResponse) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *ServiceTestGroupAssignmentPostResponse) GetGroup() ServiceTestGroupAssignmentPostGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ServiceTestGroupAssignmentPostResponse) GetGroupOk() (*ServiceTestGroupAssignmentPostGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ServiceTestGroupAssignmentPostResponse) SetGroup(v ServiceTestGroupAssignmentPostGroup)`

SetGroup sets Group field to given value.


### GetServiceTest

`func (o *ServiceTestGroupAssignmentPostResponse) GetServiceTest() ServiceTestGroupAssignmentPostServiceTest`

GetServiceTest returns the ServiceTest field if non-nil, zero value otherwise.

### GetServiceTestOk

`func (o *ServiceTestGroupAssignmentPostResponse) GetServiceTestOk() (*ServiceTestGroupAssignmentPostServiceTest, bool)`

GetServiceTestOk returns a tuple with the ServiceTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTest

`func (o *ServiceTestGroupAssignmentPostResponse) SetServiceTest(v ServiceTestGroupAssignmentPostServiceTest)`

SetServiceTest sets ServiceTest field to given value.


### GetType

`func (o *ServiceTestGroupAssignmentPostResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceTestGroupAssignmentPostResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceTestGroupAssignmentPostResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


