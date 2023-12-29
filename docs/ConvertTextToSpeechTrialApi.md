# \ConvertTextToSpeechTrialApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelConvertTextToSpeechTrial**](ConvertTextToSpeechTrialApi.md#CancelConvertTextToSpeechTrial) | **Delete** /v5.0/ocr/ConvertTextToSpeechTrial | CancelConvertTextToSpeechTrial
[**GetConvertTextToSpeechTrial**](ConvertTextToSpeechTrialApi.md#GetConvertTextToSpeechTrial) | **Get** /v5.0/ocr/ConvertTextToSpeechTrial | GetConvertTextToSpeechTrial
[**PostConvertTextToSpeechTrial**](ConvertTextToSpeechTrialApi.md#PostConvertTextToSpeechTrial) | **Post** /v5.0/ocr/ConvertTextToSpeechTrial | PostConvertTextToSpeechTrial



## CancelConvertTextToSpeechTrial

> CancelConvertTextToSpeechTrial(ctx).Id(id).Execute()

CancelConvertTextToSpeechTrial

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
    r, err := apiClient.ConvertTextToSpeechTrialApi.CancelConvertTextToSpeechTrial(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvertTextToSpeechTrialApi.CancelConvertTextToSpeechTrial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelConvertTextToSpeechTrialRequest struct via the builder pattern


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


## GetConvertTextToSpeechTrial

> TTSResponse GetConvertTextToSpeechTrial(ctx).Id(id).Execute()

GetConvertTextToSpeechTrial

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
    resp, r, err := apiClient.ConvertTextToSpeechTrialApi.GetConvertTextToSpeechTrial(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvertTextToSpeechTrialApi.GetConvertTextToSpeechTrial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConvertTextToSpeechTrial`: TTSResponse
    fmt.Fprintf(os.Stdout, "Response from `ConvertTextToSpeechTrialApi.GetConvertTextToSpeechTrial`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConvertTextToSpeechTrialRequest struct via the builder pattern


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


## PostConvertTextToSpeechTrial

> string PostConvertTextToSpeechTrial(ctx).TTSBody(tTSBody).Execute()

PostConvertTextToSpeechTrial

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
    resp, r, err := apiClient.ConvertTextToSpeechTrialApi.PostConvertTextToSpeechTrial(context.Background()).TTSBody(tTSBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConvertTextToSpeechTrialApi.PostConvertTextToSpeechTrial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostConvertTextToSpeechTrial`: string
    fmt.Fprintf(os.Stdout, "Response from `ConvertTextToSpeechTrialApi.PostConvertTextToSpeechTrial`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostConvertTextToSpeechTrialRequest struct via the builder pattern


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

