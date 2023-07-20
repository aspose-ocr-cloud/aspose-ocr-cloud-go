/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRDjVu2PDFBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRDjVu2PDFBody{}

// OCRDjVu2PDFBody struct for OCRDjVu2PDFBody
type OCRDjVu2PDFBody struct {
	Image string `json:"image"`
	Settings OCRSettingsDjVu2PDF `json:"settings"`
}

// NewOCRDjVu2PDFBody instantiates a new OCRDjVu2PDFBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRDjVu2PDFBody(image string, settings OCRSettingsDjVu2PDF) *OCRDjVu2PDFBody {
	this := OCRDjVu2PDFBody{}
	this.Image = image
	this.Settings = settings
	return &this
}

// NewOCRDjVu2PDFBodyWithDefaults instantiates a new OCRDjVu2PDFBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRDjVu2PDFBodyWithDefaults() *OCRDjVu2PDFBody {
	this := OCRDjVu2PDFBody{}
	return &this
}

// GetImage returns the Image field value
func (o *OCRDjVu2PDFBody) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *OCRDjVu2PDFBody) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *OCRDjVu2PDFBody) SetImage(v string) {
	o.Image = v
}

// GetSettings returns the Settings field value
func (o *OCRDjVu2PDFBody) GetSettings() OCRSettingsDjVu2PDF {
	if o == nil {
		var ret OCRSettingsDjVu2PDF
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *OCRDjVu2PDFBody) GetSettingsOk() (*OCRSettingsDjVu2PDF, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *OCRDjVu2PDFBody) SetSettings(v OCRSettingsDjVu2PDF) {
	o.Settings = v
}

func (o OCRDjVu2PDFBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRDjVu2PDFBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

type NullableOCRDjVu2PDFBody struct {
	value *OCRDjVu2PDFBody
	isSet bool
}

func (v NullableOCRDjVu2PDFBody) Get() *OCRDjVu2PDFBody {
	return v.value
}

func (v *NullableOCRDjVu2PDFBody) Set(val *OCRDjVu2PDFBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRDjVu2PDFBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRDjVu2PDFBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRDjVu2PDFBody(val *OCRDjVu2PDFBody) *NullableOCRDjVu2PDFBody {
	return &NullableOCRDjVu2PDFBody{value: val, isSet: true}
}

func (v NullableOCRDjVu2PDFBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRDjVu2PDFBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


