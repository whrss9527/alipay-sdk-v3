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

// checks if the AlipayMarketingActivityVoucherAppendResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityVoucherAppendResponseModel{}

// AlipayMarketingActivityVoucherAppendResponseModel struct for AlipayMarketingActivityVoucherAppendResponseModel
type AlipayMarketingActivityVoucherAppendResponseModel struct {
	// 预充值链接  限制 有效时间3天
	RechargeUrl *string `json:"recharge_url,omitempty"`
}

// NewAlipayMarketingActivityVoucherAppendResponseModel instantiates a new AlipayMarketingActivityVoucherAppendResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityVoucherAppendResponseModel() *AlipayMarketingActivityVoucherAppendResponseModel {
	this := AlipayMarketingActivityVoucherAppendResponseModel{}
	return &this
}

// NewAlipayMarketingActivityVoucherAppendResponseModelWithDefaults instantiates a new AlipayMarketingActivityVoucherAppendResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityVoucherAppendResponseModelWithDefaults() *AlipayMarketingActivityVoucherAppendResponseModel {
	this := AlipayMarketingActivityVoucherAppendResponseModel{}
	return &this
}

// GetRechargeUrl returns the RechargeUrl field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherAppendResponseModel) GetRechargeUrl() string {
	if o == nil || IsNil(o.RechargeUrl) {
		var ret string
		return ret
	}
	return *o.RechargeUrl
}

// GetRechargeUrlOk returns a tuple with the RechargeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherAppendResponseModel) GetRechargeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RechargeUrl) {
		return nil, false
	}
	return o.RechargeUrl, true
}

// HasRechargeUrl returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherAppendResponseModel) HasRechargeUrl() bool {
	if o != nil && !IsNil(o.RechargeUrl) {
		return true
	}

	return false
}

// SetRechargeUrl gets a reference to the given string and assigns it to the RechargeUrl field.
func (o *AlipayMarketingActivityVoucherAppendResponseModel) SetRechargeUrl(v string) {
	o.RechargeUrl = &v
}

func (o AlipayMarketingActivityVoucherAppendResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityVoucherAppendResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RechargeUrl) {
		toSerialize["recharge_url"] = o.RechargeUrl
	}
	return toSerialize, nil
}

type NullableAlipayMarketingActivityVoucherAppendResponseModel struct {
	value *AlipayMarketingActivityVoucherAppendResponseModel
	isSet bool
}

func (v NullableAlipayMarketingActivityVoucherAppendResponseModel) Get() *AlipayMarketingActivityVoucherAppendResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityVoucherAppendResponseModel) Set(val *AlipayMarketingActivityVoucherAppendResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityVoucherAppendResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityVoucherAppendResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityVoucherAppendResponseModel(val *AlipayMarketingActivityVoucherAppendResponseModel) *NullableAlipayMarketingActivityVoucherAppendResponseModel {
	return &NullableAlipayMarketingActivityVoucherAppendResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityVoucherAppendResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityVoucherAppendResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
