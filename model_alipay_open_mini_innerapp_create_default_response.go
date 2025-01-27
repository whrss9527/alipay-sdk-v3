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

// AlipayOpenMiniInnerappCreateDefaultResponse struct for AlipayOpenMiniInnerappCreateDefaultResponse
type AlipayOpenMiniInnerappCreateDefaultResponse struct {
	AlipayOpenMiniInnerappCreateErrorResponseModel *AlipayOpenMiniInnerappCreateErrorResponseModel
	CommonErrorType                                *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniInnerappCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniInnerappCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniInnerappCreateErrorResponseModel)
	if err == nil {
		jsonAlipayOpenMiniInnerappCreateErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniInnerappCreateErrorResponseModel)
		if string(jsonAlipayOpenMiniInnerappCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniInnerappCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniInnerappCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniInnerappCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniInnerappCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniInnerappCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniInnerappCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniInnerappCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenMiniInnerappCreateDefaultResponse struct {
	value *AlipayOpenMiniInnerappCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniInnerappCreateDefaultResponse) Get() *AlipayOpenMiniInnerappCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerappCreateDefaultResponse) Set(val *AlipayOpenMiniInnerappCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerappCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerappCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerappCreateDefaultResponse(val *AlipayOpenMiniInnerappCreateDefaultResponse) *NullableAlipayOpenMiniInnerappCreateDefaultResponse {
	return &NullableAlipayOpenMiniInnerappCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerappCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerappCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
