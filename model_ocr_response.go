/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the OCRResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OCRResponse{}

// OCRResponse Response with Recognition result for specific task ID
type OCRResponse struct {
	// The specific Task ID that result was made for
	Id NullableString `json:"id,omitempty"`
	ResponseStatusCode *ResponseStatusCode `json:"responseStatusCode,omitempty"`
	TaskStatus *OCRTaskStatus `json:"taskStatus,omitempty"`
	// List of results - Especially Text or PDF files
	Results []OCRResult `json:"results,omitempty"`
	Error NullableOCRError `json:"error,omitempty"`
}

// NewOCRResponse instantiates a new OCRResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOCRResponse() *OCRResponse {
	this := OCRResponse{}
	return &this
}

// NewOCRResponseWithDefaults instantiates a new OCRResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOCRResponseWithDefaults() *OCRResponse {
	this := OCRResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OCRResponse) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OCRResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *OCRResponse) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *OCRResponse) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *OCRResponse) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *OCRResponse) UnsetId() {
	o.Id.Unset()
}

// GetResponseStatusCode returns the ResponseStatusCode field value if set, zero value otherwise.
func (o *OCRResponse) GetResponseStatusCode() ResponseStatusCode {
	if o == nil || IsNil(o.ResponseStatusCode) {
		var ret ResponseStatusCode
		return ret
	}
	return *o.ResponseStatusCode
}

// GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRResponse) GetResponseStatusCodeOk() (*ResponseStatusCode, bool) {
	if o == nil || IsNil(o.ResponseStatusCode) {
		return nil, false
	}
	return o.ResponseStatusCode, true
}

// HasResponseStatusCode returns a boolean if a field has been set.
func (o *OCRResponse) HasResponseStatusCode() bool {
	if o != nil && !IsNil(o.ResponseStatusCode) {
		return true
	}

	return false
}

// SetResponseStatusCode gets a reference to the given ResponseStatusCode and assigns it to the ResponseStatusCode field.
func (o *OCRResponse) SetResponseStatusCode(v ResponseStatusCode) {
	o.ResponseStatusCode = &v
}

// GetTaskStatus returns the TaskStatus field value if set, zero value otherwise.
func (o *OCRResponse) GetTaskStatus() OCRTaskStatus {
	if o == nil || IsNil(o.TaskStatus) {
		var ret OCRTaskStatus
		return ret
	}
	return *o.TaskStatus
}

// GetTaskStatusOk returns a tuple with the TaskStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OCRResponse) GetTaskStatusOk() (*OCRTaskStatus, bool) {
	if o == nil || IsNil(o.TaskStatus) {
		return nil, false
	}
	return o.TaskStatus, true
}

// HasTaskStatus returns a boolean if a field has been set.
func (o *OCRResponse) HasTaskStatus() bool {
	if o != nil && !IsNil(o.TaskStatus) {
		return true
	}

	return false
}

// SetTaskStatus gets a reference to the given OCRTaskStatus and assigns it to the TaskStatus field.
func (o *OCRResponse) SetTaskStatus(v OCRTaskStatus) {
	o.TaskStatus = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OCRResponse) GetResults() []OCRResult {
	if o == nil {
		var ret []OCRResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OCRResponse) GetResultsOk() ([]OCRResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *OCRResponse) HasResults() bool {
	if o != nil && IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []OCRResult and assigns it to the Results field.
func (o *OCRResponse) SetResults(v []OCRResult) {
	o.Results = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OCRResponse) GetError() OCRError {
	if o == nil || IsNil(o.Error.Get()) {
		var ret OCRError
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OCRResponse) GetErrorOk() (*OCRError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *OCRResponse) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableOCRError and assigns it to the Error field.
func (o *OCRResponse) SetError(v OCRError) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *OCRResponse) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *OCRResponse) UnsetError() {
	o.Error.Unset()
}

func (o OCRResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OCRResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if !IsNil(o.ResponseStatusCode) {
		toSerialize["responseStatusCode"] = o.ResponseStatusCode
	}
	if !IsNil(o.TaskStatus) {
		toSerialize["taskStatus"] = o.TaskStatus
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	return toSerialize, nil
}

type NullableOCRResponse struct {
	value *OCRResponse
	isSet bool
}

func (v NullableOCRResponse) Get() *OCRResponse {
	return v.value
}

func (v *NullableOCRResponse) Set(val *OCRResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOCRResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOCRResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOCRResponse(val *OCRResponse) *NullableOCRResponse {
	return &NullableOCRResponse{value: val, isSet: true}
}

func (v NullableOCRResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOCRResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


