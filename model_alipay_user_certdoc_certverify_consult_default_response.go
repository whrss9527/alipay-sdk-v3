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

// AlipayUserCertdocCertverifyConsultDefaultResponse struct for AlipayUserCertdocCertverifyConsultDefaultResponse
type AlipayUserCertdocCertverifyConsultDefaultResponse struct {
	AlipayUserCertdocCertverifyConsultErrorResponseModel *AlipayUserCertdocCertverifyConsultErrorResponseModel
	CommonErrorType                                      *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayUserCertdocCertverifyConsultDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayUserCertdocCertverifyConsultErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayUserCertdocCertverifyConsultErrorResponseModel)
	if err == nil {
		jsonAlipayUserCertdocCertverifyConsultErrorResponseModel, _ := json.Marshal(dst.AlipayUserCertdocCertverifyConsultErrorResponseModel)
		if string(jsonAlipayUserCertdocCertverifyConsultErrorResponseModel) == "{}" { // empty struct
			dst.AlipayUserCertdocCertverifyConsultErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayUserCertdocCertverifyConsultErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayUserCertdocCertverifyConsultErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayUserCertdocCertverifyConsultDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayUserCertdocCertverifyConsultDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayUserCertdocCertverifyConsultErrorResponseModel != nil {
		return json.Marshal(&src.AlipayUserCertdocCertverifyConsultErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayUserCertdocCertverifyConsultDefaultResponse struct {
	value *AlipayUserCertdocCertverifyConsultDefaultResponse
	isSet bool
}

func (v NullableAlipayUserCertdocCertverifyConsultDefaultResponse) Get() *AlipayUserCertdocCertverifyConsultDefaultResponse {
	return v.value
}

func (v *NullableAlipayUserCertdocCertverifyConsultDefaultResponse) Set(val *AlipayUserCertdocCertverifyConsultDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayUserCertdocCertverifyConsultDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayUserCertdocCertverifyConsultDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayUserCertdocCertverifyConsultDefaultResponse(val *AlipayUserCertdocCertverifyConsultDefaultResponse) *NullableAlipayUserCertdocCertverifyConsultDefaultResponse {
	return &NullableAlipayUserCertdocCertverifyConsultDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayUserCertdocCertverifyConsultDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayUserCertdocCertverifyConsultDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
