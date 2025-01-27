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

// AlipayEbppInvoiceEcorderOrderAPIService AlipayEbppInvoiceEcorderOrderAPI service
type AlipayEbppInvoiceEcorderOrderAPIService service

type ApiAlipayEbppInvoiceEcorderOrderQueryRequest struct {
	ctx          context.Context
	ApiService   *AlipayEbppInvoiceEcorderOrderAPIService
	enterpriseId *string
	accountId    *string
	agreementNo  *string
	orderId      *string
	orderType    *string
}

// 企业ID
func (r ApiAlipayEbppInvoiceEcorderOrderQueryRequest) EnterpriseId(enterpriseId string) ApiAlipayEbppInvoiceEcorderOrderQueryRequest {
	r.enterpriseId = &enterpriseId
	return r
}

// 共同账户ID
func (r ApiAlipayEbppInvoiceEcorderOrderQueryRequest) AccountId(accountId string) ApiAlipayEbppInvoiceEcorderOrderQueryRequest {
	r.accountId = &accountId
	return r
}

// 授权签约协议号
func (r ApiAlipayEbppInvoiceEcorderOrderQueryRequest) AgreementNo(agreementNo string) ApiAlipayEbppInvoiceEcorderOrderQueryRequest {
	r.agreementNo = &agreementNo
	return r
}

// 订单ID
func (r ApiAlipayEbppInvoiceEcorderOrderQueryRequest) OrderId(orderId string) ApiAlipayEbppInvoiceEcorderOrderQueryRequest {
	r.orderId = &orderId
	return r
}

// 订单类型
func (r ApiAlipayEbppInvoiceEcorderOrderQueryRequest) OrderType(orderType string) ApiAlipayEbppInvoiceEcorderOrderQueryRequest {
	r.orderType = &orderType
	return r
}

func (r ApiAlipayEbppInvoiceEcorderOrderQueryRequest) Execute() (*AlipayEbppInvoiceEcorderOrderQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEbppInvoiceEcorderOrderQueryExecute(r)
}

/*
AlipayEbppInvoiceEcorderOrderQuery 企业码订单查询

企业码订单查询

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEbppInvoiceEcorderOrderQueryRequest
*/
func (r *AlipayEbppInvoiceEcorderOrderAPIService) AlipayEbppInvoiceEcorderOrderQuery(ctx context.Context) ApiAlipayEbppInvoiceEcorderOrderQueryRequest {
	return ApiAlipayEbppInvoiceEcorderOrderQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEbppInvoiceEcorderOrderQueryResponseModel
func (a *AlipayEbppInvoiceEcorderOrderAPIService) AlipayEbppInvoiceEcorderOrderQueryExecute(r ApiAlipayEbppInvoiceEcorderOrderQueryRequest) (*AlipayEbppInvoiceEcorderOrderQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEbppInvoiceEcorderOrderQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEbppInvoiceEcorderOrderAPIService.AlipayEbppInvoiceEcorderOrderQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/ebpp/invoice/ecorder/order/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.enterpriseId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enterprise_id", r.enterpriseId, "form", "")
	}
	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId, "form", "")
	}
	if r.agreementNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "agreement_no", r.agreementNo, "form", "")
	}
	if r.orderId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_id", r.orderId, "form", "")
	}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType, "form", "")
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
		var v AlipayEbppInvoiceEcorderOrderQueryDefaultResponse
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

func (a *AlipayEbppInvoiceEcorderOrderAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayEbppInvoiceEcorderOrderAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
