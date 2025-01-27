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

// AlipayOpenInstantdeliveryMerchantshopAPIService AlipayOpenInstantdeliveryMerchantshopAPI service
type AlipayOpenInstantdeliveryMerchantshopAPIService service

type ApiAlipayOpenInstantdeliveryMerchantshopBatchqueryRequest struct {
	ctx                                                  context.Context
	ApiService                                           *AlipayOpenInstantdeliveryMerchantshopAPIService
	alipayOpenInstantdeliveryMerchantshopBatchqueryModel *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel
}

func (r ApiAlipayOpenInstantdeliveryMerchantshopBatchqueryRequest) AlipayOpenInstantdeliveryMerchantshopBatchqueryModel(alipayOpenInstantdeliveryMerchantshopBatchqueryModel AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) ApiAlipayOpenInstantdeliveryMerchantshopBatchqueryRequest {
	r.alipayOpenInstantdeliveryMerchantshopBatchqueryModel = &alipayOpenInstantdeliveryMerchantshopBatchqueryModel
	return r
}

func (r ApiAlipayOpenInstantdeliveryMerchantshopBatchqueryRequest) Execute() (*AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenInstantdeliveryMerchantshopBatchqueryExecute(r)
}

/*
AlipayOpenInstantdeliveryMerchantshopBatchquery 即时配送商家门店分页查询

即时配送场景，商家门店信息分页查询

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayOpenInstantdeliveryMerchantshopBatchqueryRequest
*/
func (r *AlipayOpenInstantdeliveryMerchantshopAPIService) AlipayOpenInstantdeliveryMerchantshopBatchquery(ctx context.Context) ApiAlipayOpenInstantdeliveryMerchantshopBatchqueryRequest {
	return ApiAlipayOpenInstantdeliveryMerchantshopBatchqueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel
func (a *AlipayOpenInstantdeliveryMerchantshopAPIService) AlipayOpenInstantdeliveryMerchantshopBatchqueryExecute(r ApiAlipayOpenInstantdeliveryMerchantshopBatchqueryRequest) (*AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenInstantdeliveryMerchantshopAPIService.AlipayOpenInstantdeliveryMerchantshopBatchquery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/instantdelivery/merchantshop/batchquery"

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
	localVarPostBody = r.alipayOpenInstantdeliveryMerchantshopBatchqueryModel

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
		var v AlipayOpenInstantdeliveryMerchantshopBatchqueryDefaultResponse
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

type ApiAlipayOpenInstantdeliveryMerchantshopCreateRequest struct {
	ctx                                              context.Context
	ApiService                                       *AlipayOpenInstantdeliveryMerchantshopAPIService
	alipayOpenInstantdeliveryMerchantshopCreateModel *AlipayOpenInstantdeliveryMerchantshopCreateModel
}

func (r ApiAlipayOpenInstantdeliveryMerchantshopCreateRequest) AlipayOpenInstantdeliveryMerchantshopCreateModel(alipayOpenInstantdeliveryMerchantshopCreateModel AlipayOpenInstantdeliveryMerchantshopCreateModel) ApiAlipayOpenInstantdeliveryMerchantshopCreateRequest {
	r.alipayOpenInstantdeliveryMerchantshopCreateModel = &alipayOpenInstantdeliveryMerchantshopCreateModel
	return r
}

func (r ApiAlipayOpenInstantdeliveryMerchantshopCreateRequest) Execute() (*AlipayOpenInstantdeliveryMerchantshopCreateResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenInstantdeliveryMerchantshopCreateExecute(r)
}

/*
AlipayOpenInstantdeliveryMerchantshopCreate 即时配送商家门店创建

帮商家在即时配送公司创建门店。名称+地址必须全局唯一；pid下的shopNo必须唯一。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayOpenInstantdeliveryMerchantshopCreateRequest
*/
func (r *AlipayOpenInstantdeliveryMerchantshopAPIService) AlipayOpenInstantdeliveryMerchantshopCreate(ctx context.Context) ApiAlipayOpenInstantdeliveryMerchantshopCreateRequest {
	return ApiAlipayOpenInstantdeliveryMerchantshopCreateRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayOpenInstantdeliveryMerchantshopCreateResponseModel
func (a *AlipayOpenInstantdeliveryMerchantshopAPIService) AlipayOpenInstantdeliveryMerchantshopCreateExecute(r ApiAlipayOpenInstantdeliveryMerchantshopCreateRequest) (*AlipayOpenInstantdeliveryMerchantshopCreateResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayOpenInstantdeliveryMerchantshopCreateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenInstantdeliveryMerchantshopAPIService.AlipayOpenInstantdeliveryMerchantshopCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/instantdelivery/merchantshop/create"

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
	localVarPostBody = r.alipayOpenInstantdeliveryMerchantshopCreateModel

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
		var v AlipayOpenInstantdeliveryMerchantshopCreateDefaultResponse
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

type ApiAlipayOpenInstantdeliveryMerchantshopModifyRequest struct {
	ctx                                              context.Context
	ApiService                                       *AlipayOpenInstantdeliveryMerchantshopAPIService
	alipayOpenInstantdeliveryMerchantshopModifyModel *AlipayOpenInstantdeliveryMerchantshopModifyModel
}

func (r ApiAlipayOpenInstantdeliveryMerchantshopModifyRequest) AlipayOpenInstantdeliveryMerchantshopModifyModel(alipayOpenInstantdeliveryMerchantshopModifyModel AlipayOpenInstantdeliveryMerchantshopModifyModel) ApiAlipayOpenInstantdeliveryMerchantshopModifyRequest {
	r.alipayOpenInstantdeliveryMerchantshopModifyModel = &alipayOpenInstantdeliveryMerchantshopModifyModel
	return r
}

func (r ApiAlipayOpenInstantdeliveryMerchantshopModifyRequest) Execute() (*AlipayOpenInstantdeliveryMerchantshopModifyResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenInstantdeliveryMerchantshopModifyExecute(r)
}

/*
AlipayOpenInstantdeliveryMerchantshopModify 即时配送商家门店更新

帮商家更新在即时配送公司的门店

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayOpenInstantdeliveryMerchantshopModifyRequest
*/
func (r *AlipayOpenInstantdeliveryMerchantshopAPIService) AlipayOpenInstantdeliveryMerchantshopModify(ctx context.Context) ApiAlipayOpenInstantdeliveryMerchantshopModifyRequest {
	return ApiAlipayOpenInstantdeliveryMerchantshopModifyRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayOpenInstantdeliveryMerchantshopModifyResponseModel
func (a *AlipayOpenInstantdeliveryMerchantshopAPIService) AlipayOpenInstantdeliveryMerchantshopModifyExecute(r ApiAlipayOpenInstantdeliveryMerchantshopModifyRequest) (*AlipayOpenInstantdeliveryMerchantshopModifyResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayOpenInstantdeliveryMerchantshopModifyResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenInstantdeliveryMerchantshopAPIService.AlipayOpenInstantdeliveryMerchantshopModify")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/instantdelivery/merchantshop/modify"

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
	localVarPostBody = r.alipayOpenInstantdeliveryMerchantshopModifyModel

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
		var v AlipayOpenInstantdeliveryMerchantshopModifyDefaultResponse
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

type ApiAlipayOpenInstantdeliveryMerchantshopQueryRequest struct {
	ctx        context.Context
	ApiService *AlipayOpenInstantdeliveryMerchantshopAPIService
	shopNo     *string
}

// 商家门店编码。
func (r ApiAlipayOpenInstantdeliveryMerchantshopQueryRequest) ShopNo(shopNo string) ApiAlipayOpenInstantdeliveryMerchantshopQueryRequest {
	r.shopNo = &shopNo
	return r
}

func (r ApiAlipayOpenInstantdeliveryMerchantshopQueryRequest) Execute() (*AlipayOpenInstantdeliveryMerchantshopQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenInstantdeliveryMerchantshopQueryExecute(r)
}

/*
AlipayOpenInstantdeliveryMerchantshopQuery 即时配送商家门店详情查询

即时配送场景，查询商家门店详情信息。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayOpenInstantdeliveryMerchantshopQueryRequest
*/
func (r *AlipayOpenInstantdeliveryMerchantshopAPIService) AlipayOpenInstantdeliveryMerchantshopQuery(ctx context.Context) ApiAlipayOpenInstantdeliveryMerchantshopQueryRequest {
	return ApiAlipayOpenInstantdeliveryMerchantshopQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayOpenInstantdeliveryMerchantshopQueryResponseModel
func (a *AlipayOpenInstantdeliveryMerchantshopAPIService) AlipayOpenInstantdeliveryMerchantshopQueryExecute(r ApiAlipayOpenInstantdeliveryMerchantshopQueryRequest) (*AlipayOpenInstantdeliveryMerchantshopQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayOpenInstantdeliveryMerchantshopQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenInstantdeliveryMerchantshopAPIService.AlipayOpenInstantdeliveryMerchantshopQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/instantdelivery/merchantshop/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.shopNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "shop_no", r.shopNo, "form", "")
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
		var v AlipayOpenInstantdeliveryMerchantshopQueryDefaultResponse
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

func (a *AlipayOpenInstantdeliveryMerchantshopAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayOpenInstantdeliveryMerchantshopAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
