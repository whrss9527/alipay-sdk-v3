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

// AlipayOpenMiniInnerversionContentRollbackDefaultResponse struct for AlipayOpenMiniInnerversionContentRollbackDefaultResponse
type AlipayOpenMiniInnerversionContentRollbackDefaultResponse struct {
	AlipayOpenMiniInnerversionContentRollbackErrorResponseModel *AlipayOpenMiniInnerversionContentRollbackErrorResponseModel
	CommonErrorType                                             *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniInnerversionContentRollbackDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniInnerversionContentRollbackErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniInnerversionContentRollbackErrorResponseModel)
	if err == nil {
		jsonAlipayOpenMiniInnerversionContentRollbackErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniInnerversionContentRollbackErrorResponseModel)
		if string(jsonAlipayOpenMiniInnerversionContentRollbackErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniInnerversionContentRollbackErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniInnerversionContentRollbackErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniInnerversionContentRollbackErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniInnerversionContentRollbackDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniInnerversionContentRollbackDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniInnerversionContentRollbackErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniInnerversionContentRollbackErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenMiniInnerversionContentRollbackDefaultResponse struct {
	value *AlipayOpenMiniInnerversionContentRollbackDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionContentRollbackDefaultResponse) Get() *AlipayOpenMiniInnerversionContentRollbackDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionContentRollbackDefaultResponse) Set(val *AlipayOpenMiniInnerversionContentRollbackDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionContentRollbackDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionContentRollbackDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionContentRollbackDefaultResponse(val *AlipayOpenMiniInnerversionContentRollbackDefaultResponse) *NullableAlipayOpenMiniInnerversionContentRollbackDefaultResponse {
	return &NullableAlipayOpenMiniInnerversionContentRollbackDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionContentRollbackDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionContentRollbackDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
