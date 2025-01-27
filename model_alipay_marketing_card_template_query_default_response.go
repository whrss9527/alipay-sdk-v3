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

// AlipayMarketingCardTemplateQueryDefaultResponse struct for AlipayMarketingCardTemplateQueryDefaultResponse
type AlipayMarketingCardTemplateQueryDefaultResponse struct {
	AlipayMarketingCardTemplateQueryErrorResponseModel *AlipayMarketingCardTemplateQueryErrorResponseModel
	CommonErrorType                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingCardTemplateQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingCardTemplateQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingCardTemplateQueryErrorResponseModel)
	if err == nil {
		jsonAlipayMarketingCardTemplateQueryErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingCardTemplateQueryErrorResponseModel)
		if string(jsonAlipayMarketingCardTemplateQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingCardTemplateQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingCardTemplateQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingCardTemplateQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingCardTemplateQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingCardTemplateQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingCardTemplateQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingCardTemplateQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayMarketingCardTemplateQueryDefaultResponse struct {
	value *AlipayMarketingCardTemplateQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingCardTemplateQueryDefaultResponse) Get() *AlipayMarketingCardTemplateQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingCardTemplateQueryDefaultResponse) Set(val *AlipayMarketingCardTemplateQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardTemplateQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardTemplateQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardTemplateQueryDefaultResponse(val *AlipayMarketingCardTemplateQueryDefaultResponse) *NullableAlipayMarketingCardTemplateQueryDefaultResponse {
	return &NullableAlipayMarketingCardTemplateQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardTemplateQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardTemplateQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
