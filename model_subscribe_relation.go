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

// checks if the SubscribeRelation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscribeRelation{}

// SubscribeRelation struct for SubscribeRelation
type SubscribeRelation struct {
	// 订阅保持状态，即用户勾选“总是保持以上选择，不再询问”选项时勾选的订阅状态。说明：若用户“总是保持以上选择，不再询问”选项，且选择订阅消息。下次触发消息订阅组件时，支付宝将自动发起一次静默订阅（不再拉起订阅组件，无需用户手动订阅）。
	KeepState *string `json:"keep_state,omitempty"`
	// 模板是否展示在订阅组件中
	Show *bool `json:"show,omitempty"`
	// 订阅状态
	SubscribeState *string `json:"subscribe_state,omitempty"`
	// 消息模板的订阅类型
	SubscribeType *string `json:"subscribe_type,omitempty"`
	// 消息模板id
	TemplateId *string `json:"template_id,omitempty"`
}

// NewSubscribeRelation instantiates a new SubscribeRelation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribeRelation() *SubscribeRelation {
	this := SubscribeRelation{}
	return &this
}

// NewSubscribeRelationWithDefaults instantiates a new SubscribeRelation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribeRelationWithDefaults() *SubscribeRelation {
	this := SubscribeRelation{}
	return &this
}

// GetKeepState returns the KeepState field value if set, zero value otherwise.
func (o *SubscribeRelation) GetKeepState() string {
	if o == nil || IsNil(o.KeepState) {
		var ret string
		return ret
	}
	return *o.KeepState
}

// GetKeepStateOk returns a tuple with the KeepState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeRelation) GetKeepStateOk() (*string, bool) {
	if o == nil || IsNil(o.KeepState) {
		return nil, false
	}
	return o.KeepState, true
}

// HasKeepState returns a boolean if a field has been set.
func (o *SubscribeRelation) HasKeepState() bool {
	if o != nil && !IsNil(o.KeepState) {
		return true
	}

	return false
}

// SetKeepState gets a reference to the given string and assigns it to the KeepState field.
func (o *SubscribeRelation) SetKeepState(v string) {
	o.KeepState = &v
}

// GetShow returns the Show field value if set, zero value otherwise.
func (o *SubscribeRelation) GetShow() bool {
	if o == nil || IsNil(o.Show) {
		var ret bool
		return ret
	}
	return *o.Show
}

// GetShowOk returns a tuple with the Show field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeRelation) GetShowOk() (*bool, bool) {
	if o == nil || IsNil(o.Show) {
		return nil, false
	}
	return o.Show, true
}

// HasShow returns a boolean if a field has been set.
func (o *SubscribeRelation) HasShow() bool {
	if o != nil && !IsNil(o.Show) {
		return true
	}

	return false
}

// SetShow gets a reference to the given bool and assigns it to the Show field.
func (o *SubscribeRelation) SetShow(v bool) {
	o.Show = &v
}

// GetSubscribeState returns the SubscribeState field value if set, zero value otherwise.
func (o *SubscribeRelation) GetSubscribeState() string {
	if o == nil || IsNil(o.SubscribeState) {
		var ret string
		return ret
	}
	return *o.SubscribeState
}

// GetSubscribeStateOk returns a tuple with the SubscribeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeRelation) GetSubscribeStateOk() (*string, bool) {
	if o == nil || IsNil(o.SubscribeState) {
		return nil, false
	}
	return o.SubscribeState, true
}

// HasSubscribeState returns a boolean if a field has been set.
func (o *SubscribeRelation) HasSubscribeState() bool {
	if o != nil && !IsNil(o.SubscribeState) {
		return true
	}

	return false
}

// SetSubscribeState gets a reference to the given string and assigns it to the SubscribeState field.
func (o *SubscribeRelation) SetSubscribeState(v string) {
	o.SubscribeState = &v
}

// GetSubscribeType returns the SubscribeType field value if set, zero value otherwise.
func (o *SubscribeRelation) GetSubscribeType() string {
	if o == nil || IsNil(o.SubscribeType) {
		var ret string
		return ret
	}
	return *o.SubscribeType
}

// GetSubscribeTypeOk returns a tuple with the SubscribeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeRelation) GetSubscribeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubscribeType) {
		return nil, false
	}
	return o.SubscribeType, true
}

// HasSubscribeType returns a boolean if a field has been set.
func (o *SubscribeRelation) HasSubscribeType() bool {
	if o != nil && !IsNil(o.SubscribeType) {
		return true
	}

	return false
}

// SetSubscribeType gets a reference to the given string and assigns it to the SubscribeType field.
func (o *SubscribeRelation) SetSubscribeType(v string) {
	o.SubscribeType = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *SubscribeRelation) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribeRelation) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *SubscribeRelation) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *SubscribeRelation) SetTemplateId(v string) {
	o.TemplateId = &v
}

func (o SubscribeRelation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscribeRelation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeepState) {
		toSerialize["keep_state"] = o.KeepState
	}
	if !IsNil(o.Show) {
		toSerialize["show"] = o.Show
	}
	if !IsNil(o.SubscribeState) {
		toSerialize["subscribe_state"] = o.SubscribeState
	}
	if !IsNil(o.SubscribeType) {
		toSerialize["subscribe_type"] = o.SubscribeType
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	return toSerialize, nil
}

type NullableSubscribeRelation struct {
	value *SubscribeRelation
	isSet bool
}

func (v NullableSubscribeRelation) Get() *SubscribeRelation {
	return v.value
}

func (v *NullableSubscribeRelation) Set(val *SubscribeRelation) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribeRelation) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribeRelation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribeRelation(val *SubscribeRelation) *NullableSubscribeRelation {
	return &NullableSubscribeRelation{value: val, isSet: true}
}

func (v NullableSubscribeRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribeRelation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
