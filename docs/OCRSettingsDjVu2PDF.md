# OCRSettingsDjVu2PDF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to [**Language**](Language.md) |  | [optional] 
**Rotate** | Pointer to **int32** |  | [optional] 
**MakeSkewCorrect** | Pointer to **bool** | Option to enable skew correction algorithm. True by default | [optional] [default to true]
**MakeSpellCheck** | Pointer to **bool** | Option to enable spell checking and correction algorithm. False by default | [optional] [default to false]
**MakeContrastCorrection** | Pointer to **bool** | Option to enable image contrast correction algorithm. True by default | [optional] [default to true]
**MakeBinarization** | Pointer to **bool** |  | [optional] [default to true]
**MakeUpsampling** | Pointer to **bool** | Option to enable image up-sampling algorithm to improve quality. True by default | [optional] [default to false]
**DsrMode** | Pointer to [**DsrMode**](DsrMode.md) |  | [optional] 
**DsrConfidence** | Pointer to [**DsrConfidence**](DsrConfidence.md) |  | [optional] 
**ResultType** | Pointer to [**ResultType**](ResultType.md) |  | [optional] 
**ResultTypeTable** | Pointer to [**ResultTypeTable**](ResultTypeTable.md) |  | [optional] 
**Regions** | Pointer to [**[]OCRRegion**](OCRRegion.md) |  | [optional] 

## Methods

### NewOCRSettingsDjVu2PDF

`func NewOCRSettingsDjVu2PDF() *OCRSettingsDjVu2PDF`

NewOCRSettingsDjVu2PDF instantiates a new OCRSettingsDjVu2PDF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRSettingsDjVu2PDFWithDefaults

`func NewOCRSettingsDjVu2PDFWithDefaults() *OCRSettingsDjVu2PDF`

