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


// AlipayFundJointaccountDetailQueryDefaultResponse struct for AlipayFundJointaccountDetailQueryDefaultResponse
type AlipayFundJointaccountDetailQueryDefaultResponse struct {
	AlipayFundJointaccountDetailQueryErrorResponseModel *AlipayFundJointaccountDetailQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayFundJointaccountDetailQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayFundJointaccountDetailQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayFundJointaccountDetailQueryErrorResponseModel);
	if err == nil {
		jsonAlipayFundJointaccountDetailQueryErrorResponseModel, _ := json.Marshal(dst.AlipayFundJointaccountDetailQueryErrorResponseModel)
		if string(jsonAlipayFundJointaccountDetailQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayFundJointaccountDetailQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayFundJointaccountDetailQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayFundJointaccountDetailQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayFundJointaccountDetailQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayFundJointaccountDetailQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayFundJointaccountDetailQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayFundJointaccountDetailQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayFundJointaccountDetailQueryDefaultResponse struct {
	value *AlipayFundJointaccountDetailQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayFundJointaccountDetailQueryDefaultResponse) Get() *AlipayFundJointaccountDetailQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayFundJointaccountDetailQueryDefaultResponse) Set(val *AlipayFundJointaccountDetailQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundJointaccountDetailQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundJointaccountDetailQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundJointaccountDetailQueryDefaultResponse(val *AlipayFundJointaccountDetailQueryDefaultResponse) *NullableAlipayFundJointaccountDetailQueryDefaultResponse {
	return &NullableAlipayFundJointaccountDetailQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayFundJointaccountDetailQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundJointaccountDetailQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

