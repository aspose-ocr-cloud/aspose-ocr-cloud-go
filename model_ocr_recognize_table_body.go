/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRRecognizeTableBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRRecognizeTableBody{}

// OCRRecognizeTableBody Combines Image data and OCR Recognition settings for Table image
type OCRRecognizeTableBody struct {
	// Gets or Sets Image
	Image string `json:"image"`
	Settings OCRSettingsRecognizeTable `json:"settings"`
}

// NewOCRRecognizeTableBody instantiates a new OCRRecognizeTableBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRRecognizeTableBody(image string, settings OCRSettingsRecognizeTable) *OCRRecognizeTableBody {
	this := OCRRecognizeTableBody{}
	this.Image = image
	this.Settings = settings
	return &this
}

// NewOCRRecognizeTableBodyWithDefaults instantiates a new OCRRecognizeTableBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRRecognizeTableBodyWithDefaults() *OCRRecognizeTableBody {
	this := OCRRecognizeTableBody{}
	return &this
}

// GetImage returns the Image field value
func (o *OCRRecognizeTableBody) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *OCRRecognizeTableBody) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *OCRRecognizeTableBody) SetImage(v string) {
	o.Image = v
}

// GetSettings returns the Settings field value
func (o *OCRRecognizeTableBody) GetSettings() OCRSettingsRecognizeTable {
	if o == nil {
		var ret OCRSettingsRecognizeTable
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *OCRRecognizeTableBody) GetSettingsOk() (*OCRSettingsRecognizeTable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *OCRRecognizeTableBody) SetSettings(v OCRSettingsRecognizeTable) {
	o.Settings = v
}

func (o OCRRecognizeTableBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRRecognizeTableBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

type NullableOCRRecognizeTableBody struct {
	value *OCRRecognizeTableBody
	isSet bool
}

func (v NullableOCRRecognizeTableBody) Get() *OCRRecognizeTableBody {
	return v.value
}

func (v *NullableOCRRecognizeTableBody) Set(val *OCRRecognizeTableBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRRecognizeTableBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRRecognizeTableBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRRecognizeTableBody(val *OCRRecognizeTableBody) *NullableOCRRecognizeTableBody {
	return &NullableOCRRecognizeTableBody{value: val, isSet: true}
}

func (v NullableOCRRecognizeTableBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRRecognizeTableBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


