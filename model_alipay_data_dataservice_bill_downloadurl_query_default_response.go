/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)


// AlipayDataDataserviceBillDownloadurlQueryDefaultResponse struct for AlipayDataDataserviceBillDownloadurlQueryDefaultResponse
type AlipayDataDataserviceBillDownloadurlQueryDefaultResponse struct {
	AlipayDataDataserviceBillDownloadurlQueryErrorResponseModel *AlipayDataDataserviceBillDownloadurlQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayDataDataserviceBillDownloadurlQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayDataDataserviceBillDownloadurlQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayDataDataserviceBillDownloadurlQueryErrorResponseModel);
	if err == nil {
		jsonAlipayDataDataserviceBillDownloadurlQueryErrorResponseModel, _ := json.Marshal(dst.AlipayDataDataserviceBillDownloadurlQueryErrorResponseModel)
		if string(jsonAlipayDataDataserviceBillDownloadurlQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayDataDataserviceBillDownloadurlQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayDataDataserviceBillDownloadurlQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayDataDataserviceBillDownloadurlQueryErrorResponseModel = nil
	}

	// try to unmarshal JSON data into CommonErrorType
	err = json.Unmarshal(data, &dst.CommonErrorType);
	if err == nil {
		jsonCommonErrorType, _ := json.Marshal(dst.CommonErrorType)
		if string(jsonCommonErrorType) == "{}" { // empty struct
			dst.CommonErrorType = nil
		} else {
			return nil // data stored in dst.CommonErrorType, return on the first match
		}
	} else {
		dst.CommonErrorType = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayDataDataserviceBillDownloadurlQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayDataDataserviceBillDownloadurlQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayDataDataserviceBillDownloadurlQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayDataDataserviceBillDownloadurlQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayDataDataserviceBillDownloadurlQueryDefaultResponse struct {
	value *AlipayDataDataserviceBillDownloadurlQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayDataDataserviceBillDownloadurlQueryDefaultResponse) Get() *AlipayDataDataserviceBillDownloadurlQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayDataDataserviceBillDownloadurlQueryDefaultResponse) Set(val *AlipayDataDataserviceBillDownloadurlQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayDataDataserviceBillDownloadurlQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayDataDataserviceBillDownloadurlQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayDataDataserviceBillDownloadurlQueryDefaultResponse(val *AlipayDataDataserviceBillDownloadurlQueryDefaultResponse) *NullableAlipayDataDataserviceBillDownloadurlQueryDefaultResponse {
	return &NullableAlipayDataDataserviceBillDownloadurlQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayDataDataserviceBillDownloadurlQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayDataDataserviceBillDownloadurlQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

