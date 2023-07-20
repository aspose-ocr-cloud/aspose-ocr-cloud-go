# OCRSettingsRecognizeFont

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to [**Language**](Language.md) |  | [optional] 
**MakeSkewCorrect** | Pointer to **bool** |  | [optional] [default to true]
**MakeSpellCheck** | Pointer to **bool** |  | [optional] [default to false]
**MakeContrastCorrection** | Pointer to **bool** |  | [optional] [default to true]
**ResultType** | Pointer to [**ResultType**](ResultType.md) |  | [optional] 
**Rotate** | Pointer to **int32** |  | [optional] 
**MakeBinarization** | Pointer to **bool** |  | [optional] [default to true]
**MakeUpsampling** | Pointer to **bool** | Option to enable image up-sampling algorithm to improve quality. True by default | [optional] [default to false]
**DsrMode** | Pointer to [**DsrMode**](DsrMode.md) |  | [optional] 
**DsrConfidence** | Pointer to [**DsrConfidence**](DsrConfidence.md) |  | [optional] 
**ResultTypeTable** | Pointer to [**ResultTypeTable**](ResultTypeTable.md) |  | [optional] 
**Regions** | Pointer to [**[]OCRRegion**](OCRRegion.md) |  | [optional] 

## Methods

### NewOCRSettingsRecognizeFont

`func NewOCRSettingsRecognizeFont() *OCRSettingsRecognizeFont`

NewOCRSettingsRecognizeFont instantiates a new OCRSettingsRecognizeFont object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRSettingsRecognizeFontWithDefaults

`func NewOCRSettingsRecognizeFontWithDefaults() *OCRSettingsRecognizeFont`

