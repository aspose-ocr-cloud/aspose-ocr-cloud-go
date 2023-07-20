/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRSettingsDjVu2PDF type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRSettingsDjVu2PDF{}

// OCRSettingsDjVu2PDF struct for OCRSettingsDjVu2PDF
type OCRSettingsDjVu2PDF struct {
	Language *Language `json:"language,omitempty"`
	Rotate *int32 `json:"Rotate,omitempty"`
	// Option to enable skew correction algorithm. True by default
	MakeSkewCorrect *bool `json:"makeSkewCorrect,omitempty"`
	// Option to enable spell checking and correction algorithm. False by default
	MakeSpellCheck *bool `json:"makeSpellCheck,omitempty"`
	// Option to enable image contrast correction algorithm. True by default
	MakeContrastCorrection *bool `json:"makeContrastCorrection,omitempty"`
	MakeBinarization *bool `json:"makeBinarization,omitempty"`
	// Option to enable image up-sampling algorithm to improve quality. True by default
	MakeUpsampling *bool `json:"makeUpsampling,omitempty"`
	DsrMode *DsrMode `json:"dsrMode,omitempty"`
	DsrConfidence *DsrConfidence `json:"dsrConfidence,omitempty"`
	ResultType *ResultType `json:"resultType,omitempty"`
	ResultTypeTable *ResultTypeTable `json:"resultTypeTable,omitempty"`
	Regions []OCRRegion `json:"regions,omitempty"`
}

// NewOCRSettingsDjVu2PDF instantiates a new OCRSettingsDjVu2PDF object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRSettingsDjVu2PDF() *OCRSettingsDjVu2PDF {
	this := OCRSettingsDjVu2PDF{}
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

// NewOCRSettingsDjVu2PDFWithDefaults instantiates a new OCRSettingsDjVu2PDF object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRSettingsDjVu2PDFWithDefaults() *OCRSettingsDjVu2PDF {
	this := OCRSettingsDjVu2PDF{}
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
func (o *OCRSettingsDjVu2PDF) GetLanguage() Language {
	if o == nil || IsNil(o.Language) {
		var ret Language
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsDjVu2PDF) GetLanguageOk() (*Language, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *OCRSettingsDjVu2PDF) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given Language and assigns it to the Language field.
func (o *OCRSettingsDjVu2PDF) SetLanguage(v Language) {
	o.Language = &v
}

// GetRotate returns the Rotate field value if set, zero value otherwise.
func (o *OCRSettingsDjVu2PDF) GetRotate() int32 {
	if o == nil || IsNil(o.Rotate) {
		var ret int32
		return ret
	}
	return *o.Rotate
}

// GetRotateOk returns a tuple with the Rotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsDjVu2PDF) GetRotateOk() (*int32, bool) {
	if o == nil || IsNil(o.Rotate) {
		return nil, false
	}
	return o.Rotate, true
}

// HasRotate returns a boolean if a field has been set.
func (o *OCRSettingsDjVu2PDF) HasRotate() bool {
	if o != nil && !IsNil(o.Rotate) {
		return true
	}

	return false
}

// SetRotate gets a reference to the given int32 and assigns it to the Rotate field.
func (o *OCRSettingsDjVu2PDF) SetRotate(v int32) {
	o.Rotate = &v
}

// GetMakeSkewCorrect returns the MakeSkewCorrect field value if set, zero value otherwise.
func (o *OCRSettingsDjVu2PDF) GetMakeSkewCorrect() bool {
	if o == nil || IsNil(o.MakeSkewCorrect) {
		var ret bool
		return ret
	}
	return *o.MakeSkewCorrect
}

// GetMakeSkewCorrectOk returns a tuple with the MakeSkewCorrect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsDjVu2PDF) GetMakeSkewCorrectOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeSkewCorrect) {
		return nil, false
	}
	return o.MakeSkewCorrect, true
}

// HasMakeSkewCorrect returns a boolean if a field has been set.
func (o *OCRSettingsDjVu2PDF) HasMakeSkewCorrect() bool {
	if o != nil && !IsNil(o.MakeSkewCorrect) {
		return true
	}

	return false
}

// SetMakeSkewCorrect gets a reference to the given bool and assigns it to the MakeSkewCorrect field.
func (o *OCRSettingsDjVu2PDF) SetMakeSkewCorrect(v bool) {
	o.MakeSkewCorrect = &v
}

// GetMakeSpellCheck returns the MakeSpellCheck field value if set, zero value otherwise.
func (o *OCRSettingsDjVu2PDF) GetMakeSpellCheck() bool {
	if o == nil || IsNil(o.MakeSpellCheck) {
		var ret bool
		return ret
	}
	return *o.MakeSpellCheck
}

// GetMakeSpellCheckOk returns a tuple with the MakeSpellCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsDjVu2PDF) GetMakeSpellCheckOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeSpellCheck) {
		return nil, false
	}
	return o.MakeSpellCheck, true
}

