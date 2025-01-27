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

// AlipayIserviceCcmSwTreecategoryAPIService AlipayIserviceCcmSwTreecategoryAPI service
type AlipayIserviceCcmSwTreecategoryAPIService service

type ApiAlipayIserviceCcmSwTreecategoryBatchqueryRequest struct {
	ctx           context.Context
	ApiService    *AlipayIserviceCcmSwTreecategoryAPIService
	ccsInstanceId *string
	libraryId     *int32
	name          *string
	pageNum       *int32
	pageSize      *int32
}

// 子部门ID，不传为默认部门
func (r ApiAlipayIserviceCcmSwTreecategoryBatchqueryRequest) CcsInstanceId(ccsInstanceId string) ApiAlipayIserviceCcmSwTreecategoryBatchqueryRequest {
	r.ccsInstanceId = &ccsInstanceId
	return r
}

// 知识库ID
func (r ApiAlipayIserviceCcmSwTreecategoryBatchqueryRequest) LibraryId(libraryId int32) ApiAlipayIserviceCcmSwTreecategoryBatchqueryRequest {
	r.libraryId = &libraryId
	return r
}

// 节点名称
func (r ApiAlipayIserviceCcmSwTreecategoryBatchqueryRequest) Name(name string) ApiAlipayIserviceCcmSwTreecategoryBatchqueryRequest {
	r.name = &name
	return r
}

// 页数，page_size不能为空
func (r ApiAlipayIserviceCcmSwTreecategoryBatchqueryRequest) PageNum(pageNum int32) ApiAlipayIserviceCcmSwTreecategoryBatchqueryRequest {
	r.pageNum = &pageNum
	return r
}

// 页显示大小，page_num不能为空
func (r ApiAlipayIserviceCcmSwTreecategoryBatchqueryRequest) PageSize(pageSize int32) ApiAlipayIserviceCcmSwTreecategoryBatchqueryRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiAlipayIserviceCcmSwTreecategoryBatchqueryRequest) Execute() (*AlipayIserviceCcmSwTreecategoryBatchqueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayIserviceCcmSwTreecategoryBatchqueryExecute(r)
}

/*
AlipayIserviceCcmSwTreecategoryBatchquery 知识库-节点-批量查询

知识库-节点-批量查询

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayIserviceCcmSwTreecategoryBatchqueryRequest
*/
func (r *AlipayIserviceCcmSwTreecategoryAPIService) AlipayIserviceCcmSwTreecategoryBatchquery(ctx context.Context) ApiAlipayIserviceCcmSwTreecategoryBatchqueryRequest {
	return ApiAlipayIserviceCcmSwTreecategoryBatchqueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayIserviceCcmSwTreecategoryBatchqueryResponseModel
func (a *AlipayIserviceCcmSwTreecategoryAPIService) AlipayIserviceCcmSwTreecategoryBatchqueryExecute(r ApiAlipayIserviceCcmSwTreecategoryBatchqueryRequest) (*AlipayIserviceCcmSwTreecategoryBatchqueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayIserviceCcmSwTreecategoryBatchqueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayIserviceCcmSwTreecategoryAPIService.AlipayIserviceCcmSwTreecategoryBatchquery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/iservice/ccm/sw/treecategory/batchquery"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.ccsInstanceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ccs_instance_id", r.ccsInstanceId, "form", "")
	}
	if r.libraryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "library_id", r.libraryId, "form", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_num", r.pageNum, "form", "")
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
		var v AlipayIserviceCcmSwTreecategoryBatchqueryDefaultResponse
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

type ApiAlipayIserviceCcmSwTreecategoryCreateRequest struct {
	ctx                                        context.Context
	ApiService                                 *AlipayIserviceCcmSwTreecategoryAPIService
	alipayIserviceCcmSwTreecategoryCreateModel *AlipayIserviceCcmSwTreecategoryCreateModel
}

func (r ApiAlipayIserviceCcmSwTreecategoryCreateRequest) AlipayIserviceCcmSwTreecategoryCreateModel(alipayIserviceCcmSwTreecategoryCreateModel AlipayIserviceCcmSwTreecategoryCreateModel) ApiAlipayIserviceCcmSwTreecategoryCreateRequest {
	r.alipayIserviceCcmSwTreecategoryCreateModel = &alipayIserviceCcmSwTreecategoryCreateModel
	return r
}

func (r ApiAlipayIserviceCcmSwTreecategoryCreateRequest) Execute() (*AlipayIserviceCcmSwTreecategoryCreateResponseModel, *http.Response, error) {
	return r.ApiService.AlipayIserviceCcmSwTreecategoryCreateExecute(r)
}

