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

// checks if the AlipayOpenPublicAccountQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicAccountQueryResponseModel{}

// AlipayOpenPublicAccountQueryResponseModel struct for AlipayOpenPublicAccountQueryResponseModel
type AlipayOpenPublicAccountQueryResponseModel struct {
	// 绑定账户列表
	PublicBindAccounts []StdPublicBindAccount `json:"public_bind_accounts,omitempty"`
}

// NewAlipayOpenPublicAccountQueryResponseModel instantiates a new AlipayOpenPublicAccountQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicAccountQueryResponseModel() *AlipayOpenPublicAccountQueryResponseModel {
	this := AlipayOpenPublicAccountQueryResponseModel{}
	return &this
}

// NewAlipayOpenPublicAccountQueryResponseModelWithDefaults instantiates a new AlipayOpenPublicAccountQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicAccountQueryResponseModelWithDefaults() *AlipayOpenPublicAccountQueryResponseModel {
	this := AlipayOpenPublicAccountQueryResponseModel{}
	return &this
}

// GetPublicBindAccounts returns the PublicBindAccounts field value if set, zero value otherwise.
func (o *AlipayOpenPublicAccountQueryResponseModel) GetPublicBindAccounts() []StdPublicBindAccount {
	if o == nil || IsNil(o.PublicBindAccounts) {
		var ret []StdPublicBindAccount
		return ret
	}
	return o.PublicBindAccounts
}

// GetPublicBindAccountsOk returns a tuple with the PublicBindAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicAccountQueryResponseModel) GetPublicBindAccountsOk() ([]StdPublicBindAccount, bool) {
	if o == nil || IsNil(o.PublicBindAccounts) {
		return nil, false
	}
	return o.PublicBindAccounts, true
}

// HasPublicBindAccounts returns a boolean if a field has been set.
func (o *AlipayOpenPublicAccountQueryResponseModel) HasPublicBindAccounts() bool {
	if o != nil && !IsNil(o.PublicBindAccounts) {
		return true
	}

	return false
}

// SetPublicBindAccounts gets a reference to the given []StdPublicBindAccount and assigns it to the PublicBindAccounts field.
func (o *AlipayOpenPublicAccountQueryResponseModel) SetPublicBindAccounts(v []StdPublicBindAccount) {
	o.PublicBindAccounts = v
}

func (o AlipayOpenPublicAccountQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicAccountQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PublicBindAccounts) {
		toSerialize["public_bind_accounts"] = o.PublicBindAccounts
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicAccountQueryResponseModel struct {
	value *AlipayOpenPublicAccountQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenPublicAccountQueryResponseModel) Get() *AlipayOpenPublicAccountQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenPublicAccountQueryResponseModel) Set(val *AlipayOpenPublicAccountQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicAccountQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicAccountQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicAccountQueryResponseModel(val *AlipayOpenPublicAccountQueryResponseModel) *NullableAlipayOpenPublicAccountQueryResponseModel {
	return &NullableAlipayOpenPublicAccountQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicAccountQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicAccountQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

