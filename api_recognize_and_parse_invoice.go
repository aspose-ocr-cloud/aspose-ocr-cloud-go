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


// RecognizeAndParseInvoiceApiService RecognizeAndParseInvoiceApi service
type RecognizeAndParseInvoiceApiService service

type ApiCancelRecognizeAndParseInvoiceRequest struct {
	ctx context.Context
	ApiService *RecognizeAndParseInvoiceApiService
	id *string
}

func (r ApiCancelRecognizeAndParseInvoiceRequest) Id(id string) ApiCancelRecognizeAndParseInvoiceRequest {
	r.id = &id
	return r
}

func (r ApiCancelRecognizeAndParseInvoiceRequest) Execute() (*http.Response, error) {
	return r.ApiService.CancelRecognizeAndParseInvoiceExecute(r)
}

/*
CancelRecognizeAndParseInvoice CancelRecognizeAndParseInvoice

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCancelRecognizeAndParseInvoiceRequest
*/
func (a *RecognizeAndParseInvoiceApiService) CancelRecognizeAndParseInvoice(ctx context.Context) ApiCancelRecognizeAndParseInvoiceRequest {
	return ApiCancelRecognizeAndParseInvoiceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RecognizeAndParseInvoiceApiService) CancelRecognizeAndParseInvoiceExecute(r ApiCancelRecognizeAndParseInvoiceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecognizeAndParseInvoiceApiService.CancelRecognizeAndParseInvoice")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v5.0/ocr/RecognizeAndParseInvoice"

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

type ApiGetRecognizeAndParseInvoiceRequest struct {
	ctx context.Context
	ApiService *RecognizeAndParseInvoiceApiService
	id *string
}

func (r ApiGetRecognizeAndParseInvoiceRequest) Id(id string) ApiGetRecognizeAndParseInvoiceRequest {
	r.id = &id
	return r
}

func (r ApiGetRecognizeAndParseInvoiceRequest) Execute() (*OCRResponse, *http.Response, error) {
	return r.ApiService.GetRecognizeAndParseInvoiceExecute(r)
}

/*
GetRecognizeAndParseInvoice GetRecognizeAndParseInvoice

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRecognizeAndParseInvoiceRequest
*/
func (a *RecognizeAndParseInvoiceApiService) GetRecognizeAndParseInvoice(ctx context.Context) ApiGetRecognizeAndParseInvoiceRequest {
	return ApiGetRecognizeAndParseInvoiceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OCRResponse
func (a *RecognizeAndParseInvoiceApiService) GetRecognizeAndParseInvoiceExecute(r ApiGetRecognizeAndParseInvoiceRequest) (*OCRResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OCRResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecognizeAndParseInvoiceApiService.GetRecognizeAndParseInvoice")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v5.0/ocr/RecognizeAndParseInvoice"

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

type ApiPostRecognizeAndParseInvoiceRequest struct {
	ctx context.Context
	ApiService *RecognizeAndParseInvoiceApiService
	oCRRecognizeAndParseInvoiceBody *OCRRecognizeAndParseInvoiceBody
}

func (r ApiPostRecognizeAndParseInvoiceRequest) OCRRecognizeAndParseInvoiceBody(oCRRecognizeAndParseInvoiceBody OCRRecognizeAndParseInvoiceBody) ApiPostRecognizeAndParseInvoiceRequest {
	r.oCRRecognizeAndParseInvoiceBody = &oCRRecognizeAndParseInvoiceBody
	return r
}

func (r ApiPostRecognizeAndParseInvoiceRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.PostRecognizeAndParseInvoiceExecute(r)
}

/*
PostRecognizeAndParseInvoice PostRecognizeAndParseInvoice

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostRecognizeAndParseInvoiceRequest
*/
func (a *RecognizeAndParseInvoiceApiService) PostRecognizeAndParseInvoice(ctx context.Context) ApiPostRecognizeAndParseInvoiceRequest {
	return ApiPostRecognizeAndParseInvoiceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *RecognizeAndParseInvoiceApiService) PostRecognizeAndParseInvoiceExecute(r ApiPostRecognizeAndParseInvoiceRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecognizeAndParseInvoiceApiService.PostRecognizeAndParseInvoice")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v5.0/ocr/RecognizeAndParseInvoice"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oCRRecognizeAndParseInvoiceBody == nil {
		return localVarReturnValue, nil, reportError("oCRRecognizeAndParseInvoiceBody is required and must be specified")
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
	localVarPostBody = r.oCRRecognizeAndParseInvoiceBody
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
