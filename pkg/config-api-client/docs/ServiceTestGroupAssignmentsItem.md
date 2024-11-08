# ServiceTestGroupAssignmentsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Group** | [**Group**](Group.md) |  | 
**ServiceTest** | [**ServiceTest**](ServiceTest.md) |  | 
**Type** | **string** |  | 

## Methods

### NewServiceTestGroupAssignmentsItem

`func NewServiceTestGroupAssignmentsItem(id string, group Group, serviceTest ServiceTest, type_ string, ) *ServiceTestGroupAssignmentsItem`

NewServiceTestGroupAssignmentsItem instantiates a new ServiceTestGroupAssignmentsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTestGroupAssignmentsItemWithDefaults

`func NewServiceTestGroupAssignmentsItemWithDefaults() *ServiceTestGroupAssignmentsItem`

NewServiceTestGroupAssignmentsItemWithDefaults instantiates a new ServiceTestGroupAssignmentsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceTestGroupAssignmentsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceTestGroupAssignmentsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceTestGroupAssignmentsItem) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *ServiceTestGroupAssignmentsItem) GetGroup() Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ServiceTestGroupAssignmentsItem) GetGroupOk() (*Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ServiceTestGroupAssignmentsItem) SetGroup(v Group)`

SetGroup sets Group field to given value.


### GetServiceTest

`func (o *ServiceTestGroupAssignmentsItem) GetServiceTest() ServiceTest`

GetServiceTest returns the ServiceTest field if non-nil, zero value otherwise.

### GetServiceTestOk

`func (o *ServiceTestGroupAssignmentsItem) GetServiceTestOk() (*ServiceTest, bool)`

GetServiceTestOk returns a tuple with the ServiceTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTest

`func (o *ServiceTestGroupAssignmentsItem) SetServiceTest(v ServiceTest)`

SetServiceTest sets ServiceTest field to given value.


### GetType

`func (o *ServiceTestGroupAssignmentsItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceTestGroupAssignmentsItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceTestGroupAssignmentsItem) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


