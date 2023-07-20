/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the TTSBodyDeprecated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TTSBodyDeprecated{}

// TTSBodyDeprecated struct for TTSBodyDeprecated
type TTSBodyDeprecated struct {
	Text string `json:"text"`
	Language LanguageTTS `json:"language"`
	ResultType ResultTypeTTS `json:"resultType"`
}

// NewTTSBodyDeprecated instantiates a new TTSBodyDeprecated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTTSBodyDeprecated(text string, language LanguageTTS, resultType ResultTypeTTS) *TTSBodyDeprecated {
	this := TTSBodyDeprecated{}
	this.Text = text
	this.Language = language
	this.ResultType = resultType
	return &this
}

// NewTTSBodyDeprecatedWithDefaults instantiates a new TTSBodyDeprecated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTTSBodyDeprecatedWithDefaults() *TTSBodyDeprecated {
	this := TTSBodyDeprecated{}
	return &this
}

// GetText returns the Text field value
func (o *TTSBodyDeprecated) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *TTSBodyDeprecated) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *TTSBodyDeprecated) SetText(v string) {
	o.Text = v
}

// GetLanguage returns the Language field value
func (o *TTSBodyDeprecated) GetLanguage() LanguageTTS {
	if o == nil {
		var ret LanguageTTS
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *TTSBodyDeprecated) GetLanguageOk() (*LanguageTTS, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *TTSBodyDeprecated) SetLanguage(v LanguageTTS) {
	o.Language = v
}

// GetResultType returns the ResultType field value
func (o *TTSBodyDeprecated) GetResultType() ResultTypeTTS {
	if o == nil {
		var ret ResultTypeTTS
		return ret
	}

	return o.ResultType
}

// GetResultTypeOk returns a tuple with the ResultType field value
// and a boolean to check if the value has been set.
func (o *TTSBodyDeprecated) GetResultTypeOk() (*ResultTypeTTS, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultType, true
}

// SetResultType sets field value
func (o *TTSBodyDeprecated) SetResultType(v ResultTypeTTS) {
	o.ResultType = v
}

func (o TTSBodyDeprecated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TTSBodyDeprecated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	toSerialize["language"] = o.Language
	toSerialize["resultType"] = o.ResultType
	return toSerialize, nil
}

type NullableTTSBodyDeprecated struct {
	value *TTSBodyDeprecated
	isSet bool
}

func (v NullableTTSBodyDeprecated) Get() *TTSBodyDeprecated {
	return v.value
}

func (v *NullableTTSBodyDeprecated) Set(val *TTSBodyDeprecated) {
	v.value = val
	v.isSet = true
}

func (v NullableTTSBodyDeprecated) IsSet() bool {
	return v.isSet
}

func (v *NullableTTSBodyDeprecated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTTSBodyDeprecated(val *TTSBodyDeprecated) *NullableTTSBodyDeprecated {
	return &NullableTTSBodyDeprecated{value: val, isSet: true}
}

func (v NullableTTSBodyDeprecated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTTSBodyDeprecated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


