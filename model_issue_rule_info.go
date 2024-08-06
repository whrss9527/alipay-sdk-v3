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

// checks if the IssueRuleInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueRuleInfo{}

// IssueRuleInfo struct for IssueRuleInfo
type IssueRuleInfo struct {
	// 生效时间段
	EffectivePeriod *string `json:"effective_period,omitempty"`
	// 累计类型，默认为0 可选值：0（不可累计）、1（可累计）、2（累计天数）、3（累计到指定日期）
	InvalidMode *int32 `json:"invalid_mode,omitempty"`
	// 累计类型值
	InvalidModeValue *string `json:"invalid_mode_value,omitempty"`
	// 发放金额，单位元
	IssueAmountValue *string `json:"issue_amount_value,omitempty"`
	// 发放规则有效结束时间
	IssueEndDate *string `json:"issue_end_date,omitempty"`
	// 发放规则id
	IssueRuleId *string `json:"issue_rule_id,omitempty"`
	// 发放规则名称
	IssueRuleName *string `json:"issue_rule_name,omitempty"`
	// 发放规则有效起始时间
	IssueStartDate *string `json:"issue_start_date,omitempty"`
	// 发放类型
	IssueType *string `json:"issue_type,omitempty"`
	// 外部发放规则id
	OuterSourceId *string `json:"outer_source_id,omitempty"`
	// 额度类型
	QuotaType *string `json:"quota_type,omitempty"`
	// 是否可转赠
	ShareMode *int32 `json:"share_mode,omitempty"`
	// 目标id
	TargetId *string `json:"target_id,omitempty"`
	// 发放规则归属的目标类型
	TargetType *string `json:"target_type,omitempty"`
}

// NewIssueRuleInfo instantiates a new IssueRuleInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueRuleInfo() *IssueRuleInfo {
	this := IssueRuleInfo{}
	return &this
}

// NewIssueRuleInfoWithDefaults instantiates a new IssueRuleInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueRuleInfoWithDefaults() *IssueRuleInfo {
	this := IssueRuleInfo{}
	return &this
}

// GetEffectivePeriod returns the EffectivePeriod field value if set, zero value otherwise.
func (o *IssueRuleInfo) GetEffectivePeriod() string {
	if o == nil || IsNil(o.EffectivePeriod) {
		var ret string
		return ret
	}
	return *o.EffectivePeriod
}

// GetEffectivePeriodOk returns a tuple with the EffectivePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRuleInfo) GetEffectivePeriodOk() (*string, bool) {
	if o == nil || IsNil(o.EffectivePeriod) {
		return nil, false
	}
	return o.EffectivePeriod, true
}

// HasEffectivePeriod returns a boolean if a field has been set.
func (o *IssueRuleInfo) HasEffectivePeriod() bool {
	if o != nil && !IsNil(o.EffectivePeriod) {
		return true
	}

	return false
}

// SetEffectivePeriod gets a reference to the given string and assigns it to the EffectivePeriod field.
func (o *IssueRuleInfo) SetEffectivePeriod(v string) {
	o.EffectivePeriod = &v
}

// GetInvalidMode returns the InvalidMode field value if set, zero value otherwise.
func (o *IssueRuleInfo) GetInvalidMode() int32 {
	if o == nil || IsNil(o.InvalidMode) {
		var ret int32
		return ret
	}
	return *o.InvalidMode
}

// GetInvalidModeOk returns a tuple with the InvalidMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRuleInfo) GetInvalidModeOk() (*int32, bool) {
	if o == nil || IsNil(o.InvalidMode) {
		return nil, false
	}
	return o.InvalidMode, true
}

// HasInvalidMode returns a boolean if a field has been set.
func (o *IssueRuleInfo) HasInvalidMode() bool {
	if o != nil && !IsNil(o.InvalidMode) {
		return true
	}

	return false
}

