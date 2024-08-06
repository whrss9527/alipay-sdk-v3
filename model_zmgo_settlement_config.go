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

// checks if the ZMGOSettlementConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZMGOSettlementConfig{}

// ZMGOSettlementConfig struct for ZMGOSettlementConfig
type ZMGOSettlementConfig struct {
	// 会员费扣费名称
	CustomFeeName *string `json:"custom_fee_name,omitempty"`
	CycleFlexWithholdConfig *ZMGOCycleFlexConfig `json:"cycle_flex_withhold_config,omitempty"`
	CycleWithholdConfig *ZMGOCycleWithholdConfig `json:"cycle_withhold_config,omitempty"`
	// 预计结算延迟天数，当exp_stop_time_mode取值RELATIVE_DATE、ABSOLATE_DATE、ABSOLUTE_PLUS_1_DATE时，必传
	ExpStopDelayDays *int32 `json:"exp_stop_delay_days,omitempty"`
	// 预期结算时间，当exp_stop_time_mode取值APPOINT_DATE时，必传
	ExpStopTime *string `json:"exp_stop_time,omitempty"`
	// 预计结算时间模式，取值：（1）RELATIVE_DATE(\"RELATIVE_DATE\", \"相对日期\"), （2）ABSOLATE_DATE(\"ABSOLATE_DATE\", \"绝对日期(当日开始计算)\"), （3）ABSOLUTE_PLUS_1_DATE(\"ABSOLUTE_PLUS_1_DATE\", \"绝对日期(次日开始计算)\"), （4）APPOINT_DATE(\"APPOINT_DATE\", \"指定日期\")
	ExpStopTimeMode *string `json:"exp_stop_time_mode,omitempty"`
	// 按履约次数灵活扣款计划。当settlement_type取值为fulfillTimesCustomSettlement时，此属性必传
	FulfilltimesCustomSettlementPlan *string `json:"fulfilltimes_custom_settlement_plan,omitempty"`
	// 结算类型，取值：     BENEFIT_SETTLEMENT(\"benefitSettlement\", \"按权益金额结算\"),     FULFILL_TIMES_CUSTOM_SETTLEMENT(\"fulfillTimesCustomSettlement\",\"按履约次数扣款\")
	SettlementType *string `json:"settlement_type,omitempty"`
}

// NewZMGOSettlementConfig instantiates a new ZMGOSettlementConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZMGOSettlementConfig() *ZMGOSettlementConfig {
	this := ZMGOSettlementConfig{}
	return &this
}

// NewZMGOSettlementConfigWithDefaults instantiates a new ZMGOSettlementConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZMGOSettlementConfigWithDefaults() *ZMGOSettlementConfig {
	this := ZMGOSettlementConfig{}
	return &this
}

// GetCustomFeeName returns the CustomFeeName field value if set, zero value otherwise.
func (o *ZMGOSettlementConfig) GetCustomFeeName() string {
	if o == nil || IsNil(o.CustomFeeName) {
		var ret string
		return ret
	}
	return *o.CustomFeeName
}

// GetCustomFeeNameOk returns a tuple with the CustomFeeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOSettlementConfig) GetCustomFeeNameOk() (*string, bool) {
	if o == nil || IsNil(o.CustomFeeName) {
		return nil, false
	}
	return o.CustomFeeName, true
}

// HasCustomFeeName returns a boolean if a field has been set.
func (o *ZMGOSettlementConfig) HasCustomFeeName() bool {
	if o != nil && !IsNil(o.CustomFeeName) {
		return true
	}

	return false
}

// SetCustomFeeName gets a reference to the given string and assigns it to the CustomFeeName field.
func (o *ZMGOSettlementConfig) SetCustomFeeName(v string) {
	o.CustomFeeName = &v
}

// GetCycleFlexWithholdConfig returns the CycleFlexWithholdConfig field value if set, zero value otherwise.
func (o *ZMGOSettlementConfig) GetCycleFlexWithholdConfig() ZMGOCycleFlexConfig {
	if o == nil || IsNil(o.CycleFlexWithholdConfig) {
		var ret ZMGOCycleFlexConfig
		return ret
	}
	return *o.CycleFlexWithholdConfig
}

// GetCycleFlexWithholdConfigOk returns a tuple with the CycleFlexWithholdConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOSettlementConfig) GetCycleFlexWithholdConfigOk() (*ZMGOCycleFlexConfig, bool) {
	if o == nil || IsNil(o.CycleFlexWithholdConfig) {
		return nil, false
	}
	return o.CycleFlexWithholdConfig, true
}

