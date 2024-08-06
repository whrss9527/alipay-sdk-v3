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

// checks if the VoucherAvailableAppInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherAvailableAppInfo{}

// VoucherAvailableAppInfo struct for VoucherAvailableAppInfo
type VoucherAvailableAppInfo struct {
	// 可核销的支付宝小程序id
	AvailableAppIds []string `json:"available_app_ids,omitempty"`
}

// NewVoucherAvailableAppInfo instantiates a new VoucherAvailableAppInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherAvailableAppInfo() *VoucherAvailableAppInfo {
	this := VoucherAvailableAppInfo{}
	return &this
}

// NewVoucherAvailableAppInfoWithDefaults instantiates a new VoucherAvailableAppInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherAvailableAppInfoWithDefaults() *VoucherAvailableAppInfo {
	this := VoucherAvailableAppInfo{}
	return &this
}

// GetAvailableAppIds returns the AvailableAppIds field value if set, zero value otherwise.
func (o *VoucherAvailableAppInfo) GetAvailableAppIds() []string {
	if o == nil || IsNil(o.AvailableAppIds) {
		var ret []string
		return ret
	}
	return o.AvailableAppIds
}

// GetAvailableAppIdsOk returns a tuple with the AvailableAppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableAppInfo) GetAvailableAppIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailableAppIds) {
		return nil, false
	}
	return o.AvailableAppIds, true
}

// HasAvailableAppIds returns a boolean if a field has been set.
func (o *VoucherAvailableAppInfo) HasAvailableAppIds() bool {
	if o != nil && !IsNil(o.AvailableAppIds) {
		return true
	}

	return false
}

// SetAvailableAppIds gets a reference to the given []string and assigns it to the AvailableAppIds field.
func (o *VoucherAvailableAppInfo) SetAvailableAppIds(v []string) {
	o.AvailableAppIds = v
}

func (o VoucherAvailableAppInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherAvailableAppInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailableAppIds) {
		toSerialize["available_app_ids"] = o.AvailableAppIds
	}
	return toSerialize, nil
}

type NullableVoucherAvailableAppInfo struct {
	value *VoucherAvailableAppInfo
	isSet bool
}

func (v NullableVoucherAvailableAppInfo) Get() *VoucherAvailableAppInfo {
	return v.value
}

func (v *NullableVoucherAvailableAppInfo) Set(val *VoucherAvailableAppInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherAvailableAppInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherAvailableAppInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherAvailableAppInfo(val *VoucherAvailableAppInfo) *NullableVoucherAvailableAppInfo {
	return &NullableVoucherAvailableAppInfo{value: val, isSet: true}
}

func (v NullableVoucherAvailableAppInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherAvailableAppInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

