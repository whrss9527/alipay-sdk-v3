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

// AlipayOpenSearchAppkeywordApplyDefaultResponse struct for AlipayOpenSearchAppkeywordApplyDefaultResponse
type AlipayOpenSearchAppkeywordApplyDefaultResponse struct {
	AlipayOpenSearchAppkeywordApplyErrorResponseModel *AlipayOpenSearchAppkeywordApplyErrorResponseModel
	CommonErrorType                                   *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenSearchAppkeywordApplyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenSearchAppkeywordApplyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenSearchAppkeywordApplyErrorResponseModel)
	if err == nil {
		jsonAlipayOpenSearchAppkeywordApplyErrorResponseModel, _ := json.Marshal(dst.AlipayOpenSearchAppkeywordApplyErrorResponseModel)
		if string(jsonAlipayOpenSearchAppkeywordApplyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenSearchAppkeywordApplyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenSearchAppkeywordApplyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenSearchAppkeywordApplyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenSearchAppkeywordApplyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenSearchAppkeywordApplyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenSearchAppkeywordApplyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenSearchAppkeywordApplyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenSearchAppkeywordApplyDefaultResponse struct {
	value *AlipayOpenSearchAppkeywordApplyDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenSearchAppkeywordApplyDefaultResponse) Get() *AlipayOpenSearchAppkeywordApplyDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenSearchAppkeywordApplyDefaultResponse) Set(val *AlipayOpenSearchAppkeywordApplyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchAppkeywordApplyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchAppkeywordApplyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchAppkeywordApplyDefaultResponse(val *AlipayOpenSearchAppkeywordApplyDefaultResponse) *NullableAlipayOpenSearchAppkeywordApplyDefaultResponse {
	return &NullableAlipayOpenSearchAppkeywordApplyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchAppkeywordApplyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchAppkeywordApplyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
