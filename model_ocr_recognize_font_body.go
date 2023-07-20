/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRRecognizeFontBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRRecognizeFontBody{}

// OCRRecognizeFontBody struct for OCRRecognizeFontBody
type OCRRecognizeFontBody struct {
	Image string `json:"image"`
	Settings OCRSettingsRecognizeFont `json:"settings"`
}

// NewOCRRecognizeFontBody instantiates a new OCRRecognizeFontBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRRecognizeFontBody(image string, settings OCRSettingsRecognizeFont) *OCRRecognizeFontBody {
	this := OCRRecognizeFontBody{}
	this.Image = image
	this.Settings = settings
	return &this
}

// NewOCRRecognizeFontBodyWithDefaults instantiates a new OCRRecognizeFontBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRRecognizeFontBodyWithDefaults() *OCRRecognizeFontBody {
	this := OCRRecognizeFontBody{}
	return &this
}

// GetImage returns the Image field value
func (o *OCRRecognizeFontBody) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *OCRRecognizeFontBody) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *OCRRecognizeFontBody) SetImage(v string) {
	o.Image = v
}

// GetSettings returns the Settings field value
func (o *OCRRecognizeFontBody) GetSettings() OCRSettingsRecognizeFont {
	if o == nil {
		var ret OCRSettingsRecognizeFont
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *OCRRecognizeFontBody) GetSettingsOk() (*OCRSettingsRecognizeFont, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *OCRRecognizeFontBody) SetSettings(v OCRSettingsRecognizeFont) {
	o.Settings = v
}

func (o OCRRecognizeFontBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRRecognizeFontBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

type NullableOCRRecognizeFontBody struct {
	value *OCRRecognizeFontBody
	isSet bool
}

func (v NullableOCRRecognizeFontBody) Get() *OCRRecognizeFontBody {
	return v.value
}

func (v *NullableOCRRecognizeFontBody) Set(val *OCRRecognizeFontBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRRecognizeFontBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRRecognizeFontBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRRecognizeFontBody(val *OCRRecognizeFontBody) *NullableOCRRecognizeFontBody {
	return &NullableOCRRecognizeFontBody{value: val, isSet: true}
}

func (v NullableOCRRecognizeFontBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRRecognizeFontBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


