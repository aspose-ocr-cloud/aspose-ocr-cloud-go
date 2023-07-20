# \DetectRegionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDetectRegions**](DetectRegionsApi.md#CancelDetectRegions) | **Delete** /v5.0/ocr/DetectRegions | CancelDetectRegions
[**GetDetectRegions**](DetectRegionsApi.md#GetDetectRegions) | **Get** /v5.0/ocr/DetectRegions | GetDetectRegions
[**PostDetectRegions**](DetectRegionsApi.md#PostDetectRegions) | **Post** /v5.0/ocr/DetectRegions | PostDetectRegions



## CancelDetectRegions

> CancelDetectRegions(ctx).Id(id).Execute()

CancelDetectRegions

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
    r, err := apiClient.DetectRegionsApi.CancelDetectRegions(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DetectRegionsApi.CancelDetectRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelDetectRegionsRequest struct via the builder pattern


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


## GetDetectRegions

> OCRResponse GetDetectRegions(ctx).Id(id).Execute()

GetDetectRegions

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
    resp, r, err := apiClient.DetectRegionsApi.GetDetectRegions(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DetectRegionsApi.GetDetectRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDetectRegions`: OCRResponse
    fmt.Fprintf(os.Stdout, "Response from `DetectRegionsApi.GetDetectRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDetectRegionsRequest struct via the builder pattern


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


## PostDetectRegions

> string PostDetectRegions(ctx).OCRDetectRegionsBody(oCRDetectRegionsBody).Execute()

PostDetectRegions

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
    oCRDetectRegionsBody := *openapiclient.NewOCRDetectRegionsBody(string(123), *openapiclient.NewOCRSettingsDetectRegions()) // OCRDetectRegionsBody | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.DetectRegionsApi.PostDetectRegions(context.Background()).OCRDetectRegionsBody(oCRDetectRegionsBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DetectRegionsApi.PostDetectRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDetectRegions`: string
    fmt.Fprintf(os.Stdout, "Response from `DetectRegionsApi.PostDetectRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDetectRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCRDetectRegionsBody** | [**OCRDetectRegionsBody**](OCRDetectRegionsBody.md) |  | 

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

