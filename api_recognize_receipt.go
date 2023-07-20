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


// RecognizeReceiptApiService RecognizeReceiptApi service
type RecognizeReceiptApiService service

type ApiCancelRecognizeReceiptRequest struct {
	ctx context.Context
	ApiService *RecognizeReceiptApiService
	id *string
}

func (r ApiCancelRecognizeReceiptRequest) Id(id string) ApiCancelRecognizeReceiptRequest {
	r.id = &id
	return r
}

func (r ApiCancelRecognizeReceiptRequest) Execute() (*http.Response, error) {
	return r.ApiService.CancelRecognizeReceiptExecute(r)
}

/*
CancelRecognizeReceipt CancelRecognizeReceipt

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCancelRecognizeReceiptRequest
*/
func (a *RecognizeReceiptApiService) CancelRecognizeReceipt(ctx context.Context) ApiCancelRecognizeReceiptRequest {
	return ApiCancelRecognizeReceiptRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RecognizeReceiptApiService) CancelRecognizeReceiptExecute(r ApiCancelRecognizeReceiptRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecognizeReceiptApiService.CancelRecognizeReceipt")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v5.0/ocr/RecognizeReceipt"

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

type ApiGetRecognizeReceiptRequest struct {
	ctx context.Context
	ApiService *RecognizeReceiptApiService
	id *string
}

func (r ApiGetRecognizeReceiptRequest) Id(id string) ApiGetRecognizeReceiptRequest {
	r.id = &id
	return r
}

func (r ApiGetRecognizeReceiptRequest) Execute() (*OCRResponse, *http.Response, error) {
	return r.ApiService.GetRecognizeReceiptExecute(r)
}

/*
GetRecognizeReceipt GetRecognizeReceipt

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRecognizeReceiptRequest
*/
func (a *RecognizeReceiptApiService) GetRecognizeReceipt(ctx context.Context) ApiGetRecognizeReceiptRequest {
	return ApiGetRecognizeReceiptRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OCRResponse
func (a *RecognizeReceiptApiService) GetRecognizeReceiptExecute(r ApiGetRecognizeReceiptRequest) (*OCRResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OCRResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecognizeReceiptApiService.GetRecognizeReceipt")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v5.0/ocr/RecognizeReceipt"

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

type ApiPostRecognizeReceiptRequest struct {
	ctx context.Context
	ApiService *RecognizeReceiptApiService
	oCRRecognizeReceiptBody *OCRRecognizeReceiptBody
}

func (r ApiPostRecognizeReceiptRequest) OCRRecognizeReceiptBody(oCRRecognizeReceiptBody OCRRecognizeReceiptBody) ApiPostRecognizeReceiptRequest {
	r.oCRRecognizeReceiptBody = &oCRRecognizeReceiptBody
	return r
}

func (r ApiPostRecognizeReceiptRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.PostRecognizeReceiptExecute(r)
}

/*
PostRecognizeReceipt PostRecognizeReceipt

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostRecognizeReceiptRequest
*/
func (a *RecognizeReceiptApiService) PostRecognizeReceipt(ctx context.Context) ApiPostRecognizeReceiptRequest {
	return ApiPostRecognizeReceiptRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *RecognizeReceiptApiService) PostRecognizeReceiptExecute(r ApiPostRecognizeReceiptRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecognizeReceiptApiService.PostRecognizeReceipt")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v5.0/ocr/RecognizeReceipt"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oCRRecognizeReceiptBody == nil {
		return localVarReturnValue, nil, reportError("oCRRecognizeReceiptBody is required and must be specified")
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
	localVarPostBody = r.oCRRecognizeReceiptBody
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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
