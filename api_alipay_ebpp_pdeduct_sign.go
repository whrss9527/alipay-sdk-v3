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

// AlipayEbppPdeductSignAPIService AlipayEbppPdeductSignAPI service
type AlipayEbppPdeductSignAPIService service

type ApiAlipayEbppPdeductSignAddRequest struct {
	ctx                           context.Context
	ApiService                    *AlipayEbppPdeductSignAPIService
	alipayEbppPdeductSignAddModel *AlipayEbppPdeductSignAddModel
}

func (r ApiAlipayEbppPdeductSignAddRequest) AlipayEbppPdeductSignAddModel(alipayEbppPdeductSignAddModel AlipayEbppPdeductSignAddModel) ApiAlipayEbppPdeductSignAddRequest {
	r.alipayEbppPdeductSignAddModel = &alipayEbppPdeductSignAddModel
	return r
}

func (r ApiAlipayEbppPdeductSignAddRequest) Execute() (*AlipayEbppPdeductSignAddResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEbppPdeductSignAddExecute(r)
}

/*
AlipayEbppPdeductSignAdd 缴费直连代扣签约

缴费直连代扣签约

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEbppPdeductSignAddRequest
*/
func (r *AlipayEbppPdeductSignAPIService) AlipayEbppPdeductSignAdd(ctx context.Context) ApiAlipayEbppPdeductSignAddRequest {
	return ApiAlipayEbppPdeductSignAddRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEbppPdeductSignAddResponseModel
func (a *AlipayEbppPdeductSignAPIService) AlipayEbppPdeductSignAddExecute(r ApiAlipayEbppPdeductSignAddRequest) (*AlipayEbppPdeductSignAddResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEbppPdeductSignAddResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEbppPdeductSignAPIService.AlipayEbppPdeductSignAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/ebpp/pdeduct/sign/add"

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
	localVarPostBody = r.alipayEbppPdeductSignAddModel

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
		var v AlipayEbppPdeductSignAddDefaultResponse
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

type ApiAlipayEbppPdeductSignCancelRequest struct {
	ctx                              context.Context
	ApiService                       *AlipayEbppPdeductSignAPIService
	alipayEbppPdeductSignCancelModel *AlipayEbppPdeductSignCancelModel
}

func (r ApiAlipayEbppPdeductSignCancelRequest) AlipayEbppPdeductSignCancelModel(alipayEbppPdeductSignCancelModel AlipayEbppPdeductSignCancelModel) ApiAlipayEbppPdeductSignCancelRequest {
	r.alipayEbppPdeductSignCancelModel = &alipayEbppPdeductSignCancelModel
	return r
}

func (r ApiAlipayEbppPdeductSignCancelRequest) Execute() (*AlipayEbppPdeductSignCancelResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEbppPdeductSignCancelExecute(r)
}

/*
AlipayEbppPdeductSignCancel 缴费直连代扣取消签约

缴费直连代扣，用户取消签约接口

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEbppPdeductSignCancelRequest
*/
func (r *AlipayEbppPdeductSignAPIService) AlipayEbppPdeductSignCancel(ctx context.Context) ApiAlipayEbppPdeductSignCancelRequest {
	return ApiAlipayEbppPdeductSignCancelRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEbppPdeductSignCancelResponseModel
func (a *AlipayEbppPdeductSignAPIService) AlipayEbppPdeductSignCancelExecute(r ApiAlipayEbppPdeductSignCancelRequest) (*AlipayEbppPdeductSignCancelResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEbppPdeductSignCancelResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEbppPdeductSignAPIService.AlipayEbppPdeductSignCancel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/ebpp/pdeduct/sign/cancel"

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
	localVarPostBody = r.alipayEbppPdeductSignCancelModel

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
		var v AlipayEbppPdeductSignCancelDefaultResponse
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

type ApiAlipayEbppPdeductSignQueryRequest struct {
	ctx         context.Context
	ApiService  *AlipayEbppPdeductSignAPIService
	userId      *string
	openId      *string
	agreementId *string
	bizType     *string
	subBizType  *string
	chargeInst  *string
	billKey     *string
}

// 用户ID
func (r ApiAlipayEbppPdeductSignQueryRequest) UserId(userId string) ApiAlipayEbppPdeductSignQueryRequest {
	r.userId = &userId
	return r
}

// 用户UserId在应用AppId下的唯一用户标识
func (r ApiAlipayEbppPdeductSignQueryRequest) OpenId(openId string) ApiAlipayEbppPdeductSignQueryRequest {
	r.openId = &openId
	return r
}

// 支付宝代扣协议Id。若协议id不传递，则需要保证业务类型、子业务类型、出账机构、户号必传
func (r ApiAlipayEbppPdeductSignQueryRequest) AgreementId(agreementId string) ApiAlipayEbppPdeductSignQueryRequest {
	r.agreementId = &agreementId
	return r
}

// 业务类型。  JF：缴水、电、燃气、固话宽带、有线电视、交通罚款费用  WUYE：缴物业费  HK：信用卡还款  TX：手机充值
func (r ApiAlipayEbppPdeductSignQueryRequest) BizType(bizType string) ApiAlipayEbppPdeductSignQueryRequest {
	r.bizType = &bizType
	return r
}

// 业务子类型。  WATER：缴水费  ELECTRIC：缴电费  GAS：缴燃气费  COMMUN：缴固话宽带  CATV：缴有线电视费  TRAFFIC：缴交通罚款  WUYE：缴物业费  HK：信用卡还款  CZ：手机充值
func (r ApiAlipayEbppPdeductSignQueryRequest) SubBizType(subBizType string) ApiAlipayEbppPdeductSignQueryRequest {
	r.subBizType = &subBizType
	return r
}

// 支付宝缴费系统中的出账机构ID
func (r ApiAlipayEbppPdeductSignQueryRequest) ChargeInst(chargeInst string) ApiAlipayEbppPdeductSignQueryRequest {
	r.chargeInst = &chargeInst
	return r
}

// 户号，机构针对于每户的水、电都会有唯一的标识户号
func (r ApiAlipayEbppPdeductSignQueryRequest) BillKey(billKey string) ApiAlipayEbppPdeductSignQueryRequest {
	r.billKey = &billKey
	return r
}

func (r ApiAlipayEbppPdeductSignQueryRequest) Execute() (*AlipayEbppPdeductSignQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEbppPdeductSignQueryExecute(r)
}

/*
AlipayEbppPdeductSignQuery 直连代扣协议查询接口

提供给朗新查询代扣协议状态的接口，用于进行双边对账，解决单边协议问题

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEbppPdeductSignQueryRequest
*/
func (r *AlipayEbppPdeductSignAPIService) AlipayEbppPdeductSignQuery(ctx context.Context) ApiAlipayEbppPdeductSignQueryRequest {
	return ApiAlipayEbppPdeductSignQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEbppPdeductSignQueryResponseModel
func (a *AlipayEbppPdeductSignAPIService) AlipayEbppPdeductSignQueryExecute(r ApiAlipayEbppPdeductSignQueryRequest) (*AlipayEbppPdeductSignQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEbppPdeductSignQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEbppPdeductSignAPIService.AlipayEbppPdeductSignQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/ebpp/pdeduct/sign/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "user_id", r.userId, "form", "")
	}
	if r.openId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "open_id", r.openId, "form", "")
	}
	if r.agreementId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "agreement_id", r.agreementId, "form", "")
	}
	if r.bizType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "biz_type", r.bizType, "form", "")
	}
	if r.subBizType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sub_biz_type", r.subBizType, "form", "")
	}
	if r.chargeInst != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "charge_inst", r.chargeInst, "form", "")
	}
	if r.billKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "bill_key", r.billKey, "form", "")
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
		var v AlipayEbppPdeductSignQueryDefaultResponse
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

