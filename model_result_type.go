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

// ResultType Result document type for OCR process
type ResultType string

// List of ResultType
const (
	RESULTTYPE_TEXT ResultType = "Text"
	RESULTTYPE_PDF ResultType = "Pdf"
	RESULTTYPE_TEXT_AND_PDF ResultType = "TextAndPdf"
	RESULTTYPE_HOCR ResultType = "Hocr"
	RESULTTYPE_TEXT_AND_HOCR ResultType = "TextAndHocr"
	RESULTTYPE_PDF_AND_HOCR ResultType = "PdfAndHocr"
	RESULTTYPE_TEXT_AND_PDF_AND_HOCR ResultType = "TextAndPdfAndHocr"
	RESULTTYPE_IMAGE_PNG ResultType = "ImagePNG"
)

// All allowed values of ResultType enum
var AllowedResultTypeEnumValues = []ResultType{
	"Text",
	"Pdf",
	"TextAndPdf",
	"Hocr",
	"TextAndHocr",
	"PdfAndHocr",
	"TextAndPdfAndHocr",
	"ImagePNG",
}

func (v *ResultType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResultType(value)
	for _, existing := range AllowedResultTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResultType", value)
}

// NewResultTypeFromValue returns a pointer to a valid ResultType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResultTypeFromValue(v string) (*ResultType, error) {
	ev := ResultType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResultType: valid values are %v", v, AllowedResultTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResultType) IsValid() bool {
	for _, existing := range AllowedResultTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResultType value
func (v ResultType) Ptr() *ResultType {
	return &v
}

type NullableResultType struct {
	value *ResultType
	isSet bool
}

func (v NullableResultType) Get() *ResultType {
	return v.value
}

func (v *NullableResultType) Set(val *ResultType) {
	v.value = val
	v.isSet = true
}

func (v NullableResultType) IsSet() bool {
	return v.isSet
}

func (v *NullableResultType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultType(val *ResultType) *NullableResultType {
	return &NullableResultType{value: val, isSet: true}
}

func (v NullableResultType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

