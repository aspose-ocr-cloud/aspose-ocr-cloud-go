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

// DsrMode Option that sets the recognition result type or combination of some types: Text, Searchable PDF, HOCR
type DsrMode string

// List of DsrMode
const (
	DSRMODE_REGIONS DsrMode = "Regions"
	DSRMODE_DSR_NO_FILTER DsrMode = "DsrNoFilter"
	DSRMODE_DSR_AND_FILTER DsrMode = "DsrAndFilter"
	DSRMODE_NO_DSR_NO_FILTER DsrMode = "NoDsrNoFilter"
	DSRMODE_TEXT_DETECTOR DsrMode = "TextDetector"
	DSRMODE_DSR_PLUS_DETECTOR DsrMode = "DsrPlusDetector"
	DSRMODE_POLYGONAL_TEXT_DETECTOR DsrMode = "PolygonalTextDetector"
)

// All allowed values of DsrMode enum
var AllowedDsrModeEnumValues = []DsrMode{
	"Regions",
	"DsrNoFilter",
	"DsrAndFilter",
	"NoDsrNoFilter",
	"TextDetector",
	"DsrPlusDetector",
	"PolygonalTextDetector",
}

func (v *DsrMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DsrMode(value)
	for _, existing := range AllowedDsrModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DsrMode", value)
}

// NewDsrModeFromValue returns a pointer to a valid DsrMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDsrModeFromValue(v string) (*DsrMode, error) {
	ev := DsrMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DsrMode: valid values are %v", v, AllowedDsrModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DsrMode) IsValid() bool {
	for _, existing := range AllowedDsrModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DsrMode value
func (v DsrMode) Ptr() *DsrMode {
	return &v
}

type NullableDsrMode struct {
	value *DsrMode
	isSet bool
}

func (v NullableDsrMode) Get() *DsrMode {
	return v.value
}

func (v *NullableDsrMode) Set(val *DsrMode) {
	v.value = val
	v.isSet = true
}

func (v NullableDsrMode) IsSet() bool {
	return v.isSet
}

func (v *NullableDsrMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDsrMode(val *DsrMode) *NullableDsrMode {
	return &NullableDsrMode{value: val, isSet: true}
}

func (v NullableDsrMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDsrMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

