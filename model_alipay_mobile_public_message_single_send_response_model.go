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

// checks if the AlipayMobilePublicMessageSingleSendResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMobilePublicMessageSingleSendResponseModel{}

// AlipayMobilePublicMessageSingleSendResponseModel struct for AlipayMobilePublicMessageSingleSendResponseModel
type AlipayMobilePublicMessageSingleSendResponseModel struct {
	// 结果码
	Code *string `json:"code,omitempty"`
	// 结果描述
	Msg *string `json:"msg,omitempty"`
}

// NewAlipayMobilePublicMessageSingleSendResponseModel instantiates a new AlipayMobilePublicMessageSingleSendResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMobilePublicMessageSingleSendResponseModel() *AlipayMobilePublicMessageSingleSendResponseModel {
	this := AlipayMobilePublicMessageSingleSendResponseModel{}
	return &this
}

// NewAlipayMobilePublicMessageSingleSendResponseModelWithDefaults instantiates a new AlipayMobilePublicMessageSingleSendResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMobilePublicMessageSingleSendResponseModelWithDefaults() *AlipayMobilePublicMessageSingleSendResponseModel {
	this := AlipayMobilePublicMessageSingleSendResponseModel{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AlipayMobilePublicMessageSingleSendResponseModel) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMobilePublicMessageSingleSendResponseModel) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AlipayMobilePublicMessageSingleSendResponseModel) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AlipayMobilePublicMessageSingleSendResponseModel) SetCode(v string) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *AlipayMobilePublicMessageSingleSendResponseModel) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMobilePublicMessageSingleSendResponseModel) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *AlipayMobilePublicMessageSingleSendResponseModel) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *AlipayMobilePublicMessageSingleSendResponseModel) SetMsg(v string) {
	o.Msg = &v
}

func (o AlipayMobilePublicMessageSingleSendResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMobilePublicMessageSingleSendResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	return toSerialize, nil
}

type NullableAlipayMobilePublicMessageSingleSendResponseModel struct {
	value *AlipayMobilePublicMessageSingleSendResponseModel
	isSet bool
}

func (v NullableAlipayMobilePublicMessageSingleSendResponseModel) Get() *AlipayMobilePublicMessageSingleSendResponseModel {
	return v.value
}

func (v *NullableAlipayMobilePublicMessageSingleSendResponseModel) Set(val *AlipayMobilePublicMessageSingleSendResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMobilePublicMessageSingleSendResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMobilePublicMessageSingleSendResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMobilePublicMessageSingleSendResponseModel(val *AlipayMobilePublicMessageSingleSendResponseModel) *NullableAlipayMobilePublicMessageSingleSendResponseModel {
	return &NullableAlipayMobilePublicMessageSingleSendResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMobilePublicMessageSingleSendResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMobilePublicMessageSingleSendResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
