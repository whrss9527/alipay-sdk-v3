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

// checks if the StandardRuleInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardRuleInfo{}

// StandardRuleInfo struct for StandardRuleInfo
type StandardRuleInfo struct {
	// 有效期截止
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// 有效期起始
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// 费控条件列表
	ExpenseCtrlRuleInfoList []ExpenseCtrRuleInfo `json:"expense_ctrl_rule_info_list,omitempty"`
	// 当笔消费金额大于规则可用余额时，用于控制支付策略，该字段缺省时采取因公账户和个人账户组合支付策略， 枚举值：PERSONAL（全部个人账户支付）
	PaymentPolicy *string `json:"payment_policy,omitempty"`
	// 费控规则说明
	StandardDesc *string `json:"standard_desc,omitempty"`
	// 费控规则ID
	StandardId *string `json:"standard_id,omitempty"`
	// 费控规则名称
	StandardName *string `json:"standard_name,omitempty"`
}

// NewStandardRuleInfo instantiates a new StandardRuleInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardRuleInfo() *StandardRuleInfo {
	this := StandardRuleInfo{}
	return &this
}

// NewStandardRuleInfoWithDefaults instantiates a new StandardRuleInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardRuleInfoWithDefaults() *StandardRuleInfo {
	this := StandardRuleInfo{}
	return &this
}

// GetEffectiveEndDate returns the EffectiveEndDate field value if set, zero value otherwise.
func (o *StandardRuleInfo) GetEffectiveEndDate() string {
	if o == nil || IsNil(o.EffectiveEndDate) {
		var ret string
		return ret
	}
	return *o.EffectiveEndDate
}

// GetEffectiveEndDateOk returns a tuple with the EffectiveEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardRuleInfo) GetEffectiveEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveEndDate) {
		return nil, false
	}
	return o.EffectiveEndDate, true
}

// HasEffectiveEndDate returns a boolean if a field has been set.
func (o *StandardRuleInfo) HasEffectiveEndDate() bool {
	if o != nil && !IsNil(o.EffectiveEndDate) {
		return true
	}

	return false
}

// SetEffectiveEndDate gets a reference to the given string and assigns it to the EffectiveEndDate field.
func (o *StandardRuleInfo) SetEffectiveEndDate(v string) {
	o.EffectiveEndDate = &v
}

// GetEffectiveStartDate returns the EffectiveStartDate field value if set, zero value otherwise.
func (o *StandardRuleInfo) GetEffectiveStartDate() string {
	if o == nil || IsNil(o.EffectiveStartDate) {
		var ret string
		return ret
	}
	return *o.EffectiveStartDate
}

// GetEffectiveStartDateOk returns a tuple with the EffectiveStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardRuleInfo) GetEffectiveStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveStartDate) {
		return nil, false
	}
	return o.EffectiveStartDate, true
}

// HasEffectiveStartDate returns a boolean if a field has been set.
func (o *StandardRuleInfo) HasEffectiveStartDate() bool {
	if o != nil && !IsNil(o.EffectiveStartDate) {
		return true
	}

	return false
}

// SetEffectiveStartDate gets a reference to the given string and assigns it to the EffectiveStartDate field.
func (o *StandardRuleInfo) SetEffectiveStartDate(v string) {
	o.EffectiveStartDate = &v
}

// GetExpenseCtrlRuleInfoList returns the ExpenseCtrlRuleInfoList field value if set, zero value otherwise.
func (o *StandardRuleInfo) GetExpenseCtrlRuleInfoList() []ExpenseCtrRuleInfo {
	if o == nil || IsNil(o.ExpenseCtrlRuleInfoList) {
		var ret []ExpenseCtrRuleInfo
		return ret
	}
	return o.ExpenseCtrlRuleInfoList
}

// GetExpenseCtrlRuleInfoListOk returns a tuple with the ExpenseCtrlRuleInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardRuleInfo) GetExpenseCtrlRuleInfoListOk() ([]ExpenseCtrRuleInfo, bool) {
	if o == nil || IsNil(o.ExpenseCtrlRuleInfoList) {
		return nil, false
	}
	return o.ExpenseCtrlRuleInfoList, true
}

// HasExpenseCtrlRuleInfoList returns a boolean if a field has been set.
func (o *StandardRuleInfo) HasExpenseCtrlRuleInfoList() bool {
	if o != nil && !IsNil(o.ExpenseCtrlRuleInfoList) {
		return true
	}

	return false
}

// SetExpenseCtrlRuleInfoList gets a reference to the given []ExpenseCtrRuleInfo and assigns it to the ExpenseCtrlRuleInfoList field.
func (o *StandardRuleInfo) SetExpenseCtrlRuleInfoList(v []ExpenseCtrRuleInfo) {
	o.ExpenseCtrlRuleInfoList = v
}

