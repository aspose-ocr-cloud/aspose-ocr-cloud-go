/*
Aspose OCR Cloud 5.0 API

Aspose OCR Cloud 5.0 API

API version: 5.0
*/


package asposeocrcloud

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ConvertTextToSpeechTrialApiService ConvertTextToSpeechTrialApi service
type ConvertTextToSpeechTrialApiService service

type ApiCancelConvertTextToSpeechTrialRequest struct {
	ctx context.Context
	ApiService *ConvertTextToSpeechTrialApiService
	id *string
}

func (r ApiCancelConvertTextToSpeechTrialRequest) Id(id string) ApiCancelConvertTextToSpeechTrialRequest {
	r.id = &id
	return r
}

func (r ApiCancelConvertTextToSpeechTrialRequest) Execute() (*http.Response, error) {
	return r.ApiService.CancelConvertTextToSpeechTrialExecute(r)
}

/*
CancelConvertTextToSpeechTrial CancelConvertTextToSpeechTrial

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCancelConvertTextToSpeechTrialRequest
*/
func (a *ConvertTextToSpeechTrialApiService) CancelConvertTextToSpeechTrial(ctx context.Context) ApiCancelConvertTextToSpeechTrialRequest {
	return ApiCancelConvertTextToSpeechTrialRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConvertTextToSpeechTrialApiService) CancelConvertTextToSpeechTrialExecute(r ApiCancelConvertTextToSpeechTrialRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConvertTextToSpeechTrialApiService.CancelConvertTextToSpeechTrial")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v5.0/ocr/ConvertTextToSpeechTrial"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.id == nil {
		return nil, reportError("id is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetConvertTextToSpeechTrialRequest struct {
	ctx context.Context
	ApiService *ConvertTextToSpeechTrialApiService
	id *string
}

func (r ApiGetConvertTextToSpeechTrialRequest) Id(id string) ApiGetConvertTextToSpeechTrialRequest {
	r.id = &id
	return r
}

func (r ApiGetConvertTextToSpeechTrialRequest) Execute() (*TTSResponse, *http.Response, error) {
	return r.ApiService.GetConvertTextToSpeechTrialExecute(r)
}

/*
GetConvertTextToSpeechTrial GetConvertTextToSpeechTrial

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetConvertTextToSpeechTrialRequest
*/
func (a *ConvertTextToSpeechTrialApiService) GetConvertTextToSpeechTrial(ctx context.Context) ApiGetConvertTextToSpeechTrialRequest {
	return ApiGetConvertTextToSpeechTrialRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TTSResponse
func (a *ConvertTextToSpeechTrialApiService) GetConvertTextToSpeechTrialExecute(r ApiGetConvertTextToSpeechTrialRequest) (*TTSResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TTSResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConvertTextToSpeechTrialApiService.GetConvertTextToSpeechTrial")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v5.0/ocr/ConvertTextToSpeechTrial"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.id == nil {
		return localVarReturnValue, nil, reportError("id is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostConvertTextToSpeechTrialRequest struct {
	ctx context.Context
	ApiService *ConvertTextToSpeechTrialApiService
	tTSBody *TTSBody
}

func (r ApiPostConvertTextToSpeechTrialRequest) TTSBody(tTSBody TTSBody) ApiPostConvertTextToSpeechTrialRequest {
	r.tTSBody = &tTSBody
	return r
}

func (r ApiPostConvertTextToSpeechTrialRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.PostConvertTextToSpeechTrialExecute(r)
}

/*
PostConvertTextToSpeechTrial PostConvertTextToSpeechTrial

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostConvertTextToSpeechTrialRequest
*/
func (a *ConvertTextToSpeechTrialApiService) PostConvertTextToSpeechTrial(ctx context.Context) ApiPostConvertTextToSpeechTrialRequest {
	return ApiPostConvertTextToSpeechTrialRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *ConvertTextToSpeechTrialApiService) PostConvertTextToSpeechTrialExecute(r ApiPostConvertTextToSpeechTrialRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConvertTextToSpeechTrialApiService.PostConvertTextToSpeechTrial")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v5.0/ocr/ConvertTextToSpeechTrial"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tTSBody == nil {
		return localVarReturnValue, nil, reportError("tTSBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.tTSBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
