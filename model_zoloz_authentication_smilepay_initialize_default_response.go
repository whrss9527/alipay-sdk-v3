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

// ZolozAuthenticationSmilepayInitializeDefaultResponse struct for ZolozAuthenticationSmilepayInitializeDefaultResponse
type ZolozAuthenticationSmilepayInitializeDefaultResponse struct {
	CommonErrorType                                         *CommonErrorType
	ZolozAuthenticationSmilepayInitializeErrorResponseModel *ZolozAuthenticationSmilepayInitializeErrorResponseModel
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ZolozAuthenticationSmilepayInitializeDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
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

	// try to unmarshal JSON data into ZolozAuthenticationSmilepayInitializeErrorResponseModel
	err = json.Unmarshal(data, &dst.ZolozAuthenticationSmilepayInitializeErrorResponseModel)
	if err == nil {
		jsonZolozAuthenticationSmilepayInitializeErrorResponseModel, _ := json.Marshal(dst.ZolozAuthenticationSmilepayInitializeErrorResponseModel)
		if string(jsonZolozAuthenticationSmilepayInitializeErrorResponseModel) == "{}" { // empty struct
			dst.ZolozAuthenticationSmilepayInitializeErrorResponseModel = nil
		} else {
			return nil // data stored in dst.ZolozAuthenticationSmilepayInitializeErrorResponseModel, return on the first match
		}
	} else {
		dst.ZolozAuthenticationSmilepayInitializeErrorResponseModel = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ZolozAuthenticationSmilepayInitializeDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ZolozAuthenticationSmilepayInitializeDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	if src.ZolozAuthenticationSmilepayInitializeErrorResponseModel != nil {
		return json.Marshal(&src.ZolozAuthenticationSmilepayInitializeErrorResponseModel)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableZolozAuthenticationSmilepayInitializeDefaultResponse struct {
	value *ZolozAuthenticationSmilepayInitializeDefaultResponse
	isSet bool
}

func (v NullableZolozAuthenticationSmilepayInitializeDefaultResponse) Get() *ZolozAuthenticationSmilepayInitializeDefaultResponse {
	return v.value
}

func (v *NullableZolozAuthenticationSmilepayInitializeDefaultResponse) Set(val *ZolozAuthenticationSmilepayInitializeDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableZolozAuthenticationSmilepayInitializeDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableZolozAuthenticationSmilepayInitializeDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZolozAuthenticationSmilepayInitializeDefaultResponse(val *ZolozAuthenticationSmilepayInitializeDefaultResponse) *NullableZolozAuthenticationSmilepayInitializeDefaultResponse {
	return &NullableZolozAuthenticationSmilepayInitializeDefaultResponse{value: val, isSet: true}
}

func (v NullableZolozAuthenticationSmilepayInitializeDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZolozAuthenticationSmilepayInitializeDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
