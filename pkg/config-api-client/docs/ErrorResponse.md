# ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebugId** | **string** |  | 
**ErrorCode** | **string** |  | 
**HttpStatusCode** | **int32** |  | 
**Message** | **string** |  | 
**ErrorDetails** | Pointer to [**[]ErrorDetail**](ErrorDetail.md) |  | [optional] 

## Methods

### NewErrorResponse

`func NewErrorResponse(debugId string, errorCode string, httpStatusCode int32, message string, ) *ErrorResponse`

NewErrorResponse instantiates a new ErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseWithDefaults

`func NewErrorResponseWithDefaults() *ErrorResponse`

NewErrorResponseWithDefaults instantiates a new ErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebugId

`func (o *ErrorResponse) GetDebugId() string`

GetDebugId returns the DebugId field if non-nil, zero value otherwise.

### GetDebugIdOk

`func (o *ErrorResponse) GetDebugIdOk() (*string, bool)`

GetDebugIdOk returns a tuple with the DebugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugId

`func (o *ErrorResponse) SetDebugId(v string)`

SetDebugId sets DebugId field to given value.


### GetErrorCode

`func (o *ErrorResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ErrorResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ErrorResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetHttpStatusCode

`func (o *ErrorResponse) GetHttpStatusCode() int32`

GetHttpStatusCode returns the HttpStatusCode field if non-nil, zero value otherwise.

### GetHttpStatusCodeOk

`func (o *ErrorResponse) GetHttpStatusCodeOk() (*int32, bool)`

GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatusCode

`func (o *ErrorResponse) SetHttpStatusCode(v int32)`

SetHttpStatusCode sets HttpStatusCode field to given value.


### GetMessage

`func (o *ErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrorDetails

`func (o *ErrorResponse) GetErrorDetails() []ErrorDetail`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *ErrorResponse) GetErrorDetailsOk() (*[]ErrorDetail, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *ErrorResponse) SetErrorDetails(v []ErrorDetail)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *ErrorResponse) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


