/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"encoding/json"
)

// checks if the ConsultActivityResultInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsultActivityResultInfo{}

// ConsultActivityResultInfo struct for ConsultActivityResultInfo
type ConsultActivityResultInfo struct {
	// 活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// 咨询结果码
	ConsultResultCode *string `json:"consult_result_code,omitempty"`
}

// NewConsultActivityResultInfo instantiates a new ConsultActivityResultInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsultActivityResultInfo() *ConsultActivityResultInfo {
	this := ConsultActivityResultInfo{}
	return &this
}

// NewConsultActivityResultInfoWithDefaults instantiates a new ConsultActivityResultInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsultActivityResultInfoWithDefaults() *ConsultActivityResultInfo {
	this := ConsultActivityResultInfo{}
	return &this
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *ConsultActivityResultInfo) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId) {
		var ret string
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsultActivityResultInfo) GetActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *ConsultActivityResultInfo) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given string and assigns it to the ActivityId field.
func (o *ConsultActivityResultInfo) SetActivityId(v string) {
	o.ActivityId = &v
}

// GetConsultResultCode returns the ConsultResultCode field value if set, zero value otherwise.
func (o *ConsultActivityResultInfo) GetConsultResultCode() string {
	if o == nil || IsNil(o.ConsultResultCode) {
		var ret string
		return ret
	}
	return *o.ConsultResultCode
}

// GetConsultResultCodeOk returns a tuple with the ConsultResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsultActivityResultInfo) GetConsultResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ConsultResultCode) {
		return nil, false
	}
	return o.ConsultResultCode, true
}

// HasConsultResultCode returns a boolean if a field has been set.
func (o *ConsultActivityResultInfo) HasConsultResultCode() bool {
	if o != nil && !IsNil(o.ConsultResultCode) {
		return true
	}

	return false
}

// SetConsultResultCode gets a reference to the given string and assigns it to the ConsultResultCode field.
func (o *ConsultActivityResultInfo) SetConsultResultCode(v string) {
	o.ConsultResultCode = &v
}

func (o ConsultActivityResultInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsultActivityResultInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
	}
	if !IsNil(o.ConsultResultCode) {
		toSerialize["consult_result_code"] = o.ConsultResultCode
	}
	return toSerialize, nil
}

type NullableConsultActivityResultInfo struct {
	value *ConsultActivityResultInfo
	isSet bool
}

func (v NullableConsultActivityResultInfo) Get() *ConsultActivityResultInfo {
	return v.value
}

func (v *NullableConsultActivityResultInfo) Set(val *ConsultActivityResultInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableConsultActivityResultInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableConsultActivityResultInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsultActivityResultInfo(val *ConsultActivityResultInfo) *NullableConsultActivityResultInfo {
	return &NullableConsultActivityResultInfo{value: val, isSet: true}
}

func (v NullableConsultActivityResultInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsultActivityResultInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
