# OCRSettingsRecognizeAndParseInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MakeSkewCorrect** | Pointer to **bool** |  | [optional] [default to false]
**MakeBinarization** | Pointer to **bool** |  | [optional] [default to false]
**MakeUpsampling** | Pointer to **bool** |  | [optional] [default to false]
**ResultType** | Pointer to [**ResultType**](ResultType.md) |  | [optional] 
**Language** | Pointer to [**Language**](Language.md) |  | [optional] 
**Rotate** | Pointer to **int32** |  | [optional] 
**MakeSpellCheck** | Pointer to **bool** | Option to enable spell checking and correction algorithm. False by default | [optional] [default to false]
**MakeContrastCorrection** | Pointer to **bool** | Option to enable image contrast correction algorithm. True by default | [optional] [default to false]
**DsrMode** | Pointer to [**DsrMode**](DsrMode.md) |  | [optional] 
**DsrConfidence** | Pointer to [**DsrConfidence**](DsrConfidence.md) |  | [optional] 
**ResultTypeTable** | Pointer to [**ResultTypeTable**](ResultTypeTable.md) |  | [optional] 
**Regions** | Pointer to [**[]OCRRegion**](OCRRegion.md) |  | [optional] 

## Methods

### NewOCRSettingsRecognizeAndParseInvoice

`func NewOCRSettingsRecognizeAndParseInvoice() *OCRSettingsRecognizeAndParseInvoice`

NewOCRSettingsRecognizeAndParseInvoice instantiates a new OCRSettingsRecognizeAndParseInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRSettingsRecognizeAndParseInvoiceWithDefaults

`func NewOCRSettingsRecognizeAndParseInvoiceWithDefaults() *OCRSettingsRecognizeAndParseInvoice`

NewOCRSettingsRecognizeAndParseInvoiceWithDefaults instantiates a new OCRSettingsRecognizeAndParseInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMakeSkewCorrect

`func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeSkewCorrect() bool`

GetMakeSkewCorrect returns the MakeSkewCorrect field if non-nil, zero value otherwise.

### GetMakeSkewCorrectOk

`func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeSkewCorrectOk() (*bool, bool)`

GetMakeSkewCorrectOk returns a tuple with the MakeSkewCorrect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeSkewCorrect

`func (o *OCRSettingsRecognizeAndParseInvoice) SetMakeSkewCorrect(v bool)`

SetMakeSkewCorrect sets MakeSkewCorrect field to given value.

### HasMakeSkewCorrect

`func (o *OCRSettingsRecognizeAndParseInvoice) HasMakeSkewCorrect() bool`

HasMakeSkewCorrect returns a boolean if a field has been set.

### GetMakeBinarization

`func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeBinarization() bool`

GetMakeBinarization returns the MakeBinarization field if non-nil, zero value otherwise.

### GetMakeBinarizationOk

`func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeBinarizationOk() (*bool, bool)`

GetMakeBinarizationOk returns a tuple with the MakeBinarization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeBinarization

`func (o *OCRSettingsRecognizeAndParseInvoice) SetMakeBinarization(v bool)`

SetMakeBinarization sets MakeBinarization field to given value.

### HasMakeBinarization

`func (o *OCRSettingsRecognizeAndParseInvoice) HasMakeBinarization() bool`

HasMakeBinarization returns a boolean if a field has been set.

### GetMakeUpsampling

`func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeUpsampling() bool`

GetMakeUpsampling returns the MakeUpsampling field if non-nil, zero value otherwise.

### GetMakeUpsamplingOk

`func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeUpsamplingOk() (*bool, bool)`

GetMakeUpsamplingOk returns a tuple with the MakeUpsampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeUpsampling

`func (o *OCRSettingsRecognizeAndParseInvoice) SetMakeUpsampling(v bool)`

SetMakeUpsampling sets MakeUpsampling field to given value.

### HasMakeUpsampling

`func (o *OCRSettingsRecognizeAndParseInvoice) HasMakeUpsampling() bool`

HasMakeUpsampling returns a boolean if a field has been set.

### GetResultType

`func (o *OCRSettingsRecognizeAndParseInvoice) GetResultType() ResultType`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *OCRSettingsRecognizeAndParseInvoice) GetResultTypeOk() (*ResultType, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *OCRSettingsRecognizeAndParseInvoice) SetResultType(v ResultType)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *OCRSettingsRecognizeAndParseInvoice) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetLanguage

`func (o *OCRSettingsRecognizeAndParseInvoice) GetLanguage() Language`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *OCRSettingsRecognizeAndParseInvoice) GetLanguageOk() (*Language, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *OCRSettingsRecognizeAndParseInvoice) SetLanguage(v Language)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *OCRSettingsRecognizeAndParseInvoice) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRotate

`func (o *OCRSettingsRecognizeAndParseInvoice) GetRotate() int32`

GetRotate returns the Rotate field if non-nil, zero value otherwise.

### GetRotateOk

