# OCRResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The specific Task ID that result was made for | [optional] 
**ResponseStatusCode** | Pointer to [**ResponseStatusCode**](ResponseStatusCode.md) |  | [optional] 
**TaskStatus** | Pointer to [**OCRTaskStatus**](OCRTaskStatus.md) |  | [optional] 
**Results** | Pointer to [**[]OCRResult**](OCRResult.md) | List of results - Especially Text or PDF files | [optional] [readonly] 
**Error** | Pointer to [**NullableOCRError**](OCRError.md) |  | [optional] 

## Methods

### NewOCRResponse

`func NewOCRResponse() *OCRResponse`

NewOCRResponse instantiates a new OCRResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRResponseWithDefaults

`func NewOCRResponseWithDefaults() *OCRResponse`

NewOCRResponseWithDefaults instantiates a new OCRResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OCRResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OCRResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OCRResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OCRResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OCRResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OCRResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetResponseStatusCode

`func (o *OCRResponse) GetResponseStatusCode() ResponseStatusCode`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *OCRResponse) GetResponseStatusCodeOk() (*ResponseStatusCode, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *OCRResponse) SetResponseStatusCode(v ResponseStatusCode)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *OCRResponse) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetTaskStatus

`func (o *OCRResponse) GetTaskStatus() OCRTaskStatus`

GetTaskStatus returns the TaskStatus field if non-nil, zero value otherwise.

### GetTaskStatusOk

`func (o *OCRResponse) GetTaskStatusOk() (*OCRTaskStatus, bool)`

GetTaskStatusOk returns a tuple with the TaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatus

`func (o *OCRResponse) SetTaskStatus(v OCRTaskStatus)`

SetTaskStatus sets TaskStatus field to given value.

### HasTaskStatus

`func (o *OCRResponse) HasTaskStatus() bool`

HasTaskStatus returns a boolean if a field has been set.

### GetResults

`func (o *OCRResponse) GetResults() []OCRResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *OCRResponse) GetResultsOk() (*[]OCRResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *OCRResponse) SetResults(v []OCRResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *OCRResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *OCRResponse) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *OCRResponse) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetError

`func (o *OCRResponse) GetError() OCRError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OCRResponse) GetErrorOk() (*OCRError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OCRResponse) SetError(v OCRError)`

SetError sets Error field to given value.

### HasError

`func (o *OCRResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *OCRResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *OCRResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


