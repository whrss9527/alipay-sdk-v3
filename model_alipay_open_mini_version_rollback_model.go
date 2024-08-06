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

// checks if the AlipayOpenMiniVersionRollbackModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniVersionRollbackModel{}

// AlipayOpenMiniVersionRollbackModel struct for AlipayOpenMiniVersionRollbackModel
type AlipayOpenMiniVersionRollbackModel struct {
	// 商家小程序已上架版本。下架后将自动回滚至上一（已上架）小程序版本。 例如：商家小程序有 0.01（上一上架版本）、0.02（未上架版本）、0.03（当前上架版本） 三个版本，回滚时需传入商家小程序版本 0.03 表示将 0.03 版本回滚至上一个已上架版本即此处 0.01 版本，同时 0.03 版本将自动下架，0.01 版本自动上架（无需审核）。
	AppVersion *string `json:"app_version,omitempty"`
	// 小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。
	BundleId *string `json:"bundle_id,omitempty"`
}

// NewAlipayOpenMiniVersionRollbackModel instantiates a new AlipayOpenMiniVersionRollbackModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniVersionRollbackModel() *AlipayOpenMiniVersionRollbackModel {
	this := AlipayOpenMiniVersionRollbackModel{}
	return &this
}

// NewAlipayOpenMiniVersionRollbackModelWithDefaults instantiates a new AlipayOpenMiniVersionRollbackModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniVersionRollbackModelWithDefaults() *AlipayOpenMiniVersionRollbackModel {
	this := AlipayOpenMiniVersionRollbackModel{}
	return &this
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *AlipayOpenMiniVersionRollbackModel) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion) {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniVersionRollbackModel) GetAppVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *AlipayOpenMiniVersionRollbackModel) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *AlipayOpenMiniVersionRollbackModel) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *AlipayOpenMiniVersionRollbackModel) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniVersionRollbackModel) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *AlipayOpenMiniVersionRollbackModel) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *AlipayOpenMiniVersionRollbackModel) SetBundleId(v string) {
	o.BundleId = &v
}

func (o AlipayOpenMiniVersionRollbackModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniVersionRollbackModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppVersion) {
		toSerialize["app_version"] = o.AppVersion
	}
	if !IsNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniVersionRollbackModel struct {
	value *AlipayOpenMiniVersionRollbackModel
	isSet bool
}

func (v NullableAlipayOpenMiniVersionRollbackModel) Get() *AlipayOpenMiniVersionRollbackModel {
	return v.value
}

func (v *NullableAlipayOpenMiniVersionRollbackModel) Set(val *AlipayOpenMiniVersionRollbackModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniVersionRollbackModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniVersionRollbackModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniVersionRollbackModel(val *AlipayOpenMiniVersionRollbackModel) *NullableAlipayOpenMiniVersionRollbackModel {
	return &NullableAlipayOpenMiniVersionRollbackModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniVersionRollbackModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniVersionRollbackModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

