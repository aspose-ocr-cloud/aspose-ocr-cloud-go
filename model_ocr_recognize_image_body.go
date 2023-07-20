/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRRecognizeImageBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRRecognizeImageBody{}

// OCRRecognizeImageBody Combines Image data and OCR Recognition settings
type OCRRecognizeImageBody struct {
	// Gets or Sets Image
	Image string `json:"image"`
	Settings OCRSettingsRecognizeImage `json:"settings"`
}

// NewOCRRecognizeImageBody instantiates a new OCRRecognizeImageBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRRecognizeImageBody(image string, settings OCRSettingsRecognizeImage) *OCRRecognizeImageBody {
	this := OCRRecognizeImageBody{}
	this.Image = image
	this.Settings = settings
	return &this
}

// NewOCRRecognizeImageBodyWithDefaults instantiates a new OCRRecognizeImageBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRRecognizeImageBodyWithDefaults() *OCRRecognizeImageBody {
	this := OCRRecognizeImageBody{}
	return &this
}

// GetImage returns the Image field value
func (o *OCRRecognizeImageBody) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *OCRRecognizeImageBody) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *OCRRecognizeImageBody) SetImage(v string) {
	o.Image = v
}

// GetSettings returns the Settings field value
func (o *OCRRecognizeImageBody) GetSettings() OCRSettingsRecognizeImage {
	if o == nil {
		var ret OCRSettingsRecognizeImage
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *OCRRecognizeImageBody) GetSettingsOk() (*OCRSettingsRecognizeImage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *OCRRecognizeImageBody) SetSettings(v OCRSettingsRecognizeImage) {
	o.Settings = v
}

func (o OCRRecognizeImageBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRRecognizeImageBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

type NullableOCRRecognizeImageBody struct {
	value *OCRRecognizeImageBody
	isSet bool
}

func (v NullableOCRRecognizeImageBody) Get() *OCRRecognizeImageBody {
	return v.value
}

func (v *NullableOCRRecognizeImageBody) Set(val *OCRRecognizeImageBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRRecognizeImageBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRRecognizeImageBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRRecognizeImageBody(val *OCRRecognizeImageBody) *NullableOCRRecognizeImageBody {
	return &NullableOCRRecognizeImageBody{value: val, isSet: true}
}

func (v NullableOCRRecognizeImageBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRRecognizeImageBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


