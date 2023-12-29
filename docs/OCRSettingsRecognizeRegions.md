# OCRSettingsRecognizeRegions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to [**Language**](Language.md) |  | [optional] 
**MakeSkewCorrect** | Pointer to **bool** |  | [optional] [default to true]
**MakeSpellCheck** | Pointer to **bool** |  | [optional] [default to false]
**MakeContrastCorrection** | Pointer to **bool** |  | [optional] [default to false]
**MakeUpsampling** | Pointer to **bool** |  | [optional] [default to false]
**Regions** | Pointer to [**[]OCRRegion**](OCRRegion.md) |  | [optional] 
**ResultType** | Pointer to [**ResultType**](ResultType.md) |  | [optional] 
**Rotate** | Pointer to **int32** |  | [optional] 
**MakeBinarization** | Pointer to **bool** |  | [optional] [default to true]
**DsrMode** | Pointer to [**DsrMode**](DsrMode.md) |  | [optional] 
**DsrConfidence** | Pointer to [**DsrConfidence**](DsrConfidence.md) |  | [optional] 
**ResultTypeTable** | Pointer to [**ResultTypeTable**](ResultTypeTable.md) |  | [optional] 

## Methods

### NewOCRSettingsRecognizeRegions

`func NewOCRSettingsRecognizeRegions() *OCRSettingsRecognizeRegions`

NewOCRSettingsRecognizeRegions instantiates a new OCRSettingsRecognizeRegions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRSettingsRecognizeRegionsWithDefaults

`func NewOCRSettingsRecognizeRegionsWithDefaults() *OCRSettingsRecognizeRegions`

NewOCRSettingsRecognizeRegionsWithDefaults instantiates a new OCRSettingsRecognizeRegions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *OCRSettingsRecognizeRegions) GetLanguage() Language`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *OCRSettingsRecognizeRegions) GetLanguageOk() (*Language, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *OCRSettingsRecognizeRegions) SetLanguage(v Language)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *OCRSettingsRecognizeRegions) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMakeSkewCorrect

`func (o *OCRSettingsRecognizeRegions) GetMakeSkewCorrect() bool`

GetMakeSkewCorrect returns the MakeSkewCorrect field if non-nil, zero value otherwise.

### GetMakeSkewCorrectOk

`func (o *OCRSettingsRecognizeRegions) GetMakeSkewCorrectOk() (*bool, bool)`

GetMakeSkewCorrectOk returns a tuple with the MakeSkewCorrect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeSkewCorrect

`func (o *OCRSettingsRecognizeRegions) SetMakeSkewCorrect(v bool)`

SetMakeSkewCorrect sets MakeSkewCorrect field to given value.

### HasMakeSkewCorrect

`func (o *OCRSettingsRecognizeRegions) HasMakeSkewCorrect() bool`

HasMakeSkewCorrect returns a boolean if a field has been set.

### GetMakeSpellCheck

`func (o *OCRSettingsRecognizeRegions) GetMakeSpellCheck() bool`

GetMakeSpellCheck returns the MakeSpellCheck field if non-nil, zero value otherwise.

### GetMakeSpellCheckOk

`func (o *OCRSettingsRecognizeRegions) GetMakeSpellCheckOk() (*bool, bool)`

GetMakeSpellCheckOk returns a tuple with the MakeSpellCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeSpellCheck

`func (o *OCRSettingsRecognizeRegions) SetMakeSpellCheck(v bool)`

SetMakeSpellCheck sets MakeSpellCheck field to given value.

### HasMakeSpellCheck

`func (o *OCRSettingsRecognizeRegions) HasMakeSpellCheck() bool`

HasMakeSpellCheck returns a boolean if a field has been set.

### GetMakeContrastCorrection

`func (o *OCRSettingsRecognizeRegions) GetMakeContrastCorrection() bool`

GetMakeContrastCorrection returns the MakeContrastCorrection field if non-nil, zero value otherwise.

### GetMakeContrastCorrectionOk

`func (o *OCRSettingsRecognizeRegions) GetMakeContrastCorrectionOk() (*bool, bool)`

GetMakeContrastCorrectionOk returns a tuple with the MakeContrastCorrection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeContrastCorrection

`func (o *OCRSettingsRecognizeRegions) SetMakeContrastCorrection(v bool)`

SetMakeContrastCorrection sets MakeContrastCorrection field to given value.

### HasMakeContrastCorrection

`func (o *OCRSettingsRecognizeRegions) HasMakeContrastCorrection() bool`

HasMakeContrastCorrection returns a boolean if a field has been set.

### GetMakeUpsampling

`func (o *OCRSettingsRecognizeRegions) GetMakeUpsampling() bool`

GetMakeUpsampling returns the MakeUpsampling field if non-nil, zero value otherwise.

### GetMakeUpsamplingOk

`func (o *OCRSettingsRecognizeRegions) GetMakeUpsamplingOk() (*bool, bool)`

GetMakeUpsamplingOk returns a tuple with the MakeUpsampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeUpsampling

`func (o *OCRSettingsRecognizeRegions) SetMakeUpsampling(v bool)`

SetMakeUpsampling sets MakeUpsampling field to given value.

### HasMakeUpsampling

`func (o *OCRSettingsRecognizeRegions) HasMakeUpsampling() bool`

HasMakeUpsampling returns a boolean if a field has been set.

### GetRegions

`func (o *OCRSettingsRecognizeRegions) GetRegions() []OCRRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *OCRSettingsRecognizeRegions) GetRegionsOk() (*[]OCRRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *OCRSettingsRecognizeRegions) SetRegions(v []OCRRegion)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *OCRSettingsRecognizeRegions) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### SetRegionsNil

