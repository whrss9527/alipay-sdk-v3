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

// AlipayEcoEduKtBillingAPIService AlipayEcoEduKtBillingAPI service
type AlipayEcoEduKtBillingAPIService service

type ApiAlipayEcoEduKtBillingModifyRequest struct {
	ctx                              context.Context
	ApiService                       *AlipayEcoEduKtBillingAPIService
	alipayEcoEduKtBillingModifyModel *AlipayEcoEduKtBillingModifyModel
}

func (r ApiAlipayEcoEduKtBillingModifyRequest) AlipayEcoEduKtBillingModifyModel(alipayEcoEduKtBillingModifyModel AlipayEcoEduKtBillingModifyModel) ApiAlipayEcoEduKtBillingModifyRequest {
	r.alipayEcoEduKtBillingModifyModel = &alipayEcoEduKtBillingModifyModel
	return r
}

func (r ApiAlipayEcoEduKtBillingModifyRequest) Execute() (*AlipayEcoEduKtBillingModifyResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEcoEduKtBillingModifyExecute(r)
}

/*
AlipayEcoEduKtBillingModify 教育缴费账单状态同步接口

isv向支付宝－中小学－教育缴费发送教育缴费账单后，家长在支付宝－中小学－教育缴费中对账单完成支付操作且支付成功，ISV在自有系统内容对账，完成后通过此接口同步对账后的账单状态。状态分为：缴费成功、缴费失败。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEcoEduKtBillingModifyRequest
*/
func (r *AlipayEcoEduKtBillingAPIService) AlipayEcoEduKtBillingModify(ctx context.Context) ApiAlipayEcoEduKtBillingModifyRequest {
	return ApiAlipayEcoEduKtBillingModifyRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEcoEduKtBillingModifyResponseModel
func (a *AlipayEcoEduKtBillingAPIService) AlipayEcoEduKtBillingModifyExecute(r ApiAlipayEcoEduKtBillingModifyRequest) (*AlipayEcoEduKtBillingModifyResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEcoEduKtBillingModifyResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEcoEduKtBillingAPIService.AlipayEcoEduKtBillingModify")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/eco/edu/kt/billing/modify"

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
	localVarPostBody = r.alipayEcoEduKtBillingModifyModel

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
		var v AlipayEcoEduKtBillingModifyDefaultResponse
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

type ApiAlipayEcoEduKtBillingQueryRequest struct {
	ctx        context.Context
	ApiService *AlipayEcoEduKtBillingAPIService
	isvPid     *string
	schoolPid  *string
	outTradeNo *string
}

// Isv pid
func (r ApiAlipayEcoEduKtBillingQueryRequest) IsvPid(isvPid string) ApiAlipayEcoEduKtBillingQueryRequest {
	r.isvPid = &isvPid
	return r
}

// 学校支付宝pid，直付通填写smid
func (r ApiAlipayEcoEduKtBillingQueryRequest) SchoolPid(schoolPid string) ApiAlipayEcoEduKtBillingQueryRequest {
	r.schoolPid = &schoolPid
	return r
}

// ISV调用发送账单接口，返回给商户的order_no
func (r ApiAlipayEcoEduKtBillingQueryRequest) OutTradeNo(outTradeNo string) ApiAlipayEcoEduKtBillingQueryRequest {
	r.outTradeNo = &outTradeNo
	return r
}

func (r ApiAlipayEcoEduKtBillingQueryRequest) Execute() (*AlipayEcoEduKtBillingQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEcoEduKtBillingQueryExecute(r)
}

/*
AlipayEcoEduKtBillingQuery 缴费账单查询

缴费账单查询

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEcoEduKtBillingQueryRequest
*/
func (r *AlipayEcoEduKtBillingAPIService) AlipayEcoEduKtBillingQuery(ctx context.Context) ApiAlipayEcoEduKtBillingQueryRequest {
	return ApiAlipayEcoEduKtBillingQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEcoEduKtBillingQueryResponseModel
func (a *AlipayEcoEduKtBillingAPIService) AlipayEcoEduKtBillingQueryExecute(r ApiAlipayEcoEduKtBillingQueryRequest) (*AlipayEcoEduKtBillingQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEcoEduKtBillingQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEcoEduKtBillingAPIService.AlipayEcoEduKtBillingQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/eco/edu/kt/billing/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isvPid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isv_pid", r.isvPid, "form", "")
	}
	if r.schoolPid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "school_pid", r.schoolPid, "form", "")
	}
	if r.outTradeNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "out_trade_no", r.outTradeNo, "form", "")
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
		var v AlipayEcoEduKtBillingQueryDefaultResponse
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

type ApiAlipayEcoEduKtBillingSendRequest struct {
	ctx                            context.Context
	ApiService                     *AlipayEcoEduKtBillingAPIService
	alipayEcoEduKtBillingSendModel *AlipayEcoEduKtBillingSendModel
}

func (r ApiAlipayEcoEduKtBillingSendRequest) AlipayEcoEduKtBillingSendModel(alipayEcoEduKtBillingSendModel AlipayEcoEduKtBillingSendModel) ApiAlipayEcoEduKtBillingSendRequest {
	r.alipayEcoEduKtBillingSendModel = &alipayEcoEduKtBillingSendModel
	return r
}

func (r ApiAlipayEcoEduKtBillingSendRequest) Execute() (*AlipayEcoEduKtBillingSendResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEcoEduKtBillingSendExecute(r)
}

/*
AlipayEcoEduKtBillingSend 教育缴费账单发送接口

商家发送教育缴费账单给孩子，家长在支付宝可以查看自己绑定的孩子的缴费账单。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEcoEduKtBillingSendRequest
*/
func (r *AlipayEcoEduKtBillingAPIService) AlipayEcoEduKtBillingSend(ctx context.Context) ApiAlipayEcoEduKtBillingSendRequest {
	return ApiAlipayEcoEduKtBillingSendRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEcoEduKtBillingSendResponseModel
func (a *AlipayEcoEduKtBillingAPIService) AlipayEcoEduKtBillingSendExecute(r ApiAlipayEcoEduKtBillingSendRequest) (*AlipayEcoEduKtBillingSendResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEcoEduKtBillingSendResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEcoEduKtBillingAPIService.AlipayEcoEduKtBillingSend")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/eco/edu/kt/billing/send"

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
	localVarPostBody = r.alipayEcoEduKtBillingSendModel

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
		var v AlipayEcoEduKtBillingSendDefaultResponse
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

func (a *AlipayEcoEduKtBillingAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayEcoEduKtBillingAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
