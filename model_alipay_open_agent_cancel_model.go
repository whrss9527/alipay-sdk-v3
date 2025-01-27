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

// checks if the AlipayOpenAgentCancelModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenAgentCancelModel{}

// AlipayOpenAgentCancelModel struct for AlipayOpenAgentCancelModel
type AlipayOpenAgentCancelModel struct {
	// ISV 代商户操作事务编号，通过事务开启接口alipay.open.agent.create调用返回。
	BatchNo *string `json:"batch_no,omitempty"`
}

// NewAlipayOpenAgentCancelModel instantiates a new AlipayOpenAgentCancelModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenAgentCancelModel() *AlipayOpenAgentCancelModel {
	this := AlipayOpenAgentCancelModel{}
	return &this
}

// NewAlipayOpenAgentCancelModelWithDefaults instantiates a new AlipayOpenAgentCancelModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenAgentCancelModelWithDefaults() *AlipayOpenAgentCancelModel {
	this := AlipayOpenAgentCancelModel{}
	return &this
}

// GetBatchNo returns the BatchNo field value if set, zero value otherwise.
func (o *AlipayOpenAgentCancelModel) GetBatchNo() string {
	if o == nil || IsNil(o.BatchNo) {
		var ret string
		return ret
	}
	return *o.BatchNo
}

// GetBatchNoOk returns a tuple with the BatchNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentCancelModel) GetBatchNoOk() (*string, bool) {
	if o == nil || IsNil(o.BatchNo) {
		return nil, false
	}
	return o.BatchNo, true
}

// HasBatchNo returns a boolean if a field has been set.
func (o *AlipayOpenAgentCancelModel) HasBatchNo() bool {
	if o != nil && !IsNil(o.BatchNo) {
		return true
	}

	return false
}

// SetBatchNo gets a reference to the given string and assigns it to the BatchNo field.
func (o *AlipayOpenAgentCancelModel) SetBatchNo(v string) {
	o.BatchNo = &v
}

func (o AlipayOpenAgentCancelModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenAgentCancelModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BatchNo) {
		toSerialize["batch_no"] = o.BatchNo
	}
	return toSerialize, nil
}

type NullableAlipayOpenAgentCancelModel struct {
	value *AlipayOpenAgentCancelModel
	isSet bool
}

func (v NullableAlipayOpenAgentCancelModel) Get() *AlipayOpenAgentCancelModel {
	return v.value
}

func (v *NullableAlipayOpenAgentCancelModel) Set(val *AlipayOpenAgentCancelModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAgentCancelModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAgentCancelModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAgentCancelModel(val *AlipayOpenAgentCancelModel) *NullableAlipayOpenAgentCancelModel {
	return &NullableAlipayOpenAgentCancelModel{value: val, isSet: true}
}

func (v NullableAlipayOpenAgentCancelModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAgentCancelModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
