# \RecognizePdfApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRecognizePdf**](RecognizePdfApi.md#CancelRecognizePdf) | **Delete** /v5.0/ocr/RecognizePdf | CancelRecognizePdf
[**GetRecognizePdf**](RecognizePdfApi.md#GetRecognizePdf) | **Get** /v5.0/ocr/RecognizePdf | GetRecognizePdf
[**PostRecognizePdf**](RecognizePdfApi.md#PostRecognizePdf) | **Post** /v5.0/ocr/RecognizePdf | PostRecognizePdf



## CancelRecognizePdf

> CancelRecognizePdf(ctx).Id(id).Execute()

CancelRecognizePdf

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
    r, err := apiClient.RecognizePdfApi.CancelRecognizePdf(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizePdfApi.CancelRecognizePdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelRecognizePdfRequest struct via the builder pattern


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


## GetRecognizePdf

> OCRResponse GetRecognizePdf(ctx).Id(id).Execute()

GetRecognizePdf

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
    resp, r, err := apiClient.RecognizePdfApi.GetRecognizePdf(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizePdfApi.GetRecognizePdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecognizePdf`: OCRResponse
    fmt.Fprintf(os.Stdout, "Response from `RecognizePdfApi.GetRecognizePdf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecognizePdfRequest struct via the builder pattern


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


## PostRecognizePdf

> string PostRecognizePdf(ctx).OCRRecognizePdfBody(oCRRecognizePdfBody).Execute()

PostRecognizePdf

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
    oCRRecognizePdfBody := *openapiclient.NewOCRRecognizePdfBody(string(123), *openapiclient.NewOCRSettingsRecognizePdf()) // OCRRecognizePdfBody | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.RecognizePdfApi.PostRecognizePdf(context.Background()).OCRRecognizePdfBody(oCRRecognizePdfBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizePdfApi.PostRecognizePdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRecognizePdf`: string
    fmt.Fprintf(os.Stdout, "Response from `RecognizePdfApi.PostRecognizePdf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRecognizePdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCRRecognizePdfBody** | [**OCRRecognizePdfBody**](OCRRecognizePdfBody.md) |  | 

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

