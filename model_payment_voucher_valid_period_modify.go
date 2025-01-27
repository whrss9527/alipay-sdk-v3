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

// checks if the PaymentVoucherValidPeriodModify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentVoucherValidPeriodModify{}

// PaymentVoucherValidPeriodModify struct for PaymentVoucherValidPeriodModify
type PaymentVoucherValidPeriodModify struct {
	// 券生效后N天内可以使用。 限制： type为RELATIVE时可修改。 valid_days_after_receive必须大于0。  修改该时间，只允许延长，不允许缩短。
	ValidDaysAfterReceive *int32 `json:"valid_days_after_receive,omitempty"`
	// 券可使用的结束时间。 格式为yyyy-MM-dd HH:mm:ss 限制： type为ABSOLUTE可修改。 券可使用的结束时间valid_end_time 必须大于 券的发放结束时间 publish_end_time 修改券可使用的结束时间，只能延长，不允许缩短。
	ValidEndTime *string `json:"valid_end_time,omitempty"`
}

// NewPaymentVoucherValidPeriodModify instantiates a new PaymentVoucherValidPeriodModify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentVoucherValidPeriodModify() *PaymentVoucherValidPeriodModify {
	this := PaymentVoucherValidPeriodModify{}
	return &this
}

// NewPaymentVoucherValidPeriodModifyWithDefaults instantiates a new PaymentVoucherValidPeriodModify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentVoucherValidPeriodModifyWithDefaults() *PaymentVoucherValidPeriodModify {
	this := PaymentVoucherValidPeriodModify{}
	return &this
}

// GetValidDaysAfterReceive returns the ValidDaysAfterReceive field value if set, zero value otherwise.
func (o *PaymentVoucherValidPeriodModify) GetValidDaysAfterReceive() int32 {
	if o == nil || IsNil(o.ValidDaysAfterReceive) {
		var ret int32
		return ret
	}
	return *o.ValidDaysAfterReceive
}

// GetValidDaysAfterReceiveOk returns a tuple with the ValidDaysAfterReceive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherValidPeriodModify) GetValidDaysAfterReceiveOk() (*int32, bool) {
	if o == nil || IsNil(o.ValidDaysAfterReceive) {
		return nil, false
	}
	return o.ValidDaysAfterReceive, true
}

// HasValidDaysAfterReceive returns a boolean if a field has been set.
func (o *PaymentVoucherValidPeriodModify) HasValidDaysAfterReceive() bool {
	if o != nil && !IsNil(o.ValidDaysAfterReceive) {
		return true
	}

	return false
}

// SetValidDaysAfterReceive gets a reference to the given int32 and assigns it to the ValidDaysAfterReceive field.
func (o *PaymentVoucherValidPeriodModify) SetValidDaysAfterReceive(v int32) {
	o.ValidDaysAfterReceive = &v
}

// GetValidEndTime returns the ValidEndTime field value if set, zero value otherwise.
func (o *PaymentVoucherValidPeriodModify) GetValidEndTime() string {
	if o == nil || IsNil(o.ValidEndTime) {
		var ret string
		return ret
	}
	return *o.ValidEndTime
}

// GetValidEndTimeOk returns a tuple with the ValidEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherValidPeriodModify) GetValidEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ValidEndTime) {
		return nil, false
	}
	return o.ValidEndTime, true
}

// HasValidEndTime returns a boolean if a field has been set.
func (o *PaymentVoucherValidPeriodModify) HasValidEndTime() bool {
	if o != nil && !IsNil(o.ValidEndTime) {
		return true
	}

	return false
}

// SetValidEndTime gets a reference to the given string and assigns it to the ValidEndTime field.
func (o *PaymentVoucherValidPeriodModify) SetValidEndTime(v string) {
	o.ValidEndTime = &v
}

func (o PaymentVoucherValidPeriodModify) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentVoucherValidPeriodModify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ValidDaysAfterReceive) {
		toSerialize["valid_days_after_receive"] = o.ValidDaysAfterReceive
	}
	if !IsNil(o.ValidEndTime) {
		toSerialize["valid_end_time"] = o.ValidEndTime
	}
	return toSerialize, nil
}

type NullablePaymentVoucherValidPeriodModify struct {
	value *PaymentVoucherValidPeriodModify
	isSet bool
}

func (v NullablePaymentVoucherValidPeriodModify) Get() *PaymentVoucherValidPeriodModify {
	return v.value
}

func (v *NullablePaymentVoucherValidPeriodModify) Set(val *PaymentVoucherValidPeriodModify) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentVoucherValidPeriodModify) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentVoucherValidPeriodModify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentVoucherValidPeriodModify(val *PaymentVoucherValidPeriodModify) *NullablePaymentVoucherValidPeriodModify {
	return &NullablePaymentVoucherValidPeriodModify{value: val, isSet: true}
}

func (v NullablePaymentVoucherValidPeriodModify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentVoucherValidPeriodModify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
