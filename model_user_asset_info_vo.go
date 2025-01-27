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

// checks if the UserAssetInfoVO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAssetInfoVO{}

// UserAssetInfoVO struct for UserAssetInfoVO
type UserAssetInfoVO struct {
	// 资产ID
	AssetId *string `json:"asset_id,omitempty"`
	// 资产类型
	AssetType *string `json:"asset_type,omitempty"`
	// 支付宝用户ID
	UserId *string `json:"user_id,omitempty"`
	// 支付宝用户开放ID
	UserOpenId *string `json:"user_open_id,omitempty"`
}

// NewUserAssetInfoVO instantiates a new UserAssetInfoVO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAssetInfoVO() *UserAssetInfoVO {
	this := UserAssetInfoVO{}
	return &this
}

// NewUserAssetInfoVOWithDefaults instantiates a new UserAssetInfoVO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAssetInfoVOWithDefaults() *UserAssetInfoVO {
	this := UserAssetInfoVO{}
	return &this
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *UserAssetInfoVO) GetAssetId() string {
	if o == nil || IsNil(o.AssetId) {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAssetInfoVO) GetAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssetId) {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *UserAssetInfoVO) HasAssetId() bool {
	if o != nil && !IsNil(o.AssetId) {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *UserAssetInfoVO) SetAssetId(v string) {
	o.AssetId = &v
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *UserAssetInfoVO) GetAssetType() string {
	if o == nil || IsNil(o.AssetType) {
		var ret string
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAssetInfoVO) GetAssetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssetType) {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *UserAssetInfoVO) HasAssetType() bool {
	if o != nil && !IsNil(o.AssetType) {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given string and assigns it to the AssetType field.
func (o *UserAssetInfoVO) SetAssetType(v string) {
	o.AssetType = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UserAssetInfoVO) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAssetInfoVO) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UserAssetInfoVO) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *UserAssetInfoVO) SetUserId(v string) {
	o.UserId = &v
}

// GetUserOpenId returns the UserOpenId field value if set, zero value otherwise.
func (o *UserAssetInfoVO) GetUserOpenId() string {
	if o == nil || IsNil(o.UserOpenId) {
		var ret string
		return ret
	}
	return *o.UserOpenId
}

// GetUserOpenIdOk returns a tuple with the UserOpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAssetInfoVO) GetUserOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserOpenId) {
		return nil, false
	}
	return o.UserOpenId, true
}

// HasUserOpenId returns a boolean if a field has been set.
func (o *UserAssetInfoVO) HasUserOpenId() bool {
	if o != nil && !IsNil(o.UserOpenId) {
		return true
	}

	return false
}

// SetUserOpenId gets a reference to the given string and assigns it to the UserOpenId field.
func (o *UserAssetInfoVO) SetUserOpenId(v string) {
	o.UserOpenId = &v
}

func (o UserAssetInfoVO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAssetInfoVO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetId) {
		toSerialize["asset_id"] = o.AssetId
	}
	if !IsNil(o.AssetType) {
		toSerialize["asset_type"] = o.AssetType
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.UserOpenId) {
		toSerialize["user_open_id"] = o.UserOpenId
	}
	return toSerialize, nil
}

type NullableUserAssetInfoVO struct {
	value *UserAssetInfoVO
	isSet bool
}

func (v NullableUserAssetInfoVO) Get() *UserAssetInfoVO {
	return v.value
}

func (v *NullableUserAssetInfoVO) Set(val *UserAssetInfoVO) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAssetInfoVO) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAssetInfoVO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAssetInfoVO(val *UserAssetInfoVO) *NullableUserAssetInfoVO {
	return &NullableUserAssetInfoVO{value: val, isSet: true}
}

func (v NullableUserAssetInfoVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAssetInfoVO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
