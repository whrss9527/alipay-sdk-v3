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

// AlipayOpenMiniMiniappBrandQueryDefaultResponse struct for AlipayOpenMiniMiniappBrandQueryDefaultResponse
type AlipayOpenMiniMiniappBrandQueryDefaultResponse struct {
	AlipayOpenMiniMiniappBrandQueryErrorResponseModel *AlipayOpenMiniMiniappBrandQueryErrorResponseModel
	CommonErrorType                                   *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniMiniappBrandQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniMiniappBrandQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniMiniappBrandQueryErrorResponseModel)
	if err == nil {
		jsonAlipayOpenMiniMiniappBrandQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniMiniappBrandQueryErrorResponseModel)
		if string(jsonAlipayOpenMiniMiniappBrandQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniMiniappBrandQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniMiniappBrandQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniMiniappBrandQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniMiniappBrandQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniMiniappBrandQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniMiniappBrandQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniMiniappBrandQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenMiniMiniappBrandQueryDefaultResponse struct {
	value *AlipayOpenMiniMiniappBrandQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniMiniappBrandQueryDefaultResponse) Get() *AlipayOpenMiniMiniappBrandQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniMiniappBrandQueryDefaultResponse) Set(val *AlipayOpenMiniMiniappBrandQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniMiniappBrandQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniMiniappBrandQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniMiniappBrandQueryDefaultResponse(val *AlipayOpenMiniMiniappBrandQueryDefaultResponse) *NullableAlipayOpenMiniMiniappBrandQueryDefaultResponse {
	return &NullableAlipayOpenMiniMiniappBrandQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniMiniappBrandQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniMiniappBrandQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
