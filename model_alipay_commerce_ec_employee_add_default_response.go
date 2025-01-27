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

// AlipayCommerceEcEmployeeAddDefaultResponse struct for AlipayCommerceEcEmployeeAddDefaultResponse
type AlipayCommerceEcEmployeeAddDefaultResponse struct {
	AlipayCommerceEcEmployeeAddErrorResponseModel *AlipayCommerceEcEmployeeAddErrorResponseModel
	CommonErrorType                               *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceEcEmployeeAddDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceEcEmployeeAddErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceEcEmployeeAddErrorResponseModel)
	if err == nil {
		jsonAlipayCommerceEcEmployeeAddErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceEcEmployeeAddErrorResponseModel)
		if string(jsonAlipayCommerceEcEmployeeAddErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceEcEmployeeAddErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceEcEmployeeAddErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceEcEmployeeAddErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceEcEmployeeAddDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceEcEmployeeAddDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceEcEmployeeAddErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceEcEmployeeAddErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayCommerceEcEmployeeAddDefaultResponse struct {
	value *AlipayCommerceEcEmployeeAddDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceEcEmployeeAddDefaultResponse) Get() *AlipayCommerceEcEmployeeAddDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceEcEmployeeAddDefaultResponse) Set(val *AlipayCommerceEcEmployeeAddDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcEmployeeAddDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcEmployeeAddDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcEmployeeAddDefaultResponse(val *AlipayCommerceEcEmployeeAddDefaultResponse) *NullableAlipayCommerceEcEmployeeAddDefaultResponse {
	return &NullableAlipayCommerceEcEmployeeAddDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcEmployeeAddDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcEmployeeAddDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
