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

// AlipayFundAuthOrderVoucherCreateDefaultResponse struct for AlipayFundAuthOrderVoucherCreateDefaultResponse
type AlipayFundAuthOrderVoucherCreateDefaultResponse struct {
	AlipayFundAuthOrderVoucherCreateErrorResponseModel *AlipayFundAuthOrderVoucherCreateErrorResponseModel
	CommonErrorType                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayFundAuthOrderVoucherCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayFundAuthOrderVoucherCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayFundAuthOrderVoucherCreateErrorResponseModel)
	if err == nil {
		jsonAlipayFundAuthOrderVoucherCreateErrorResponseModel, _ := json.Marshal(dst.AlipayFundAuthOrderVoucherCreateErrorResponseModel)
		if string(jsonAlipayFundAuthOrderVoucherCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayFundAuthOrderVoucherCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayFundAuthOrderVoucherCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayFundAuthOrderVoucherCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayFundAuthOrderVoucherCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayFundAuthOrderVoucherCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayFundAuthOrderVoucherCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayFundAuthOrderVoucherCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayFundAuthOrderVoucherCreateDefaultResponse struct {
	value *AlipayFundAuthOrderVoucherCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayFundAuthOrderVoucherCreateDefaultResponse) Get() *AlipayFundAuthOrderVoucherCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayFundAuthOrderVoucherCreateDefaultResponse) Set(val *AlipayFundAuthOrderVoucherCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundAuthOrderVoucherCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundAuthOrderVoucherCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundAuthOrderVoucherCreateDefaultResponse(val *AlipayFundAuthOrderVoucherCreateDefaultResponse) *NullableAlipayFundAuthOrderVoucherCreateDefaultResponse {
	return &NullableAlipayFundAuthOrderVoucherCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayFundAuthOrderVoucherCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundAuthOrderVoucherCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
