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

// AlipayCommerceOperationActivityMerchantModifyDefaultResponse struct for AlipayCommerceOperationActivityMerchantModifyDefaultResponse
type AlipayCommerceOperationActivityMerchantModifyDefaultResponse struct {
	AlipayCommerceOperationActivityMerchantModifyErrorResponseModel *AlipayCommerceOperationActivityMerchantModifyErrorResponseModel
	CommonErrorType                                                 *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceOperationActivityMerchantModifyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceOperationActivityMerchantModifyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceOperationActivityMerchantModifyErrorResponseModel)
	if err == nil {
		jsonAlipayCommerceOperationActivityMerchantModifyErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceOperationActivityMerchantModifyErrorResponseModel)
		if string(jsonAlipayCommerceOperationActivityMerchantModifyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceOperationActivityMerchantModifyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceOperationActivityMerchantModifyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceOperationActivityMerchantModifyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceOperationActivityMerchantModifyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceOperationActivityMerchantModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceOperationActivityMerchantModifyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceOperationActivityMerchantModifyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayCommerceOperationActivityMerchantModifyDefaultResponse struct {
	value *AlipayCommerceOperationActivityMerchantModifyDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceOperationActivityMerchantModifyDefaultResponse) Get() *AlipayCommerceOperationActivityMerchantModifyDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceOperationActivityMerchantModifyDefaultResponse) Set(val *AlipayCommerceOperationActivityMerchantModifyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceOperationActivityMerchantModifyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceOperationActivityMerchantModifyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceOperationActivityMerchantModifyDefaultResponse(val *AlipayCommerceOperationActivityMerchantModifyDefaultResponse) *NullableAlipayCommerceOperationActivityMerchantModifyDefaultResponse {
	return &NullableAlipayCommerceOperationActivityMerchantModifyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceOperationActivityMerchantModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceOperationActivityMerchantModifyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
