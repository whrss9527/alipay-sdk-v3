/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

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


// AlipayMarketingActivityVoucherpackageAPIService AlipayMarketingActivityVoucherpackageAPI service
type AlipayMarketingActivityVoucherpackageAPIService service

type ApiAlipayMarketingActivityVoucherpackageBatchqueryRequest struct {
	ctx context.Context
	ApiService *AlipayMarketingActivityVoucherpackageAPIService
	alipayMarketingActivityVoucherpackageBatchqueryModel *AlipayMarketingActivityVoucherpackageBatchqueryModel
}

func (r ApiAlipayMarketingActivityVoucherpackageBatchqueryRequest) AlipayMarketingActivityVoucherpackageBatchqueryModel(alipayMarketingActivityVoucherpackageBatchqueryModel AlipayMarketingActivityVoucherpackageBatchqueryModel) ApiAlipayMarketingActivityVoucherpackageBatchqueryRequest {
	r.alipayMarketingActivityVoucherpackageBatchqueryModel = &alipayMarketingActivityVoucherpackageBatchqueryModel
	return r
}

func (r ApiAlipayMarketingActivityVoucherpackageBatchqueryRequest) Execute() (*AlipayMarketingActivityVoucherpackageBatchqueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayMarketingActivityVoucherpackageBatchqueryExecute(r)
}

/*
AlipayMarketingActivityVoucherpackageBatchquery 券包批量查询

券包批量查询

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayMarketingActivityVoucherpackageBatchqueryRequest
*/
func (r *AlipayMarketingActivityVoucherpackageAPIService) AlipayMarketingActivityVoucherpackageBatchquery(ctx context.Context) ApiAlipayMarketingActivityVoucherpackageBatchqueryRequest {
	return ApiAlipayMarketingActivityVoucherpackageBatchqueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayMarketingActivityVoucherpackageBatchqueryResponseModel
func (a *AlipayMarketingActivityVoucherpackageAPIService) AlipayMarketingActivityVoucherpackageBatchqueryExecute(r ApiAlipayMarketingActivityVoucherpackageBatchqueryRequest) (*AlipayMarketingActivityVoucherpackageBatchqueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayMarketingActivityVoucherpackageAPIService.AlipayMarketingActivityVoucherpackageBatchquery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/marketing/activity/voucherpackage/batchquery"

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
	localVarPostBody = r.alipayMarketingActivityVoucherpackageBatchqueryModel

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
			var v AlipayMarketingActivityVoucherpackageBatchqueryDefaultResponse
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


type ApiAlipayMarketingActivityVoucherpackageConsultRequest struct {
	ctx context.Context
	ApiService *AlipayMarketingActivityVoucherpackageAPIService
	authToken *string
	alipayMarketingActivityVoucherpackageConsultModel *AlipayMarketingActivityVoucherpackageConsultModel
}

// 用户授权令牌
func (r ApiAlipayMarketingActivityVoucherpackageConsultRequest) AuthToken(authToken string) ApiAlipayMarketingActivityVoucherpackageConsultRequest {
	r.authToken = &authToken
	return r
}

func (r ApiAlipayMarketingActivityVoucherpackageConsultRequest) AlipayMarketingActivityVoucherpackageConsultModel(alipayMarketingActivityVoucherpackageConsultModel AlipayMarketingActivityVoucherpackageConsultModel) ApiAlipayMarketingActivityVoucherpackageConsultRequest {
	r.alipayMarketingActivityVoucherpackageConsultModel = &alipayMarketingActivityVoucherpackageConsultModel
	return r
}

func (r ApiAlipayMarketingActivityVoucherpackageConsultRequest) Execute() (*AlipayMarketingActivityVoucherpackageConsultResponseModel, *http.Response, error) {
	return r.ApiService.AlipayMarketingActivityVoucherpackageConsultExecute(r)
}

/*
AlipayMarketingActivityVoucherpackageConsult 券包购买咨询

券包购买咨询，咨询当前用户是否可以购买指定的券包

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayMarketingActivityVoucherpackageConsultRequest
*/
func (r *AlipayMarketingActivityVoucherpackageAPIService) AlipayMarketingActivityVoucherpackageConsult(ctx context.Context) ApiAlipayMarketingActivityVoucherpackageConsultRequest {
	return ApiAlipayMarketingActivityVoucherpackageConsultRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayMarketingActivityVoucherpackageConsultResponseModel
func (a *AlipayMarketingActivityVoucherpackageAPIService) AlipayMarketingActivityVoucherpackageConsultExecute(r ApiAlipayMarketingActivityVoucherpackageConsultRequest) (*AlipayMarketingActivityVoucherpackageConsultResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayMarketingActivityVoucherpackageConsultResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayMarketingActivityVoucherpackageAPIService.AlipayMarketingActivityVoucherpackageConsult")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/marketing/activity/voucherpackage/consult"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.authToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "auth_token", r.authToken, "form", "")
	}

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
	localVarPostBody = r.alipayMarketingActivityVoucherpackageConsultModel

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
			var v AlipayMarketingActivityVoucherpackageConsultDefaultResponse
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


type ApiAlipayMarketingActivityVoucherpackageQueryRequest struct {
	ctx context.Context
	ApiService *AlipayMarketingActivityVoucherpackageAPIService
	voucherPackageId *string
}

// 券包id
func (r ApiAlipayMarketingActivityVoucherpackageQueryRequest) VoucherPackageId(voucherPackageId string) ApiAlipayMarketingActivityVoucherpackageQueryRequest {
	r.voucherPackageId = &voucherPackageId
	return r
}

func (r ApiAlipayMarketingActivityVoucherpackageQueryRequest) Execute() (*AlipayMarketingActivityVoucherpackageQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayMarketingActivityVoucherpackageQueryExecute(r)
}

/*
AlipayMarketingActivityVoucherpackageQuery 券包详情查询

查询券包详情，包括券包信息和券包下面券信息

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayMarketingActivityVoucherpackageQueryRequest
*/
func (r *AlipayMarketingActivityVoucherpackageAPIService) AlipayMarketingActivityVoucherpackageQuery(ctx context.Context) ApiAlipayMarketingActivityVoucherpackageQueryRequest {
	return ApiAlipayMarketingActivityVoucherpackageQueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayMarketingActivityVoucherpackageQueryResponseModel
func (a *AlipayMarketingActivityVoucherpackageAPIService) AlipayMarketingActivityVoucherpackageQueryExecute(r ApiAlipayMarketingActivityVoucherpackageQueryRequest) (*AlipayMarketingActivityVoucherpackageQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayMarketingActivityVoucherpackageQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayMarketingActivityVoucherpackageAPIService.AlipayMarketingActivityVoucherpackageQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/marketing/activity/voucherpackage/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.voucherPackageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "voucher_package_id", r.voucherPackageId, "form", "")
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
			var v AlipayMarketingActivityVoucherpackageQueryDefaultResponse
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


func (a *AlipayMarketingActivityVoucherpackageAPIService) signRequest(req *http.Request) error {
	appID := a.client.cfg.AppID
	appCertSN := a.client.cfg.AppCertSN
	privateKey := a.client.cfg.PrivateKey

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

	signature, err := signWithRSA(content, privateKey)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("ALIPAY-SHA256withRSA %s,sign=%s", authString, signature))
	return nil
}

func (a *AlipayMarketingActivityVoucherpackageAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}

