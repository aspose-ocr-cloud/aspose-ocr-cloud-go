# OCRSettingsDetectRegions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MakeSkewCorrect** | Pointer to **bool** |  | [optional] [default to true]
**MakeContrastCorrection** | Pointer to **bool** |  | [optional] [default to true]
**MakeUpsampling** | Pointer to **bool** |  | [optional] [default to false]
**DsrConfidence** | Pointer to [**DsrConfidence**](DsrConfidence.md) |  | [optional] 
**Language** | Pointer to [**Language**](Language.md) |  | [optional] 
**Rotate** | Pointer to **int32** |  | [optional] 
**MakeSpellCheck** | Pointer to **bool** | Option to enable spell checking and correction algorithm. False by default | [optional] [default to false]
**MakeBinarization** | Pointer to **bool** |  | [optional] [default to true]
**DsrMode** | Pointer to [**DsrMode**](DsrMode.md) |  | [optional] 
**ResultType** | Pointer to [**ResultType**](ResultType.md) |  | [optional] 
**ResultTypeTable** | Pointer to [**ResultTypeTable**](ResultTypeTable.md) |  | [optional] 
**Regions** | Pointer to [**[]OCRRegion**](OCRRegion.md) |  | [optional] 

## Methods

### NewOCRSettingsDetectRegions

`func NewOCRSettingsDetectRegions() *OCRSettingsDetectRegions`

NewOCRSettingsDetectRegions instantiates a new OCRSettingsDetectRegions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRSettingsDetectRegionsWithDefaults

`func NewOCRSettingsDetectRegionsWithDefaults() *OCRSettingsDetectRegions`

NewOCRSettingsDetectRegionsWithDefaults instantiates a new OCRSettingsDetectRegions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMakeSkewCorrect

`func (o *OCRSettingsDetectRegions) GetMakeSkewCorrect() bool`

GetMakeSkewCorrect returns the MakeSkewCorrect field if non-nil, zero value otherwise.

### GetMakeSkewCorrectOk

`func (o *OCRSettingsDetectRegions) GetMakeSkewCorrectOk() (*bool, bool)`

GetMakeSkewCorrectOk returns a tuple with the MakeSkewCorrect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeSkewCorrect

`func (o *OCRSettingsDetectRegions) SetMakeSkewCorrect(v bool)`

SetMakeSkewCorrect sets MakeSkewCorrect field to given value.

### HasMakeSkewCorrect

`func (o *OCRSettingsDetectRegions) HasMakeSkewCorrect() bool`

HasMakeSkewCorrect returns a boolean if a field has been set.

### GetMakeContrastCorrection

`func (o *OCRSettingsDetectRegions) GetMakeContrastCorrection() bool`

GetMakeContrastCorrection returns the MakeContrastCorrection field if non-nil, zero value otherwise.

### GetMakeContrastCorrectionOk

`func (o *OCRSettingsDetectRegions) GetMakeContrastCorrectionOk() (*bool, bool)`

GetMakeContrastCorrectionOk returns a tuple with the MakeContrastCorrection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeContrastCorrection

`func (o *OCRSettingsDetectRegions) SetMakeContrastCorrection(v bool)`

SetMakeContrastCorrection sets MakeContrastCorrection field to given value.

### HasMakeContrastCorrection

`func (o *OCRSettingsDetectRegions) HasMakeContrastCorrection() bool`

HasMakeContrastCorrection returns a boolean if a field has been set.

### GetMakeUpsampling

`func (o *OCRSettingsDetectRegions) GetMakeUpsampling() bool`

GetMakeUpsampling returns the MakeUpsampling field if non-nil, zero value otherwise.

### GetMakeUpsamplingOk

`func (o *OCRSettingsDetectRegions) GetMakeUpsamplingOk() (*bool, bool)`

GetMakeUpsamplingOk returns a tuple with the MakeUpsampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeUpsampling

`func (o *OCRSettingsDetectRegions) SetMakeUpsampling(v bool)`

SetMakeUpsampling sets MakeUpsampling field to given value.

### HasMakeUpsampling

`func (o *OCRSettingsDetectRegions) HasMakeUpsampling() bool`

HasMakeUpsampling returns a boolean if a field has been set.

### GetDsrConfidence

`func (o *OCRSettingsDetectRegions) GetDsrConfidence() DsrConfidence`

GetDsrConfidence returns the DsrConfidence field if non-nil, zero value otherwise.

### GetDsrConfidenceOk

`func (o *OCRSettingsDetectRegions) GetDsrConfidenceOk() (*DsrConfidence, bool)`

GetDsrConfidenceOk returns a tuple with the DsrConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsrConfidence

`func (o *OCRSettingsDetectRegions) SetDsrConfidence(v DsrConfidence)`

