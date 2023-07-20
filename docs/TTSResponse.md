# TTSResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**ResponseStatusCode** | Pointer to [**ResponseStatusCode**](ResponseStatusCode.md) |  | [optional] 
**TaskStatus** | Pointer to [**TTSTaskStatus**](TTSTaskStatus.md) |  | [optional] 
**Results** | Pointer to [**[]TTSResult**](TTSResult.md) |  | [optional] [readonly] 
**Error** | Pointer to [**NullableTTSError**](TTSError.md) |  | [optional] 

## Methods

### NewTTSResponse

`func NewTTSResponse() *TTSResponse`

NewTTSResponse instantiates a new TTSResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTTSResponseWithDefaults

`func NewTTSResponseWithDefaults() *TTSResponse`

NewTTSResponseWithDefaults instantiates a new TTSResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TTSResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TTSResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TTSResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TTSResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TTSResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TTSResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetResponseStatusCode

`func (o *TTSResponse) GetResponseStatusCode() ResponseStatusCode`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *TTSResponse) GetResponseStatusCodeOk() (*ResponseStatusCode, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *TTSResponse) SetResponseStatusCode(v ResponseStatusCode)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *TTSResponse) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetTaskStatus

`func (o *TTSResponse) GetTaskStatus() TTSTaskStatus`

GetTaskStatus returns the TaskStatus field if non-nil, zero value otherwise.

### GetTaskStatusOk

`func (o *TTSResponse) GetTaskStatusOk() (*TTSTaskStatus, bool)`

GetTaskStatusOk returns a tuple with the TaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatus

`func (o *TTSResponse) SetTaskStatus(v TTSTaskStatus)`

SetTaskStatus sets TaskStatus field to given value.

### HasTaskStatus

`func (o *TTSResponse) HasTaskStatus() bool`

HasTaskStatus returns a boolean if a field has been set.

### GetResults

`func (o *TTSResponse) GetResults() []TTSResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TTSResponse) GetResultsOk() (*[]TTSResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TTSResponse) SetResults(v []TTSResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *TTSResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *TTSResponse) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *TTSResponse) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetError

`func (o *TTSResponse) GetError() TTSError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TTSResponse) GetErrorOk() (*TTSError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TTSResponse) SetError(v TTSError)`

SetError sets Error field to given value.

### HasError

`func (o *TTSResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *TTSResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *TTSResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


