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

// AlipayOpenPublicPersonalizedMenuCreateDefaultResponse struct for AlipayOpenPublicPersonalizedMenuCreateDefaultResponse
type AlipayOpenPublicPersonalizedMenuCreateDefaultResponse struct {
	AlipayOpenPublicPersonalizedMenuCreateErrorResponseModel *AlipayOpenPublicPersonalizedMenuCreateErrorResponseModel
	CommonErrorType                                          *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicPersonalizedMenuCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicPersonalizedMenuCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicPersonalizedMenuCreateErrorResponseModel)
	if err == nil {
		jsonAlipayOpenPublicPersonalizedMenuCreateErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicPersonalizedMenuCreateErrorResponseModel)
		if string(jsonAlipayOpenPublicPersonalizedMenuCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicPersonalizedMenuCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicPersonalizedMenuCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicPersonalizedMenuCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicPersonalizedMenuCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicPersonalizedMenuCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicPersonalizedMenuCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicPersonalizedMenuCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenPublicPersonalizedMenuCreateDefaultResponse struct {
	value *AlipayOpenPublicPersonalizedMenuCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicPersonalizedMenuCreateDefaultResponse) Get() *AlipayOpenPublicPersonalizedMenuCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicPersonalizedMenuCreateDefaultResponse) Set(val *AlipayOpenPublicPersonalizedMenuCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicPersonalizedMenuCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicPersonalizedMenuCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicPersonalizedMenuCreateDefaultResponse(val *AlipayOpenPublicPersonalizedMenuCreateDefaultResponse) *NullableAlipayOpenPublicPersonalizedMenuCreateDefaultResponse {
	return &NullableAlipayOpenPublicPersonalizedMenuCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicPersonalizedMenuCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicPersonalizedMenuCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
