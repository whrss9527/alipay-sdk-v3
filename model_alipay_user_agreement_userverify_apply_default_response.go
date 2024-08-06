/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)


// AlipayUserAgreementUserverifyApplyDefaultResponse struct for AlipayUserAgreementUserverifyApplyDefaultResponse
type AlipayUserAgreementUserverifyApplyDefaultResponse struct {
	AlipayUserAgreementUserverifyApplyErrorResponseModel *AlipayUserAgreementUserverifyApplyErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayUserAgreementUserverifyApplyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayUserAgreementUserverifyApplyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayUserAgreementUserverifyApplyErrorResponseModel);
	if err == nil {
		jsonAlipayUserAgreementUserverifyApplyErrorResponseModel, _ := json.Marshal(dst.AlipayUserAgreementUserverifyApplyErrorResponseModel)
		if string(jsonAlipayUserAgreementUserverifyApplyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayUserAgreementUserverifyApplyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayUserAgreementUserverifyApplyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayUserAgreementUserverifyApplyErrorResponseModel = nil
	}

	// try to unmarshal JSON data into CommonErrorType
	err = json.Unmarshal(data, &dst.CommonErrorType);
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayUserAgreementUserverifyApplyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayUserAgreementUserverifyApplyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayUserAgreementUserverifyApplyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayUserAgreementUserverifyApplyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayUserAgreementUserverifyApplyDefaultResponse struct {
	value *AlipayUserAgreementUserverifyApplyDefaultResponse
	isSet bool
}

func (v NullableAlipayUserAgreementUserverifyApplyDefaultResponse) Get() *AlipayUserAgreementUserverifyApplyDefaultResponse {
	return v.value
}

func (v *NullableAlipayUserAgreementUserverifyApplyDefaultResponse) Set(val *AlipayUserAgreementUserverifyApplyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayUserAgreementUserverifyApplyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayUserAgreementUserverifyApplyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayUserAgreementUserverifyApplyDefaultResponse(val *AlipayUserAgreementUserverifyApplyDefaultResponse) *NullableAlipayUserAgreementUserverifyApplyDefaultResponse {
	return &NullableAlipayUserAgreementUserverifyApplyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayUserAgreementUserverifyApplyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayUserAgreementUserverifyApplyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

