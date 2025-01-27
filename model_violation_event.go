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

// checks if the ViolationEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViolationEvent{}

// ViolationEvent struct for ViolationEvent
type ViolationEvent struct {
	// 商家是否可以申诉
	CanAppeal *bool `json:"can_appeal,omitempty"`
	// 商家是否可以整改
	CanRectify *bool `json:"can_rectify,omitempty"`
	// 处罚动作及有效期
	PunishAction *string `json:"punish_action,omitempty"`
	// 违规工单状态枚举
	Status *string `json:"status,omitempty"`
	// 违规对象ID
	TargetId *string `json:"target_id,omitempty"`
	// 违规对象名称
	TargetName *string `json:"target_name,omitempty"`
	// 违规对象类型 小程序ID:APPID  生活号ID:PUBLICID
	TargetType *string `json:"target_type,omitempty"`
	// 支付宝侧生成的违规记录唯一标识
	ViolationRecordId *string `json:"violation_record_id,omitempty"`
	// 违规时间，格式为 yyyy-MM-dd HH:mm:ss
	ViolationTime *string `json:"violation_time,omitempty"`
	// 即平台依据平台规范/规则，判定商户的违规类型
	ViolationType *string `json:"violation_type,omitempty"`
}

// NewViolationEvent instantiates a new ViolationEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViolationEvent() *ViolationEvent {
	this := ViolationEvent{}
	return &this
}

// NewViolationEventWithDefaults instantiates a new ViolationEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViolationEventWithDefaults() *ViolationEvent {
	this := ViolationEvent{}
	return &this
}

// GetCanAppeal returns the CanAppeal field value if set, zero value otherwise.
func (o *ViolationEvent) GetCanAppeal() bool {
	if o == nil || IsNil(o.CanAppeal) {
		var ret bool
		return ret
	}
	return *o.CanAppeal
}

// GetCanAppealOk returns a tuple with the CanAppeal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetCanAppealOk() (*bool, bool) {
	if o == nil || IsNil(o.CanAppeal) {
		return nil, false
	}
	return o.CanAppeal, true
}

// HasCanAppeal returns a boolean if a field has been set.
func (o *ViolationEvent) HasCanAppeal() bool {
	if o != nil && !IsNil(o.CanAppeal) {
		return true
	}

	return false
}

// SetCanAppeal gets a reference to the given bool and assigns it to the CanAppeal field.
func (o *ViolationEvent) SetCanAppeal(v bool) {
	o.CanAppeal = &v
}

// GetCanRectify returns the CanRectify field value if set, zero value otherwise.
func (o *ViolationEvent) GetCanRectify() bool {
	if o == nil || IsNil(o.CanRectify) {
		var ret bool
		return ret
	}
	return *o.CanRectify
}

// GetCanRectifyOk returns a tuple with the CanRectify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetCanRectifyOk() (*bool, bool) {
	if o == nil || IsNil(o.CanRectify) {
		return nil, false
	}
	return o.CanRectify, true
}

// HasCanRectify returns a boolean if a field has been set.
func (o *ViolationEvent) HasCanRectify() bool {
	if o != nil && !IsNil(o.CanRectify) {
		return true
	}

	return false
}

// SetCanRectify gets a reference to the given bool and assigns it to the CanRectify field.
func (o *ViolationEvent) SetCanRectify(v bool) {
	o.CanRectify = &v
}

// GetPunishAction returns the PunishAction field value if set, zero value otherwise.
func (o *ViolationEvent) GetPunishAction() string {
	if o == nil || IsNil(o.PunishAction) {
		var ret string
		return ret
	}
	return *o.PunishAction
}

// GetPunishActionOk returns a tuple with the PunishAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetPunishActionOk() (*string, bool) {
	if o == nil || IsNil(o.PunishAction) {
		return nil, false
	}
	return o.PunishAction, true
}

// HasPunishAction returns a boolean if a field has been set.
func (o *ViolationEvent) HasPunishAction() bool {
	if o != nil && !IsNil(o.PunishAction) {
		return true
	}

	return false
}

