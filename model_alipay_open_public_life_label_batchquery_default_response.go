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

// AlipayOpenPublicLifeLabelBatchqueryDefaultResponse struct for AlipayOpenPublicLifeLabelBatchqueryDefaultResponse
type AlipayOpenPublicLifeLabelBatchqueryDefaultResponse struct {
	AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel
	CommonErrorType                                       *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicLifeLabelBatchqueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel)
	if err == nil {
		jsonAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel)
		if string(jsonAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicLifeLabelBatchqueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicLifeLabelBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenPublicLifeLabelBatchqueryDefaultResponse struct {
	value *AlipayOpenPublicLifeLabelBatchqueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicLifeLabelBatchqueryDefaultResponse) Get() *AlipayOpenPublicLifeLabelBatchqueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicLifeLabelBatchqueryDefaultResponse) Set(val *AlipayOpenPublicLifeLabelBatchqueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicLifeLabelBatchqueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicLifeLabelBatchqueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicLifeLabelBatchqueryDefaultResponse(val *AlipayOpenPublicLifeLabelBatchqueryDefaultResponse) *NullableAlipayOpenPublicLifeLabelBatchqueryDefaultResponse {
	return &NullableAlipayOpenPublicLifeLabelBatchqueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicLifeLabelBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicLifeLabelBatchqueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
