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

// AlipayEcoMycarParkingConfigAPIService AlipayEcoMycarParkingConfigAPI service
type AlipayEcoMycarParkingConfigAPIService service

type ApiAlipayEcoMycarParkingConfigQueryRequest struct {
	ctx           context.Context
	ApiService    *AlipayEcoMycarParkingConfigAPIService
	interfaceName *string
	interfaceType *string
}

// 接口名称。H5传入参数固定值：alipay.eco.mycar.parking.userpage.query； 小程序传入参数固定值：alipay.eco.mycar.parking.isv.homepage； 协议状态变更通知传入参数固定值： alipay.eco.mycar.parking.agreement.notify。
func (r ApiAlipayEcoMycarParkingConfigQueryRequest) InterfaceName(interfaceName string) ApiAlipayEcoMycarParkingConfigQueryRequest {
	r.interfaceName = &interfaceName
	return r
}

// 接口类型。H5、小程序传入参数固定值：interface_page； 通知接口传入参数固定值：interface_service。
func (r ApiAlipayEcoMycarParkingConfigQueryRequest) InterfaceType(interfaceType string) ApiAlipayEcoMycarParkingConfigQueryRequest {
	r.interfaceType = &interfaceType
	return r
}

func (r ApiAlipayEcoMycarParkingConfigQueryRequest) Execute() (*AlipayEcoMycarParkingConfigQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEcoMycarParkingConfigQueryExecute(r)
}

/*
AlipayEcoMycarParkingConfigQuery ISV系统配置查询接口

ISV通过该接口，查询ISV已注册到车主平台停车业务系统配置信息。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEcoMycarParkingConfigQueryRequest
*/
func (r *AlipayEcoMycarParkingConfigAPIService) AlipayEcoMycarParkingConfigQuery(ctx context.Context) ApiAlipayEcoMycarParkingConfigQueryRequest {
	return ApiAlipayEcoMycarParkingConfigQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEcoMycarParkingConfigQueryResponseModel
func (a *AlipayEcoMycarParkingConfigAPIService) AlipayEcoMycarParkingConfigQueryExecute(r ApiAlipayEcoMycarParkingConfigQueryRequest) (*AlipayEcoMycarParkingConfigQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEcoMycarParkingConfigQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEcoMycarParkingConfigAPIService.AlipayEcoMycarParkingConfigQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/eco/mycar/parking/config/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.interfaceName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interface_name", r.interfaceName, "form", "")
	}
	if r.interfaceType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interface_type", r.interfaceType, "form", "")
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
		var v AlipayEcoMycarParkingConfigQueryDefaultResponse
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

type ApiAlipayEcoMycarParkingConfigSetRequest struct {
	ctx                                 context.Context
	ApiService                          *AlipayEcoMycarParkingConfigAPIService
	alipayEcoMycarParkingConfigSetModel *AlipayEcoMycarParkingConfigSetModel
}

func (r ApiAlipayEcoMycarParkingConfigSetRequest) AlipayEcoMycarParkingConfigSetModel(alipayEcoMycarParkingConfigSetModel AlipayEcoMycarParkingConfigSetModel) ApiAlipayEcoMycarParkingConfigSetRequest {
	r.alipayEcoMycarParkingConfigSetModel = &alipayEcoMycarParkingConfigSetModel
	return r
}

func (r ApiAlipayEcoMycarParkingConfigSetRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayEcoMycarParkingConfigSetExecute(r)
}

/*
AlipayEcoMycarParkingConfigSet 停车ISV系统配置接口

ISV通过该接口，配置车主平台各业务所需要的ISV的系统配置信息。如果接口配置存在则覆盖原有的接口信息。通过该接口设置的配置信息是立刻生效的，在调用该接口修改配置信息时请先评估对线上业务的影响，并做好处理方案。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEcoMycarParkingConfigSetRequest
*/
func (r *AlipayEcoMycarParkingConfigAPIService) AlipayEcoMycarParkingConfigSet(ctx context.Context) ApiAlipayEcoMycarParkingConfigSetRequest {
	return ApiAlipayEcoMycarParkingConfigSetRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *AlipayEcoMycarParkingConfigAPIService) AlipayEcoMycarParkingConfigSetExecute(r ApiAlipayEcoMycarParkingConfigSetRequest) (map[string]interface{}, *http.Response, error) {
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

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEcoMycarParkingConfigAPIService.AlipayEcoMycarParkingConfigSet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/eco/mycar/parking/config/set"

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
	localVarPostBody = r.alipayEcoMycarParkingConfigSetModel

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
		var v AlipayEcoMycarParkingConfigSetDefaultResponse
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

func (a *AlipayEcoMycarParkingConfigAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayEcoMycarParkingConfigAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
