/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRRegion{}

// OCRRegion Represents information about strict regions to recognize text
type OCRRegion struct {
	Rect *OCRRect `json:"rect,omitempty"`
	// The serial number of the region for the formation of the text
	Order *int32 `json:"order,omitempty"`
}

// NewOCRRegion instantiates a new OCRRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRRegion() *OCRRegion {
	this := OCRRegion{}
	return &this
}

// NewOCRRegionWithDefaults instantiates a new OCRRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRRegionWithDefaults() *OCRRegion {
	this := OCRRegion{}
	return &this
}

// GetRect returns the Rect field value if set, zero value otherwise.
func (o *OCRRegion) GetRect() OCRRect {
	if o == nil || IsNil(o.Rect) {
		var ret OCRRect
		return ret
	}
	return *o.Rect
}

// GetRectOk returns a tuple with the Rect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRRegion) GetRectOk() (*OCRRect, bool) {
	if o == nil || IsNil(o.Rect) {
		return nil, false
	}
	return o.Rect, true
}

// HasRect returns a boolean if a field has been set.
func (o *OCRRegion) HasRect() bool {
	if o != nil && !IsNil(o.Rect) {
		return true
	}

	return false
}

// SetRect gets a reference to the given OCRRect and assigns it to the Rect field.
func (o *OCRRegion) SetRect(v OCRRect) {
	o.Rect = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *OCRRegion) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRRegion) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *OCRRegion) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *OCRRegion) SetOrder(v int32) {
	o.Order = &v
}

func (o OCRRegion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rect) {
		toSerialize["rect"] = o.Rect
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return toSerialize, nil
}

type NullableOCRRegion struct {
	value *OCRRegion
	isSet bool
}

func (v NullableOCRRegion) Get() *OCRRegion {
	return v.value
}

func (v *NullableOCRRegion) Set(val *OCRRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRRegion(val *OCRRegion) *NullableOCRRegion {
	return &NullableOCRRegion{value: val, isSet: true}
}

func (v NullableOCRRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


