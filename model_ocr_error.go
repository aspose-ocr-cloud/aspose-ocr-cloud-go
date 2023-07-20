/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRError{}

// OCRError Error to return to SDK client
type OCRError struct {
	// A list of various clear descriptions of the errors
	Messages []string `json:"messages,omitempty"`
	// Warning messages - non critical errors: e.g. some data lost
	Warnings []string `json:"warnings,omitempty"`
}

// NewOCRError instantiates a new OCRError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRError() *OCRError {
	this := OCRError{}
	return &this
}

// NewOCRErrorWithDefaults instantiates a new OCRError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRErrorWithDefaults() *OCRError {
	this := OCRError{}
	return &this
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OCRError) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OCRError) GetMessagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *OCRError) HasMessages() bool {
	if o != nil && IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *OCRError) SetMessages(v []string) {
	o.Messages = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OCRError) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OCRError) GetWarningsOk() ([]string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *OCRError) HasWarnings() bool {
	if o != nil && IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *OCRError) SetWarnings(v []string) {
	o.Warnings = v
}

func (o OCRError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableOCRError struct {
	value *OCRError
	isSet bool
}

func (v NullableOCRError) Get() *OCRError {
	return v.value
}

func (v *NullableOCRError) Set(val *OCRError) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRError) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRError(val *OCRError) *NullableOCRError {
	return &NullableOCRError{value: val, isSet: true}
}

func (v NullableOCRError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


