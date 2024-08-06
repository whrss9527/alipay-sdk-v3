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


// AlipayFundJointaccountMemberQueryDefaultResponse struct for AlipayFundJointaccountMemberQueryDefaultResponse
type AlipayFundJointaccountMemberQueryDefaultResponse struct {
	AlipayFundJointaccountMemberQueryErrorResponseModel *AlipayFundJointaccountMemberQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayFundJointaccountMemberQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayFundJointaccountMemberQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayFundJointaccountMemberQueryErrorResponseModel);
	if err == nil {
		jsonAlipayFundJointaccountMemberQueryErrorResponseModel, _ := json.Marshal(dst.AlipayFundJointaccountMemberQueryErrorResponseModel)
		if string(jsonAlipayFundJointaccountMemberQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayFundJointaccountMemberQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayFundJointaccountMemberQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayFundJointaccountMemberQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayFundJointaccountMemberQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayFundJointaccountMemberQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayFundJointaccountMemberQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayFundJointaccountMemberQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayFundJointaccountMemberQueryDefaultResponse struct {
	value *AlipayFundJointaccountMemberQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayFundJointaccountMemberQueryDefaultResponse) Get() *AlipayFundJointaccountMemberQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayFundJointaccountMemberQueryDefaultResponse) Set(val *AlipayFundJointaccountMemberQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundJointaccountMemberQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundJointaccountMemberQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundJointaccountMemberQueryDefaultResponse(val *AlipayFundJointaccountMemberQueryDefaultResponse) *NullableAlipayFundJointaccountMemberQueryDefaultResponse {
	return &NullableAlipayFundJointaccountMemberQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayFundJointaccountMemberQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundJointaccountMemberQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

