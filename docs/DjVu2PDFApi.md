# \DjVu2PDFApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDjVu2PDF**](DjVu2PDFApi.md#CancelDjVu2PDF) | **Delete** /v5.0/ocr/DjVu2PDF | CancelDjVu2PDF
[**GetDjVu2PDF**](DjVu2PDFApi.md#GetDjVu2PDF) | **Get** /v5.0/ocr/DjVu2PDF | GetDjVu2PDF
[**PostDjVu2PDF**](DjVu2PDFApi.md#PostDjVu2PDF) | **Post** /v5.0/ocr/DjVu2PDF | PostDjVu2PDF



## CancelDjVu2PDF

> CancelDjVu2PDF(ctx).Id(id).Execute()

CancelDjVu2PDF

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
    r, err := apiClient.DjVu2PDFApi.CancelDjVu2PDF(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DjVu2PDFApi.CancelDjVu2PDF``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelDjVu2PDFRequest struct via the builder pattern


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


## GetDjVu2PDF

> OCRResponse GetDjVu2PDF(ctx).Id(id).Execute()

GetDjVu2PDF

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
    resp, r, err := apiClient.DjVu2PDFApi.GetDjVu2PDF(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DjVu2PDFApi.GetDjVu2PDF``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDjVu2PDF`: OCRResponse
    fmt.Fprintf(os.Stdout, "Response from `DjVu2PDFApi.GetDjVu2PDF`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDjVu2PDFRequest struct via the builder pattern


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


## PostDjVu2PDF

> string PostDjVu2PDF(ctx).OCRDjVu2PDFBody(oCRDjVu2PDFBody).Execute()

PostDjVu2PDF

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
    oCRDjVu2PDFBody := *openapiclient.NewOCRDjVu2PDFBody(string(123), *openapiclient.NewOCRSettingsDjVu2PDF()) // OCRDjVu2PDFBody | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.DjVu2PDFApi.PostDjVu2PDF(context.Background()).OCRDjVu2PDFBody(oCRDjVu2PDFBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DjVu2PDFApi.PostDjVu2PDF``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDjVu2PDF`: string
    fmt.Fprintf(os.Stdout, "Response from `DjVu2PDFApi.PostDjVu2PDF`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDjVu2PDFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCRDjVu2PDFBody** | [**OCRDjVu2PDFBody**](OCRDjVu2PDFBody.md) |  | 

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

