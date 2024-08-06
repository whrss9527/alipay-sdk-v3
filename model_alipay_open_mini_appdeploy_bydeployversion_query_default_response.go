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


// AlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse struct for AlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse
type AlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse struct {
	AlipayOpenMiniAppdeployBydeployversionQueryErrorResponseModel *AlipayOpenMiniAppdeployBydeployversionQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniAppdeployBydeployversionQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniAppdeployBydeployversionQueryErrorResponseModel);
	if err == nil {
		jsonAlipayOpenMiniAppdeployBydeployversionQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniAppdeployBydeployversionQueryErrorResponseModel)
		if string(jsonAlipayOpenMiniAppdeployBydeployversionQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniAppdeployBydeployversionQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniAppdeployBydeployversionQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniAppdeployBydeployversionQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniAppdeployBydeployversionQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniAppdeployBydeployversionQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse struct {
	value *AlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse) Get() *AlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse) Set(val *AlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse(val *AlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse) *NullableAlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse {
	return &NullableAlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniAppdeployBydeployversionQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

