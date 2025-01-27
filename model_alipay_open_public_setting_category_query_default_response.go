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

// AlipayOpenPublicSettingCategoryQueryDefaultResponse struct for AlipayOpenPublicSettingCategoryQueryDefaultResponse
type AlipayOpenPublicSettingCategoryQueryDefaultResponse struct {
	AlipayOpenPublicSettingCategoryQueryErrorResponseModel *AlipayOpenPublicSettingCategoryQueryErrorResponseModel
	CommonErrorType                                        *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicSettingCategoryQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicSettingCategoryQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicSettingCategoryQueryErrorResponseModel)
	if err == nil {
		jsonAlipayOpenPublicSettingCategoryQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicSettingCategoryQueryErrorResponseModel)
		if string(jsonAlipayOpenPublicSettingCategoryQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicSettingCategoryQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicSettingCategoryQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicSettingCategoryQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicSettingCategoryQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicSettingCategoryQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicSettingCategoryQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicSettingCategoryQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenPublicSettingCategoryQueryDefaultResponse struct {
	value *AlipayOpenPublicSettingCategoryQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicSettingCategoryQueryDefaultResponse) Get() *AlipayOpenPublicSettingCategoryQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicSettingCategoryQueryDefaultResponse) Set(val *AlipayOpenPublicSettingCategoryQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicSettingCategoryQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicSettingCategoryQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicSettingCategoryQueryDefaultResponse(val *AlipayOpenPublicSettingCategoryQueryDefaultResponse) *NullableAlipayOpenPublicSettingCategoryQueryDefaultResponse {
	return &NullableAlipayOpenPublicSettingCategoryQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicSettingCategoryQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicSettingCategoryQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
