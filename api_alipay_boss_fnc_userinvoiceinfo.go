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

// AlipayBossFncUserinvoiceinfoAPIService AlipayBossFncUserinvoiceinfoAPI service
type AlipayBossFncUserinvoiceinfoAPIService service

type ApiAlipayBossFncUserinvoiceinfoCreateRequest struct {
	ctx                                     context.Context
	ApiService                              *AlipayBossFncUserinvoiceinfoAPIService
	alipayBossFncUserinvoiceinfoCreateModel *AlipayBossFncUserinvoiceinfoCreateModel
}

func (r ApiAlipayBossFncUserinvoiceinfoCreateRequest) AlipayBossFncUserinvoiceinfoCreateModel(alipayBossFncUserinvoiceinfoCreateModel AlipayBossFncUserinvoiceinfoCreateModel) ApiAlipayBossFncUserinvoiceinfoCreateRequest {
	r.alipayBossFncUserinvoiceinfoCreateModel = &alipayBossFncUserinvoiceinfoCreateModel
	return r
}

func (r ApiAlipayBossFncUserinvoiceinfoCreateRequest) Execute() (*AlipayBossFncUserinvoiceinfoCreateResponseModel, *http.Response, error) {
	return r.ApiService.AlipayBossFncUserinvoiceinfoCreateExecute(r)
}

/*
AlipayBossFncUserinvoiceinfoCreate 新增客户开票资料

申请开票alipay.boss.fnc.invoice.create之前需要维护的开票资料信息

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayBossFncUserinvoiceinfoCreateRequest
*/
func (r *AlipayBossFncUserinvoiceinfoAPIService) AlipayBossFncUserinvoiceinfoCreate(ctx context.Context) ApiAlipayBossFncUserinvoiceinfoCreateRequest {
	return ApiAlipayBossFncUserinvoiceinfoCreateRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayBossFncUserinvoiceinfoCreateResponseModel
func (a *AlipayBossFncUserinvoiceinfoAPIService) AlipayBossFncUserinvoiceinfoCreateExecute(r ApiAlipayBossFncUserinvoiceinfoCreateRequest) (*AlipayBossFncUserinvoiceinfoCreateResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayBossFncUserinvoiceinfoCreateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayBossFncUserinvoiceinfoAPIService.AlipayBossFncUserinvoiceinfoCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/boss/fnc/userinvoiceinfo/create"

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
	localVarPostBody = r.alipayBossFncUserinvoiceinfoCreateModel

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
		var v AlipayBossFncUserinvoiceinfoCreateDefaultResponse
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

type ApiAlipayBossFncUserinvoiceinfoModifyRequest struct {
	ctx                                     context.Context
	ApiService                              *AlipayBossFncUserinvoiceinfoAPIService
	alipayBossFncUserinvoiceinfoModifyModel *AlipayBossFncUserinvoiceinfoModifyModel
}

func (r ApiAlipayBossFncUserinvoiceinfoModifyRequest) AlipayBossFncUserinvoiceinfoModifyModel(alipayBossFncUserinvoiceinfoModifyModel AlipayBossFncUserinvoiceinfoModifyModel) ApiAlipayBossFncUserinvoiceinfoModifyRequest {
	r.alipayBossFncUserinvoiceinfoModifyModel = &alipayBossFncUserinvoiceinfoModifyModel
	return r
}

func (r ApiAlipayBossFncUserinvoiceinfoModifyRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayBossFncUserinvoiceinfoModifyExecute(r)
}

/*
AlipayBossFncUserinvoiceinfoModify 修改用户开票资料

通过查询用户的开票资料（alipay.boss.fnc.userinvoiceinfo.query）后进行开票资料的修改

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayBossFncUserinvoiceinfoModifyRequest
*/
func (r *AlipayBossFncUserinvoiceinfoAPIService) AlipayBossFncUserinvoiceinfoModify(ctx context.Context) ApiAlipayBossFncUserinvoiceinfoModifyRequest {
	return ApiAlipayBossFncUserinvoiceinfoModifyRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *AlipayBossFncUserinvoiceinfoAPIService) AlipayBossFncUserinvoiceinfoModifyExecute(r ApiAlipayBossFncUserinvoiceinfoModifyRequest) (map[string]interface{}, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayBossFncUserinvoiceinfoAPIService.AlipayBossFncUserinvoiceinfoModify")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/boss/fnc/userinvoiceinfo/modify"

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
	localVarPostBody = r.alipayBossFncUserinvoiceinfoModifyModel

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
		var v AlipayBossFncUserinvoiceinfoModifyDefaultResponse
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

type ApiAlipayBossFncUserinvoiceinfoQueryRequest struct {
	ctx        context.Context
	ApiService *AlipayBossFncUserinvoiceinfoAPIService
	pid        *string
}

// 开票pid/mid/ou，唯一标识商户信息/集团用户信息的ID
func (r ApiAlipayBossFncUserinvoiceinfoQueryRequest) Pid(pid string) ApiAlipayBossFncUserinvoiceinfoQueryRequest {
	r.pid = &pid
	return r
}

func (r ApiAlipayBossFncUserinvoiceinfoQueryRequest) Execute() (*AlipayBossFncUserinvoiceinfoQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayBossFncUserinvoiceinfoQueryExecute(r)
}

/*
AlipayBossFncUserinvoiceinfoQuery 根据PID获取相关的开票资料

根据PID获取相关的开票资料

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayBossFncUserinvoiceinfoQueryRequest
*/
func (r *AlipayBossFncUserinvoiceinfoAPIService) AlipayBossFncUserinvoiceinfoQuery(ctx context.Context) ApiAlipayBossFncUserinvoiceinfoQueryRequest {
	return ApiAlipayBossFncUserinvoiceinfoQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayBossFncUserinvoiceinfoQueryResponseModel
func (a *AlipayBossFncUserinvoiceinfoAPIService) AlipayBossFncUserinvoiceinfoQueryExecute(r ApiAlipayBossFncUserinvoiceinfoQueryRequest) (*AlipayBossFncUserinvoiceinfoQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayBossFncUserinvoiceinfoQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayBossFncUserinvoiceinfoAPIService.AlipayBossFncUserinvoiceinfoQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/boss/fnc/userinvoiceinfo/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pid", r.pid, "form", "")
	}

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
		var v AlipayBossFncUserinvoiceinfoQueryDefaultResponse
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

func (a *AlipayBossFncUserinvoiceinfoAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayBossFncUserinvoiceinfoAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
