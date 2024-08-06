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


// AlipayOpenPublicMenuCreateDefaultResponse struct for AlipayOpenPublicMenuCreateDefaultResponse
type AlipayOpenPublicMenuCreateDefaultResponse struct {
	AlipayOpenPublicMenuCreateErrorResponseModel *AlipayOpenPublicMenuCreateErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicMenuCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicMenuCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicMenuCreateErrorResponseModel);
	if err == nil {
		jsonAlipayOpenPublicMenuCreateErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicMenuCreateErrorResponseModel)
		if string(jsonAlipayOpenPublicMenuCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicMenuCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicMenuCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicMenuCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicMenuCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicMenuCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicMenuCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicMenuCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenPublicMenuCreateDefaultResponse struct {
	value *AlipayOpenPublicMenuCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicMenuCreateDefaultResponse) Get() *AlipayOpenPublicMenuCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicMenuCreateDefaultResponse) Set(val *AlipayOpenPublicMenuCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicMenuCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicMenuCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicMenuCreateDefaultResponse(val *AlipayOpenPublicMenuCreateDefaultResponse) *NullableAlipayOpenPublicMenuCreateDefaultResponse {
	return &NullableAlipayOpenPublicMenuCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicMenuCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicMenuCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

