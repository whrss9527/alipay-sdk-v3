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

// AlipayPcreditHuabeiAuthOrderAPIService AlipayPcreditHuabeiAuthOrderAPI service
type AlipayPcreditHuabeiAuthOrderAPIService service

type ApiAlipayPcreditHuabeiAuthOrderQueryRequest struct {
	ctx          context.Context
	ApiService   *AlipayPcreditHuabeiAuthOrderAPIService
	authOptId    *string
	alipayUserId *string
	openId       *string
	outRequestNo *string
}

// 支付宝侧花呗冻结、解冻操作单据id。在原先的冻结或者解冻接口调用中同步返回给商户，或者通过商户通知返回给商户。按订单号查询时，此字段不可为空。
func (r ApiAlipayPcreditHuabeiAuthOrderQueryRequest) AuthOptId(authOptId string) ApiAlipayPcreditHuabeiAuthOrderQueryRequest {
	r.authOptId = &authOptId
	return r
}

// 买家在支付宝的用户id。通过userid+请求流水号组合查询时，此字段不可为空。
func (r ApiAlipayPcreditHuabeiAuthOrderQueryRequest) AlipayUserId(alipayUserId string) ApiAlipayPcreditHuabeiAuthOrderQueryRequest {
	r.alipayUserId = &alipayUserId
	return r
}

// 买家在支付宝的用户id
func (r ApiAlipayPcreditHuabeiAuthOrderQueryRequest) OpenId(openId string) ApiAlipayPcreditHuabeiAuthOrderQueryRequest {
	r.openId = &openId
	return r
}

// 商户原先调用冻结、解冻接口传入的请求流水号。按照流水号查询订单时，此字段不能为空。
func (r ApiAlipayPcreditHuabeiAuthOrderQueryRequest) OutRequestNo(outRequestNo string) ApiAlipayPcreditHuabeiAuthOrderQueryRequest {
	r.outRequestNo = &outRequestNo
	return r
}

func (r ApiAlipayPcreditHuabeiAuthOrderQueryRequest) Execute() (*AlipayPcreditHuabeiAuthOrderQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayPcreditHuabeiAuthOrderQueryExecute(r)
}

/*
AlipayPcreditHuabeiAuthOrderQuery 花呗先享订单查询接口

查询花呗先享冻结、解冻订单内容及状态。有3种查询方式。推荐商户优先使用auth_opt_id查询；其次是按照(alipay_user_id,out_request_no)组合方式查询；最后是单独通过out_request_no方式查询。
注意：最后一种方式，仅支持2019年2月15日开始的订单。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayPcreditHuabeiAuthOrderQueryRequest
*/
func (r *AlipayPcreditHuabeiAuthOrderAPIService) AlipayPcreditHuabeiAuthOrderQuery(ctx context.Context) ApiAlipayPcreditHuabeiAuthOrderQueryRequest {
	return ApiAlipayPcreditHuabeiAuthOrderQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayPcreditHuabeiAuthOrderQueryResponseModel
func (a *AlipayPcreditHuabeiAuthOrderAPIService) AlipayPcreditHuabeiAuthOrderQueryExecute(r ApiAlipayPcreditHuabeiAuthOrderQueryRequest) (*AlipayPcreditHuabeiAuthOrderQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayPcreditHuabeiAuthOrderQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayPcreditHuabeiAuthOrderAPIService.AlipayPcreditHuabeiAuthOrderQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/pcredit/huabei/auth/order/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.authOptId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "auth_opt_id", r.authOptId, "form", "")
	}
	if r.alipayUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "alipay_user_id", r.alipayUserId, "form", "")
	}
	if r.openId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "open_id", r.openId, "form", "")
	}
	if r.outRequestNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "out_request_no", r.outRequestNo, "form", "")
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
		var v AlipayPcreditHuabeiAuthOrderQueryDefaultResponse
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

type ApiAlipayPcreditHuabeiAuthOrderUnfreezeRequest struct {
	ctx                                       context.Context
	ApiService                                *AlipayPcreditHuabeiAuthOrderAPIService
	alipayPcreditHuabeiAuthOrderUnfreezeModel *AlipayPcreditHuabeiAuthOrderUnfreezeModel
}

func (r ApiAlipayPcreditHuabeiAuthOrderUnfreezeRequest) AlipayPcreditHuabeiAuthOrderUnfreezeModel(alipayPcreditHuabeiAuthOrderUnfreezeModel AlipayPcreditHuabeiAuthOrderUnfreezeModel) ApiAlipayPcreditHuabeiAuthOrderUnfreezeRequest {
	r.alipayPcreditHuabeiAuthOrderUnfreezeModel = &alipayPcreditHuabeiAuthOrderUnfreezeModel
	return r
}

func (r ApiAlipayPcreditHuabeiAuthOrderUnfreezeRequest) Execute() (*AlipayPcreditHuabeiAuthOrderUnfreezeResponseModel, *http.Response, error) {
	return r.ApiService.AlipayPcreditHuabeiAuthOrderUnfreezeExecute(r)
}

/*
AlipayPcreditHuabeiAuthOrderUnfreeze 花呗先享解冻或解约接口

用户已经开通花呗先享协议后，商户通过此接口解冻用户资金池金额，也可以解冻并解约。
如果是解约操作，则要求传入的解冻金额必须等于用户资金池余额。
注意：商户在发起解约前，请务必保证已经结算过用户会员费，一旦解约后，无法发起结算用户会员费操作。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayPcreditHuabeiAuthOrderUnfreezeRequest
*/
func (r *AlipayPcreditHuabeiAuthOrderAPIService) AlipayPcreditHuabeiAuthOrderUnfreeze(ctx context.Context) ApiAlipayPcreditHuabeiAuthOrderUnfreezeRequest {
	return ApiAlipayPcreditHuabeiAuthOrderUnfreezeRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayPcreditHuabeiAuthOrderUnfreezeResponseModel
func (a *AlipayPcreditHuabeiAuthOrderAPIService) AlipayPcreditHuabeiAuthOrderUnfreezeExecute(r ApiAlipayPcreditHuabeiAuthOrderUnfreezeRequest) (*AlipayPcreditHuabeiAuthOrderUnfreezeResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayPcreditHuabeiAuthOrderUnfreezeResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayPcreditHuabeiAuthOrderAPIService.AlipayPcreditHuabeiAuthOrderUnfreeze")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/pcredit/huabei/auth/order/unfreeze"

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
	localVarPostBody = r.alipayPcreditHuabeiAuthOrderUnfreezeModel

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
		var v AlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse
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

func (a *AlipayPcreditHuabeiAuthOrderAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayPcreditHuabeiAuthOrderAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
