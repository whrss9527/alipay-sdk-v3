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

// checks if the AlipayOpenSearchAppkeywordquotaQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSearchAppkeywordquotaQueryResponseModel{}

// AlipayOpenSearchAppkeywordquotaQueryResponseModel struct for AlipayOpenSearchAppkeywordquotaQueryResponseModel
type AlipayOpenSearchAppkeywordquotaQueryResponseModel struct {
	// 剩余可配置额度数量，返回具体数字
	QuotaNum *string `json:"quota_num,omitempty"`
}

// NewAlipayOpenSearchAppkeywordquotaQueryResponseModel instantiates a new AlipayOpenSearchAppkeywordquotaQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSearchAppkeywordquotaQueryResponseModel() *AlipayOpenSearchAppkeywordquotaQueryResponseModel {
	this := AlipayOpenSearchAppkeywordquotaQueryResponseModel{}
	return &this
}

// NewAlipayOpenSearchAppkeywordquotaQueryResponseModelWithDefaults instantiates a new AlipayOpenSearchAppkeywordquotaQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSearchAppkeywordquotaQueryResponseModelWithDefaults() *AlipayOpenSearchAppkeywordquotaQueryResponseModel {
	this := AlipayOpenSearchAppkeywordquotaQueryResponseModel{}
	return &this
}

// GetQuotaNum returns the QuotaNum field value if set, zero value otherwise.
func (o *AlipayOpenSearchAppkeywordquotaQueryResponseModel) GetQuotaNum() string {
	if o == nil || IsNil(o.QuotaNum) {
		var ret string
		return ret
	}
	return *o.QuotaNum
}

// GetQuotaNumOk returns a tuple with the QuotaNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchAppkeywordquotaQueryResponseModel) GetQuotaNumOk() (*string, bool) {
	if o == nil || IsNil(o.QuotaNum) {
		return nil, false
	}
	return o.QuotaNum, true
}

// HasQuotaNum returns a boolean if a field has been set.
func (o *AlipayOpenSearchAppkeywordquotaQueryResponseModel) HasQuotaNum() bool {
	if o != nil && !IsNil(o.QuotaNum) {
		return true
	}

	return false
}

// SetQuotaNum gets a reference to the given string and assigns it to the QuotaNum field.
func (o *AlipayOpenSearchAppkeywordquotaQueryResponseModel) SetQuotaNum(v string) {
	o.QuotaNum = &v
}

func (o AlipayOpenSearchAppkeywordquotaQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSearchAppkeywordquotaQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QuotaNum) {
		toSerialize["quota_num"] = o.QuotaNum
	}
	return toSerialize, nil
}

type NullableAlipayOpenSearchAppkeywordquotaQueryResponseModel struct {
	value *AlipayOpenSearchAppkeywordquotaQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenSearchAppkeywordquotaQueryResponseModel) Get() *AlipayOpenSearchAppkeywordquotaQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenSearchAppkeywordquotaQueryResponseModel) Set(val *AlipayOpenSearchAppkeywordquotaQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchAppkeywordquotaQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchAppkeywordquotaQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchAppkeywordquotaQueryResponseModel(val *AlipayOpenSearchAppkeywordquotaQueryResponseModel) *NullableAlipayOpenSearchAppkeywordquotaQueryResponseModel {
	return &NullableAlipayOpenSearchAppkeywordquotaQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchAppkeywordquotaQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchAppkeywordquotaQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