// SetPunishAction gets a reference to the given string and assigns it to the PunishAction field.
func (o *ViolationEvent) SetPunishAction(v string) {
	o.PunishAction = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ViolationEvent) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ViolationEvent) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ViolationEvent) SetStatus(v string) {
	o.Status = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *ViolationEvent) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *ViolationEvent) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *ViolationEvent) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *ViolationEvent) GetTargetName() string {
	if o == nil || IsNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetTargetNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetName) {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *ViolationEvent) HasTargetName() bool {
	if o != nil && !IsNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *ViolationEvent) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *ViolationEvent) GetTargetType() string {
	if o == nil || IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *ViolationEvent) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *ViolationEvent) SetTargetType(v string) {
	o.TargetType = &v
}

// GetViolationRecordId returns the ViolationRecordId field value if set, zero value otherwise.
func (o *ViolationEvent) GetViolationRecordId() string {
	if o == nil || IsNil(o.ViolationRecordId) {
		var ret string
		return ret
	}
	return *o.ViolationRecordId
}

// GetViolationRecordIdOk returns a tuple with the ViolationRecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetViolationRecordIdOk() (*string, bool) {
	if o == nil || IsNil(o.ViolationRecordId) {
		return nil, false
	}
	return o.ViolationRecordId, true
}

// HasViolationRecordId returns a boolean if a field has been set.
func (o *ViolationEvent) HasViolationRecordId() bool {
	if o != nil && !IsNil(o.ViolationRecordId) {
		return true
	}

	return false
}

// SetViolationRecordId gets a reference to the given string and assigns it to the ViolationRecordId field.
func (o *ViolationEvent) SetViolationRecordId(v string) {
	o.ViolationRecordId = &v
}

// GetViolationTime returns the ViolationTime field value if set, zero value otherwise.
func (o *ViolationEvent) GetViolationTime() string {
	if o == nil || IsNil(o.ViolationTime) {
		var ret string
		return ret
	}
	return *o.ViolationTime
}

// GetViolationTimeOk returns a tuple with the ViolationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetViolationTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ViolationTime) {
		return nil, false
	}
	return o.ViolationTime, true
}

// HasViolationTime returns a boolean if a field has been set.
func (o *ViolationEvent) HasViolationTime() bool {
	if o != nil && !IsNil(o.ViolationTime) {
		return true
	}

	return false
}

// SetViolationTime gets a reference to the given string and assigns it to the ViolationTime field.
func (o *ViolationEvent) SetViolationTime(v string) {
	o.ViolationTime = &v
}

// GetViolationType returns the ViolationType field value if set, zero value otherwise.
func (o *ViolationEvent) GetViolationType() string {
	if o == nil || IsNil(o.ViolationType) {
		var ret string
		return ret
	}
	return *o.ViolationType
}

// GetViolationTypeOk returns a tuple with the ViolationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetViolationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ViolationType) {
		return nil, false
	}
	return o.ViolationType, true
}

// HasViolationType returns a boolean if a field has been set.
func (o *ViolationEvent) HasViolationType() bool {
	if o != nil && !IsNil(o.ViolationType) {
		return true
	}

	return false
}

// SetViolationType gets a reference to the given string and assigns it to the ViolationType field.
func (o *ViolationEvent) SetViolationType(v string) {
	o.ViolationType = &v
}

func (o ViolationEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViolationEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanAppeal) {
		toSerialize["can_appeal"] = o.CanAppeal
	}
	if !IsNil(o.CanRectify) {
		toSerialize["can_rectify"] = o.CanRectify
	}
	if !IsNil(o.PunishAction) {
		toSerialize["punish_action"] = o.PunishAction
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TargetId) {
		toSerialize["target_id"] = o.TargetId
	}
	if !IsNil(o.TargetName) {
		toSerialize["target_name"] = o.TargetName
	}
	if !IsNil(o.TargetType) {
		toSerialize["target_type"] = o.TargetType
	}
	if !IsNil(o.ViolationRecordId) {
		toSerialize["violation_record_id"] = o.ViolationRecordId
	}
	if !IsNil(o.ViolationTime) {
		toSerialize["violation_time"] = o.ViolationTime
	}
	if !IsNil(o.ViolationType) {
		toSerialize["violation_type"] = o.ViolationType
	}
	return toSerialize, nil
}

type NullableViolationEvent struct {
	value *ViolationEvent
	isSet bool
}

func (v NullableViolationEvent) Get() *ViolationEvent {
	return v.value
}

func (v *NullableViolationEvent) Set(val *ViolationEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableViolationEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableViolationEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViolationEvent(val *ViolationEvent) *NullableViolationEvent {
	return &NullableViolationEvent{value: val, isSet: true}
}

func (v NullableViolationEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViolationEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
