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

// AlipayOfflineMaterialImageUploadDefaultResponse struct for AlipayOfflineMaterialImageUploadDefaultResponse
type AlipayOfflineMaterialImageUploadDefaultResponse struct {
	AlipayOfflineMaterialImageUploadErrorResponseModel *AlipayOfflineMaterialImageUploadErrorResponseModel
	CommonErrorType                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOfflineMaterialImageUploadDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOfflineMaterialImageUploadErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOfflineMaterialImageUploadErrorResponseModel)
	if err == nil {
		jsonAlipayOfflineMaterialImageUploadErrorResponseModel, _ := json.Marshal(dst.AlipayOfflineMaterialImageUploadErrorResponseModel)
		if string(jsonAlipayOfflineMaterialImageUploadErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOfflineMaterialImageUploadErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOfflineMaterialImageUploadErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOfflineMaterialImageUploadErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOfflineMaterialImageUploadDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOfflineMaterialImageUploadDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOfflineMaterialImageUploadErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOfflineMaterialImageUploadErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOfflineMaterialImageUploadDefaultResponse struct {
	value *AlipayOfflineMaterialImageUploadDefaultResponse
	isSet bool
}

func (v NullableAlipayOfflineMaterialImageUploadDefaultResponse) Get() *AlipayOfflineMaterialImageUploadDefaultResponse {
	return v.value
}

func (v *NullableAlipayOfflineMaterialImageUploadDefaultResponse) Set(val *AlipayOfflineMaterialImageUploadDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOfflineMaterialImageUploadDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOfflineMaterialImageUploadDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOfflineMaterialImageUploadDefaultResponse(val *AlipayOfflineMaterialImageUploadDefaultResponse) *NullableAlipayOfflineMaterialImageUploadDefaultResponse {
	return &NullableAlipayOfflineMaterialImageUploadDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOfflineMaterialImageUploadDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOfflineMaterialImageUploadDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
