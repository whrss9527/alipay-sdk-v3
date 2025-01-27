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

// AlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse struct for AlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse
type AlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse struct {
	AlipayEbppInvoiceExpensecontrolAggregationCreateErrorResponseModel *AlipayEbppInvoiceExpensecontrolAggregationCreateErrorResponseModel
	CommonErrorType                                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEbppInvoiceExpensecontrolAggregationCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEbppInvoiceExpensecontrolAggregationCreateErrorResponseModel)
	if err == nil {
		jsonAlipayEbppInvoiceExpensecontrolAggregationCreateErrorResponseModel, _ := json.Marshal(dst.AlipayEbppInvoiceExpensecontrolAggregationCreateErrorResponseModel)
		if string(jsonAlipayEbppInvoiceExpensecontrolAggregationCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEbppInvoiceExpensecontrolAggregationCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEbppInvoiceExpensecontrolAggregationCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEbppInvoiceExpensecontrolAggregationCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEbppInvoiceExpensecontrolAggregationCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEbppInvoiceExpensecontrolAggregationCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse struct {
	value *AlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse) Get() *AlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse) Set(val *AlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse(val *AlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse) *NullableAlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse {
	return &NullableAlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpensecontrolAggregationCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
