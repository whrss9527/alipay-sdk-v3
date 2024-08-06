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

// checks if the VoucherPackageModeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherPackageModeInfo{}

// VoucherPackageModeInfo struct for VoucherPackageModeInfo
type VoucherPackageModeInfo struct {
	// 券包id，对应alipay.marketing.activity.voucherpackage.query中voucher_package_id。
	VoucherPackageId *string `json:"voucher_package_id,omitempty"`
}

// NewVoucherPackageModeInfo instantiates a new VoucherPackageModeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherPackageModeInfo() *VoucherPackageModeInfo {
	this := VoucherPackageModeInfo{}
	return &this
}

// NewVoucherPackageModeInfoWithDefaults instantiates a new VoucherPackageModeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherPackageModeInfoWithDefaults() *VoucherPackageModeInfo {
	this := VoucherPackageModeInfo{}
	return &this
}

// GetVoucherPackageId returns the VoucherPackageId field value if set, zero value otherwise.
func (o *VoucherPackageModeInfo) GetVoucherPackageId() string {
	if o == nil || IsNil(o.VoucherPackageId) {
		var ret string
		return ret
	}
	return *o.VoucherPackageId
}

// GetVoucherPackageIdOk returns a tuple with the VoucherPackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherPackageModeInfo) GetVoucherPackageIdOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherPackageId) {
		return nil, false
	}
	return o.VoucherPackageId, true
}

// HasVoucherPackageId returns a boolean if a field has been set.
func (o *VoucherPackageModeInfo) HasVoucherPackageId() bool {
	if o != nil && !IsNil(o.VoucherPackageId) {
		return true
	}

	return false
}

// SetVoucherPackageId gets a reference to the given string and assigns it to the VoucherPackageId field.
func (o *VoucherPackageModeInfo) SetVoucherPackageId(v string) {
	o.VoucherPackageId = &v
}

func (o VoucherPackageModeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherPackageModeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VoucherPackageId) {
		toSerialize["voucher_package_id"] = o.VoucherPackageId
	}
	return toSerialize, nil
}

type NullableVoucherPackageModeInfo struct {
	value *VoucherPackageModeInfo
	isSet bool
}

func (v NullableVoucherPackageModeInfo) Get() *VoucherPackageModeInfo {
	return v.value
}

func (v *NullableVoucherPackageModeInfo) Set(val *VoucherPackageModeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherPackageModeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherPackageModeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherPackageModeInfo(val *VoucherPackageModeInfo) *NullableVoucherPackageModeInfo {
	return &NullableVoucherPackageModeInfo{value: val, isSet: true}
}

func (v NullableVoucherPackageModeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherPackageModeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

