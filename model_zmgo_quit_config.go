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

// checks if the ZMGOQuitConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZMGOQuitConfig{}

// ZMGOQuitConfig struct for ZMGOQuitConfig
type ZMGOQuitConfig struct {
	// 退出提示
	QuitDesc *string `json:"quit_desc,omitempty"`
	// 退出跳转商家页地址，当quit_type取值为：QUIT_MERCHANT_JUMP，必传
	QuitJumpUrl *string `json:"quit_jump_url,omitempty"`
	// 退出芝麻GO方式，取值：   （1）QUIT_COMMON(\"QUIT_COMMON\", \"标准退出\"), （2）QUIT_MERCHANT_JUMP(\"QUIT_MERCHANT_JUMP\", \"商家页退出\"), （3）QUIT_ONLY_TIPS(\"QUIT_ONLY_TIPS\", \"仅退出提示\")
	QuitType *string `json:"quit_type,omitempty"`
}

// NewZMGOQuitConfig instantiates a new ZMGOQuitConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZMGOQuitConfig() *ZMGOQuitConfig {
	this := ZMGOQuitConfig{}
	return &this
}

// NewZMGOQuitConfigWithDefaults instantiates a new ZMGOQuitConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZMGOQuitConfigWithDefaults() *ZMGOQuitConfig {
	this := ZMGOQuitConfig{}
	return &this
}

// GetQuitDesc returns the QuitDesc field value if set, zero value otherwise.
func (o *ZMGOQuitConfig) GetQuitDesc() string {
	if o == nil || IsNil(o.QuitDesc) {
		var ret string
		return ret
	}
	return *o.QuitDesc
}

// GetQuitDescOk returns a tuple with the QuitDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOQuitConfig) GetQuitDescOk() (*string, bool) {
	if o == nil || IsNil(o.QuitDesc) {
		return nil, false
	}
	return o.QuitDesc, true
}

// HasQuitDesc returns a boolean if a field has been set.
func (o *ZMGOQuitConfig) HasQuitDesc() bool {
	if o != nil && !IsNil(o.QuitDesc) {
		return true
	}

	return false
}

// SetQuitDesc gets a reference to the given string and assigns it to the QuitDesc field.
func (o *ZMGOQuitConfig) SetQuitDesc(v string) {
	o.QuitDesc = &v
}

// GetQuitJumpUrl returns the QuitJumpUrl field value if set, zero value otherwise.
func (o *ZMGOQuitConfig) GetQuitJumpUrl() string {
	if o == nil || IsNil(o.QuitJumpUrl) {
		var ret string
		return ret
	}
	return *o.QuitJumpUrl
}

// GetQuitJumpUrlOk returns a tuple with the QuitJumpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOQuitConfig) GetQuitJumpUrlOk() (*string, bool) {
	if o == nil || IsNil(o.QuitJumpUrl) {
		return nil, false
	}
	return o.QuitJumpUrl, true
}

// HasQuitJumpUrl returns a boolean if a field has been set.
func (o *ZMGOQuitConfig) HasQuitJumpUrl() bool {
	if o != nil && !IsNil(o.QuitJumpUrl) {
		return true
	}

	return false
}

// SetQuitJumpUrl gets a reference to the given string and assigns it to the QuitJumpUrl field.
func (o *ZMGOQuitConfig) SetQuitJumpUrl(v string) {
	o.QuitJumpUrl = &v
}

// GetQuitType returns the QuitType field value if set, zero value otherwise.
func (o *ZMGOQuitConfig) GetQuitType() string {
	if o == nil || IsNil(o.QuitType) {
		var ret string
		return ret
	}
	return *o.QuitType
}

// GetQuitTypeOk returns a tuple with the QuitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOQuitConfig) GetQuitTypeOk() (*string, bool) {
	if o == nil || IsNil(o.QuitType) {
		return nil, false
	}
	return o.QuitType, true
}

// HasQuitType returns a boolean if a field has been set.
func (o *ZMGOQuitConfig) HasQuitType() bool {
	if o != nil && !IsNil(o.QuitType) {
		return true
	}

	return false
}

// SetQuitType gets a reference to the given string and assigns it to the QuitType field.
func (o *ZMGOQuitConfig) SetQuitType(v string) {
	o.QuitType = &v
}

func (o ZMGOQuitConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZMGOQuitConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QuitDesc) {
		toSerialize["quit_desc"] = o.QuitDesc
	}
	if !IsNil(o.QuitJumpUrl) {
		toSerialize["quit_jump_url"] = o.QuitJumpUrl
	}
	if !IsNil(o.QuitType) {
		toSerialize["quit_type"] = o.QuitType
	}
	return toSerialize, nil
}

type NullableZMGOQuitConfig struct {
	value *ZMGOQuitConfig
	isSet bool
}

func (v NullableZMGOQuitConfig) Get() *ZMGOQuitConfig {
	return v.value
}

func (v *NullableZMGOQuitConfig) Set(val *ZMGOQuitConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableZMGOQuitConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableZMGOQuitConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZMGOQuitConfig(val *ZMGOQuitConfig) *NullableZMGOQuitConfig {
	return &NullableZMGOQuitConfig{value: val, isSet: true}
}

func (v NullableZMGOQuitConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZMGOQuitConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