type ApiAlipayEbppPdeductSignValidateRequest struct {
	ctx                                context.Context
	ApiService                         *AlipayEbppPdeductSignAPIService
	alipayEbppPdeductSignValidateModel *AlipayEbppPdeductSignValidateModel
}

func (r ApiAlipayEbppPdeductSignValidateRequest) AlipayEbppPdeductSignValidateModel(alipayEbppPdeductSignValidateModel AlipayEbppPdeductSignValidateModel) ApiAlipayEbppPdeductSignValidateRequest {
	r.alipayEbppPdeductSignValidateModel = &alipayEbppPdeductSignValidateModel
	return r
}

func (r ApiAlipayEbppPdeductSignValidateRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayEbppPdeductSignValidateExecute(r)
}

/*
AlipayEbppPdeductSignValidate 缴费直连代扣签约前置校验

服务窗代扣签约的前置校验接口，签约前，可以先调用该接口进行校验是否可以签约

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEbppPdeductSignValidateRequest
*/
func (r *AlipayEbppPdeductSignAPIService) AlipayEbppPdeductSignValidate(ctx context.Context) ApiAlipayEbppPdeductSignValidateRequest {
	return ApiAlipayEbppPdeductSignValidateRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *AlipayEbppPdeductSignAPIService) AlipayEbppPdeductSignValidateExecute(r ApiAlipayEbppPdeductSignValidateRequest) (map[string]interface{}, *http.Response, error) {
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

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEbppPdeductSignAPIService.AlipayEbppPdeductSignValidate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/ebpp/pdeduct/sign/validate"

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
	localVarPostBody = r.alipayEbppPdeductSignValidateModel

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
		var v AlipayEbppPdeductSignValidateDefaultResponse
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

func (a *AlipayEbppPdeductSignAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayEbppPdeductSignAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
