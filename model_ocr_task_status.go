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

// OCRTaskStatus Task status
type OCRTaskStatus string

// List of OCRTaskStatus
const (
	OCRTASKSTATUS_PENDING OCRTaskStatus = "Pending"
	OCRTASKSTATUS_PROCESSING OCRTaskStatus = "Processing"
	OCRTASKSTATUS_COMPLETED OCRTaskStatus = "Completed"
	OCRTASKSTATUS_ERROR OCRTaskStatus = "Error"
	OCRTASKSTATUS_NOT_EXIST OCRTaskStatus = "NotExist"
)

// All allowed values of OCRTaskStatus enum
var AllowedOCRTaskStatusEnumValues = []OCRTaskStatus{
	"Pending",
	"Processing",
	"Completed",
	"Error",
	"NotExist",
}

func (v *OCRTaskStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OCRTaskStatus(value)
	for _, existing := range AllowedOCRTaskStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OCRTaskStatus", value)
}

// NewOCRTaskStatusFromValue returns a pointer to a valid OCRTaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOCRTaskStatusFromValue(v string) (*OCRTaskStatus, error) {
	ev := OCRTaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OCRTaskStatus: valid values are %v", v, AllowedOCRTaskStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OCRTaskStatus) IsValid() bool {
	for _, existing := range AllowedOCRTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OCRTaskStatus value
func (v OCRTaskStatus) Ptr() *OCRTaskStatus {
	return &v
}

type NullableOCRTaskStatus struct {
	value *OCRTaskStatus
	isSet bool
}

func (v NullableOCRTaskStatus) Get() *OCRTaskStatus {
	return v.value
}

func (v *NullableOCRTaskStatus) Set(val *OCRTaskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRTaskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRTaskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRTaskStatus(val *OCRTaskStatus) *NullableOCRTaskStatus {
	return &NullableOCRTaskStatus{value: val, isSet: true}
}

func (v NullableOCRTaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRTaskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

