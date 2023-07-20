# TTSBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**Settings** | [**TTSSettings**](TTSSettings.md) |  | 

## Methods

### NewTTSBody

`func NewTTSBody(text string, settings TTSSettings, ) *TTSBody`

NewTTSBody instantiates a new TTSBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTTSBodyWithDefaults

`func NewTTSBodyWithDefaults() *TTSBody`

NewTTSBodyWithDefaults instantiates a new TTSBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *TTSBody) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TTSBody) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TTSBody) SetText(v string)`

SetText sets Text field to given value.


### GetSettings

`func (o *TTSBody) GetSettings() TTSSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *TTSBody) GetSettingsOk() (*TTSSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *TTSBody) SetSettings(v TTSSettings)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


