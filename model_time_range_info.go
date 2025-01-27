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

// checks if the TimeRangeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeRangeInfo{}

// TimeRangeInfo struct for TimeRangeInfo
type TimeRangeInfo struct {
	// 开始时间  格式：HH:mm:ss
	BeginTime   *string      `json:"begin_time,omitempty"`
	EndTimeInfo *EndTimeInfo `json:"end_time_info,omitempty"`
}

// NewTimeRangeInfo instantiates a new TimeRangeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeRangeInfo() *TimeRangeInfo {
	this := TimeRangeInfo{}
	return &this
}

// NewTimeRangeInfoWithDefaults instantiates a new TimeRangeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeRangeInfoWithDefaults() *TimeRangeInfo {
	this := TimeRangeInfo{}
	return &this
}

// GetBeginTime returns the BeginTime field value if set, zero value otherwise.
func (o *TimeRangeInfo) GetBeginTime() string {
	if o == nil || IsNil(o.BeginTime) {
		var ret string
		return ret
	}
	return *o.BeginTime
}

// GetBeginTimeOk returns a tuple with the BeginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeRangeInfo) GetBeginTimeOk() (*string, bool) {
	if o == nil || IsNil(o.BeginTime) {
		return nil, false
	}
	return o.BeginTime, true
}

// HasBeginTime returns a boolean if a field has been set.
func (o *TimeRangeInfo) HasBeginTime() bool {
	if o != nil && !IsNil(o.BeginTime) {
		return true
	}

	return false
}

// SetBeginTime gets a reference to the given string and assigns it to the BeginTime field.
func (o *TimeRangeInfo) SetBeginTime(v string) {
	o.BeginTime = &v
}

// GetEndTimeInfo returns the EndTimeInfo field value if set, zero value otherwise.
func (o *TimeRangeInfo) GetEndTimeInfo() EndTimeInfo {
	if o == nil || IsNil(o.EndTimeInfo) {
		var ret EndTimeInfo
		return ret
	}
	return *o.EndTimeInfo
}

// GetEndTimeInfoOk returns a tuple with the EndTimeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeRangeInfo) GetEndTimeInfoOk() (*EndTimeInfo, bool) {
	if o == nil || IsNil(o.EndTimeInfo) {
		return nil, false
	}
	return o.EndTimeInfo, true
}

// HasEndTimeInfo returns a boolean if a field has been set.
func (o *TimeRangeInfo) HasEndTimeInfo() bool {
	if o != nil && !IsNil(o.EndTimeInfo) {
		return true
	}

	return false
}

// SetEndTimeInfo gets a reference to the given EndTimeInfo and assigns it to the EndTimeInfo field.
func (o *TimeRangeInfo) SetEndTimeInfo(v EndTimeInfo) {
	o.EndTimeInfo = &v
}

func (o TimeRangeInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeRangeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BeginTime) {
		toSerialize["begin_time"] = o.BeginTime
	}
	if !IsNil(o.EndTimeInfo) {
		toSerialize["end_time_info"] = o.EndTimeInfo
	}
	return toSerialize, nil
}

type NullableTimeRangeInfo struct {
	value *TimeRangeInfo
	isSet bool
}

func (v NullableTimeRangeInfo) Get() *TimeRangeInfo {
	return v.value
}

func (v *NullableTimeRangeInfo) Set(val *TimeRangeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeRangeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeRangeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeRangeInfo(val *TimeRangeInfo) *NullableTimeRangeInfo {
	return &NullableTimeRangeInfo{value: val, isSet: true}
}

func (v NullableTimeRangeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeRangeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
