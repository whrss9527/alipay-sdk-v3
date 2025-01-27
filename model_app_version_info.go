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

// checks if the AppVersionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppVersionInfo{}

// AppVersionInfo struct for AppVersionInfo
type AppVersionInfo struct {
	// 版本号，一般为x.y.x格式，其中x、y、z都是数字
	AppVersion *string `json:"app_version,omitempty"`
	// 准入审核结果 （提审时 audit_rule 传 BASE_PROMOTE 时有该字段）
	BaseAudit *string `json:"base_audit,omitempty"`
	// 小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。
	BundleId *string `json:"bundle_id,omitempty"`
	// 是否可上架 true：可上架 false：不可上架 （version_status 为PROMOTE_AUDIT_REJECT有值） （提审时 audit_rule 传 BASE_PROMOTE 时有该字段）
	CanRelease *string `json:"can_release,omitempty"`
	// 版本创建时间，精确到秒
	CreateTime *string `json:"create_time,omitempty"`
	// 营销审核结果： PASS：通过 REJECT：驳回 （提审时 audit_rule 传 BASE_PROMOTE 时有该字段）
	PromoteAudit *string `json:"promote_audit,omitempty"`
	// 小程序版本描述，介绍此版本的主要变更和功能，5-500个字符。
	VersionDescription *string `json:"version_description,omitempty"`
	// 版本状态，可选值为：INIT: 开发中, AUDITING: 审核中, AUDIT_REJECT: 审核驳回, WAIT_RELEASE: 待上架, BASE_AUDIT_PASS: 准入不可营销, GRAY: 灰度中, RELEASE: 已上架, OFFLINE: 已下架, AUDIT_OFFLINE: 被强制下架;
	VersionStatus *string `json:"version_status,omitempty"`
}

// NewAppVersionInfo instantiates a new AppVersionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppVersionInfo() *AppVersionInfo {
	this := AppVersionInfo{}
	return &this
}

// NewAppVersionInfoWithDefaults instantiates a new AppVersionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppVersionInfoWithDefaults() *AppVersionInfo {
	this := AppVersionInfo{}
	return &this
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *AppVersionInfo) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion) {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionInfo) GetAppVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *AppVersionInfo) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *AppVersionInfo) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetBaseAudit returns the BaseAudit field value if set, zero value otherwise.
func (o *AppVersionInfo) GetBaseAudit() string {
	if o == nil || IsNil(o.BaseAudit) {
		var ret string
		return ret
	}
	return *o.BaseAudit
}

// GetBaseAuditOk returns a tuple with the BaseAudit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionInfo) GetBaseAuditOk() (*string, bool) {
	if o == nil || IsNil(o.BaseAudit) {
		return nil, false
	}
	return o.BaseAudit, true
}

// HasBaseAudit returns a boolean if a field has been set.
func (o *AppVersionInfo) HasBaseAudit() bool {
	if o != nil && !IsNil(o.BaseAudit) {
		return true
	}

	return false
}

// SetBaseAudit gets a reference to the given string and assigns it to the BaseAudit field.
func (o *AppVersionInfo) SetBaseAudit(v string) {
	o.BaseAudit = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *AppVersionInfo) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionInfo) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *AppVersionInfo) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *AppVersionInfo) SetBundleId(v string) {
	o.BundleId = &v
}

// GetCanRelease returns the CanRelease field value if set, zero value otherwise.
func (o *AppVersionInfo) GetCanRelease() string {
	if o == nil || IsNil(o.CanRelease) {
		var ret string
		return ret
	}
	return *o.CanRelease
}

// GetCanReleaseOk returns a tuple with the CanRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionInfo) GetCanReleaseOk() (*string, bool) {
	if o == nil || IsNil(o.CanRelease) {
		return nil, false
	}
	return o.CanRelease, true
}

// HasCanRelease returns a boolean if a field has been set.
func (o *AppVersionInfo) HasCanRelease() bool {
	if o != nil && !IsNil(o.CanRelease) {
		return true
	}

	return false
}

