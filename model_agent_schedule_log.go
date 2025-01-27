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

// checks if the AgentScheduleLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentScheduleLog{}

// AgentScheduleLog struct for AgentScheduleLog
type AgentScheduleLog struct {
	// 客服id
	AgentId *string `json:"agent_id,omitempty"`
	// 客服名称
	AgentName *string `json:"agent_name,omitempty"`
	// 状态变更发生时间,采用UTC时间，按照ISO8601标准表示，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'
	CreateTime *string `json:"create_time,omitempty"`
	// 状态持续时长,单位秒
	Duration *int32 `json:"duration,omitempty"`
	// 状态变更结束时间,采用UTC时间，按照ISO8601标准表示，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'
	EndTime *string `json:"end_time,omitempty"`
	// isv或商户系统中对应的客服工号
	ExternalUserNo *string `json:"external_user_no,omitempty"`
	// 客服状态变更流水id
	Id *string `json:"id,omitempty"`
	// 变更前状态
	LastStatus *string `json:"last_status,omitempty"`
	// 状态变更开始时间,采用UTC时间，按照ISO8601标准表示，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'
	StartTime *string `json:"start_time,omitempty"`
	// 变更后状态
	Status *string `json:"status,omitempty"`
}

// NewAgentScheduleLog instantiates a new AgentScheduleLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentScheduleLog() *AgentScheduleLog {
	this := AgentScheduleLog{}
	return &this
}

// NewAgentScheduleLogWithDefaults instantiates a new AgentScheduleLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentScheduleLogWithDefaults() *AgentScheduleLog {
	this := AgentScheduleLog{}
	return &this
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *AgentScheduleLog) GetAgentId() string {
	if o == nil || IsNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentScheduleLog) GetAgentIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *AgentScheduleLog) HasAgentId() bool {
	if o != nil && !IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *AgentScheduleLog) SetAgentId(v string) {
	o.AgentId = &v
}

// GetAgentName returns the AgentName field value if set, zero value otherwise.
func (o *AgentScheduleLog) GetAgentName() string {
	if o == nil || IsNil(o.AgentName) {
		var ret string
		return ret
	}
	return *o.AgentName
}

// GetAgentNameOk returns a tuple with the AgentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentScheduleLog) GetAgentNameOk() (*string, bool) {
	if o == nil || IsNil(o.AgentName) {
		return nil, false
	}
	return o.AgentName, true
}

// HasAgentName returns a boolean if a field has been set.
func (o *AgentScheduleLog) HasAgentName() bool {
	if o != nil && !IsNil(o.AgentName) {
		return true
	}

	return false
}

// SetAgentName gets a reference to the given string and assigns it to the AgentName field.
func (o *AgentScheduleLog) SetAgentName(v string) {
	o.AgentName = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *AgentScheduleLog) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentScheduleLog) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *AgentScheduleLog) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *AgentScheduleLog) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *AgentScheduleLog) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentScheduleLog) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *AgentScheduleLog) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *AgentScheduleLog) SetDuration(v int32) {
	o.Duration = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *AgentScheduleLog) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentScheduleLog) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *AgentScheduleLog) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *AgentScheduleLog) SetEndTime(v string) {
	o.EndTime = &v
}

// GetExternalUserNo returns the ExternalUserNo field value if set, zero value otherwise.
func (o *AgentScheduleLog) GetExternalUserNo() string {
	if o == nil || IsNil(o.ExternalUserNo) {
		var ret string
		return ret
	}
	return *o.ExternalUserNo
}

// GetExternalUserNoOk returns a tuple with the ExternalUserNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentScheduleLog) GetExternalUserNoOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalUserNo) {
		return nil, false
	}
	return o.ExternalUserNo, true
}

// HasExternalUserNo returns a boolean if a field has been set.
func (o *AgentScheduleLog) HasExternalUserNo() bool {
	if o != nil && !IsNil(o.ExternalUserNo) {
		return true
	}

	return false
}

// SetExternalUserNo gets a reference to the given string and assigns it to the ExternalUserNo field.
func (o *AgentScheduleLog) SetExternalUserNo(v string) {
	o.ExternalUserNo = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgentScheduleLog) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentScheduleLog) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgentScheduleLog) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgentScheduleLog) SetId(v string) {
	o.Id = &v
}

// GetLastStatus returns the LastStatus field value if set, zero value otherwise.
func (o *AgentScheduleLog) GetLastStatus() string {
	if o == nil || IsNil(o.LastStatus) {
		var ret string
		return ret
	}
	return *o.LastStatus
}

// GetLastStatusOk returns a tuple with the LastStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentScheduleLog) GetLastStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LastStatus) {
		return nil, false
	}
	return o.LastStatus, true
}

// HasLastStatus returns a boolean if a field has been set.
func (o *AgentScheduleLog) HasLastStatus() bool {
	if o != nil && !IsNil(o.LastStatus) {
		return true
	}

	return false
}

// SetLastStatus gets a reference to the given string and assigns it to the LastStatus field.
func (o *AgentScheduleLog) SetLastStatus(v string) {
	o.LastStatus = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *AgentScheduleLog) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentScheduleLog) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *AgentScheduleLog) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *AgentScheduleLog) SetStartTime(v string) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AgentScheduleLog) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentScheduleLog) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AgentScheduleLog) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AgentScheduleLog) SetStatus(v string) {
	o.Status = &v
}

func (o AgentScheduleLog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentScheduleLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgentId) {
		toSerialize["agent_id"] = o.AgentId
	}
	if !IsNil(o.AgentName) {
		toSerialize["agent_name"] = o.AgentName
	}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.ExternalUserNo) {
		toSerialize["external_user_no"] = o.ExternalUserNo
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastStatus) {
		toSerialize["last_status"] = o.LastStatus
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAgentScheduleLog struct {
	value *AgentScheduleLog
	isSet bool
}

func (v NullableAgentScheduleLog) Get() *AgentScheduleLog {
	return v.value
}

func (v *NullableAgentScheduleLog) Set(val *AgentScheduleLog) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentScheduleLog) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentScheduleLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentScheduleLog(val *AgentScheduleLog) *NullableAgentScheduleLog {
	return &NullableAgentScheduleLog{value: val, isSet: true}
}

func (v NullableAgentScheduleLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentScheduleLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
