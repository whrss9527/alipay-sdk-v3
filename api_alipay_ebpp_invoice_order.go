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

// AlipayEbppInvoiceOrderAPIService AlipayEbppInvoiceOrderAPI service
type AlipayEbppInvoiceOrderAPIService service

type ApiAlipayEbppInvoiceOrderQueryRequest struct {
	ctx           context.Context
	ApiService    *AlipayEbppInvoiceOrderAPIService
	orderNo       *string
	mShortName    *string
	subMShortName *string
}

// 开票申请时所传入订单号，不限于支付宝体内交易订单号。如：20200520110046966071
func (r ApiAlipayEbppInvoiceOrderQueryRequest) OrderNo(orderNo string) ApiAlipayEbppInvoiceOrderQueryRequest {
	r.orderNo = &orderNo
	return r
}

// 定义商户的一级简称,用于标识商户品牌，对应于商户入驻时填写的\&quot;商户品牌简称\&quot;。 如：肯德基：KFC
func (r ApiAlipayEbppInvoiceOrderQueryRequest) MShortName(mShortName string) ApiAlipayEbppInvoiceOrderQueryRequest {
	r.mShortName = &mShortName
	return r
}

// 定义商户的二级简称,用于标识商户品牌下的分支机构，如门店，对应于商户入驻时填写的\&quot;商户门店简称\&quot;。 如：肯德基-杭州西湖区文一西路店：KFC-HZ-19003 要求：\&quot;商户品牌简称+商户门店简称\&quot;作为确定商户及其下属机构的唯一标识，不可重复。
func (r ApiAlipayEbppInvoiceOrderQueryRequest) SubMShortName(subMShortName string) ApiAlipayEbppInvoiceOrderQueryRequest {
	r.subMShortName = &subMShortName
	return r
}

func (r ApiAlipayEbppInvoiceOrderQueryRequest) Execute() (*AlipayEbppInvoiceOrderQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEbppInvoiceOrderQueryExecute(r)
}

/*
AlipayEbppInvoiceOrderQuery 根据外部订单号查询发票信息

根据外部订单号查询发票详情信息，适用于外部商户无开票申请ID场景

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEbppInvoiceOrderQueryRequest
*/
func (r *AlipayEbppInvoiceOrderAPIService) AlipayEbppInvoiceOrderQuery(ctx context.Context) ApiAlipayEbppInvoiceOrderQueryRequest {
	return ApiAlipayEbppInvoiceOrderQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEbppInvoiceOrderQueryResponseModel
func (a *AlipayEbppInvoiceOrderAPIService) AlipayEbppInvoiceOrderQueryExecute(r ApiAlipayEbppInvoiceOrderQueryRequest) (*AlipayEbppInvoiceOrderQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEbppInvoiceOrderQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEbppInvoiceOrderAPIService.AlipayEbppInvoiceOrderQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/ebpp/invoice/order/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.orderNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_no", r.orderNo, "form", "")
	}
	if r.mShortName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "m_short_name", r.mShortName, "form", "")
	}
	if r.subMShortName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sub_m_short_name", r.subMShortName, "form", "")
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
		var v AlipayEbppInvoiceOrderQueryDefaultResponse
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

func (a *AlipayEbppInvoiceOrderAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayEbppInvoiceOrderAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
