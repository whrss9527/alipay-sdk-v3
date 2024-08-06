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

// checks if the AlipayOpenAgentCommonsignConfirmResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenAgentCommonsignConfirmResponseModel{}

// AlipayOpenAgentCommonsignConfirmResponseModel struct for AlipayOpenAgentCommonsignConfirmResponseModel
type AlipayOpenAgentCommonsignConfirmResponseModel struct {
	// 签约单号
	OrderNo *string `json:"order_no,omitempty"`
}

// NewAlipayOpenAgentCommonsignConfirmResponseModel instantiates a new AlipayOpenAgentCommonsignConfirmResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenAgentCommonsignConfirmResponseModel() *AlipayOpenAgentCommonsignConfirmResponseModel {
	this := AlipayOpenAgentCommonsignConfirmResponseModel{}
	return &this
}

// NewAlipayOpenAgentCommonsignConfirmResponseModelWithDefaults instantiates a new AlipayOpenAgentCommonsignConfirmResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenAgentCommonsignConfirmResponseModelWithDefaults() *AlipayOpenAgentCommonsignConfirmResponseModel {
	this := AlipayOpenAgentCommonsignConfirmResponseModel{}
	return &this
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *AlipayOpenAgentCommonsignConfirmResponseModel) GetOrderNo() string {
	if o == nil || IsNil(o.OrderNo) {
		var ret string
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAgentCommonsignConfirmResponseModel) GetOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *AlipayOpenAgentCommonsignConfirmResponseModel) HasOrderNo() bool {
	if o != nil && !IsNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given string and assigns it to the OrderNo field.
func (o *AlipayOpenAgentCommonsignConfirmResponseModel) SetOrderNo(v string) {
	o.OrderNo = &v
}

func (o AlipayOpenAgentCommonsignConfirmResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenAgentCommonsignConfirmResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderNo) {
		toSerialize["order_no"] = o.OrderNo
	}
	return toSerialize, nil
}

type NullableAlipayOpenAgentCommonsignConfirmResponseModel struct {
	value *AlipayOpenAgentCommonsignConfirmResponseModel
	isSet bool
}

func (v NullableAlipayOpenAgentCommonsignConfirmResponseModel) Get() *AlipayOpenAgentCommonsignConfirmResponseModel {
	return v.value
}

func (v *NullableAlipayOpenAgentCommonsignConfirmResponseModel) Set(val *AlipayOpenAgentCommonsignConfirmResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAgentCommonsignConfirmResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAgentCommonsignConfirmResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAgentCommonsignConfirmResponseModel(val *AlipayOpenAgentCommonsignConfirmResponseModel) *NullableAlipayOpenAgentCommonsignConfirmResponseModel {
	return &NullableAlipayOpenAgentCommonsignConfirmResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenAgentCommonsignConfirmResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAgentCommonsignConfirmResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

