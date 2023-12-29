/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRSettingsRecognizeAndParseInvoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRSettingsRecognizeAndParseInvoice{}

// OCRSettingsRecognizeAndParseInvoice struct for OCRSettingsRecognizeAndParseInvoice
type OCRSettingsRecognizeAndParseInvoice struct {
	MakeSkewCorrect *bool `json:"makeSkewCorrect,omitempty"`
	MakeBinarization *bool `json:"makeBinarization,omitempty"`
	MakeUpsampling *bool `json:"makeUpsampling,omitempty"`
	ResultType *ResultType `json:"resultType,omitempty"`
	Language *Language `json:"language,omitempty"`
	Rotate *int32 `json:"Rotate,omitempty"`
	// Option to enable spell checking and correction algorithm. False by default
	MakeSpellCheck *bool `json:"makeSpellCheck,omitempty"`
	// Option to enable image contrast correction algorithm. True by default
	// Deprecated
	MakeContrastCorrection *bool `json:"makeContrastCorrection,omitempty"`
	DsrMode *DsrMode `json:"dsrMode,omitempty"`
	DsrConfidence *DsrConfidence `json:"dsrConfidence,omitempty"`
	ResultTypeTable *ResultTypeTable `json:"resultTypeTable,omitempty"`
	Regions []OCRRegion `json:"regions,omitempty"`
}

// NewOCRSettingsRecognizeAndParseInvoice instantiates a new OCRSettingsRecognizeAndParseInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRSettingsRecognizeAndParseInvoice() *OCRSettingsRecognizeAndParseInvoice {
	this := OCRSettingsRecognizeAndParseInvoice{}
	var makeSkewCorrect bool = false
	this.MakeSkewCorrect = &makeSkewCorrect
	var makeBinarization bool = false
	this.MakeBinarization = &makeBinarization
	var makeUpsampling bool = false
	this.MakeUpsampling = &makeUpsampling
	var makeSpellCheck bool = false
	this.MakeSpellCheck = &makeSpellCheck
	var makeContrastCorrection bool = false
	this.MakeContrastCorrection = &makeContrastCorrection
	return &this
}

// NewOCRSettingsRecognizeAndParseInvoiceWithDefaults instantiates a new OCRSettingsRecognizeAndParseInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRSettingsRecognizeAndParseInvoiceWithDefaults() *OCRSettingsRecognizeAndParseInvoice {
	this := OCRSettingsRecognizeAndParseInvoice{}
	var makeSkewCorrect bool = false
	this.MakeSkewCorrect = &makeSkewCorrect
	var makeBinarization bool = false
	this.MakeBinarization = &makeBinarization
	var makeUpsampling bool = false
	this.MakeUpsampling = &makeUpsampling
	var makeSpellCheck bool = false
	this.MakeSpellCheck = &makeSpellCheck
	var makeContrastCorrection bool = false
	this.MakeContrastCorrection = &makeContrastCorrection
	return &this
}

// GetMakeSkewCorrect returns the MakeSkewCorrect field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeSkewCorrect() bool {
	if o == nil || IsNil(o.MakeSkewCorrect) {
		var ret bool
		return ret
	}
	return *o.MakeSkewCorrect
}

// GetMakeSkewCorrectOk returns a tuple with the MakeSkewCorrect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeSkewCorrectOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeSkewCorrect) {
		return nil, false
	}
	return o.MakeSkewCorrect, true
}

// HasMakeSkewCorrect returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) HasMakeSkewCorrect() bool {
	if o != nil && !IsNil(o.MakeSkewCorrect) {
		return true
	}

	return false
}

// SetMakeSkewCorrect gets a reference to the given bool and assigns it to the MakeSkewCorrect field.
func (o *OCRSettingsRecognizeAndParseInvoice) SetMakeSkewCorrect(v bool) {
	o.MakeSkewCorrect = &v
}

// GetMakeBinarization returns the MakeBinarization field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeBinarization() bool {
	if o == nil || IsNil(o.MakeBinarization) {
		var ret bool
		return ret
	}
	return *o.MakeBinarization
}

// GetMakeBinarizationOk returns a tuple with the MakeBinarization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeBinarizationOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeBinarization) {
		return nil, false
	}
	return o.MakeBinarization, true
}

// HasMakeBinarization returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) HasMakeBinarization() bool {
	if o != nil && !IsNil(o.MakeBinarization) {
		return true
	}

	return false
}

// SetMakeBinarization gets a reference to the given bool and assigns it to the MakeBinarization field.
func (o *OCRSettingsRecognizeAndParseInvoice) SetMakeBinarization(v bool) {
	o.MakeBinarization = &v
}

// GetMakeUpsampling returns the MakeUpsampling field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeUpsampling() bool {
	if o == nil || IsNil(o.MakeUpsampling) {
		var ret bool
		return ret
	}
	return *o.MakeUpsampling
}

// GetMakeUpsamplingOk returns a tuple with the MakeUpsampling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeUpsamplingOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeUpsampling) {
		return nil, false
	}
	return o.MakeUpsampling, true
}

