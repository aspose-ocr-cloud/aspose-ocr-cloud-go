# \UpscaleImageApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelUpscaleImage**](UpscaleImageApi.md#CancelUpscaleImage) | **Delete** /v5.0/ocr/UpscaleImage | CancelUpscaleImage
[**GetUpscaleImage**](UpscaleImageApi.md#GetUpscaleImage) | **Get** /v5.0/ocr/UpscaleImage | GetUpscaleImage
[**PostUpscaleImage**](UpscaleImageApi.md#PostUpscaleImage) | **Post** /v5.0/ocr/UpscaleImage | PostUpscaleImage



## CancelUpscaleImage

> CancelUpscaleImage(ctx).Id(id).Execute()

CancelUpscaleImage

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
    r, err := apiClient.UpscaleImageApi.CancelUpscaleImage(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpscaleImageApi.CancelUpscaleImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelUpscaleImageRequest struct via the builder pattern


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


## GetUpscaleImage

> OCRResponse GetUpscaleImage(ctx).Id(id).Execute()

GetUpscaleImage

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
    resp, r, err := apiClient.UpscaleImageApi.GetUpscaleImage(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpscaleImageApi.GetUpscaleImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUpscaleImage`: OCRResponse
    fmt.Fprintf(os.Stdout, "Response from `UpscaleImageApi.GetUpscaleImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUpscaleImageRequest struct via the builder pattern


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


## PostUpscaleImage

> string PostUpscaleImage(ctx).OCRUpscaleImageBody(oCRUpscaleImageBody).Execute()

PostUpscaleImage

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
    oCRUpscaleImageBody := *openapiclient.NewOCRUpscaleImageBody(string(123)) // OCRUpscaleImageBody | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.UpscaleImageApi.PostUpscaleImage(context.Background()).OCRUpscaleImageBody(oCRUpscaleImageBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpscaleImageApi.PostUpscaleImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUpscaleImage`: string
    fmt.Fprintf(os.Stdout, "Response from `UpscaleImageApi.PostUpscaleImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUpscaleImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCRUpscaleImageBody** | [**OCRUpscaleImageBody**](OCRUpscaleImageBody.md) |  | 

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

