# \RecognizeLabelApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRecognizeLabel**](RecognizeLabelApi.md#CancelRecognizeLabel) | **Delete** /v5.0/ocr/RecognizeLabel | CancelRecognizeLabel
[**GetRecognizeLabel**](RecognizeLabelApi.md#GetRecognizeLabel) | **Get** /v5.0/ocr/RecognizeLabel | GetRecognizeLabel
[**PostRecognizeLabel**](RecognizeLabelApi.md#PostRecognizeLabel) | **Post** /v5.0/ocr/RecognizeLabel | PostRecognizeLabel



## CancelRecognizeLabel

> CancelRecognizeLabel(ctx).Id(id).Execute()

CancelRecognizeLabel

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    id := "id_example" // string | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    r, err := apiClient.RecognizeLabelApi.CancelRecognizeLabel(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeLabelApi.CancelRecognizeLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelRecognizeLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecognizeLabel

> OCRResponse GetRecognizeLabel(ctx).Id(id).Execute()

GetRecognizeLabel

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    id := "id_example" // string | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.RecognizeLabelApi.GetRecognizeLabel(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeLabelApi.GetRecognizeLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecognizeLabel`: OCRResponse
    fmt.Fprintf(os.Stdout, "Response from `RecognizeLabelApi.GetRecognizeLabel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecognizeLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

### Return type

[**OCRResponse**](OCRResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRecognizeLabel

> string PostRecognizeLabel(ctx).OCRRecognizeLabelBody(oCRRecognizeLabelBody).Execute()

PostRecognizeLabel

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {
    oCRRecognizeLabelBody := *openapiclient.NewOCRRecognizeLabelBody(string(123), *openapiclient.NewOCRSettingsRecognizeLabel()) // OCRRecognizeLabelBody | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.RecognizeLabelApi.PostRecognizeLabel(context.Background()).OCRRecognizeLabelBody(oCRRecognizeLabelBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeLabelApi.PostRecognizeLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRecognizeLabel`: string
    fmt.Fprintf(os.Stdout, "Response from `RecognizeLabelApi.PostRecognizeLabel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRecognizeLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCRRecognizeLabelBody** | [**OCRRecognizeLabelBody**](OCRRecognizeLabelBody.md) |  | 

### Return type

**string**

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

