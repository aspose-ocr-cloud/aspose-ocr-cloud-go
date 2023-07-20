/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRSettingsRecognizeFont type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRSettingsRecognizeFont{}

// OCRSettingsRecognizeFont struct for OCRSettingsRecognizeFont
type OCRSettingsRecognizeFont struct {
	Language *Language `json:"language,omitempty"`
	MakeSkewCorrect *bool `json:"makeSkewCorrect,omitempty"`
	MakeSpellCheck *bool `json:"makeSpellCheck,omitempty"`
	MakeContrastCorrection *bool `json:"makeContrastCorrection,omitempty"`
	ResultType *ResultType `json:"resultType,omitempty"`
	Rotate *int32 `json:"Rotate,omitempty"`
	MakeBinarization *bool `json:"makeBinarization,omitempty"`
	// Option to enable image up-sampling algorithm to improve quality. True by default
	MakeUpsampling *bool `json:"makeUpsampling,omitempty"`
	DsrMode *DsrMode `json:"dsrMode,omitempty"`
	DsrConfidence *DsrConfidence `json:"dsrConfidence,omitempty"`
	ResultTypeTable *ResultTypeTable `json:"resultTypeTable,omitempty"`
	Regions []OCRRegion `json:"regions,omitempty"`
}

// NewOCRSettingsRecognizeFont instantiates a new OCRSettingsRecognizeFont object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRSettingsRecognizeFont() *OCRSettingsRecognizeFont {
	this := OCRSettingsRecognizeFont{}
	var makeSkewCorrect bool = true
	this.MakeSkewCorrect = &makeSkewCorrect
	var makeSpellCheck bool = false
	this.MakeSpellCheck = &makeSpellCheck
	var makeContrastCorrection bool = true
	this.MakeContrastCorrection = &makeContrastCorrection
	var makeBinarization bool = true
	this.MakeBinarization = &makeBinarization
	var makeUpsampling bool = false
	this.MakeUpsampling = &makeUpsampling
	return &this
}

// NewOCRSettingsRecognizeFontWithDefaults instantiates a new OCRSettingsRecognizeFont object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRSettingsRecognizeFontWithDefaults() *OCRSettingsRecognizeFont {
	this := OCRSettingsRecognizeFont{}
	var makeSkewCorrect bool = true
	this.MakeSkewCorrect = &makeSkewCorrect
	var makeSpellCheck bool = false
	this.MakeSpellCheck = &makeSpellCheck
	var makeContrastCorrection bool = true
	this.MakeContrastCorrection = &makeContrastCorrection
	var makeBinarization bool = true
	this.MakeBinarization = &makeBinarization
	var makeUpsampling bool = false
	this.MakeUpsampling = &makeUpsampling
	return &this
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeFont) GetLanguage() Language {
	if o == nil || IsNil(o.Language) {
		var ret Language
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeFont) GetLanguageOk() (*Language, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeFont) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given Language and assigns it to the Language field.
func (o *OCRSettingsRecognizeFont) SetLanguage(v Language) {
	o.Language = &v
}

// GetMakeSkewCorrect returns the MakeSkewCorrect field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeFont) GetMakeSkewCorrect() bool {
	if o == nil || IsNil(o.MakeSkewCorrect) {
		var ret bool
		return ret
	}
	return *o.MakeSkewCorrect
}

// GetMakeSkewCorrectOk returns a tuple with the MakeSkewCorrect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeFont) GetMakeSkewCorrectOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeSkewCorrect) {
		return nil, false
	}
	return o.MakeSkewCorrect, true
}

// HasMakeSkewCorrect returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeFont) HasMakeSkewCorrect() bool {
	if o != nil && !IsNil(o.MakeSkewCorrect) {
		return true
	}

	return false
}

// SetMakeSkewCorrect gets a reference to the given bool and assigns it to the MakeSkewCorrect field.
func (o *OCRSettingsRecognizeFont) SetMakeSkewCorrect(v bool) {
	o.MakeSkewCorrect = &v
}

// GetMakeSpellCheck returns the MakeSpellCheck field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeFont) GetMakeSpellCheck() bool {
	if o == nil || IsNil(o.MakeSpellCheck) {
		var ret bool
		return ret
	}
	return *o.MakeSpellCheck
}

// GetMakeSpellCheckOk returns a tuple with the MakeSpellCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeFont) GetMakeSpellCheckOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeSpellCheck) {
		return nil, false
	}
	return o.MakeSpellCheck, true
}

// HasMakeSpellCheck returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeFont) HasMakeSpellCheck() bool {
	if o != nil && !IsNil(o.MakeSpellCheck) {
		return true
	}

	return false
}

