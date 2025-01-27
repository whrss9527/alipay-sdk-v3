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

// AlipayOpenMiniIsvFastregisterCreateDefaultResponse struct for AlipayOpenMiniIsvFastregisterCreateDefaultResponse
type AlipayOpenMiniIsvFastregisterCreateDefaultResponse struct {
	AlipayOpenMiniIsvFastregisterCreateErrorResponseModel *AlipayOpenMiniIsvFastregisterCreateErrorResponseModel
	CommonErrorType                                       *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniIsvFastregisterCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniIsvFastregisterCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniIsvFastregisterCreateErrorResponseModel)
	if err == nil {
		jsonAlipayOpenMiniIsvFastregisterCreateErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniIsvFastregisterCreateErrorResponseModel)
		if string(jsonAlipayOpenMiniIsvFastregisterCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniIsvFastregisterCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniIsvFastregisterCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniIsvFastregisterCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniIsvFastregisterCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniIsvFastregisterCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniIsvFastregisterCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniIsvFastregisterCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenMiniIsvFastregisterCreateDefaultResponse struct {
	value *AlipayOpenMiniIsvFastregisterCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniIsvFastregisterCreateDefaultResponse) Get() *AlipayOpenMiniIsvFastregisterCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniIsvFastregisterCreateDefaultResponse) Set(val *AlipayOpenMiniIsvFastregisterCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniIsvFastregisterCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniIsvFastregisterCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniIsvFastregisterCreateDefaultResponse(val *AlipayOpenMiniIsvFastregisterCreateDefaultResponse) *NullableAlipayOpenMiniIsvFastregisterCreateDefaultResponse {
	return &NullableAlipayOpenMiniIsvFastregisterCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniIsvFastregisterCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniIsvFastregisterCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
