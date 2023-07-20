/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRRecognizePdfBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRRecognizePdfBody{}

// OCRRecognizePdfBody Combines Image data and OCR Recognition settings for PDF
type OCRRecognizePdfBody struct {
	// Gets or Sets Image
	Image string `json:"image"`
	Settings OCRSettingsRecognizePdf `json:"settings"`
}

// NewOCRRecognizePdfBody instantiates a new OCRRecognizePdfBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRRecognizePdfBody(image string, settings OCRSettingsRecognizePdf) *OCRRecognizePdfBody {
	this := OCRRecognizePdfBody{}
	this.Image = image
	this.Settings = settings
	return &this
}

// NewOCRRecognizePdfBodyWithDefaults instantiates a new OCRRecognizePdfBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRRecognizePdfBodyWithDefaults() *OCRRecognizePdfBody {
	this := OCRRecognizePdfBody{}
	return &this
}

// GetImage returns the Image field value
func (o *OCRRecognizePdfBody) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *OCRRecognizePdfBody) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *OCRRecognizePdfBody) SetImage(v string) {
	o.Image = v
}

// GetSettings returns the Settings field value
func (o *OCRRecognizePdfBody) GetSettings() OCRSettingsRecognizePdf {
	if o == nil {
		var ret OCRSettingsRecognizePdf
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *OCRRecognizePdfBody) GetSettingsOk() (*OCRSettingsRecognizePdf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *OCRRecognizePdfBody) SetSettings(v OCRSettingsRecognizePdf) {
	o.Settings = v
}

func (o OCRRecognizePdfBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRRecognizePdfBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

type NullableOCRRecognizePdfBody struct {
	value *OCRRecognizePdfBody
	isSet bool
}

func (v NullableOCRRecognizePdfBody) Get() *OCRRecognizePdfBody {
	return v.value
}

func (v *NullableOCRRecognizePdfBody) Set(val *OCRRecognizePdfBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRRecognizePdfBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRRecognizePdfBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRRecognizePdfBody(val *OCRRecognizePdfBody) *NullableOCRRecognizePdfBody {
	return &NullableOCRRecognizePdfBody{value: val, isSet: true}
}

func (v NullableOCRRecognizePdfBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRRecognizePdfBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


