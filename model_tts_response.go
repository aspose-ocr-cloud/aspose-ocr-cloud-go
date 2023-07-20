/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
)

// checks if the TTSResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TTSResponse{}

// TTSResponse struct for TTSResponse
type TTSResponse struct {
	Id NullableString `json:"id,omitempty"`
	ResponseStatusCode *ResponseStatusCode `json:"responseStatusCode,omitempty"`
	TaskStatus *TTSTaskStatus `json:"taskStatus,omitempty"`
	Results []TTSResult `json:"results,omitempty"`
	Error NullableTTSError `json:"error,omitempty"`
}

// NewTTSResponse instantiates a new TTSResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTTSResponse() *TTSResponse {
	this := TTSResponse{}
	return &this
}

// NewTTSResponseWithDefaults instantiates a new TTSResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTTSResponseWithDefaults() *TTSResponse {
	this := TTSResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TTSResponse) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TTSResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *TTSResponse) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *TTSResponse) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *TTSResponse) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *TTSResponse) UnsetId() {
	o.Id.Unset()
}

// GetResponseStatusCode returns the ResponseStatusCode field value if set, zero value otherwise.
func (o *TTSResponse) GetResponseStatusCode() ResponseStatusCode {
	if o == nil || IsNil(o.ResponseStatusCode) {
		var ret ResponseStatusCode
		return ret
	}
	return *o.ResponseStatusCode
}

// GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TTSResponse) GetResponseStatusCodeOk() (*ResponseStatusCode, bool) {
	if o == nil || IsNil(o.ResponseStatusCode) {
		return nil, false
	}
	return o.ResponseStatusCode, true
}

// HasResponseStatusCode returns a boolean if a field has been set.
func (o *TTSResponse) HasResponseStatusCode() bool {
	if o != nil && !IsNil(o.ResponseStatusCode) {
		return true
	}

	return false
}

// SetResponseStatusCode gets a reference to the given ResponseStatusCode and assigns it to the ResponseStatusCode field.
func (o *TTSResponse) SetResponseStatusCode(v ResponseStatusCode) {
	o.ResponseStatusCode = &v
}

// GetTaskStatus returns the TaskStatus field value if set, zero value otherwise.
func (o *TTSResponse) GetTaskStatus() TTSTaskStatus {
	if o == nil || IsNil(o.TaskStatus) {
		var ret TTSTaskStatus
		return ret
	}
	return *o.TaskStatus
}

// GetTaskStatusOk returns a tuple with the TaskStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TTSResponse) GetTaskStatusOk() (*TTSTaskStatus, bool) {
	if o == nil || IsNil(o.TaskStatus) {
		return nil, false
	}
	return o.TaskStatus, true
}

// HasTaskStatus returns a boolean if a field has been set.
func (o *TTSResponse) HasTaskStatus() bool {
	if o != nil && !IsNil(o.TaskStatus) {
		return true
	}

	return false
}

// SetTaskStatus gets a reference to the given TTSTaskStatus and assigns it to the TaskStatus field.
func (o *TTSResponse) SetTaskStatus(v TTSTaskStatus) {
	o.TaskStatus = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TTSResponse) GetResults() []TTSResult {
	if o == nil {
		var ret []TTSResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TTSResponse) GetResultsOk() ([]TTSResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *TTSResponse) HasResults() bool {
	if o != nil && IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []TTSResult and assigns it to the Results field.
func (o *TTSResponse) SetResults(v []TTSResult) {
	o.Results = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TTSResponse) GetError() TTSError {
	if o == nil || IsNil(o.Error.Get()) {
		var ret TTSError
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TTSResponse) GetErrorOk() (*TTSError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *TTSResponse) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableTTSError and assigns it to the Error field.
func (o *TTSResponse) SetError(v TTSError) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *TTSResponse) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *TTSResponse) UnsetError() {
	o.Error.Unset()
}

func (o TTSResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TTSResponse) ToMap() (map[string]interface{}, error) {
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

type NullableTTSResponse struct {
	value *TTSResponse
	isSet bool
}

func (v NullableTTSResponse) Get() *TTSResponse {
	return v.value
}

func (v *NullableTTSResponse) Set(val *TTSResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTTSResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTTSResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTTSResponse(val *TTSResponse) *NullableTTSResponse {
	return &NullableTTSResponse{value: val, isSet: true}
}

func (v NullableTTSResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTTSResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


