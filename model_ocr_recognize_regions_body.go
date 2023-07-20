/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRRecognizeRegionsBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRRecognizeRegionsBody{}

// OCRRecognizeRegionsBody struct for OCRRecognizeRegionsBody
type OCRRecognizeRegionsBody struct {
	Image string `json:"image"`
	Settings OCRSettingsRecognizeRegions `json:"settings"`
}

// NewOCRRecognizeRegionsBody instantiates a new OCRRecognizeRegionsBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRRecognizeRegionsBody(image string, settings OCRSettingsRecognizeRegions) *OCRRecognizeRegionsBody {
	this := OCRRecognizeRegionsBody{}
	this.Image = image
	this.Settings = settings
	return &this
}

// NewOCRRecognizeRegionsBodyWithDefaults instantiates a new OCRRecognizeRegionsBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRRecognizeRegionsBodyWithDefaults() *OCRRecognizeRegionsBody {
	this := OCRRecognizeRegionsBody{}
	return &this
}

// GetImage returns the Image field value
func (o *OCRRecognizeRegionsBody) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *OCRRecognizeRegionsBody) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *OCRRecognizeRegionsBody) SetImage(v string) {
	o.Image = v
}

// GetSettings returns the Settings field value
func (o *OCRRecognizeRegionsBody) GetSettings() OCRSettingsRecognizeRegions {
	if o == nil {
		var ret OCRSettingsRecognizeRegions
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *OCRRecognizeRegionsBody) GetSettingsOk() (*OCRSettingsRecognizeRegions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *OCRRecognizeRegionsBody) SetSettings(v OCRSettingsRecognizeRegions) {
	o.Settings = v
}

func (o OCRRecognizeRegionsBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRRecognizeRegionsBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

type NullableOCRRecognizeRegionsBody struct {
	value *OCRRecognizeRegionsBody
	isSet bool
}

func (v NullableOCRRecognizeRegionsBody) Get() *OCRRecognizeRegionsBody {
	return v.value
}

func (v *NullableOCRRecognizeRegionsBody) Set(val *OCRRecognizeRegionsBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRRecognizeRegionsBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRRecognizeRegionsBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRRecognizeRegionsBody(val *OCRRecognizeRegionsBody) *NullableOCRRecognizeRegionsBody {
	return &NullableOCRRecognizeRegionsBody{value: val, isSet: true}
}

func (v NullableOCRRecognizeRegionsBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRRecognizeRegionsBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


