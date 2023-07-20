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

// ResultTypeTTS the model 'ResultTypeTTS'
type ResultTypeTTS string

// List of ResultTypeTTS
const (
	RESULTTYPETTS_WAV ResultTypeTTS = "Wav"
)

// All allowed values of ResultTypeTTS enum
var AllowedResultTypeTTSEnumValues = []ResultTypeTTS{
	"Wav",
}

func (v *ResultTypeTTS) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResultTypeTTS(value)
	for _, existing := range AllowedResultTypeTTSEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResultTypeTTS", value)
}

// NewResultTypeTTSFromValue returns a pointer to a valid ResultTypeTTS
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResultTypeTTSFromValue(v string) (*ResultTypeTTS, error) {
	ev := ResultTypeTTS(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResultTypeTTS: valid values are %v", v, AllowedResultTypeTTSEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResultTypeTTS) IsValid() bool {
	for _, existing := range AllowedResultTypeTTSEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResultTypeTTS value
func (v ResultTypeTTS) Ptr() *ResultTypeTTS {
	return &v
}

type NullableResultTypeTTS struct {
	value *ResultTypeTTS
	isSet bool
}

func (v NullableResultTypeTTS) Get() *ResultTypeTTS {
	return v.value
}

func (v *NullableResultTypeTTS) Set(val *ResultTypeTTS) {
	v.value = val
	v.isSet = true
}

func (v NullableResultTypeTTS) IsSet() bool {
	return v.isSet
}

func (v *NullableResultTypeTTS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultTypeTTS(val *ResultTypeTTS) *NullableResultTypeTTS {
	return &NullableResultTypeTTS{value: val, isSet: true}
}

func (v NullableResultTypeTTS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultTypeTTS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

