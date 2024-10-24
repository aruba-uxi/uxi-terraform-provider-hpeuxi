# ServiceTestsListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Category** | **string** |  | 
**Name** | **string** |  | 
**Target** | **NullableString** |  | 
**Template** | **string** |  | 
**IsEnabled** | **bool** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "uxi/service-test"]

## Methods

### NewServiceTestsListItem

`func NewServiceTestsListItem(id string, category string, name string, target NullableString, template string, isEnabled bool, ) *ServiceTestsListItem`

NewServiceTestsListItem instantiates a new ServiceTestsListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTestsListItemWithDefaults

`func NewServiceTestsListItemWithDefaults() *ServiceTestsListItem`

NewServiceTestsListItemWithDefaults instantiates a new ServiceTestsListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceTestsListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceTestsListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceTestsListItem) SetId(v string)`

SetId sets Id field to given value.


### GetCategory

`func (o *ServiceTestsListItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ServiceTestsListItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ServiceTestsListItem) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetName

`func (o *ServiceTestsListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceTestsListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceTestsListItem) SetName(v string)`

SetName sets Name field to given value.


### GetTarget

`func (o *ServiceTestsListItem) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ServiceTestsListItem) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ServiceTestsListItem) SetTarget(v string)`

SetTarget sets Target field to given value.


### SetTargetNil

`func (o *ServiceTestsListItem) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *ServiceTestsListItem) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetTemplate

`func (o *ServiceTestsListItem) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ServiceTestsListItem) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ServiceTestsListItem) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetIsEnabled

`func (o *ServiceTestsListItem) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ServiceTestsListItem) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ServiceTestsListItem) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetType

`func (o *ServiceTestsListItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceTestsListItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceTestsListItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceTestsListItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


