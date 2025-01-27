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
	"strings"
	"time"
)

// AlipayMarketingActivityGoodsAPIService AlipayMarketingActivityGoodsAPI service
type AlipayMarketingActivityGoodsAPIService service

type ApiAlipayMarketingActivityGoodsBatchqueryRequest struct {
	ctx                context.Context
	ApiService         *AlipayMarketingActivityGoodsAPIService
	activityId         string
	merchantId         *string
	goodsUseType       *string
	pageNum            *int32
	pageSize           *int32
	productVersion     *string
	merchantAccessMode *string
}

// 商户PID,默认为当前接口调用商户。
func (r ApiAlipayMarketingActivityGoodsBatchqueryRequest) MerchantId(merchantId string) ApiAlipayMarketingActivityGoodsBatchqueryRequest {
	r.merchantId = &merchantId
	return r
}

// 活动单品类型。
func (r ApiAlipayMarketingActivityGoodsBatchqueryRequest) GoodsUseType(goodsUseType string) ApiAlipayMarketingActivityGoodsBatchqueryRequest {
	r.goodsUseType = &goodsUseType
	return r
}

// 分页查询页码。 限制: 必须为大于0的整数
func (r ApiAlipayMarketingActivityGoodsBatchqueryRequest) PageNum(pageNum int32) ApiAlipayMarketingActivityGoodsBatchqueryRequest {
	r.pageNum = &pageNum
	return r
}

// 分页查询单页数据条数。 限制: 1.必须为大于0的整数 2.每页最大值为20
func (r ApiAlipayMarketingActivityGoodsBatchqueryRequest) PageSize(pageSize int32) ApiAlipayMarketingActivityGoodsBatchqueryRequest {
	r.pageSize = &pageSize
	return r
}

// 版本号  枚举值: 2.0.0
func (r ApiAlipayMarketingActivityGoodsBatchqueryRequest) ProductVersion(productVersion string) ApiAlipayMarketingActivityGoodsBatchqueryRequest {
	r.productVersion = &productVersion
	return r
}

// 商户接入模式
func (r ApiAlipayMarketingActivityGoodsBatchqueryRequest) MerchantAccessMode(merchantAccessMode string) ApiAlipayMarketingActivityGoodsBatchqueryRequest {
	r.merchantAccessMode = &merchantAccessMode
	return r
}

func (r ApiAlipayMarketingActivityGoodsBatchqueryRequest) Execute() (*AlipayMarketingActivityGoodsBatchqueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayMarketingActivityGoodsBatchqueryExecute(r)
}

/*
AlipayMarketingActivityGoodsBatchquery 查询活动适用商品

通过此接口可查询活动的可用或不可用商品,判断是否在该商品可用，来决定是否展示。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param activityId 活动id
	@return ApiAlipayMarketingActivityGoodsBatchqueryRequest
*/
func (r *AlipayMarketingActivityGoodsAPIService) AlipayMarketingActivityGoodsBatchquery(ctx context.Context, activityId string) ApiAlipayMarketingActivityGoodsBatchqueryRequest {
	return ApiAlipayMarketingActivityGoodsBatchqueryRequest{
		ApiService: r,
		ctx:        ctx,
		activityId: activityId,
	}
}

// Execute executes the request
//
//	@return AlipayMarketingActivityGoodsBatchqueryResponseModel
func (a *AlipayMarketingActivityGoodsAPIService) AlipayMarketingActivityGoodsBatchqueryExecute(r ApiAlipayMarketingActivityGoodsBatchqueryRequest) (*AlipayMarketingActivityGoodsBatchqueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayMarketingActivityGoodsBatchqueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayMarketingActivityGoodsAPIService.AlipayMarketingActivityGoodsBatchquery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/marketing/activity/{activity_id}/goods/batchquery"
	localVarPath = strings.Replace(localVarPath, "{"+"activity_id"+"}", url.PathEscape(parameterValueToString(r.activityId, "activityId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.merchantId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "merchant_id", r.merchantId, "form", "")
	}
	if r.goodsUseType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "goods_use_type", r.goodsUseType, "form", "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_num", r.pageNum, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
	}
	if r.productVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_version", r.productVersion, "form", "")
	}
	if r.merchantAccessMode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "merchant_access_mode", r.merchantAccessMode, "form", "")
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
		var v AlipayMarketingActivityGoodsBatchqueryDefaultResponse
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

func (a *AlipayMarketingActivityGoodsAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayMarketingActivityGoodsAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
