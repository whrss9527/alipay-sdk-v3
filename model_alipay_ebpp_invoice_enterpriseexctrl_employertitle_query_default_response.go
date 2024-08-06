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


// AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse struct for AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse
type AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse struct {
	AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryErrorResponseModel *AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryErrorResponseModel);
	if err == nil {
		jsonAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryErrorResponseModel, _ := json.Marshal(dst.AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryErrorResponseModel)
		if string(jsonAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse struct {
	value *AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse) Get() *AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse) Set(val *AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse(val *AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse) *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse {
	return &NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

