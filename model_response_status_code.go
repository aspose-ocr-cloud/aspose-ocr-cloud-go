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

// ResponseStatusCode Status code showing the status of the request, operation, and result processing.
type ResponseStatusCode string

// List of ResponseStatusCode
const (
	RESPONSESTATUSCODE_OK ResponseStatusCode = "Ok"
	RESPONSESTATUSCODE_PARTIALLY_NOT_FOUND ResponseStatusCode = "PartiallyNotFound"
	RESPONSESTATUSCODE_NO_ANY_RESULT_DATA ResponseStatusCode = "NoAnyResultData"
	RESPONSESTATUSCODE_RESULT_DATA_LOST ResponseStatusCode = "ResultDataLost"
	RESPONSESTATUSCODE_TASK_NOT_FOUND ResponseStatusCode = "TaskNotFound"
	RESPONSESTATUSCODE_NOT_READY ResponseStatusCode = "NotReady"
	RESPONSESTATUSCODE_ERROR ResponseStatusCode = "Error"
)

// All allowed values of ResponseStatusCode enum
var AllowedResponseStatusCodeEnumValues = []ResponseStatusCode{
	"Ok",
	"PartiallyNotFound",
	"NoAnyResultData",
	"ResultDataLost",
	"TaskNotFound",
	"NotReady",
	"Error",
}

func (v *ResponseStatusCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResponseStatusCode(value)
	for _, existing := range AllowedResponseStatusCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResponseStatusCode", value)
}

// NewResponseStatusCodeFromValue returns a pointer to a valid ResponseStatusCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResponseStatusCodeFromValue(v string) (*ResponseStatusCode, error) {
	ev := ResponseStatusCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResponseStatusCode: valid values are %v", v, AllowedResponseStatusCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResponseStatusCode) IsValid() bool {
	for _, existing := range AllowedResponseStatusCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResponseStatusCode value
func (v ResponseStatusCode) Ptr() *ResponseStatusCode {
	return &v
}

type NullableResponseStatusCode struct {
	value *ResponseStatusCode
	isSet bool
}

func (v NullableResponseStatusCode) Get() *ResponseStatusCode {
	return v.value
}

func (v *NullableResponseStatusCode) Set(val *ResponseStatusCode) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseStatusCode) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseStatusCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseStatusCode(val *ResponseStatusCode) *NullableResponseStatusCode {
	return &NullableResponseStatusCode{value: val, isSet: true}
}

func (v NullableResponseStatusCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseStatusCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

