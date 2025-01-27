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

// AlipayOpenPublicMessageSingleSendDefaultResponse struct for AlipayOpenPublicMessageSingleSendDefaultResponse
type AlipayOpenPublicMessageSingleSendDefaultResponse struct {
	AlipayOpenPublicMessageSingleSendErrorResponseModel *AlipayOpenPublicMessageSingleSendErrorResponseModel
	CommonErrorType                                     *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicMessageSingleSendDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicMessageSingleSendErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicMessageSingleSendErrorResponseModel)
	if err == nil {
		jsonAlipayOpenPublicMessageSingleSendErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicMessageSingleSendErrorResponseModel)
		if string(jsonAlipayOpenPublicMessageSingleSendErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicMessageSingleSendErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicMessageSingleSendErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicMessageSingleSendErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicMessageSingleSendDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicMessageSingleSendDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicMessageSingleSendErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicMessageSingleSendErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenPublicMessageSingleSendDefaultResponse struct {
	value *AlipayOpenPublicMessageSingleSendDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicMessageSingleSendDefaultResponse) Get() *AlipayOpenPublicMessageSingleSendDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicMessageSingleSendDefaultResponse) Set(val *AlipayOpenPublicMessageSingleSendDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicMessageSingleSendDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicMessageSingleSendDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicMessageSingleSendDefaultResponse(val *AlipayOpenPublicMessageSingleSendDefaultResponse) *NullableAlipayOpenPublicMessageSingleSendDefaultResponse {
	return &NullableAlipayOpenPublicMessageSingleSendDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicMessageSingleSendDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicMessageSingleSendDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
