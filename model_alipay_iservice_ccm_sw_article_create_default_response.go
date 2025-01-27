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

// AlipayIserviceCcmSwArticleCreateDefaultResponse struct for AlipayIserviceCcmSwArticleCreateDefaultResponse
type AlipayIserviceCcmSwArticleCreateDefaultResponse struct {
	AlipayIserviceCcmSwArticleCreateErrorResponseModel *AlipayIserviceCcmSwArticleCreateErrorResponseModel
	CommonErrorType                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayIserviceCcmSwArticleCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayIserviceCcmSwArticleCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayIserviceCcmSwArticleCreateErrorResponseModel)
	if err == nil {
		jsonAlipayIserviceCcmSwArticleCreateErrorResponseModel, _ := json.Marshal(dst.AlipayIserviceCcmSwArticleCreateErrorResponseModel)
		if string(jsonAlipayIserviceCcmSwArticleCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayIserviceCcmSwArticleCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayIserviceCcmSwArticleCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayIserviceCcmSwArticleCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayIserviceCcmSwArticleCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayIserviceCcmSwArticleCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayIserviceCcmSwArticleCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayIserviceCcmSwArticleCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayIserviceCcmSwArticleCreateDefaultResponse struct {
	value *AlipayIserviceCcmSwArticleCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayIserviceCcmSwArticleCreateDefaultResponse) Get() *AlipayIserviceCcmSwArticleCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayIserviceCcmSwArticleCreateDefaultResponse) Set(val *AlipayIserviceCcmSwArticleCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmSwArticleCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmSwArticleCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmSwArticleCreateDefaultResponse(val *AlipayIserviceCcmSwArticleCreateDefaultResponse) *NullableAlipayIserviceCcmSwArticleCreateDefaultResponse {
	return &NullableAlipayIserviceCcmSwArticleCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmSwArticleCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmSwArticleCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
