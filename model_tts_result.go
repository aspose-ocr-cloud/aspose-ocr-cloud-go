/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the TTSResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TTSResult{}

// TTSResult struct for TTSResult
type TTSResult struct {
	Type NullableString `json:"type,omitempty"`
	Data NullableString `json:"data,omitempty"`
}

// NewTTSResult instantiates a new TTSResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTTSResult() *TTSResult {
	this := TTSResult{}
	return &this
}

// NewTTSResultWithDefaults instantiates a new TTSResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTTSResultWithDefaults() *TTSResult {
	this := TTSResult{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TTSResult) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TTSResult) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *TTSResult) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *TTSResult) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *TTSResult) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *TTSResult) UnsetType() {
	o.Type.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TTSResult) GetData() string {
	if o == nil || IsNil(o.Data.Get()) {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TTSResult) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *TTSResult) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableString and assigns it to the Data field.
func (o *TTSResult) SetData(v string) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *TTSResult) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *TTSResult) UnsetData() {
	o.Data.Unset()
}

func (o TTSResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TTSResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return toSerialize, nil
}

type NullableTTSResult struct {
	value *TTSResult
	isSet bool
}

func (v NullableTTSResult) Get() *TTSResult {
	return v.value
}

func (v *NullableTTSResult) Set(val *TTSResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTTSResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTTSResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTTSResult(val *TTSResult) *NullableTTSResult {
	return &NullableTTSResult{value: val, isSet: true}
}

func (v NullableTTSResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTTSResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


