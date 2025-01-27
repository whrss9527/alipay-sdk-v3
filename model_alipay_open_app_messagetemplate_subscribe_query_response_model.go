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

// checks if the AlipayOpenAppMessagetemplateSubscribeQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenAppMessagetemplateSubscribeQueryResponseModel{}

// AlipayOpenAppMessagetemplateSubscribeQueryResponseModel struct for AlipayOpenAppMessagetemplateSubscribeQueryResponseModel
type AlipayOpenAppMessagetemplateSubscribeQueryResponseModel struct {
	// 是否显示订阅组件
	ShowComponent *bool `json:"show_component,omitempty"`
	// 用户对消息模板的订阅关系列表，为入参中的用户id对消息模板id的订阅关系。 限制：用户未订阅消息，该参数不返回。
	SubscribeRelations []SubscribeRelation `json:"subscribe_relations,omitempty"`
}

// NewAlipayOpenAppMessagetemplateSubscribeQueryResponseModel instantiates a new AlipayOpenAppMessagetemplateSubscribeQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenAppMessagetemplateSubscribeQueryResponseModel() *AlipayOpenAppMessagetemplateSubscribeQueryResponseModel {
	this := AlipayOpenAppMessagetemplateSubscribeQueryResponseModel{}
	return &this
}

// NewAlipayOpenAppMessagetemplateSubscribeQueryResponseModelWithDefaults instantiates a new AlipayOpenAppMessagetemplateSubscribeQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenAppMessagetemplateSubscribeQueryResponseModelWithDefaults() *AlipayOpenAppMessagetemplateSubscribeQueryResponseModel {
	this := AlipayOpenAppMessagetemplateSubscribeQueryResponseModel{}
	return &this
}

// GetShowComponent returns the ShowComponent field value if set, zero value otherwise.
func (o *AlipayOpenAppMessagetemplateSubscribeQueryResponseModel) GetShowComponent() bool {
	if o == nil || IsNil(o.ShowComponent) {
		var ret bool
		return ret
	}
	return *o.ShowComponent
}

// GetShowComponentOk returns a tuple with the ShowComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppMessagetemplateSubscribeQueryResponseModel) GetShowComponentOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowComponent) {
		return nil, false
	}
	return o.ShowComponent, true
}

// HasShowComponent returns a boolean if a field has been set.
func (o *AlipayOpenAppMessagetemplateSubscribeQueryResponseModel) HasShowComponent() bool {
	if o != nil && !IsNil(o.ShowComponent) {
		return true
	}

	return false
}

// SetShowComponent gets a reference to the given bool and assigns it to the ShowComponent field.
func (o *AlipayOpenAppMessagetemplateSubscribeQueryResponseModel) SetShowComponent(v bool) {
	o.ShowComponent = &v
}

// GetSubscribeRelations returns the SubscribeRelations field value if set, zero value otherwise.
func (o *AlipayOpenAppMessagetemplateSubscribeQueryResponseModel) GetSubscribeRelations() []SubscribeRelation {
	if o == nil || IsNil(o.SubscribeRelations) {
		var ret []SubscribeRelation
		return ret
	}
	return o.SubscribeRelations
}

// GetSubscribeRelationsOk returns a tuple with the SubscribeRelations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppMessagetemplateSubscribeQueryResponseModel) GetSubscribeRelationsOk() ([]SubscribeRelation, bool) {
	if o == nil || IsNil(o.SubscribeRelations) {
		return nil, false
	}
	return o.SubscribeRelations, true
}

// HasSubscribeRelations returns a boolean if a field has been set.
func (o *AlipayOpenAppMessagetemplateSubscribeQueryResponseModel) HasSubscribeRelations() bool {
	if o != nil && !IsNil(o.SubscribeRelations) {
		return true
	}

	return false
}

// SetSubscribeRelations gets a reference to the given []SubscribeRelation and assigns it to the SubscribeRelations field.
func (o *AlipayOpenAppMessagetemplateSubscribeQueryResponseModel) SetSubscribeRelations(v []SubscribeRelation) {
	o.SubscribeRelations = v
}

func (o AlipayOpenAppMessagetemplateSubscribeQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenAppMessagetemplateSubscribeQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShowComponent) {
		toSerialize["show_component"] = o.ShowComponent
	}
	if !IsNil(o.SubscribeRelations) {
		toSerialize["subscribe_relations"] = o.SubscribeRelations
	}
	return toSerialize, nil
}

type NullableAlipayOpenAppMessagetemplateSubscribeQueryResponseModel struct {
	value *AlipayOpenAppMessagetemplateSubscribeQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenAppMessagetemplateSubscribeQueryResponseModel) Get() *AlipayOpenAppMessagetemplateSubscribeQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenAppMessagetemplateSubscribeQueryResponseModel) Set(val *AlipayOpenAppMessagetemplateSubscribeQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAppMessagetemplateSubscribeQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAppMessagetemplateSubscribeQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAppMessagetemplateSubscribeQueryResponseModel(val *AlipayOpenAppMessagetemplateSubscribeQueryResponseModel) *NullableAlipayOpenAppMessagetemplateSubscribeQueryResponseModel {
	return &NullableAlipayOpenAppMessagetemplateSubscribeQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenAppMessagetemplateSubscribeQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAppMessagetemplateSubscribeQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
