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

// AlipayFundAccountbookNotifyQueryDefaultResponse struct for AlipayFundAccountbookNotifyQueryDefaultResponse
type AlipayFundAccountbookNotifyQueryDefaultResponse struct {
	AlipayFundAccountbookNotifyQueryErrorResponseModel *AlipayFundAccountbookNotifyQueryErrorResponseModel
	CommonErrorType                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayFundAccountbookNotifyQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayFundAccountbookNotifyQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayFundAccountbookNotifyQueryErrorResponseModel)
	if err == nil {
		jsonAlipayFundAccountbookNotifyQueryErrorResponseModel, _ := json.Marshal(dst.AlipayFundAccountbookNotifyQueryErrorResponseModel)
		if string(jsonAlipayFundAccountbookNotifyQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayFundAccountbookNotifyQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayFundAccountbookNotifyQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayFundAccountbookNotifyQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayFundAccountbookNotifyQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayFundAccountbookNotifyQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayFundAccountbookNotifyQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayFundAccountbookNotifyQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayFundAccountbookNotifyQueryDefaultResponse struct {
	value *AlipayFundAccountbookNotifyQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayFundAccountbookNotifyQueryDefaultResponse) Get() *AlipayFundAccountbookNotifyQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayFundAccountbookNotifyQueryDefaultResponse) Set(val *AlipayFundAccountbookNotifyQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundAccountbookNotifyQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundAccountbookNotifyQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundAccountbookNotifyQueryDefaultResponse(val *AlipayFundAccountbookNotifyQueryDefaultResponse) *NullableAlipayFundAccountbookNotifyQueryDefaultResponse {
	return &NullableAlipayFundAccountbookNotifyQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayFundAccountbookNotifyQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundAccountbookNotifyQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
