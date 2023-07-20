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


// DeskewImageApiService DeskewImageApi service
type DeskewImageApiService service

type ApiCancelDeskewImageRequest struct {
	ctx context.Context
	ApiService *DeskewImageApiService
	id *string
}

func (r ApiCancelDeskewImageRequest) Id(id string) ApiCancelDeskewImageRequest {
	r.id = &id
	return r
}

func (r ApiCancelDeskewImageRequest) Execute() (*http.Response, error) {
	return r.ApiService.CancelDeskewImageExecute(r)
}

/*
CancelDeskewImage CancelDeskewImage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCancelDeskewImageRequest
*/
func (a *DeskewImageApiService) CancelDeskewImage(ctx context.Context) ApiCancelDeskewImageRequest {
	return ApiCancelDeskewImageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DeskewImageApiService) CancelDeskewImageExecute(r ApiCancelDeskewImageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeskewImageApiService.CancelDeskewImage")
	if err != nil {
		return nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v5.0/ocr/DeskewImage"

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

type ApiGetDeskewImageRequest struct {
	ctx context.Context
	ApiService *DeskewImageApiService
	id *string
}

func (r ApiGetDeskewImageRequest) Id(id string) ApiGetDeskewImageRequest {
	r.id = &id
	return r
}

func (r ApiGetDeskewImageRequest) Execute() (*OCRResponse, *http.Response, error) {
	return r.ApiService.GetDeskewImageExecute(r)
}

/*
GetDeskewImage GetDeskewImage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDeskewImageRequest
*/
func (a *DeskewImageApiService) GetDeskewImage(ctx context.Context) ApiGetDeskewImageRequest {
	return ApiGetDeskewImageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OCRResponse
func (a *DeskewImageApiService) GetDeskewImageExecute(r ApiGetDeskewImageRequest) (*OCRResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OCRResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeskewImageApiService.GetDeskewImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v5.0/ocr/DeskewImage"

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

type ApiPostDeskewImageRequest struct {
	ctx context.Context
	ApiService *DeskewImageApiService
	oCRDeskewImageBody *OCRDeskewImageBody
}

func (r ApiPostDeskewImageRequest) OCRDeskewImageBody(oCRDeskewImageBody OCRDeskewImageBody) ApiPostDeskewImageRequest {
	r.oCRDeskewImageBody = &oCRDeskewImageBody
	return r
}

func (r ApiPostDeskewImageRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.PostDeskewImageExecute(r)
}

/*
PostDeskewImage PostDeskewImage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostDeskewImageRequest
*/
func (a *DeskewImageApiService) PostDeskewImage(ctx context.Context) ApiPostDeskewImageRequest {
	return ApiPostDeskewImageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *DeskewImageApiService) PostDeskewImageExecute(r ApiPostDeskewImageRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeskewImageApiService.PostDeskewImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v5.0/ocr/DeskewImage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oCRDeskewImageBody == nil {
		return localVarReturnValue, nil, reportError("oCRDeskewImageBody is required and must be specified")
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
	localVarPostBody = r.oCRDeskewImageBody
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