`func (o *OCRSettingsRecognizeRegions) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *OCRSettingsRecognizeRegions) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil
### GetResultType

`func (o *OCRSettingsRecognizeRegions) GetResultType() ResultType`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *OCRSettingsRecognizeRegions) GetResultTypeOk() (*ResultType, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *OCRSettingsRecognizeRegions) SetResultType(v ResultType)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *OCRSettingsRecognizeRegions) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetRotate

`func (o *OCRSettingsRecognizeRegions) GetRotate() int32`

GetRotate returns the Rotate field if non-nil, zero value otherwise.

### GetRotateOk

`func (o *OCRSettingsRecognizeRegions) GetRotateOk() (*int32, bool)`

GetRotateOk returns a tuple with the Rotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotate

`func (o *OCRSettingsRecognizeRegions) SetRotate(v int32)`

SetRotate sets Rotate field to given value.

### HasRotate

`func (o *OCRSettingsRecognizeRegions) HasRotate() bool`

HasRotate returns a boolean if a field has been set.

### GetMakeBinarization

`func (o *OCRSettingsRecognizeRegions) GetMakeBinarization() bool`

GetMakeBinarization returns the MakeBinarization field if non-nil, zero value otherwise.

### GetMakeBinarizationOk

`func (o *OCRSettingsRecognizeRegions) GetMakeBinarizationOk() (*bool, bool)`

GetMakeBinarizationOk returns a tuple with the MakeBinarization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeBinarization

`func (o *OCRSettingsRecognizeRegions) SetMakeBinarization(v bool)`

SetMakeBinarization sets MakeBinarization field to given value.

### HasMakeBinarization

`func (o *OCRSettingsRecognizeRegions) HasMakeBinarization() bool`

HasMakeBinarization returns a boolean if a field has been set.

### GetDsrMode

`func (o *OCRSettingsRecognizeRegions) GetDsrMode() DsrMode`

GetDsrMode returns the DsrMode field if non-nil, zero value otherwise.

### GetDsrModeOk

`func (o *OCRSettingsRecognizeRegions) GetDsrModeOk() (*DsrMode, bool)`

GetDsrModeOk returns a tuple with the DsrMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsrMode

`func (o *OCRSettingsRecognizeRegions) SetDsrMode(v DsrMode)`

SetDsrMode sets DsrMode field to given value.

### HasDsrMode

`func (o *OCRSettingsRecognizeRegions) HasDsrMode() bool`

HasDsrMode returns a boolean if a field has been set.

### GetDsrConfidence

`func (o *OCRSettingsRecognizeRegions) GetDsrConfidence() DsrConfidence`

GetDsrConfidence returns the DsrConfidence field if non-nil, zero value otherwise.

### GetDsrConfidenceOk

`func (o *OCRSettingsRecognizeRegions) GetDsrConfidenceOk() (*DsrConfidence, bool)`

GetDsrConfidenceOk returns a tuple with the DsrConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsrConfidence

`func (o *OCRSettingsRecognizeRegions) SetDsrConfidence(v DsrConfidence)`

SetDsrConfidence sets DsrConfidence field to given value.

### HasDsrConfidence

`func (o *OCRSettingsRecognizeRegions) HasDsrConfidence() bool`

HasDsrConfidence returns a boolean if a field has been set.

### GetResultTypeTable

`func (o *OCRSettingsRecognizeRegions) GetResultTypeTable() ResultTypeTable`

GetResultTypeTable returns the ResultTypeTable field if non-nil, zero value otherwise.

### GetResultTypeTableOk

`func (o *OCRSettingsRecognizeRegions) GetResultTypeTableOk() (*ResultTypeTable, bool)`

GetResultTypeTableOk returns a tuple with the ResultTypeTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTypeTable

`func (o *OCRSettingsRecognizeRegions) SetResultTypeTable(v ResultTypeTable)`

SetResultTypeTable sets ResultTypeTable field to given value.

### HasResultTypeTable

`func (o *OCRSettingsRecognizeRegions) HasResultTypeTable() bool`

HasResultTypeTable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


