# \TextToSpeechApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTextToSpeechResult**](TextToSpeechApi.md#GetTextToSpeechResult) | **Get** /v5.0/ocr/TextToSpeech/GetTextToSpeechResult | GetTextToSpeechResult
[**GetTextToSpeechResultFile**](TextToSpeechApi.md#GetTextToSpeechResultFile) | **Get** /v5.0/ocr/TextToSpeech/GetTextToSpeechResultFile | GetTextToSpeechResultFile
[**PostTextToSpeech**](TextToSpeechApi.md#PostTextToSpeech) | **Post** /v5.0/ocr/TextToSpeech/PostTextToSpeech | PostTextToSpeech



## GetTextToSpeechResult

> TTSResponse GetTextToSpeechResult(ctx).Id(id).Execute()

GetTextToSpeechResult

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
    resp, r, err := apiClient.TextToSpeechApi.GetTextToSpeechResult(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TextToSpeechApi.GetTextToSpeechResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTextToSpeechResult`: TTSResponse
    fmt.Fprintf(os.Stdout, "Response from `TextToSpeechApi.GetTextToSpeechResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTextToSpeechResultRequest struct via the builder pattern


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


## GetTextToSpeechResultFile

> map[string]interface{} GetTextToSpeechResultFile(ctx).Id(id).Execute()

GetTextToSpeechResultFile

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
    resp, r, err := apiClient.TextToSpeechApi.GetTextToSpeechResultFile(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TextToSpeechApi.GetTextToSpeechResultFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTextToSpeechResultFile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TextToSpeechApi.GetTextToSpeechResultFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTextToSpeechResultFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTextToSpeech

> string PostTextToSpeech(ctx).TTSBodyDeprecated(tTSBodyDeprecated).Execute()

PostTextToSpeech

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
    tTSBodyDeprecated := *openapiclient.NewTTSBodyDeprecated("Text_example", openapiclient.LanguageTTS("English"), openapiclient.ResultTypeTTS("Wav")) // TTSBodyDeprecated | 

    configuration := asposeocrcloud.NewConfiguration()
    apiClient := asposeocrcloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TextToSpeechApi.PostTextToSpeech(context.Background()).TTSBodyDeprecated(tTSBodyDeprecated).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TextToSpeechApi.PostTextToSpeech``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTextToSpeech`: string
    fmt.Fprintf(os.Stdout, "Response from `TextToSpeechApi.PostTextToSpeech`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTextToSpeechRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tTSBodyDeprecated** | [**TTSBodyDeprecated**](TTSBodyDeprecated.md) |  | 

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

