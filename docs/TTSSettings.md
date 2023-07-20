# TTSSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | [**LanguageTTS**](LanguageTTS.md) |  | 
**ResultType** | [**ResultTypeTTS**](ResultTypeTTS.md) |  | 

## Methods

### NewTTSSettings

`func NewTTSSettings(language LanguageTTS, resultType ResultTypeTTS, ) *TTSSettings`

NewTTSSettings instantiates a new TTSSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTTSSettingsWithDefaults

`func NewTTSSettingsWithDefaults() *TTSSettings`

NewTTSSettingsWithDefaults instantiates a new TTSSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *TTSSettings) GetLanguage() LanguageTTS`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *TTSSettings) GetLanguageOk() (*LanguageTTS, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *TTSSettings) SetLanguage(v LanguageTTS)`

SetLanguage sets Language field to given value.


### GetResultType

`func (o *TTSSettings) GetResultType() ResultTypeTTS`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *TTSSettings) GetResultTypeOk() (*ResultTypeTTS, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *TTSSettings) SetResultType(v ResultTypeTTS)`

SetResultType sets ResultType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