/*
AlipayIserviceCcmSwTreecategoryCreate 知识库-节点-创建

知识库-节点-创建

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayIserviceCcmSwTreecategoryCreateRequest
*/
func (r *AlipayIserviceCcmSwTreecategoryAPIService) AlipayIserviceCcmSwTreecategoryCreate(ctx context.Context) ApiAlipayIserviceCcmSwTreecategoryCreateRequest {
	return ApiAlipayIserviceCcmSwTreecategoryCreateRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayIserviceCcmSwTreecategoryCreateResponseModel
func (a *AlipayIserviceCcmSwTreecategoryAPIService) AlipayIserviceCcmSwTreecategoryCreateExecute(r ApiAlipayIserviceCcmSwTreecategoryCreateRequest) (*AlipayIserviceCcmSwTreecategoryCreateResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayIserviceCcmSwTreecategoryCreateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayIserviceCcmSwTreecategoryAPIService.AlipayIserviceCcmSwTreecategoryCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/iservice/ccm/sw/treecategory/create"

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
	localVarPostBody = r.alipayIserviceCcmSwTreecategoryCreateModel

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
		var v AlipayIserviceCcmSwTreecategoryCreateDefaultResponse
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

type ApiAlipayIserviceCcmSwTreecategoryDeleteRequest struct {
	ctx        context.Context
	ApiService *AlipayIserviceCcmSwTreecategoryAPIService
}

func (r ApiAlipayIserviceCcmSwTreecategoryDeleteRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayIserviceCcmSwTreecategoryDeleteExecute(r)
}

/*
AlipayIserviceCcmSwTreecategoryDelete 知识库-节点-删除

知识库-节点-删除

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayIserviceCcmSwTreecategoryDeleteRequest
*/
func (r *AlipayIserviceCcmSwTreecategoryAPIService) AlipayIserviceCcmSwTreecategoryDelete(ctx context.Context) ApiAlipayIserviceCcmSwTreecategoryDeleteRequest {
	return ApiAlipayIserviceCcmSwTreecategoryDeleteRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *AlipayIserviceCcmSwTreecategoryAPIService) AlipayIserviceCcmSwTreecategoryDeleteExecute(r ApiAlipayIserviceCcmSwTreecategoryDeleteRequest) (map[string]interface{}, *http.Response, error) {
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

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayIserviceCcmSwTreecategoryAPIService.AlipayIserviceCcmSwTreecategoryDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/iservice/ccm/sw/treecategory/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		var v AlipayIserviceCcmSwTreecategoryDeleteDefaultResponse
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

type ApiAlipayIserviceCcmSwTreecategoryModifyRequest struct {
	ctx                                        context.Context
	ApiService                                 *AlipayIserviceCcmSwTreecategoryAPIService
	alipayIserviceCcmSwTreecategoryModifyModel *AlipayIserviceCcmSwTreecategoryModifyModel
}

func (r ApiAlipayIserviceCcmSwTreecategoryModifyRequest) AlipayIserviceCcmSwTreecategoryModifyModel(alipayIserviceCcmSwTreecategoryModifyModel AlipayIserviceCcmSwTreecategoryModifyModel) ApiAlipayIserviceCcmSwTreecategoryModifyRequest {
	r.alipayIserviceCcmSwTreecategoryModifyModel = &alipayIserviceCcmSwTreecategoryModifyModel
	return r
}

func (r ApiAlipayIserviceCcmSwTreecategoryModifyRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayIserviceCcmSwTreecategoryModifyExecute(r)
}

/*
AlipayIserviceCcmSwTreecategoryModify 知识库-节点-修改

知识库-节点-修改

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayIserviceCcmSwTreecategoryModifyRequest
*/
func (r *AlipayIserviceCcmSwTreecategoryAPIService) AlipayIserviceCcmSwTreecategoryModify(ctx context.Context) ApiAlipayIserviceCcmSwTreecategoryModifyRequest {
	return ApiAlipayIserviceCcmSwTreecategoryModifyRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *AlipayIserviceCcmSwTreecategoryAPIService) AlipayIserviceCcmSwTreecategoryModifyExecute(r ApiAlipayIserviceCcmSwTreecategoryModifyRequest) (map[string]interface{}, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayIserviceCcmSwTreecategoryAPIService.AlipayIserviceCcmSwTreecategoryModify")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/iservice/ccm/sw/treecategory/modify"

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
	localVarPostBody = r.alipayIserviceCcmSwTreecategoryModifyModel

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
		var v AlipayIserviceCcmSwTreecategoryModifyDefaultResponse
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

func (a *AlipayIserviceCcmSwTreecategoryAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayIserviceCcmSwTreecategoryAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
