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

// checks if the PaymentVoucherSendRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentVoucherSendRule{}

// PaymentVoucherSendRule struct for PaymentVoucherSendRule
type PaymentVoucherSendRule struct {
	// 设置此字段，允许指定单天最大发券数量。  限制: 每天发放张数*活动天数应小于等于优惠券发放总量
	MaxQuantityByDay *int32 `json:"max_quantity_by_day,omitempty"`
	// 限制相同身份证号领取次数(voucher_quantity_limit_per_user)。默认false不限制。 枚举值 true：是 false：否
	NaturalPersonLimit *bool `json:"natural_person_limit,omitempty"`
	// 限制相同手机号领取次数(voucher_quantity_limit_per_user)。默认false不限制 枚举值 true：是 false：否
	PhoneNumberLimit *bool `json:"phone_number_limit,omitempty"`
	// 限制支付宝实名用户才能领取支付券,默认为false表示不限制 枚举值 true\\false
	RealNameLimit *bool `json:"real_name_limit,omitempty"`
	// 发行券的总数量。 限制： 1、发放总个数最少1个 2、发放总个数最多99999999个
	VoucherQuantity *int32 `json:"voucher_quantity,omitempty"`
	// 每人领取限制。 默认按照支付宝账号进行领取限制;  不填写或填入0表示没有领取限制.
	VoucherQuantityLimitPerUser *int32 `json:"voucher_quantity_limit_per_user,omitempty"`
	// 周期限领配置,限制每人在固定周期内领取张数(voucher_quantity_limit_per_user),默认LIFE_CYCLE
	VoucherQuantityLimitPerUserPeriodType *string `json:"voucher_quantity_limit_per_user_period_type,omitempty"`
}

// NewPaymentVoucherSendRule instantiates a new PaymentVoucherSendRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentVoucherSendRule() *PaymentVoucherSendRule {
	this := PaymentVoucherSendRule{}
	return &this
}

// NewPaymentVoucherSendRuleWithDefaults instantiates a new PaymentVoucherSendRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentVoucherSendRuleWithDefaults() *PaymentVoucherSendRule {
	this := PaymentVoucherSendRule{}
	return &this
}

// GetMaxQuantityByDay returns the MaxQuantityByDay field value if set, zero value otherwise.
func (o *PaymentVoucherSendRule) GetMaxQuantityByDay() int32 {
	if o == nil || IsNil(o.MaxQuantityByDay) {
		var ret int32
		return ret
	}
	return *o.MaxQuantityByDay
}

// GetMaxQuantityByDayOk returns a tuple with the MaxQuantityByDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherSendRule) GetMaxQuantityByDayOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxQuantityByDay) {
		return nil, false
	}
	return o.MaxQuantityByDay, true
}

// HasMaxQuantityByDay returns a boolean if a field has been set.
func (o *PaymentVoucherSendRule) HasMaxQuantityByDay() bool {
	if o != nil && !IsNil(o.MaxQuantityByDay) {
		return true
	}

	return false
}

// SetMaxQuantityByDay gets a reference to the given int32 and assigns it to the MaxQuantityByDay field.
func (o *PaymentVoucherSendRule) SetMaxQuantityByDay(v int32) {
	o.MaxQuantityByDay = &v
}

// GetNaturalPersonLimit returns the NaturalPersonLimit field value if set, zero value otherwise.
func (o *PaymentVoucherSendRule) GetNaturalPersonLimit() bool {
	if o == nil || IsNil(o.NaturalPersonLimit) {
		var ret bool
		return ret
	}
	return *o.NaturalPersonLimit
}

// GetNaturalPersonLimitOk returns a tuple with the NaturalPersonLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherSendRule) GetNaturalPersonLimitOk() (*bool, bool) {
	if o == nil || IsNil(o.NaturalPersonLimit) {
		return nil, false
	}
	return o.NaturalPersonLimit, true
}

// HasNaturalPersonLimit returns a boolean if a field has been set.
func (o *PaymentVoucherSendRule) HasNaturalPersonLimit() bool {
	if o != nil && !IsNil(o.NaturalPersonLimit) {
		return true
	}

	return false
}

