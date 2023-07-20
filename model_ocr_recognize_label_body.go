/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRRecognizeLabelBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRRecognizeLabelBody{}

// OCRRecognizeLabelBody struct for OCRRecognizeLabelBody
type OCRRecognizeLabelBody struct {
	Image string `json:"image"`
	Settings OCRSettingsRecognizeLabel `json:"settings"`
}

// NewOCRRecognizeLabelBody instantiates a new OCRRecognizeLabelBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRRecognizeLabelBody(image string, settings OCRSettingsRecognizeLabel) *OCRRecognizeLabelBody {
	this := OCRRecognizeLabelBody{}
	this.Image = image
	this.Settings = settings
	return &this
}

// NewOCRRecognizeLabelBodyWithDefaults instantiates a new OCRRecognizeLabelBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRRecognizeLabelBodyWithDefaults() *OCRRecognizeLabelBody {
	this := OCRRecognizeLabelBody{}
	return &this
}

// GetImage returns the Image field value
func (o *OCRRecognizeLabelBody) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *OCRRecognizeLabelBody) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *OCRRecognizeLabelBody) SetImage(v string) {
	o.Image = v
}

// GetSettings returns the Settings field value
func (o *OCRRecognizeLabelBody) GetSettings() OCRSettingsRecognizeLabel {
	if o == nil {
		var ret OCRSettingsRecognizeLabel
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *OCRRecognizeLabelBody) GetSettingsOk() (*OCRSettingsRecognizeLabel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *OCRRecognizeLabelBody) SetSettings(v OCRSettingsRecognizeLabel) {
	o.Settings = v
}

func (o OCRRecognizeLabelBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRRecognizeLabelBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

type NullableOCRRecognizeLabelBody struct {
	value *OCRRecognizeLabelBody
	isSet bool
}

func (v NullableOCRRecognizeLabelBody) Get() *OCRRecognizeLabelBody {
	return v.value
}

func (v *NullableOCRRecognizeLabelBody) Set(val *OCRRecognizeLabelBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRRecognizeLabelBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRRecognizeLabelBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRRecognizeLabelBody(val *OCRRecognizeLabelBody) *NullableOCRRecognizeLabelBody {
	return &NullableOCRRecognizeLabelBody{value: val, isSet: true}
}

func (v NullableOCRRecognizeLabelBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRRecognizeLabelBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


