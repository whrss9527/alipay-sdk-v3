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

// checks if the AlipayMobilePublicTemplateMessageDeleteModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMobilePublicTemplateMessageDeleteModel{}

// AlipayMobilePublicTemplateMessageDeleteModel struct for AlipayMobilePublicTemplateMessageDeleteModel
type AlipayMobilePublicTemplateMessageDeleteModel struct {
	// 模板id
	TemplateId *string `json:"template_id,omitempty"`
}

// NewAlipayMobilePublicTemplateMessageDeleteModel instantiates a new AlipayMobilePublicTemplateMessageDeleteModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMobilePublicTemplateMessageDeleteModel() *AlipayMobilePublicTemplateMessageDeleteModel {
	this := AlipayMobilePublicTemplateMessageDeleteModel{}
	return &this
}

// NewAlipayMobilePublicTemplateMessageDeleteModelWithDefaults instantiates a new AlipayMobilePublicTemplateMessageDeleteModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMobilePublicTemplateMessageDeleteModelWithDefaults() *AlipayMobilePublicTemplateMessageDeleteModel {
	this := AlipayMobilePublicTemplateMessageDeleteModel{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *AlipayMobilePublicTemplateMessageDeleteModel) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMobilePublicTemplateMessageDeleteModel) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *AlipayMobilePublicTemplateMessageDeleteModel) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *AlipayMobilePublicTemplateMessageDeleteModel) SetTemplateId(v string) {
	o.TemplateId = &v
}

func (o AlipayMobilePublicTemplateMessageDeleteModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMobilePublicTemplateMessageDeleteModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	return toSerialize, nil
}

type NullableAlipayMobilePublicTemplateMessageDeleteModel struct {
	value *AlipayMobilePublicTemplateMessageDeleteModel
	isSet bool
}

func (v NullableAlipayMobilePublicTemplateMessageDeleteModel) Get() *AlipayMobilePublicTemplateMessageDeleteModel {
	return v.value
}

func (v *NullableAlipayMobilePublicTemplateMessageDeleteModel) Set(val *AlipayMobilePublicTemplateMessageDeleteModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMobilePublicTemplateMessageDeleteModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMobilePublicTemplateMessageDeleteModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMobilePublicTemplateMessageDeleteModel(val *AlipayMobilePublicTemplateMessageDeleteModel) *NullableAlipayMobilePublicTemplateMessageDeleteModel {
	return &NullableAlipayMobilePublicTemplateMessageDeleteModel{value: val, isSet: true}
}

func (v NullableAlipayMobilePublicTemplateMessageDeleteModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMobilePublicTemplateMessageDeleteModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

