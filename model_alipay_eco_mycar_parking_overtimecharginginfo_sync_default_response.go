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

// AlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse struct for AlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse
type AlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse struct {
	AlipayEcoMycarParkingOvertimecharginginfoSyncErrorResponseModel *AlipayEcoMycarParkingOvertimecharginginfoSyncErrorResponseModel
	CommonErrorType                                                 *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEcoMycarParkingOvertimecharginginfoSyncErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEcoMycarParkingOvertimecharginginfoSyncErrorResponseModel)
	if err == nil {
		jsonAlipayEcoMycarParkingOvertimecharginginfoSyncErrorResponseModel, _ := json.Marshal(dst.AlipayEcoMycarParkingOvertimecharginginfoSyncErrorResponseModel)
		if string(jsonAlipayEcoMycarParkingOvertimecharginginfoSyncErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEcoMycarParkingOvertimecharginginfoSyncErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEcoMycarParkingOvertimecharginginfoSyncErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEcoMycarParkingOvertimecharginginfoSyncErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEcoMycarParkingOvertimecharginginfoSyncErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEcoMycarParkingOvertimecharginginfoSyncErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse struct {
	value *AlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse
	isSet bool
}

func (v NullableAlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse) Get() *AlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse {
	return v.value
}

func (v *NullableAlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse) Set(val *AlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse(val *AlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse) *NullableAlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse {
	return &NullableAlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoMycarParkingOvertimecharginginfoSyncDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
