# ServiceTestsGetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the service test | 
**Category** | **string** | The category of the service test | 
**Name** | **string** | The name of the service test | 
**Target** | **NullableString** | The target of the service test | 
**Template** | **string** | The template of the service test | 
**IsEnabled** | **bool** | Indicates if the service test is enabled | 
**Type** | **string** | The type of the resource. | 

## Methods

### NewServiceTestsGetItem

`func NewServiceTestsGetItem(id string, category string, name string, target NullableString, template string, isEnabled bool, type_ string, ) *ServiceTestsGetItem`

NewServiceTestsGetItem instantiates a new ServiceTestsGetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTestsGetItemWithDefaults

`func NewServiceTestsGetItemWithDefaults() *ServiceTestsGetItem`

NewServiceTestsGetItemWithDefaults instantiates a new ServiceTestsGetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceTestsGetItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceTestsGetItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceTestsGetItem) SetId(v string)`

SetId sets Id field to given value.


### GetCategory

`func (o *ServiceTestsGetItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ServiceTestsGetItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ServiceTestsGetItem) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetName

`func (o *ServiceTestsGetItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceTestsGetItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceTestsGetItem) SetName(v string)`

SetName sets Name field to given value.


### GetTarget

`func (o *ServiceTestsGetItem) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ServiceTestsGetItem) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ServiceTestsGetItem) SetTarget(v string)`

SetTarget sets Target field to given value.


### SetTargetNil

`func (o *ServiceTestsGetItem) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *ServiceTestsGetItem) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetTemplate

`func (o *ServiceTestsGetItem) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ServiceTestsGetItem) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ServiceTestsGetItem) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetIsEnabled

`func (o *ServiceTestsGetItem) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ServiceTestsGetItem) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ServiceTestsGetItem) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetType

`func (o *ServiceTestsGetItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceTestsGetItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceTestsGetItem) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


