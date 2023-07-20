# \DeskewImageApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDeskewImage**](DeskewImageApi.md#CancelDeskewImage) | **Delete** /v5.0/ocr/DeskewImage | CancelDeskewImage
[**GetDeskewImage**](DeskewImageApi.md#GetDeskewImage) | **Get** /v5.0/ocr/DeskewImage | GetDeskewImage
[**PostDeskewImage**](DeskewImageApi.md#PostDeskewImage) | **Post** /v5.0/ocr/DeskewImage | PostDeskewImage



## CancelDeskewImage

> CancelDeskewImage(ctx).Id(id).Execute()

CancelDeskewImage

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
    r, err := apiClient.DeskewImageApi.CancelDeskewImage(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeskewImageApi.CancelDeskewImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelDeskewImageRequest struct via the builder pattern


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


## GetDeskewImage

> OCRResponse GetDeskewImage(ctx).Id(id).Execute()

GetDeskewImage

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
    resp, r, err := apiClient.DeskewImageApi.GetDeskewImage(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeskewImageApi.GetDeskewImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeskewImage`: OCRResponse
    fmt.Fprintf(os.Stdout, "Response from `DeskewImageApi.GetDeskewImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeskewImageRequest struct via the builder pattern


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


## PostDeskewImage

> string PostDeskewImage(ctx).OCRDeskewImageBody(oCRDeskewImageBody).Execute()

PostDeskewImage

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
    oCRDeskewImageBody := *openapiclient.NewOCRDeskewImageBody(string(123)) // OCRDeskewImageBody | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.DeskewImageApi.PostDeskewImage(context.Background()).OCRDeskewImageBody(oCRDeskewImageBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeskewImageApi.PostDeskewImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDeskewImage`: string
    fmt.Fprintf(os.Stdout, "Response from `DeskewImageApi.PostDeskewImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDeskewImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCRDeskewImageBody** | [**OCRDeskewImageBody**](OCRDeskewImageBody.md) |  | 

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