NewOCRSettingsDjVu2PDFWithDefaults instantiates a new OCRSettingsDjVu2PDF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *OCRSettingsDjVu2PDF) GetLanguage() Language`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *OCRSettingsDjVu2PDF) GetLanguageOk() (*Language, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *OCRSettingsDjVu2PDF) SetLanguage(v Language)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *OCRSettingsDjVu2PDF) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRotate

`func (o *OCRSettingsDjVu2PDF) GetRotate() int32`

GetRotate returns the Rotate field if non-nil, zero value otherwise.

### GetRotateOk

`func (o *OCRSettingsDjVu2PDF) GetRotateOk() (*int32, bool)`

GetRotateOk returns a tuple with the Rotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotate

`func (o *OCRSettingsDjVu2PDF) SetRotate(v int32)`

SetRotate sets Rotate field to given value.

### HasRotate

`func (o *OCRSettingsDjVu2PDF) HasRotate() bool`

HasRotate returns a boolean if a field has been set.

### GetMakeSkewCorrect

`func (o *OCRSettingsDjVu2PDF) GetMakeSkewCorrect() bool`

GetMakeSkewCorrect returns the MakeSkewCorrect field if non-nil, zero value otherwise.

### GetMakeSkewCorrectOk

`func (o *OCRSettingsDjVu2PDF) GetMakeSkewCorrectOk() (*bool, bool)`

GetMakeSkewCorrectOk returns a tuple with the MakeSkewCorrect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeSkewCorrect

`func (o *OCRSettingsDjVu2PDF) SetMakeSkewCorrect(v bool)`

SetMakeSkewCorrect sets MakeSkewCorrect field to given value.

### HasMakeSkewCorrect

`func (o *OCRSettingsDjVu2PDF) HasMakeSkewCorrect() bool`

HasMakeSkewCorrect returns a boolean if a field has been set.

### GetMakeSpellCheck

`func (o *OCRSettingsDjVu2PDF) GetMakeSpellCheck() bool`

GetMakeSpellCheck returns the MakeSpellCheck field if non-nil, zero value otherwise.

### GetMakeSpellCheckOk

`func (o *OCRSettingsDjVu2PDF) GetMakeSpellCheckOk() (*bool, bool)`

GetMakeSpellCheckOk returns a tuple with the MakeSpellCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeSpellCheck

`func (o *OCRSettingsDjVu2PDF) SetMakeSpellCheck(v bool)`

SetMakeSpellCheck sets MakeSpellCheck field to given value.

### HasMakeSpellCheck

`func (o *OCRSettingsDjVu2PDF) HasMakeSpellCheck() bool`

HasMakeSpellCheck returns a boolean if a field has been set.

### GetMakeContrastCorrection

`func (o *OCRSettingsDjVu2PDF) GetMakeContrastCorrection() bool`

GetMakeContrastCorrection returns the MakeContrastCorrection field if non-nil, zero value otherwise.

### GetMakeContrastCorrectionOk

`func (o *OCRSettingsDjVu2PDF) GetMakeContrastCorrectionOk() (*bool, bool)`

GetMakeContrastCorrectionOk returns a tuple with the MakeContrastCorrection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeContrastCorrection

`func (o *OCRSettingsDjVu2PDF) SetMakeContrastCorrection(v bool)`

SetMakeContrastCorrection sets MakeContrastCorrection field to given value.

### HasMakeContrastCorrection

`func (o *OCRSettingsDjVu2PDF) HasMakeContrastCorrection() bool`

HasMakeContrastCorrection returns a boolean if a field has been set.

### GetMakeBinarization

`func (o *OCRSettingsDjVu2PDF) GetMakeBinarization() bool`

GetMakeBinarization returns the MakeBinarization field if non-nil, zero value otherwise.

### GetMakeBinarizationOk

`func (o *OCRSettingsDjVu2PDF) GetMakeBinarizationOk() (*bool, bool)`

GetMakeBinarizationOk returns a tuple with the MakeBinarization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeBinarization

`func (o *OCRSettingsDjVu2PDF) SetMakeBinarization(v bool)`

SetMakeBinarization sets MakeBinarization field to given value.

### HasMakeBinarization

`func (o *OCRSettingsDjVu2PDF) HasMakeBinarization() bool`

HasMakeBinarization returns a boolean if a field has been set.

### GetMakeUpsampling

`func (o *OCRSettingsDjVu2PDF) GetMakeUpsampling() bool`

GetMakeUpsampling returns the MakeUpsampling field if non-nil, zero value otherwise.

### GetMakeUpsamplingOk

`func (o *OCRSettingsDjVu2PDF) GetMakeUpsamplingOk() (*bool, bool)`

GetMakeUpsamplingOk returns a tuple with the MakeUpsampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeUpsampling

`func (o *OCRSettingsDjVu2PDF) SetMakeUpsampling(v bool)`

SetMakeUpsampling sets MakeUpsampling field to given value.

### HasMakeUpsampling

`func (o *OCRSettingsDjVu2PDF) HasMakeUpsampling() bool`

HasMakeUpsampling returns a boolean if a field has been set.

### GetDsrMode

`func (o *OCRSettingsDjVu2PDF) GetDsrMode() DsrMode`

GetDsrMode returns the DsrMode field if non-nil, zero value otherwise.

### GetDsrModeOk

`func (o *OCRSettingsDjVu2PDF) GetDsrModeOk() (*DsrMode, bool)`

GetDsrModeOk returns a tuple with the DsrMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsrMode

`func (o *OCRSettingsDjVu2PDF) SetDsrMode(v DsrMode)`

SetDsrMode sets DsrMode field to given value.

### HasDsrMode

`func (o *OCRSettingsDjVu2PDF) HasDsrMode() bool`

HasDsrMode returns a boolean if a field has been set.

### GetDsrConfidence

`func (o *OCRSettingsDjVu2PDF) GetDsrConfidence() DsrConfidence`

GetDsrConfidence returns the DsrConfidence field if non-nil, zero value otherwise.

### GetDsrConfidenceOk

`func (o *OCRSettingsDjVu2PDF) GetDsrConfidenceOk() (*DsrConfidence, bool)`

GetDsrConfidenceOk returns a tuple with the DsrConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsrConfidence

`func (o *OCRSettingsDjVu2PDF) SetDsrConfidence(v DsrConfidence)`

SetDsrConfidence sets DsrConfidence field to given value.

### HasDsrConfidence

`func (o *OCRSettingsDjVu2PDF) HasDsrConfidence() bool`

HasDsrConfidence returns a boolean if a field has been set.

### GetResultType

`func (o *OCRSettingsDjVu2PDF) GetResultType() ResultType`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *OCRSettingsDjVu2PDF) GetResultTypeOk() (*ResultType, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *OCRSettingsDjVu2PDF) SetResultType(v ResultType)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *OCRSettingsDjVu2PDF) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetResultTypeTable

`func (o *OCRSettingsDjVu2PDF) GetResultTypeTable() ResultTypeTable`

GetResultTypeTable returns the ResultTypeTable field if non-nil, zero value otherwise.

### GetResultTypeTableOk

`func (o *OCRSettingsDjVu2PDF) GetResultTypeTableOk() (*ResultTypeTable, bool)`

GetResultTypeTableOk returns a tuple with the ResultTypeTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTypeTable

`func (o *OCRSettingsDjVu2PDF) SetResultTypeTable(v ResultTypeTable)`

SetResultTypeTable sets ResultTypeTable field to given value.

### HasResultTypeTable

`func (o *OCRSettingsDjVu2PDF) HasResultTypeTable() bool`

HasResultTypeTable returns a boolean if a field has been set.

### GetRegions

`func (o *OCRSettingsDjVu2PDF) GetRegions() []OCRRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *OCRSettingsDjVu2PDF) GetRegionsOk() (*[]OCRRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *OCRSettingsDjVu2PDF) SetRegions(v []OCRRegion)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *OCRSettingsDjVu2PDF) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### SetRegionsNil

`func (o *OCRSettingsDjVu2PDF) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *OCRSettingsDjVu2PDF) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


