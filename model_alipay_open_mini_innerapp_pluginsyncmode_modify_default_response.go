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

// AlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse struct for AlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse
type AlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse struct {
	AlipayOpenMiniInnerappPluginsyncmodeModifyErrorResponseModel *AlipayOpenMiniInnerappPluginsyncmodeModifyErrorResponseModel
	CommonErrorType                                              *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniInnerappPluginsyncmodeModifyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniInnerappPluginsyncmodeModifyErrorResponseModel)
	if err == nil {
		jsonAlipayOpenMiniInnerappPluginsyncmodeModifyErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniInnerappPluginsyncmodeModifyErrorResponseModel)
		if string(jsonAlipayOpenMiniInnerappPluginsyncmodeModifyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniInnerappPluginsyncmodeModifyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniInnerappPluginsyncmodeModifyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniInnerappPluginsyncmodeModifyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniInnerappPluginsyncmodeModifyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniInnerappPluginsyncmodeModifyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse struct {
	value *AlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse) Get() *AlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse) Set(val *AlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse(val *AlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse) *NullableAlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse {
	return &NullableAlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerappPluginsyncmodeModifyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
