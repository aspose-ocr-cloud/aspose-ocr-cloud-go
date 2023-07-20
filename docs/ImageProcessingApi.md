# \ImageProcessingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetResultFile**](ImageProcessingApi.md#GetResultFile) | **Get** /v5.0/ocr/ImageProcessing/GetResultFile | GetResultFile
[**GetResultTask**](ImageProcessingApi.md#GetResultTask) | **Get** /v5.0/ocr/ImageProcessing/GetResultTask | GetResultTask
[**PostBinarizationFile**](ImageProcessingApi.md#PostBinarizationFile) | **Post** /v5.0/ocr/ImageProcessing/PostBinarizationFile | PostBinarizationFile
[**PostDewarpingFile**](ImageProcessingApi.md#PostDewarpingFile) | **Post** /v5.0/ocr/ImageProcessing/PostDewarpingFile | PostDewarpingFile
[**PostSkewCorrectionFile**](ImageProcessingApi.md#PostSkewCorrectionFile) | **Post** /v5.0/ocr/ImageProcessing/PostSkewCorrectionFile | PostSkewCorrectionFile
[**PostUpsamplingFile**](ImageProcessingApi.md#PostUpsamplingFile) | **Post** /v5.0/ocr/ImageProcessing/PostUpsamplingImageFile | PostUpsamplingImageFile



## GetResultFile

> map[string]interface{} GetResultFile(ctx).Id(id).Execute()

GetResultFile

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
    resp, r, err := apiClient.ImageProcessingApi.GetResultFile(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageProcessingApi.GetResultFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResultFile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ImageProcessingApi.GetResultFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResultFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResultTask

> OCRResponse GetResultTask(ctx).Id(id).Execute()

GetResultTask

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
    resp, r, err := apiClient.ImageProcessingApi.GetResultTask(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageProcessingApi.GetResultTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResultTask`: OCRResponse
    fmt.Fprintf(os.Stdout, "Response from `ImageProcessingApi.GetResultTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResultTaskRequest struct via the builder pattern


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


## PostBinarizationFile

> string PostBinarizationFile(ctx).File(file).Execute()

PostBinarizationFile

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
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageProcessingApi.PostBinarizationFile(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageProcessingApi.PostBinarizationFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostBinarizationFile`: string
    fmt.Fprintf(os.Stdout, "Response from `ImageProcessingApi.PostBinarizationFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBinarizationFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

**string**

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDewarpingFile

> string PostDewarpingFile(ctx).File(file).Execute()

PostDewarpingFile

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
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageProcessingApi.PostDewarpingFile(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageProcessingApi.PostDewarpingFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDewarpingFile`: string
    fmt.Fprintf(os.Stdout, "Response from `ImageProcessingApi.PostDewarpingFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDewarpingFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

**string**

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSkewCorrectionFile

> string PostSkewCorrectionFile(ctx).File(file).Execute()

PostSkewCorrectionFile

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
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageProcessingApi.PostSkewCorrectionFile(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageProcessingApi.PostSkewCorrectionFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSkewCorrectionFile`: string
    fmt.Fprintf(os.Stdout, "Response from `ImageProcessingApi.PostSkewCorrectionFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSkewCorrectionFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

**string**

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUpsamplingFile

> string PostUpsamplingFile(ctx).File(file).Execute()

PostUpsamplingImageFile

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
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageProcessingApi.PostUpsamplingFile(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageProcessingApi.PostUpsamplingFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUpsamplingFile`: string
    fmt.Fprintf(os.Stdout, "Response from `ImageProcessingApi.PostUpsamplingFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUpsamplingFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

**string**

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

