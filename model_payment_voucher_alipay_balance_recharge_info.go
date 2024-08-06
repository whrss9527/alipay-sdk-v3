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

// checks if the PaymentVoucherAlipayBalanceRechargeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentVoucherAlipayBalanceRechargeInfo{}

// PaymentVoucherAlipayBalanceRechargeInfo struct for PaymentVoucherAlipayBalanceRechargeInfo
type PaymentVoucherAlipayBalanceRechargeInfo struct {
	// 出资的支付宝登录账号 限制: 1、登录账号和用户ID必须且只能二选一
	LogonId *string `json:"logon_id,omitempty"`
	// 出资的支付宝用户id 限制: 1、登录账号和用户ID必须且只能二选一
	UserId *string `json:"user_id,omitempty"`
}

// NewPaymentVoucherAlipayBalanceRechargeInfo instantiates a new PaymentVoucherAlipayBalanceRechargeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentVoucherAlipayBalanceRechargeInfo() *PaymentVoucherAlipayBalanceRechargeInfo {
	this := PaymentVoucherAlipayBalanceRechargeInfo{}
	return &this
}

// NewPaymentVoucherAlipayBalanceRechargeInfoWithDefaults instantiates a new PaymentVoucherAlipayBalanceRechargeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentVoucherAlipayBalanceRechargeInfoWithDefaults() *PaymentVoucherAlipayBalanceRechargeInfo {
	this := PaymentVoucherAlipayBalanceRechargeInfo{}
	return &this
}

// GetLogonId returns the LogonId field value if set, zero value otherwise.
func (o *PaymentVoucherAlipayBalanceRechargeInfo) GetLogonId() string {
	if o == nil || IsNil(o.LogonId) {
		var ret string
		return ret
	}
	return *o.LogonId
}

// GetLogonIdOk returns a tuple with the LogonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherAlipayBalanceRechargeInfo) GetLogonIdOk() (*string, bool) {
	if o == nil || IsNil(o.LogonId) {
		return nil, false
	}
	return o.LogonId, true
}

// HasLogonId returns a boolean if a field has been set.
func (o *PaymentVoucherAlipayBalanceRechargeInfo) HasLogonId() bool {
	if o != nil && !IsNil(o.LogonId) {
		return true
	}

	return false
}

// SetLogonId gets a reference to the given string and assigns it to the LogonId field.
func (o *PaymentVoucherAlipayBalanceRechargeInfo) SetLogonId(v string) {
	o.LogonId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PaymentVoucherAlipayBalanceRechargeInfo) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherAlipayBalanceRechargeInfo) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PaymentVoucherAlipayBalanceRechargeInfo) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *PaymentVoucherAlipayBalanceRechargeInfo) SetUserId(v string) {
	o.UserId = &v
}

func (o PaymentVoucherAlipayBalanceRechargeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentVoucherAlipayBalanceRechargeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LogonId) {
		toSerialize["logon_id"] = o.LogonId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullablePaymentVoucherAlipayBalanceRechargeInfo struct {
	value *PaymentVoucherAlipayBalanceRechargeInfo
	isSet bool
}

func (v NullablePaymentVoucherAlipayBalanceRechargeInfo) Get() *PaymentVoucherAlipayBalanceRechargeInfo {
	return v.value
}

func (v *NullablePaymentVoucherAlipayBalanceRechargeInfo) Set(val *PaymentVoucherAlipayBalanceRechargeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentVoucherAlipayBalanceRechargeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentVoucherAlipayBalanceRechargeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentVoucherAlipayBalanceRechargeInfo(val *PaymentVoucherAlipayBalanceRechargeInfo) *NullablePaymentVoucherAlipayBalanceRechargeInfo {
	return &NullablePaymentVoucherAlipayBalanceRechargeInfo{value: val, isSet: true}
}

func (v NullablePaymentVoucherAlipayBalanceRechargeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentVoucherAlipayBalanceRechargeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

