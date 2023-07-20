/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRDeskewImageBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRDeskewImageBody{}

// OCRDeskewImageBody struct for OCRDeskewImageBody
type OCRDeskewImageBody struct {
	Image string `json:"image"`
}

// NewOCRDeskewImageBody instantiates a new OCRDeskewImageBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRDeskewImageBody(image string) *OCRDeskewImageBody {
	this := OCRDeskewImageBody{}
	this.Image = image
	return &this
}

// NewOCRDeskewImageBodyWithDefaults instantiates a new OCRDeskewImageBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRDeskewImageBodyWithDefaults() *OCRDeskewImageBody {
	this := OCRDeskewImageBody{}
	return &this
}

// GetImage returns the Image field value
func (o *OCRDeskewImageBody) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *OCRDeskewImageBody) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *OCRDeskewImageBody) SetImage(v string) {
	o.Image = v
}

func (o OCRDeskewImageBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRDeskewImageBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	return toSerialize, nil
}

type NullableOCRDeskewImageBody struct {
	value *OCRDeskewImageBody
	isSet bool
}

func (v NullableOCRDeskewImageBody) Get() *OCRDeskewImageBody {
	return v.value
}

func (v *NullableOCRDeskewImageBody) Set(val *OCRDeskewImageBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRDeskewImageBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRDeskewImageBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRDeskewImageBody(val *OCRDeskewImageBody) *NullableOCRDeskewImageBody {
	return &NullableOCRDeskewImageBody{value: val, isSet: true}
}

func (v NullableOCRDeskewImageBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRDeskewImageBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


