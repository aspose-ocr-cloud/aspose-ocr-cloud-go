/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRDewarpImageBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRDewarpImageBody{}

// OCRDewarpImageBody struct for OCRDewarpImageBody
type OCRDewarpImageBody struct {
	Image string `json:"image"`
}

// NewOCRDewarpImageBody instantiates a new OCRDewarpImageBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRDewarpImageBody(image string) *OCRDewarpImageBody {
	this := OCRDewarpImageBody{}
	this.Image = image
	return &this
}

// NewOCRDewarpImageBodyWithDefaults instantiates a new OCRDewarpImageBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRDewarpImageBodyWithDefaults() *OCRDewarpImageBody {
	this := OCRDewarpImageBody{}
	return &this
}

// GetImage returns the Image field value
func (o *OCRDewarpImageBody) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *OCRDewarpImageBody) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *OCRDewarpImageBody) SetImage(v string) {
	o.Image = v
}

func (o OCRDewarpImageBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRDewarpImageBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	return toSerialize, nil
}

type NullableOCRDewarpImageBody struct {
	value *OCRDewarpImageBody
	isSet bool
}

func (v NullableOCRDewarpImageBody) Get() *OCRDewarpImageBody {
	return v.value
}

func (v *NullableOCRDewarpImageBody) Set(val *OCRDewarpImageBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRDewarpImageBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRDewarpImageBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRDewarpImageBody(val *OCRDewarpImageBody) *NullableOCRDewarpImageBody {
	return &NullableOCRDewarpImageBody{value: val, isSet: true}
}

func (v NullableOCRDewarpImageBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRDewarpImageBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


