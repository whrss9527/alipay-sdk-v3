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

// checks if the McardNotifyMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &McardNotifyMessage{}

// McardNotifyMessage struct for McardNotifyMessage
type McardNotifyMessage struct {
	// 用户提醒信息，按如下格式拼装，需要ISV提供change_reason。  积分变动模板：{change_reason}，您的积分有变动  余额变动模板：{change_reason}，您的余额有变动  等级变更无需提供原因。
	ChangeReason *string `json:"change_reason,omitempty"`
	// JSON格式扩展信息，主要是发送消息中的变量
	ExtInfo *string `json:"ext_info,omitempty"`
	// 消息类型，每种消息都定义了固定消息模板
	MessageType *string `json:"message_type,omitempty"`
}

// NewMcardNotifyMessage instantiates a new McardNotifyMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMcardNotifyMessage() *McardNotifyMessage {
	this := McardNotifyMessage{}
	return &this
}

// NewMcardNotifyMessageWithDefaults instantiates a new McardNotifyMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMcardNotifyMessageWithDefaults() *McardNotifyMessage {
	this := McardNotifyMessage{}
	return &this
}

// GetChangeReason returns the ChangeReason field value if set, zero value otherwise.
func (o *McardNotifyMessage) GetChangeReason() string {
	if o == nil || IsNil(o.ChangeReason) {
		var ret string
		return ret
	}
	return *o.ChangeReason
}

// GetChangeReasonOk returns a tuple with the ChangeReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McardNotifyMessage) GetChangeReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeReason) {
		return nil, false
	}
	return o.ChangeReason, true
}

// HasChangeReason returns a boolean if a field has been set.
func (o *McardNotifyMessage) HasChangeReason() bool {
	if o != nil && !IsNil(o.ChangeReason) {
		return true
	}

	return false
}

// SetChangeReason gets a reference to the given string and assigns it to the ChangeReason field.
func (o *McardNotifyMessage) SetChangeReason(v string) {
	o.ChangeReason = &v
}

// GetExtInfo returns the ExtInfo field value if set, zero value otherwise.
func (o *McardNotifyMessage) GetExtInfo() string {
	if o == nil || IsNil(o.ExtInfo) {
		var ret string
		return ret
	}
	return *o.ExtInfo
}

// GetExtInfoOk returns a tuple with the ExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McardNotifyMessage) GetExtInfoOk() (*string, bool) {
	if o == nil || IsNil(o.ExtInfo) {
		return nil, false
	}
	return o.ExtInfo, true
}

// HasExtInfo returns a boolean if a field has been set.
func (o *McardNotifyMessage) HasExtInfo() bool {
	if o != nil && !IsNil(o.ExtInfo) {
		return true
	}

	return false
}

// SetExtInfo gets a reference to the given string and assigns it to the ExtInfo field.
func (o *McardNotifyMessage) SetExtInfo(v string) {
	o.ExtInfo = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *McardNotifyMessage) GetMessageType() string {
	if o == nil || IsNil(o.MessageType) {
		var ret string
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McardNotifyMessage) GetMessageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *McardNotifyMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given string and assigns it to the MessageType field.
func (o *McardNotifyMessage) SetMessageType(v string) {
	o.MessageType = &v
}

func (o McardNotifyMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o McardNotifyMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChangeReason) {
		toSerialize["change_reason"] = o.ChangeReason
	}
	if !IsNil(o.ExtInfo) {
		toSerialize["ext_info"] = o.ExtInfo
	}
	if !IsNil(o.MessageType) {
		toSerialize["message_type"] = o.MessageType
	}
	return toSerialize, nil
}

type NullableMcardNotifyMessage struct {
	value *McardNotifyMessage
	isSet bool
}

func (v NullableMcardNotifyMessage) Get() *McardNotifyMessage {
	return v.value
}

func (v *NullableMcardNotifyMessage) Set(val *McardNotifyMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableMcardNotifyMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMcardNotifyMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMcardNotifyMessage(val *McardNotifyMessage) *NullableMcardNotifyMessage {
	return &NullableMcardNotifyMessage{value: val, isSet: true}
}

func (v NullableMcardNotifyMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMcardNotifyMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

