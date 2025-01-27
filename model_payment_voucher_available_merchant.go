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

// checks if the PaymentVoucherAvailableMerchant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentVoucherAvailableMerchant{}

// PaymentVoucherAvailableMerchant struct for PaymentVoucherAvailableMerchant
type PaymentVoucherAvailableMerchant struct {
	// 优惠券可核销的直连商户PID  限制： 1、available_pids和available_smids至少二选一。
	AvailablePids []string `json:"available_pids,omitempty"`
	// 优惠券可核销的间连商户SMID  限制：  1、available_pids和available_smids至少二选一。
	AvailableSmids []string `json:"available_smids,omitempty"`
}

// NewPaymentVoucherAvailableMerchant instantiates a new PaymentVoucherAvailableMerchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentVoucherAvailableMerchant() *PaymentVoucherAvailableMerchant {
	this := PaymentVoucherAvailableMerchant{}
	return &this
}

// NewPaymentVoucherAvailableMerchantWithDefaults instantiates a new PaymentVoucherAvailableMerchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentVoucherAvailableMerchantWithDefaults() *PaymentVoucherAvailableMerchant {
	this := PaymentVoucherAvailableMerchant{}
	return &this
}

// GetAvailablePids returns the AvailablePids field value if set, zero value otherwise.
func (o *PaymentVoucherAvailableMerchant) GetAvailablePids() []string {
	if o == nil || IsNil(o.AvailablePids) {
		var ret []string
		return ret
	}
	return o.AvailablePids
}

// GetAvailablePidsOk returns a tuple with the AvailablePids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherAvailableMerchant) GetAvailablePidsOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailablePids) {
		return nil, false
	}
	return o.AvailablePids, true
}

// HasAvailablePids returns a boolean if a field has been set.
func (o *PaymentVoucherAvailableMerchant) HasAvailablePids() bool {
	if o != nil && !IsNil(o.AvailablePids) {
		return true
	}

	return false
}

// SetAvailablePids gets a reference to the given []string and assigns it to the AvailablePids field.
func (o *PaymentVoucherAvailableMerchant) SetAvailablePids(v []string) {
	o.AvailablePids = v
}

// GetAvailableSmids returns the AvailableSmids field value if set, zero value otherwise.
func (o *PaymentVoucherAvailableMerchant) GetAvailableSmids() []string {
	if o == nil || IsNil(o.AvailableSmids) {
		var ret []string
		return ret
	}
	return o.AvailableSmids
}

// GetAvailableSmidsOk returns a tuple with the AvailableSmids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherAvailableMerchant) GetAvailableSmidsOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailableSmids) {
		return nil, false
	}
	return o.AvailableSmids, true
}

// HasAvailableSmids returns a boolean if a field has been set.
func (o *PaymentVoucherAvailableMerchant) HasAvailableSmids() bool {
	if o != nil && !IsNil(o.AvailableSmids) {
		return true
	}

	return false
}

// SetAvailableSmids gets a reference to the given []string and assigns it to the AvailableSmids field.
func (o *PaymentVoucherAvailableMerchant) SetAvailableSmids(v []string) {
	o.AvailableSmids = v
}

func (o PaymentVoucherAvailableMerchant) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentVoucherAvailableMerchant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailablePids) {
		toSerialize["available_pids"] = o.AvailablePids
	}
	if !IsNil(o.AvailableSmids) {
		toSerialize["available_smids"] = o.AvailableSmids
	}
	return toSerialize, nil
}

type NullablePaymentVoucherAvailableMerchant struct {
	value *PaymentVoucherAvailableMerchant
	isSet bool
}

func (v NullablePaymentVoucherAvailableMerchant) Get() *PaymentVoucherAvailableMerchant {
	return v.value
}

func (v *NullablePaymentVoucherAvailableMerchant) Set(val *PaymentVoucherAvailableMerchant) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentVoucherAvailableMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentVoucherAvailableMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentVoucherAvailableMerchant(val *PaymentVoucherAvailableMerchant) *NullablePaymentVoucherAvailableMerchant {
	return &NullablePaymentVoucherAvailableMerchant{value: val, isSet: true}
}

func (v NullablePaymentVoucherAvailableMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentVoucherAvailableMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
