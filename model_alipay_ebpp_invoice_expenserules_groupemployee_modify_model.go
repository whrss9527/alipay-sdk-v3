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

// checks if the AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel{}

// AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel struct for AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel
type AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel struct {
	// 企业共同账户id
	AccountId *string `json:"account_id,omitempty"`
	// 未切换 open_id 时请使用此字段： 需要添加的员工UID列表 特殊说明：一次最多50个
	AddEmployeeList []string `json:"add_employee_list,omitempty"`
	// 切换 open_id 后请使用此字段： 需要添加的open_id/企业码员工ID列表 特殊说明：一次最多50个
	AddEmployeeOpenIdList []string `json:"add_employee_open_id_list,omitempty"`
	// 授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 企业码企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 费控规则ID列表
	GroupIdList []string `json:"group_id_list,omitempty"`
	// 未切换 open_id 时请使用此字段： 需要移除的员工UID列表 特殊说明：一次最多50个
	RemoveEmployeeList []string `json:"remove_employee_list,omitempty"`
	// 切换 open_id 后请使用此字段： 需要移除的open_id/企业码员工ID列表 特殊说明：一次最多50个
	RemoveEmployeeOpenIdList []string `json:"remove_employee_open_id_list,omitempty"`
}

// NewAlipayEbppInvoiceExpenserulesGroupemployeeModifyModel instantiates a new AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceExpenserulesGroupemployeeModifyModel() *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel {
	this := AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel{}
	return &this
}

// NewAlipayEbppInvoiceExpenserulesGroupemployeeModifyModelWithDefaults instantiates a new AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceExpenserulesGroupemployeeModifyModelWithDefaults() *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel {
	this := AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAddEmployeeList returns the AddEmployeeList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) GetAddEmployeeList() []string {
	if o == nil || IsNil(o.AddEmployeeList) {
		var ret []string
		return ret
	}
	return o.AddEmployeeList
}

// GetAddEmployeeListOk returns a tuple with the AddEmployeeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) GetAddEmployeeListOk() ([]string, bool) {
	if o == nil || IsNil(o.AddEmployeeList) {
		return nil, false
	}
	return o.AddEmployeeList, true
}

// HasAddEmployeeList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) HasAddEmployeeList() bool {
	if o != nil && !IsNil(o.AddEmployeeList) {
		return true
	}

	return false
}

// SetAddEmployeeList gets a reference to the given []string and assigns it to the AddEmployeeList field.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) SetAddEmployeeList(v []string) {
	o.AddEmployeeList = v
}

// GetAddEmployeeOpenIdList returns the AddEmployeeOpenIdList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) GetAddEmployeeOpenIdList() []string {
	if o == nil || IsNil(o.AddEmployeeOpenIdList) {
		var ret []string
		return ret
	}
	return o.AddEmployeeOpenIdList
}

// GetAddEmployeeOpenIdListOk returns a tuple with the AddEmployeeOpenIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) GetAddEmployeeOpenIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.AddEmployeeOpenIdList) {
		return nil, false
	}
	return o.AddEmployeeOpenIdList, true
}

// HasAddEmployeeOpenIdList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) HasAddEmployeeOpenIdList() bool {
	if o != nil && !IsNil(o.AddEmployeeOpenIdList) {
		return true
	}

	return false
}

// SetAddEmployeeOpenIdList gets a reference to the given []string and assigns it to the AddEmployeeOpenIdList field.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) SetAddEmployeeOpenIdList(v []string) {
	o.AddEmployeeOpenIdList = v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetGroupIdList returns the GroupIdList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) GetGroupIdList() []string {
	if o == nil || IsNil(o.GroupIdList) {
		var ret []string
		return ret
	}
	return o.GroupIdList
}

// GetGroupIdListOk returns a tuple with the GroupIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) GetGroupIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupIdList) {
		return nil, false
	}
	return o.GroupIdList, true
}

// HasGroupIdList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) HasGroupIdList() bool {
	if o != nil && !IsNil(o.GroupIdList) {
		return true
	}

	return false
}

// SetGroupIdList gets a reference to the given []string and assigns it to the GroupIdList field.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) SetGroupIdList(v []string) {
	o.GroupIdList = v
}

// GetRemoveEmployeeList returns the RemoveEmployeeList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) GetRemoveEmployeeList() []string {
	if o == nil || IsNil(o.RemoveEmployeeList) {
		var ret []string
		return ret
	}
	return o.RemoveEmployeeList
}

// GetRemoveEmployeeListOk returns a tuple with the RemoveEmployeeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) GetRemoveEmployeeListOk() ([]string, bool) {
	if o == nil || IsNil(o.RemoveEmployeeList) {
		return nil, false
	}
	return o.RemoveEmployeeList, true
}

// HasRemoveEmployeeList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) HasRemoveEmployeeList() bool {
	if o != nil && !IsNil(o.RemoveEmployeeList) {
		return true
	}

	return false
}

// SetRemoveEmployeeList gets a reference to the given []string and assigns it to the RemoveEmployeeList field.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) SetRemoveEmployeeList(v []string) {
	o.RemoveEmployeeList = v
}

// GetRemoveEmployeeOpenIdList returns the RemoveEmployeeOpenIdList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) GetRemoveEmployeeOpenIdList() []string {
	if o == nil || IsNil(o.RemoveEmployeeOpenIdList) {
		var ret []string
		return ret
	}
	return o.RemoveEmployeeOpenIdList
}

// GetRemoveEmployeeOpenIdListOk returns a tuple with the RemoveEmployeeOpenIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) GetRemoveEmployeeOpenIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.RemoveEmployeeOpenIdList) {
		return nil, false
	}
	return o.RemoveEmployeeOpenIdList, true
}

// HasRemoveEmployeeOpenIdList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) HasRemoveEmployeeOpenIdList() bool {
	if o != nil && !IsNil(o.RemoveEmployeeOpenIdList) {
		return true
	}

	return false
}

// SetRemoveEmployeeOpenIdList gets a reference to the given []string and assigns it to the RemoveEmployeeOpenIdList field.
func (o *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) SetRemoveEmployeeOpenIdList(v []string) {
	o.RemoveEmployeeOpenIdList = v
}

func (o AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AddEmployeeList) {
		toSerialize["add_employee_list"] = o.AddEmployeeList
	}
	if !IsNil(o.AddEmployeeOpenIdList) {
		toSerialize["add_employee_open_id_list"] = o.AddEmployeeOpenIdList
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if !IsNil(o.GroupIdList) {
		toSerialize["group_id_list"] = o.GroupIdList
	}
	if !IsNil(o.RemoveEmployeeList) {
		toSerialize["remove_employee_list"] = o.RemoveEmployeeList
	}
	if !IsNil(o.RemoveEmployeeOpenIdList) {
		toSerialize["remove_employee_open_id_list"] = o.RemoveEmployeeOpenIdList
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyModel struct {
	value *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) Get() *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) Set(val *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyModel(val *AlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) *NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyModel {
	return &NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