// SetNaturalPersonLimit gets a reference to the given bool and assigns it to the NaturalPersonLimit field.
func (o *PaymentVoucherSendRule) SetNaturalPersonLimit(v bool) {
	o.NaturalPersonLimit = &v
}

// GetPhoneNumberLimit returns the PhoneNumberLimit field value if set, zero value otherwise.
func (o *PaymentVoucherSendRule) GetPhoneNumberLimit() bool {
	if o == nil || IsNil(o.PhoneNumberLimit) {
		var ret bool
		return ret
	}
	return *o.PhoneNumberLimit
}

// GetPhoneNumberLimitOk returns a tuple with the PhoneNumberLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherSendRule) GetPhoneNumberLimitOk() (*bool, bool) {
	if o == nil || IsNil(o.PhoneNumberLimit) {
		return nil, false
	}
	return o.PhoneNumberLimit, true
}

// HasPhoneNumberLimit returns a boolean if a field has been set.
func (o *PaymentVoucherSendRule) HasPhoneNumberLimit() bool {
	if o != nil && !IsNil(o.PhoneNumberLimit) {
		return true
	}

	return false
}

// SetPhoneNumberLimit gets a reference to the given bool and assigns it to the PhoneNumberLimit field.
func (o *PaymentVoucherSendRule) SetPhoneNumberLimit(v bool) {
	o.PhoneNumberLimit = &v
}

// GetRealNameLimit returns the RealNameLimit field value if set, zero value otherwise.
func (o *PaymentVoucherSendRule) GetRealNameLimit() bool {
	if o == nil || IsNil(o.RealNameLimit) {
		var ret bool
		return ret
	}
	return *o.RealNameLimit
}

// GetRealNameLimitOk returns a tuple with the RealNameLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherSendRule) GetRealNameLimitOk() (*bool, bool) {
	if o == nil || IsNil(o.RealNameLimit) {
		return nil, false
	}
	return o.RealNameLimit, true
}

// HasRealNameLimit returns a boolean if a field has been set.
func (o *PaymentVoucherSendRule) HasRealNameLimit() bool {
	if o != nil && !IsNil(o.RealNameLimit) {
		return true
	}

	return false
}

// SetRealNameLimit gets a reference to the given bool and assigns it to the RealNameLimit field.
func (o *PaymentVoucherSendRule) SetRealNameLimit(v bool) {
	o.RealNameLimit = &v
}

// GetVoucherQuantity returns the VoucherQuantity field value if set, zero value otherwise.
func (o *PaymentVoucherSendRule) GetVoucherQuantity() int32 {
	if o == nil || IsNil(o.VoucherQuantity) {
		var ret int32
		return ret
	}
	return *o.VoucherQuantity
}

// GetVoucherQuantityOk returns a tuple with the VoucherQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherSendRule) GetVoucherQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.VoucherQuantity) {
		return nil, false
	}
	return o.VoucherQuantity, true
}

// HasVoucherQuantity returns a boolean if a field has been set.
func (o *PaymentVoucherSendRule) HasVoucherQuantity() bool {
	if o != nil && !IsNil(o.VoucherQuantity) {
		return true
	}

	return false
}

// SetVoucherQuantity gets a reference to the given int32 and assigns it to the VoucherQuantity field.
func (o *PaymentVoucherSendRule) SetVoucherQuantity(v int32) {
	o.VoucherQuantity = &v
}

// GetVoucherQuantityLimitPerUser returns the VoucherQuantityLimitPerUser field value if set, zero value otherwise.
func (o *PaymentVoucherSendRule) GetVoucherQuantityLimitPerUser() int32 {
	if o == nil || IsNil(o.VoucherQuantityLimitPerUser) {
		var ret int32
		return ret
	}
	return *o.VoucherQuantityLimitPerUser
}

