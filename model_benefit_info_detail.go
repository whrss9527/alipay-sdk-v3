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

// checks if the BenefitInfoDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BenefitInfoDetail{}

// BenefitInfoDetail struct for BenefitInfoDetail
type BenefitInfoDetail struct {
	// PRE_FUND：实际核销或者商户赠送的金额  DISCOUNT：实际折扣掉的金额（获取权益不支持该类型）  COUPON：实际核销或者商户赠送的券
	Amount *string `json:"amount,omitempty"`
	// 权益类型
	BenefitType *string `json:"benefit_type,omitempty"`
	// COUPON：当核销或者赠送券时，可以设置该值
	Count *string `json:"count,omitempty"`
	// 产生核销或者赠送权益的描述
	Description *string `json:"description,omitempty"`
}

// NewBenefitInfoDetail instantiates a new BenefitInfoDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBenefitInfoDetail() *BenefitInfoDetail {
	this := BenefitInfoDetail{}
	return &this
}

// NewBenefitInfoDetailWithDefaults instantiates a new BenefitInfoDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBenefitInfoDetailWithDefaults() *BenefitInfoDetail {
	this := BenefitInfoDetail{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *BenefitInfoDetail) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BenefitInfoDetail) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *BenefitInfoDetail) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *BenefitInfoDetail) SetAmount(v string) {
	o.Amount = &v
}

// GetBenefitType returns the BenefitType field value if set, zero value otherwise.
func (o *BenefitInfoDetail) GetBenefitType() string {
	if o == nil || IsNil(o.BenefitType) {
		var ret string
		return ret
	}
	return *o.BenefitType
}

// GetBenefitTypeOk returns a tuple with the BenefitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BenefitInfoDetail) GetBenefitTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BenefitType) {
		return nil, false
	}
	return o.BenefitType, true
}

// HasBenefitType returns a boolean if a field has been set.
func (o *BenefitInfoDetail) HasBenefitType() bool {
	if o != nil && !IsNil(o.BenefitType) {
		return true
	}

	return false
}

// SetBenefitType gets a reference to the given string and assigns it to the BenefitType field.
func (o *BenefitInfoDetail) SetBenefitType(v string) {
	o.BenefitType = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *BenefitInfoDetail) GetCount() string {
	if o == nil || IsNil(o.Count) {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BenefitInfoDetail) GetCountOk() (*string, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *BenefitInfoDetail) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *BenefitInfoDetail) SetCount(v string) {
	o.Count = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BenefitInfoDetail) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BenefitInfoDetail) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BenefitInfoDetail) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BenefitInfoDetail) SetDescription(v string) {
	o.Description = &v
}

func (o BenefitInfoDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BenefitInfoDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BenefitType) {
		toSerialize["benefit_type"] = o.BenefitType
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableBenefitInfoDetail struct {
	value *BenefitInfoDetail
	isSet bool
}

func (v NullableBenefitInfoDetail) Get() *BenefitInfoDetail {
	return v.value
}

func (v *NullableBenefitInfoDetail) Set(val *BenefitInfoDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableBenefitInfoDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableBenefitInfoDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBenefitInfoDetail(val *BenefitInfoDetail) *NullableBenefitInfoDetail {
	return &NullableBenefitInfoDetail{value: val, isSet: true}
}

func (v NullableBenefitInfoDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBenefitInfoDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
