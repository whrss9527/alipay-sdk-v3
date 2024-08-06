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


// AlipayOpenAppApiFieldQueryDefaultResponse struct for AlipayOpenAppApiFieldQueryDefaultResponse
type AlipayOpenAppApiFieldQueryDefaultResponse struct {
	AlipayOpenAppApiFieldQueryErrorResponseModel *AlipayOpenAppApiFieldQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenAppApiFieldQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenAppApiFieldQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenAppApiFieldQueryErrorResponseModel);
	if err == nil {
		jsonAlipayOpenAppApiFieldQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenAppApiFieldQueryErrorResponseModel)
		if string(jsonAlipayOpenAppApiFieldQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenAppApiFieldQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenAppApiFieldQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenAppApiFieldQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenAppApiFieldQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenAppApiFieldQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenAppApiFieldQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenAppApiFieldQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenAppApiFieldQueryDefaultResponse struct {
	value *AlipayOpenAppApiFieldQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenAppApiFieldQueryDefaultResponse) Get() *AlipayOpenAppApiFieldQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenAppApiFieldQueryDefaultResponse) Set(val *AlipayOpenAppApiFieldQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAppApiFieldQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAppApiFieldQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAppApiFieldQueryDefaultResponse(val *AlipayOpenAppApiFieldQueryDefaultResponse) *NullableAlipayOpenAppApiFieldQueryDefaultResponse {
	return &NullableAlipayOpenAppApiFieldQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenAppApiFieldQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAppApiFieldQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

