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

// AlipayMobilePublicMenuGetDefaultResponse struct for AlipayMobilePublicMenuGetDefaultResponse
type AlipayMobilePublicMenuGetDefaultResponse struct {
	AlipayMobilePublicMenuGetErrorResponseModel *AlipayMobilePublicMenuGetErrorResponseModel
	CommonErrorType                             *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMobilePublicMenuGetDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMobilePublicMenuGetErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMobilePublicMenuGetErrorResponseModel)
	if err == nil {
		jsonAlipayMobilePublicMenuGetErrorResponseModel, _ := json.Marshal(dst.AlipayMobilePublicMenuGetErrorResponseModel)
		if string(jsonAlipayMobilePublicMenuGetErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMobilePublicMenuGetErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMobilePublicMenuGetErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMobilePublicMenuGetErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMobilePublicMenuGetDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMobilePublicMenuGetDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMobilePublicMenuGetErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMobilePublicMenuGetErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayMobilePublicMenuGetDefaultResponse struct {
	value *AlipayMobilePublicMenuGetDefaultResponse
	isSet bool
}

func (v NullableAlipayMobilePublicMenuGetDefaultResponse) Get() *AlipayMobilePublicMenuGetDefaultResponse {
	return v.value
}

func (v *NullableAlipayMobilePublicMenuGetDefaultResponse) Set(val *AlipayMobilePublicMenuGetDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMobilePublicMenuGetDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMobilePublicMenuGetDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMobilePublicMenuGetDefaultResponse(val *AlipayMobilePublicMenuGetDefaultResponse) *NullableAlipayMobilePublicMenuGetDefaultResponse {
	return &NullableAlipayMobilePublicMenuGetDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMobilePublicMenuGetDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMobilePublicMenuGetDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
