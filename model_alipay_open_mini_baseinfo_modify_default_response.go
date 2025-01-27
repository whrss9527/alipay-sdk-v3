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

// AlipayOpenMiniBaseinfoModifyDefaultResponse struct for AlipayOpenMiniBaseinfoModifyDefaultResponse
type AlipayOpenMiniBaseinfoModifyDefaultResponse struct {
	AlipayOpenMiniBaseinfoModifyErrorResponseModel *AlipayOpenMiniBaseinfoModifyErrorResponseModel
	CommonErrorType                                *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniBaseinfoModifyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniBaseinfoModifyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniBaseinfoModifyErrorResponseModel)
	if err == nil {
		jsonAlipayOpenMiniBaseinfoModifyErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniBaseinfoModifyErrorResponseModel)
		if string(jsonAlipayOpenMiniBaseinfoModifyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniBaseinfoModifyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniBaseinfoModifyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniBaseinfoModifyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniBaseinfoModifyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniBaseinfoModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniBaseinfoModifyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniBaseinfoModifyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenMiniBaseinfoModifyDefaultResponse struct {
	value *AlipayOpenMiniBaseinfoModifyDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniBaseinfoModifyDefaultResponse) Get() *AlipayOpenMiniBaseinfoModifyDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniBaseinfoModifyDefaultResponse) Set(val *AlipayOpenMiniBaseinfoModifyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniBaseinfoModifyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniBaseinfoModifyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniBaseinfoModifyDefaultResponse(val *AlipayOpenMiniBaseinfoModifyDefaultResponse) *NullableAlipayOpenMiniBaseinfoModifyDefaultResponse {
	return &NullableAlipayOpenMiniBaseinfoModifyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniBaseinfoModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniBaseinfoModifyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
