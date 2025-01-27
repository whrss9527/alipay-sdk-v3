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

// checks if the EcVoucherInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EcVoucherInfo{}

// EcVoucherInfo struct for EcVoucherInfo
type EcVoucherInfo struct {
	// 共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// 员工ID，汇总发票该字段无效
	EmployeeId *string `json:"employee_id,omitempty"`
	// 企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 员工支付宝UID
	OpenId *string `json:"open_id,omitempty"`
	// 员工支付宝UID
	UserId *string `json:"user_id,omitempty"`
	// 凭证内容
	VoucherContent *string `json:"voucher_content,omitempty"`
	// 凭证创建时间，格式：yyyy-MM-dd HH:mm:ss
	VoucherDate *string `json:"voucher_date,omitempty"`
	// 凭证ID，幂等用
	VoucherId *string `json:"voucher_id,omitempty"`
	// 凭证来源
	VoucherSource *string `json:"voucher_source,omitempty"`
	// 凭证类型
	VoucherType *string `json:"voucher_type,omitempty"`
}

// NewEcVoucherInfo instantiates a new EcVoucherInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEcVoucherInfo() *EcVoucherInfo {
	this := EcVoucherInfo{}
	return &this
}

// NewEcVoucherInfoWithDefaults instantiates a new EcVoucherInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEcVoucherInfoWithDefaults() *EcVoucherInfo {
	this := EcVoucherInfo{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *EcVoucherInfo) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcVoucherInfo) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *EcVoucherInfo) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *EcVoucherInfo) SetAccountId(v string) {
	o.AccountId = &v
}

// GetEmployeeId returns the EmployeeId field value if set, zero value otherwise.
func (o *EcVoucherInfo) GetEmployeeId() string {
	if o == nil || IsNil(o.EmployeeId) {
		var ret string
		return ret
	}
	return *o.EmployeeId
}

// GetEmployeeIdOk returns a tuple with the EmployeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcVoucherInfo) GetEmployeeIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmployeeId) {
		return nil, false
	}
	return o.EmployeeId, true
}

// HasEmployeeId returns a boolean if a field has been set.
func (o *EcVoucherInfo) HasEmployeeId() bool {
	if o != nil && !IsNil(o.EmployeeId) {
		return true
	}

	return false
}

// SetEmployeeId gets a reference to the given string and assigns it to the EmployeeId field.
func (o *EcVoucherInfo) SetEmployeeId(v string) {
	o.EmployeeId = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *EcVoucherInfo) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcVoucherInfo) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *EcVoucherInfo) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *EcVoucherInfo) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *EcVoucherInfo) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcVoucherInfo) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *EcVoucherInfo) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *EcVoucherInfo) SetOpenId(v string) {
	o.OpenId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *EcVoucherInfo) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcVoucherInfo) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *EcVoucherInfo) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *EcVoucherInfo) SetUserId(v string) {
	o.UserId = &v
}

// GetVoucherContent returns the VoucherContent field value if set, zero value otherwise.
func (o *EcVoucherInfo) GetVoucherContent() string {
	if o == nil || IsNil(o.VoucherContent) {
		var ret string
		return ret
	}
	return *o.VoucherContent
}

// GetVoucherContentOk returns a tuple with the VoucherContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcVoucherInfo) GetVoucherContentOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherContent) {
		return nil, false
	}
	return o.VoucherContent, true
}

// HasVoucherContent returns a boolean if a field has been set.
func (o *EcVoucherInfo) HasVoucherContent() bool {
	if o != nil && !IsNil(o.VoucherContent) {
		return true
	}

	return false
}

// SetVoucherContent gets a reference to the given string and assigns it to the VoucherContent field.
func (o *EcVoucherInfo) SetVoucherContent(v string) {
	o.VoucherContent = &v
}

// GetVoucherDate returns the VoucherDate field value if set, zero value otherwise.
func (o *EcVoucherInfo) GetVoucherDate() string {
	if o == nil || IsNil(o.VoucherDate) {
		var ret string
		return ret
	}
	return *o.VoucherDate
}

// GetVoucherDateOk returns a tuple with the VoucherDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcVoucherInfo) GetVoucherDateOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherDate) {
		return nil, false
	}
	return o.VoucherDate, true
}

