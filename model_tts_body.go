/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the TTSBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TTSBody{}

// TTSBody struct for TTSBody
type TTSBody struct {
	Text string `json:"text"`
	Settings TTSSettings `json:"settings"`
}

// NewTTSBody instantiates a new TTSBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTTSBody(text string, settings TTSSettings) *TTSBody {
	this := TTSBody{}
	this.Text = text
	this.Settings = settings
	return &this
}

// NewTTSBodyWithDefaults instantiates a new TTSBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTTSBodyWithDefaults() *TTSBody {
	this := TTSBody{}
	return &this
}

// GetText returns the Text field value
func (o *TTSBody) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *TTSBody) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *TTSBody) SetText(v string) {
	o.Text = v
}

// GetSettings returns the Settings field value
func (o *TTSBody) GetSettings() TTSSettings {
	if o == nil {
		var ret TTSSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *TTSBody) GetSettingsOk() (*TTSSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *TTSBody) SetSettings(v TTSSettings) {
	o.Settings = v
}

func (o TTSBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TTSBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

type NullableTTSBody struct {
	value *TTSBody
	isSet bool
}

func (v NullableTTSBody) Get() *TTSBody {
	return v.value
}

func (v *NullableTTSBody) Set(val *TTSBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTTSBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTTSBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTTSBody(val *TTSBody) *NullableTTSBody {
	return &NullableTTSBody{value: val, isSet: true}
}

func (v NullableTTSBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTTSBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


