# \ConvertTextToSpeechApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelConvertTextToSpeech**](ConvertTextToSpeechApi.md#CancelConvertTextToSpeech) | **Delete** /v5.0/ocr/ConvertTextToSpeech | CancelConvertTextToSpeech
[**GetConvertTextToSpeech**](ConvertTextToSpeechApi.md#GetConvertTextToSpeech) | **Get** /v5.0/ocr/ConvertTextToSpeech | GetConvertTextToSpeech
[**PostConvertTextToSpeech**](ConvertTextToSpeechApi.md#PostConvertTextToSpeech) | **Post** /v5.0/ocr/ConvertTextToSpeech | PostConvertTextToSpeech



## CancelConvertTextToSpeech

> CancelConvertTextToSpeech(ctx).Id(id).Execute()

CancelConvertTextToSpeech

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
    r, err := apiClient.ConvertTextToSpeechApi.CancelConvertTextToSpeech(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvertTextToSpeechApi.CancelConvertTextToSpeech``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelConvertTextToSpeechRequest struct via the builder pattern


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


## GetConvertTextToSpeech

> TTSResponse GetConvertTextToSpeech(ctx).Id(id).Execute()

GetConvertTextToSpeech

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
    resp, r, err := apiClient.ConvertTextToSpeechApi.GetConvertTextToSpeech(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvertTextToSpeechApi.GetConvertTextToSpeech``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConvertTextToSpeech`: TTSResponse
    fmt.Fprintf(os.Stdout, "Response from `ConvertTextToSpeechApi.GetConvertTextToSpeech`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConvertTextToSpeechRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

### Return type

[**TTSResponse**](TTSResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConvertTextToSpeech

> string PostConvertTextToSpeech(ctx).TTSBody(tTSBody).Execute()

PostConvertTextToSpeech

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
    tTSBody := *openapiclient.NewTTSBody("Text_example", *openapiclient.NewTTSSettings(openapiclient.LanguageTTS("English"), openapiclient.ResultTypeTTS("Wav"))) // TTSBody | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ConvertTextToSpeechApi.PostConvertTextToSpeech(context.Background()).TTSBody(tTSBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvertTextToSpeechApi.PostConvertTextToSpeech``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostConvertTextToSpeech`: string
    fmt.Fprintf(os.Stdout, "Response from `ConvertTextToSpeechApi.PostConvertTextToSpeech`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostConvertTextToSpeechRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tTSBody** | [**TTSBody**](TTSBody.md) |  | 

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

