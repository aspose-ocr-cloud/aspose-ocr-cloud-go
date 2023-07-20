# \BinarizeImageApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelBinarizeImage**](BinarizeImageApi.md#CancelBinarizeImage) | **Delete** /v5.0/ocr/BinarizeImage | CancelBinarizeImage
[**GetBinarizeImage**](BinarizeImageApi.md#GetBinarizeImage) | **Get** /v5.0/ocr/BinarizeImage | GetBinarizeImage
[**PostBinarizeImage**](BinarizeImageApi.md#PostBinarizeImage) | **Post** /v5.0/ocr/BinarizeImage | PostBinarizeImage



## CancelBinarizeImage

> CancelBinarizeImage(ctx).Id(id).Execute()

CancelBinarizeImage

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
    r, err := apiClient.BinarizeImageApi.CancelBinarizeImage(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BinarizeImageApi.CancelBinarizeImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelBinarizeImageRequest struct via the builder pattern


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


## GetBinarizeImage

> OCRResponse GetBinarizeImage(ctx).Id(id).Execute()

GetBinarizeImage

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
    resp, r, err := apiClient.BinarizeImageApi.GetBinarizeImage(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BinarizeImageApi.GetBinarizeImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBinarizeImage`: OCRResponse
    fmt.Fprintf(os.Stdout, "Response from `BinarizeImageApi.GetBinarizeImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBinarizeImageRequest struct via the builder pattern


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


## PostBinarizeImage

> string PostBinarizeImage(ctx).OCRBinarizeImageBody(oCRBinarizeImageBody).Execute()

PostBinarizeImage

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
    oCRBinarizeImageBody := *openapiclient.NewOCRBinarizeImageBody(string(123)) // OCRBinarizeImageBody | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.BinarizeImageApi.PostBinarizeImage(context.Background()).OCRBinarizeImageBody(oCRBinarizeImageBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BinarizeImageApi.PostBinarizeImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostBinarizeImage`: string
    fmt.Fprintf(os.Stdout, "Response from `BinarizeImageApi.PostBinarizeImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBinarizeImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCRBinarizeImageBody** | [**OCRBinarizeImageBody**](OCRBinarizeImageBody.md) |  | 

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