// GetPaymentPolicy returns the PaymentPolicy field value if set, zero value otherwise.
func (o *StandardRuleInfo) GetPaymentPolicy() string {
	if o == nil || IsNil(o.PaymentPolicy) {
		var ret string
		return ret
	}
	return *o.PaymentPolicy
}

// GetPaymentPolicyOk returns a tuple with the PaymentPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardRuleInfo) GetPaymentPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentPolicy) {
		return nil, false
	}
	return o.PaymentPolicy, true
}

// HasPaymentPolicy returns a boolean if a field has been set.
func (o *StandardRuleInfo) HasPaymentPolicy() bool {
	if o != nil && !IsNil(o.PaymentPolicy) {
		return true
	}

	return false
}

// SetPaymentPolicy gets a reference to the given string and assigns it to the PaymentPolicy field.
func (o *StandardRuleInfo) SetPaymentPolicy(v string) {
	o.PaymentPolicy = &v
}

// GetStandardDesc returns the StandardDesc field value if set, zero value otherwise.
func (o *StandardRuleInfo) GetStandardDesc() string {
	if o == nil || IsNil(o.StandardDesc) {
		var ret string
		return ret
	}
	return *o.StandardDesc
}

// GetStandardDescOk returns a tuple with the StandardDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardRuleInfo) GetStandardDescOk() (*string, bool) {
	if o == nil || IsNil(o.StandardDesc) {
		return nil, false
	}
	return o.StandardDesc, true
}

// HasStandardDesc returns a boolean if a field has been set.
func (o *StandardRuleInfo) HasStandardDesc() bool {
	if o != nil && !IsNil(o.StandardDesc) {
		return true
	}

	return false
}

// SetStandardDesc gets a reference to the given string and assigns it to the StandardDesc field.
func (o *StandardRuleInfo) SetStandardDesc(v string) {
	o.StandardDesc = &v
}

// GetStandardId returns the StandardId field value if set, zero value otherwise.
func (o *StandardRuleInfo) GetStandardId() string {
	if o == nil || IsNil(o.StandardId) {
		var ret string
		return ret
	}
	return *o.StandardId
}

// GetStandardIdOk returns a tuple with the StandardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardRuleInfo) GetStandardIdOk() (*string, bool) {
	if o == nil || IsNil(o.StandardId) {
		return nil, false
	}
	return o.StandardId, true
}

// HasStandardId returns a boolean if a field has been set.
func (o *StandardRuleInfo) HasStandardId() bool {
	if o != nil && !IsNil(o.StandardId) {
		return true
	}

	return false
}

// SetStandardId gets a reference to the given string and assigns it to the StandardId field.
func (o *StandardRuleInfo) SetStandardId(v string) {
	o.StandardId = &v
}

// GetStandardName returns the StandardName field value if set, zero value otherwise.
func (o *StandardRuleInfo) GetStandardName() string {
	if o == nil || IsNil(o.StandardName) {
		var ret string
		return ret
	}
	return *o.StandardName
}

// GetStandardNameOk returns a tuple with the StandardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardRuleInfo) GetStandardNameOk() (*string, bool) {
	if o == nil || IsNil(o.StandardName) {
		return nil, false
	}
	return o.StandardName, true
}

// HasStandardName returns a boolean if a field has been set.
func (o *StandardRuleInfo) HasStandardName() bool {
	if o != nil && !IsNil(o.StandardName) {
		return true
	}

	return false
}

// SetStandardName gets a reference to the given string and assigns it to the StandardName field.
func (o *StandardRuleInfo) SetStandardName(v string) {
	o.StandardName = &v
}

func (o StandardRuleInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandardRuleInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EffectiveEndDate) {
		toSerialize["effective_end_date"] = o.EffectiveEndDate
	}
	if !IsNil(o.EffectiveStartDate) {
		toSerialize["effective_start_date"] = o.EffectiveStartDate
	}
	if !IsNil(o.ExpenseCtrlRuleInfoList) {
		toSerialize["expense_ctrl_rule_info_list"] = o.ExpenseCtrlRuleInfoList
	}
	if !IsNil(o.PaymentPolicy) {
		toSerialize["payment_policy"] = o.PaymentPolicy
	}
	if !IsNil(o.StandardDesc) {
		toSerialize["standard_desc"] = o.StandardDesc
	}
	if !IsNil(o.StandardId) {
		toSerialize["standard_id"] = o.StandardId
	}
	if !IsNil(o.StandardName) {
		toSerialize["standard_name"] = o.StandardName
	}
	return toSerialize, nil
}

type NullableStandardRuleInfo struct {
	value *StandardRuleInfo
	isSet bool
}

func (v NullableStandardRuleInfo) Get() *StandardRuleInfo {
	return v.value
}

func (v *NullableStandardRuleInfo) Set(val *StandardRuleInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardRuleInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardRuleInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardRuleInfo(val *StandardRuleInfo) *NullableStandardRuleInfo {
	return &NullableStandardRuleInfo{value: val, isSet: true}
}

func (v NullableStandardRuleInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardRuleInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
