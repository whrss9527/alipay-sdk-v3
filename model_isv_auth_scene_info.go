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

// checks if the IsvAuthSceneInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IsvAuthSceneInfo{}

// IsvAuthSceneInfo struct for IsvAuthSceneInfo
type IsvAuthSceneInfo struct {
	// 运营场景编码 OPERATION_POINTS：管理运营积分 SHOP_MANAGE：管理门店信息 MINI_APP_OPER：运营支付宝小程序 PROMOTION_MANAGE：运营营销活动
	SceneCode *string `json:"scene_code,omitempty"`
	// 运营场景下的权限编码，多个权限编码以,隔开 1、管理门店信息：SHOP_MANAGE；基础权限（升级）：SHOP_MANAGE_BASE 2、运营营销活动：PROMOTION_MANAGE ；基础权限（升级）：PROMOTION_MANAGE_BASE 3、运营支付宝小程序：MINI_APP_OPER；基础权限（升级）：MINI_APP_OPER_BASE 4、管理运营积分：OPERATION_POINTS；基础权限（升级）：OPERATION_POINTS_BASE
	ScenePermissions *string `json:"scene_permissions,omitempty"`
}

// NewIsvAuthSceneInfo instantiates a new IsvAuthSceneInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsvAuthSceneInfo() *IsvAuthSceneInfo {
	this := IsvAuthSceneInfo{}
	return &this
}

// NewIsvAuthSceneInfoWithDefaults instantiates a new IsvAuthSceneInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsvAuthSceneInfoWithDefaults() *IsvAuthSceneInfo {
	this := IsvAuthSceneInfo{}
	return &this
}

// GetSceneCode returns the SceneCode field value if set, zero value otherwise.
func (o *IsvAuthSceneInfo) GetSceneCode() string {
	if o == nil || IsNil(o.SceneCode) {
		var ret string
		return ret
	}
	return *o.SceneCode
}

// GetSceneCodeOk returns a tuple with the SceneCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvAuthSceneInfo) GetSceneCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SceneCode) {
		return nil, false
	}
	return o.SceneCode, true
}

// HasSceneCode returns a boolean if a field has been set.
func (o *IsvAuthSceneInfo) HasSceneCode() bool {
	if o != nil && !IsNil(o.SceneCode) {
		return true
	}

	return false
}

// SetSceneCode gets a reference to the given string and assigns it to the SceneCode field.
func (o *IsvAuthSceneInfo) SetSceneCode(v string) {
	o.SceneCode = &v
}

// GetScenePermissions returns the ScenePermissions field value if set, zero value otherwise.
func (o *IsvAuthSceneInfo) GetScenePermissions() string {
	if o == nil || IsNil(o.ScenePermissions) {
		var ret string
		return ret
	}
	return *o.ScenePermissions
}

// GetScenePermissionsOk returns a tuple with the ScenePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsvAuthSceneInfo) GetScenePermissionsOk() (*string, bool) {
	if o == nil || IsNil(o.ScenePermissions) {
		return nil, false
	}
	return o.ScenePermissions, true
}

// HasScenePermissions returns a boolean if a field has been set.
func (o *IsvAuthSceneInfo) HasScenePermissions() bool {
	if o != nil && !IsNil(o.ScenePermissions) {
		return true
	}

	return false
}

// SetScenePermissions gets a reference to the given string and assigns it to the ScenePermissions field.
func (o *IsvAuthSceneInfo) SetScenePermissions(v string) {
	o.ScenePermissions = &v
}

func (o IsvAuthSceneInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IsvAuthSceneInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SceneCode) {
		toSerialize["scene_code"] = o.SceneCode
	}
	if !IsNil(o.ScenePermissions) {
		toSerialize["scene_permissions"] = o.ScenePermissions
	}
	return toSerialize, nil
}

type NullableIsvAuthSceneInfo struct {
	value *IsvAuthSceneInfo
	isSet bool
}

func (v NullableIsvAuthSceneInfo) Get() *IsvAuthSceneInfo {
	return v.value
}

func (v *NullableIsvAuthSceneInfo) Set(val *IsvAuthSceneInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIsvAuthSceneInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIsvAuthSceneInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsvAuthSceneInfo(val *IsvAuthSceneInfo) *NullableIsvAuthSceneInfo {
	return &NullableIsvAuthSceneInfo{value: val, isSet: true}
}

func (v NullableIsvAuthSceneInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsvAuthSceneInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
