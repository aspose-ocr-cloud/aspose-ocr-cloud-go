# \RecognizeRegionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRecognizeRegions**](RecognizeRegionsApi.md#CancelRecognizeRegions) | **Delete** /v5.0/ocr/RecognizeRegions | CancelRecognizeRegions
[**GetRecognizeRegions**](RecognizeRegionsApi.md#GetRecognizeRegions) | **Get** /v5.0/ocr/RecognizeRegions | GetRecognizeRegions
[**PostRecognizeRegions**](RecognizeRegionsApi.md#PostRecognizeRegions) | **Post** /v5.0/ocr/RecognizeRegions | PostRecognizeRegions



## CancelRecognizeRegions

> CancelRecognizeRegions(ctx).Id(id).Execute()

CancelRecognizeRegions

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
    r, err := apiClient.RecognizeRegionsApi.CancelRecognizeRegions(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeRegionsApi.CancelRecognizeRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelRecognizeRegionsRequest struct via the builder pattern


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


## GetRecognizeRegions

> OCRResponse GetRecognizeRegions(ctx).Id(id).Execute()

GetRecognizeRegions

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
    resp, r, err := apiClient.RecognizeRegionsApi.GetRecognizeRegions(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeRegionsApi.GetRecognizeRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecognizeRegions`: OCRResponse
    fmt.Fprintf(os.Stdout, "Response from `RecognizeRegionsApi.GetRecognizeRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecognizeRegionsRequest struct via the builder pattern


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


## PostRecognizeRegions

> string PostRecognizeRegions(ctx).OCRRecognizeRegionsBody(oCRRecognizeRegionsBody).Execute()

PostRecognizeRegions

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
    oCRRecognizeRegionsBody := *openapiclient.NewOCRRecognizeRegionsBody(string(123), *openapiclient.NewOCRSettingsRecognizeRegions()) // OCRRecognizeRegionsBody | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.RecognizeRegionsApi.PostRecognizeRegions(context.Background()).OCRRecognizeRegionsBody(oCRRecognizeRegionsBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeRegionsApi.PostRecognizeRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRecognizeRegions`: string
    fmt.Fprintf(os.Stdout, "Response from `RecognizeRegionsApi.PostRecognizeRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRecognizeRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCRRecognizeRegionsBody** | [**OCRRecognizeRegionsBody**](OCRRecognizeRegionsBody.md) |  | 

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

