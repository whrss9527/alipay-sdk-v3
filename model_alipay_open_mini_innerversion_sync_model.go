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

// checks if the AlipayOpenMiniInnerversionSyncModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerversionSyncModel{}

// AlipayOpenMiniInnerversionSyncModel struct for AlipayOpenMiniInnerversionSyncModel
type AlipayOpenMiniInnerversionSyncModel struct {
	// 业务参数来源
	AppOrigin *string `json:"app_origin,omitempty"`
	// 推送的小程序版本号
	AppVersion *string `json:"app_version,omitempty"`
	// 端信息
	BundleId *string `json:"bundle_id,omitempty"`
	// 操作人ID
	DevId *string `json:"dev_id,omitempty"`
	// 小程序ID，仅特殊场景使用，普通业务方无需关注该参数。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// 同送方式
	SyncType *string `json:"sync_type,omitempty"`
}

// NewAlipayOpenMiniInnerversionSyncModel instantiates a new AlipayOpenMiniInnerversionSyncModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerversionSyncModel() *AlipayOpenMiniInnerversionSyncModel {
	this := AlipayOpenMiniInnerversionSyncModel{}
	return &this
}

// NewAlipayOpenMiniInnerversionSyncModelWithDefaults instantiates a new AlipayOpenMiniInnerversionSyncModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerversionSyncModelWithDefaults() *AlipayOpenMiniInnerversionSyncModel {
	this := AlipayOpenMiniInnerversionSyncModel{}
	return &this
}

// GetAppOrigin returns the AppOrigin field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionSyncModel) GetAppOrigin() string {
	if o == nil || IsNil(o.AppOrigin) {
		var ret string
		return ret
	}
	return *o.AppOrigin
}

// GetAppOriginOk returns a tuple with the AppOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionSyncModel) GetAppOriginOk() (*string, bool) {
	if o == nil || IsNil(o.AppOrigin) {
		return nil, false
	}
	return o.AppOrigin, true
}

// HasAppOrigin returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionSyncModel) HasAppOrigin() bool {
	if o != nil && !IsNil(o.AppOrigin) {
		return true
	}

	return false
}

// SetAppOrigin gets a reference to the given string and assigns it to the AppOrigin field.
func (o *AlipayOpenMiniInnerversionSyncModel) SetAppOrigin(v string) {
	o.AppOrigin = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionSyncModel) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion) {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionSyncModel) GetAppVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionSyncModel) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *AlipayOpenMiniInnerversionSyncModel) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionSyncModel) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionSyncModel) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionSyncModel) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *AlipayOpenMiniInnerversionSyncModel) SetBundleId(v string) {
	o.BundleId = &v
}

// GetDevId returns the DevId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionSyncModel) GetDevId() string {
	if o == nil || IsNil(o.DevId) {
		var ret string
		return ret
	}
	return *o.DevId
}

// GetDevIdOk returns a tuple with the DevId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionSyncModel) GetDevIdOk() (*string, bool) {
	if o == nil || IsNil(o.DevId) {
		return nil, false
	}
	return o.DevId, true
}

// HasDevId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionSyncModel) HasDevId() bool {
	if o != nil && !IsNil(o.DevId) {
		return true
	}

	return false
}

// SetDevId gets a reference to the given string and assigns it to the DevId field.
func (o *AlipayOpenMiniInnerversionSyncModel) SetDevId(v string) {
	o.DevId = &v
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionSyncModel) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionSyncModel) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionSyncModel) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *AlipayOpenMiniInnerversionSyncModel) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

// GetSyncType returns the SyncType field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionSyncModel) GetSyncType() string {
	if o == nil || IsNil(o.SyncType) {
		var ret string
		return ret
	}
	return *o.SyncType
}

// GetSyncTypeOk returns a tuple with the SyncType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionSyncModel) GetSyncTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SyncType) {
		return nil, false
	}
	return o.SyncType, true
}

// HasSyncType returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionSyncModel) HasSyncType() bool {
	if o != nil && !IsNil(o.SyncType) {
		return true
	}

	return false
}

// SetSyncType gets a reference to the given string and assigns it to the SyncType field.
func (o *AlipayOpenMiniInnerversionSyncModel) SetSyncType(v string) {
	o.SyncType = &v
}

func (o AlipayOpenMiniInnerversionSyncModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerversionSyncModel) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DevId) {
		toSerialize["dev_id"] = o.DevId
	}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	if !IsNil(o.SyncType) {
		toSerialize["sync_type"] = o.SyncType
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniInnerversionSyncModel struct {
	value *AlipayOpenMiniInnerversionSyncModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionSyncModel) Get() *AlipayOpenMiniInnerversionSyncModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionSyncModel) Set(val *AlipayOpenMiniInnerversionSyncModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionSyncModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionSyncModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionSyncModel(val *AlipayOpenMiniInnerversionSyncModel) *NullableAlipayOpenMiniInnerversionSyncModel {
	return &NullableAlipayOpenMiniInnerversionSyncModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionSyncModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionSyncModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
