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

// checks if the LogisticsDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogisticsDetail{}

// LogisticsDetail struct for LogisticsDetail
type LogisticsDetail struct {
	// 物流类型
	LogisticsType *string `json:"logistics_type,omitempty"`
}

// NewLogisticsDetail instantiates a new LogisticsDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogisticsDetail() *LogisticsDetail {
	this := LogisticsDetail{}
	return &this
}

// NewLogisticsDetailWithDefaults instantiates a new LogisticsDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogisticsDetailWithDefaults() *LogisticsDetail {
	this := LogisticsDetail{}
	return &this
}

// GetLogisticsType returns the LogisticsType field value if set, zero value otherwise.
func (o *LogisticsDetail) GetLogisticsType() string {
	if o == nil || IsNil(o.LogisticsType) {
		var ret string
		return ret
	}
	return *o.LogisticsType
}

// GetLogisticsTypeOk returns a tuple with the LogisticsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogisticsDetail) GetLogisticsTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LogisticsType) {
		return nil, false
	}
	return o.LogisticsType, true
}

// HasLogisticsType returns a boolean if a field has been set.
func (o *LogisticsDetail) HasLogisticsType() bool {
	if o != nil && !IsNil(o.LogisticsType) {
		return true
	}

	return false
}

// SetLogisticsType gets a reference to the given string and assigns it to the LogisticsType field.
func (o *LogisticsDetail) SetLogisticsType(v string) {
	o.LogisticsType = &v
}

func (o LogisticsDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogisticsDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LogisticsType) {
		toSerialize["logistics_type"] = o.LogisticsType
	}
	return toSerialize, nil
}

type NullableLogisticsDetail struct {
	value *LogisticsDetail
	isSet bool
}

func (v NullableLogisticsDetail) Get() *LogisticsDetail {
	return v.value
}

func (v *NullableLogisticsDetail) Set(val *LogisticsDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableLogisticsDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableLogisticsDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogisticsDetail(val *LogisticsDetail) *NullableLogisticsDetail {
	return &NullableLogisticsDetail{value: val, isSet: true}
}

func (v NullableLogisticsDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogisticsDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
