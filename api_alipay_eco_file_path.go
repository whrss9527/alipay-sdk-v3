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

// AlipayEcoFilePathAPIService AlipayEcoFilePathAPI service
type AlipayEcoFilePathAPIService service

type ApiAlipayEcoFilePathQueryRequest struct {
	ctx         context.Context
	ApiService  *AlipayEcoFilePathAPIService
	targetAppId *string
	contentMd5  *string
	contentType *string
	fileName    *string
	fileSize    *int32
}

// 目标isv应用ID
func (r ApiAlipayEcoFilePathQueryRequest) TargetAppId(targetAppId string) ApiAlipayEcoFilePathQueryRequest {
	r.targetAppId = &targetAppId
	return r
}

// 先计算文件md5值，在对该md5值进行base64编码
func (r ApiAlipayEcoFilePathQueryRequest) ContentMd5(contentMd5 string) ApiAlipayEcoFilePathQueryRequest {
	r.contentMd5 = &contentMd5
	return r
}

// 目标文件的MIME类型
func (r ApiAlipayEcoFilePathQueryRequest) ContentType(contentType string) ApiAlipayEcoFilePathQueryRequest {
	r.contentType = &contentType
	return r
}

// 文件名称。 注意：必须带上文件扩展名，不然会导致后续发起流程校验不通过。示例：合同.pdf。
func (r ApiAlipayEcoFilePathQueryRequest) FileName(fileName string) ApiAlipayEcoFilePathQueryRequest {
	r.fileName = &fileName
	return r
}

// 文件大小，单位byte。最大允许上传30Mb
func (r ApiAlipayEcoFilePathQueryRequest) FileSize(fileSize int32) ApiAlipayEcoFilePathQueryRequest {
	r.fileSize = &fileSize
	return r
}

func (r ApiAlipayEcoFilePathQueryRequest) Execute() (*AlipayEcoFilePathQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEcoFilePathQueryExecute(r)
}

/*
AlipayEcoFilePathQuery 获取文件直传地址（E签宝）

获取文件直传地址（E签宝）

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEcoFilePathQueryRequest
*/
func (r *AlipayEcoFilePathAPIService) AlipayEcoFilePathQuery(ctx context.Context) ApiAlipayEcoFilePathQueryRequest {
	return ApiAlipayEcoFilePathQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEcoFilePathQueryResponseModel
func (a *AlipayEcoFilePathAPIService) AlipayEcoFilePathQueryExecute(r ApiAlipayEcoFilePathQueryRequest) (*AlipayEcoFilePathQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEcoFilePathQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEcoFilePathAPIService.AlipayEcoFilePathQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/eco/file/path/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.targetAppId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "target_app_id", r.targetAppId, "form", "")
	}
	if r.contentMd5 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content_md_5", r.contentMd5, "form", "")
	}
	if r.contentType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content_type", r.contentType, "form", "")
	}
	if r.fileName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "file_name", r.fileName, "form", "")
	}
	if r.fileSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "file_size", r.fileSize, "form", "")
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
		var v AlipayEcoFilePathQueryDefaultResponse
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

func (a *AlipayEcoFilePathAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayEcoFilePathAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
