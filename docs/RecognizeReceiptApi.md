# \RecognizeReceiptApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRecognizeReceipt**](RecognizeReceiptApi.md#CancelRecognizeReceipt) | **Delete** /v5.0/ocr/RecognizeReceipt | CancelRecognizeReceipt
[**GetRecognizeReceipt**](RecognizeReceiptApi.md#GetRecognizeReceipt) | **Get** /v5.0/ocr/RecognizeReceipt | GetRecognizeReceipt
[**PostRecognizeReceipt**](RecognizeReceiptApi.md#PostRecognizeReceipt) | **Post** /v5.0/ocr/RecognizeReceipt | PostRecognizeReceipt



## CancelRecognizeReceipt

> CancelRecognizeReceipt(ctx).Id(id).Execute()

CancelRecognizeReceipt

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
    r, err := apiClient.RecognizeReceiptApi.CancelRecognizeReceipt(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeReceiptApi.CancelRecognizeReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelRecognizeReceiptRequest struct via the builder pattern


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


## GetRecognizeReceipt

> OCRResponse GetRecognizeReceipt(ctx).Id(id).Execute()

GetRecognizeReceipt

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
    resp, r, err := apiClient.RecognizeReceiptApi.GetRecognizeReceipt(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeReceiptApi.GetRecognizeReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecognizeReceipt`: OCRResponse
    fmt.Fprintf(os.Stdout, "Response from `RecognizeReceiptApi.GetRecognizeReceipt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecognizeReceiptRequest struct via the builder pattern


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


## PostRecognizeReceipt

> string PostRecognizeReceipt(ctx).OCRRecognizeReceiptBody(oCRRecognizeReceiptBody).Execute()

PostRecognizeReceipt

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
    oCRRecognizeReceiptBody := *openapiclient.NewOCRRecognizeReceiptBody(string(123), *openapiclient.NewOCRSettingsRecognizeReceipt()) // OCRRecognizeReceiptBody | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.RecognizeReceiptApi.PostRecognizeReceipt(context.Background()).OCRRecognizeReceiptBody(oCRRecognizeReceiptBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeReceiptApi.PostRecognizeReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRecognizeReceipt`: string
    fmt.Fprintf(os.Stdout, "Response from `RecognizeReceiptApi.PostRecognizeReceipt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRecognizeReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCRRecognizeReceiptBody** | [**OCRRecognizeReceiptBody**](OCRRecognizeReceiptBody.md) |  | 

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

