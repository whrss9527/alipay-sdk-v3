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

// checks if the VoucherBalanceRechargeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherBalanceRechargeInfo{}

// VoucherBalanceRechargeInfo struct for VoucherBalanceRechargeInfo
type VoucherBalanceRechargeInfo struct {
	// 支付宝余额充值金额  限制：   1.币种为人民币，单位元。   2. 总预算=优惠金额*总发券张数
	Amount *string `json:"amount,omitempty"`
	// 出资的商户支付宝登录账号
	LogonId *string `json:"logon_id,omitempty"`
	// 出资的商户支付宝ID
	PartnerId *string `json:"partner_id,omitempty"`
}

// NewVoucherBalanceRechargeInfo instantiates a new VoucherBalanceRechargeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherBalanceRechargeInfo() *VoucherBalanceRechargeInfo {
	this := VoucherBalanceRechargeInfo{}
	return &this
}

// NewVoucherBalanceRechargeInfoWithDefaults instantiates a new VoucherBalanceRechargeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherBalanceRechargeInfoWithDefaults() *VoucherBalanceRechargeInfo {
	this := VoucherBalanceRechargeInfo{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *VoucherBalanceRechargeInfo) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherBalanceRechargeInfo) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *VoucherBalanceRechargeInfo) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *VoucherBalanceRechargeInfo) SetAmount(v string) {
	o.Amount = &v
}

// GetLogonId returns the LogonId field value if set, zero value otherwise.
func (o *VoucherBalanceRechargeInfo) GetLogonId() string {
	if o == nil || IsNil(o.LogonId) {
		var ret string
		return ret
	}
	return *o.LogonId
}

// GetLogonIdOk returns a tuple with the LogonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherBalanceRechargeInfo) GetLogonIdOk() (*string, bool) {
	if o == nil || IsNil(o.LogonId) {
		return nil, false
	}
	return o.LogonId, true
}

// HasLogonId returns a boolean if a field has been set.
func (o *VoucherBalanceRechargeInfo) HasLogonId() bool {
	if o != nil && !IsNil(o.LogonId) {
		return true
	}

	return false
}

// SetLogonId gets a reference to the given string and assigns it to the LogonId field.
func (o *VoucherBalanceRechargeInfo) SetLogonId(v string) {
	o.LogonId = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *VoucherBalanceRechargeInfo) GetPartnerId() string {
	if o == nil || IsNil(o.PartnerId) {
		var ret string
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherBalanceRechargeInfo) GetPartnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *VoucherBalanceRechargeInfo) HasPartnerId() bool {
	if o != nil && !IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given string and assigns it to the PartnerId field.
func (o *VoucherBalanceRechargeInfo) SetPartnerId(v string) {
	o.PartnerId = &v
}

func (o VoucherBalanceRechargeInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherBalanceRechargeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.LogonId) {
		toSerialize["logon_id"] = o.LogonId
	}
	if !IsNil(o.PartnerId) {
		toSerialize["partner_id"] = o.PartnerId
	}
	return toSerialize, nil
}

type NullableVoucherBalanceRechargeInfo struct {
	value *VoucherBalanceRechargeInfo
	isSet bool
}

func (v NullableVoucherBalanceRechargeInfo) Get() *VoucherBalanceRechargeInfo {
	return v.value
}

func (v *NullableVoucherBalanceRechargeInfo) Set(val *VoucherBalanceRechargeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherBalanceRechargeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherBalanceRechargeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherBalanceRechargeInfo(val *VoucherBalanceRechargeInfo) *NullableVoucherBalanceRechargeInfo {
	return &NullableVoucherBalanceRechargeInfo{value: val, isSet: true}
}

func (v NullableVoucherBalanceRechargeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherBalanceRechargeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