`func (o *OCRSettingsRecognizeAndParseInvoice) GetRotateOk() (*int32, bool)`

GetRotateOk returns a tuple with the Rotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotate

`func (o *OCRSettingsRecognizeAndParseInvoice) SetRotate(v int32)`

SetRotate sets Rotate field to given value.

### HasRotate

`func (o *OCRSettingsRecognizeAndParseInvoice) HasRotate() bool`

HasRotate returns a boolean if a field has been set.

### GetMakeSpellCheck

`func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeSpellCheck() bool`

GetMakeSpellCheck returns the MakeSpellCheck field if non-nil, zero value otherwise.

### GetMakeSpellCheckOk

`func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeSpellCheckOk() (*bool, bool)`

GetMakeSpellCheckOk returns a tuple with the MakeSpellCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeSpellCheck

`func (o *OCRSettingsRecognizeAndParseInvoice) SetMakeSpellCheck(v bool)`

SetMakeSpellCheck sets MakeSpellCheck field to given value.

### HasMakeSpellCheck

`func (o *OCRSettingsRecognizeAndParseInvoice) HasMakeSpellCheck() bool`

HasMakeSpellCheck returns a boolean if a field has been set.

### GetMakeContrastCorrection

`func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeContrastCorrection() bool`

GetMakeContrastCorrection returns the MakeContrastCorrection field if non-nil, zero value otherwise.

### GetMakeContrastCorrectionOk

`func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeContrastCorrectionOk() (*bool, bool)`

GetMakeContrastCorrectionOk returns a tuple with the MakeContrastCorrection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeContrastCorrection

`func (o *OCRSettingsRecognizeAndParseInvoice) SetMakeContrastCorrection(v bool)`

SetMakeContrastCorrection sets MakeContrastCorrection field to given value.

### HasMakeContrastCorrection

`func (o *OCRSettingsRecognizeAndParseInvoice) HasMakeContrastCorrection() bool`

HasMakeContrastCorrection returns a boolean if a field has been set.

### GetDsrMode

`func (o *OCRSettingsRecognizeAndParseInvoice) GetDsrMode() DsrMode`

GetDsrMode returns the DsrMode field if non-nil, zero value otherwise.

### GetDsrModeOk

`func (o *OCRSettingsRecognizeAndParseInvoice) GetDsrModeOk() (*DsrMode, bool)`

GetDsrModeOk returns a tuple with the DsrMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsrMode

`func (o *OCRSettingsRecognizeAndParseInvoice) SetDsrMode(v DsrMode)`

SetDsrMode sets DsrMode field to given value.

### HasDsrMode

`func (o *OCRSettingsRecognizeAndParseInvoice) HasDsrMode() bool`

HasDsrMode returns a boolean if a field has been set.

### GetDsrConfidence

`func (o *OCRSettingsRecognizeAndParseInvoice) GetDsrConfidence() DsrConfidence`

GetDsrConfidence returns the DsrConfidence field if non-nil, zero value otherwise.

### GetDsrConfidenceOk

`func (o *OCRSettingsRecognizeAndParseInvoice) GetDsrConfidenceOk() (*DsrConfidence, bool)`

GetDsrConfidenceOk returns a tuple with the DsrConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsrConfidence

`func (o *OCRSettingsRecognizeAndParseInvoice) SetDsrConfidence(v DsrConfidence)`

SetDsrConfidence sets DsrConfidence field to given value.

### HasDsrConfidence

`func (o *OCRSettingsRecognizeAndParseInvoice) HasDsrConfidence() bool`

HasDsrConfidence returns a boolean if a field has been set.

### GetResultTypeTable

`func (o *OCRSettingsRecognizeAndParseInvoice) GetResultTypeTable() ResultTypeTable`

GetResultTypeTable returns the ResultTypeTable field if non-nil, zero value otherwise.

### GetResultTypeTableOk

`func (o *OCRSettingsRecognizeAndParseInvoice) GetResultTypeTableOk() (*ResultTypeTable, bool)`

GetResultTypeTableOk returns a tuple with the ResultTypeTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTypeTable

`func (o *OCRSettingsRecognizeAndParseInvoice) SetResultTypeTable(v ResultTypeTable)`

SetResultTypeTable sets ResultTypeTable field to given value.

### HasResultTypeTable

`func (o *OCRSettingsRecognizeAndParseInvoice) HasResultTypeTable() bool`

HasResultTypeTable returns a boolean if a field has been set.

### GetRegions

`func (o *OCRSettingsRecognizeAndParseInvoice) GetRegions() []OCRRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *OCRSettingsRecognizeAndParseInvoice) GetRegionsOk() (*[]OCRRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *OCRSettingsRecognizeAndParseInvoice) SetRegions(v []OCRRegion)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *OCRSettingsRecognizeAndParseInvoice) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### SetRegionsNil

`func (o *OCRSettingsRecognizeAndParseInvoice) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *OCRSettingsRecognizeAndParseInvoice) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


