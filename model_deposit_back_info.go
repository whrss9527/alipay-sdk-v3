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

// checks if the DepositBackInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DepositBackInfo{}

// DepositBackInfo struct for DepositBackInfo
type DepositBackInfo struct {
	// 银行响应时间，格式为yyyy-MM-dd HH:mm:ss
	BankAckTime *string `json:"bank_ack_time,omitempty"`
	// 银行卡冲退金额。单位：元。
	DbackAmount *string `json:"dback_amount,omitempty"`
	// 银行卡冲退状态。S-成功，F-失败，P-处理中。银行卡冲退失败，资金自动转入用户支付宝余额。
	DbackStatus *string `json:"dback_status,omitempty"`
	// 预估银行到账时间，格式为yyyy-MM-dd HH:mm:ss
	EstBankReceiptTime *string `json:"est_bank_receipt_time,omitempty"`
	// 是否存在银行卡冲退信息。
	HasDepositBack *string `json:"has_deposit_back,omitempty"`
}

// NewDepositBackInfo instantiates a new DepositBackInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositBackInfo() *DepositBackInfo {
	this := DepositBackInfo{}
	return &this
}

// NewDepositBackInfoWithDefaults instantiates a new DepositBackInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositBackInfoWithDefaults() *DepositBackInfo {
	this := DepositBackInfo{}
	return &this
}

// GetBankAckTime returns the BankAckTime field value if set, zero value otherwise.
func (o *DepositBackInfo) GetBankAckTime() string {
	if o == nil || IsNil(o.BankAckTime) {
		var ret string
		return ret
	}
	return *o.BankAckTime
}

// GetBankAckTimeOk returns a tuple with the BankAckTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositBackInfo) GetBankAckTimeOk() (*string, bool) {
	if o == nil || IsNil(o.BankAckTime) {
		return nil, false
	}
	return o.BankAckTime, true
}

// HasBankAckTime returns a boolean if a field has been set.
func (o *DepositBackInfo) HasBankAckTime() bool {
	if o != nil && !IsNil(o.BankAckTime) {
		return true
	}

	return false
}

// SetBankAckTime gets a reference to the given string and assigns it to the BankAckTime field.
func (o *DepositBackInfo) SetBankAckTime(v string) {
	o.BankAckTime = &v
}

// GetDbackAmount returns the DbackAmount field value if set, zero value otherwise.
func (o *DepositBackInfo) GetDbackAmount() string {
	if o == nil || IsNil(o.DbackAmount) {
		var ret string
		return ret
	}
	return *o.DbackAmount
}

// GetDbackAmountOk returns a tuple with the DbackAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositBackInfo) GetDbackAmountOk() (*string, bool) {
	if o == nil || IsNil(o.DbackAmount) {
		return nil, false
	}
	return o.DbackAmount, true
}

// HasDbackAmount returns a boolean if a field has been set.
func (o *DepositBackInfo) HasDbackAmount() bool {
	if o != nil && !IsNil(o.DbackAmount) {
		return true
	}

	return false
}

// SetDbackAmount gets a reference to the given string and assigns it to the DbackAmount field.
func (o *DepositBackInfo) SetDbackAmount(v string) {
	o.DbackAmount = &v
}

// GetDbackStatus returns the DbackStatus field value if set, zero value otherwise.
func (o *DepositBackInfo) GetDbackStatus() string {
	if o == nil || IsNil(o.DbackStatus) {
		var ret string
		return ret
	}
	return *o.DbackStatus
}

// GetDbackStatusOk returns a tuple with the DbackStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositBackInfo) GetDbackStatusOk() (*string, bool) {
	if o == nil || IsNil(o.DbackStatus) {
		return nil, false
	}
	return o.DbackStatus, true
}

// HasDbackStatus returns a boolean if a field has been set.
func (o *DepositBackInfo) HasDbackStatus() bool {
	if o != nil && !IsNil(o.DbackStatus) {
		return true
	}

	return false
}

// SetDbackStatus gets a reference to the given string and assigns it to the DbackStatus field.
func (o *DepositBackInfo) SetDbackStatus(v string) {
	o.DbackStatus = &v
}

// GetEstBankReceiptTime returns the EstBankReceiptTime field value if set, zero value otherwise.
func (o *DepositBackInfo) GetEstBankReceiptTime() string {
	if o == nil || IsNil(o.EstBankReceiptTime) {
		var ret string
		return ret
	}
	return *o.EstBankReceiptTime
}

// GetEstBankReceiptTimeOk returns a tuple with the EstBankReceiptTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositBackInfo) GetEstBankReceiptTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EstBankReceiptTime) {
		return nil, false
	}
	return o.EstBankReceiptTime, true
}

// HasEstBankReceiptTime returns a boolean if a field has been set.
func (o *DepositBackInfo) HasEstBankReceiptTime() bool {
	if o != nil && !IsNil(o.EstBankReceiptTime) {
		return true
	}

	return false
}

// SetEstBankReceiptTime gets a reference to the given string and assigns it to the EstBankReceiptTime field.
func (o *DepositBackInfo) SetEstBankReceiptTime(v string) {
	o.EstBankReceiptTime = &v
}

// GetHasDepositBack returns the HasDepositBack field value if set, zero value otherwise.
func (o *DepositBackInfo) GetHasDepositBack() string {
	if o == nil || IsNil(o.HasDepositBack) {
		var ret string
		return ret
	}
	return *o.HasDepositBack
}

// GetHasDepositBackOk returns a tuple with the HasDepositBack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositBackInfo) GetHasDepositBackOk() (*string, bool) {
	if o == nil || IsNil(o.HasDepositBack) {
		return nil, false
	}
	return o.HasDepositBack, true
}

// HasHasDepositBack returns a boolean if a field has been set.
func (o *DepositBackInfo) HasHasDepositBack() bool {
	if o != nil && !IsNil(o.HasDepositBack) {
		return true
	}

	return false
}

// SetHasDepositBack gets a reference to the given string and assigns it to the HasDepositBack field.
func (o *DepositBackInfo) SetHasDepositBack(v string) {
	o.HasDepositBack = &v
}

func (o DepositBackInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositBackInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BankAckTime) {
		toSerialize["bank_ack_time"] = o.BankAckTime
	}
	if !IsNil(o.DbackAmount) {
		toSerialize["dback_amount"] = o.DbackAmount
	}
	if !IsNil(o.DbackStatus) {
		toSerialize["dback_status"] = o.DbackStatus
	}
	if !IsNil(o.EstBankReceiptTime) {
		toSerialize["est_bank_receipt_time"] = o.EstBankReceiptTime
	}
	if !IsNil(o.HasDepositBack) {
		toSerialize["has_deposit_back"] = o.HasDepositBack
	}
	return toSerialize, nil
}

type NullableDepositBackInfo struct {
	value *DepositBackInfo
	isSet bool
}

func (v NullableDepositBackInfo) Get() *DepositBackInfo {
	return v.value
}

func (v *NullableDepositBackInfo) Set(val *DepositBackInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositBackInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositBackInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositBackInfo(val *DepositBackInfo) *NullableDepositBackInfo {
	return &NullableDepositBackInfo{value: val, isSet: true}
}

func (v NullableDepositBackInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositBackInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
