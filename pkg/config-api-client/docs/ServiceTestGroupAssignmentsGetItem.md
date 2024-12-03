# ServiceTestGroupAssignmentsGetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Group** | [**ServiceTestGroupAssignmentsGetGroup**](ServiceTestGroupAssignmentsGetGroup.md) |  | 
**ServiceTest** | [**ServiceTestGroupAssignmentsGetServiceTest**](ServiceTestGroupAssignmentsGetServiceTest.md) |  | 
**Type** | **string** |  | 

## Methods

### NewServiceTestGroupAssignmentsGetItem

`func NewServiceTestGroupAssignmentsGetItem(id string, group ServiceTestGroupAssignmentsGetGroup, serviceTest ServiceTestGroupAssignmentsGetServiceTest, type_ string, ) *ServiceTestGroupAssignmentsGetItem`

NewServiceTestGroupAssignmentsGetItem instantiates a new ServiceTestGroupAssignmentsGetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTestGroupAssignmentsGetItemWithDefaults

`func NewServiceTestGroupAssignmentsGetItemWithDefaults() *ServiceTestGroupAssignmentsGetItem`

NewServiceTestGroupAssignmentsGetItemWithDefaults instantiates a new ServiceTestGroupAssignmentsGetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceTestGroupAssignmentsGetItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceTestGroupAssignmentsGetItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceTestGroupAssignmentsGetItem) SetId(v string)`

SetId sets Id field to given value.


### GetGroup

`func (o *ServiceTestGroupAssignmentsGetItem) GetGroup() ServiceTestGroupAssignmentsGetGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ServiceTestGroupAssignmentsGetItem) GetGroupOk() (*ServiceTestGroupAssignmentsGetGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ServiceTestGroupAssignmentsGetItem) SetGroup(v ServiceTestGroupAssignmentsGetGroup)`

SetGroup sets Group field to given value.


### GetServiceTest

`func (o *ServiceTestGroupAssignmentsGetItem) GetServiceTest() ServiceTestGroupAssignmentsGetServiceTest`

GetServiceTest returns the ServiceTest field if non-nil, zero value otherwise.

### GetServiceTestOk

`func (o *ServiceTestGroupAssignmentsGetItem) GetServiceTestOk() (*ServiceTestGroupAssignmentsGetServiceTest, bool)`

GetServiceTestOk returns a tuple with the ServiceTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTest

`func (o *ServiceTestGroupAssignmentsGetItem) SetServiceTest(v ServiceTestGroupAssignmentsGetServiceTest)`

SetServiceTest sets ServiceTest field to given value.


### GetType

`func (o *ServiceTestGroupAssignmentsGetItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceTestGroupAssignmentsGetItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceTestGroupAssignmentsGetItem) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


