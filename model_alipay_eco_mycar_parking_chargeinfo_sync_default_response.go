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

// AlipayEcoMycarParkingChargeinfoSyncDefaultResponse struct for AlipayEcoMycarParkingChargeinfoSyncDefaultResponse
type AlipayEcoMycarParkingChargeinfoSyncDefaultResponse struct {
	AlipayEcoMycarParkingChargeinfoSyncErrorResponseModel *AlipayEcoMycarParkingChargeinfoSyncErrorResponseModel
	CommonErrorType                                       *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEcoMycarParkingChargeinfoSyncDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEcoMycarParkingChargeinfoSyncErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEcoMycarParkingChargeinfoSyncErrorResponseModel)
	if err == nil {
		jsonAlipayEcoMycarParkingChargeinfoSyncErrorResponseModel, _ := json.Marshal(dst.AlipayEcoMycarParkingChargeinfoSyncErrorResponseModel)
		if string(jsonAlipayEcoMycarParkingChargeinfoSyncErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEcoMycarParkingChargeinfoSyncErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEcoMycarParkingChargeinfoSyncErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEcoMycarParkingChargeinfoSyncErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEcoMycarParkingChargeinfoSyncDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEcoMycarParkingChargeinfoSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEcoMycarParkingChargeinfoSyncErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEcoMycarParkingChargeinfoSyncErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayEcoMycarParkingChargeinfoSyncDefaultResponse struct {
	value *AlipayEcoMycarParkingChargeinfoSyncDefaultResponse
	isSet bool
}

func (v NullableAlipayEcoMycarParkingChargeinfoSyncDefaultResponse) Get() *AlipayEcoMycarParkingChargeinfoSyncDefaultResponse {
	return v.value
}

func (v *NullableAlipayEcoMycarParkingChargeinfoSyncDefaultResponse) Set(val *AlipayEcoMycarParkingChargeinfoSyncDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoMycarParkingChargeinfoSyncDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoMycarParkingChargeinfoSyncDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoMycarParkingChargeinfoSyncDefaultResponse(val *AlipayEcoMycarParkingChargeinfoSyncDefaultResponse) *NullableAlipayEcoMycarParkingChargeinfoSyncDefaultResponse {
	return &NullableAlipayEcoMycarParkingChargeinfoSyncDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEcoMycarParkingChargeinfoSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoMycarParkingChargeinfoSyncDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