// HasMakeSpellCheck returns a boolean if a field has been set.
func (o *OCRSettingsDjVu2PDF) HasMakeSpellCheck() bool {
	if o != nil && !IsNil(o.MakeSpellCheck) {
		return true
	}

	return false
}

// SetMakeSpellCheck gets a reference to the given bool and assigns it to the MakeSpellCheck field.
func (o *OCRSettingsDjVu2PDF) SetMakeSpellCheck(v bool) {
	o.MakeSpellCheck = &v
}

// GetMakeContrastCorrection returns the MakeContrastCorrection field value if set, zero value otherwise.
func (o *OCRSettingsDjVu2PDF) GetMakeContrastCorrection() bool {
	if o == nil || IsNil(o.MakeContrastCorrection) {
		var ret bool
		return ret
	}
	return *o.MakeContrastCorrection
}

// GetMakeContrastCorrectionOk returns a tuple with the MakeContrastCorrection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsDjVu2PDF) GetMakeContrastCorrectionOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeContrastCorrection) {
		return nil, false
	}
	return o.MakeContrastCorrection, true
}

// HasMakeContrastCorrection returns a boolean if a field has been set.
func (o *OCRSettingsDjVu2PDF) HasMakeContrastCorrection() bool {
	if o != nil && !IsNil(o.MakeContrastCorrection) {
		return true
	}

	return false
}

// SetMakeContrastCorrection gets a reference to the given bool and assigns it to the MakeContrastCorrection field.
func (o *OCRSettingsDjVu2PDF) SetMakeContrastCorrection(v bool) {
	o.MakeContrastCorrection = &v
}

// GetMakeBinarization returns the MakeBinarization field value if set, zero value otherwise.
func (o *OCRSettingsDjVu2PDF) GetMakeBinarization() bool {
	if o == nil || IsNil(o.MakeBinarization) {
		var ret bool
		return ret
	}
	return *o.MakeBinarization
}

// GetMakeBinarizationOk returns a tuple with the MakeBinarization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsDjVu2PDF) GetMakeBinarizationOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeBinarization) {
		return nil, false
	}
	return o.MakeBinarization, true
}

// HasMakeBinarization returns a boolean if a field has been set.
func (o *OCRSettingsDjVu2PDF) HasMakeBinarization() bool {
	if o != nil && !IsNil(o.MakeBinarization) {
		return true
	}

	return false
}

// SetMakeBinarization gets a reference to the given bool and assigns it to the MakeBinarization field.
func (o *OCRSettingsDjVu2PDF) SetMakeBinarization(v bool) {
	o.MakeBinarization = &v
}

// GetMakeUpsampling returns the MakeUpsampling field value if set, zero value otherwise.
func (o *OCRSettingsDjVu2PDF) GetMakeUpsampling() bool {
	if o == nil || IsNil(o.MakeUpsampling) {
		var ret bool
		return ret
	}
	return *o.MakeUpsampling
}

// GetMakeUpsamplingOk returns a tuple with the MakeUpsampling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsDjVu2PDF) GetMakeUpsamplingOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeUpsampling) {
		return nil, false
	}
	return o.MakeUpsampling, true
}

// HasMakeUpsampling returns a boolean if a field has been set.
func (o *OCRSettingsDjVu2PDF) HasMakeUpsampling() bool {
	if o != nil && !IsNil(o.MakeUpsampling) {
		return true
	}

	return false
}

// SetMakeUpsampling gets a reference to the given bool and assigns it to the MakeUpsampling field.
func (o *OCRSettingsDjVu2PDF) SetMakeUpsampling(v bool) {
	o.MakeUpsampling = &v
}

// GetDsrMode returns the DsrMode field value if set, zero value otherwise.
func (o *OCRSettingsDjVu2PDF) GetDsrMode() DsrMode {
	if o == nil || IsNil(o.DsrMode) {
		var ret DsrMode
		return ret
	}
	return *o.DsrMode
}

// GetDsrModeOk returns a tuple with the DsrMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsDjVu2PDF) GetDsrModeOk() (*DsrMode, bool) {
	if o == nil || IsNil(o.DsrMode) {
		return nil, false
	}
	return o.DsrMode, true
}

// HasDsrMode returns a boolean if a field has been set.
func (o *OCRSettingsDjVu2PDF) HasDsrMode() bool {
	if o != nil && !IsNil(o.DsrMode) {
		return true
	}

	return false
}

// SetDsrMode gets a reference to the given DsrMode and assigns it to the DsrMode field.
func (o *OCRSettingsDjVu2PDF) SetDsrMode(v DsrMode) {
	o.DsrMode = &v
}

