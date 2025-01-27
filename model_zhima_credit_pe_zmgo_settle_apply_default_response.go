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

// ZhimaCreditPeZmgoSettleApplyDefaultResponse struct for ZhimaCreditPeZmgoSettleApplyDefaultResponse
type ZhimaCreditPeZmgoSettleApplyDefaultResponse struct {
	CommonErrorType                                *CommonErrorType
	ZhimaCreditPeZmgoSettleApplyErrorResponseModel *ZhimaCreditPeZmgoSettleApplyErrorResponseModel
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ZhimaCreditPeZmgoSettleApplyDefaultResponse) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into ZhimaCreditPeZmgoSettleApplyErrorResponseModel
	err = json.Unmarshal(data, &dst.ZhimaCreditPeZmgoSettleApplyErrorResponseModel)
	if err == nil {
		jsonZhimaCreditPeZmgoSettleApplyErrorResponseModel, _ := json.Marshal(dst.ZhimaCreditPeZmgoSettleApplyErrorResponseModel)
		if string(jsonZhimaCreditPeZmgoSettleApplyErrorResponseModel) == "{}" { // empty struct
			dst.ZhimaCreditPeZmgoSettleApplyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.ZhimaCreditPeZmgoSettleApplyErrorResponseModel, return on the first match
		}
	} else {
		dst.ZhimaCreditPeZmgoSettleApplyErrorResponseModel = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ZhimaCreditPeZmgoSettleApplyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ZhimaCreditPeZmgoSettleApplyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	if src.ZhimaCreditPeZmgoSettleApplyErrorResponseModel != nil {
		return json.Marshal(&src.ZhimaCreditPeZmgoSettleApplyErrorResponseModel)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableZhimaCreditPeZmgoSettleApplyDefaultResponse struct {
	value *ZhimaCreditPeZmgoSettleApplyDefaultResponse
	isSet bool
}

func (v NullableZhimaCreditPeZmgoSettleApplyDefaultResponse) Get() *ZhimaCreditPeZmgoSettleApplyDefaultResponse {
	return v.value
}

func (v *NullableZhimaCreditPeZmgoSettleApplyDefaultResponse) Set(val *ZhimaCreditPeZmgoSettleApplyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaCreditPeZmgoSettleApplyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaCreditPeZmgoSettleApplyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaCreditPeZmgoSettleApplyDefaultResponse(val *ZhimaCreditPeZmgoSettleApplyDefaultResponse) *NullableZhimaCreditPeZmgoSettleApplyDefaultResponse {
	return &NullableZhimaCreditPeZmgoSettleApplyDefaultResponse{value: val, isSet: true}
}

func (v NullableZhimaCreditPeZmgoSettleApplyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaCreditPeZmgoSettleApplyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
