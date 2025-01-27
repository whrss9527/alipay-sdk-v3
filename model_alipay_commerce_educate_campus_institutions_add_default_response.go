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

// AlipayCommerceEducateCampusInstitutionsAddDefaultResponse struct for AlipayCommerceEducateCampusInstitutionsAddDefaultResponse
type AlipayCommerceEducateCampusInstitutionsAddDefaultResponse struct {
	AlipayCommerceEducateCampusInstitutionsAddErrorResponseModel *AlipayCommerceEducateCampusInstitutionsAddErrorResponseModel
	CommonErrorType                                              *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceEducateCampusInstitutionsAddDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceEducateCampusInstitutionsAddErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceEducateCampusInstitutionsAddErrorResponseModel)
	if err == nil {
		jsonAlipayCommerceEducateCampusInstitutionsAddErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceEducateCampusInstitutionsAddErrorResponseModel)
		if string(jsonAlipayCommerceEducateCampusInstitutionsAddErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceEducateCampusInstitutionsAddErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceEducateCampusInstitutionsAddErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceEducateCampusInstitutionsAddErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceEducateCampusInstitutionsAddDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceEducateCampusInstitutionsAddDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceEducateCampusInstitutionsAddErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceEducateCampusInstitutionsAddErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayCommerceEducateCampusInstitutionsAddDefaultResponse struct {
	value *AlipayCommerceEducateCampusInstitutionsAddDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceEducateCampusInstitutionsAddDefaultResponse) Get() *AlipayCommerceEducateCampusInstitutionsAddDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceEducateCampusInstitutionsAddDefaultResponse) Set(val *AlipayCommerceEducateCampusInstitutionsAddDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEducateCampusInstitutionsAddDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEducateCampusInstitutionsAddDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEducateCampusInstitutionsAddDefaultResponse(val *AlipayCommerceEducateCampusInstitutionsAddDefaultResponse) *NullableAlipayCommerceEducateCampusInstitutionsAddDefaultResponse {
	return &NullableAlipayCommerceEducateCampusInstitutionsAddDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceEducateCampusInstitutionsAddDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEducateCampusInstitutionsAddDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
