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

// AlipayEcoMycarParkingSpaceinfoSyncDefaultResponse struct for AlipayEcoMycarParkingSpaceinfoSyncDefaultResponse
type AlipayEcoMycarParkingSpaceinfoSyncDefaultResponse struct {
	AlipayEcoMycarParkingSpaceinfoSyncErrorResponseModel *AlipayEcoMycarParkingSpaceinfoSyncErrorResponseModel
	CommonErrorType                                      *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEcoMycarParkingSpaceinfoSyncDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEcoMycarParkingSpaceinfoSyncErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEcoMycarParkingSpaceinfoSyncErrorResponseModel)
	if err == nil {
		jsonAlipayEcoMycarParkingSpaceinfoSyncErrorResponseModel, _ := json.Marshal(dst.AlipayEcoMycarParkingSpaceinfoSyncErrorResponseModel)
		if string(jsonAlipayEcoMycarParkingSpaceinfoSyncErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEcoMycarParkingSpaceinfoSyncErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEcoMycarParkingSpaceinfoSyncErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEcoMycarParkingSpaceinfoSyncErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEcoMycarParkingSpaceinfoSyncDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEcoMycarParkingSpaceinfoSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEcoMycarParkingSpaceinfoSyncErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEcoMycarParkingSpaceinfoSyncErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayEcoMycarParkingSpaceinfoSyncDefaultResponse struct {
	value *AlipayEcoMycarParkingSpaceinfoSyncDefaultResponse
	isSet bool
}

func (v NullableAlipayEcoMycarParkingSpaceinfoSyncDefaultResponse) Get() *AlipayEcoMycarParkingSpaceinfoSyncDefaultResponse {
	return v.value
}

func (v *NullableAlipayEcoMycarParkingSpaceinfoSyncDefaultResponse) Set(val *AlipayEcoMycarParkingSpaceinfoSyncDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoMycarParkingSpaceinfoSyncDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoMycarParkingSpaceinfoSyncDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoMycarParkingSpaceinfoSyncDefaultResponse(val *AlipayEcoMycarParkingSpaceinfoSyncDefaultResponse) *NullableAlipayEcoMycarParkingSpaceinfoSyncDefaultResponse {
	return &NullableAlipayEcoMycarParkingSpaceinfoSyncDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEcoMycarParkingSpaceinfoSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoMycarParkingSpaceinfoSyncDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
