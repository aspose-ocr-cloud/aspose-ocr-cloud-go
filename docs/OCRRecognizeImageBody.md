# OCRRecognizeImageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** | Gets or Sets Image | 
**Settings** | [**OCRSettingsRecognizeImage**](OCRSettingsRecognizeImage.md) |  | 

## Methods

### NewOCRRecognizeImageBody

`func NewOCRRecognizeImageBody(image string, settings OCRSettingsRecognizeImage, ) *OCRRecognizeImageBody`

NewOCRRecognizeImageBody instantiates a new OCRRecognizeImageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRRecognizeImageBodyWithDefaults

`func NewOCRRecognizeImageBodyWithDefaults() *OCRRecognizeImageBody`

NewOCRRecognizeImageBodyWithDefaults instantiates a new OCRRecognizeImageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *OCRRecognizeImageBody) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *OCRRecognizeImageBody) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *OCRRecognizeImageBody) SetImage(v string)`

SetImage sets Image field to given value.


### GetSettings

`func (o *OCRRecognizeImageBody) GetSettings() OCRSettingsRecognizeImage`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *OCRRecognizeImageBody) GetSettingsOk() (*OCRSettingsRecognizeImage, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *OCRRecognizeImageBody) SetSettings(v OCRSettingsRecognizeImage)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


