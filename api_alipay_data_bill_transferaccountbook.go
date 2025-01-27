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

// AlipayDataBillTransferaccountbookAPIService AlipayDataBillTransferaccountbookAPI service
type AlipayDataBillTransferaccountbookAPIService service

type ApiAlipayDataBillTransferaccountbookQueryRequest struct {
	ctx         context.Context
	ApiService  *AlipayDataBillTransferaccountbookAPIService
	startTime   *string
	endTime     *string
	type_       *string
	agreementNo *string
	storeNo     *string
	pageNo      *string
	pageSize    *string
}

// 充值、转账、提现流水业务时间的起始范围
func (r ApiAlipayDataBillTransferaccountbookQueryRequest) StartTime(startTime string) ApiAlipayDataBillTransferaccountbookQueryRequest {
	r.startTime = &startTime
	return r
}

// 充值、转账、提现流水业务时间的结束范围。与起始时间间隔不超过31天。查询结果为起始时间至结束时间的左闭右开区间
func (r ApiAlipayDataBillTransferaccountbookQueryRequest) EndTime(endTime string) ApiAlipayDataBillTransferaccountbookQueryRequest {
	r.endTime = &endTime
	return r
}

// 转账类型：充值-DEPOSIT，提现-WITHDRAW，转账-TRANSFER。
func (r ApiAlipayDataBillTransferaccountbookQueryRequest) Type_(type_ string) ApiAlipayDataBillTransferaccountbookQueryRequest {
	r.type_ = &type_
	return r
}

// 协议号
func (r ApiAlipayDataBillTransferaccountbookQueryRequest) AgreementNo(agreementNo string) ApiAlipayDataBillTransferaccountbookQueryRequest {
	r.agreementNo = &agreementNo
	return r
}

// 子账本号，或者子账本名称。模糊查询
func (r ApiAlipayDataBillTransferaccountbookQueryRequest) StoreNo(storeNo string) ApiAlipayDataBillTransferaccountbookQueryRequest {
	r.storeNo = &storeNo
	return r
}

// 分页号，从1开始
func (r ApiAlipayDataBillTransferaccountbookQueryRequest) PageNo(pageNo string) ApiAlipayDataBillTransferaccountbookQueryRequest {
	r.pageNo = &pageNo
	return r
}

// 分页大小1000-2000，默认2000
func (r ApiAlipayDataBillTransferaccountbookQueryRequest) PageSize(pageSize string) ApiAlipayDataBillTransferaccountbookQueryRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiAlipayDataBillTransferaccountbookQueryRequest) Execute() (*AlipayDataBillTransferaccountbookQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayDataBillTransferaccountbookQueryExecute(r)
}

/*
AlipayDataBillTransferaccountbookQuery 子账本充提转账单查询(incubating)

子账本充提转账单查询（子账本业务定制接口）

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayDataBillTransferaccountbookQueryRequest
*/
func (r *AlipayDataBillTransferaccountbookAPIService) AlipayDataBillTransferaccountbookQuery(ctx context.Context) ApiAlipayDataBillTransferaccountbookQueryRequest {
	return ApiAlipayDataBillTransferaccountbookQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayDataBillTransferaccountbookQueryResponseModel
func (a *AlipayDataBillTransferaccountbookAPIService) AlipayDataBillTransferaccountbookQueryExecute(r ApiAlipayDataBillTransferaccountbookQueryRequest) (*AlipayDataBillTransferaccountbookQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayDataBillTransferaccountbookQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayDataBillTransferaccountbookAPIService.AlipayDataBillTransferaccountbookQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/data/bill/transferaccountbook/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime, "form", "")
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime, "form", "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.agreementNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "agreement_no", r.agreementNo, "form", "")
	}
	if r.storeNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "store_no", r.storeNo, "form", "")
	}
	if r.pageNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_no", r.pageNo, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
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
		var v AlipayDataBillTransferaccountbookQueryDefaultResponse
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

func (a *AlipayDataBillTransferaccountbookAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayDataBillTransferaccountbookAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
