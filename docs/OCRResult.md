# OCRResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | File data type (extension) | [optional] 
**Data** | Pointer to **NullableString** | File binary data | [optional] 

## Methods

### NewOCRResult

`func NewOCRResult() *OCRResult`

NewOCRResult instantiates a new OCRResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRResultWithDefaults

`func NewOCRResultWithDefaults() *OCRResult`

NewOCRResultWithDefaults instantiates a new OCRResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OCRResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OCRResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OCRResult) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OCRResult) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *OCRResult) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OCRResult) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetData

`func (o *OCRResult) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OCRResult) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OCRResult) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *OCRResult) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *OCRResult) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *OCRResult) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


