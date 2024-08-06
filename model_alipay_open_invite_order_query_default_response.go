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


// AlipayOpenInviteOrderQueryDefaultResponse struct for AlipayOpenInviteOrderQueryDefaultResponse
type AlipayOpenInviteOrderQueryDefaultResponse struct {
	AlipayOpenInviteOrderQueryErrorResponseModel *AlipayOpenInviteOrderQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenInviteOrderQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenInviteOrderQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenInviteOrderQueryErrorResponseModel);
	if err == nil {
		jsonAlipayOpenInviteOrderQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenInviteOrderQueryErrorResponseModel)
		if string(jsonAlipayOpenInviteOrderQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenInviteOrderQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenInviteOrderQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenInviteOrderQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenInviteOrderQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenInviteOrderQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenInviteOrderQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenInviteOrderQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenInviteOrderQueryDefaultResponse struct {
	value *AlipayOpenInviteOrderQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenInviteOrderQueryDefaultResponse) Get() *AlipayOpenInviteOrderQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenInviteOrderQueryDefaultResponse) Set(val *AlipayOpenInviteOrderQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenInviteOrderQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenInviteOrderQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenInviteOrderQueryDefaultResponse(val *AlipayOpenInviteOrderQueryDefaultResponse) *NullableAlipayOpenInviteOrderQueryDefaultResponse {
	return &NullableAlipayOpenInviteOrderQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenInviteOrderQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenInviteOrderQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

