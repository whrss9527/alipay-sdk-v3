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

// AlipayCommerceEcEnterpriseCreateDefaultResponse struct for AlipayCommerceEcEnterpriseCreateDefaultResponse
type AlipayCommerceEcEnterpriseCreateDefaultResponse struct {
	AlipayCommerceEcEnterpriseCreateErrorResponseModel *AlipayCommerceEcEnterpriseCreateErrorResponseModel
	CommonErrorType                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceEcEnterpriseCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceEcEnterpriseCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceEcEnterpriseCreateErrorResponseModel)
	if err == nil {
		jsonAlipayCommerceEcEnterpriseCreateErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceEcEnterpriseCreateErrorResponseModel)
		if string(jsonAlipayCommerceEcEnterpriseCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceEcEnterpriseCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceEcEnterpriseCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceEcEnterpriseCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceEcEnterpriseCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceEcEnterpriseCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceEcEnterpriseCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceEcEnterpriseCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayCommerceEcEnterpriseCreateDefaultResponse struct {
	value *AlipayCommerceEcEnterpriseCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceEcEnterpriseCreateDefaultResponse) Get() *AlipayCommerceEcEnterpriseCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceEcEnterpriseCreateDefaultResponse) Set(val *AlipayCommerceEcEnterpriseCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcEnterpriseCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcEnterpriseCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcEnterpriseCreateDefaultResponse(val *AlipayCommerceEcEnterpriseCreateDefaultResponse) *NullableAlipayCommerceEcEnterpriseCreateDefaultResponse {
	return &NullableAlipayCommerceEcEnterpriseCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcEnterpriseCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcEnterpriseCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