// SetInvalidMode gets a reference to the given int32 and assigns it to the InvalidMode field.
func (o *IssueRuleInfo) SetInvalidMode(v int32) {
	o.InvalidMode = &v
}

// GetInvalidModeValue returns the InvalidModeValue field value if set, zero value otherwise.
func (o *IssueRuleInfo) GetInvalidModeValue() string {
	if o == nil || IsNil(o.InvalidModeValue) {
		var ret string
		return ret
	}
	return *o.InvalidModeValue
}

// GetInvalidModeValueOk returns a tuple with the InvalidModeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRuleInfo) GetInvalidModeValueOk() (*string, bool) {
	if o == nil || IsNil(o.InvalidModeValue) {
		return nil, false
	}
	return o.InvalidModeValue, true
}

// HasInvalidModeValue returns a boolean if a field has been set.
func (o *IssueRuleInfo) HasInvalidModeValue() bool {
	if o != nil && !IsNil(o.InvalidModeValue) {
		return true
	}

	return false
}

// SetInvalidModeValue gets a reference to the given string and assigns it to the InvalidModeValue field.
func (o *IssueRuleInfo) SetInvalidModeValue(v string) {
	o.InvalidModeValue = &v
}

// GetIssueAmountValue returns the IssueAmountValue field value if set, zero value otherwise.
func (o *IssueRuleInfo) GetIssueAmountValue() string {
	if o == nil || IsNil(o.IssueAmountValue) {
		var ret string
		return ret
	}
	return *o.IssueAmountValue
}

// GetIssueAmountValueOk returns a tuple with the IssueAmountValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRuleInfo) GetIssueAmountValueOk() (*string, bool) {
	if o == nil || IsNil(o.IssueAmountValue) {
		return nil, false
	}
	return o.IssueAmountValue, true
}

// HasIssueAmountValue returns a boolean if a field has been set.
func (o *IssueRuleInfo) HasIssueAmountValue() bool {
	if o != nil && !IsNil(o.IssueAmountValue) {
		return true
	}

	return false
}

// SetIssueAmountValue gets a reference to the given string and assigns it to the IssueAmountValue field.
func (o *IssueRuleInfo) SetIssueAmountValue(v string) {
	o.IssueAmountValue = &v
}

// GetIssueEndDate returns the IssueEndDate field value if set, zero value otherwise.
func (o *IssueRuleInfo) GetIssueEndDate() string {
	if o == nil || IsNil(o.IssueEndDate) {
		var ret string
		return ret
	}
	return *o.IssueEndDate
}

// GetIssueEndDateOk returns a tuple with the IssueEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRuleInfo) GetIssueEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.IssueEndDate) {
		return nil, false
	}
	return o.IssueEndDate, true
}

// HasIssueEndDate returns a boolean if a field has been set.
func (o *IssueRuleInfo) HasIssueEndDate() bool {
	if o != nil && !IsNil(o.IssueEndDate) {
		return true
	}

	return false
}

// SetIssueEndDate gets a reference to the given string and assigns it to the IssueEndDate field.
func (o *IssueRuleInfo) SetIssueEndDate(v string) {
	o.IssueEndDate = &v
}

// GetIssueRuleId returns the IssueRuleId field value if set, zero value otherwise.
func (o *IssueRuleInfo) GetIssueRuleId() string {
	if o == nil || IsNil(o.IssueRuleId) {
		var ret string
		return ret
	}
	return *o.IssueRuleId
}

// GetIssueRuleIdOk returns a tuple with the IssueRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRuleInfo) GetIssueRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.IssueRuleId) {
		return nil, false
	}
	return o.IssueRuleId, true
}

// HasIssueRuleId returns a boolean if a field has been set.
func (o *IssueRuleInfo) HasIssueRuleId() bool {
	if o != nil && !IsNil(o.IssueRuleId) {
		return true
	}

	return false
}

// SetIssueRuleId gets a reference to the given string and assigns it to the IssueRuleId field.
func (o *IssueRuleInfo) SetIssueRuleId(v string) {
	o.IssueRuleId = &v
}

