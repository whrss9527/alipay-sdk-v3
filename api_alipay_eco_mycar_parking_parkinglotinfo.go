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

// AlipayEcoMycarParkingParkinglotinfoAPIService AlipayEcoMycarParkingParkinglotinfoAPI service
type AlipayEcoMycarParkingParkinglotinfoAPIService service

type ApiAlipayEcoMycarParkingParkinglotinfoCreateRequest struct {
	ctx                                            context.Context
	ApiService                                     *AlipayEcoMycarParkingParkinglotinfoAPIService
	alipayEcoMycarParkingParkinglotinfoCreateModel *AlipayEcoMycarParkingParkinglotinfoCreateModel
}

func (r ApiAlipayEcoMycarParkingParkinglotinfoCreateRequest) AlipayEcoMycarParkingParkinglotinfoCreateModel(alipayEcoMycarParkingParkinglotinfoCreateModel AlipayEcoMycarParkingParkinglotinfoCreateModel) ApiAlipayEcoMycarParkingParkinglotinfoCreateRequest {
	r.alipayEcoMycarParkingParkinglotinfoCreateModel = &alipayEcoMycarParkingParkinglotinfoCreateModel
	return r
}

func (r ApiAlipayEcoMycarParkingParkinglotinfoCreateRequest) Execute() (*AlipayEcoMycarParkingParkinglotinfoCreateResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEcoMycarParkingParkinglotinfoCreateExecute(r)
}