// HasMakeUpsampling returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) HasMakeUpsampling() bool {
	if o != nil && !IsNil(o.MakeUpsampling) {
		return true
	}

	return false
}

// SetMakeUpsampling gets a reference to the given bool and assigns it to the MakeUpsampling field.
func (o *OCRSettingsRecognizeAndParseInvoice) SetMakeUpsampling(v bool) {
	o.MakeUpsampling = &v
}

// GetResultType returns the ResultType field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeAndParseInvoice) GetResultType() ResultType {
	if o == nil || IsNil(o.ResultType) {
		var ret ResultType
		return ret
	}
	return *o.ResultType
}

// GetResultTypeOk returns a tuple with the ResultType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) GetResultTypeOk() (*ResultType, bool) {
	if o == nil || IsNil(o.ResultType) {
		return nil, false
	}
	return o.ResultType, true
}

// HasResultType returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) HasResultType() bool {
	if o != nil && !IsNil(o.ResultType) {
		return true
	}

	return false
}

// SetResultType gets a reference to the given ResultType and assigns it to the ResultType field.
func (o *OCRSettingsRecognizeAndParseInvoice) SetResultType(v ResultType) {
	o.ResultType = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeAndParseInvoice) GetLanguage() Language {
	if o == nil || IsNil(o.Language) {
		var ret Language
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) GetLanguageOk() (*Language, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given Language and assigns it to the Language field.
func (o *OCRSettingsRecognizeAndParseInvoice) SetLanguage(v Language) {
	o.Language = &v
}

// GetRotate returns the Rotate field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeAndParseInvoice) GetRotate() int32 {
	if o == nil || IsNil(o.Rotate) {
		var ret int32
		return ret
	}
	return *o.Rotate
}

// GetRotateOk returns a tuple with the Rotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) GetRotateOk() (*int32, bool) {
	if o == nil || IsNil(o.Rotate) {
		return nil, false
	}
	return o.Rotate, true
}

// HasRotate returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) HasRotate() bool {
	if o != nil && !IsNil(o.Rotate) {
		return true
	}

	return false
}

// SetRotate gets a reference to the given int32 and assigns it to the Rotate field.
func (o *OCRSettingsRecognizeAndParseInvoice) SetRotate(v int32) {
	o.Rotate = &v
}

// GetMakeSpellCheck returns the MakeSpellCheck field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeSpellCheck() bool {
	if o == nil || IsNil(o.MakeSpellCheck) {
		var ret bool
		return ret
	}
	return *o.MakeSpellCheck
}

// GetMakeSpellCheckOk returns a tuple with the MakeSpellCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeSpellCheckOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeSpellCheck) {
		return nil, false
	}
	return o.MakeSpellCheck, true
}

// HasMakeSpellCheck returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) HasMakeSpellCheck() bool {
	if o != nil && !IsNil(o.MakeSpellCheck) {
		return true
	}

	return false
}

// SetMakeSpellCheck gets a reference to the given bool and assigns it to the MakeSpellCheck field.
func (o *OCRSettingsRecognizeAndParseInvoice) SetMakeSpellCheck(v bool) {
	o.MakeSpellCheck = &v
}

// GetMakeContrastCorrection returns the MakeContrastCorrection field value if set, zero value otherwise.
// Deprecated
func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeContrastCorrection() bool {
	if o == nil || IsNil(o.MakeContrastCorrection) {
		var ret bool
		return ret
	}
	return *o.MakeContrastCorrection
}

// GetMakeContrastCorrectionOk returns a tuple with the MakeContrastCorrection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *OCRSettingsRecognizeAndParseInvoice) GetMakeContrastCorrectionOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeContrastCorrection) {
		return nil, false
	}
	return o.MakeContrastCorrection, true
}

// HasMakeContrastCorrection returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) HasMakeContrastCorrection() bool {
	if o != nil && !IsNil(o.MakeContrastCorrection) {
		return true
	}

	return false
}

// SetMakeContrastCorrection gets a reference to the given bool and assigns it to the MakeContrastCorrection field.
// Deprecated
func (o *OCRSettingsRecognizeAndParseInvoice) SetMakeContrastCorrection(v bool) {
	o.MakeContrastCorrection = &v
}

// GetDsrMode returns the DsrMode field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeAndParseInvoice) GetDsrMode() DsrMode {
	if o == nil || IsNil(o.DsrMode) {
		var ret DsrMode
		return ret
	}
	return *o.DsrMode
}

// GetDsrModeOk returns a tuple with the DsrMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) GetDsrModeOk() (*DsrMode, bool) {
	if o == nil || IsNil(o.DsrMode) {
		return nil, false
	}
	return o.DsrMode, true
}

// HasDsrMode returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) HasDsrMode() bool {
	if o != nil && !IsNil(o.DsrMode) {
		return true
	}

	return false
}