// GetIssueRuleName returns the IssueRuleName field value if set, zero value otherwise.
func (o *IssueRuleInfo) GetIssueRuleName() string {
	if o == nil || IsNil(o.IssueRuleName) {
		var ret string
		return ret
	}
	return *o.IssueRuleName
}

// GetIssueRuleNameOk returns a tuple with the IssueRuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRuleInfo) GetIssueRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.IssueRuleName) {
		return nil, false
	}
	return o.IssueRuleName, true
}

// HasIssueRuleName returns a boolean if a field has been set.
func (o *IssueRuleInfo) HasIssueRuleName() bool {
	if o != nil && !IsNil(o.IssueRuleName) {
		return true
	}

	return false
}

// SetIssueRuleName gets a reference to the given string and assigns it to the IssueRuleName field.
func (o *IssueRuleInfo) SetIssueRuleName(v string) {
	o.IssueRuleName = &v
}

// GetIssueStartDate returns the IssueStartDate field value if set, zero value otherwise.
func (o *IssueRuleInfo) GetIssueStartDate() string {
	if o == nil || IsNil(o.IssueStartDate) {
		var ret string
		return ret
	}
	return *o.IssueStartDate
}

// GetIssueStartDateOk returns a tuple with the IssueStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRuleInfo) GetIssueStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.IssueStartDate) {
		return nil, false
	}
	return o.IssueStartDate, true
}

// HasIssueStartDate returns a boolean if a field has been set.
func (o *IssueRuleInfo) HasIssueStartDate() bool {
	if o != nil && !IsNil(o.IssueStartDate) {
		return true
	}

	return false
}

// SetIssueStartDate gets a reference to the given string and assigns it to the IssueStartDate field.
func (o *IssueRuleInfo) SetIssueStartDate(v string) {
	o.IssueStartDate = &v
}

// GetIssueType returns the IssueType field value if set, zero value otherwise.
func (o *IssueRuleInfo) GetIssueType() string {
	if o == nil || IsNil(o.IssueType) {
		var ret string
		return ret
	}
	return *o.IssueType
}

// GetIssueTypeOk returns a tuple with the IssueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRuleInfo) GetIssueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IssueType) {
		return nil, false
	}
	return o.IssueType, true
}

// HasIssueType returns a boolean if a field has been set.
func (o *IssueRuleInfo) HasIssueType() bool {
	if o != nil && !IsNil(o.IssueType) {
		return true
	}

	return false
}

// SetIssueType gets a reference to the given string and assigns it to the IssueType field.
func (o *IssueRuleInfo) SetIssueType(v string) {
	o.IssueType = &v
}

// GetOuterSourceId returns the OuterSourceId field value if set, zero value otherwise.
func (o *IssueRuleInfo) GetOuterSourceId() string {
	if o == nil || IsNil(o.OuterSourceId) {
		var ret string
		return ret
	}
	return *o.OuterSourceId
}

// GetOuterSourceIdOk returns a tuple with the OuterSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRuleInfo) GetOuterSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.OuterSourceId) {
		return nil, false
	}
	return o.OuterSourceId, true
}

// HasOuterSourceId returns a boolean if a field has been set.
func (o *IssueRuleInfo) HasOuterSourceId() bool {
	if o != nil && !IsNil(o.OuterSourceId) {
		return true
	}

	return false
}

// SetOuterSourceId gets a reference to the given string and assigns it to the OuterSourceId field.
func (o *IssueRuleInfo) SetOuterSourceId(v string) {
	o.OuterSourceId = &v
}

// GetQuotaType returns the QuotaType field value if set, zero value otherwise.
func (o *IssueRuleInfo) GetQuotaType() string {
	if o == nil || IsNil(o.QuotaType) {
		var ret string
		return ret
	}
	return *o.QuotaType
}

// GetQuotaTypeOk returns a tuple with the QuotaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRuleInfo) GetQuotaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.QuotaType) {
		return nil, false
	}
	return o.QuotaType, true
}

