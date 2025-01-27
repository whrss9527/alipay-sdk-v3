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

// AlipayFundEnterprisepayGroupQueryDefaultResponse struct for AlipayFundEnterprisepayGroupQueryDefaultResponse
type AlipayFundEnterprisepayGroupQueryDefaultResponse struct {
	AlipayFundEnterprisepayGroupQueryErrorResponseModel *AlipayFundEnterprisepayGroupQueryErrorResponseModel
	CommonErrorType                                     *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayFundEnterprisepayGroupQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayFundEnterprisepayGroupQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayFundEnterprisepayGroupQueryErrorResponseModel)
	if err == nil {
		jsonAlipayFundEnterprisepayGroupQueryErrorResponseModel, _ := json.Marshal(dst.AlipayFundEnterprisepayGroupQueryErrorResponseModel)
		if string(jsonAlipayFundEnterprisepayGroupQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayFundEnterprisepayGroupQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayFundEnterprisepayGroupQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayFundEnterprisepayGroupQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayFundEnterprisepayGroupQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayFundEnterprisepayGroupQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayFundEnterprisepayGroupQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayFundEnterprisepayGroupQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayFundEnterprisepayGroupQueryDefaultResponse struct {
	value *AlipayFundEnterprisepayGroupQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayFundEnterprisepayGroupQueryDefaultResponse) Get() *AlipayFundEnterprisepayGroupQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayFundEnterprisepayGroupQueryDefaultResponse) Set(val *AlipayFundEnterprisepayGroupQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundEnterprisepayGroupQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundEnterprisepayGroupQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundEnterprisepayGroupQueryDefaultResponse(val *AlipayFundEnterprisepayGroupQueryDefaultResponse) *NullableAlipayFundEnterprisepayGroupQueryDefaultResponse {
	return &NullableAlipayFundEnterprisepayGroupQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayFundEnterprisepayGroupQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundEnterprisepayGroupQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
