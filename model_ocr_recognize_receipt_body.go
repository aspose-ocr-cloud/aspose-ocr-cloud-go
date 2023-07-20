/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRRecognizeReceiptBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRRecognizeReceiptBody{}

// OCRRecognizeReceiptBody Combines Image data and OCR Recognition settings fro Receipt image
type OCRRecognizeReceiptBody struct {
	// Gets or Sets Image
	Image string `json:"image"`
	Settings OCRSettingsRecognizeReceipt `json:"settings"`
}

// NewOCRRecognizeReceiptBody instantiates a new OCRRecognizeReceiptBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRRecognizeReceiptBody(image string, settings OCRSettingsRecognizeReceipt) *OCRRecognizeReceiptBody {
	this := OCRRecognizeReceiptBody{}
	this.Image = image
	this.Settings = settings
	return &this
}

// NewOCRRecognizeReceiptBodyWithDefaults instantiates a new OCRRecognizeReceiptBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRRecognizeReceiptBodyWithDefaults() *OCRRecognizeReceiptBody {
	this := OCRRecognizeReceiptBody{}
	return &this
}

// GetImage returns the Image field value
func (o *OCRRecognizeReceiptBody) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *OCRRecognizeReceiptBody) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *OCRRecognizeReceiptBody) SetImage(v string) {
	o.Image = v
}

// GetSettings returns the Settings field value
func (o *OCRRecognizeReceiptBody) GetSettings() OCRSettingsRecognizeReceipt {
	if o == nil {
		var ret OCRSettingsRecognizeReceipt
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *OCRRecognizeReceiptBody) GetSettingsOk() (*OCRSettingsRecognizeReceipt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *OCRRecognizeReceiptBody) SetSettings(v OCRSettingsRecognizeReceipt) {
	o.Settings = v
}

func (o OCRRecognizeReceiptBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRRecognizeReceiptBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

type NullableOCRRecognizeReceiptBody struct {
	value *OCRRecognizeReceiptBody
	isSet bool
}

func (v NullableOCRRecognizeReceiptBody) Get() *OCRRecognizeReceiptBody {
	return v.value
}

func (v *NullableOCRRecognizeReceiptBody) Set(val *OCRRecognizeReceiptBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRRecognizeReceiptBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRRecognizeReceiptBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRRecognizeReceiptBody(val *OCRRecognizeReceiptBody) *NullableOCRRecognizeReceiptBody {
	return &NullableOCRRecognizeReceiptBody{value: val, isSet: true}
}

func (v NullableOCRRecognizeReceiptBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRRecognizeReceiptBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