// HasQuotaType returns a boolean if a field has been set.
func (o *IssueRuleInfo) HasQuotaType() bool {
	if o != nil && !IsNil(o.QuotaType) {
		return true
	}

	return false
}

// SetQuotaType gets a reference to the given string and assigns it to the QuotaType field.
func (o *IssueRuleInfo) SetQuotaType(v string) {
	o.QuotaType = &v
}

// GetShareMode returns the ShareMode field value if set, zero value otherwise.
func (o *IssueRuleInfo) GetShareMode() int32 {
	if o == nil || IsNil(o.ShareMode) {
		var ret int32
		return ret
	}
	return *o.ShareMode
}

// GetShareModeOk returns a tuple with the ShareMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRuleInfo) GetShareModeOk() (*int32, bool) {
	if o == nil || IsNil(o.ShareMode) {
		return nil, false
	}
	return o.ShareMode, true
}

// HasShareMode returns a boolean if a field has been set.
func (o *IssueRuleInfo) HasShareMode() bool {
	if o != nil && !IsNil(o.ShareMode) {
		return true
	}

	return false
}

// SetShareMode gets a reference to the given int32 and assigns it to the ShareMode field.
func (o *IssueRuleInfo) SetShareMode(v int32) {
	o.ShareMode = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *IssueRuleInfo) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRuleInfo) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *IssueRuleInfo) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *IssueRuleInfo) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *IssueRuleInfo) GetTargetType() string {
	if o == nil || IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRuleInfo) GetTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *IssueRuleInfo) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *IssueRuleInfo) SetTargetType(v string) {
	o.TargetType = &v
}

func (o IssueRuleInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueRuleInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EffectivePeriod) {
		toSerialize["effective_period"] = o.EffectivePeriod
	}
	if !IsNil(o.InvalidMode) {
		toSerialize["invalid_mode"] = o.InvalidMode
	}
	if !IsNil(o.InvalidModeValue) {
		toSerialize["invalid_mode_value"] = o.InvalidModeValue
	}
	if !IsNil(o.IssueAmountValue) {
		toSerialize["issue_amount_value"] = o.IssueAmountValue
	}
	if !IsNil(o.IssueEndDate) {
		toSerialize["issue_end_date"] = o.IssueEndDate
	}
	if !IsNil(o.IssueRuleId) {
		toSerialize["issue_rule_id"] = o.IssueRuleId
	}
	if !IsNil(o.IssueRuleName) {
		toSerialize["issue_rule_name"] = o.IssueRuleName
	}
	if !IsNil(o.IssueStartDate) {
		toSerialize["issue_start_date"] = o.IssueStartDate
	}
	if !IsNil(o.IssueType) {
		toSerialize["issue_type"] = o.IssueType
	}
	if !IsNil(o.OuterSourceId) {
		toSerialize["outer_source_id"] = o.OuterSourceId
	}
	if !IsNil(o.QuotaType) {
		toSerialize["quota_type"] = o.QuotaType
	}
	if !IsNil(o.ShareMode) {
		toSerialize["share_mode"] = o.ShareMode
	}
	if !IsNil(o.TargetId) {
		toSerialize["target_id"] = o.TargetId
	}
	if !IsNil(o.TargetType) {
		toSerialize["target_type"] = o.TargetType
	}
	return toSerialize, nil
}

type NullableIssueRuleInfo struct {
	value *IssueRuleInfo
	isSet bool
}

func (v NullableIssueRuleInfo) Get() *IssueRuleInfo {
	return v.value
}

func (v *NullableIssueRuleInfo) Set(val *IssueRuleInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueRuleInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueRuleInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueRuleInfo(val *IssueRuleInfo) *NullableIssueRuleInfo {
	return &NullableIssueRuleInfo{value: val, isSet: true}
}

func (v NullableIssueRuleInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueRuleInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

