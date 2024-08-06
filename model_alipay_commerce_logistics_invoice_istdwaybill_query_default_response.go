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


// AlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse struct for AlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse
type AlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse struct {
	AlipayCommerceLogisticsInvoiceIstdwaybillQueryErrorResponseModel *AlipayCommerceLogisticsInvoiceIstdwaybillQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceLogisticsInvoiceIstdwaybillQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceLogisticsInvoiceIstdwaybillQueryErrorResponseModel);
	if err == nil {
		jsonAlipayCommerceLogisticsInvoiceIstdwaybillQueryErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceLogisticsInvoiceIstdwaybillQueryErrorResponseModel)
		if string(jsonAlipayCommerceLogisticsInvoiceIstdwaybillQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceLogisticsInvoiceIstdwaybillQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceLogisticsInvoiceIstdwaybillQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceLogisticsInvoiceIstdwaybillQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceLogisticsInvoiceIstdwaybillQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceLogisticsInvoiceIstdwaybillQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse struct {
	value *AlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse) Get() *AlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse) Set(val *AlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse(val *AlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse) *NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse {
	return &NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

