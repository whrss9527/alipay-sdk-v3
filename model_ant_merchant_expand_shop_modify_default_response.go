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

// AntMerchantExpandShopModifyDefaultResponse struct for AntMerchantExpandShopModifyDefaultResponse
type AntMerchantExpandShopModifyDefaultResponse struct {
	AntMerchantExpandShopModifyErrorResponseModel *AntMerchantExpandShopModifyErrorResponseModel
	CommonErrorType                               *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AntMerchantExpandShopModifyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AntMerchantExpandShopModifyErrorResponseModel
	err = json.Unmarshal(data, &dst.AntMerchantExpandShopModifyErrorResponseModel)
	if err == nil {
		jsonAntMerchantExpandShopModifyErrorResponseModel, _ := json.Marshal(dst.AntMerchantExpandShopModifyErrorResponseModel)
		if string(jsonAntMerchantExpandShopModifyErrorResponseModel) == "{}" { // empty struct
			dst.AntMerchantExpandShopModifyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AntMerchantExpandShopModifyErrorResponseModel, return on the first match
		}
	} else {
		dst.AntMerchantExpandShopModifyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AntMerchantExpandShopModifyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AntMerchantExpandShopModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AntMerchantExpandShopModifyErrorResponseModel != nil {
		return json.Marshal(&src.AntMerchantExpandShopModifyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAntMerchantExpandShopModifyDefaultResponse struct {
	value *AntMerchantExpandShopModifyDefaultResponse
	isSet bool
}

func (v NullableAntMerchantExpandShopModifyDefaultResponse) Get() *AntMerchantExpandShopModifyDefaultResponse {
	return v.value
}

func (v *NullableAntMerchantExpandShopModifyDefaultResponse) Set(val *AntMerchantExpandShopModifyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandShopModifyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandShopModifyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandShopModifyDefaultResponse(val *AntMerchantExpandShopModifyDefaultResponse) *NullableAntMerchantExpandShopModifyDefaultResponse {
	return &NullableAntMerchantExpandShopModifyDefaultResponse{value: val, isSet: true}
}

func (v NullableAntMerchantExpandShopModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandShopModifyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
