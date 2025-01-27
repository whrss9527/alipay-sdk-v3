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

// checks if the ExecutionPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutionPlan{}

// ExecutionPlan struct for ExecutionPlan
type ExecutionPlan struct {
	// 周期扣预期执行时间
	ExecuteTime *string `json:"execute_time,omitempty"`
	// 周期扣执行计划最晚执行时间
	LatestExecuteTime *string `json:"latest_execute_time,omitempty"`
	// 周期扣期数
	PeriodId *string `json:"period_id,omitempty"`
	// 周期扣中单笔金额，单位是元
	SingleAmount *string `json:"single_amount,omitempty"`
}

// NewExecutionPlan instantiates a new ExecutionPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionPlan() *ExecutionPlan {
	this := ExecutionPlan{}
	return &this
}

// NewExecutionPlanWithDefaults instantiates a new ExecutionPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionPlanWithDefaults() *ExecutionPlan {
	this := ExecutionPlan{}
	return &this
}

// GetExecuteTime returns the ExecuteTime field value if set, zero value otherwise.
func (o *ExecutionPlan) GetExecuteTime() string {
	if o == nil || IsNil(o.ExecuteTime) {
		var ret string
		return ret
	}
	return *o.ExecuteTime
}

// GetExecuteTimeOk returns a tuple with the ExecuteTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionPlan) GetExecuteTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExecuteTime) {
		return nil, false
	}
	return o.ExecuteTime, true
}

// HasExecuteTime returns a boolean if a field has been set.
func (o *ExecutionPlan) HasExecuteTime() bool {
	if o != nil && !IsNil(o.ExecuteTime) {
		return true
	}

	return false
}

// SetExecuteTime gets a reference to the given string and assigns it to the ExecuteTime field.
func (o *ExecutionPlan) SetExecuteTime(v string) {
	o.ExecuteTime = &v
}

// GetLatestExecuteTime returns the LatestExecuteTime field value if set, zero value otherwise.
func (o *ExecutionPlan) GetLatestExecuteTime() string {
	if o == nil || IsNil(o.LatestExecuteTime) {
		var ret string
		return ret
	}
	return *o.LatestExecuteTime
}

// GetLatestExecuteTimeOk returns a tuple with the LatestExecuteTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionPlan) GetLatestExecuteTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LatestExecuteTime) {
		return nil, false
	}
	return o.LatestExecuteTime, true
}

// HasLatestExecuteTime returns a boolean if a field has been set.
func (o *ExecutionPlan) HasLatestExecuteTime() bool {
	if o != nil && !IsNil(o.LatestExecuteTime) {
		return true
	}

	return false
}

// SetLatestExecuteTime gets a reference to the given string and assigns it to the LatestExecuteTime field.
func (o *ExecutionPlan) SetLatestExecuteTime(v string) {
	o.LatestExecuteTime = &v
}

// GetPeriodId returns the PeriodId field value if set, zero value otherwise.
func (o *ExecutionPlan) GetPeriodId() string {
	if o == nil || IsNil(o.PeriodId) {
		var ret string
		return ret
	}
	return *o.PeriodId
}

// GetPeriodIdOk returns a tuple with the PeriodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionPlan) GetPeriodIdOk() (*string, bool) {
	if o == nil || IsNil(o.PeriodId) {
		return nil, false
	}
	return o.PeriodId, true
}

// HasPeriodId returns a boolean if a field has been set.
func (o *ExecutionPlan) HasPeriodId() bool {
	if o != nil && !IsNil(o.PeriodId) {
		return true
	}

	return false
}

// SetPeriodId gets a reference to the given string and assigns it to the PeriodId field.
func (o *ExecutionPlan) SetPeriodId(v string) {
	o.PeriodId = &v
}

// GetSingleAmount returns the SingleAmount field value if set, zero value otherwise.
func (o *ExecutionPlan) GetSingleAmount() string {
	if o == nil || IsNil(o.SingleAmount) {
		var ret string
		return ret
	}
	return *o.SingleAmount
}

// GetSingleAmountOk returns a tuple with the SingleAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionPlan) GetSingleAmountOk() (*string, bool) {
	if o == nil || IsNil(o.SingleAmount) {
		return nil, false
	}
	return o.SingleAmount, true
}

// HasSingleAmount returns a boolean if a field has been set.
func (o *ExecutionPlan) HasSingleAmount() bool {
	if o != nil && !IsNil(o.SingleAmount) {
		return true
	}

	return false
}

// SetSingleAmount gets a reference to the given string and assigns it to the SingleAmount field.
func (o *ExecutionPlan) SetSingleAmount(v string) {
	o.SingleAmount = &v
}

func (o ExecutionPlan) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExecuteTime) {
		toSerialize["execute_time"] = o.ExecuteTime
	}
	if !IsNil(o.LatestExecuteTime) {
		toSerialize["latest_execute_time"] = o.LatestExecuteTime
	}
	if !IsNil(o.PeriodId) {
		toSerialize["period_id"] = o.PeriodId
	}
	if !IsNil(o.SingleAmount) {
		toSerialize["single_amount"] = o.SingleAmount
	}
	return toSerialize, nil
}

type NullableExecutionPlan struct {
	value *ExecutionPlan
	isSet bool
}

func (v NullableExecutionPlan) Get() *ExecutionPlan {
	return v.value
}

func (v *NullableExecutionPlan) Set(val *ExecutionPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionPlan(val *ExecutionPlan) *NullableExecutionPlan {
	return &NullableExecutionPlan{value: val, isSet: true}
}

func (v NullableExecutionPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
