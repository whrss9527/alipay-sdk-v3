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
	"os"
)


// AlipayOpenAgentFacetofaceAPIService AlipayOpenAgentFacetofaceAPI service
type AlipayOpenAgentFacetofaceAPIService service

type ApiAlipayOpenAgentFacetofaceSignRequest struct {
	ctx context.Context
	ApiService *AlipayOpenAgentFacetofaceAPIService
	businessLicenseAuthPic *os.File
	businessLicensePic *os.File
	data *AlipayOpenAgentFacetofaceSignModel
	shopScenePic *os.File
	shopSignBoardPic *os.File
	specialLicensePic *os.File
}

func (r ApiAlipayOpenAgentFacetofaceSignRequest) BusinessLicenseAuthPic(businessLicenseAuthPic *os.File) ApiAlipayOpenAgentFacetofaceSignRequest {
	r.businessLicenseAuthPic = businessLicenseAuthPic
	return r
}

func (r ApiAlipayOpenAgentFacetofaceSignRequest) BusinessLicensePic(businessLicensePic *os.File) ApiAlipayOpenAgentFacetofaceSignRequest {
	r.businessLicensePic = businessLicensePic
	return r
}

func (r ApiAlipayOpenAgentFacetofaceSignRequest) Data(data AlipayOpenAgentFacetofaceSignModel) ApiAlipayOpenAgentFacetofaceSignRequest {
	r.data = &data
	return r
}

func (r ApiAlipayOpenAgentFacetofaceSignRequest) ShopScenePic(shopScenePic *os.File) ApiAlipayOpenAgentFacetofaceSignRequest {
	r.shopScenePic = shopScenePic
	return r
}

func (r ApiAlipayOpenAgentFacetofaceSignRequest) ShopSignBoardPic(shopSignBoardPic *os.File) ApiAlipayOpenAgentFacetofaceSignRequest {
	r.shopSignBoardPic = shopSignBoardPic
	return r
}

func (r ApiAlipayOpenAgentFacetofaceSignRequest) SpecialLicensePic(specialLicensePic *os.File) ApiAlipayOpenAgentFacetofaceSignRequest {
	r.specialLicensePic = specialLicensePic
	return r
}

func (r ApiAlipayOpenAgentFacetofaceSignRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayOpenAgentFacetofaceSignExecute(r)
}

/*
AlipayOpenAgentFacetofaceSign 代签约当面付产品

三方应用代理签约当面付产品，需要配合开启事务接口使用

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayOpenAgentFacetofaceSignRequest
*/
func (r *AlipayOpenAgentFacetofaceAPIService) AlipayOpenAgentFacetofaceSign(ctx context.Context) ApiAlipayOpenAgentFacetofaceSignRequest {
	return ApiAlipayOpenAgentFacetofaceSignRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AlipayOpenAgentFacetofaceAPIService) AlipayOpenAgentFacetofaceSignExecute(r ApiAlipayOpenAgentFacetofaceSignRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenAgentFacetofaceAPIService.AlipayOpenAgentFacetofaceSign")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/agent/facetoface/sign"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}



	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var businessLicenseAuthPicLocalVarFormFileName string
	var businessLicenseAuthPicLocalVarFileName     string
	var businessLicenseAuthPicLocalVarFileBytes    []byte

	businessLicenseAuthPicLocalVarFormFileName = "business_license_auth_pic"
	businessLicenseAuthPicLocalVarFile := r.businessLicenseAuthPic

	if businessLicenseAuthPicLocalVarFile != nil {
		fbs, _ := io.ReadAll(businessLicenseAuthPicLocalVarFile)

		businessLicenseAuthPicLocalVarFileBytes = fbs
		businessLicenseAuthPicLocalVarFileName = businessLicenseAuthPicLocalVarFile.Name()
		businessLicenseAuthPicLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: businessLicenseAuthPicLocalVarFileBytes, fileName: businessLicenseAuthPicLocalVarFileName, formFileName: businessLicenseAuthPicLocalVarFormFileName})
	}
	var businessLicensePicLocalVarFormFileName string
	var businessLicensePicLocalVarFileName     string
	var businessLicensePicLocalVarFileBytes    []byte

	businessLicensePicLocalVarFormFileName = "business_license_pic"
	businessLicensePicLocalVarFile := r.businessLicensePic

	if businessLicensePicLocalVarFile != nil {
		fbs, _ := io.ReadAll(businessLicensePicLocalVarFile)

		businessLicensePicLocalVarFileBytes = fbs
		businessLicensePicLocalVarFileName = businessLicensePicLocalVarFile.Name()
		businessLicensePicLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: businessLicensePicLocalVarFileBytes, fileName: businessLicensePicLocalVarFileName, formFileName: businessLicensePicLocalVarFormFileName})
	}
	if r.data != nil {
		paramJson, err := parameterToJson(*r.data)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("data", paramJson)
	}
	var shopScenePicLocalVarFormFileName string
	var shopScenePicLocalVarFileName     string
	var shopScenePicLocalVarFileBytes    []byte

	shopScenePicLocalVarFormFileName = "shop_scene_pic"
	shopScenePicLocalVarFile := r.shopScenePic

	if shopScenePicLocalVarFile != nil {
		fbs, _ := io.ReadAll(shopScenePicLocalVarFile)

		shopScenePicLocalVarFileBytes = fbs
		shopScenePicLocalVarFileName = shopScenePicLocalVarFile.Name()
		shopScenePicLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: shopScenePicLocalVarFileBytes, fileName: shopScenePicLocalVarFileName, formFileName: shopScenePicLocalVarFormFileName})
	}
	var shopSignBoardPicLocalVarFormFileName string
	var shopSignBoardPicLocalVarFileName     string
	var shopSignBoardPicLocalVarFileBytes    []byte

	shopSignBoardPicLocalVarFormFileName = "shop_sign_board_pic"
	shopSignBoardPicLocalVarFile := r.shopSignBoardPic

	if shopSignBoardPicLocalVarFile != nil {
		fbs, _ := io.ReadAll(shopSignBoardPicLocalVarFile)

		shopSignBoardPicLocalVarFileBytes = fbs
		shopSignBoardPicLocalVarFileName = shopSignBoardPicLocalVarFile.Name()
		shopSignBoardPicLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: shopSignBoardPicLocalVarFileBytes, fileName: shopSignBoardPicLocalVarFileName, formFileName: shopSignBoardPicLocalVarFormFileName})
	}
	var specialLicensePicLocalVarFormFileName string
	var specialLicensePicLocalVarFileName     string
	var specialLicensePicLocalVarFileBytes    []byte

	specialLicensePicLocalVarFormFileName = "special_license_pic"
	specialLicensePicLocalVarFile := r.specialLicensePic

	if specialLicensePicLocalVarFile != nil {
		fbs, _ := io.ReadAll(specialLicensePicLocalVarFile)

		specialLicensePicLocalVarFileBytes = fbs
		specialLicensePicLocalVarFileName = specialLicensePicLocalVarFile.Name()
		specialLicensePicLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: specialLicensePicLocalVarFileBytes, fileName: specialLicensePicLocalVarFileName, formFileName: specialLicensePicLocalVarFormFileName})
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
			var v AlipayOpenAgentFacetofaceSignDefaultResponse
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


func (a *AlipayOpenAgentFacetofaceAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayOpenAgentFacetofaceAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}

