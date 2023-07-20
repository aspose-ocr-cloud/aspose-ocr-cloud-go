# \RecognizeImageApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRecognizeImage**](RecognizeImageApi.md#CancelRecognizeImage) | **Delete** /v5.0/ocr/RecognizeImage | CancelRecognizeImage
[**GetRecognizeImage**](RecognizeImageApi.md#GetRecognizeImage) | **Get** /v5.0/ocr/RecognizeImage | GetRecognizeImage
[**PostRecognizeImage**](RecognizeImageApi.md#PostRecognizeImage) | **Post** /v5.0/ocr/RecognizeImage | PostRecognizeImage



## CancelRecognizeImage

> CancelRecognizeImage(ctx).Id(id).Execute()

CancelRecognizeImage

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
    r, err := apiClient.RecognizeImageApi.CancelRecognizeImage(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeImageApi.CancelRecognizeImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelRecognizeImageRequest struct via the builder pattern


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


## GetRecognizeImage

> OCRResponse GetRecognizeImage(ctx).Id(id).Execute()

GetRecognizeImage

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
    resp, r, err := apiClient.RecognizeImageApi.GetRecognizeImage(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeImageApi.GetRecognizeImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecognizeImage`: OCRResponse
    fmt.Fprintf(os.Stdout, "Response from `RecognizeImageApi.GetRecognizeImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecognizeImageRequest struct via the builder pattern


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


## PostRecognizeImage

> string PostRecognizeImage(ctx).OCRRecognizeImageBody(oCRRecognizeImageBody).Execute()

PostRecognizeImage

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
    oCRRecognizeImageBody := *openapiclient.NewOCRRecognizeImageBody(string(123), *openapiclient.NewOCRSettingsRecognizeImage()) // OCRRecognizeImageBody | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.RecognizeImageApi.PostRecognizeImage(context.Background()).OCRRecognizeImageBody(oCRRecognizeImageBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeImageApi.PostRecognizeImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRecognizeImage`: string
    fmt.Fprintf(os.Stdout, "Response from `RecognizeImageApi.PostRecognizeImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRecognizeImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCRRecognizeImageBody** | [**OCRRecognizeImageBody**](OCRRecognizeImageBody.md) |  | 

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

