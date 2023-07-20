# \DewarpImageApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDewarpImage**](DewarpImageApi.md#CancelDewarpImage) | **Delete** /v5.0/ocr/DewarpImage | CancelDewarpImage
[**GetDewarpImage**](DewarpImageApi.md#GetDewarpImage) | **Get** /v5.0/ocr/DewarpImage | GetDewarpImage
[**PostDewarpImage**](DewarpImageApi.md#PostDewarpImage) | **Post** /v5.0/ocr/DewarpImage | PostDewarpImage



## CancelDewarpImage

> CancelDewarpImage(ctx).Id(id).Execute()

CancelDewarpImage

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
    r, err := apiClient.DewarpImageApi.CancelDewarpImage(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DewarpImageApi.CancelDewarpImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelDewarpImageRequest struct via the builder pattern


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


## GetDewarpImage

> OCRResponse GetDewarpImage(ctx).Id(id).Execute()

GetDewarpImage

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
    resp, r, err := apiClient.DewarpImageApi.GetDewarpImage(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DewarpImageApi.GetDewarpImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDewarpImage`: OCRResponse
    fmt.Fprintf(os.Stdout, "Response from `DewarpImageApi.GetDewarpImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDewarpImageRequest struct via the builder pattern


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


## PostDewarpImage

> string PostDewarpImage(ctx).OCRDewarpImageBody(oCRDewarpImageBody).Execute()

PostDewarpImage

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
    oCRDewarpImageBody := *openapiclient.NewOCRDewarpImageBody(string(123)) // OCRDewarpImageBody | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.DewarpImageApi.PostDewarpImage(context.Background()).OCRDewarpImageBody(oCRDewarpImageBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DewarpImageApi.PostDewarpImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDewarpImage`: string
    fmt.Fprintf(os.Stdout, "Response from `DewarpImageApi.PostDewarpImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDewarpImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCRDewarpImageBody** | [**OCRDewarpImageBody**](OCRDewarpImageBody.md) |  | 

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

