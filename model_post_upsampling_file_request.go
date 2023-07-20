/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"encoding/json"
	"os"
)

// checks if the PostUpsamplingFileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostUpsamplingFileRequest{}

// PostUpsamplingFileRequest struct for PostUpsamplingFileRequest
type PostUpsamplingFileRequest struct {
	File **os.File `json:"file,omitempty"`
}

// NewPostUpsamplingFileRequest instantiates a new PostUpsamplingFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostUpsamplingFileRequest() *PostUpsamplingFileRequest {
	this := PostUpsamplingFileRequest{}
	return &this
}

// NewPostUpsamplingFileRequestWithDefaults instantiates a new PostUpsamplingFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostUpsamplingFileRequestWithDefaults() *PostUpsamplingFileRequest {
	this := PostUpsamplingFileRequest{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *PostUpsamplingFileRequest) GetFile() *os.File {
	if o == nil || IsNil(o.File) {
		var ret *os.File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostUpsamplingFileRequest) GetFileOk() (**os.File, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *PostUpsamplingFileRequest) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *PostUpsamplingFileRequest) SetFile(v *os.File) {
	o.File = &v
}

func (o PostUpsamplingFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostUpsamplingFileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	return toSerialize, nil
}

type NullablePostUpsamplingFileRequest struct {
	value *PostUpsamplingFileRequest
	isSet bool
}

func (v NullablePostUpsamplingFileRequest) Get() *PostUpsamplingFileRequest {
	return v.value
}

func (v *NullablePostUpsamplingFileRequest) Set(val *PostUpsamplingFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostUpsamplingFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostUpsamplingFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostUpsamplingFileRequest(val *PostUpsamplingFileRequest) *NullablePostUpsamplingFileRequest {
	return &NullablePostUpsamplingFileRequest{value: val, isSet: true}
}

func (v NullablePostUpsamplingFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostUpsamplingFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