// SetCanRelease gets a reference to the given string and assigns it to the CanRelease field.
func (o *AppVersionInfo) SetCanRelease(v string) {
	o.CanRelease = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *AppVersionInfo) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionInfo) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *AppVersionInfo) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *AppVersionInfo) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetPromoteAudit returns the PromoteAudit field value if set, zero value otherwise.
func (o *AppVersionInfo) GetPromoteAudit() string {
	if o == nil || IsNil(o.PromoteAudit) {
		var ret string
		return ret
	}
	return *o.PromoteAudit
}

// GetPromoteAuditOk returns a tuple with the PromoteAudit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionInfo) GetPromoteAuditOk() (*string, bool) {
	if o == nil || IsNil(o.PromoteAudit) {
		return nil, false
	}
	return o.PromoteAudit, true
}

// HasPromoteAudit returns a boolean if a field has been set.
func (o *AppVersionInfo) HasPromoteAudit() bool {
	if o != nil && !IsNil(o.PromoteAudit) {
		return true
	}

	return false
}

// SetPromoteAudit gets a reference to the given string and assigns it to the PromoteAudit field.
func (o *AppVersionInfo) SetPromoteAudit(v string) {
	o.PromoteAudit = &v
}

// GetVersionDescription returns the VersionDescription field value if set, zero value otherwise.
func (o *AppVersionInfo) GetVersionDescription() string {
	if o == nil || IsNil(o.VersionDescription) {
		var ret string
		return ret
	}
	return *o.VersionDescription
}

// GetVersionDescriptionOk returns a tuple with the VersionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionInfo) GetVersionDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.VersionDescription) {
		return nil, false
	}
	return o.VersionDescription, true
}

// HasVersionDescription returns a boolean if a field has been set.
func (o *AppVersionInfo) HasVersionDescription() bool {
	if o != nil && !IsNil(o.VersionDescription) {
		return true
	}

	return false
}

// SetVersionDescription gets a reference to the given string and assigns it to the VersionDescription field.
func (o *AppVersionInfo) SetVersionDescription(v string) {
	o.VersionDescription = &v
}

// GetVersionStatus returns the VersionStatus field value if set, zero value otherwise.
func (o *AppVersionInfo) GetVersionStatus() string {
	if o == nil || IsNil(o.VersionStatus) {
		var ret string
		return ret
	}
	return *o.VersionStatus
}

// GetVersionStatusOk returns a tuple with the VersionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppVersionInfo) GetVersionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.VersionStatus) {
		return nil, false
	}
	return o.VersionStatus, true
}

// HasVersionStatus returns a boolean if a field has been set.
func (o *AppVersionInfo) HasVersionStatus() bool {
	if o != nil && !IsNil(o.VersionStatus) {
		return true
	}

	return false
}

// SetVersionStatus gets a reference to the given string and assigns it to the VersionStatus field.
func (o *AppVersionInfo) SetVersionStatus(v string) {
	o.VersionStatus = &v
}

func (o AppVersionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppVersionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppVersion) {
		toSerialize["app_version"] = o.AppVersion
	}
	if !IsNil(o.BaseAudit) {
		toSerialize["base_audit"] = o.BaseAudit
	}
	if !IsNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	if !IsNil(o.CanRelease) {
		toSerialize["can_release"] = o.CanRelease
	}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.PromoteAudit) {
		toSerialize["promote_audit"] = o.PromoteAudit
	}
	if !IsNil(o.VersionDescription) {
		toSerialize["version_description"] = o.VersionDescription
	}
	if !IsNil(o.VersionStatus) {
		toSerialize["version_status"] = o.VersionStatus
	}
	return toSerialize, nil
}

type NullableAppVersionInfo struct {
	value *AppVersionInfo
	isSet bool
}

func (v NullableAppVersionInfo) Get() *AppVersionInfo {
	return v.value
}

func (v *NullableAppVersionInfo) Set(val *AppVersionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAppVersionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAppVersionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppVersionInfo(val *AppVersionInfo) *NullableAppVersionInfo {
	return &NullableAppVersionInfo{value: val, isSet: true}
}

func (v NullableAppVersionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppVersionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
