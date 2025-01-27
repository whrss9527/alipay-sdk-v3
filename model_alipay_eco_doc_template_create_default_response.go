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

// AlipayEcoDocTemplateCreateDefaultResponse struct for AlipayEcoDocTemplateCreateDefaultResponse
type AlipayEcoDocTemplateCreateDefaultResponse struct {
	AlipayEcoDocTemplateCreateErrorResponseModel *AlipayEcoDocTemplateCreateErrorResponseModel
	CommonErrorType                              *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEcoDocTemplateCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEcoDocTemplateCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEcoDocTemplateCreateErrorResponseModel)
	if err == nil {
		jsonAlipayEcoDocTemplateCreateErrorResponseModel, _ := json.Marshal(dst.AlipayEcoDocTemplateCreateErrorResponseModel)
		if string(jsonAlipayEcoDocTemplateCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEcoDocTemplateCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEcoDocTemplateCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEcoDocTemplateCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEcoDocTemplateCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEcoDocTemplateCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEcoDocTemplateCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEcoDocTemplateCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayEcoDocTemplateCreateDefaultResponse struct {
	value *AlipayEcoDocTemplateCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayEcoDocTemplateCreateDefaultResponse) Get() *AlipayEcoDocTemplateCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayEcoDocTemplateCreateDefaultResponse) Set(val *AlipayEcoDocTemplateCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoDocTemplateCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoDocTemplateCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoDocTemplateCreateDefaultResponse(val *AlipayEcoDocTemplateCreateDefaultResponse) *NullableAlipayEcoDocTemplateCreateDefaultResponse {
	return &NullableAlipayEcoDocTemplateCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEcoDocTemplateCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoDocTemplateCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
