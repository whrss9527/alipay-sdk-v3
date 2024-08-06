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


// AlipayEcoSignflowsUrlQueryDefaultResponse struct for AlipayEcoSignflowsUrlQueryDefaultResponse
type AlipayEcoSignflowsUrlQueryDefaultResponse struct {
	AlipayEcoSignflowsUrlQueryErrorResponseModel *AlipayEcoSignflowsUrlQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEcoSignflowsUrlQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEcoSignflowsUrlQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEcoSignflowsUrlQueryErrorResponseModel);
	if err == nil {
		jsonAlipayEcoSignflowsUrlQueryErrorResponseModel, _ := json.Marshal(dst.AlipayEcoSignflowsUrlQueryErrorResponseModel)
		if string(jsonAlipayEcoSignflowsUrlQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEcoSignflowsUrlQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEcoSignflowsUrlQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEcoSignflowsUrlQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEcoSignflowsUrlQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEcoSignflowsUrlQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEcoSignflowsUrlQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEcoSignflowsUrlQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEcoSignflowsUrlQueryDefaultResponse struct {
	value *AlipayEcoSignflowsUrlQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayEcoSignflowsUrlQueryDefaultResponse) Get() *AlipayEcoSignflowsUrlQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayEcoSignflowsUrlQueryDefaultResponse) Set(val *AlipayEcoSignflowsUrlQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoSignflowsUrlQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoSignflowsUrlQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoSignflowsUrlQueryDefaultResponse(val *AlipayEcoSignflowsUrlQueryDefaultResponse) *NullableAlipayEcoSignflowsUrlQueryDefaultResponse {
	return &NullableAlipayEcoSignflowsUrlQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEcoSignflowsUrlQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoSignflowsUrlQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