// SetMakeSpellCheck gets a reference to the given bool and assigns it to the MakeSpellCheck field.
func (o *OCRSettingsRecognizeFont) SetMakeSpellCheck(v bool) {
	o.MakeSpellCheck = &v
}

// GetMakeContrastCorrection returns the MakeContrastCorrection field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeFont) GetMakeContrastCorrection() bool {
	if o == nil || IsNil(o.MakeContrastCorrection) {
		var ret bool
		return ret
	}
	return *o.MakeContrastCorrection
}

// GetMakeContrastCorrectionOk returns a tuple with the MakeContrastCorrection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeFont) GetMakeContrastCorrectionOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeContrastCorrection) {
		return nil, false
	}
	return o.MakeContrastCorrection, true
}

// HasMakeContrastCorrection returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeFont) HasMakeContrastCorrection() bool {
	if o != nil && !IsNil(o.MakeContrastCorrection) {
		return true
	}

	return false
}

// SetMakeContrastCorrection gets a reference to the given bool and assigns it to the MakeContrastCorrection field.
func (o *OCRSettingsRecognizeFont) SetMakeContrastCorrection(v bool) {
	o.MakeContrastCorrection = &v
}

// GetResultType returns the ResultType field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeFont) GetResultType() ResultType {
	if o == nil || IsNil(o.ResultType) {
		var ret ResultType
		return ret
	}
	return *o.ResultType
}

// GetResultTypeOk returns a tuple with the ResultType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeFont) GetResultTypeOk() (*ResultType, bool) {
	if o == nil || IsNil(o.ResultType) {
		return nil, false
	}
	return o.ResultType, true
}

// HasResultType returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeFont) HasResultType() bool {
	if o != nil && !IsNil(o.ResultType) {
		return true
	}

	return false
}

// SetResultType gets a reference to the given ResultType and assigns it to the ResultType field.
func (o *OCRSettingsRecognizeFont) SetResultType(v ResultType) {
	o.ResultType = &v
}

// GetRotate returns the Rotate field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeFont) GetRotate() int32 {
	if o == nil || IsNil(o.Rotate) {
		var ret int32
		return ret
	}
	return *o.Rotate
}

// GetRotateOk returns a tuple with the Rotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeFont) GetRotateOk() (*int32, bool) {
	if o == nil || IsNil(o.Rotate) {
		return nil, false
	}
	return o.Rotate, true
}

// HasRotate returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeFont) HasRotate() bool {
	if o != nil && !IsNil(o.Rotate) {
		return true
	}

	return false
}

// SetRotate gets a reference to the given int32 and assigns it to the Rotate field.
func (o *OCRSettingsRecognizeFont) SetRotate(v int32) {
	o.Rotate = &v
}

// GetMakeBinarization returns the MakeBinarization field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeFont) GetMakeBinarization() bool {
	if o == nil || IsNil(o.MakeBinarization) {
		var ret bool
		return ret
	}
	return *o.MakeBinarization
}

// GetMakeBinarizationOk returns a tuple with the MakeBinarization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeFont) GetMakeBinarizationOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeBinarization) {
		return nil, false
	}
	return o.MakeBinarization, true
}

// HasMakeBinarization returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeFont) HasMakeBinarization() bool {
	if o != nil && !IsNil(o.MakeBinarization) {
		return true
	}

	return false
}

// SetMakeBinarization gets a reference to the given bool and assigns it to the MakeBinarization field.
func (o *OCRSettingsRecognizeFont) SetMakeBinarization(v bool) {
	o.MakeBinarization = &v
}

// GetMakeUpsampling returns the MakeUpsampling field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeFont) GetMakeUpsampling() bool {
	if o == nil || IsNil(o.MakeUpsampling) {
		var ret bool
		return ret
	}
	return *o.MakeUpsampling
}

// GetMakeUpsamplingOk returns a tuple with the MakeUpsampling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeFont) GetMakeUpsamplingOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeUpsampling) {
		return nil, false
	}
	return o.MakeUpsampling, true
}

// HasMakeUpsampling returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeFont) HasMakeUpsampling() bool {
	if o != nil && !IsNil(o.MakeUpsampling) {
		return true
	}

	return false
}

// SetMakeUpsampling gets a reference to the given bool and assigns it to the MakeUpsampling field.
func (o *OCRSettingsRecognizeFont) SetMakeUpsampling(v bool) {
	o.MakeUpsampling = &v
}

// GetDsrMode returns the DsrMode field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeFont) GetDsrMode() DsrMode {
	if o == nil || IsNil(o.DsrMode) {
		var ret DsrMode
		return ret
	}
	return *o.DsrMode
}

