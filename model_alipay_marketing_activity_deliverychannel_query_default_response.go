/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)


// AlipayMarketingActivityDeliverychannelQueryDefaultResponse struct for AlipayMarketingActivityDeliverychannelQueryDefaultResponse
type AlipayMarketingActivityDeliverychannelQueryDefaultResponse struct {
	AlipayMarketingActivityDeliverychannelQueryErrorResponseModel *AlipayMarketingActivityDeliverychannelQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingActivityDeliverychannelQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingActivityDeliverychannelQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingActivityDeliverychannelQueryErrorResponseModel);
	if err == nil {
		jsonAlipayMarketingActivityDeliverychannelQueryErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingActivityDeliverychannelQueryErrorResponseModel)
		if string(jsonAlipayMarketingActivityDeliverychannelQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingActivityDeliverychannelQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingActivityDeliverychannelQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingActivityDeliverychannelQueryErrorResponseModel = nil
	}

	// try to unmarshal JSON data into CommonErrorType
	err = json.Unmarshal(data, &dst.CommonErrorType);
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingActivityDeliverychannelQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingActivityDeliverychannelQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingActivityDeliverychannelQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingActivityDeliverychannelQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMarketingActivityDeliverychannelQueryDefaultResponse struct {
	value *AlipayMarketingActivityDeliverychannelQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingActivityDeliverychannelQueryDefaultResponse) Get() *AlipayMarketingActivityDeliverychannelQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingActivityDeliverychannelQueryDefaultResponse) Set(val *AlipayMarketingActivityDeliverychannelQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityDeliverychannelQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityDeliverychannelQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityDeliverychannelQueryDefaultResponse(val *AlipayMarketingActivityDeliverychannelQueryDefaultResponse) *NullableAlipayMarketingActivityDeliverychannelQueryDefaultResponse {
	return &NullableAlipayMarketingActivityDeliverychannelQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityDeliverychannelQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityDeliverychannelQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

