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

// checks if the AlipayOpenSearchSubservicekeywordQuerystatusResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSearchSubservicekeywordQuerystatusResponseModel{}

// AlipayOpenSearchSubservicekeywordQuerystatusResponseModel struct for AlipayOpenSearchSubservicekeywordQuerystatusResponseModel
type AlipayOpenSearchSubservicekeywordQuerystatusResponseModel struct {
	// 关键词工单审核状态返回值
	KeyWords []KeyWordInfo `json:"key_words,omitempty"`
}

// NewAlipayOpenSearchSubservicekeywordQuerystatusResponseModel instantiates a new AlipayOpenSearchSubservicekeywordQuerystatusResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSearchSubservicekeywordQuerystatusResponseModel() *AlipayOpenSearchSubservicekeywordQuerystatusResponseModel {
	this := AlipayOpenSearchSubservicekeywordQuerystatusResponseModel{}
	return &this
}

// NewAlipayOpenSearchSubservicekeywordQuerystatusResponseModelWithDefaults instantiates a new AlipayOpenSearchSubservicekeywordQuerystatusResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSearchSubservicekeywordQuerystatusResponseModelWithDefaults() *AlipayOpenSearchSubservicekeywordQuerystatusResponseModel {
	this := AlipayOpenSearchSubservicekeywordQuerystatusResponseModel{}
	return &this
}

// GetKeyWords returns the KeyWords field value if set, zero value otherwise.
func (o *AlipayOpenSearchSubservicekeywordQuerystatusResponseModel) GetKeyWords() []KeyWordInfo {
	if o == nil || IsNil(o.KeyWords) {
		var ret []KeyWordInfo
		return ret
	}
	return o.KeyWords
}

// GetKeyWordsOk returns a tuple with the KeyWords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchSubservicekeywordQuerystatusResponseModel) GetKeyWordsOk() ([]KeyWordInfo, bool) {
	if o == nil || IsNil(o.KeyWords) {
		return nil, false
	}
	return o.KeyWords, true
}

// HasKeyWords returns a boolean if a field has been set.
func (o *AlipayOpenSearchSubservicekeywordQuerystatusResponseModel) HasKeyWords() bool {
	if o != nil && !IsNil(o.KeyWords) {
		return true
	}

	return false
}

// SetKeyWords gets a reference to the given []KeyWordInfo and assigns it to the KeyWords field.
func (o *AlipayOpenSearchSubservicekeywordQuerystatusResponseModel) SetKeyWords(v []KeyWordInfo) {
	o.KeyWords = v
}

func (o AlipayOpenSearchSubservicekeywordQuerystatusResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSearchSubservicekeywordQuerystatusResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeyWords) {
		toSerialize["key_words"] = o.KeyWords
	}
	return toSerialize, nil
}

type NullableAlipayOpenSearchSubservicekeywordQuerystatusResponseModel struct {
	value *AlipayOpenSearchSubservicekeywordQuerystatusResponseModel
	isSet bool
}

func (v NullableAlipayOpenSearchSubservicekeywordQuerystatusResponseModel) Get() *AlipayOpenSearchSubservicekeywordQuerystatusResponseModel {
	return v.value
}

func (v *NullableAlipayOpenSearchSubservicekeywordQuerystatusResponseModel) Set(val *AlipayOpenSearchSubservicekeywordQuerystatusResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchSubservicekeywordQuerystatusResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchSubservicekeywordQuerystatusResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchSubservicekeywordQuerystatusResponseModel(val *AlipayOpenSearchSubservicekeywordQuerystatusResponseModel) *NullableAlipayOpenSearchSubservicekeywordQuerystatusResponseModel {
	return &NullableAlipayOpenSearchSubservicekeywordQuerystatusResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchSubservicekeywordQuerystatusResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchSubservicekeywordQuerystatusResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
