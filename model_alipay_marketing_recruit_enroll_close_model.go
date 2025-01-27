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

// checks if the AlipayMarketingRecruitEnrollCloseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingRecruitEnrollCloseModel{}

// AlipayMarketingRecruitEnrollCloseModel struct for AlipayMarketingRecruitEnrollCloseModel
type AlipayMarketingRecruitEnrollCloseModel struct {
	// 活动报名ID
	EnrollId *string `json:"enroll_id,omitempty"`
}

// NewAlipayMarketingRecruitEnrollCloseModel instantiates a new AlipayMarketingRecruitEnrollCloseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingRecruitEnrollCloseModel() *AlipayMarketingRecruitEnrollCloseModel {
	this := AlipayMarketingRecruitEnrollCloseModel{}
	return &this
}

// NewAlipayMarketingRecruitEnrollCloseModelWithDefaults instantiates a new AlipayMarketingRecruitEnrollCloseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingRecruitEnrollCloseModelWithDefaults() *AlipayMarketingRecruitEnrollCloseModel {
	this := AlipayMarketingRecruitEnrollCloseModel{}
	return &this
}

// GetEnrollId returns the EnrollId field value if set, zero value otherwise.
func (o *AlipayMarketingRecruitEnrollCloseModel) GetEnrollId() string {
	if o == nil || IsNil(o.EnrollId) {
		var ret string
		return ret
	}
	return *o.EnrollId
}

// GetEnrollIdOk returns a tuple with the EnrollId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingRecruitEnrollCloseModel) GetEnrollIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnrollId) {
		return nil, false
	}
	return o.EnrollId, true
}

// HasEnrollId returns a boolean if a field has been set.
func (o *AlipayMarketingRecruitEnrollCloseModel) HasEnrollId() bool {
	if o != nil && !IsNil(o.EnrollId) {
		return true
	}

	return false
}

// SetEnrollId gets a reference to the given string and assigns it to the EnrollId field.
func (o *AlipayMarketingRecruitEnrollCloseModel) SetEnrollId(v string) {
	o.EnrollId = &v
}

func (o AlipayMarketingRecruitEnrollCloseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingRecruitEnrollCloseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnrollId) {
		toSerialize["enroll_id"] = o.EnrollId
	}
	return toSerialize, nil
}

type NullableAlipayMarketingRecruitEnrollCloseModel struct {
	value *AlipayMarketingRecruitEnrollCloseModel
	isSet bool
}

func (v NullableAlipayMarketingRecruitEnrollCloseModel) Get() *AlipayMarketingRecruitEnrollCloseModel {
	return v.value
}

func (v *NullableAlipayMarketingRecruitEnrollCloseModel) Set(val *AlipayMarketingRecruitEnrollCloseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingRecruitEnrollCloseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingRecruitEnrollCloseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingRecruitEnrollCloseModel(val *AlipayMarketingRecruitEnrollCloseModel) *NullableAlipayMarketingRecruitEnrollCloseModel {
	return &NullableAlipayMarketingRecruitEnrollCloseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingRecruitEnrollCloseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingRecruitEnrollCloseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
