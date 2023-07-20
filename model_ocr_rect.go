/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRRect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRRect{}

// OCRRect Represents a rectangle: Left-Top (X1-Y1) to Right-Bottom (X2-Y2)
type OCRRect struct {
	// X-Coordinate of top left corner
	TopLeftX *int32 `json:"topLeftX,omitempty"`
	// Y-Coordinate of top left corner
	TopLeftY *int32 `json:"topLeftY,omitempty"`
	// X-Coordinate of bottom right corner
	BottomRightX *int32 `json:"bottomRightX,omitempty"`
	// Y-Coordinate of bottom right corner
	BottomRightY *int32 `json:"bottomRightY,omitempty"`
}

// NewOCRRect instantiates a new OCRRect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRRect() *OCRRect {
	this := OCRRect{}
	return &this
}

// NewOCRRectWithDefaults instantiates a new OCRRect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRRectWithDefaults() *OCRRect {
	this := OCRRect{}
	return &this
}

// GetTopLeftX returns the TopLeftX field value if set, zero value otherwise.
func (o *OCRRect) GetTopLeftX() int32 {
	if o == nil || IsNil(o.TopLeftX) {
		var ret int32
		return ret
	}
	return *o.TopLeftX
}

// GetTopLeftXOk returns a tuple with the TopLeftX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRRect) GetTopLeftXOk() (*int32, bool) {
	if o == nil || IsNil(o.TopLeftX) {
		return nil, false
	}
	return o.TopLeftX, true
}

// HasTopLeftX returns a boolean if a field has been set.
func (o *OCRRect) HasTopLeftX() bool {
	if o != nil && !IsNil(o.TopLeftX) {
		return true
	}

	return false
}

// SetTopLeftX gets a reference to the given int32 and assigns it to the TopLeftX field.
func (o *OCRRect) SetTopLeftX(v int32) {
	o.TopLeftX = &v
}

// GetTopLeftY returns the TopLeftY field value if set, zero value otherwise.
func (o *OCRRect) GetTopLeftY() int32 {
	if o == nil || IsNil(o.TopLeftY) {
		var ret int32
		return ret
	}
	return *o.TopLeftY
}

// GetTopLeftYOk returns a tuple with the TopLeftY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRRect) GetTopLeftYOk() (*int32, bool) {
	if o == nil || IsNil(o.TopLeftY) {
		return nil, false
	}
	return o.TopLeftY, true
}

// HasTopLeftY returns a boolean if a field has been set.
func (o *OCRRect) HasTopLeftY() bool {
	if o != nil && !IsNil(o.TopLeftY) {
		return true
	}

	return false
}

// SetTopLeftY gets a reference to the given int32 and assigns it to the TopLeftY field.
func (o *OCRRect) SetTopLeftY(v int32) {
	o.TopLeftY = &v
}

// GetBottomRightX returns the BottomRightX field value if set, zero value otherwise.
func (o *OCRRect) GetBottomRightX() int32 {
	if o == nil || IsNil(o.BottomRightX) {
		var ret int32
		return ret
	}
	return *o.BottomRightX
}

// GetBottomRightXOk returns a tuple with the BottomRightX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRRect) GetBottomRightXOk() (*int32, bool) {
	if o == nil || IsNil(o.BottomRightX) {
		return nil, false
	}
	return o.BottomRightX, true
}

// HasBottomRightX returns a boolean if a field has been set.
func (o *OCRRect) HasBottomRightX() bool {
	if o != nil && !IsNil(o.BottomRightX) {
		return true
	}

	return false
}

// SetBottomRightX gets a reference to the given int32 and assigns it to the BottomRightX field.
func (o *OCRRect) SetBottomRightX(v int32) {
	o.BottomRightX = &v
}

// GetBottomRightY returns the BottomRightY field value if set, zero value otherwise.
func (o *OCRRect) GetBottomRightY() int32 {
	if o == nil || IsNil(o.BottomRightY) {
		var ret int32
		return ret
	}
	return *o.BottomRightY
}

// GetBottomRightYOk returns a tuple with the BottomRightY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRRect) GetBottomRightYOk() (*int32, bool) {
	if o == nil || IsNil(o.BottomRightY) {
		return nil, false
	}
	return o.BottomRightY, true
}

// HasBottomRightY returns a boolean if a field has been set.
func (o *OCRRect) HasBottomRightY() bool {
	if o != nil && !IsNil(o.BottomRightY) {
		return true
	}

	return false
}

// SetBottomRightY gets a reference to the given int32 and assigns it to the BottomRightY field.
func (o *OCRRect) SetBottomRightY(v int32) {
	o.BottomRightY = &v
}

func (o OCRRect) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRRect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TopLeftX) {
		toSerialize["topLeftX"] = o.TopLeftX
	}
	if !IsNil(o.TopLeftY) {
		toSerialize["topLeftY"] = o.TopLeftY
	}
	if !IsNil(o.BottomRightX) {
		toSerialize["bottomRightX"] = o.BottomRightX
	}
	if !IsNil(o.BottomRightY) {
		toSerialize["bottomRightY"] = o.BottomRightY
	}
	return toSerialize, nil
}

type NullableOCRRect struct {
	value *OCRRect
	isSet bool
}

func (v NullableOCRRect) Get() *OCRRect {
	return v.value
}

func (v *NullableOCRRect) Set(val *OCRRect) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRRect) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRRect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRRect(val *OCRRect) *NullableOCRRect {
	return &NullableOCRRect{value: val, isSet: true}
}

func (v NullableOCRRect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRRect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


