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

// DsrConfidence Region filtering threshold. Filtering by the algorithm confidence scale. Higher mode - more aggressive filtering. Default == Medium
type DsrConfidence string

// List of DsrConfidence
const (
	DSRCONFIDENCE_DEFAULT DsrConfidence = "Default"
	DSRCONFIDENCE_LOW DsrConfidence = "Low"
	DSRCONFIDENCE_LOW_MID DsrConfidence = "LowMid"
	DSRCONFIDENCE_MID DsrConfidence = "Mid"
	DSRCONFIDENCE_MID_HIGH DsrConfidence = "MidHigh"
	DSRCONFIDENCE_HIGH DsrConfidence = "High"
	DSRCONFIDENCE_ULTRA DsrConfidence = "Ultra"
	DSRCONFIDENCE_ALL DsrConfidence = "All"
)

// All allowed values of DsrConfidence enum
var AllowedDsrConfidenceEnumValues = []DsrConfidence{
	"Default",
	"Low",
	"LowMid",
	"Mid",
	"MidHigh",
	"High",
	"Ultra",
	"All",
}

func (v *DsrConfidence) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DsrConfidence(value)
	for _, existing := range AllowedDsrConfidenceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DsrConfidence", value)
}

// NewDsrConfidenceFromValue returns a pointer to a valid DsrConfidence
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDsrConfidenceFromValue(v string) (*DsrConfidence, error) {
	ev := DsrConfidence(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DsrConfidence: valid values are %v", v, AllowedDsrConfidenceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DsrConfidence) IsValid() bool {
	for _, existing := range AllowedDsrConfidenceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DsrConfidence value
func (v DsrConfidence) Ptr() *DsrConfidence {
	return &v
}

type NullableDsrConfidence struct {
	value *DsrConfidence
	isSet bool
}

func (v NullableDsrConfidence) Get() *DsrConfidence {
	return v.value
}

func (v *NullableDsrConfidence) Set(val *DsrConfidence) {
	v.value = val
	v.isSet = true
}

func (v NullableDsrConfidence) IsSet() bool {
	return v.isSet
}

func (v *NullableDsrConfidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDsrConfidence(val *DsrConfidence) *NullableDsrConfidence {
	return &NullableDsrConfidence{value: val, isSet: true}
}

func (v NullableDsrConfidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDsrConfidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

