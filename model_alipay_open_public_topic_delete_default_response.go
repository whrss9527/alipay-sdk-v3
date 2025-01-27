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

// AlipayOpenPublicTopicDeleteDefaultResponse struct for AlipayOpenPublicTopicDeleteDefaultResponse
type AlipayOpenPublicTopicDeleteDefaultResponse struct {
	AlipayOpenPublicTopicDeleteErrorResponseModel *AlipayOpenPublicTopicDeleteErrorResponseModel
	CommonErrorType                               *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicTopicDeleteDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicTopicDeleteErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicTopicDeleteErrorResponseModel)
	if err == nil {
		jsonAlipayOpenPublicTopicDeleteErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicTopicDeleteErrorResponseModel)
		if string(jsonAlipayOpenPublicTopicDeleteErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicTopicDeleteErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicTopicDeleteErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicTopicDeleteErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicTopicDeleteDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicTopicDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicTopicDeleteErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicTopicDeleteErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenPublicTopicDeleteDefaultResponse struct {
	value *AlipayOpenPublicTopicDeleteDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicTopicDeleteDefaultResponse) Get() *AlipayOpenPublicTopicDeleteDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicTopicDeleteDefaultResponse) Set(val *AlipayOpenPublicTopicDeleteDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicTopicDeleteDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicTopicDeleteDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicTopicDeleteDefaultResponse(val *AlipayOpenPublicTopicDeleteDefaultResponse) *NullableAlipayOpenPublicTopicDeleteDefaultResponse {
	return &NullableAlipayOpenPublicTopicDeleteDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicTopicDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicTopicDeleteDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
