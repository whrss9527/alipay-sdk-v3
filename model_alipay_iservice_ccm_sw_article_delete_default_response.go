/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"encoding/json"
	"fmt"
)

// AlipayIserviceCcmSwArticleDeleteDefaultResponse struct for AlipayIserviceCcmSwArticleDeleteDefaultResponse
type AlipayIserviceCcmSwArticleDeleteDefaultResponse struct {
	AlipayIserviceCcmSwArticleDeleteErrorResponseModel *AlipayIserviceCcmSwArticleDeleteErrorResponseModel
	CommonErrorType                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayIserviceCcmSwArticleDeleteDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayIserviceCcmSwArticleDeleteErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayIserviceCcmSwArticleDeleteErrorResponseModel)
	if err == nil {
		jsonAlipayIserviceCcmSwArticleDeleteErrorResponseModel, _ := json.Marshal(dst.AlipayIserviceCcmSwArticleDeleteErrorResponseModel)
		if string(jsonAlipayIserviceCcmSwArticleDeleteErrorResponseModel) == "{}" { // empty struct
			dst.AlipayIserviceCcmSwArticleDeleteErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayIserviceCcmSwArticleDeleteErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayIserviceCcmSwArticleDeleteErrorResponseModel = nil
	}

	// try to unmarshal JSON data into CommonErrorType
	err = json.Unmarshal(data, &dst.CommonErrorType)
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayIserviceCcmSwArticleDeleteDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayIserviceCcmSwArticleDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayIserviceCcmSwArticleDeleteErrorResponseModel != nil {
		return json.Marshal(&src.AlipayIserviceCcmSwArticleDeleteErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayIserviceCcmSwArticleDeleteDefaultResponse struct {
	value *AlipayIserviceCcmSwArticleDeleteDefaultResponse
	isSet bool
}

func (v NullableAlipayIserviceCcmSwArticleDeleteDefaultResponse) Get() *AlipayIserviceCcmSwArticleDeleteDefaultResponse {
	return v.value
}

func (v *NullableAlipayIserviceCcmSwArticleDeleteDefaultResponse) Set(val *AlipayIserviceCcmSwArticleDeleteDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmSwArticleDeleteDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmSwArticleDeleteDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmSwArticleDeleteDefaultResponse(val *AlipayIserviceCcmSwArticleDeleteDefaultResponse) *NullableAlipayIserviceCcmSwArticleDeleteDefaultResponse {
	return &NullableAlipayIserviceCcmSwArticleDeleteDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmSwArticleDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmSwArticleDeleteDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
