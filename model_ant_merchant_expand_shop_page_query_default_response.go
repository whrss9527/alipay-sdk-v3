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

// AntMerchantExpandShopPageQueryDefaultResponse struct for AntMerchantExpandShopPageQueryDefaultResponse
type AntMerchantExpandShopPageQueryDefaultResponse struct {
	AntMerchantExpandShopPageQueryErrorResponseModel *AntMerchantExpandShopPageQueryErrorResponseModel
	CommonErrorType                                  *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AntMerchantExpandShopPageQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AntMerchantExpandShopPageQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AntMerchantExpandShopPageQueryErrorResponseModel)
	if err == nil {
		jsonAntMerchantExpandShopPageQueryErrorResponseModel, _ := json.Marshal(dst.AntMerchantExpandShopPageQueryErrorResponseModel)
		if string(jsonAntMerchantExpandShopPageQueryErrorResponseModel) == "{}" { // empty struct
			dst.AntMerchantExpandShopPageQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AntMerchantExpandShopPageQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AntMerchantExpandShopPageQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AntMerchantExpandShopPageQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AntMerchantExpandShopPageQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AntMerchantExpandShopPageQueryErrorResponseModel != nil {
		return json.Marshal(&src.AntMerchantExpandShopPageQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAntMerchantExpandShopPageQueryDefaultResponse struct {
	value *AntMerchantExpandShopPageQueryDefaultResponse
	isSet bool
}

func (v NullableAntMerchantExpandShopPageQueryDefaultResponse) Get() *AntMerchantExpandShopPageQueryDefaultResponse {
	return v.value
}

func (v *NullableAntMerchantExpandShopPageQueryDefaultResponse) Set(val *AntMerchantExpandShopPageQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandShopPageQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandShopPageQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandShopPageQueryDefaultResponse(val *AntMerchantExpandShopPageQueryDefaultResponse) *NullableAntMerchantExpandShopPageQueryDefaultResponse {
	return &NullableAntMerchantExpandShopPageQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAntMerchantExpandShopPageQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandShopPageQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
