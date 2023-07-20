/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
	"fmt"
)

// LanguageTTS the model 'LanguageTTS'
type LanguageTTS string

// List of LanguageTTS
const (
	LANGUAGETTS_ENGLISH LanguageTTS = "English"
)

// All allowed values of LanguageTTS enum
var AllowedLanguageTTSEnumValues = []LanguageTTS{
	"English",
}

func (v *LanguageTTS) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LanguageTTS(value)
	for _, existing := range AllowedLanguageTTSEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LanguageTTS", value)
}

// NewLanguageTTSFromValue returns a pointer to a valid LanguageTTS
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLanguageTTSFromValue(v string) (*LanguageTTS, error) {
	ev := LanguageTTS(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LanguageTTS: valid values are %v", v, AllowedLanguageTTSEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LanguageTTS) IsValid() bool {
	for _, existing := range AllowedLanguageTTSEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LanguageTTS value
func (v LanguageTTS) Ptr() *LanguageTTS {
	return &v
}

type NullableLanguageTTS struct {
	value *LanguageTTS
	isSet bool
}

func (v NullableLanguageTTS) Get() *LanguageTTS {
	return v.value
}

func (v *NullableLanguageTTS) Set(val *LanguageTTS) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguageTTS) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguageTTS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguageTTS(val *LanguageTTS) *NullableLanguageTTS {
	return &NullableLanguageTTS{value: val, isSet: true}
}

func (v NullableLanguageTTS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguageTTS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

