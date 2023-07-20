/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRResult{}

// OCRResult Represents information about response after OCR.
type OCRResult struct {
	// File data type (extension)
	Type NullableString `json:"type,omitempty"`
	// File binary data
	Data NullableString `json:"data,omitempty"`
}

// NewOCRResult instantiates a new OCRResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRResult() *OCRResult {
	this := OCRResult{}
	return &this
}

// NewOCRResultWithDefaults instantiates a new OCRResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRResultWithDefaults() *OCRResult {
	this := OCRResult{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OCRResult) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OCRResult) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *OCRResult) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *OCRResult) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *OCRResult) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *OCRResult) UnsetType() {
	o.Type.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OCRResult) GetData() string {
	if o == nil || IsNil(o.Data.Get()) {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OCRResult) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *OCRResult) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableString and assigns it to the Data field.
func (o *OCRResult) SetData(v string) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *OCRResult) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *OCRResult) UnsetData() {
	o.Data.Unset()
}

func (o OCRResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return toSerialize, nil
}

type NullableOCRResult struct {
	value *OCRResult
	isSet bool
}

func (v NullableOCRResult) Get() *OCRResult {
	return v.value
}

func (v *NullableOCRResult) Set(val *OCRResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRResult(val *OCRResult) *NullableOCRResult {
	return &NullableOCRResult{value: val, isSet: true}
}

func (v NullableOCRResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


