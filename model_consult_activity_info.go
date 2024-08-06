/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ConsultActivityInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsultActivityInfo{}

// ConsultActivityInfo struct for ConsultActivityInfo
type ConsultActivityInfo struct {
	// 活动id
	ActivityId *string `json:"activity_id,omitempty"`
}

// NewConsultActivityInfo instantiates a new ConsultActivityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsultActivityInfo() *ConsultActivityInfo {
	this := ConsultActivityInfo{}
	return &this
}

// NewConsultActivityInfoWithDefaults instantiates a new ConsultActivityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsultActivityInfoWithDefaults() *ConsultActivityInfo {
	this := ConsultActivityInfo{}
	return &this
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *ConsultActivityInfo) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId) {
		var ret string
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsultActivityInfo) GetActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *ConsultActivityInfo) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given string and assigns it to the ActivityId field.
func (o *ConsultActivityInfo) SetActivityId(v string) {
	o.ActivityId = &v
}

func (o ConsultActivityInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsultActivityInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
	}
	return toSerialize, nil
}

type NullableConsultActivityInfo struct {
	value *ConsultActivityInfo
	isSet bool
}

func (v NullableConsultActivityInfo) Get() *ConsultActivityInfo {
	return v.value
}

func (v *NullableConsultActivityInfo) Set(val *ConsultActivityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableConsultActivityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableConsultActivityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsultActivityInfo(val *ConsultActivityInfo) *NullableConsultActivityInfo {
	return &NullableConsultActivityInfo{value: val, isSet: true}
}

func (v NullableConsultActivityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsultActivityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

