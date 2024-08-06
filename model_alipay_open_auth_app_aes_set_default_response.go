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


// AlipayOpenAuthAppAesSetDefaultResponse struct for AlipayOpenAuthAppAesSetDefaultResponse
type AlipayOpenAuthAppAesSetDefaultResponse struct {
	AlipayOpenAuthAppAesSetErrorResponseModel *AlipayOpenAuthAppAesSetErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenAuthAppAesSetDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenAuthAppAesSetErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenAuthAppAesSetErrorResponseModel);
	if err == nil {
		jsonAlipayOpenAuthAppAesSetErrorResponseModel, _ := json.Marshal(dst.AlipayOpenAuthAppAesSetErrorResponseModel)
		if string(jsonAlipayOpenAuthAppAesSetErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenAuthAppAesSetErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenAuthAppAesSetErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenAuthAppAesSetErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenAuthAppAesSetDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenAuthAppAesSetDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenAuthAppAesSetErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenAuthAppAesSetErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenAuthAppAesSetDefaultResponse struct {
	value *AlipayOpenAuthAppAesSetDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenAuthAppAesSetDefaultResponse) Get() *AlipayOpenAuthAppAesSetDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenAuthAppAesSetDefaultResponse) Set(val *AlipayOpenAuthAppAesSetDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAuthAppAesSetDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAuthAppAesSetDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAuthAppAesSetDefaultResponse(val *AlipayOpenAuthAppAesSetDefaultResponse) *NullableAlipayOpenAuthAppAesSetDefaultResponse {
	return &NullableAlipayOpenAuthAppAesSetDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenAuthAppAesSetDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAuthAppAesSetDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

