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


// AlipayOpenPublicLifeMsgRecallDefaultResponse struct for AlipayOpenPublicLifeMsgRecallDefaultResponse
type AlipayOpenPublicLifeMsgRecallDefaultResponse struct {
	AlipayOpenPublicLifeMsgRecallErrorResponseModel *AlipayOpenPublicLifeMsgRecallErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicLifeMsgRecallDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicLifeMsgRecallErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicLifeMsgRecallErrorResponseModel);
	if err == nil {
		jsonAlipayOpenPublicLifeMsgRecallErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicLifeMsgRecallErrorResponseModel)
		if string(jsonAlipayOpenPublicLifeMsgRecallErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicLifeMsgRecallErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicLifeMsgRecallErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicLifeMsgRecallErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicLifeMsgRecallDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicLifeMsgRecallDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicLifeMsgRecallErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicLifeMsgRecallErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenPublicLifeMsgRecallDefaultResponse struct {
	value *AlipayOpenPublicLifeMsgRecallDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicLifeMsgRecallDefaultResponse) Get() *AlipayOpenPublicLifeMsgRecallDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicLifeMsgRecallDefaultResponse) Set(val *AlipayOpenPublicLifeMsgRecallDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicLifeMsgRecallDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicLifeMsgRecallDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicLifeMsgRecallDefaultResponse(val *AlipayOpenPublicLifeMsgRecallDefaultResponse) *NullableAlipayOpenPublicLifeMsgRecallDefaultResponse {
	return &NullableAlipayOpenPublicLifeMsgRecallDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicLifeMsgRecallDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicLifeMsgRecallDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

