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

// AlipayEbppInvoiceIssueruleCreateDefaultResponse struct for AlipayEbppInvoiceIssueruleCreateDefaultResponse
type AlipayEbppInvoiceIssueruleCreateDefaultResponse struct {
	AlipayEbppInvoiceIssueruleCreateErrorResponseModel *AlipayEbppInvoiceIssueruleCreateErrorResponseModel
	CommonErrorType                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEbppInvoiceIssueruleCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEbppInvoiceIssueruleCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEbppInvoiceIssueruleCreateErrorResponseModel)
	if err == nil {
		jsonAlipayEbppInvoiceIssueruleCreateErrorResponseModel, _ := json.Marshal(dst.AlipayEbppInvoiceIssueruleCreateErrorResponseModel)
		if string(jsonAlipayEbppInvoiceIssueruleCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEbppInvoiceIssueruleCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEbppInvoiceIssueruleCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEbppInvoiceIssueruleCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEbppInvoiceIssueruleCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEbppInvoiceIssueruleCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEbppInvoiceIssueruleCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEbppInvoiceIssueruleCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayEbppInvoiceIssueruleCreateDefaultResponse struct {
	value *AlipayEbppInvoiceIssueruleCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayEbppInvoiceIssueruleCreateDefaultResponse) Get() *AlipayEbppInvoiceIssueruleCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayEbppInvoiceIssueruleCreateDefaultResponse) Set(val *AlipayEbppInvoiceIssueruleCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceIssueruleCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceIssueruleCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceIssueruleCreateDefaultResponse(val *AlipayEbppInvoiceIssueruleCreateDefaultResponse) *NullableAlipayEbppInvoiceIssueruleCreateDefaultResponse {
	return &NullableAlipayEbppInvoiceIssueruleCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceIssueruleCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceIssueruleCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