SetDsrConfidence sets DsrConfidence field to given value.

### HasDsrConfidence

`func (o *OCRSettingsDetectRegions) HasDsrConfidence() bool`

HasDsrConfidence returns a boolean if a field has been set.

### GetLanguage

`func (o *OCRSettingsDetectRegions) GetLanguage() Language`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *OCRSettingsDetectRegions) GetLanguageOk() (*Language, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *OCRSettingsDetectRegions) SetLanguage(v Language)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *OCRSettingsDetectRegions) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRotate

`func (o *OCRSettingsDetectRegions) GetRotate() int32`

GetRotate returns the Rotate field if non-nil, zero value otherwise.

### GetRotateOk

`func (o *OCRSettingsDetectRegions) GetRotateOk() (*int32, bool)`

GetRotateOk returns a tuple with the Rotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotate

`func (o *OCRSettingsDetectRegions) SetRotate(v int32)`

SetRotate sets Rotate field to given value.

### HasRotate

`func (o *OCRSettingsDetectRegions) HasRotate() bool`

HasRotate returns a boolean if a field has been set.

### GetMakeSpellCheck

`func (o *OCRSettingsDetectRegions) GetMakeSpellCheck() bool`

GetMakeSpellCheck returns the MakeSpellCheck field if non-nil, zero value otherwise.

### GetMakeSpellCheckOk

`func (o *OCRSettingsDetectRegions) GetMakeSpellCheckOk() (*bool, bool)`

GetMakeSpellCheckOk returns a tuple with the MakeSpellCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeSpellCheck

`func (o *OCRSettingsDetectRegions) SetMakeSpellCheck(v bool)`

SetMakeSpellCheck sets MakeSpellCheck field to given value.

### HasMakeSpellCheck

`func (o *OCRSettingsDetectRegions) HasMakeSpellCheck() bool`

HasMakeSpellCheck returns a boolean if a field has been set.

### GetMakeBinarization

`func (o *OCRSettingsDetectRegions) GetMakeBinarization() bool`

GetMakeBinarization returns the MakeBinarization field if non-nil, zero value otherwise.

### GetMakeBinarizationOk

`func (o *OCRSettingsDetectRegions) GetMakeBinarizationOk() (*bool, bool)`

GetMakeBinarizationOk returns a tuple with the MakeBinarization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeBinarization

`func (o *OCRSettingsDetectRegions) SetMakeBinarization(v bool)`

SetMakeBinarization sets MakeBinarization field to given value.

### HasMakeBinarization

`func (o *OCRSettingsDetectRegions) HasMakeBinarization() bool`

HasMakeBinarization returns a boolean if a field has been set.

### GetDsrMode

`func (o *OCRSettingsDetectRegions) GetDsrMode() DsrMode`

GetDsrMode returns the DsrMode field if non-nil, zero value otherwise.

### GetDsrModeOk

`func (o *OCRSettingsDetectRegions) GetDsrModeOk() (*DsrMode, bool)`

GetDsrModeOk returns a tuple with the DsrMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsrMode

`func (o *OCRSettingsDetectRegions) SetDsrMode(v DsrMode)`

SetDsrMode sets DsrMode field to given value.

### HasDsrMode

`func (o *OCRSettingsDetectRegions) HasDsrMode() bool`

HasDsrMode returns a boolean if a field has been set.

### GetResultType

`func (o *OCRSettingsDetectRegions) GetResultType() ResultType`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *OCRSettingsDetectRegions) GetResultTypeOk() (*ResultType, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *OCRSettingsDetectRegions) SetResultType(v ResultType)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *OCRSettingsDetectRegions) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetResultTypeTable

`func (o *OCRSettingsDetectRegions) GetResultTypeTable() ResultTypeTable`

GetResultTypeTable returns the ResultTypeTable field if non-nil, zero value otherwise.

### GetResultTypeTableOk

`func (o *OCRSettingsDetectRegions) GetResultTypeTableOk() (*ResultTypeTable, bool)`

GetResultTypeTableOk returns a tuple with the ResultTypeTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTypeTable

`func (o *OCRSettingsDetectRegions) SetResultTypeTable(v ResultTypeTable)`

SetResultTypeTable sets ResultTypeTable field to given value.

### HasResultTypeTable

`func (o *OCRSettingsDetectRegions) HasResultTypeTable() bool`

HasResultTypeTable returns a boolean if a field has been set.

### GetRegions

`func (o *OCRSettingsDetectRegions) GetRegions() []OCRRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *OCRSettingsDetectRegions) GetRegionsOk() (*[]OCRRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *OCRSettingsDetectRegions) SetRegions(v []OCRRegion)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *OCRSettingsDetectRegions) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### SetRegionsNil

`func (o *OCRSettingsDetectRegions) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *OCRSettingsDetectRegions) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


