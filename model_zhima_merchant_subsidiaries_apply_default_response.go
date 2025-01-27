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

// ZhimaMerchantSubsidiariesApplyDefaultResponse struct for ZhimaMerchantSubsidiariesApplyDefaultResponse
type ZhimaMerchantSubsidiariesApplyDefaultResponse struct {
	CommonErrorType                                  *CommonErrorType
	ZhimaMerchantSubsidiariesApplyErrorResponseModel *ZhimaMerchantSubsidiariesApplyErrorResponseModel
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ZhimaMerchantSubsidiariesApplyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
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

	// try to unmarshal JSON data into ZhimaMerchantSubsidiariesApplyErrorResponseModel
	err = json.Unmarshal(data, &dst.ZhimaMerchantSubsidiariesApplyErrorResponseModel)
	if err == nil {
		jsonZhimaMerchantSubsidiariesApplyErrorResponseModel, _ := json.Marshal(dst.ZhimaMerchantSubsidiariesApplyErrorResponseModel)
		if string(jsonZhimaMerchantSubsidiariesApplyErrorResponseModel) == "{}" { // empty struct
			dst.ZhimaMerchantSubsidiariesApplyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.ZhimaMerchantSubsidiariesApplyErrorResponseModel, return on the first match
		}
	} else {
		dst.ZhimaMerchantSubsidiariesApplyErrorResponseModel = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ZhimaMerchantSubsidiariesApplyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ZhimaMerchantSubsidiariesApplyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	if src.ZhimaMerchantSubsidiariesApplyErrorResponseModel != nil {
		return json.Marshal(&src.ZhimaMerchantSubsidiariesApplyErrorResponseModel)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableZhimaMerchantSubsidiariesApplyDefaultResponse struct {
	value *ZhimaMerchantSubsidiariesApplyDefaultResponse
	isSet bool
}

func (v NullableZhimaMerchantSubsidiariesApplyDefaultResponse) Get() *ZhimaMerchantSubsidiariesApplyDefaultResponse {
	return v.value
}

func (v *NullableZhimaMerchantSubsidiariesApplyDefaultResponse) Set(val *ZhimaMerchantSubsidiariesApplyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaMerchantSubsidiariesApplyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaMerchantSubsidiariesApplyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaMerchantSubsidiariesApplyDefaultResponse(val *ZhimaMerchantSubsidiariesApplyDefaultResponse) *NullableZhimaMerchantSubsidiariesApplyDefaultResponse {
	return &NullableZhimaMerchantSubsidiariesApplyDefaultResponse{value: val, isSet: true}
}

func (v NullableZhimaMerchantSubsidiariesApplyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaMerchantSubsidiariesApplyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
