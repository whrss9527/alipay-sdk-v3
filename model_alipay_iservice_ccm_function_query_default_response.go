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

// AlipayIserviceCcmFunctionQueryDefaultResponse struct for AlipayIserviceCcmFunctionQueryDefaultResponse
type AlipayIserviceCcmFunctionQueryDefaultResponse struct {
	AlipayIserviceCcmFunctionQueryErrorResponseModel *AlipayIserviceCcmFunctionQueryErrorResponseModel
	CommonErrorType                                  *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayIserviceCcmFunctionQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayIserviceCcmFunctionQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayIserviceCcmFunctionQueryErrorResponseModel)
	if err == nil {
		jsonAlipayIserviceCcmFunctionQueryErrorResponseModel, _ := json.Marshal(dst.AlipayIserviceCcmFunctionQueryErrorResponseModel)
		if string(jsonAlipayIserviceCcmFunctionQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayIserviceCcmFunctionQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayIserviceCcmFunctionQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayIserviceCcmFunctionQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayIserviceCcmFunctionQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayIserviceCcmFunctionQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayIserviceCcmFunctionQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayIserviceCcmFunctionQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayIserviceCcmFunctionQueryDefaultResponse struct {
	value *AlipayIserviceCcmFunctionQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayIserviceCcmFunctionQueryDefaultResponse) Get() *AlipayIserviceCcmFunctionQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayIserviceCcmFunctionQueryDefaultResponse) Set(val *AlipayIserviceCcmFunctionQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmFunctionQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmFunctionQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmFunctionQueryDefaultResponse(val *AlipayIserviceCcmFunctionQueryDefaultResponse) *NullableAlipayIserviceCcmFunctionQueryDefaultResponse {
	return &NullableAlipayIserviceCcmFunctionQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmFunctionQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmFunctionQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
