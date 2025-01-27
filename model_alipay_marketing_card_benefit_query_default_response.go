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

// AlipayMarketingCardBenefitQueryDefaultResponse struct for AlipayMarketingCardBenefitQueryDefaultResponse
type AlipayMarketingCardBenefitQueryDefaultResponse struct {
	AlipayMarketingCardBenefitQueryErrorResponseModel *AlipayMarketingCardBenefitQueryErrorResponseModel
	CommonErrorType                                   *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingCardBenefitQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingCardBenefitQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingCardBenefitQueryErrorResponseModel)
	if err == nil {
		jsonAlipayMarketingCardBenefitQueryErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingCardBenefitQueryErrorResponseModel)
		if string(jsonAlipayMarketingCardBenefitQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingCardBenefitQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingCardBenefitQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingCardBenefitQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingCardBenefitQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingCardBenefitQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingCardBenefitQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingCardBenefitQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayMarketingCardBenefitQueryDefaultResponse struct {
	value *AlipayMarketingCardBenefitQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingCardBenefitQueryDefaultResponse) Get() *AlipayMarketingCardBenefitQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingCardBenefitQueryDefaultResponse) Set(val *AlipayMarketingCardBenefitQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardBenefitQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardBenefitQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardBenefitQueryDefaultResponse(val *AlipayMarketingCardBenefitQueryDefaultResponse) *NullableAlipayMarketingCardBenefitQueryDefaultResponse {
	return &NullableAlipayMarketingCardBenefitQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardBenefitQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardBenefitQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
