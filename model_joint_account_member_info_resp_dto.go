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

// checks if the JointAccountMemberInfoRespDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JointAccountMemberInfoRespDTO{}

// JointAccountMemberInfoRespDTO struct for JointAccountMemberInfoRespDTO
type JointAccountMemberInfoRespDTO struct {
	// （群成员）支付宝侧用户唯一标识
	OpenId *string `json:"open_id,omitempty"`
	// 成员角色：<br> -MASTER：创建人<br> -ADMIN：管理员<br> -MEMBER：群成员<br>
	OperateRole *string `json:"operate_role,omitempty"`
	// （群成员）支付宝侧用户唯一标识
	UserId *string `json:"user_id,omitempty"`
}

// NewJointAccountMemberInfoRespDTO instantiates a new JointAccountMemberInfoRespDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJointAccountMemberInfoRespDTO() *JointAccountMemberInfoRespDTO {
	this := JointAccountMemberInfoRespDTO{}
	return &this
}

// NewJointAccountMemberInfoRespDTOWithDefaults instantiates a new JointAccountMemberInfoRespDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJointAccountMemberInfoRespDTOWithDefaults() *JointAccountMemberInfoRespDTO {
	this := JointAccountMemberInfoRespDTO{}
	return &this
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *JointAccountMemberInfoRespDTO) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JointAccountMemberInfoRespDTO) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *JointAccountMemberInfoRespDTO) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *JointAccountMemberInfoRespDTO) SetOpenId(v string) {
	o.OpenId = &v
}

// GetOperateRole returns the OperateRole field value if set, zero value otherwise.
func (o *JointAccountMemberInfoRespDTO) GetOperateRole() string {
	if o == nil || IsNil(o.OperateRole) {
		var ret string
		return ret
	}
	return *o.OperateRole
}

// GetOperateRoleOk returns a tuple with the OperateRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JointAccountMemberInfoRespDTO) GetOperateRoleOk() (*string, bool) {
	if o == nil || IsNil(o.OperateRole) {
		return nil, false
	}
	return o.OperateRole, true
}

// HasOperateRole returns a boolean if a field has been set.
func (o *JointAccountMemberInfoRespDTO) HasOperateRole() bool {
	if o != nil && !IsNil(o.OperateRole) {
		return true
	}

	return false
}

// SetOperateRole gets a reference to the given string and assigns it to the OperateRole field.
func (o *JointAccountMemberInfoRespDTO) SetOperateRole(v string) {
	o.OperateRole = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *JointAccountMemberInfoRespDTO) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JointAccountMemberInfoRespDTO) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *JointAccountMemberInfoRespDTO) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *JointAccountMemberInfoRespDTO) SetUserId(v string) {
	o.UserId = &v
}

func (o JointAccountMemberInfoRespDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JointAccountMemberInfoRespDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.OperateRole) {
		toSerialize["operate_role"] = o.OperateRole
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableJointAccountMemberInfoRespDTO struct {
	value *JointAccountMemberInfoRespDTO
	isSet bool
}

func (v NullableJointAccountMemberInfoRespDTO) Get() *JointAccountMemberInfoRespDTO {
	return v.value
}

func (v *NullableJointAccountMemberInfoRespDTO) Set(val *JointAccountMemberInfoRespDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableJointAccountMemberInfoRespDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableJointAccountMemberInfoRespDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJointAccountMemberInfoRespDTO(val *JointAccountMemberInfoRespDTO) *NullableJointAccountMemberInfoRespDTO {
	return &NullableJointAccountMemberInfoRespDTO{value: val, isSet: true}
}

func (v NullableJointAccountMemberInfoRespDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJointAccountMemberInfoRespDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
