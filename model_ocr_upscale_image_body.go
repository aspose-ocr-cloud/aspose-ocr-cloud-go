/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRUpscaleImageBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRUpscaleImageBody{}

// OCRUpscaleImageBody struct for OCRUpscaleImageBody
type OCRUpscaleImageBody struct {
	Image string `json:"image"`
}

// NewOCRUpscaleImageBody instantiates a new OCRUpscaleImageBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRUpscaleImageBody(image string) *OCRUpscaleImageBody {
	this := OCRUpscaleImageBody{}
	this.Image = image
	return &this
}

// NewOCRUpscaleImageBodyWithDefaults instantiates a new OCRUpscaleImageBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRUpscaleImageBodyWithDefaults() *OCRUpscaleImageBody {
	this := OCRUpscaleImageBody{}
	return &this
}

// GetImage returns the Image field value
func (o *OCRUpscaleImageBody) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *OCRUpscaleImageBody) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *OCRUpscaleImageBody) SetImage(v string) {
	o.Image = v
}

func (o OCRUpscaleImageBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRUpscaleImageBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	return toSerialize, nil
}

type NullableOCRUpscaleImageBody struct {
	value *OCRUpscaleImageBody
	isSet bool
}

func (v NullableOCRUpscaleImageBody) Get() *OCRUpscaleImageBody {
	return v.value
}

func (v *NullableOCRUpscaleImageBody) Set(val *OCRUpscaleImageBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRUpscaleImageBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRUpscaleImageBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRUpscaleImageBody(val *OCRUpscaleImageBody) *NullableOCRUpscaleImageBody {
	return &NullableOCRUpscaleImageBody{value: val, isSet: true}
}

func (v NullableOCRUpscaleImageBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRUpscaleImageBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


