/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRRecognizeAndParseInvoiceBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRRecognizeAndParseInvoiceBody{}

// OCRRecognizeAndParseInvoiceBody struct for OCRRecognizeAndParseInvoiceBody
type OCRRecognizeAndParseInvoiceBody struct {
	Image string `json:"image"`
	Settings OCRSettingsRecognizeAndParseInvoice `json:"settings"`
}

// NewOCRRecognizeAndParseInvoiceBody instantiates a new OCRRecognizeAndParseInvoiceBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRRecognizeAndParseInvoiceBody(image string, settings OCRSettingsRecognizeAndParseInvoice) *OCRRecognizeAndParseInvoiceBody {
	this := OCRRecognizeAndParseInvoiceBody{}
	this.Image = image
	this.Settings = settings
	return &this
}

// NewOCRRecognizeAndParseInvoiceBodyWithDefaults instantiates a new OCRRecognizeAndParseInvoiceBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRRecognizeAndParseInvoiceBodyWithDefaults() *OCRRecognizeAndParseInvoiceBody {
	this := OCRRecognizeAndParseInvoiceBody{}
	return &this
}

// GetImage returns the Image field value
func (o *OCRRecognizeAndParseInvoiceBody) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *OCRRecognizeAndParseInvoiceBody) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *OCRRecognizeAndParseInvoiceBody) SetImage(v string) {
	o.Image = v
}

// GetSettings returns the Settings field value
func (o *OCRRecognizeAndParseInvoiceBody) GetSettings() OCRSettingsRecognizeAndParseInvoice {
	if o == nil {
		var ret OCRSettingsRecognizeAndParseInvoice
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *OCRRecognizeAndParseInvoiceBody) GetSettingsOk() (*OCRSettingsRecognizeAndParseInvoice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *OCRRecognizeAndParseInvoiceBody) SetSettings(v OCRSettingsRecognizeAndParseInvoice) {
	o.Settings = v
}

func (o OCRRecognizeAndParseInvoiceBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRRecognizeAndParseInvoiceBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

type NullableOCRRecognizeAndParseInvoiceBody struct {
	value *OCRRecognizeAndParseInvoiceBody
	isSet bool
}

func (v NullableOCRRecognizeAndParseInvoiceBody) Get() *OCRRecognizeAndParseInvoiceBody {
	return v.value
}

func (v *NullableOCRRecognizeAndParseInvoiceBody) Set(val *OCRRecognizeAndParseInvoiceBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRRecognizeAndParseInvoiceBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRRecognizeAndParseInvoiceBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRRecognizeAndParseInvoiceBody(val *OCRRecognizeAndParseInvoiceBody) *NullableOCRRecognizeAndParseInvoiceBody {
	return &NullableOCRRecognizeAndParseInvoiceBody{value: val, isSet: true}
}

func (v NullableOCRRecognizeAndParseInvoiceBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRRecognizeAndParseInvoiceBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


