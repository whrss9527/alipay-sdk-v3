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

// checks if the DateRangeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateRangeInfo{}

// DateRangeInfo struct for DateRangeInfo
type DateRangeInfo struct {
	// 开始日期  格式：yyyy-MM-dd
	BeginDate *string `json:"begin_date,omitempty"`
	// 结束日期  格式：yyyy-MM-dd
	EndDate *string `json:"end_date,omitempty"`
}

// NewDateRangeInfo instantiates a new DateRangeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateRangeInfo() *DateRangeInfo {
	this := DateRangeInfo{}
	return &this
}

// NewDateRangeInfoWithDefaults instantiates a new DateRangeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateRangeInfoWithDefaults() *DateRangeInfo {
	this := DateRangeInfo{}
	return &this
}

// GetBeginDate returns the BeginDate field value if set, zero value otherwise.
func (o *DateRangeInfo) GetBeginDate() string {
	if o == nil || IsNil(o.BeginDate) {
		var ret string
		return ret
	}
	return *o.BeginDate
}

// GetBeginDateOk returns a tuple with the BeginDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateRangeInfo) GetBeginDateOk() (*string, bool) {
	if o == nil || IsNil(o.BeginDate) {
		return nil, false
	}
	return o.BeginDate, true
}

// HasBeginDate returns a boolean if a field has been set.
func (o *DateRangeInfo) HasBeginDate() bool {
	if o != nil && !IsNil(o.BeginDate) {
		return true
	}

	return false
}

// SetBeginDate gets a reference to the given string and assigns it to the BeginDate field.
func (o *DateRangeInfo) SetBeginDate(v string) {
	o.BeginDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *DateRangeInfo) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateRangeInfo) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *DateRangeInfo) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *DateRangeInfo) SetEndDate(v string) {
	o.EndDate = &v
}

func (o DateRangeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DateRangeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BeginDate) {
		toSerialize["begin_date"] = o.BeginDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableDateRangeInfo struct {
	value *DateRangeInfo
	isSet bool
}

func (v NullableDateRangeInfo) Get() *DateRangeInfo {
	return v.value
}

func (v *NullableDateRangeInfo) Set(val *DateRangeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDateRangeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDateRangeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateRangeInfo(val *DateRangeInfo) *NullableDateRangeInfo {
	return &NullableDateRangeInfo{value: val, isSet: true}
}

func (v NullableDateRangeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateRangeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

