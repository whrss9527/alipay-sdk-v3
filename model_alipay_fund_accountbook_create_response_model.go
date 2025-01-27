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

// checks if the AlipayFundAccountbookCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundAccountbookCreateResponseModel{}

// AlipayFundAccountbookCreateResponseModel struct for AlipayFundAccountbookCreateResponseModel
type AlipayFundAccountbookCreateResponseModel struct {
	// 开通的资金记账本id
	AccountBookId *string      `json:"account_book_id,omitempty"`
	ExtCardInfo   *ExtCardInfo `json:"ext_card_info,omitempty"`
}

// NewAlipayFundAccountbookCreateResponseModel instantiates a new AlipayFundAccountbookCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundAccountbookCreateResponseModel() *AlipayFundAccountbookCreateResponseModel {
	this := AlipayFundAccountbookCreateResponseModel{}
	return &this
}

// NewAlipayFundAccountbookCreateResponseModelWithDefaults instantiates a new AlipayFundAccountbookCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundAccountbookCreateResponseModelWithDefaults() *AlipayFundAccountbookCreateResponseModel {
	this := AlipayFundAccountbookCreateResponseModel{}
	return &this
}

// GetAccountBookId returns the AccountBookId field value if set, zero value otherwise.
func (o *AlipayFundAccountbookCreateResponseModel) GetAccountBookId() string {
	if o == nil || IsNil(o.AccountBookId) {
		var ret string
		return ret
	}
	return *o.AccountBookId
}

// GetAccountBookIdOk returns a tuple with the AccountBookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAccountbookCreateResponseModel) GetAccountBookIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountBookId) {
		return nil, false
	}
	return o.AccountBookId, true
}

// HasAccountBookId returns a boolean if a field has been set.
func (o *AlipayFundAccountbookCreateResponseModel) HasAccountBookId() bool {
	if o != nil && !IsNil(o.AccountBookId) {
		return true
	}

	return false
}

// SetAccountBookId gets a reference to the given string and assigns it to the AccountBookId field.
func (o *AlipayFundAccountbookCreateResponseModel) SetAccountBookId(v string) {
	o.AccountBookId = &v
}

// GetExtCardInfo returns the ExtCardInfo field value if set, zero value otherwise.
func (o *AlipayFundAccountbookCreateResponseModel) GetExtCardInfo() ExtCardInfo {
	if o == nil || IsNil(o.ExtCardInfo) {
		var ret ExtCardInfo
		return ret
	}
	return *o.ExtCardInfo
}

// GetExtCardInfoOk returns a tuple with the ExtCardInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAccountbookCreateResponseModel) GetExtCardInfoOk() (*ExtCardInfo, bool) {
	if o == nil || IsNil(o.ExtCardInfo) {
		return nil, false
	}
	return o.ExtCardInfo, true
}

// HasExtCardInfo returns a boolean if a field has been set.
func (o *AlipayFundAccountbookCreateResponseModel) HasExtCardInfo() bool {
	if o != nil && !IsNil(o.ExtCardInfo) {
		return true
	}

	return false
}

// SetExtCardInfo gets a reference to the given ExtCardInfo and assigns it to the ExtCardInfo field.
func (o *AlipayFundAccountbookCreateResponseModel) SetExtCardInfo(v ExtCardInfo) {
	o.ExtCardInfo = &v
}

func (o AlipayFundAccountbookCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundAccountbookCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountBookId) {
		toSerialize["account_book_id"] = o.AccountBookId
	}
	if !IsNil(o.ExtCardInfo) {
		toSerialize["ext_card_info"] = o.ExtCardInfo
	}
	return toSerialize, nil
}

type NullableAlipayFundAccountbookCreateResponseModel struct {
	value *AlipayFundAccountbookCreateResponseModel
	isSet bool
}

func (v NullableAlipayFundAccountbookCreateResponseModel) Get() *AlipayFundAccountbookCreateResponseModel {
	return v.value
}

func (v *NullableAlipayFundAccountbookCreateResponseModel) Set(val *AlipayFundAccountbookCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundAccountbookCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundAccountbookCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundAccountbookCreateResponseModel(val *AlipayFundAccountbookCreateResponseModel) *NullableAlipayFundAccountbookCreateResponseModel {
	return &NullableAlipayFundAccountbookCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayFundAccountbookCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundAccountbookCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
