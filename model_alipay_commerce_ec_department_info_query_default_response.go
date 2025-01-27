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

// AlipayCommerceEcDepartmentInfoQueryDefaultResponse struct for AlipayCommerceEcDepartmentInfoQueryDefaultResponse
type AlipayCommerceEcDepartmentInfoQueryDefaultResponse struct {
	AlipayCommerceEcDepartmentInfoQueryErrorResponseModel *AlipayCommerceEcDepartmentInfoQueryErrorResponseModel
	CommonErrorType                                       *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceEcDepartmentInfoQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceEcDepartmentInfoQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceEcDepartmentInfoQueryErrorResponseModel)
	if err == nil {
		jsonAlipayCommerceEcDepartmentInfoQueryErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceEcDepartmentInfoQueryErrorResponseModel)
		if string(jsonAlipayCommerceEcDepartmentInfoQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceEcDepartmentInfoQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceEcDepartmentInfoQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceEcDepartmentInfoQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceEcDepartmentInfoQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceEcDepartmentInfoQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceEcDepartmentInfoQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceEcDepartmentInfoQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayCommerceEcDepartmentInfoQueryDefaultResponse struct {
	value *AlipayCommerceEcDepartmentInfoQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceEcDepartmentInfoQueryDefaultResponse) Get() *AlipayCommerceEcDepartmentInfoQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceEcDepartmentInfoQueryDefaultResponse) Set(val *AlipayCommerceEcDepartmentInfoQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcDepartmentInfoQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcDepartmentInfoQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcDepartmentInfoQueryDefaultResponse(val *AlipayCommerceEcDepartmentInfoQueryDefaultResponse) *NullableAlipayCommerceEcDepartmentInfoQueryDefaultResponse {
	return &NullableAlipayCommerceEcDepartmentInfoQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcDepartmentInfoQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcDepartmentInfoQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
