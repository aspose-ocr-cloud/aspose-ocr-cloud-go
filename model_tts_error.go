/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the TTSError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TTSError{}

// TTSError struct for TTSError
type TTSError struct {
	Messages []string `json:"messages,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

// NewTTSError instantiates a new TTSError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTTSError() *TTSError {
	this := TTSError{}
	return &this
}

// NewTTSErrorWithDefaults instantiates a new TTSError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTTSErrorWithDefaults() *TTSError {
	this := TTSError{}
	return &this
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TTSError) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TTSError) GetMessagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *TTSError) HasMessages() bool {
	if o != nil && IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *TTSError) SetMessages(v []string) {
	o.Messages = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TTSError) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TTSError) GetWarningsOk() ([]string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *TTSError) HasWarnings() bool {
	if o != nil && IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *TTSError) SetWarnings(v []string) {
	o.Warnings = v
}

func (o TTSError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TTSError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableTTSError struct {
	value *TTSError
	isSet bool
}

func (v NullableTTSError) Get() *TTSError {
	return v.value
}

func (v *NullableTTSError) Set(val *TTSError) {
	v.value = val
	v.isSet = true
}

func (v NullableTTSError) IsSet() bool {
	return v.isSet
}

func (v *NullableTTSError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTTSError(val *TTSError) *NullableTTSError {
	return &NullableTTSError{value: val, isSet: true}
}

func (v NullableTTSError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTTSError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


