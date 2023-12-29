# \RecognizeAndParseInvoiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRecognizeAndParseInvoice**](RecognizeAndParseInvoiceApi.md#CancelRecognizeAndParseInvoice) | **Delete** /v5.0/ocr/RecognizeAndParseInvoice | CancelRecognizeAndParseInvoice
[**GetRecognizeAndParseInvoice**](RecognizeAndParseInvoiceApi.md#GetRecognizeAndParseInvoice) | **Get** /v5.0/ocr/RecognizeAndParseInvoice | GetRecognizeAndParseInvoice
[**PostRecognizeAndParseInvoice**](RecognizeAndParseInvoiceApi.md#PostRecognizeAndParseInvoice) | **Post** /v5.0/ocr/RecognizeAndParseInvoice | PostRecognizeAndParseInvoice



## CancelRecognizeAndParseInvoice

> CancelRecognizeAndParseInvoice(ctx).Id(id).Execute()

CancelRecognizeAndParseInvoice

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
    r, err := apiClient.RecognizeAndParseInvoiceApi.CancelRecognizeAndParseInvoice(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeAndParseInvoiceApi.CancelRecognizeAndParseInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelRecognizeAndParseInvoiceRequest struct via the builder pattern


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


## GetRecognizeAndParseInvoice

> OCRResponse GetRecognizeAndParseInvoice(ctx).Id(id).Execute()

GetRecognizeAndParseInvoice

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
    resp, r, err := apiClient.RecognizeAndParseInvoiceApi.GetRecognizeAndParseInvoice(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeAndParseInvoiceApi.GetRecognizeAndParseInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecognizeAndParseInvoice`: OCRResponse
    fmt.Fprintf(os.Stdout, "Response from `RecognizeAndParseInvoiceApi.GetRecognizeAndParseInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecognizeAndParseInvoiceRequest struct via the builder pattern


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


## PostRecognizeAndParseInvoice

> string PostRecognizeAndParseInvoice(ctx).OCRRecognizeAndParseInvoiceBody(oCRRecognizeAndParseInvoiceBody).Execute()

PostRecognizeAndParseInvoice

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
    oCRRecognizeAndParseInvoiceBody := *openapiclient.NewOCRRecognizeAndParseInvoiceBody(string(123), *openapiclient.NewOCRSettingsRecognizeAndParseInvoice()) // OCRRecognizeAndParseInvoiceBody | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.RecognizeAndParseInvoiceApi.PostRecognizeAndParseInvoice(context.Background()).OCRRecognizeAndParseInvoiceBody(oCRRecognizeAndParseInvoiceBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeAndParseInvoiceApi.PostRecognizeAndParseInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRecognizeAndParseInvoice`: string
    fmt.Fprintf(os.Stdout, "Response from `RecognizeAndParseInvoiceApi.PostRecognizeAndParseInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRecognizeAndParseInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCRRecognizeAndParseInvoiceBody** | [**OCRRecognizeAndParseInvoiceBody**](OCRRecognizeAndParseInvoiceBody.md) |  | 

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

