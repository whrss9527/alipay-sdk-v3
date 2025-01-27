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

// AlipayMarketingCampaignCashListAPIService AlipayMarketingCampaignCashListAPI service
type AlipayMarketingCampaignCashListAPIService service

type ApiAlipayMarketingCampaignCashListQueryRequest struct {
	ctx        context.Context
	ApiService *AlipayMarketingCampaignCashListAPIService
	campStatus *string
	pageSize   *string
	pageIndex  *string
}

// 要查询的活动状态,不填默认为 ALL 返回所有类型。枚举支持: *ALL：所有类型的活动。 *CREATED：已创建未打款。 *PAID：已打款。 *READY：活动已开始。 *PAUSE：活动已暂停。 *CLOSED：活动已结束。 *SETTLE：活动已清算。
func (r ApiAlipayMarketingCampaignCashListQueryRequest) CampStatus(campStatus string) ApiAlipayMarketingCampaignCashListQueryRequest {
	r.campStatus = &campStatus
	return r
}

// 分页查询时每页返回的列表大小，每页数据最大为 50。
func (r ApiAlipayMarketingCampaignCashListQueryRequest) PageSize(pageSize string) ApiAlipayMarketingCampaignCashListQueryRequest {
	r.pageSize = &pageSize
	return r
}

// 分页查询时的页码，从1开始
func (r ApiAlipayMarketingCampaignCashListQueryRequest) PageIndex(pageIndex string) ApiAlipayMarketingCampaignCashListQueryRequest {
	r.pageIndex = &pageIndex
	return r
}

func (r ApiAlipayMarketingCampaignCashListQueryRequest) Execute() (*AlipayMarketingCampaignCashListQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayMarketingCampaignCashListQueryExecute(r)
}

/*
AlipayMarketingCampaignCashListQuery 现金活动列表查询

商户通过开放平台查询自己创建的现金活动列表

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayMarketingCampaignCashListQueryRequest
*/
func (r *AlipayMarketingCampaignCashListAPIService) AlipayMarketingCampaignCashListQuery(ctx context.Context) ApiAlipayMarketingCampaignCashListQueryRequest {
	return ApiAlipayMarketingCampaignCashListQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayMarketingCampaignCashListQueryResponseModel
func (a *AlipayMarketingCampaignCashListAPIService) AlipayMarketingCampaignCashListQueryExecute(r ApiAlipayMarketingCampaignCashListQueryRequest) (*AlipayMarketingCampaignCashListQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayMarketingCampaignCashListQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayMarketingCampaignCashListAPIService.AlipayMarketingCampaignCashListQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/marketing/campaign/cash/list/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.campStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "camp_status", r.campStatus, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
	}
	if r.pageIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_index", r.pageIndex, "form", "")
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
		var v AlipayMarketingCampaignCashListQueryDefaultResponse
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

func (a *AlipayMarketingCampaignCashListAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayMarketingCampaignCashListAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
