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

// AlipayCommerceEcEnterpriseDeleteDefaultResponse struct for AlipayCommerceEcEnterpriseDeleteDefaultResponse
type AlipayCommerceEcEnterpriseDeleteDefaultResponse struct {
	AlipayCommerceEcEnterpriseDeleteErrorResponseModel *AlipayCommerceEcEnterpriseDeleteErrorResponseModel
	CommonErrorType                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceEcEnterpriseDeleteDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceEcEnterpriseDeleteErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceEcEnterpriseDeleteErrorResponseModel)
	if err == nil {
		jsonAlipayCommerceEcEnterpriseDeleteErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceEcEnterpriseDeleteErrorResponseModel)
		if string(jsonAlipayCommerceEcEnterpriseDeleteErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceEcEnterpriseDeleteErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceEcEnterpriseDeleteErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceEcEnterpriseDeleteErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceEcEnterpriseDeleteDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceEcEnterpriseDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceEcEnterpriseDeleteErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceEcEnterpriseDeleteErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayCommerceEcEnterpriseDeleteDefaultResponse struct {
	value *AlipayCommerceEcEnterpriseDeleteDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceEcEnterpriseDeleteDefaultResponse) Get() *AlipayCommerceEcEnterpriseDeleteDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceEcEnterpriseDeleteDefaultResponse) Set(val *AlipayCommerceEcEnterpriseDeleteDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcEnterpriseDeleteDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcEnterpriseDeleteDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcEnterpriseDeleteDefaultResponse(val *AlipayCommerceEcEnterpriseDeleteDefaultResponse) *NullableAlipayCommerceEcEnterpriseDeleteDefaultResponse {
	return &NullableAlipayCommerceEcEnterpriseDeleteDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcEnterpriseDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcEnterpriseDeleteDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
