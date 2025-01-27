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

// AlipayFundEnterprisepayGroupAPIService AlipayFundEnterprisepayGroupAPI service
type AlipayFundEnterprisepayGroupAPIService service

type ApiAlipayFundEnterprisepayGroupAddRequest struct {
	ctx                                  context.Context
	ApiService                           *AlipayFundEnterprisepayGroupAPIService
	alipayFundEnterprisepayGroupAddModel *AlipayFundEnterprisepayGroupAddModel
}

func (r ApiAlipayFundEnterprisepayGroupAddRequest) AlipayFundEnterprisepayGroupAddModel(alipayFundEnterprisepayGroupAddModel AlipayFundEnterprisepayGroupAddModel) ApiAlipayFundEnterprisepayGroupAddRequest {
	r.alipayFundEnterprisepayGroupAddModel = &alipayFundEnterprisepayGroupAddModel
	return r
}

func (r ApiAlipayFundEnterprisepayGroupAddRequest) Execute() (*AlipayFundEnterprisepayGroupAddResponseModel, *http.Response, error) {
	return r.ApiService.AlipayFundEnterprisepayGroupAddExecute(r)
}

/*
AlipayFundEnterprisepayGroupAdd 因公付新增账户下群组

创建因公付群组，包含群组信息和出资信息

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayFundEnterprisepayGroupAddRequest
*/
func (r *AlipayFundEnterprisepayGroupAPIService) AlipayFundEnterprisepayGroupAdd(ctx context.Context) ApiAlipayFundEnterprisepayGroupAddRequest {
	return ApiAlipayFundEnterprisepayGroupAddRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayFundEnterprisepayGroupAddResponseModel
func (a *AlipayFundEnterprisepayGroupAPIService) AlipayFundEnterprisepayGroupAddExecute(r ApiAlipayFundEnterprisepayGroupAddRequest) (*AlipayFundEnterprisepayGroupAddResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayFundEnterprisepayGroupAddResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayFundEnterprisepayGroupAPIService.AlipayFundEnterprisepayGroupAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/fund/enterprisepay/group/add"

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
	localVarPostBody = r.alipayFundEnterprisepayGroupAddModel

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
		var v AlipayFundEnterprisepayGroupAddDefaultResponse
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

type ApiAlipayFundEnterprisepayGroupDeleteRequest struct {
	ctx         context.Context
	ApiService  *AlipayFundEnterprisepayGroupAPIService
	productCode *string
	bizScene    *string
	accountId   *string
	agreementNo *string
	outBizNo    *string
}

// 产品码，默认值 ENTERPRISE_PAY
func (r ApiAlipayFundEnterprisepayGroupDeleteRequest) ProductCode(productCode string) ApiAlipayFundEnterprisepayGroupDeleteRequest {
	r.productCode = &productCode
	return r
}

// 场景码，联系支付宝分配
func (r ApiAlipayFundEnterprisepayGroupDeleteRequest) BizScene(bizScene string) ApiAlipayFundEnterprisepayGroupDeleteRequest {
	r.bizScene = &bizScene
	return r
}

// 企业签约账户ID
func (r ApiAlipayFundEnterprisepayGroupDeleteRequest) AccountId(accountId string) ApiAlipayFundEnterprisepayGroupDeleteRequest {
	r.accountId = &accountId
	return r
}

// 平台和企业的三方授权协议号
func (r ApiAlipayFundEnterprisepayGroupDeleteRequest) AgreementNo(agreementNo string) ApiAlipayFundEnterprisepayGroupDeleteRequest {
	r.agreementNo = &agreementNo
	return r
}

// 外部业务号
func (r ApiAlipayFundEnterprisepayGroupDeleteRequest) OutBizNo(outBizNo string) ApiAlipayFundEnterprisepayGroupDeleteRequest {
	r.outBizNo = &outBizNo
	return r
}

func (r ApiAlipayFundEnterprisepayGroupDeleteRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayFundEnterprisepayGroupDeleteExecute(r)
}

/*
AlipayFundEnterprisepayGroupDelete 因公付删除账户下群组

因公付删除账户下群组

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayFundEnterprisepayGroupDeleteRequest
*/
func (r *AlipayFundEnterprisepayGroupAPIService) AlipayFundEnterprisepayGroupDelete(ctx context.Context) ApiAlipayFundEnterprisepayGroupDeleteRequest {
	return ApiAlipayFundEnterprisepayGroupDeleteRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *AlipayFundEnterprisepayGroupAPIService) AlipayFundEnterprisepayGroupDeleteExecute(r ApiAlipayFundEnterprisepayGroupDeleteRequest) (map[string]interface{}, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayFundEnterprisepayGroupAPIService.AlipayFundEnterprisepayGroupDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/fund/enterprisepay/group/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.productCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_code", r.productCode, "form", "")
	}
	if r.bizScene != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "biz_scene", r.bizScene, "form", "")
	}
	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId, "form", "")
	}
	if r.agreementNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "agreement_no", r.agreementNo, "form", "")
	}
	if r.outBizNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "out_biz_no", r.outBizNo, "form", "")
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
		var v AlipayFundEnterprisepayGroupDeleteDefaultResponse
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

