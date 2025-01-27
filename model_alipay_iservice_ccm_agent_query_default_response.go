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

// AlipayIserviceCcmAgentQueryDefaultResponse struct for AlipayIserviceCcmAgentQueryDefaultResponse
type AlipayIserviceCcmAgentQueryDefaultResponse struct {
	AlipayIserviceCcmAgentQueryErrorResponseModel *AlipayIserviceCcmAgentQueryErrorResponseModel
	CommonErrorType                               *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayIserviceCcmAgentQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayIserviceCcmAgentQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayIserviceCcmAgentQueryErrorResponseModel)
	if err == nil {
		jsonAlipayIserviceCcmAgentQueryErrorResponseModel, _ := json.Marshal(dst.AlipayIserviceCcmAgentQueryErrorResponseModel)
		if string(jsonAlipayIserviceCcmAgentQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayIserviceCcmAgentQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayIserviceCcmAgentQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayIserviceCcmAgentQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayIserviceCcmAgentQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayIserviceCcmAgentQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayIserviceCcmAgentQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayIserviceCcmAgentQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayIserviceCcmAgentQueryDefaultResponse struct {
	value *AlipayIserviceCcmAgentQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayIserviceCcmAgentQueryDefaultResponse) Get() *AlipayIserviceCcmAgentQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayIserviceCcmAgentQueryDefaultResponse) Set(val *AlipayIserviceCcmAgentQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmAgentQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmAgentQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmAgentQueryDefaultResponse(val *AlipayIserviceCcmAgentQueryDefaultResponse) *NullableAlipayIserviceCcmAgentQueryDefaultResponse {
	return &NullableAlipayIserviceCcmAgentQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmAgentQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmAgentQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
