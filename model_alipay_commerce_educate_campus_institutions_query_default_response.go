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


// AlipayCommerceEducateCampusInstitutionsQueryDefaultResponse struct for AlipayCommerceEducateCampusInstitutionsQueryDefaultResponse
type AlipayCommerceEducateCampusInstitutionsQueryDefaultResponse struct {
	AlipayCommerceEducateCampusInstitutionsQueryErrorResponseModel *AlipayCommerceEducateCampusInstitutionsQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceEducateCampusInstitutionsQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceEducateCampusInstitutionsQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceEducateCampusInstitutionsQueryErrorResponseModel);
	if err == nil {
		jsonAlipayCommerceEducateCampusInstitutionsQueryErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceEducateCampusInstitutionsQueryErrorResponseModel)
		if string(jsonAlipayCommerceEducateCampusInstitutionsQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceEducateCampusInstitutionsQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceEducateCampusInstitutionsQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceEducateCampusInstitutionsQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceEducateCampusInstitutionsQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceEducateCampusInstitutionsQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceEducateCampusInstitutionsQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceEducateCampusInstitutionsQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayCommerceEducateCampusInstitutionsQueryDefaultResponse struct {
	value *AlipayCommerceEducateCampusInstitutionsQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceEducateCampusInstitutionsQueryDefaultResponse) Get() *AlipayCommerceEducateCampusInstitutionsQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceEducateCampusInstitutionsQueryDefaultResponse) Set(val *AlipayCommerceEducateCampusInstitutionsQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEducateCampusInstitutionsQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEducateCampusInstitutionsQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEducateCampusInstitutionsQueryDefaultResponse(val *AlipayCommerceEducateCampusInstitutionsQueryDefaultResponse) *NullableAlipayCommerceEducateCampusInstitutionsQueryDefaultResponse {
	return &NullableAlipayCommerceEducateCampusInstitutionsQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceEducateCampusInstitutionsQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEducateCampusInstitutionsQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

