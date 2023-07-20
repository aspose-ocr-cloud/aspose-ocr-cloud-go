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

// Language Recognition Language
type Language string

// List of Language
const (
	LANGUAGE_ENGLISH Language = "English"
	LANGUAGE_GERMAN Language = "German"
	LANGUAGE_FRENCH Language = "French"
	LANGUAGE_ITALIAN Language = "Italian"
	LANGUAGE_SPANISH Language = "Spanish"
	LANGUAGE_PORTUGUESE Language = "Portuguese"
	LANGUAGE_POLISH Language = "Polish"
	LANGUAGE_SLOVENE Language = "Slovene"
	LANGUAGE_SLOVAK Language = "Slovak"
	LANGUAGE_NETHERLANDS Language = "Netherlands"
	LANGUAGE_LITHUANIAN Language = "Lithuanian"
	LANGUAGE_LATVIAN Language = "Latvian"
	LANGUAGE_DANISH Language = "Danish"
	LANGUAGE_NORWEGIAN Language = "Norwegian"
	LANGUAGE_FINNISH Language = "Finnish"
	LANGUAGE_SERBIAN Language = "Serbian"
	LANGUAGE_CROATIAN Language = "Croatian"
	LANGUAGE_CZECH Language = "Czech"
	LANGUAGE_SWEDISH Language = "Swedish"
	LANGUAGE_ESTONIAN Language = "Estonian"
	LANGUAGE_ROMANIAN Language = "Romanian"
	LANGUAGE_CHINESE Language = "Chinese"
	LANGUAGE_RUSSIAN Language = "Russian"
	LANGUAGE_ARABIC Language = "Arabic"
	LANGUAGE_HINDI Language = "Hindi"
	LANGUAGE_UKRAINAN Language = "Ukrainan"
	LANGUAGE_BENGALI Language = "Bengali"
	LANGUAGE_TIBETAN Language = "Tibetan"
	LANGUAGE_THAI Language = "Thai"
	LANGUAGE_URDU Language = "Urdu"
	LANGUAGE_TURKISH Language = "Turkish"
	LANGUAGE_KOREAN Language = "Korean"
	LANGUAGE_INDONESIAN Language = "Indonesian"
	LANGUAGE_HEBREW Language = "Hebrew"
	LANGUAGE_JAVANESE Language = "Javanese"
	LANGUAGE_GREEK Language = "Greek"
	LANGUAGE_JAPANESE Language = "Japanese"
	LANGUAGE_PERSIAN Language = "Persian"
	LANGUAGE_ALBANIAN Language = "Albanian"
	LANGUAGE_LATIN Language = "Latin"
	LANGUAGE_VIETNAMESE Language = "Vietnamese"
	LANGUAGE_UZBEK Language = "Uzbek"
	LANGUAGE_GEORGIAN Language = "Georgian"
	LANGUAGE_BULGARIAN Language = "Bulgarian"
	LANGUAGE_AZERBAIJANI Language = "Azerbaijani"
	LANGUAGE_KAZAH Language = "Kazah"
	LANGUAGE_MACEDONIAN Language = "Macedonian"
	LANGUAGE_BELORUSSIAN Language = "Belorussian"
	LANGUAGE_HWT_ENG Language = "HWT_eng"
)

// All allowed values of Language enum
var AllowedLanguageEnumValues = []Language{
	"English",
	"German",
	"French",
	"Italian",
	"Spanish",
	"Portuguese",
	"Polish",
	"Slovene",
	"Slovak",
	"Netherlands",
	"Lithuanian",
	"Latvian",
	"Danish",
	"Norwegian",
	"Finnish",
	"Serbian",
	"Croatian",
	"Czech",
	"Swedish",
	"Estonian",
	"Romanian",
	"Chinese",
	"Russian",
	"Arabic",
	"Hindi",
	"Ukrainan",
	"Bengali",
	"Tibetan",
	"Thai",
	"Urdu",
	"Turkish",
	"Korean",
	"Indonesian",
	"Hebrew",
	"Javanese",
	"Greek",
	"Japanese",
	"Persian",
	"Albanian",
	"Latin",
	"Vietnamese",
	"Uzbek",
	"Georgian",
	"Bulgarian",
	"Azerbaijani",
	"Kazah",
	"Macedonian",
	"Belorussian",
	"HWT_eng",
}

func (v *Language) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Language(value)
	for _, existing := range AllowedLanguageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Language", value)
}

// NewLanguageFromValue returns a pointer to a valid Language
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLanguageFromValue(v string) (*Language, error) {
	ev := Language(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Language: valid values are %v", v, AllowedLanguageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Language) IsValid() bool {
	for _, existing := range AllowedLanguageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Language value
func (v Language) Ptr() *Language {
	return &v
}

type NullableLanguage struct {
	value *Language
	isSet bool
}

func (v NullableLanguage) Get() *Language {
	return v.value
}

func (v *NullableLanguage) Set(val *Language) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguage(val *Language) *NullableLanguage {
	return &NullableLanguage{value: val, isSet: true}
}

func (v NullableLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

