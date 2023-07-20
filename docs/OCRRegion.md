# OCRRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rect** | Pointer to [**OCRRect**](OCRRect.md) |  | [optional] 
**Order** | Pointer to **int32** | The serial number of the region for the formation of the text | [optional] 

## Methods

### NewOCRRegion

`func NewOCRRegion() *OCRRegion`

NewOCRRegion instantiates a new OCRRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRRegionWithDefaults

`func NewOCRRegionWithDefaults() *OCRRegion`

NewOCRRegionWithDefaults instantiates a new OCRRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRect

`func (o *OCRRegion) GetRect() OCRRect`

GetRect returns the Rect field if non-nil, zero value otherwise.

### GetRectOk

`func (o *OCRRegion) GetRectOk() (*OCRRect, bool)`

GetRectOk returns a tuple with the Rect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRect

`func (o *OCRRegion) SetRect(v OCRRect)`

SetRect sets Rect field to given value.

### HasRect

`func (o *OCRRegion) HasRect() bool`

HasRect returns a boolean if a field has been set.

### GetOrder

`func (o *OCRRegion) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OCRRegion) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OCRRegion) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *OCRRegion) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


