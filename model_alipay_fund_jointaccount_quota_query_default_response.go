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

// AlipayFundJointaccountQuotaQueryDefaultResponse struct for AlipayFundJointaccountQuotaQueryDefaultResponse
type AlipayFundJointaccountQuotaQueryDefaultResponse struct {
	AlipayFundJointaccountQuotaQueryErrorResponseModel *AlipayFundJointaccountQuotaQueryErrorResponseModel
	CommonErrorType                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayFundJointaccountQuotaQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayFundJointaccountQuotaQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayFundJointaccountQuotaQueryErrorResponseModel)
	if err == nil {
		jsonAlipayFundJointaccountQuotaQueryErrorResponseModel, _ := json.Marshal(dst.AlipayFundJointaccountQuotaQueryErrorResponseModel)
		if string(jsonAlipayFundJointaccountQuotaQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayFundJointaccountQuotaQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayFundJointaccountQuotaQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayFundJointaccountQuotaQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayFundJointaccountQuotaQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayFundJointaccountQuotaQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayFundJointaccountQuotaQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayFundJointaccountQuotaQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayFundJointaccountQuotaQueryDefaultResponse struct {
	value *AlipayFundJointaccountQuotaQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayFundJointaccountQuotaQueryDefaultResponse) Get() *AlipayFundJointaccountQuotaQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayFundJointaccountQuotaQueryDefaultResponse) Set(val *AlipayFundJointaccountQuotaQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundJointaccountQuotaQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundJointaccountQuotaQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundJointaccountQuotaQueryDefaultResponse(val *AlipayFundJointaccountQuotaQueryDefaultResponse) *NullableAlipayFundJointaccountQuotaQueryDefaultResponse {
	return &NullableAlipayFundJointaccountQuotaQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayFundJointaccountQuotaQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundJointaccountQuotaQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
