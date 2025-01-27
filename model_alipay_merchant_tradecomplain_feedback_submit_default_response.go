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

// AlipayMerchantTradecomplainFeedbackSubmitDefaultResponse struct for AlipayMerchantTradecomplainFeedbackSubmitDefaultResponse
type AlipayMerchantTradecomplainFeedbackSubmitDefaultResponse struct {
	AlipayMerchantTradecomplainFeedbackSubmitErrorResponseModel *AlipayMerchantTradecomplainFeedbackSubmitErrorResponseModel
	CommonErrorType                                             *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMerchantTradecomplainFeedbackSubmitDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMerchantTradecomplainFeedbackSubmitErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMerchantTradecomplainFeedbackSubmitErrorResponseModel)
	if err == nil {
		jsonAlipayMerchantTradecomplainFeedbackSubmitErrorResponseModel, _ := json.Marshal(dst.AlipayMerchantTradecomplainFeedbackSubmitErrorResponseModel)
		if string(jsonAlipayMerchantTradecomplainFeedbackSubmitErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMerchantTradecomplainFeedbackSubmitErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMerchantTradecomplainFeedbackSubmitErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMerchantTradecomplainFeedbackSubmitErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMerchantTradecomplainFeedbackSubmitDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMerchantTradecomplainFeedbackSubmitDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMerchantTradecomplainFeedbackSubmitErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMerchantTradecomplainFeedbackSubmitErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayMerchantTradecomplainFeedbackSubmitDefaultResponse struct {
	value *AlipayMerchantTradecomplainFeedbackSubmitDefaultResponse
	isSet bool
}

func (v NullableAlipayMerchantTradecomplainFeedbackSubmitDefaultResponse) Get() *AlipayMerchantTradecomplainFeedbackSubmitDefaultResponse {
	return v.value
}

func (v *NullableAlipayMerchantTradecomplainFeedbackSubmitDefaultResponse) Set(val *AlipayMerchantTradecomplainFeedbackSubmitDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMerchantTradecomplainFeedbackSubmitDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMerchantTradecomplainFeedbackSubmitDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMerchantTradecomplainFeedbackSubmitDefaultResponse(val *AlipayMerchantTradecomplainFeedbackSubmitDefaultResponse) *NullableAlipayMerchantTradecomplainFeedbackSubmitDefaultResponse {
	return &NullableAlipayMerchantTradecomplainFeedbackSubmitDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMerchantTradecomplainFeedbackSubmitDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMerchantTradecomplainFeedbackSubmitDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
