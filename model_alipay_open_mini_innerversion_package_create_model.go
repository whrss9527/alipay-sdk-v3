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

// checks if the AlipayOpenMiniInnerversionPackageCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerversionPackageCreateModel{}

// AlipayOpenMiniInnerversionPackageCreateModel struct for AlipayOpenMiniInnerversionPackageCreateModel
type AlipayOpenMiniInnerversionPackageCreateModel struct {
	// 业务来源
	AppOrigin *string `json:"app_origin,omitempty"`
	// 小程序版本
	AppVersion *string `json:"app_version,omitempty"`
	// 端信息
	BundleId *string `json:"bundle_id,omitempty"`
	// 小程序ID，仅特殊场景使用，普通业务方无需关注该参数。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// 打包类型。预发包和覆盖率包
	PackageType *string `json:"package_type,omitempty"`
	// pid. 伙伴链路（app_origin='HUOBAN'）时必填
	Pid *string `json:"pid,omitempty"`
}

// NewAlipayOpenMiniInnerversionPackageCreateModel instantiates a new AlipayOpenMiniInnerversionPackageCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerversionPackageCreateModel() *AlipayOpenMiniInnerversionPackageCreateModel {
	this := AlipayOpenMiniInnerversionPackageCreateModel{}
	return &this
}

// NewAlipayOpenMiniInnerversionPackageCreateModelWithDefaults instantiates a new AlipayOpenMiniInnerversionPackageCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerversionPackageCreateModelWithDefaults() *AlipayOpenMiniInnerversionPackageCreateModel {
	this := AlipayOpenMiniInnerversionPackageCreateModel{}
	return &this
}

// GetAppOrigin returns the AppOrigin field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) GetAppOrigin() string {
	if o == nil || IsNil(o.AppOrigin) {
		var ret string
		return ret
	}
	return *o.AppOrigin
}

// GetAppOriginOk returns a tuple with the AppOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) GetAppOriginOk() (*string, bool) {
	if o == nil || IsNil(o.AppOrigin) {
		return nil, false
	}
	return o.AppOrigin, true
}

// HasAppOrigin returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) HasAppOrigin() bool {
	if o != nil && !IsNil(o.AppOrigin) {
		return true
	}

	return false
}

// SetAppOrigin gets a reference to the given string and assigns it to the AppOrigin field.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) SetAppOrigin(v string) {
	o.AppOrigin = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion) {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) GetAppVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) SetBundleId(v string) {
	o.BundleId = &v
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

// GetPackageType returns the PackageType field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) GetPackageType() string {
	if o == nil || IsNil(o.PackageType) {
		var ret string
		return ret
	}
	return *o.PackageType
}

// GetPackageTypeOk returns a tuple with the PackageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) GetPackageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PackageType) {
		return nil, false
	}
	return o.PackageType, true
}

// HasPackageType returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) HasPackageType() bool {
	if o != nil && !IsNil(o.PackageType) {
		return true
	}

	return false
}

// SetPackageType gets a reference to the given string and assigns it to the PackageType field.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) SetPackageType(v string) {
	o.PackageType = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *AlipayOpenMiniInnerversionPackageCreateModel) SetPid(v string) {
	o.Pid = &v
}

func (o AlipayOpenMiniInnerversionPackageCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerversionPackageCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppOrigin) {
		toSerialize["app_origin"] = o.AppOrigin
	}
	if !IsNil(o.AppVersion) {
		toSerialize["app_version"] = o.AppVersion
	}
	if !IsNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	if !IsNil(o.PackageType) {
		toSerialize["package_type"] = o.PackageType
	}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniInnerversionPackageCreateModel struct {
	value *AlipayOpenMiniInnerversionPackageCreateModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionPackageCreateModel) Get() *AlipayOpenMiniInnerversionPackageCreateModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionPackageCreateModel) Set(val *AlipayOpenMiniInnerversionPackageCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionPackageCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionPackageCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionPackageCreateModel(val *AlipayOpenMiniInnerversionPackageCreateModel) *NullableAlipayOpenMiniInnerversionPackageCreateModel {
	return &NullableAlipayOpenMiniInnerversionPackageCreateModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionPackageCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionPackageCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