// HasCycleFlexWithholdConfig returns a boolean if a field has been set.
func (o *ZMGOSettlementConfig) HasCycleFlexWithholdConfig() bool {
	if o != nil && !IsNil(o.CycleFlexWithholdConfig) {
		return true
	}

	return false
}

// SetCycleFlexWithholdConfig gets a reference to the given ZMGOCycleFlexConfig and assigns it to the CycleFlexWithholdConfig field.
func (o *ZMGOSettlementConfig) SetCycleFlexWithholdConfig(v ZMGOCycleFlexConfig) {
	o.CycleFlexWithholdConfig = &v
}

// GetCycleWithholdConfig returns the CycleWithholdConfig field value if set, zero value otherwise.
func (o *ZMGOSettlementConfig) GetCycleWithholdConfig() ZMGOCycleWithholdConfig {
	if o == nil || IsNil(o.CycleWithholdConfig) {
		var ret ZMGOCycleWithholdConfig
		return ret
	}
	return *o.CycleWithholdConfig
}

// GetCycleWithholdConfigOk returns a tuple with the CycleWithholdConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOSettlementConfig) GetCycleWithholdConfigOk() (*ZMGOCycleWithholdConfig, bool) {
	if o == nil || IsNil(o.CycleWithholdConfig) {
		return nil, false
	}
	return o.CycleWithholdConfig, true
}

// HasCycleWithholdConfig returns a boolean if a field has been set.
func (o *ZMGOSettlementConfig) HasCycleWithholdConfig() bool {
	if o != nil && !IsNil(o.CycleWithholdConfig) {
		return true
	}

	return false
}

// SetCycleWithholdConfig gets a reference to the given ZMGOCycleWithholdConfig and assigns it to the CycleWithholdConfig field.
func (o *ZMGOSettlementConfig) SetCycleWithholdConfig(v ZMGOCycleWithholdConfig) {
	o.CycleWithholdConfig = &v
}

// GetExpStopDelayDays returns the ExpStopDelayDays field value if set, zero value otherwise.
func (o *ZMGOSettlementConfig) GetExpStopDelayDays() int32 {
	if o == nil || IsNil(o.ExpStopDelayDays) {
		var ret int32
		return ret
	}
	return *o.ExpStopDelayDays
}

// GetExpStopDelayDaysOk returns a tuple with the ExpStopDelayDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOSettlementConfig) GetExpStopDelayDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpStopDelayDays) {
		return nil, false
	}
	return o.ExpStopDelayDays, true
}

// HasExpStopDelayDays returns a boolean if a field has been set.
func (o *ZMGOSettlementConfig) HasExpStopDelayDays() bool {
	if o != nil && !IsNil(o.ExpStopDelayDays) {
		return true
	}

	return false
}

// SetExpStopDelayDays gets a reference to the given int32 and assigns it to the ExpStopDelayDays field.
func (o *ZMGOSettlementConfig) SetExpStopDelayDays(v int32) {
	o.ExpStopDelayDays = &v
}

// GetExpStopTime returns the ExpStopTime field value if set, zero value otherwise.
func (o *ZMGOSettlementConfig) GetExpStopTime() string {
	if o == nil || IsNil(o.ExpStopTime) {
		var ret string
		return ret
	}
	return *o.ExpStopTime
}

// GetExpStopTimeOk returns a tuple with the ExpStopTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOSettlementConfig) GetExpStopTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpStopTime) {
		return nil, false
	}
	return o.ExpStopTime, true
}

// HasExpStopTime returns a boolean if a field has been set.
func (o *ZMGOSettlementConfig) HasExpStopTime() bool {
	if o != nil && !IsNil(o.ExpStopTime) {
		return true
	}

	return false
}

// SetExpStopTime gets a reference to the given string and assigns it to the ExpStopTime field.
func (o *ZMGOSettlementConfig) SetExpStopTime(v string) {
	o.ExpStopTime = &v
}

// GetExpStopTimeMode returns the ExpStopTimeMode field value if set, zero value otherwise.
func (o *ZMGOSettlementConfig) GetExpStopTimeMode() string {
	if o == nil || IsNil(o.ExpStopTimeMode) {
		var ret string
		return ret
	}
	return *o.ExpStopTimeMode
}

// GetExpStopTimeModeOk returns a tuple with the ExpStopTimeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOSettlementConfig) GetExpStopTimeModeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpStopTimeMode) {
		return nil, false
	}
	return o.ExpStopTimeMode, true
}

