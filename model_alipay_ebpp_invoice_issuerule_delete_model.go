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

// checks if the AlipayEbppInvoiceIssueruleDeleteModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceIssueruleDeleteModel{}

// AlipayEbppInvoiceIssueruleDeleteModel struct for AlipayEbppInvoiceIssueruleDeleteModel
type AlipayEbppInvoiceIssueruleDeleteModel struct {
	// 共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	// Deprecated
	AccountId *string `json:"account_id,omitempty"`
	// 授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	// Deprecated
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 需要删除的发放规则id列表
	IssueRuleIdList []string `json:"issue_rule_id_list,omitempty"`
	// 目标值id
	TargetId *string `json:"target_id,omitempty"`
	// 发放规则关联的目标类型
	TargetType *string `json:"target_type,omitempty"`
}

// NewAlipayEbppInvoiceIssueruleDeleteModel instantiates a new AlipayEbppInvoiceIssueruleDeleteModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceIssueruleDeleteModel() *AlipayEbppInvoiceIssueruleDeleteModel {
	this := AlipayEbppInvoiceIssueruleDeleteModel{}
	return &this
}

// NewAlipayEbppInvoiceIssueruleDeleteModelWithDefaults instantiates a new AlipayEbppInvoiceIssueruleDeleteModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceIssueruleDeleteModelWithDefaults() *AlipayEbppInvoiceIssueruleDeleteModel {
	this := AlipayEbppInvoiceIssueruleDeleteModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
// Deprecated
func (o *AlipayEbppInvoiceIssueruleDeleteModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AlipayEbppInvoiceIssueruleDeleteModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
// Deprecated
func (o *AlipayEbppInvoiceIssueruleDeleteModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
// Deprecated
func (o *AlipayEbppInvoiceIssueruleDeleteModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AlipayEbppInvoiceIssueruleDeleteModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
// Deprecated
func (o *AlipayEbppInvoiceIssueruleDeleteModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetIssueRuleIdList returns the IssueRuleIdList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) GetIssueRuleIdList() []string {
	if o == nil || IsNil(o.IssueRuleIdList) {
		var ret []string
		return ret
	}
	return o.IssueRuleIdList
}

// GetIssueRuleIdListOk returns a tuple with the IssueRuleIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) GetIssueRuleIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.IssueRuleIdList) {
		return nil, false
	}
	return o.IssueRuleIdList, true
}

// HasIssueRuleIdList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) HasIssueRuleIdList() bool {
	if o != nil && !IsNil(o.IssueRuleIdList) {
		return true
	}

	return false
}

// SetIssueRuleIdList gets a reference to the given []string and assigns it to the IssueRuleIdList field.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) SetIssueRuleIdList(v []string) {
	o.IssueRuleIdList = v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) GetTargetType() string {
	if o == nil || IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) GetTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *AlipayEbppInvoiceIssueruleDeleteModel) SetTargetType(v string) {
	o.TargetType = &v
}

func (o AlipayEbppInvoiceIssueruleDeleteModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceIssueruleDeleteModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if !IsNil(o.IssueRuleIdList) {
		toSerialize["issue_rule_id_list"] = o.IssueRuleIdList
	}
	if !IsNil(o.TargetId) {
		toSerialize["target_id"] = o.TargetId
	}
	if !IsNil(o.TargetType) {
		toSerialize["target_type"] = o.TargetType
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceIssueruleDeleteModel struct {
	value *AlipayEbppInvoiceIssueruleDeleteModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceIssueruleDeleteModel) Get() *AlipayEbppInvoiceIssueruleDeleteModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceIssueruleDeleteModel) Set(val *AlipayEbppInvoiceIssueruleDeleteModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceIssueruleDeleteModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceIssueruleDeleteModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceIssueruleDeleteModel(val *AlipayEbppInvoiceIssueruleDeleteModel) *NullableAlipayEbppInvoiceIssueruleDeleteModel {
	return &NullableAlipayEbppInvoiceIssueruleDeleteModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceIssueruleDeleteModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceIssueruleDeleteModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

