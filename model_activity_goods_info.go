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

// checks if the ActivityGoodsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityGoodsInfo{}

// ActivityGoodsInfo struct for ActivityGoodsInfo
type ActivityGoodsInfo struct {
	// 商品编码
	GoodsId *string `json:"goods_id,omitempty"`
	// 活动单品类型。
	GoodsUseType *string `json:"goods_use_type,omitempty"`
}

// NewActivityGoodsInfo instantiates a new ActivityGoodsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityGoodsInfo() *ActivityGoodsInfo {
	this := ActivityGoodsInfo{}
	return &this
}

// NewActivityGoodsInfoWithDefaults instantiates a new ActivityGoodsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityGoodsInfoWithDefaults() *ActivityGoodsInfo {
	this := ActivityGoodsInfo{}
	return &this
}

// GetGoodsId returns the GoodsId field value if set, zero value otherwise.
func (o *ActivityGoodsInfo) GetGoodsId() string {
	if o == nil || IsNil(o.GoodsId) {
		var ret string
		return ret
	}
	return *o.GoodsId
}

// GetGoodsIdOk returns a tuple with the GoodsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityGoodsInfo) GetGoodsIdOk() (*string, bool) {
	if o == nil || IsNil(o.GoodsId) {
		return nil, false
	}
	return o.GoodsId, true
}

// HasGoodsId returns a boolean if a field has been set.
func (o *ActivityGoodsInfo) HasGoodsId() bool {
	if o != nil && !IsNil(o.GoodsId) {
		return true
	}

	return false
}

// SetGoodsId gets a reference to the given string and assigns it to the GoodsId field.
func (o *ActivityGoodsInfo) SetGoodsId(v string) {
	o.GoodsId = &v
}

// GetGoodsUseType returns the GoodsUseType field value if set, zero value otherwise.
func (o *ActivityGoodsInfo) GetGoodsUseType() string {
	if o == nil || IsNil(o.GoodsUseType) {
		var ret string
		return ret
	}
	return *o.GoodsUseType
}

// GetGoodsUseTypeOk returns a tuple with the GoodsUseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityGoodsInfo) GetGoodsUseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GoodsUseType) {
		return nil, false
	}
	return o.GoodsUseType, true
}

// HasGoodsUseType returns a boolean if a field has been set.
func (o *ActivityGoodsInfo) HasGoodsUseType() bool {
	if o != nil && !IsNil(o.GoodsUseType) {
		return true
	}

	return false
}

// SetGoodsUseType gets a reference to the given string and assigns it to the GoodsUseType field.
func (o *ActivityGoodsInfo) SetGoodsUseType(v string) {
	o.GoodsUseType = &v
}

func (o ActivityGoodsInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityGoodsInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GoodsId) {
		toSerialize["goods_id"] = o.GoodsId
	}
	if !IsNil(o.GoodsUseType) {
		toSerialize["goods_use_type"] = o.GoodsUseType
	}
	return toSerialize, nil
}

type NullableActivityGoodsInfo struct {
	value *ActivityGoodsInfo
	isSet bool
}

func (v NullableActivityGoodsInfo) Get() *ActivityGoodsInfo {
	return v.value
}

func (v *NullableActivityGoodsInfo) Set(val *ActivityGoodsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityGoodsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityGoodsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityGoodsInfo(val *ActivityGoodsInfo) *NullableActivityGoodsInfo {
	return &NullableActivityGoodsInfo{value: val, isSet: true}
}

func (v NullableActivityGoodsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityGoodsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

