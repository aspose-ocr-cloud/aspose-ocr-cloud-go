/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRBinarizeImageBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRBinarizeImageBody{}

// OCRBinarizeImageBody struct for OCRBinarizeImageBody
type OCRBinarizeImageBody struct {
	Image string `json:"image"`
}

// NewOCRBinarizeImageBody instantiates a new OCRBinarizeImageBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRBinarizeImageBody(image string) *OCRBinarizeImageBody {
	this := OCRBinarizeImageBody{}
	this.Image = image
	return &this
}

// NewOCRBinarizeImageBodyWithDefaults instantiates a new OCRBinarizeImageBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRBinarizeImageBodyWithDefaults() *OCRBinarizeImageBody {
	this := OCRBinarizeImageBody{}
	return &this
}

// GetImage returns the Image field value
func (o *OCRBinarizeImageBody) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *OCRBinarizeImageBody) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *OCRBinarizeImageBody) SetImage(v string) {
	o.Image = v
}

func (o OCRBinarizeImageBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRBinarizeImageBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	return toSerialize, nil
}

type NullableOCRBinarizeImageBody struct {
	value *OCRBinarizeImageBody
	isSet bool
}

func (v NullableOCRBinarizeImageBody) Get() *OCRBinarizeImageBody {
	return v.value
}

func (v *NullableOCRBinarizeImageBody) Set(val *OCRBinarizeImageBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRBinarizeImageBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRBinarizeImageBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRBinarizeImageBody(val *OCRBinarizeImageBody) *NullableOCRBinarizeImageBody {
	return &NullableOCRBinarizeImageBody{value: val, isSet: true}
}

func (v NullableOCRBinarizeImageBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRBinarizeImageBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


