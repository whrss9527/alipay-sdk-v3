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

// AlipayCommerceLogisticsOrderInstantdeliveryAPIService AlipayCommerceLogisticsOrderInstantdeliveryAPI service
type AlipayCommerceLogisticsOrderInstantdeliveryAPIService service

type ApiAlipayCommerceLogisticsOrderInstantdeliveryCancelRequest struct {
	ctx                                                    context.Context
	ApiService                                             *AlipayCommerceLogisticsOrderInstantdeliveryAPIService
	alipayCommerceLogisticsOrderInstantdeliveryCancelModel *AlipayCommerceLogisticsOrderInstantdeliveryCancelModel
}

func (r ApiAlipayCommerceLogisticsOrderInstantdeliveryCancelRequest) AlipayCommerceLogisticsOrderInstantdeliveryCancelModel(alipayCommerceLogisticsOrderInstantdeliveryCancelModel AlipayCommerceLogisticsOrderInstantdeliveryCancelModel) ApiAlipayCommerceLogisticsOrderInstantdeliveryCancelRequest {
	r.alipayCommerceLogisticsOrderInstantdeliveryCancelModel = &alipayCommerceLogisticsOrderInstantdeliveryCancelModel
	return r
}

func (r ApiAlipayCommerceLogisticsOrderInstantdeliveryCancelRequest) Execute() (*AlipayCommerceLogisticsOrderInstantdeliveryCancelResponseModel, *http.Response, error) {
	return r.ApiService.AlipayCommerceLogisticsOrderInstantdeliveryCancelExecute(r)
}

