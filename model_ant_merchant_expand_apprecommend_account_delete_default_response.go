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

// AntMerchantExpandApprecommendAccountDeleteDefaultResponse struct for AntMerchantExpandApprecommendAccountDeleteDefaultResponse
type AntMerchantExpandApprecommendAccountDeleteDefaultResponse struct {
	AntMerchantExpandApprecommendAccountDeleteErrorResponseModel *AntMerchantExpandApprecommendAccountDeleteErrorResponseModel
	CommonErrorType                                              *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AntMerchantExpandApprecommendAccountDeleteDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AntMerchantExpandApprecommendAccountDeleteErrorResponseModel
	err = json.Unmarshal(data, &dst.AntMerchantExpandApprecommendAccountDeleteErrorResponseModel)
	if err == nil {
		jsonAntMerchantExpandApprecommendAccountDeleteErrorResponseModel, _ := json.Marshal(dst.AntMerchantExpandApprecommendAccountDeleteErrorResponseModel)
		if string(jsonAntMerchantExpandApprecommendAccountDeleteErrorResponseModel) == "{}" { // empty struct
			dst.AntMerchantExpandApprecommendAccountDeleteErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AntMerchantExpandApprecommendAccountDeleteErrorResponseModel, return on the first match
		}
	} else {
		dst.AntMerchantExpandApprecommendAccountDeleteErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AntMerchantExpandApprecommendAccountDeleteDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AntMerchantExpandApprecommendAccountDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AntMerchantExpandApprecommendAccountDeleteErrorResponseModel != nil {
		return json.Marshal(&src.AntMerchantExpandApprecommendAccountDeleteErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAntMerchantExpandApprecommendAccountDeleteDefaultResponse struct {
	value *AntMerchantExpandApprecommendAccountDeleteDefaultResponse
	isSet bool
}

func (v NullableAntMerchantExpandApprecommendAccountDeleteDefaultResponse) Get() *AntMerchantExpandApprecommendAccountDeleteDefaultResponse {
	return v.value
}

func (v *NullableAntMerchantExpandApprecommendAccountDeleteDefaultResponse) Set(val *AntMerchantExpandApprecommendAccountDeleteDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandApprecommendAccountDeleteDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandApprecommendAccountDeleteDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandApprecommendAccountDeleteDefaultResponse(val *AntMerchantExpandApprecommendAccountDeleteDefaultResponse) *NullableAntMerchantExpandApprecommendAccountDeleteDefaultResponse {
	return &NullableAntMerchantExpandApprecommendAccountDeleteDefaultResponse{value: val, isSet: true}
}

func (v NullableAntMerchantExpandApprecommendAccountDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandApprecommendAccountDeleteDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
