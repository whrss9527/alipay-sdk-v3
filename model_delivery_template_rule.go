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

// checks if the DeliveryTemplateRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryTemplateRule{}

// DeliveryTemplateRule struct for DeliveryTemplateRule
type DeliveryTemplateRule struct {
	// 指定商家消息区域
	TemplateId *string `json:"template_id,omitempty"`
}

// NewDeliveryTemplateRule instantiates a new DeliveryTemplateRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryTemplateRule() *DeliveryTemplateRule {
	this := DeliveryTemplateRule{}
	return &this
}

// NewDeliveryTemplateRuleWithDefaults instantiates a new DeliveryTemplateRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryTemplateRuleWithDefaults() *DeliveryTemplateRule {
	this := DeliveryTemplateRule{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *DeliveryTemplateRule) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryTemplateRule) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *DeliveryTemplateRule) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *DeliveryTemplateRule) SetTemplateId(v string) {
	o.TemplateId = &v
}

func (o DeliveryTemplateRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryTemplateRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	return toSerialize, nil
}

type NullableDeliveryTemplateRule struct {
	value *DeliveryTemplateRule
	isSet bool
}

func (v NullableDeliveryTemplateRule) Get() *DeliveryTemplateRule {
	return v.value
}

func (v *NullableDeliveryTemplateRule) Set(val *DeliveryTemplateRule) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryTemplateRule) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryTemplateRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryTemplateRule(val *DeliveryTemplateRule) *NullableDeliveryTemplateRule {
	return &NullableDeliveryTemplateRule{value: val, isSet: true}
}

func (v NullableDeliveryTemplateRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryTemplateRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