// GetDsrConfidence returns the DsrConfidence field value if set, zero value otherwise.
func (o *OCRSettingsDjVu2PDF) GetDsrConfidence() DsrConfidence {
	if o == nil || IsNil(o.DsrConfidence) {
		var ret DsrConfidence
		return ret
	}
	return *o.DsrConfidence
}

// GetDsrConfidenceOk returns a tuple with the DsrConfidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsDjVu2PDF) GetDsrConfidenceOk() (*DsrConfidence, bool) {
	if o == nil || IsNil(o.DsrConfidence) {
		return nil, false
	}
	return o.DsrConfidence, true
}

// HasDsrConfidence returns a boolean if a field has been set.
func (o *OCRSettingsDjVu2PDF) HasDsrConfidence() bool {
	if o != nil && !IsNil(o.DsrConfidence) {
		return true
	}

	return false
}

// SetDsrConfidence gets a reference to the given DsrConfidence and assigns it to the DsrConfidence field.
func (o *OCRSettingsDjVu2PDF) SetDsrConfidence(v DsrConfidence) {
	o.DsrConfidence = &v
}

// GetResultType returns the ResultType field value if set, zero value otherwise.
func (o *OCRSettingsDjVu2PDF) GetResultType() ResultType {
	if o == nil || IsNil(o.ResultType) {
		var ret ResultType
		return ret
	}
	return *o.ResultType
}

// GetResultTypeOk returns a tuple with the ResultType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsDjVu2PDF) GetResultTypeOk() (*ResultType, bool) {
	if o == nil || IsNil(o.ResultType) {
		return nil, false
	}
	return o.ResultType, true
}

// HasResultType returns a boolean if a field has been set.
func (o *OCRSettingsDjVu2PDF) HasResultType() bool {
	if o != nil && !IsNil(o.ResultType) {
		return true
	}

	return false
}

// SetResultType gets a reference to the given ResultType and assigns it to the ResultType field.
func (o *OCRSettingsDjVu2PDF) SetResultType(v ResultType) {
	o.ResultType = &v
}

// GetResultTypeTable returns the ResultTypeTable field value if set, zero value otherwise.
func (o *OCRSettingsDjVu2PDF) GetResultTypeTable() ResultTypeTable {
	if o == nil || IsNil(o.ResultTypeTable) {
		var ret ResultTypeTable
		return ret
	}
	return *o.ResultTypeTable
}

// GetResultTypeTableOk returns a tuple with the ResultTypeTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRSettingsDjVu2PDF) GetResultTypeTableOk() (*ResultTypeTable, bool) {
	if o == nil || IsNil(o.ResultTypeTable) {
		return nil, false
	}
	return o.ResultTypeTable, true
}

// HasResultTypeTable returns a boolean if a field has been set.
func (o *OCRSettingsDjVu2PDF) HasResultTypeTable() bool {
	if o != nil && !IsNil(o.ResultTypeTable) {
		return true
	}

	return false
}

// SetResultTypeTable gets a reference to the given ResultTypeTable and assigns it to the ResultTypeTable field.
func (o *OCRSettingsDjVu2PDF) SetResultTypeTable(v ResultTypeTable) {
	o.ResultTypeTable = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OCRSettingsDjVu2PDF) GetRegions() []OCRRegion {
	if o == nil {
		var ret []OCRRegion
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OCRSettingsDjVu2PDF) GetRegionsOk() ([]OCRRegion, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *OCRSettingsDjVu2PDF) HasRegions() bool {
	if o != nil && IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []OCRRegion and assigns it to the Regions field.
func (o *OCRSettingsDjVu2PDF) SetRegions(v []OCRRegion) {
	o.Regions = v
}

func (o OCRSettingsDjVu2PDF) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRSettingsDjVu2PDF) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Rotate) {
		toSerialize["Rotate"] = o.Rotate
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
	if !IsNil(o.ResultType) {
		toSerialize["resultType"] = o.ResultType
	}
	if !IsNil(o.ResultTypeTable) {
		toSerialize["resultTypeTable"] = o.ResultTypeTable
	}
	if o.Regions != nil {
		toSerialize["regions"] = o.Regions
	}
	return toSerialize, nil
}

type NullableOCRSettingsDjVu2PDF struct {
	value *OCRSettingsDjVu2PDF
	isSet bool
}

func (v NullableOCRSettingsDjVu2PDF) Get() *OCRSettingsDjVu2PDF {
	return v.value
}

func (v *NullableOCRSettingsDjVu2PDF) Set(val *OCRSettingsDjVu2PDF) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRSettingsDjVu2PDF) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRSettingsDjVu2PDF) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRSettingsDjVu2PDF(val *OCRSettingsDjVu2PDF) *NullableOCRSettingsDjVu2PDF {
	return &NullableOCRSettingsDjVu2PDF{value: val, isSet: true}
}

func (v NullableOCRSettingsDjVu2PDF) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRSettingsDjVu2PDF) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


