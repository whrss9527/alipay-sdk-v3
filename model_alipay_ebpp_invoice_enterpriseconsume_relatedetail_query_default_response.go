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

// AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse struct for AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse
type AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse struct {
	AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryErrorResponseModel *AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryErrorResponseModel
	CommonErrorType                                                       *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryErrorResponseModel)
	if err == nil {
		jsonAlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryErrorResponseModel, _ := json.Marshal(dst.AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryErrorResponseModel)
		if string(jsonAlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse struct {
	value *AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse) Get() *AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse) Set(val *AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse(val *AlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse) *NullableAlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse {
	return &NullableAlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeRelatedetailQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