/*
AlipayEcoMycarParkingParkinglotinfoCreate 录入停车场信息

录入停车场信息，内容信息通过该接口提交到支付宝，支付宝会生成支付宝这边停车场ID反馈给开发者，用户后续更新和上送车辆信息。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEcoMycarParkingParkinglotinfoCreateRequest
*/
func (r *AlipayEcoMycarParkingParkinglotinfoAPIService) AlipayEcoMycarParkingParkinglotinfoCreate(ctx context.Context) ApiAlipayEcoMycarParkingParkinglotinfoCreateRequest {
	return ApiAlipayEcoMycarParkingParkinglotinfoCreateRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEcoMycarParkingParkinglotinfoCreateResponseModel
func (a *AlipayEcoMycarParkingParkinglotinfoAPIService) AlipayEcoMycarParkingParkinglotinfoCreateExecute(r ApiAlipayEcoMycarParkingParkinglotinfoCreateRequest) (*AlipayEcoMycarParkingParkinglotinfoCreateResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEcoMycarParkingParkinglotinfoCreateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEcoMycarParkingParkinglotinfoAPIService.AlipayEcoMycarParkingParkinglotinfoCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/eco/mycar/parking/parkinglotinfo/create"

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
	localVarPostBody = r.alipayEcoMycarParkingParkinglotinfoCreateModel

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
		var v AlipayEcoMycarParkingParkinglotinfoCreateDefaultResponse
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

type ApiAlipayEcoMycarParkingParkinglotinfoQueryRequest struct {
	ctx          context.Context
	ApiService   *AlipayEcoMycarParkingParkinglotinfoAPIService
	parkingId    *string
	outParkingId *string
}

// 支付宝停车平台ID，由支付宝定义的该停车场标识，同一个ISV或商户范围内唯一。通过 &lt;a href&#x3D;\&quot;https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.parkinglotinfo.create\&quot;&gt;alipay.eco.mycar.parking.parkinglotinfo.create&lt;/a&gt;(录入停车场信息)接口获取。 注意：parking_id和out_parking_id不能同时为空。
func (r ApiAlipayEcoMycarParkingParkinglotinfoQueryRequest) ParkingId(parkingId string) ApiAlipayEcoMycarParkingParkinglotinfoQueryRequest {
	r.parkingId = &parkingId
	return r
}

// ISV停车场ID，由ISV定义的停车场标识，同一个ISV或商户范围内唯一。需与 &lt;a href&#x3D;\&quot;https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.parkinglotinfo.create\&quot;&gt;alipay.eco.mycar.parking.parkinglotinfo.create&lt;/a&gt;(录入停车场信息)接口传入值一致。 注意：parking_id和out_parking_id不能同时为空。
func (r ApiAlipayEcoMycarParkingParkinglotinfoQueryRequest) OutParkingId(outParkingId string) ApiAlipayEcoMycarParkingParkinglotinfoQueryRequest {
	r.outParkingId = &outParkingId
	return r
}

func (r ApiAlipayEcoMycarParkingParkinglotinfoQueryRequest) Execute() (*AlipayEcoMycarParkingParkinglotinfoQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEcoMycarParkingParkinglotinfoQueryExecute(r)
}

/*
AlipayEcoMycarParkingParkinglotinfoQuery 停车场信息查询

停车场信息查询，通过停车场id或者ISV停车场ID查询停车场信息。只能查询正在调用接口商户自己创建的停车场，限制在归属PID层面上。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEcoMycarParkingParkinglotinfoQueryRequest
*/
func (r *AlipayEcoMycarParkingParkinglotinfoAPIService) AlipayEcoMycarParkingParkinglotinfoQuery(ctx context.Context) ApiAlipayEcoMycarParkingParkinglotinfoQueryRequest {
	return ApiAlipayEcoMycarParkingParkinglotinfoQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEcoMycarParkingParkinglotinfoQueryResponseModel
func (a *AlipayEcoMycarParkingParkinglotinfoAPIService) AlipayEcoMycarParkingParkinglotinfoQueryExecute(r ApiAlipayEcoMycarParkingParkinglotinfoQueryRequest) (*AlipayEcoMycarParkingParkinglotinfoQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEcoMycarParkingParkinglotinfoQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEcoMycarParkingParkinglotinfoAPIService.AlipayEcoMycarParkingParkinglotinfoQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/eco/mycar/parking/parkinglotinfo/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.parkingId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "parking_id", r.parkingId, "form", "")
	}
	if r.outParkingId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "out_parking_id", r.outParkingId, "form", "")
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
		var v AlipayEcoMycarParkingParkinglotinfoQueryDefaultResponse
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

type ApiAlipayEcoMycarParkingParkinglotinfoUpdateRequest struct {
	ctx                                            context.Context
	ApiService                                     *AlipayEcoMycarParkingParkinglotinfoAPIService
	alipayEcoMycarParkingParkinglotinfoUpdateModel *AlipayEcoMycarParkingParkinglotinfoUpdateModel
}

func (r ApiAlipayEcoMycarParkingParkinglotinfoUpdateRequest) AlipayEcoMycarParkingParkinglotinfoUpdateModel(alipayEcoMycarParkingParkinglotinfoUpdateModel AlipayEcoMycarParkingParkinglotinfoUpdateModel) ApiAlipayEcoMycarParkingParkinglotinfoUpdateRequest {
	r.alipayEcoMycarParkingParkinglotinfoUpdateModel = &alipayEcoMycarParkingParkinglotinfoUpdateModel
	return r
}

func (r ApiAlipayEcoMycarParkingParkinglotinfoUpdateRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayEcoMycarParkingParkinglotinfoUpdateExecute(r)
}

/*
AlipayEcoMycarParkingParkinglotinfoUpdate 修改停车场信息

录入停车场信息，内容信息通过该接口提交到支付宝，支付宝会生成支付宝这边停车场ID反馈给开发者，用于后续更新和上送车辆信息，停车场名称不允许修改。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEcoMycarParkingParkinglotinfoUpdateRequest
*/
func (r *AlipayEcoMycarParkingParkinglotinfoAPIService) AlipayEcoMycarParkingParkinglotinfoUpdate(ctx context.Context) ApiAlipayEcoMycarParkingParkinglotinfoUpdateRequest {
	return ApiAlipayEcoMycarParkingParkinglotinfoUpdateRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *AlipayEcoMycarParkingParkinglotinfoAPIService) AlipayEcoMycarParkingParkinglotinfoUpdateExecute(r ApiAlipayEcoMycarParkingParkinglotinfoUpdateRequest) (map[string]interface{}, *http.Response, error) {
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

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEcoMycarParkingParkinglotinfoAPIService.AlipayEcoMycarParkingParkinglotinfoUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/eco/mycar/parking/parkinglotinfo/update"

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
	localVarPostBody = r.alipayEcoMycarParkingParkinglotinfoUpdateModel

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
		var v AlipayEcoMycarParkingParkinglotinfoUpdateDefaultResponse
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

func (a *AlipayEcoMycarParkingParkinglotinfoAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayEcoMycarParkingParkinglotinfoAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