NewOCRSettingsRecognizeFontWithDefaults instantiates a new OCRSettingsRecognizeFont object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *OCRSettingsRecognizeFont) GetLanguage() Language`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *OCRSettingsRecognizeFont) GetLanguageOk() (*Language, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *OCRSettingsRecognizeFont) SetLanguage(v Language)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *OCRSettingsRecognizeFont) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMakeSkewCorrect

`func (o *OCRSettingsRecognizeFont) GetMakeSkewCorrect() bool`

GetMakeSkewCorrect returns the MakeSkewCorrect field if non-nil, zero value otherwise.

### GetMakeSkewCorrectOk

`func (o *OCRSettingsRecognizeFont) GetMakeSkewCorrectOk() (*bool, bool)`

GetMakeSkewCorrectOk returns a tuple with the MakeSkewCorrect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeSkewCorrect

`func (o *OCRSettingsRecognizeFont) SetMakeSkewCorrect(v bool)`

SetMakeSkewCorrect sets MakeSkewCorrect field to given value.

### HasMakeSkewCorrect

`func (o *OCRSettingsRecognizeFont) HasMakeSkewCorrect() bool`

HasMakeSkewCorrect returns a boolean if a field has been set.

### GetMakeSpellCheck

`func (o *OCRSettingsRecognizeFont) GetMakeSpellCheck() bool`

GetMakeSpellCheck returns the MakeSpellCheck field if non-nil, zero value otherwise.

### GetMakeSpellCheckOk

`func (o *OCRSettingsRecognizeFont) GetMakeSpellCheckOk() (*bool, bool)`

GetMakeSpellCheckOk returns a tuple with the MakeSpellCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeSpellCheck

`func (o *OCRSettingsRecognizeFont) SetMakeSpellCheck(v bool)`

SetMakeSpellCheck sets MakeSpellCheck field to given value.

### HasMakeSpellCheck

`func (o *OCRSettingsRecognizeFont) HasMakeSpellCheck() bool`

HasMakeSpellCheck returns a boolean if a field has been set.

### GetMakeContrastCorrection

`func (o *OCRSettingsRecognizeFont) GetMakeContrastCorrection() bool`

GetMakeContrastCorrection returns the MakeContrastCorrection field if non-nil, zero value otherwise.

### GetMakeContrastCorrectionOk

`func (o *OCRSettingsRecognizeFont) GetMakeContrastCorrectionOk() (*bool, bool)`

GetMakeContrastCorrectionOk returns a tuple with the MakeContrastCorrection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeContrastCorrection

`func (o *OCRSettingsRecognizeFont) SetMakeContrastCorrection(v bool)`

SetMakeContrastCorrection sets MakeContrastCorrection field to given value.

### HasMakeContrastCorrection

`func (o *OCRSettingsRecognizeFont) HasMakeContrastCorrection() bool`

HasMakeContrastCorrection returns a boolean if a field has been set.

### GetResultType

`func (o *OCRSettingsRecognizeFont) GetResultType() ResultType`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *OCRSettingsRecognizeFont) GetResultTypeOk() (*ResultType, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *OCRSettingsRecognizeFont) SetResultType(v ResultType)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *OCRSettingsRecognizeFont) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetRotate

`func (o *OCRSettingsRecognizeFont) GetRotate() int32`

GetRotate returns the Rotate field if non-nil, zero value otherwise.

### GetRotateOk

`func (o *OCRSettingsRecognizeFont) GetRotateOk() (*int32, bool)`

GetRotateOk returns a tuple with the Rotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotate

`func (o *OCRSettingsRecognizeFont) SetRotate(v int32)`

SetRotate sets Rotate field to given value.

### HasRotate

`func (o *OCRSettingsRecognizeFont) HasRotate() bool`

HasRotate returns a boolean if a field has been set.

### GetMakeBinarization

`func (o *OCRSettingsRecognizeFont) GetMakeBinarization() bool`

GetMakeBinarization returns the MakeBinarization field if non-nil, zero value otherwise.

### GetMakeBinarizationOk

`func (o *OCRSettingsRecognizeFont) GetMakeBinarizationOk() (*bool, bool)`

GetMakeBinarizationOk returns a tuple with the MakeBinarization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeBinarization

`func (o *OCRSettingsRecognizeFont) SetMakeBinarization(v bool)`

SetMakeBinarization sets MakeBinarization field to given value.

### HasMakeBinarization

`func (o *OCRSettingsRecognizeFont) HasMakeBinarization() bool`

HasMakeBinarization returns a boolean if a field has been set.

### GetMakeUpsampling

`func (o *OCRSettingsRecognizeFont) GetMakeUpsampling() bool`

GetMakeUpsampling returns the MakeUpsampling field if non-nil, zero value otherwise.

### GetMakeUpsamplingOk

`func (o *OCRSettingsRecognizeFont) GetMakeUpsamplingOk() (*bool, bool)`

GetMakeUpsamplingOk returns a tuple with the MakeUpsampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeUpsampling

`func (o *OCRSettingsRecognizeFont) SetMakeUpsampling(v bool)`

SetMakeUpsampling sets MakeUpsampling field to given value.

### HasMakeUpsampling

`func (o *OCRSettingsRecognizeFont) HasMakeUpsampling() bool`

HasMakeUpsampling returns a boolean if a field has been set.

### GetDsrMode

`func (o *OCRSettingsRecognizeFont) GetDsrMode() DsrMode`

GetDsrMode returns the DsrMode field if non-nil, zero value otherwise.

### GetDsrModeOk

`func (o *OCRSettingsRecognizeFont) GetDsrModeOk() (*DsrMode, bool)`

GetDsrModeOk returns a tuple with the DsrMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsrMode

`func (o *OCRSettingsRecognizeFont) SetDsrMode(v DsrMode)`

SetDsrMode sets DsrMode field to given value.

### HasDsrMode

`func (o *OCRSettingsRecognizeFont) HasDsrMode() bool`

HasDsrMode returns a boolean if a field has been set.

### GetDsrConfidence

`func (o *OCRSettingsRecognizeFont) GetDsrConfidence() DsrConfidence`

GetDsrConfidence returns the DsrConfidence field if non-nil, zero value otherwise.

### GetDsrConfidenceOk

`func (o *OCRSettingsRecognizeFont) GetDsrConfidenceOk() (*DsrConfidence, bool)`

GetDsrConfidenceOk returns a tuple with the DsrConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsrConfidence

`func (o *OCRSettingsRecognizeFont) SetDsrConfidence(v DsrConfidence)`

SetDsrConfidence sets DsrConfidence field to given value.

### HasDsrConfidence

`func (o *OCRSettingsRecognizeFont) HasDsrConfidence() bool`

HasDsrConfidence returns a boolean if a field has been set.

### GetResultTypeTable

`func (o *OCRSettingsRecognizeFont) GetResultTypeTable() ResultTypeTable`

GetResultTypeTable returns the ResultTypeTable field if non-nil, zero value otherwise.

### GetResultTypeTableOk

`func (o *OCRSettingsRecognizeFont) GetResultTypeTableOk() (*ResultTypeTable, bool)`

GetResultTypeTableOk returns a tuple with the ResultTypeTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTypeTable

`func (o *OCRSettingsRecognizeFont) SetResultTypeTable(v ResultTypeTable)`

SetResultTypeTable sets ResultTypeTable field to given value.

### HasResultTypeTable

`func (o *OCRSettingsRecognizeFont) HasResultTypeTable() bool`

HasResultTypeTable returns a boolean if a field has been set.

### GetRegions

`func (o *OCRSettingsRecognizeFont) GetRegions() []OCRRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *OCRSettingsRecognizeFont) GetRegionsOk() (*[]OCRRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *OCRSettingsRecognizeFont) SetRegions(v []OCRRegion)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *OCRSettingsRecognizeFont) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### SetRegionsNil

`func (o *OCRSettingsRecognizeFont) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *OCRSettingsRecognizeFont) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


