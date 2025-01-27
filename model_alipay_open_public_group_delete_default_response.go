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

// AlipayOpenPublicGroupDeleteDefaultResponse struct for AlipayOpenPublicGroupDeleteDefaultResponse
type AlipayOpenPublicGroupDeleteDefaultResponse struct {
	AlipayOpenPublicGroupDeleteErrorResponseModel *AlipayOpenPublicGroupDeleteErrorResponseModel
	CommonErrorType                               *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicGroupDeleteDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicGroupDeleteErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicGroupDeleteErrorResponseModel)
	if err == nil {
		jsonAlipayOpenPublicGroupDeleteErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicGroupDeleteErrorResponseModel)
		if string(jsonAlipayOpenPublicGroupDeleteErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicGroupDeleteErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicGroupDeleteErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicGroupDeleteErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicGroupDeleteDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicGroupDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicGroupDeleteErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicGroupDeleteErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenPublicGroupDeleteDefaultResponse struct {
	value *AlipayOpenPublicGroupDeleteDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicGroupDeleteDefaultResponse) Get() *AlipayOpenPublicGroupDeleteDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicGroupDeleteDefaultResponse) Set(val *AlipayOpenPublicGroupDeleteDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicGroupDeleteDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicGroupDeleteDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicGroupDeleteDefaultResponse(val *AlipayOpenPublicGroupDeleteDefaultResponse) *NullableAlipayOpenPublicGroupDeleteDefaultResponse {
	return &NullableAlipayOpenPublicGroupDeleteDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicGroupDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicGroupDeleteDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
