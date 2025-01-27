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

// checks if the VoucherExchangeGoodsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherExchangeGoodsInfo{}

// VoucherExchangeGoodsInfo struct for VoucherExchangeGoodsInfo
type VoucherExchangeGoodsInfo struct {
	// 兑换商品名称
	ExchangeGoodsName *string `json:"exchange_goods_name,omitempty"`
}

// NewVoucherExchangeGoodsInfo instantiates a new VoucherExchangeGoodsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherExchangeGoodsInfo() *VoucherExchangeGoodsInfo {
	this := VoucherExchangeGoodsInfo{}
	return &this
}

// NewVoucherExchangeGoodsInfoWithDefaults instantiates a new VoucherExchangeGoodsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherExchangeGoodsInfoWithDefaults() *VoucherExchangeGoodsInfo {
	this := VoucherExchangeGoodsInfo{}
	return &this
}

// GetExchangeGoodsName returns the ExchangeGoodsName field value if set, zero value otherwise.
func (o *VoucherExchangeGoodsInfo) GetExchangeGoodsName() string {
	if o == nil || IsNil(o.ExchangeGoodsName) {
		var ret string
		return ret
	}
	return *o.ExchangeGoodsName
}

// GetExchangeGoodsNameOk returns a tuple with the ExchangeGoodsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherExchangeGoodsInfo) GetExchangeGoodsNameOk() (*string, bool) {
	if o == nil || IsNil(o.ExchangeGoodsName) {
		return nil, false
	}
	return o.ExchangeGoodsName, true
}

// HasExchangeGoodsName returns a boolean if a field has been set.
func (o *VoucherExchangeGoodsInfo) HasExchangeGoodsName() bool {
	if o != nil && !IsNil(o.ExchangeGoodsName) {
		return true
	}

	return false
}

// SetExchangeGoodsName gets a reference to the given string and assigns it to the ExchangeGoodsName field.
func (o *VoucherExchangeGoodsInfo) SetExchangeGoodsName(v string) {
	o.ExchangeGoodsName = &v
}

func (o VoucherExchangeGoodsInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherExchangeGoodsInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExchangeGoodsName) {
		toSerialize["exchange_goods_name"] = o.ExchangeGoodsName
	}
	return toSerialize, nil
}

type NullableVoucherExchangeGoodsInfo struct {
	value *VoucherExchangeGoodsInfo
	isSet bool
}

func (v NullableVoucherExchangeGoodsInfo) Get() *VoucherExchangeGoodsInfo {
	return v.value
}

func (v *NullableVoucherExchangeGoodsInfo) Set(val *VoucherExchangeGoodsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherExchangeGoodsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherExchangeGoodsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherExchangeGoodsInfo(val *VoucherExchangeGoodsInfo) *NullableVoucherExchangeGoodsInfo {
	return &NullableVoucherExchangeGoodsInfo{value: val, isSet: true}
}

func (v NullableVoucherExchangeGoodsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherExchangeGoodsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
