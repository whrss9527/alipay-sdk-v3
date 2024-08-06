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

// checks if the VoucherSendGuideModify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherSendGuideModify{}

// VoucherSendGuideModify struct for VoucherSendGuideModify
type VoucherSendGuideModify struct {
	// 领（购）券详情页链接，从支付宝公域跳转到服务商（商户）自定义领(购)券详情页。说明：当 voucher_type=EXCHANGE_VOUCHER 时，该字段可修改，其他不允许修改。
	VoucherDetailUrl *string `json:"voucher_detail_url,omitempty"`
}

// NewVoucherSendGuideModify instantiates a new VoucherSendGuideModify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherSendGuideModify() *VoucherSendGuideModify {
	this := VoucherSendGuideModify{}
	return &this
}

// NewVoucherSendGuideModifyWithDefaults instantiates a new VoucherSendGuideModify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherSendGuideModifyWithDefaults() *VoucherSendGuideModify {
	this := VoucherSendGuideModify{}
	return &this
}

// GetVoucherDetailUrl returns the VoucherDetailUrl field value if set, zero value otherwise.
func (o *VoucherSendGuideModify) GetVoucherDetailUrl() string {
	if o == nil || IsNil(o.VoucherDetailUrl) {
		var ret string
		return ret
	}
	return *o.VoucherDetailUrl
}

// GetVoucherDetailUrlOk returns a tuple with the VoucherDetailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherSendGuideModify) GetVoucherDetailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherDetailUrl) {
		return nil, false
	}
	return o.VoucherDetailUrl, true
}

// HasVoucherDetailUrl returns a boolean if a field has been set.
func (o *VoucherSendGuideModify) HasVoucherDetailUrl() bool {
	if o != nil && !IsNil(o.VoucherDetailUrl) {
		return true
	}

	return false
}

// SetVoucherDetailUrl gets a reference to the given string and assigns it to the VoucherDetailUrl field.
func (o *VoucherSendGuideModify) SetVoucherDetailUrl(v string) {
	o.VoucherDetailUrl = &v
}

func (o VoucherSendGuideModify) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherSendGuideModify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VoucherDetailUrl) {
		toSerialize["voucher_detail_url"] = o.VoucherDetailUrl
	}
	return toSerialize, nil
}

type NullableVoucherSendGuideModify struct {
	value *VoucherSendGuideModify
	isSet bool
}

func (v NullableVoucherSendGuideModify) Get() *VoucherSendGuideModify {
	return v.value
}

func (v *NullableVoucherSendGuideModify) Set(val *VoucherSendGuideModify) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherSendGuideModify) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherSendGuideModify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherSendGuideModify(val *VoucherSendGuideModify) *NullableVoucherSendGuideModify {
	return &NullableVoucherSendGuideModify{value: val, isSet: true}
}

func (v NullableVoucherSendGuideModify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherSendGuideModify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