// GetVoucherQuantityLimitPerUserOk returns a tuple with the VoucherQuantityLimitPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherSendRule) GetVoucherQuantityLimitPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.VoucherQuantityLimitPerUser) {
		return nil, false
	}
	return o.VoucherQuantityLimitPerUser, true
}

// HasVoucherQuantityLimitPerUser returns a boolean if a field has been set.
func (o *PaymentVoucherSendRule) HasVoucherQuantityLimitPerUser() bool {
	if o != nil && !IsNil(o.VoucherQuantityLimitPerUser) {
		return true
	}

	return false
}

// SetVoucherQuantityLimitPerUser gets a reference to the given int32 and assigns it to the VoucherQuantityLimitPerUser field.
func (o *PaymentVoucherSendRule) SetVoucherQuantityLimitPerUser(v int32) {
	o.VoucherQuantityLimitPerUser = &v
}

// GetVoucherQuantityLimitPerUserPeriodType returns the VoucherQuantityLimitPerUserPeriodType field value if set, zero value otherwise.
func (o *PaymentVoucherSendRule) GetVoucherQuantityLimitPerUserPeriodType() string {
	if o == nil || IsNil(o.VoucherQuantityLimitPerUserPeriodType) {
		var ret string
		return ret
	}
	return *o.VoucherQuantityLimitPerUserPeriodType
}

// GetVoucherQuantityLimitPerUserPeriodTypeOk returns a tuple with the VoucherQuantityLimitPerUserPeriodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherSendRule) GetVoucherQuantityLimitPerUserPeriodTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherQuantityLimitPerUserPeriodType) {
		return nil, false
	}
	return o.VoucherQuantityLimitPerUserPeriodType, true
}

// HasVoucherQuantityLimitPerUserPeriodType returns a boolean if a field has been set.
func (o *PaymentVoucherSendRule) HasVoucherQuantityLimitPerUserPeriodType() bool {
	if o != nil && !IsNil(o.VoucherQuantityLimitPerUserPeriodType) {
		return true
	}

	return false
}

// SetVoucherQuantityLimitPerUserPeriodType gets a reference to the given string and assigns it to the VoucherQuantityLimitPerUserPeriodType field.
func (o *PaymentVoucherSendRule) SetVoucherQuantityLimitPerUserPeriodType(v string) {
	o.VoucherQuantityLimitPerUserPeriodType = &v
}

func (o PaymentVoucherSendRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentVoucherSendRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxQuantityByDay) {
		toSerialize["max_quantity_by_day"] = o.MaxQuantityByDay
	}
	if !IsNil(o.NaturalPersonLimit) {
		toSerialize["natural_person_limit"] = o.NaturalPersonLimit
	}
	if !IsNil(o.PhoneNumberLimit) {
		toSerialize["phone_number_limit"] = o.PhoneNumberLimit
	}
	if !IsNil(o.RealNameLimit) {
		toSerialize["real_name_limit"] = o.RealNameLimit
	}
	if !IsNil(o.VoucherQuantity) {
		toSerialize["voucher_quantity"] = o.VoucherQuantity
	}
	if !IsNil(o.VoucherQuantityLimitPerUser) {
		toSerialize["voucher_quantity_limit_per_user"] = o.VoucherQuantityLimitPerUser
	}
	if !IsNil(o.VoucherQuantityLimitPerUserPeriodType) {
		toSerialize["voucher_quantity_limit_per_user_period_type"] = o.VoucherQuantityLimitPerUserPeriodType
	}
	return toSerialize, nil
}

type NullablePaymentVoucherSendRule struct {
	value *PaymentVoucherSendRule
	isSet bool
}

func (v NullablePaymentVoucherSendRule) Get() *PaymentVoucherSendRule {
	return v.value
}

func (v *NullablePaymentVoucherSendRule) Set(val *PaymentVoucherSendRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentVoucherSendRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentVoucherSendRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentVoucherSendRule(val *PaymentVoucherSendRule) *NullablePaymentVoucherSendRule {
	return &NullablePaymentVoucherSendRule{value: val, isSet: true}
}

func (v NullablePaymentVoucherSendRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentVoucherSendRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

