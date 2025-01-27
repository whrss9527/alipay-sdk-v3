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

// DatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse struct for DatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse
type DatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse struct {
	CommonErrorType                                                *CommonErrorType
	DatadigitalFincloudGeneralsaasFaceCheckQueryErrorResponseModel *DatadigitalFincloudGeneralsaasFaceCheckQueryErrorResponseModel
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into DatadigitalFincloudGeneralsaasFaceCheckQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.DatadigitalFincloudGeneralsaasFaceCheckQueryErrorResponseModel)
	if err == nil {
		jsonDatadigitalFincloudGeneralsaasFaceCheckQueryErrorResponseModel, _ := json.Marshal(dst.DatadigitalFincloudGeneralsaasFaceCheckQueryErrorResponseModel)
		if string(jsonDatadigitalFincloudGeneralsaasFaceCheckQueryErrorResponseModel) == "{}" { // empty struct
			dst.DatadigitalFincloudGeneralsaasFaceCheckQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.DatadigitalFincloudGeneralsaasFaceCheckQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.DatadigitalFincloudGeneralsaasFaceCheckQueryErrorResponseModel = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(DatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	if src.DatadigitalFincloudGeneralsaasFaceCheckQueryErrorResponseModel != nil {
		return json.Marshal(&src.DatadigitalFincloudGeneralsaasFaceCheckQueryErrorResponseModel)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse struct {
	value *DatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse
	isSet bool
}

func (v NullableDatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse) Get() *DatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse {
	return v.value
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse) Set(val *DatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse(val *DatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse) *NullableDatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse {
	return &NullableDatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableDatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatadigitalFincloudGeneralsaasFaceCheckQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