type ApiAlipayFundEnterprisepayGroupModifyRequest struct {
	ctx                                     context.Context
	ApiService                              *AlipayFundEnterprisepayGroupAPIService
	alipayFundEnterprisepayGroupModifyModel *AlipayFundEnterprisepayGroupModifyModel
}

func (r ApiAlipayFundEnterprisepayGroupModifyRequest) AlipayFundEnterprisepayGroupModifyModel(alipayFundEnterprisepayGroupModifyModel AlipayFundEnterprisepayGroupModifyModel) ApiAlipayFundEnterprisepayGroupModifyRequest {
	r.alipayFundEnterprisepayGroupModifyModel = &alipayFundEnterprisepayGroupModifyModel
	return r
}

func (r ApiAlipayFundEnterprisepayGroupModifyRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayFundEnterprisepayGroupModifyExecute(r)
}

/*
AlipayFundEnterprisepayGroupModify 因公付更新账户下群组

因公付更新账户下群组

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayFundEnterprisepayGroupModifyRequest
*/
func (r *AlipayFundEnterprisepayGroupAPIService) AlipayFundEnterprisepayGroupModify(ctx context.Context) ApiAlipayFundEnterprisepayGroupModifyRequest {
	return ApiAlipayFundEnterprisepayGroupModifyRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *AlipayFundEnterprisepayGroupAPIService) AlipayFundEnterprisepayGroupModifyExecute(r ApiAlipayFundEnterprisepayGroupModifyRequest) (map[string]interface{}, *http.Response, error) {
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

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayFundEnterprisepayGroupAPIService.AlipayFundEnterprisepayGroupModify")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/fund/enterprisepay/group/modify"

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
	localVarPostBody = r.alipayFundEnterprisepayGroupModifyModel

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
		var v AlipayFundEnterprisepayGroupModifyDefaultResponse
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

type ApiAlipayFundEnterprisepayGroupQueryRequest struct {
	ctx         context.Context
	ApiService  *AlipayFundEnterprisepayGroupAPIService
	productCode *string
	bizScene    *string
	accountId   *string
	agreementNo *string
	outBizNo    *string
}

// 产品码，默认值 ENTERPRISE_PAY
func (r ApiAlipayFundEnterprisepayGroupQueryRequest) ProductCode(productCode string) ApiAlipayFundEnterprisepayGroupQueryRequest {
	r.productCode = &productCode
	return r
}

// 场景码，联系支付宝分配
func (r ApiAlipayFundEnterprisepayGroupQueryRequest) BizScene(bizScene string) ApiAlipayFundEnterprisepayGroupQueryRequest {
	r.bizScene = &bizScene
	return r
}

// 企业签约账户ID
func (r ApiAlipayFundEnterprisepayGroupQueryRequest) AccountId(accountId string) ApiAlipayFundEnterprisepayGroupQueryRequest {
	r.accountId = &accountId
	return r
}

// 平台和企业的三方授权协议号
func (r ApiAlipayFundEnterprisepayGroupQueryRequest) AgreementNo(agreementNo string) ApiAlipayFundEnterprisepayGroupQueryRequest {
	r.agreementNo = &agreementNo
	return r
}

// 外部业务号，外部群组号
func (r ApiAlipayFundEnterprisepayGroupQueryRequest) OutBizNo(outBizNo string) ApiAlipayFundEnterprisepayGroupQueryRequest {
	r.outBizNo = &outBizNo
	return r
}

func (r ApiAlipayFundEnterprisepayGroupQueryRequest) Execute() (*AlipayFundEnterprisepayGroupQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayFundEnterprisepayGroupQueryExecute(r)
}

/*
AlipayFundEnterprisepayGroupQuery 因公付查询账户下群组

查询因公付群组，包含群组信息和对应的出资主体

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayFundEnterprisepayGroupQueryRequest
*/
func (r *AlipayFundEnterprisepayGroupAPIService) AlipayFundEnterprisepayGroupQuery(ctx context.Context) ApiAlipayFundEnterprisepayGroupQueryRequest {
	return ApiAlipayFundEnterprisepayGroupQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayFundEnterprisepayGroupQueryResponseModel
func (a *AlipayFundEnterprisepayGroupAPIService) AlipayFundEnterprisepayGroupQueryExecute(r ApiAlipayFundEnterprisepayGroupQueryRequest) (*AlipayFundEnterprisepayGroupQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayFundEnterprisepayGroupQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayFundEnterprisepayGroupAPIService.AlipayFundEnterprisepayGroupQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/fund/enterprisepay/group/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.productCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_code", r.productCode, "form", "")
	}
	if r.bizScene != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "biz_scene", r.bizScene, "form", "")
	}
	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId, "form", "")
	}
	if r.agreementNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "agreement_no", r.agreementNo, "form", "")
	}
	if r.outBizNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "out_biz_no", r.outBizNo, "form", "")
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
		var v AlipayFundEnterprisepayGroupQueryDefaultResponse
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

func (a *AlipayFundEnterprisepayGroupAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayFundEnterprisepayGroupAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
