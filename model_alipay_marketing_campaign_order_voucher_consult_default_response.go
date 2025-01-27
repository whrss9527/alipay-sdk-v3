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

// AlipayMarketingCampaignOrderVoucherConsultDefaultResponse struct for AlipayMarketingCampaignOrderVoucherConsultDefaultResponse
type AlipayMarketingCampaignOrderVoucherConsultDefaultResponse struct {
	AlipayMarketingCampaignOrderVoucherConsultErrorResponseModel *AlipayMarketingCampaignOrderVoucherConsultErrorResponseModel
	CommonErrorType                                              *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingCampaignOrderVoucherConsultDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingCampaignOrderVoucherConsultErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingCampaignOrderVoucherConsultErrorResponseModel)
	if err == nil {
		jsonAlipayMarketingCampaignOrderVoucherConsultErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingCampaignOrderVoucherConsultErrorResponseModel)
		if string(jsonAlipayMarketingCampaignOrderVoucherConsultErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingCampaignOrderVoucherConsultErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingCampaignOrderVoucherConsultErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingCampaignOrderVoucherConsultErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingCampaignOrderVoucherConsultDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingCampaignOrderVoucherConsultDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingCampaignOrderVoucherConsultErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingCampaignOrderVoucherConsultErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayMarketingCampaignOrderVoucherConsultDefaultResponse struct {
	value *AlipayMarketingCampaignOrderVoucherConsultDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingCampaignOrderVoucherConsultDefaultResponse) Get() *AlipayMarketingCampaignOrderVoucherConsultDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingCampaignOrderVoucherConsultDefaultResponse) Set(val *AlipayMarketingCampaignOrderVoucherConsultDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCampaignOrderVoucherConsultDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCampaignOrderVoucherConsultDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCampaignOrderVoucherConsultDefaultResponse(val *AlipayMarketingCampaignOrderVoucherConsultDefaultResponse) *NullableAlipayMarketingCampaignOrderVoucherConsultDefaultResponse {
	return &NullableAlipayMarketingCampaignOrderVoucherConsultDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingCampaignOrderVoucherConsultDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCampaignOrderVoucherConsultDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
