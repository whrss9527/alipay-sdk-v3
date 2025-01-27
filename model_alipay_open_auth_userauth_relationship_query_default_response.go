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

// AlipayOpenAuthUserauthRelationshipQueryDefaultResponse struct for AlipayOpenAuthUserauthRelationshipQueryDefaultResponse
type AlipayOpenAuthUserauthRelationshipQueryDefaultResponse struct {
	AlipayOpenAuthUserauthRelationshipQueryErrorResponseModel *AlipayOpenAuthUserauthRelationshipQueryErrorResponseModel
	CommonErrorType                                           *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenAuthUserauthRelationshipQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenAuthUserauthRelationshipQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenAuthUserauthRelationshipQueryErrorResponseModel)
	if err == nil {
		jsonAlipayOpenAuthUserauthRelationshipQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenAuthUserauthRelationshipQueryErrorResponseModel)
		if string(jsonAlipayOpenAuthUserauthRelationshipQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenAuthUserauthRelationshipQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenAuthUserauthRelationshipQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenAuthUserauthRelationshipQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenAuthUserauthRelationshipQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenAuthUserauthRelationshipQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenAuthUserauthRelationshipQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenAuthUserauthRelationshipQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenAuthUserauthRelationshipQueryDefaultResponse struct {
	value *AlipayOpenAuthUserauthRelationshipQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenAuthUserauthRelationshipQueryDefaultResponse) Get() *AlipayOpenAuthUserauthRelationshipQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenAuthUserauthRelationshipQueryDefaultResponse) Set(val *AlipayOpenAuthUserauthRelationshipQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAuthUserauthRelationshipQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAuthUserauthRelationshipQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAuthUserauthRelationshipQueryDefaultResponse(val *AlipayOpenAuthUserauthRelationshipQueryDefaultResponse) *NullableAlipayOpenAuthUserauthRelationshipQueryDefaultResponse {
	return &NullableAlipayOpenAuthUserauthRelationshipQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenAuthUserauthRelationshipQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAuthUserauthRelationshipQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