// GetDsrModeOk returns a tuple with the DsrMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeFont) GetDsrModeOk() (*DsrMode, bool) {
	if o == nil || IsNil(o.DsrMode) {
		return nil, false
	}
	return o.DsrMode, true
}

// HasDsrMode returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeFont) HasDsrMode() bool {
	if o != nil && !IsNil(o.DsrMode) {
		return true
	}

	return false
}

// SetDsrMode gets a reference to the given DsrMode and assigns it to the DsrMode field.
func (o *OCRSettingsRecognizeFont) SetDsrMode(v DsrMode) {
	o.DsrMode = &v
}

// GetDsrConfidence returns the DsrConfidence field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeFont) GetDsrConfidence() DsrConfidence {
	if o == nil || IsNil(o.DsrConfidence) {
		var ret DsrConfidence
		return ret
	}
	return *o.DsrConfidence
}

// GetDsrConfidenceOk returns a tuple with the DsrConfidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeFont) GetDsrConfidenceOk() (*DsrConfidence, bool) {
	if o == nil || IsNil(o.DsrConfidence) {
		return nil, false
	}
	return o.DsrConfidence, true
}

// HasDsrConfidence returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeFont) HasDsrConfidence() bool {
	if o != nil && !IsNil(o.DsrConfidence) {
		return true
	}

	return false
}

// SetDsrConfidence gets a reference to the given DsrConfidence and assigns it to the DsrConfidence field.
func (o *OCRSettingsRecognizeFont) SetDsrConfidence(v DsrConfidence) {
	o.DsrConfidence = &v
}

// GetResultTypeTable returns the ResultTypeTable field value if set, zero value otherwise.
func (o *OCRSettingsRecognizeFont) GetResultTypeTable() ResultTypeTable {
	if o == nil || IsNil(o.ResultTypeTable) {
		var ret ResultTypeTable
		return ret
	}
	return *o.ResultTypeTable
}

// GetResultTypeTableOk returns a tuple with the ResultTypeTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsRecognizeFont) GetResultTypeTableOk() (*ResultTypeTable, bool) {
	if o == nil || IsNil(o.ResultTypeTable) {
		return nil, false
	}
	return o.ResultTypeTable, true
}

// HasResultTypeTable returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeFont) HasResultTypeTable() bool {
	if o != nil && !IsNil(o.ResultTypeTable) {
		return true
	}

	return false
}

// SetResultTypeTable gets a reference to the given ResultTypeTable and assigns it to the ResultTypeTable field.
func (o *OCRSettingsRecognizeFont) SetResultTypeTable(v ResultTypeTable) {
	o.ResultTypeTable = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OCRSettingsRecognizeFont) GetRegions() []OCRRegion {
	if o == nil {
		var ret []OCRRegion
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OCRSettingsRecognizeFont) GetRegionsOk() ([]OCRRegion, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *OCRSettingsRecognizeFont) HasRegions() bool {
	if o != nil && IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []OCRRegion and assigns it to the Regions field.
func (o *OCRSettingsRecognizeFont) SetRegions(v []OCRRegion) {
	o.Regions = v
}

func (o OCRSettingsRecognizeFont) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRSettingsRecognizeFont) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.MakeSkewCorrect) {
		toSerialize["makeSkewCorrect"] = o.MakeSkewCorrect
	}
	if !IsNil(o.MakeSpellCheck) {
		toSerialize["makeSpellCheck"] = o.MakeSpellCheck
	}
	if !IsNil(o.MakeContrastCorrection) {
		toSerialize["makeContrastCorrection"] = o.MakeContrastCorrection
	}
	if !IsNil(o.ResultType) {
		toSerialize["resultType"] = o.ResultType
	}
	if !IsNil(o.Rotate) {
		toSerialize["Rotate"] = o.Rotate
	}
	if !IsNil(o.MakeBinarization) {
		toSerialize["makeBinarization"] = o.MakeBinarization
	}
	if !IsNil(o.MakeUpsampling) {
		toSerialize["makeUpsampling"] = o.MakeUpsampling
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

type NullableOCRSettingsRecognizeFont struct {
	value *OCRSettingsRecognizeFont
	isSet bool
}

func (v NullableOCRSettingsRecognizeFont) Get() *OCRSettingsRecognizeFont {
	return v.value
}

func (v *NullableOCRSettingsRecognizeFont) Set(val *OCRSettingsRecognizeFont) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRSettingsRecognizeFont) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRSettingsRecognizeFont) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRSettingsRecognizeFont(val *OCRSettingsRecognizeFont) *NullableOCRSettingsRecognizeFont {
	return &NullableOCRSettingsRecognizeFont{value: val, isSet: true}
}

func (v NullableOCRSettingsRecognizeFont) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRSettingsRecognizeFont) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


