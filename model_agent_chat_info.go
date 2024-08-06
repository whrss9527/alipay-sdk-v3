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

// checks if the AgentChatInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentChatInfo{}

// AgentChatInfo struct for AgentChatInfo
type AgentChatInfo struct {
	// 数据权限id(租户实例id)
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// 在线扩展技能组id列表
	ExtendedGroupIds []string `json:"extended_group_ids,omitempty"`
	// 在线技能组id
	GroupId *string `json:"group_id,omitempty"`
	// 在线服务等级
	LevelId *string `json:"level_id,omitempty"`
}

// NewAgentChatInfo instantiates a new AgentChatInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentChatInfo() *AgentChatInfo {
	this := AgentChatInfo{}
	return &this
}

// NewAgentChatInfoWithDefaults instantiates a new AgentChatInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentChatInfoWithDefaults() *AgentChatInfo {
	this := AgentChatInfo{}
	return &this
}

// GetCcsInstanceId returns the CcsInstanceId field value if set, zero value otherwise.
func (o *AgentChatInfo) GetCcsInstanceId() string {
	if o == nil || IsNil(o.CcsInstanceId) {
		var ret string
		return ret
	}
	return *o.CcsInstanceId
}

// GetCcsInstanceIdOk returns a tuple with the CcsInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentChatInfo) GetCcsInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.CcsInstanceId) {
		return nil, false
	}
	return o.CcsInstanceId, true
}

// HasCcsInstanceId returns a boolean if a field has been set.
func (o *AgentChatInfo) HasCcsInstanceId() bool {
	if o != nil && !IsNil(o.CcsInstanceId) {
		return true
	}

	return false
}

// SetCcsInstanceId gets a reference to the given string and assigns it to the CcsInstanceId field.
func (o *AgentChatInfo) SetCcsInstanceId(v string) {
	o.CcsInstanceId = &v
}

// GetExtendedGroupIds returns the ExtendedGroupIds field value if set, zero value otherwise.
func (o *AgentChatInfo) GetExtendedGroupIds() []string {
	if o == nil || IsNil(o.ExtendedGroupIds) {
		var ret []string
		return ret
	}
	return o.ExtendedGroupIds
}

// GetExtendedGroupIdsOk returns a tuple with the ExtendedGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentChatInfo) GetExtendedGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtendedGroupIds) {
		return nil, false
	}
	return o.ExtendedGroupIds, true
}

// HasExtendedGroupIds returns a boolean if a field has been set.
func (o *AgentChatInfo) HasExtendedGroupIds() bool {
	if o != nil && !IsNil(o.ExtendedGroupIds) {
		return true
	}

	return false
}

// SetExtendedGroupIds gets a reference to the given []string and assigns it to the ExtendedGroupIds field.
func (o *AgentChatInfo) SetExtendedGroupIds(v []string) {
	o.ExtendedGroupIds = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *AgentChatInfo) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentChatInfo) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *AgentChatInfo) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *AgentChatInfo) SetGroupId(v string) {
	o.GroupId = &v
}

// GetLevelId returns the LevelId field value if set, zero value otherwise.
func (o *AgentChatInfo) GetLevelId() string {
	if o == nil || IsNil(o.LevelId) {
		var ret string
		return ret
	}
	return *o.LevelId
}

// GetLevelIdOk returns a tuple with the LevelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentChatInfo) GetLevelIdOk() (*string, bool) {
	if o == nil || IsNil(o.LevelId) {
		return nil, false
	}
	return o.LevelId, true
}

// HasLevelId returns a boolean if a field has been set.
func (o *AgentChatInfo) HasLevelId() bool {
	if o != nil && !IsNil(o.LevelId) {
		return true
	}

	return false
}

// SetLevelId gets a reference to the given string and assigns it to the LevelId field.
func (o *AgentChatInfo) SetLevelId(v string) {
	o.LevelId = &v
}

func (o AgentChatInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentChatInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CcsInstanceId) {
		toSerialize["ccs_instance_id"] = o.CcsInstanceId
	}
	if !IsNil(o.ExtendedGroupIds) {
		toSerialize["extended_group_ids"] = o.ExtendedGroupIds
	}
	if !IsNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	if !IsNil(o.LevelId) {
		toSerialize["level_id"] = o.LevelId
	}
	return toSerialize, nil
}

type NullableAgentChatInfo struct {
	value *AgentChatInfo
	isSet bool
}

func (v NullableAgentChatInfo) Get() *AgentChatInfo {
	return v.value
}

func (v *NullableAgentChatInfo) Set(val *AgentChatInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentChatInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentChatInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentChatInfo(val *AgentChatInfo) *NullableAgentChatInfo {
	return &NullableAgentChatInfo{value: val, isSet: true}
}

func (v NullableAgentChatInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentChatInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

