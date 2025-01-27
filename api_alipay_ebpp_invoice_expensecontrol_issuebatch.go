/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// AlipayEbppInvoiceExpensecontrolIssuebatchAPIService AlipayEbppInvoiceExpensecontrolIssuebatchAPI service
type AlipayEbppInvoiceExpensecontrolIssuebatchAPIService service

type ApiAlipayEbppInvoiceExpensecontrolIssuebatchCancelRequest struct {
	ctx                                                  context.Context
	ApiService                                           *AlipayEbppInvoiceExpensecontrolIssuebatchAPIService
	alipayEbppInvoiceExpensecontrolIssuebatchCancelModel *AlipayEbppInvoiceExpensecontrolIssuebatchCancelModel
}

func (r ApiAlipayEbppInvoiceExpensecontrolIssuebatchCancelRequest) AlipayEbppInvoiceExpensecontrolIssuebatchCancelModel(alipayEbppInvoiceExpensecontrolIssuebatchCancelModel AlipayEbppInvoiceExpensecontrolIssuebatchCancelModel) ApiAlipayEbppInvoiceExpensecontrolIssuebatchCancelRequest {
	r.alipayEbppInvoiceExpensecontrolIssuebatchCancelModel = &alipayEbppInvoiceExpensecontrolIssuebatchCancelModel
	return r
}

func (r ApiAlipayEbppInvoiceExpensecontrolIssuebatchCancelRequest) Execute() (*AlipayEbppInvoiceExpensecontrolIssuebatchCancelResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEbppInvoiceExpensecontrolIssuebatchCancelExecute(r)
}

/*
AlipayEbppInvoiceExpensecontrolIssuebatchCancel 作废额度发放

通过发放批次id，作废当前批次下发放的额度

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEbppInvoiceExpensecontrolIssuebatchCancelRequest
*/
func (r *AlipayEbppInvoiceExpensecontrolIssuebatchAPIService) AlipayEbppInvoiceExpensecontrolIssuebatchCancel(ctx context.Context) ApiAlipayEbppInvoiceExpensecontrolIssuebatchCancelRequest {
	return ApiAlipayEbppInvoiceExpensecontrolIssuebatchCancelRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEbppInvoiceExpensecontrolIssuebatchCancelResponseModel
func (a *AlipayEbppInvoiceExpensecontrolIssuebatchAPIService) AlipayEbppInvoiceExpensecontrolIssuebatchCancelExecute(r ApiAlipayEbppInvoiceExpensecontrolIssuebatchCancelRequest) (*AlipayEbppInvoiceExpensecontrolIssuebatchCancelResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEbppInvoiceExpensecontrolIssuebatchCancelResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEbppInvoiceExpensecontrolIssuebatchAPIService.AlipayEbppInvoiceExpensecontrolIssuebatchCancel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/ebpp/invoice/expensecontrol/issuebatch/cancel"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.alipayEbppInvoiceExpensecontrolIssuebatchCancelModel

	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	// Add signing logic
	err = a.signRequest(req)
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

	// Add verification logic
	err = a.verifyResponse(localVarHTTPResponse, localVarBody)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v AlipayEbppInvoiceExpensecontrolIssuebatchCancelDefaultResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAlipayEbppInvoiceExpensecontrolIssuebatchCreateRequest struct {
	ctx                                                  context.Context
	ApiService                                           *AlipayEbppInvoiceExpensecontrolIssuebatchAPIService
	alipayEbppInvoiceExpensecontrolIssuebatchCreateModel *AlipayEbppInvoiceExpensecontrolIssuebatchCreateModel
}

func (r ApiAlipayEbppInvoiceExpensecontrolIssuebatchCreateRequest) AlipayEbppInvoiceExpensecontrolIssuebatchCreateModel(alipayEbppInvoiceExpensecontrolIssuebatchCreateModel AlipayEbppInvoiceExpensecontrolIssuebatchCreateModel) ApiAlipayEbppInvoiceExpensecontrolIssuebatchCreateRequest {
	r.alipayEbppInvoiceExpensecontrolIssuebatchCreateModel = &alipayEbppInvoiceExpensecontrolIssuebatchCreateModel
	return r
}

func (r ApiAlipayEbppInvoiceExpensecontrolIssuebatchCreateRequest) Execute() (*AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEbppInvoiceExpensecontrolIssuebatchCreateExecute(r)
}

/*
AlipayEbppInvoiceExpensecontrolIssuebatchCreate 手动发放额度接口

通过该接口对企业下的员工进行批量的额度发放。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEbppInvoiceExpensecontrolIssuebatchCreateRequest
*/
func (r *AlipayEbppInvoiceExpensecontrolIssuebatchAPIService) AlipayEbppInvoiceExpensecontrolIssuebatchCreate(ctx context.Context) ApiAlipayEbppInvoiceExpensecontrolIssuebatchCreateRequest {
	return ApiAlipayEbppInvoiceExpensecontrolIssuebatchCreateRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel
func (a *AlipayEbppInvoiceExpensecontrolIssuebatchAPIService) AlipayEbppInvoiceExpensecontrolIssuebatchCreateExecute(r ApiAlipayEbppInvoiceExpensecontrolIssuebatchCreateRequest) (*AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEbppInvoiceExpensecontrolIssuebatchCreateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEbppInvoiceExpensecontrolIssuebatchAPIService.AlipayEbppInvoiceExpensecontrolIssuebatchCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/ebpp/invoice/expensecontrol/issuebatch/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.alipayEbppInvoiceExpensecontrolIssuebatchCreateModel

	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	// Add signing logic
	err = a.signRequest(req)
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

	// Add verification logic
	err = a.verifyResponse(localVarHTTPResponse, localVarBody)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v AlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

func (a *AlipayEbppInvoiceExpensecontrolIssuebatchAPIService) signRequest(req *http.Request) error {
	appID := a.client.cfg.AppID
	appCertSN := a.client.cfg.AppCertSN
	nonce := generateUUID()
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)

	authString := fmt.Sprintf("app_id=%s", appID)
	if appCertSN != "" {
		authString += fmt.Sprintf(",app_cert_sn=%s", appCertSN)
	}
	authString += fmt.Sprintf(",nonce=%s,timestamp=%s", nonce, timestamp)

	httpMethod := req.Method
	httpRequestUri := req.URL.Path
	if req.URL.RawQuery != "" {
		httpRequestUri += "?" + req.URL.RawQuery
	}

	var httpRequestBody string
	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return err
		}
		httpRequestBody = string(bodyBytes)
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	content := authString + "\n" +
		httpMethod + "\n" +
		httpRequestUri + "\n" +
		httpRequestBody + "\n"

	if appAuthToken := req.Header.Get("alipay-app-auth-token"); appAuthToken != "" {
		content += appAuthToken + "\n"
	}

	signature, err := signWithRSA(content, a.client.cfg.privateKey)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("ALIPAY-SHA256withRSA %s,sign=%s", authString, signature))
	return nil
}

func (a *AlipayEbppInvoiceExpensecontrolIssuebatchAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
