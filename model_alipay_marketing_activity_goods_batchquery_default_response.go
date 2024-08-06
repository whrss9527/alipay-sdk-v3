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


// AlipayMarketingActivityGoodsBatchqueryDefaultResponse struct for AlipayMarketingActivityGoodsBatchqueryDefaultResponse
type AlipayMarketingActivityGoodsBatchqueryDefaultResponse struct {
	AlipayMarketingActivityGoodsBatchqueryErrorResponseModel *AlipayMarketingActivityGoodsBatchqueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingActivityGoodsBatchqueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingActivityGoodsBatchqueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingActivityGoodsBatchqueryErrorResponseModel);
	if err == nil {
		jsonAlipayMarketingActivityGoodsBatchqueryErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingActivityGoodsBatchqueryErrorResponseModel)
		if string(jsonAlipayMarketingActivityGoodsBatchqueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingActivityGoodsBatchqueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingActivityGoodsBatchqueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingActivityGoodsBatchqueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingActivityGoodsBatchqueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingActivityGoodsBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingActivityGoodsBatchqueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingActivityGoodsBatchqueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMarketingActivityGoodsBatchqueryDefaultResponse struct {
	value *AlipayMarketingActivityGoodsBatchqueryDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingActivityGoodsBatchqueryDefaultResponse) Get() *AlipayMarketingActivityGoodsBatchqueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingActivityGoodsBatchqueryDefaultResponse) Set(val *AlipayMarketingActivityGoodsBatchqueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityGoodsBatchqueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityGoodsBatchqueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityGoodsBatchqueryDefaultResponse(val *AlipayMarketingActivityGoodsBatchqueryDefaultResponse) *NullableAlipayMarketingActivityGoodsBatchqueryDefaultResponse {
	return &NullableAlipayMarketingActivityGoodsBatchqueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityGoodsBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityGoodsBatchqueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

