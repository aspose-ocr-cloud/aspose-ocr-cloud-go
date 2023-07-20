# \IdentifyFontApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelIdentifyFont**](IdentifyFontApi.md#CancelIdentifyFont) | **Delete** /v5.0/ocr/IdentifyFont | CancelIdentifyFont
[**GetIdentifyFont**](IdentifyFontApi.md#GetIdentifyFont) | **Get** /v5.0/ocr/IdentifyFont | GetIdentifyFont
[**PostIdentifyFont**](IdentifyFontApi.md#PostIdentifyFont) | **Post** /v5.0/ocr/IdentifyFont | PostIdentifyFont



## CancelIdentifyFont

> CancelIdentifyFont(ctx).Id(id).Execute()

CancelIdentifyFont

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
    r, err := apiClient.IdentifyFontApi.CancelIdentifyFont(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentifyFontApi.CancelIdentifyFont``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelIdentifyFontRequest struct via the builder pattern


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


## GetIdentifyFont

> OCRResponse GetIdentifyFont(ctx).Id(id).Execute()

GetIdentifyFont

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
    resp, r, err := apiClient.IdentifyFontApi.GetIdentifyFont(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentifyFontApi.GetIdentifyFont``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentifyFont`: OCRResponse
    fmt.Fprintf(os.Stdout, "Response from `IdentifyFontApi.GetIdentifyFont`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentifyFontRequest struct via the builder pattern


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


## PostIdentifyFont

> string PostIdentifyFont(ctx).OCRRecognizeFontBody(oCRRecognizeFontBody).Execute()

PostIdentifyFont

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
    oCRRecognizeFontBody := *openapiclient.NewOCRRecognizeFontBody(string(123), *openapiclient.NewOCRSettingsRecognizeFont()) // OCRRecognizeFontBody | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentifyFontApi.PostIdentifyFont(context.Background()).OCRRecognizeFontBody(oCRRecognizeFontBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentifyFontApi.PostIdentifyFont``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostIdentifyFont`: string
    fmt.Fprintf(os.Stdout, "Response from `IdentifyFontApi.PostIdentifyFont`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIdentifyFontRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCRRecognizeFontBody** | [**OCRRecognizeFontBody**](OCRRecognizeFontBody.md) |  | 

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

