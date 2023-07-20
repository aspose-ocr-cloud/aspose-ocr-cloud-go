/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRDetectRegionsBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRDetectRegionsBody{}

// OCRDetectRegionsBody struct for OCRDetectRegionsBody
type OCRDetectRegionsBody struct {
	Image string `json:"image"`
	Settings OCRSettingsDetectRegions `json:"settings"`
}

// NewOCRDetectRegionsBody instantiates a new OCRDetectRegionsBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRDetectRegionsBody(image string, settings OCRSettingsDetectRegions) *OCRDetectRegionsBody {
	this := OCRDetectRegionsBody{}
	this.Image = image
	this.Settings = settings
	return &this
}

// NewOCRDetectRegionsBodyWithDefaults instantiates a new OCRDetectRegionsBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRDetectRegionsBodyWithDefaults() *OCRDetectRegionsBody {
	this := OCRDetectRegionsBody{}
	return &this
}

// GetImage returns the Image field value
func (o *OCRDetectRegionsBody) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *OCRDetectRegionsBody) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *OCRDetectRegionsBody) SetImage(v string) {
	o.Image = v
}

// GetSettings returns the Settings field value
func (o *OCRDetectRegionsBody) GetSettings() OCRSettingsDetectRegions {
	if o == nil {
		var ret OCRSettingsDetectRegions
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *OCRDetectRegionsBody) GetSettingsOk() (*OCRSettingsDetectRegions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *OCRDetectRegionsBody) SetSettings(v OCRSettingsDetectRegions) {
	o.Settings = v
}

func (o OCRDetectRegionsBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRDetectRegionsBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

type NullableOCRDetectRegionsBody struct {
	value *OCRDetectRegionsBody
	isSet bool
}

func (v NullableOCRDetectRegionsBody) Get() *OCRDetectRegionsBody {
	return v.value
}

func (v *NullableOCRDetectRegionsBody) Set(val *OCRDetectRegionsBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRDetectRegionsBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRDetectRegionsBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRDetectRegionsBody(val *OCRDetectRegionsBody) *NullableOCRDetectRegionsBody {
	return &NullableOCRDetectRegionsBody{value: val, isSet: true}
}

func (v NullableOCRDetectRegionsBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRDetectRegionsBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