// SetDsrMode gets a reference to the given DsrMode and assigns it to the DsrMode field.
func (o *OCRSettingsRecognizeAndParseInvoice) SetDsrMode(v DsrMode) {
	o.DsrMode = &v
}

// GetDsrConfidence returns the DsrConfidence field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeAndParseInvoice) GetDsrConfidence() DsrConfidence {
	if o == nil || IsNil(o.DsrConfidence) {
		var ret DsrConfidence
		return ret
	}
	return *o.DsrConfidence
}

// GetDsrConfidenceOk returns a tuple with the DsrConfidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) GetDsrConfidenceOk() (*DsrConfidence, bool) {
	if o == nil || IsNil(o.DsrConfidence) {
		return nil, false
	}
	return o.DsrConfidence, true
}

// HasDsrConfidence returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) HasDsrConfidence() bool {
	if o != nil && !IsNil(o.DsrConfidence) {
		return true
	}

	return false
}

// SetDsrConfidence gets a reference to the given DsrConfidence and assigns it to the DsrConfidence field.
func (o *OCRSettingsRecognizeAndParseInvoice) SetDsrConfidence(v DsrConfidence) {
	o.DsrConfidence = &v
}

// GetResultTypeTable returns the ResultTypeTable field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeAndParseInvoice) GetResultTypeTable() ResultTypeTable {
	if o == nil || IsNil(o.ResultTypeTable) {
		var ret ResultTypeTable
		return ret
	}
	return *o.ResultTypeTable
}

// GetResultTypeTableOk returns a tuple with the ResultTypeTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) GetResultTypeTableOk() (*ResultTypeTable, bool) {
	if o == nil || IsNil(o.ResultTypeTable) {
		return nil, false
	}
	return o.ResultTypeTable, true
}

// HasResultTypeTable returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) HasResultTypeTable() bool {
	if o != nil && !IsNil(o.ResultTypeTable) {
		return true
	}

	return false
}

// SetResultTypeTable gets a reference to the given ResultTypeTable and assigns it to the ResultTypeTable field.
func (o *OCRSettingsRecognizeAndParseInvoice) SetResultTypeTable(v ResultTypeTable) {
	o.ResultTypeTable = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OCRSettingsRecognizeAndParseInvoice) GetRegions() []OCRRegion {
	if o == nil {
		var ret []OCRRegion
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OCRSettingsRecognizeAndParseInvoice) GetRegionsOk() ([]OCRRegion, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeAndParseInvoice) HasRegions() bool {
	if o != nil && IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []OCRRegion and assigns it to the Regions field.
func (o *OCRSettingsRecognizeAndParseInvoice) SetRegions(v []OCRRegion) {
	o.Regions = v
}

func (o OCRSettingsRecognizeAndParseInvoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRSettingsRecognizeAndParseInvoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MakeSkewCorrect) {
		toSerialize["makeSkewCorrect"] = o.MakeSkewCorrect
	}
	if !IsNil(o.MakeBinarization) {
		toSerialize["makeBinarization"] = o.MakeBinarization
	}
	if !IsNil(o.MakeUpsampling) {
		toSerialize["makeUpsampling"] = o.MakeUpsampling
	}
	if !IsNil(o.ResultType) {
		toSerialize["resultType"] = o.ResultType
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Rotate) {
		toSerialize["Rotate"] = o.Rotate
	}
	if !IsNil(o.MakeSpellCheck) {
		toSerialize["makeSpellCheck"] = o.MakeSpellCheck
	}
	if !IsNil(o.MakeContrastCorrection) {
		toSerialize["makeContrastCorrection"] = o.MakeContrastCorrection
	}
	if !IsNil(o.DsrMode) {
		toSerialize["dsrMode"] = o.DsrMode
	}
	if !IsNil(o.DsrConfidence) {
		toSerialize["dsrConfidence"] = o.DsrConfidence
	}
	if !IsNil(o.ResultTypeTable) {
		toSerialize["resultTypeTable"] = o.ResultTypeTable
	}
	if o.Regions != nil {
		toSerialize["regions"] = o.Regions
	}
	return toSerialize, nil
}

type NullableOCRSettingsRecognizeAndParseInvoice struct {
	value *OCRSettingsRecognizeAndParseInvoice
	isSet bool
}

func (v NullableOCRSettingsRecognizeAndParseInvoice) Get() *OCRSettingsRecognizeAndParseInvoice {
	return v.value
}

func (v *NullableOCRSettingsRecognizeAndParseInvoice) Set(val *OCRSettingsRecognizeAndParseInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRSettingsRecognizeAndParseInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRSettingsRecognizeAndParseInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRSettingsRecognizeAndParseInvoice(val *OCRSettingsRecognizeAndParseInvoice) *NullableOCRSettingsRecognizeAndParseInvoice {
	return &NullableOCRSettingsRecognizeAndParseInvoice{value: val, isSet: true}
}

func (v NullableOCRSettingsRecognizeAndParseInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRSettingsRecognizeAndParseInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


