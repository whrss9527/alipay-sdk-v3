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

// AlipayEbppInvoiceTitleListGetDefaultResponse struct for AlipayEbppInvoiceTitleListGetDefaultResponse
type AlipayEbppInvoiceTitleListGetDefaultResponse struct {
	AlipayEbppInvoiceTitleListGetErrorResponseModel *AlipayEbppInvoiceTitleListGetErrorResponseModel
	CommonErrorType                                 *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEbppInvoiceTitleListGetDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEbppInvoiceTitleListGetErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEbppInvoiceTitleListGetErrorResponseModel)
	if err == nil {
		jsonAlipayEbppInvoiceTitleListGetErrorResponseModel, _ := json.Marshal(dst.AlipayEbppInvoiceTitleListGetErrorResponseModel)
		if string(jsonAlipayEbppInvoiceTitleListGetErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEbppInvoiceTitleListGetErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEbppInvoiceTitleListGetErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEbppInvoiceTitleListGetErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEbppInvoiceTitleListGetDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEbppInvoiceTitleListGetDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEbppInvoiceTitleListGetErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEbppInvoiceTitleListGetErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayEbppInvoiceTitleListGetDefaultResponse struct {
	value *AlipayEbppInvoiceTitleListGetDefaultResponse
	isSet bool
}

func (v NullableAlipayEbppInvoiceTitleListGetDefaultResponse) Get() *AlipayEbppInvoiceTitleListGetDefaultResponse {
	return v.value
}

func (v *NullableAlipayEbppInvoiceTitleListGetDefaultResponse) Set(val *AlipayEbppInvoiceTitleListGetDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceTitleListGetDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceTitleListGetDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceTitleListGetDefaultResponse(val *AlipayEbppInvoiceTitleListGetDefaultResponse) *NullableAlipayEbppInvoiceTitleListGetDefaultResponse {
	return &NullableAlipayEbppInvoiceTitleListGetDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceTitleListGetDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceTitleListGetDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
