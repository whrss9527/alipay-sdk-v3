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

// AlipayMarketingActivityVoucherCreateDefaultResponse struct for AlipayMarketingActivityVoucherCreateDefaultResponse
type AlipayMarketingActivityVoucherCreateDefaultResponse struct {
	AlipayMarketingActivityVoucherCreateErrorResponseModel *AlipayMarketingActivityVoucherCreateErrorResponseModel
	CommonErrorType                                        *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingActivityVoucherCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingActivityVoucherCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingActivityVoucherCreateErrorResponseModel)
	if err == nil {
		jsonAlipayMarketingActivityVoucherCreateErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingActivityVoucherCreateErrorResponseModel)
		if string(jsonAlipayMarketingActivityVoucherCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingActivityVoucherCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingActivityVoucherCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingActivityVoucherCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingActivityVoucherCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingActivityVoucherCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingActivityVoucherCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingActivityVoucherCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayMarketingActivityVoucherCreateDefaultResponse struct {
	value *AlipayMarketingActivityVoucherCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingActivityVoucherCreateDefaultResponse) Get() *AlipayMarketingActivityVoucherCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingActivityVoucherCreateDefaultResponse) Set(val *AlipayMarketingActivityVoucherCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityVoucherCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityVoucherCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityVoucherCreateDefaultResponse(val *AlipayMarketingActivityVoucherCreateDefaultResponse) *NullableAlipayMarketingActivityVoucherCreateDefaultResponse {
	return &NullableAlipayMarketingActivityVoucherCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityVoucherCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityVoucherCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
