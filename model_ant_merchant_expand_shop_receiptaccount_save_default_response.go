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

// AntMerchantExpandShopReceiptaccountSaveDefaultResponse struct for AntMerchantExpandShopReceiptaccountSaveDefaultResponse
type AntMerchantExpandShopReceiptaccountSaveDefaultResponse struct {
	AntMerchantExpandShopReceiptaccountSaveErrorResponseModel *AntMerchantExpandShopReceiptaccountSaveErrorResponseModel
	CommonErrorType                                           *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AntMerchantExpandShopReceiptaccountSaveDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AntMerchantExpandShopReceiptaccountSaveErrorResponseModel
	err = json.Unmarshal(data, &dst.AntMerchantExpandShopReceiptaccountSaveErrorResponseModel)
	if err == nil {
		jsonAntMerchantExpandShopReceiptaccountSaveErrorResponseModel, _ := json.Marshal(dst.AntMerchantExpandShopReceiptaccountSaveErrorResponseModel)
		if string(jsonAntMerchantExpandShopReceiptaccountSaveErrorResponseModel) == "{}" { // empty struct
			dst.AntMerchantExpandShopReceiptaccountSaveErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AntMerchantExpandShopReceiptaccountSaveErrorResponseModel, return on the first match
		}
	} else {
		dst.AntMerchantExpandShopReceiptaccountSaveErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AntMerchantExpandShopReceiptaccountSaveDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AntMerchantExpandShopReceiptaccountSaveDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AntMerchantExpandShopReceiptaccountSaveErrorResponseModel != nil {
		return json.Marshal(&src.AntMerchantExpandShopReceiptaccountSaveErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAntMerchantExpandShopReceiptaccountSaveDefaultResponse struct {
	value *AntMerchantExpandShopReceiptaccountSaveDefaultResponse
	isSet bool
}

func (v NullableAntMerchantExpandShopReceiptaccountSaveDefaultResponse) Get() *AntMerchantExpandShopReceiptaccountSaveDefaultResponse {
	return v.value
}

func (v *NullableAntMerchantExpandShopReceiptaccountSaveDefaultResponse) Set(val *AntMerchantExpandShopReceiptaccountSaveDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandShopReceiptaccountSaveDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandShopReceiptaccountSaveDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandShopReceiptaccountSaveDefaultResponse(val *AntMerchantExpandShopReceiptaccountSaveDefaultResponse) *NullableAntMerchantExpandShopReceiptaccountSaveDefaultResponse {
	return &NullableAntMerchantExpandShopReceiptaccountSaveDefaultResponse{value: val, isSet: true}
}

func (v NullableAntMerchantExpandShopReceiptaccountSaveDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandShopReceiptaccountSaveDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