// HasVoucherDate returns a boolean if a field has been set.
func (o *EcVoucherInfo) HasVoucherDate() bool {
	if o != nil && !IsNil(o.VoucherDate) {
		return true
	}

	return false
}

// SetVoucherDate gets a reference to the given string and assigns it to the VoucherDate field.
func (o *EcVoucherInfo) SetVoucherDate(v string) {
	o.VoucherDate = &v
}

// GetVoucherId returns the VoucherId field value if set, zero value otherwise.
func (o *EcVoucherInfo) GetVoucherId() string {
	if o == nil || IsNil(o.VoucherId) {
		var ret string
		return ret
	}
	return *o.VoucherId
}

// GetVoucherIdOk returns a tuple with the VoucherId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcVoucherInfo) GetVoucherIdOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherId) {
		return nil, false
	}
	return o.VoucherId, true
}

// HasVoucherId returns a boolean if a field has been set.
func (o *EcVoucherInfo) HasVoucherId() bool {
	if o != nil && !IsNil(o.VoucherId) {
		return true
	}

	return false
}

// SetVoucherId gets a reference to the given string and assigns it to the VoucherId field.
func (o *EcVoucherInfo) SetVoucherId(v string) {
	o.VoucherId = &v
}

// GetVoucherSource returns the VoucherSource field value if set, zero value otherwise.
func (o *EcVoucherInfo) GetVoucherSource() string {
	if o == nil || IsNil(o.VoucherSource) {
		var ret string
		return ret
	}
	return *o.VoucherSource
}

// GetVoucherSourceOk returns a tuple with the VoucherSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcVoucherInfo) GetVoucherSourceOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherSource) {
		return nil, false
	}
	return o.VoucherSource, true
}

// HasVoucherSource returns a boolean if a field has been set.
func (o *EcVoucherInfo) HasVoucherSource() bool {
	if o != nil && !IsNil(o.VoucherSource) {
		return true
	}

	return false
}

// SetVoucherSource gets a reference to the given string and assigns it to the VoucherSource field.
func (o *EcVoucherInfo) SetVoucherSource(v string) {
	o.VoucherSource = &v
}

// GetVoucherType returns the VoucherType field value if set, zero value otherwise.
func (o *EcVoucherInfo) GetVoucherType() string {
	if o == nil || IsNil(o.VoucherType) {
		var ret string
		return ret
	}
	return *o.VoucherType
}

// GetVoucherTypeOk returns a tuple with the VoucherType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcVoucherInfo) GetVoucherTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherType) {
		return nil, false
	}
	return o.VoucherType, true
}

// HasVoucherType returns a boolean if a field has been set.
func (o *EcVoucherInfo) HasVoucherType() bool {
	if o != nil && !IsNil(o.VoucherType) {
		return true
	}

	return false
}

// SetVoucherType gets a reference to the given string and assigns it to the VoucherType field.
func (o *EcVoucherInfo) SetVoucherType(v string) {
	o.VoucherType = &v
}

func (o EcVoucherInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EcVoucherInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.EmployeeId) {
		toSerialize["employee_id"] = o.EmployeeId
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.VoucherContent) {
		toSerialize["voucher_content"] = o.VoucherContent
	}
	if !IsNil(o.VoucherDate) {
		toSerialize["voucher_date"] = o.VoucherDate
	}
	if !IsNil(o.VoucherId) {
		toSerialize["voucher_id"] = o.VoucherId
	}
	if !IsNil(o.VoucherSource) {
		toSerialize["voucher_source"] = o.VoucherSource
	}
	if !IsNil(o.VoucherType) {
		toSerialize["voucher_type"] = o.VoucherType
	}
	return toSerialize, nil
}

type NullableEcVoucherInfo struct {
	value *EcVoucherInfo
	isSet bool
}

func (v NullableEcVoucherInfo) Get() *EcVoucherInfo {
	return v.value
}

func (v *NullableEcVoucherInfo) Set(val *EcVoucherInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEcVoucherInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEcVoucherInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcVoucherInfo(val *EcVoucherInfo) *NullableEcVoucherInfo {
	return &NullableEcVoucherInfo{value: val, isSet: true}
}

func (v NullableEcVoucherInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcVoucherInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