/*
AlipayCommerceLogisticsOrderInstantdeliveryCancel 取消即时配送订单

取消即时配送订单

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayCommerceLogisticsOrderInstantdeliveryCancelRequest
*/
func (r *AlipayCommerceLogisticsOrderInstantdeliveryAPIService) AlipayCommerceLogisticsOrderInstantdeliveryCancel(ctx context.Context) ApiAlipayCommerceLogisticsOrderInstantdeliveryCancelRequest {
	return ApiAlipayCommerceLogisticsOrderInstantdeliveryCancelRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayCommerceLogisticsOrderInstantdeliveryCancelResponseModel
func (a *AlipayCommerceLogisticsOrderInstantdeliveryAPIService) AlipayCommerceLogisticsOrderInstantdeliveryCancelExecute(r ApiAlipayCommerceLogisticsOrderInstantdeliveryCancelRequest) (*AlipayCommerceLogisticsOrderInstantdeliveryCancelResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayCommerceLogisticsOrderInstantdeliveryCancelResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceLogisticsOrderInstantdeliveryAPIService.AlipayCommerceLogisticsOrderInstantdeliveryCancel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/logistics/order/instantdelivery/cancel"

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
	localVarPostBody = r.alipayCommerceLogisticsOrderInstantdeliveryCancelModel

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
		var v AlipayCommerceLogisticsOrderInstantdeliveryCancelDefaultResponse
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

type ApiAlipayCommerceLogisticsOrderInstantdeliveryCreateRequest struct {
	ctx                                                    context.Context
	ApiService                                             *AlipayCommerceLogisticsOrderInstantdeliveryAPIService
	alipayCommerceLogisticsOrderInstantdeliveryCreateModel *AlipayCommerceLogisticsOrderInstantdeliveryCreateModel
}

func (r ApiAlipayCommerceLogisticsOrderInstantdeliveryCreateRequest) AlipayCommerceLogisticsOrderInstantdeliveryCreateModel(alipayCommerceLogisticsOrderInstantdeliveryCreateModel AlipayCommerceLogisticsOrderInstantdeliveryCreateModel) ApiAlipayCommerceLogisticsOrderInstantdeliveryCreateRequest {
	r.alipayCommerceLogisticsOrderInstantdeliveryCreateModel = &alipayCommerceLogisticsOrderInstantdeliveryCreateModel
	return r
}

func (r ApiAlipayCommerceLogisticsOrderInstantdeliveryCreateRequest) Execute() (*AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel, *http.Response, error) {
	return r.ApiService.AlipayCommerceLogisticsOrderInstantdeliveryCreateExecute(r)
}

/*
AlipayCommerceLogisticsOrderInstantdeliveryCreate 下即时配送订单

下即时配送订单

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayCommerceLogisticsOrderInstantdeliveryCreateRequest
*/
func (r *AlipayCommerceLogisticsOrderInstantdeliveryAPIService) AlipayCommerceLogisticsOrderInstantdeliveryCreate(ctx context.Context) ApiAlipayCommerceLogisticsOrderInstantdeliveryCreateRequest {
	return ApiAlipayCommerceLogisticsOrderInstantdeliveryCreateRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel
func (a *AlipayCommerceLogisticsOrderInstantdeliveryAPIService) AlipayCommerceLogisticsOrderInstantdeliveryCreateExecute(r ApiAlipayCommerceLogisticsOrderInstantdeliveryCreateRequest) (*AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayCommerceLogisticsOrderInstantdeliveryCreateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceLogisticsOrderInstantdeliveryAPIService.AlipayCommerceLogisticsOrderInstantdeliveryCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/logistics/order/instantdelivery/create"

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
	localVarPostBody = r.alipayCommerceLogisticsOrderInstantdeliveryCreateModel

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
		var v AlipayCommerceLogisticsOrderInstantdeliveryCreateDefaultResponse
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

type ApiAlipayCommerceLogisticsOrderInstantdeliveryPrecreateRequest struct {
	ctx                                                       context.Context
	ApiService                                                *AlipayCommerceLogisticsOrderInstantdeliveryAPIService
	alipayCommerceLogisticsOrderInstantdeliveryPrecreateModel *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateModel
}

func (r ApiAlipayCommerceLogisticsOrderInstantdeliveryPrecreateRequest) AlipayCommerceLogisticsOrderInstantdeliveryPrecreateModel(alipayCommerceLogisticsOrderInstantdeliveryPrecreateModel AlipayCommerceLogisticsOrderInstantdeliveryPrecreateModel) ApiAlipayCommerceLogisticsOrderInstantdeliveryPrecreateRequest {
	r.alipayCommerceLogisticsOrderInstantdeliveryPrecreateModel = &alipayCommerceLogisticsOrderInstantdeliveryPrecreateModel
	return r
}

func (r ApiAlipayCommerceLogisticsOrderInstantdeliveryPrecreateRequest) Execute() (*AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel, *http.Response, error) {
	return r.ApiService.AlipayCommerceLogisticsOrderInstantdeliveryPrecreateExecute(r)
}

/*
AlipayCommerceLogisticsOrderInstantdeliveryPrecreate 预下即时配送订单

预下即时配送订单

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayCommerceLogisticsOrderInstantdeliveryPrecreateRequest
*/
func (r *AlipayCommerceLogisticsOrderInstantdeliveryAPIService) AlipayCommerceLogisticsOrderInstantdeliveryPrecreate(ctx context.Context) ApiAlipayCommerceLogisticsOrderInstantdeliveryPrecreateRequest {
	return ApiAlipayCommerceLogisticsOrderInstantdeliveryPrecreateRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel
func (a *AlipayCommerceLogisticsOrderInstantdeliveryAPIService) AlipayCommerceLogisticsOrderInstantdeliveryPrecreateExecute(r ApiAlipayCommerceLogisticsOrderInstantdeliveryPrecreateRequest) (*AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceLogisticsOrderInstantdeliveryAPIService.AlipayCommerceLogisticsOrderInstantdeliveryPrecreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/logistics/order/instantdelivery/precreate"

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
	localVarPostBody = r.alipayCommerceLogisticsOrderInstantdeliveryPrecreateModel

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
		var v AlipayCommerceLogisticsOrderInstantdeliveryPrecreateDefaultResponse
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

func (a *AlipayCommerceLogisticsOrderInstantdeliveryAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayCommerceLogisticsOrderInstantdeliveryAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
