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

// checks if the AlipayOpenMiniInnerappPluginuseconfigCancelModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerappPluginuseconfigCancelModel{}

// AlipayOpenMiniInnerappPluginuseconfigCancelModel struct for AlipayOpenMiniInnerappPluginuseconfigCancelModel
type AlipayOpenMiniInnerappPluginuseconfigCancelModel struct {
	// 来源
	AppOrigin *string `json:"app_origin,omitempty"`
	// 端id
	BundleId *string `json:"bundle_id,omitempty"`
	// 小程序应用 ID
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// 插件研发版本
	PluginDevVersion *string `json:"plugin_dev_version,omitempty"`
	// 插件id
	PluginId *string `json:"plugin_id,omitempty"`
}

// NewAlipayOpenMiniInnerappPluginuseconfigCancelModel instantiates a new AlipayOpenMiniInnerappPluginuseconfigCancelModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerappPluginuseconfigCancelModel() *AlipayOpenMiniInnerappPluginuseconfigCancelModel {
	this := AlipayOpenMiniInnerappPluginuseconfigCancelModel{}
	return &this
}

// NewAlipayOpenMiniInnerappPluginuseconfigCancelModelWithDefaults instantiates a new AlipayOpenMiniInnerappPluginuseconfigCancelModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerappPluginuseconfigCancelModelWithDefaults() *AlipayOpenMiniInnerappPluginuseconfigCancelModel {
	this := AlipayOpenMiniInnerappPluginuseconfigCancelModel{}
	return &this
}

// GetAppOrigin returns the AppOrigin field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) GetAppOrigin() string {
	if o == nil || IsNil(o.AppOrigin) {
		var ret string
		return ret
	}
	return *o.AppOrigin
}

// GetAppOriginOk returns a tuple with the AppOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) GetAppOriginOk() (*string, bool) {
	if o == nil || IsNil(o.AppOrigin) {
		return nil, false
	}
	return o.AppOrigin, true
}

// HasAppOrigin returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) HasAppOrigin() bool {
	if o != nil && !IsNil(o.AppOrigin) {
		return true
	}

	return false
}

// SetAppOrigin gets a reference to the given string and assigns it to the AppOrigin field.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) SetAppOrigin(v string) {
	o.AppOrigin = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) SetBundleId(v string) {
	o.BundleId = &v
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

// GetPluginDevVersion returns the PluginDevVersion field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) GetPluginDevVersion() string {
	if o == nil || IsNil(o.PluginDevVersion) {
		var ret string
		return ret
	}
	return *o.PluginDevVersion
}

// GetPluginDevVersionOk returns a tuple with the PluginDevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) GetPluginDevVersionOk() (*string, bool) {
	if o == nil || IsNil(o.PluginDevVersion) {
		return nil, false
	}
	return o.PluginDevVersion, true
}

// HasPluginDevVersion returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) HasPluginDevVersion() bool {
	if o != nil && !IsNil(o.PluginDevVersion) {
		return true
	}

	return false
}

// SetPluginDevVersion gets a reference to the given string and assigns it to the PluginDevVersion field.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) SetPluginDevVersion(v string) {
	o.PluginDevVersion = &v
}

// GetPluginId returns the PluginId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) GetPluginId() string {
	if o == nil || IsNil(o.PluginId) {
		var ret string
		return ret
	}
	return *o.PluginId
}

// GetPluginIdOk returns a tuple with the PluginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) GetPluginIdOk() (*string, bool) {
	if o == nil || IsNil(o.PluginId) {
		return nil, false
	}
	return o.PluginId, true
}

// HasPluginId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) HasPluginId() bool {
	if o != nil && !IsNil(o.PluginId) {
		return true
	}

	return false
}

// SetPluginId gets a reference to the given string and assigns it to the PluginId field.
func (o *AlipayOpenMiniInnerappPluginuseconfigCancelModel) SetPluginId(v string) {
	o.PluginId = &v
}

func (o AlipayOpenMiniInnerappPluginuseconfigCancelModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerappPluginuseconfigCancelModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppOrigin) {
		toSerialize["app_origin"] = o.AppOrigin
	}
	if !IsNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	if !IsNil(o.PluginDevVersion) {
		toSerialize["plugin_dev_version"] = o.PluginDevVersion
	}
	if !IsNil(o.PluginId) {
		toSerialize["plugin_id"] = o.PluginId
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniInnerappPluginuseconfigCancelModel struct {
	value *AlipayOpenMiniInnerappPluginuseconfigCancelModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerappPluginuseconfigCancelModel) Get() *AlipayOpenMiniInnerappPluginuseconfigCancelModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerappPluginuseconfigCancelModel) Set(val *AlipayOpenMiniInnerappPluginuseconfigCancelModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerappPluginuseconfigCancelModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerappPluginuseconfigCancelModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerappPluginuseconfigCancelModel(val *AlipayOpenMiniInnerappPluginuseconfigCancelModel) *NullableAlipayOpenMiniInnerappPluginuseconfigCancelModel {
	return &NullableAlipayOpenMiniInnerappPluginuseconfigCancelModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerappPluginuseconfigCancelModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerappPluginuseconfigCancelModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