// HasExpStopTimeMode returns a boolean if a field has been set.
func (o *ZMGOSettlementConfig) HasExpStopTimeMode() bool {
	if o != nil && !IsNil(o.ExpStopTimeMode) {
		return true
	}

	return false
}

// SetExpStopTimeMode gets a reference to the given string and assigns it to the ExpStopTimeMode field.
func (o *ZMGOSettlementConfig) SetExpStopTimeMode(v string) {
	o.ExpStopTimeMode = &v
}

// GetFulfilltimesCustomSettlementPlan returns the FulfilltimesCustomSettlementPlan field value if set, zero value otherwise.
func (o *ZMGOSettlementConfig) GetFulfilltimesCustomSettlementPlan() string {
	if o == nil || IsNil(o.FulfilltimesCustomSettlementPlan) {
		var ret string
		return ret
	}
	return *o.FulfilltimesCustomSettlementPlan
}

// GetFulfilltimesCustomSettlementPlanOk returns a tuple with the FulfilltimesCustomSettlementPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOSettlementConfig) GetFulfilltimesCustomSettlementPlanOk() (*string, bool) {
	if o == nil || IsNil(o.FulfilltimesCustomSettlementPlan) {
		return nil, false
	}
	return o.FulfilltimesCustomSettlementPlan, true
}

// HasFulfilltimesCustomSettlementPlan returns a boolean if a field has been set.
func (o *ZMGOSettlementConfig) HasFulfilltimesCustomSettlementPlan() bool {
	if o != nil && !IsNil(o.FulfilltimesCustomSettlementPlan) {
		return true
	}

	return false
}

// SetFulfilltimesCustomSettlementPlan gets a reference to the given string and assigns it to the FulfilltimesCustomSettlementPlan field.
func (o *ZMGOSettlementConfig) SetFulfilltimesCustomSettlementPlan(v string) {
	o.FulfilltimesCustomSettlementPlan = &v
}

// GetSettlementType returns the SettlementType field value if set, zero value otherwise.
func (o *ZMGOSettlementConfig) GetSettlementType() string {
	if o == nil || IsNil(o.SettlementType) {
		var ret string
		return ret
	}
	return *o.SettlementType
}

// GetSettlementTypeOk returns a tuple with the SettlementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOSettlementConfig) GetSettlementTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SettlementType) {
		return nil, false
	}
	return o.SettlementType, true
}

// HasSettlementType returns a boolean if a field has been set.
func (o *ZMGOSettlementConfig) HasSettlementType() bool {
	if o != nil && !IsNil(o.SettlementType) {
		return true
	}

	return false
}

// SetSettlementType gets a reference to the given string and assigns it to the SettlementType field.
func (o *ZMGOSettlementConfig) SetSettlementType(v string) {
	o.SettlementType = &v
}

func (o ZMGOSettlementConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZMGOSettlementConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomFeeName) {
		toSerialize["custom_fee_name"] = o.CustomFeeName
	}
	if !IsNil(o.CycleFlexWithholdConfig) {
		toSerialize["cycle_flex_withhold_config"] = o.CycleFlexWithholdConfig
	}
	if !IsNil(o.CycleWithholdConfig) {
		toSerialize["cycle_withhold_config"] = o.CycleWithholdConfig
	}
	if !IsNil(o.ExpStopDelayDays) {
		toSerialize["exp_stop_delay_days"] = o.ExpStopDelayDays
	}
	if !IsNil(o.ExpStopTime) {
		toSerialize["exp_stop_time"] = o.ExpStopTime
	}
	if !IsNil(o.ExpStopTimeMode) {
		toSerialize["exp_stop_time_mode"] = o.ExpStopTimeMode
	}
	if !IsNil(o.FulfilltimesCustomSettlementPlan) {
		toSerialize["fulfilltimes_custom_settlement_plan"] = o.FulfilltimesCustomSettlementPlan
	}
	if !IsNil(o.SettlementType) {
		toSerialize["settlement_type"] = o.SettlementType
	}
	return toSerialize, nil
}

type NullableZMGOSettlementConfig struct {
	value *ZMGOSettlementConfig
	isSet bool
}

func (v NullableZMGOSettlementConfig) Get() *ZMGOSettlementConfig {
	return v.value
}

func (v *NullableZMGOSettlementConfig) Set(val *ZMGOSettlementConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableZMGOSettlementConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableZMGOSettlementConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZMGOSettlementConfig(val *ZMGOSettlementConfig) *NullableZMGOSettlementConfig {
	return &NullableZMGOSettlementConfig{value: val, isSet: true}
}

func (v NullableZMGOSettlementConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZMGOSettlementConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

