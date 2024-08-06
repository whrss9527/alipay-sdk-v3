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

// checks if the AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel{}

// AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel struct for AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel
type AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel struct {
	// 员工已绑定费控规则列表
	EmployeeRules []ExpenseCtrlEmployeeRuleInfo `json:"employee_rules,omitempty"`
	// 客户端输入的页码
	PageNum *int32 `json:"page_num,omitempty"`
	// 客户端输入的每页行数
	PageSize *int32 `json:"page_size,omitempty"`
	// 查询到的实例总数
	TotalCount *int32 `json:"total_count,omitempty"`
}

// NewAlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel instantiates a new AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel() *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel {
	this := AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModelWithDefaults instantiates a new AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModelWithDefaults() *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel {
	this := AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel{}
	return &this
}

// GetEmployeeRules returns the EmployeeRules field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) GetEmployeeRules() []ExpenseCtrlEmployeeRuleInfo {
	if o == nil || IsNil(o.EmployeeRules) {
		var ret []ExpenseCtrlEmployeeRuleInfo
		return ret
	}
	return o.EmployeeRules
}

// GetEmployeeRulesOk returns a tuple with the EmployeeRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) GetEmployeeRulesOk() ([]ExpenseCtrlEmployeeRuleInfo, bool) {
	if o == nil || IsNil(o.EmployeeRules) {
		return nil, false
	}
	return o.EmployeeRules, true
}

// HasEmployeeRules returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) HasEmployeeRules() bool {
	if o != nil && !IsNil(o.EmployeeRules) {
		return true
	}

	return false
}

// SetEmployeeRules gets a reference to the given []ExpenseCtrlEmployeeRuleInfo and assigns it to the EmployeeRules field.
func (o *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) SetEmployeeRules(v []ExpenseCtrlEmployeeRuleInfo) {
	o.EmployeeRules = v
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) GetPageNum() int32 {
	if o == nil || IsNil(o.PageNum) {
		var ret int32
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) GetPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int32 and assigns it to the PageNum field.
func (o *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) SetPageNum(v int32) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmployeeRules) {
		toSerialize["employee_rules"] = o.EmployeeRules
	}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel struct {
	value *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) Get() *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) Set(val *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel(val *AlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) *NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel {
	return &NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

