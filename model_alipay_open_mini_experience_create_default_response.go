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

// AlipayOpenMiniExperienceCreateDefaultResponse struct for AlipayOpenMiniExperienceCreateDefaultResponse
type AlipayOpenMiniExperienceCreateDefaultResponse struct {
	AlipayOpenMiniExperienceCreateErrorResponseModel *AlipayOpenMiniExperienceCreateErrorResponseModel
	CommonErrorType                                  *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniExperienceCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniExperienceCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniExperienceCreateErrorResponseModel)
	if err == nil {
		jsonAlipayOpenMiniExperienceCreateErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniExperienceCreateErrorResponseModel)
		if string(jsonAlipayOpenMiniExperienceCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniExperienceCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniExperienceCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniExperienceCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniExperienceCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniExperienceCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniExperienceCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniExperienceCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenMiniExperienceCreateDefaultResponse struct {
	value *AlipayOpenMiniExperienceCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniExperienceCreateDefaultResponse) Get() *AlipayOpenMiniExperienceCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniExperienceCreateDefaultResponse) Set(val *AlipayOpenMiniExperienceCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniExperienceCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniExperienceCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniExperienceCreateDefaultResponse(val *AlipayOpenMiniExperienceCreateDefaultResponse) *NullableAlipayOpenMiniExperienceCreateDefaultResponse {
	return &NullableAlipayOpenMiniExperienceCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniExperienceCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniExperienceCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
