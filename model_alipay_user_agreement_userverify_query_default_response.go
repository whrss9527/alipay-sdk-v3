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

// AlipayUserAgreementUserverifyQueryDefaultResponse struct for AlipayUserAgreementUserverifyQueryDefaultResponse
type AlipayUserAgreementUserverifyQueryDefaultResponse struct {
	AlipayUserAgreementUserverifyQueryErrorResponseModel *AlipayUserAgreementUserverifyQueryErrorResponseModel
	CommonErrorType                                      *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayUserAgreementUserverifyQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayUserAgreementUserverifyQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayUserAgreementUserverifyQueryErrorResponseModel)
	if err == nil {
		jsonAlipayUserAgreementUserverifyQueryErrorResponseModel, _ := json.Marshal(dst.AlipayUserAgreementUserverifyQueryErrorResponseModel)
		if string(jsonAlipayUserAgreementUserverifyQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayUserAgreementUserverifyQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayUserAgreementUserverifyQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayUserAgreementUserverifyQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayUserAgreementUserverifyQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayUserAgreementUserverifyQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayUserAgreementUserverifyQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayUserAgreementUserverifyQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayUserAgreementUserverifyQueryDefaultResponse struct {
	value *AlipayUserAgreementUserverifyQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayUserAgreementUserverifyQueryDefaultResponse) Get() *AlipayUserAgreementUserverifyQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayUserAgreementUserverifyQueryDefaultResponse) Set(val *AlipayUserAgreementUserverifyQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayUserAgreementUserverifyQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayUserAgreementUserverifyQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayUserAgreementUserverifyQueryDefaultResponse(val *AlipayUserAgreementUserverifyQueryDefaultResponse) *NullableAlipayUserAgreementUserverifyQueryDefaultResponse {
	return &NullableAlipayUserAgreementUserverifyQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayUserAgreementUserverifyQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayUserAgreementUserverifyQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
