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

// TTSTaskStatus the model 'TTSTaskStatus'
type TTSTaskStatus string

// List of TTSTaskStatus
const (
	TTSTASKSTATUS_PENDING TTSTaskStatus = "Pending"
	TTSTASKSTATUS_PROCESSING TTSTaskStatus = "Processing"
	TTSTASKSTATUS_COMPLETED TTSTaskStatus = "Completed"
	TTSTASKSTATUS_ERROR TTSTaskStatus = "Error"
	TTSTASKSTATUS_NOT_EXIST TTSTaskStatus = "NotExist"
)

// All allowed values of TTSTaskStatus enum
var AllowedTTSTaskStatusEnumValues = []TTSTaskStatus{
	"Pending",
	"Processing",
	"Completed",
	"Error",
	"NotExist",
}

func (v *TTSTaskStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TTSTaskStatus(value)
	for _, existing := range AllowedTTSTaskStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TTSTaskStatus", value)
}

// NewTTSTaskStatusFromValue returns a pointer to a valid TTSTaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTTSTaskStatusFromValue(v string) (*TTSTaskStatus, error) {
	ev := TTSTaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TTSTaskStatus: valid values are %v", v, AllowedTTSTaskStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TTSTaskStatus) IsValid() bool {
	for _, existing := range AllowedTTSTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TTSTaskStatus value
func (v TTSTaskStatus) Ptr() *TTSTaskStatus {
	return &v
}

type NullableTTSTaskStatus struct {
	value *TTSTaskStatus
	isSet bool
}

func (v NullableTTSTaskStatus) Get() *TTSTaskStatus {
	return v.value
}

func (v *NullableTTSTaskStatus) Set(val *TTSTaskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTTSTaskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTTSTaskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTTSTaskStatus(val *TTSTaskStatus) *NullableTTSTaskStatus {
	return &NullableTTSTaskStatus{value: val, isSet: true}
}

func (v NullableTTSTaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTTSTaskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

