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

// ResultTypeTable Result document type for Table OCR process
type ResultTypeTable string

// List of ResultTypeTable
const (
	RESULTTYPETABLE_TEXT ResultTypeTable = "Text"
	RESULTTYPETABLE_EXCEL ResultTypeTable = "Excel"
	RESULTTYPETABLE_CSV ResultTypeTable = "Csv"
	RESULTTYPETABLE_CSV_AND_EXCEL ResultTypeTable = "CsvAndExcel"
)

// All allowed values of ResultTypeTable enum
var AllowedResultTypeTableEnumValues = []ResultTypeTable{
	"Text",
	"Excel",
	"Csv",
	"CsvAndExcel",
}

func (v *ResultTypeTable) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResultTypeTable(value)
	for _, existing := range AllowedResultTypeTableEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResultTypeTable", value)
}

// NewResultTypeTableFromValue returns a pointer to a valid ResultTypeTable
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResultTypeTableFromValue(v string) (*ResultTypeTable, error) {
	ev := ResultTypeTable(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResultTypeTable: valid values are %v", v, AllowedResultTypeTableEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResultTypeTable) IsValid() bool {
	for _, existing := range AllowedResultTypeTableEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResultTypeTable value
func (v ResultTypeTable) Ptr() *ResultTypeTable {
	return &v
}

type NullableResultTypeTable struct {
	value *ResultTypeTable
	isSet bool
}

func (v NullableResultTypeTable) Get() *ResultTypeTable {
	return v.value
}

func (v *NullableResultTypeTable) Set(val *ResultTypeTable) {
	v.value = val
	v.isSet = true
}

func (v NullableResultTypeTable) IsSet() bool {
	return v.isSet
}

func (v *NullableResultTypeTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultTypeTable(val *ResultTypeTable) *NullableResultTypeTable {
	return &NullableResultTypeTable{value: val, isSet: true}
}

func (v NullableResultTypeTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultTypeTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

