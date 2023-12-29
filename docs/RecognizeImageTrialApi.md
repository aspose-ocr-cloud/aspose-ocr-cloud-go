# \RecognizeImageTrialApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRecognizeImageTrial**](RecognizeImageTrialApi.md#CancelRecognizeImageTrial) | **Delete** /v5.0/ocr/RecognizeImageTrial | CancelRecognizeImageTrial
[**GetRecognizeImageTrial**](RecognizeImageTrialApi.md#GetRecognizeImageTrial) | **Get** /v5.0/ocr/RecognizeImageTrial | GetRecognizeImageTrial
[**PostRecognizeImageTrial**](RecognizeImageTrialApi.md#PostRecognizeImageTrial) | **Post** /v5.0/ocr/RecognizeImageTrial | PostRecognizeImageTrial



## CancelRecognizeImageTrial

> CancelRecognizeImageTrial(ctx).Id(id).Execute()

CancelRecognizeImageTrial

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
    r, err := apiClient.RecognizeImageTrialApi.CancelRecognizeImageTrial(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeImageTrialApi.CancelRecognizeImageTrial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelRecognizeImageTrialRequest struct via the builder pattern


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


## GetRecognizeImageTrial

> OCRResponse GetRecognizeImageTrial(ctx).Id(id).Execute()

GetRecognizeImageTrial

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
    resp, r, err := apiClient.RecognizeImageTrialApi.GetRecognizeImageTrial(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeImageTrialApi.GetRecognizeImageTrial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecognizeImageTrial`: OCRResponse
    fmt.Fprintf(os.Stdout, "Response from `RecognizeImageTrialApi.GetRecognizeImageTrial`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecognizeImageTrialRequest struct via the builder pattern


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


## PostRecognizeImageTrial

> string PostRecognizeImageTrial(ctx).OCRRecognizeImageBody(oCRRecognizeImageBody).Execute()

PostRecognizeImageTrial

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
    oCRRecognizeImageBody := *openapiclient.NewOCRRecognizeImageBody(string(123), *openapiclient.NewOCRSettingsRecognizeImage()) // OCRRecognizeImageBody | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.RecognizeImageTrialApi.PostRecognizeImageTrial(context.Background()).OCRRecognizeImageBody(oCRRecognizeImageBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecognizeImageTrialApi.PostRecognizeImageTrial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRecognizeImageTrial`: string
    fmt.Fprintf(os.Stdout, "Response from `RecognizeImageTrialApi.PostRecognizeImageTrial`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRecognizeImageTrialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCRRecognizeImageBody** | [**OCRRecognizeImageBody**](OCRRecognizeImageBody.md) |  | 

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

