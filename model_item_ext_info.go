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

// checks if the ItemExtInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemExtInfo{}

// ItemExtInfo struct for ItemExtInfo
type ItemExtInfo struct {
	// 扩展信息的key
	ExtKey *string `json:"ext_key,omitempty"`
	// 扩展信息的值
	ExtValue *string `json:"ext_value,omitempty"`
}

// NewItemExtInfo instantiates a new ItemExtInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemExtInfo() *ItemExtInfo {
	this := ItemExtInfo{}
	return &this
}

// NewItemExtInfoWithDefaults instantiates a new ItemExtInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemExtInfoWithDefaults() *ItemExtInfo {
	this := ItemExtInfo{}
	return &this
}

// GetExtKey returns the ExtKey field value if set, zero value otherwise.
func (o *ItemExtInfo) GetExtKey() string {
	if o == nil || IsNil(o.ExtKey) {
		var ret string
		return ret
	}
	return *o.ExtKey
}

// GetExtKeyOk returns a tuple with the ExtKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemExtInfo) GetExtKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ExtKey) {
		return nil, false
	}
	return o.ExtKey, true
}

// HasExtKey returns a boolean if a field has been set.
func (o *ItemExtInfo) HasExtKey() bool {
	if o != nil && !IsNil(o.ExtKey) {
		return true
	}

	return false
}

// SetExtKey gets a reference to the given string and assigns it to the ExtKey field.
func (o *ItemExtInfo) SetExtKey(v string) {
	o.ExtKey = &v
}

// GetExtValue returns the ExtValue field value if set, zero value otherwise.
func (o *ItemExtInfo) GetExtValue() string {
	if o == nil || IsNil(o.ExtValue) {
		var ret string
		return ret
	}
	return *o.ExtValue
}

// GetExtValueOk returns a tuple with the ExtValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemExtInfo) GetExtValueOk() (*string, bool) {
	if o == nil || IsNil(o.ExtValue) {
		return nil, false
	}
	return o.ExtValue, true
}

// HasExtValue returns a boolean if a field has been set.
func (o *ItemExtInfo) HasExtValue() bool {
	if o != nil && !IsNil(o.ExtValue) {
		return true
	}

	return false
}

// SetExtValue gets a reference to the given string and assigns it to the ExtValue field.
func (o *ItemExtInfo) SetExtValue(v string) {
	o.ExtValue = &v
}

func (o ItemExtInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemExtInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtKey) {
		toSerialize["ext_key"] = o.ExtKey
	}
	if !IsNil(o.ExtValue) {
		toSerialize["ext_value"] = o.ExtValue
	}
	return toSerialize, nil
}

type NullableItemExtInfo struct {
	value *ItemExtInfo
	isSet bool
}

func (v NullableItemExtInfo) Get() *ItemExtInfo {
	return v.value
}

func (v *NullableItemExtInfo) Set(val *ItemExtInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableItemExtInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableItemExtInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemExtInfo(val *ItemExtInfo) *NullableItemExtInfo {
	return &NullableItemExtInfo{value: val, isSet: true}
}

func (v NullableItemExtInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemExtInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
