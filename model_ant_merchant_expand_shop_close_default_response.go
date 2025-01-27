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

// AntMerchantExpandShopCloseDefaultResponse struct for AntMerchantExpandShopCloseDefaultResponse
type AntMerchantExpandShopCloseDefaultResponse struct {
	AntMerchantExpandShopCloseErrorResponseModel *AntMerchantExpandShopCloseErrorResponseModel
	CommonErrorType                              *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AntMerchantExpandShopCloseDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AntMerchantExpandShopCloseErrorResponseModel
	err = json.Unmarshal(data, &dst.AntMerchantExpandShopCloseErrorResponseModel)
	if err == nil {
		jsonAntMerchantExpandShopCloseErrorResponseModel, _ := json.Marshal(dst.AntMerchantExpandShopCloseErrorResponseModel)
		if string(jsonAntMerchantExpandShopCloseErrorResponseModel) == "{}" { // empty struct
			dst.AntMerchantExpandShopCloseErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AntMerchantExpandShopCloseErrorResponseModel, return on the first match
		}
	} else {
		dst.AntMerchantExpandShopCloseErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AntMerchantExpandShopCloseDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AntMerchantExpandShopCloseDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AntMerchantExpandShopCloseErrorResponseModel != nil {
		return json.Marshal(&src.AntMerchantExpandShopCloseErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAntMerchantExpandShopCloseDefaultResponse struct {
	value *AntMerchantExpandShopCloseDefaultResponse
	isSet bool
}

func (v NullableAntMerchantExpandShopCloseDefaultResponse) Get() *AntMerchantExpandShopCloseDefaultResponse {
	return v.value
}

func (v *NullableAntMerchantExpandShopCloseDefaultResponse) Set(val *AntMerchantExpandShopCloseDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandShopCloseDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandShopCloseDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandShopCloseDefaultResponse(val *AntMerchantExpandShopCloseDefaultResponse) *NullableAntMerchantExpandShopCloseDefaultResponse {
	return &NullableAntMerchantExpandShopCloseDefaultResponse{value: val, isSet: true}
}

func (v NullableAntMerchantExpandShopCloseDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandShopCloseDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
