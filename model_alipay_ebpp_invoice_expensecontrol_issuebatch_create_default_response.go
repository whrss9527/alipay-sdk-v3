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

// AlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse struct for AlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse
type AlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse struct {
	AlipayEbppInvoiceExpensecontrolIssuebatchCreateErrorResponseModel *AlipayEbppInvoiceExpensecontrolIssuebatchCreateErrorResponseModel
	CommonErrorType                                                   *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEbppInvoiceExpensecontrolIssuebatchCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEbppInvoiceExpensecontrolIssuebatchCreateErrorResponseModel)
	if err == nil {
		jsonAlipayEbppInvoiceExpensecontrolIssuebatchCreateErrorResponseModel, _ := json.Marshal(dst.AlipayEbppInvoiceExpensecontrolIssuebatchCreateErrorResponseModel)
		if string(jsonAlipayEbppInvoiceExpensecontrolIssuebatchCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEbppInvoiceExpensecontrolIssuebatchCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEbppInvoiceExpensecontrolIssuebatchCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEbppInvoiceExpensecontrolIssuebatchCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEbppInvoiceExpensecontrolIssuebatchCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEbppInvoiceExpensecontrolIssuebatchCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse struct {
	value *AlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse) Get() *AlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse) Set(val *AlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse(val *AlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse) *NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse {
	return &NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpensecontrolIssuebatchCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
