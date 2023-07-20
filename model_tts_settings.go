/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the TTSSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TTSSettings{}

// TTSSettings struct for TTSSettings
type TTSSettings struct {
	Language LanguageTTS `json:"language"`
	ResultType ResultTypeTTS `json:"resultType"`
}

// NewTTSSettings instantiates a new TTSSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTTSSettings(language LanguageTTS, resultType ResultTypeTTS) *TTSSettings {
	this := TTSSettings{}
	this.Language = language
	this.ResultType = resultType
	return &this
}

// NewTTSSettingsWithDefaults instantiates a new TTSSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTTSSettingsWithDefaults() *TTSSettings {
	this := TTSSettings{}
	return &this
}

// GetLanguage returns the Language field value
func (o *TTSSettings) GetLanguage() LanguageTTS {
	if o == nil {
		var ret LanguageTTS
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *TTSSettings) GetLanguageOk() (*LanguageTTS, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *TTSSettings) SetLanguage(v LanguageTTS) {
	o.Language = v
}

// GetResultType returns the ResultType field value
func (o *TTSSettings) GetResultType() ResultTypeTTS {
	if o == nil {
		var ret ResultTypeTTS
		return ret
	}

	return o.ResultType
}

// GetResultTypeOk returns a tuple with the ResultType field value
// and a boolean to check if the value has been set.
func (o *TTSSettings) GetResultTypeOk() (*ResultTypeTTS, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultType, true
}

// SetResultType sets field value
func (o *TTSSettings) SetResultType(v ResultTypeTTS) {
	o.ResultType = v
}

func (o TTSSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TTSSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["language"] = o.Language
	toSerialize["resultType"] = o.ResultType
	return toSerialize, nil
}

type NullableTTSSettings struct {
	value *TTSSettings
	isSet bool
}

func (v NullableTTSSettings) Get() *TTSSettings {
	return v.value
}

func (v *NullableTTSSettings) Set(val *TTSSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTTSSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTTSSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTTSSettings(val *TTSSettings) *NullableTTSSettings {
	return &NullableTTSSettings{value: val, isSet: true}
}

func (v NullableTTSSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTTSSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


