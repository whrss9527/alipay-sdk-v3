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

// AlipayCommerceEcDepartmentDeleteDefaultResponse struct for AlipayCommerceEcDepartmentDeleteDefaultResponse
type AlipayCommerceEcDepartmentDeleteDefaultResponse struct {
	AlipayCommerceEcDepartmentDeleteErrorResponseModel *AlipayCommerceEcDepartmentDeleteErrorResponseModel
	CommonErrorType                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceEcDepartmentDeleteDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceEcDepartmentDeleteErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceEcDepartmentDeleteErrorResponseModel)
	if err == nil {
		jsonAlipayCommerceEcDepartmentDeleteErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceEcDepartmentDeleteErrorResponseModel)
		if string(jsonAlipayCommerceEcDepartmentDeleteErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceEcDepartmentDeleteErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceEcDepartmentDeleteErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceEcDepartmentDeleteErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceEcDepartmentDeleteDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceEcDepartmentDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceEcDepartmentDeleteErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceEcDepartmentDeleteErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayCommerceEcDepartmentDeleteDefaultResponse struct {
	value *AlipayCommerceEcDepartmentDeleteDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceEcDepartmentDeleteDefaultResponse) Get() *AlipayCommerceEcDepartmentDeleteDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceEcDepartmentDeleteDefaultResponse) Set(val *AlipayCommerceEcDepartmentDeleteDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcDepartmentDeleteDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcDepartmentDeleteDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcDepartmentDeleteDefaultResponse(val *AlipayCommerceEcDepartmentDeleteDefaultResponse) *NullableAlipayCommerceEcDepartmentDeleteDefaultResponse {
	return &NullableAlipayCommerceEcDepartmentDeleteDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcDepartmentDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcDepartmentDeleteDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
