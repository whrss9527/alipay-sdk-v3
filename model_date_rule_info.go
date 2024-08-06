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

// checks if the DateRuleInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateRuleInfo{}

// DateRuleInfo struct for DateRuleInfo
type DateRuleInfo struct {
	DateRangeInfo *DateRangeInfo `json:"date_range_info,omitempty"`
	TimeRangeInfo *TimeRangeInfo `json:"time_range_info,omitempty"`
}

// NewDateRuleInfo instantiates a new DateRuleInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateRuleInfo() *DateRuleInfo {
	this := DateRuleInfo{}
	return &this
}

// NewDateRuleInfoWithDefaults instantiates a new DateRuleInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateRuleInfoWithDefaults() *DateRuleInfo {
	this := DateRuleInfo{}
	return &this
}

// GetDateRangeInfo returns the DateRangeInfo field value if set, zero value otherwise.
func (o *DateRuleInfo) GetDateRangeInfo() DateRangeInfo {
	if o == nil || IsNil(o.DateRangeInfo) {
		var ret DateRangeInfo
		return ret
	}
	return *o.DateRangeInfo
}

// GetDateRangeInfoOk returns a tuple with the DateRangeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateRuleInfo) GetDateRangeInfoOk() (*DateRangeInfo, bool) {
	if o == nil || IsNil(o.DateRangeInfo) {
		return nil, false
	}
	return o.DateRangeInfo, true
}

// HasDateRangeInfo returns a boolean if a field has been set.
func (o *DateRuleInfo) HasDateRangeInfo() bool {
	if o != nil && !IsNil(o.DateRangeInfo) {
		return true
	}

	return false
}

// SetDateRangeInfo gets a reference to the given DateRangeInfo and assigns it to the DateRangeInfo field.
func (o *DateRuleInfo) SetDateRangeInfo(v DateRangeInfo) {
	o.DateRangeInfo = &v
}

// GetTimeRangeInfo returns the TimeRangeInfo field value if set, zero value otherwise.
func (o *DateRuleInfo) GetTimeRangeInfo() TimeRangeInfo {
	if o == nil || IsNil(o.TimeRangeInfo) {
		var ret TimeRangeInfo
		return ret
	}
	return *o.TimeRangeInfo
}

// GetTimeRangeInfoOk returns a tuple with the TimeRangeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateRuleInfo) GetTimeRangeInfoOk() (*TimeRangeInfo, bool) {
	if o == nil || IsNil(o.TimeRangeInfo) {
		return nil, false
	}
	return o.TimeRangeInfo, true
}

// HasTimeRangeInfo returns a boolean if a field has been set.
func (o *DateRuleInfo) HasTimeRangeInfo() bool {
	if o != nil && !IsNil(o.TimeRangeInfo) {
		return true
	}

	return false
}

// SetTimeRangeInfo gets a reference to the given TimeRangeInfo and assigns it to the TimeRangeInfo field.
func (o *DateRuleInfo) SetTimeRangeInfo(v TimeRangeInfo) {
	o.TimeRangeInfo = &v
}

func (o DateRuleInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DateRuleInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateRangeInfo) {
		toSerialize["date_range_info"] = o.DateRangeInfo
	}
	if !IsNil(o.TimeRangeInfo) {
		toSerialize["time_range_info"] = o.TimeRangeInfo
	}
	return toSerialize, nil
}

type NullableDateRuleInfo struct {
	value *DateRuleInfo
	isSet bool
}

func (v NullableDateRuleInfo) Get() *DateRuleInfo {
	return v.value
}

func (v *NullableDateRuleInfo) Set(val *DateRuleInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDateRuleInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDateRuleInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateRuleInfo(val *DateRuleInfo) *NullableDateRuleInfo {
	return &NullableDateRuleInfo{value: val, isSet: true}
}

func (v NullableDateRuleInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateRuleInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

