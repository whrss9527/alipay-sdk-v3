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

// checks if the AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel{}

// AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel struct for AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel
type AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel struct {
	// 企业共同账户id
	AccountId *string `json:"account_id,omitempty"`
	// 授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 消费金额
	Amount *int32 `json:"amount,omitempty"`
	// 交易发生时间
	DealTime *string `json:"deal_time,omitempty"`
	// 员工id
	EmployeeId *string `json:"employee_id,omitempty"`
	// 员工账号类型
	EmployeeIdType *int32 `json:"employee_id_type,omitempty"`
	// 员工开放id
	EmployeeOpenId *string `json:"employee_open_id,omitempty"`
	// 企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 拓展参数
	Extend *string `json:"extend,omitempty"`
	// 交易类型
	IsOffSet *int32 `json:"is_off_set,omitempty"`
	// 外部交易支付单号
	OutSourceId *string `json:"out_source_id,omitempty"`
	// 外部平台编码
	Platform *string `json:"platform,omitempty"`
	// 外部交易退款单号
	RelateNo *string `json:"relate_no,omitempty"`
	// 费控规则ID
	StandardId *string `json:"standard_id,omitempty"`
}

// NewAlipayEbppInvoiceExpensecomsueOutsourceNotifyModel instantiates a new AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceExpensecomsueOutsourceNotifyModel() *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel {
	this := AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel{}
	return &this
}

// NewAlipayEbppInvoiceExpensecomsueOutsourceNotifyModelWithDefaults instantiates a new AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceExpensecomsueOutsourceNotifyModelWithDefaults() *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel {
	this := AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) SetAmount(v int32) {
	o.Amount = &v
}

// GetDealTime returns the DealTime field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetDealTime() string {
	if o == nil || IsNil(o.DealTime) {
		var ret string
		return ret
	}
	return *o.DealTime
}

// GetDealTimeOk returns a tuple with the DealTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetDealTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DealTime) {
		return nil, false
	}
	return o.DealTime, true
}

// HasDealTime returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) HasDealTime() bool {
	if o != nil && !IsNil(o.DealTime) {
		return true
	}

	return false
}

// SetDealTime gets a reference to the given string and assigns it to the DealTime field.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) SetDealTime(v string) {
	o.DealTime = &v
}

// GetEmployeeId returns the EmployeeId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetEmployeeId() string {
	if o == nil || IsNil(o.EmployeeId) {
		var ret string
		return ret
	}
	return *o.EmployeeId
}

// GetEmployeeIdOk returns a tuple with the EmployeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetEmployeeIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmployeeId) {
		return nil, false
	}
	return o.EmployeeId, true
}

// HasEmployeeId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) HasEmployeeId() bool {
	if o != nil && !IsNil(o.EmployeeId) {
		return true
	}

	return false
}

// SetEmployeeId gets a reference to the given string and assigns it to the EmployeeId field.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) SetEmployeeId(v string) {
	o.EmployeeId = &v
}

// GetEmployeeIdType returns the EmployeeIdType field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetEmployeeIdType() int32 {
	if o == nil || IsNil(o.EmployeeIdType) {
		var ret int32
		return ret
	}
	return *o.EmployeeIdType
}

// GetEmployeeIdTypeOk returns a tuple with the EmployeeIdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetEmployeeIdTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.EmployeeIdType) {
		return nil, false
	}
	return o.EmployeeIdType, true
}

// HasEmployeeIdType returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) HasEmployeeIdType() bool {
	if o != nil && !IsNil(o.EmployeeIdType) {
		return true
	}

	return false
}

// SetEmployeeIdType gets a reference to the given int32 and assigns it to the EmployeeIdType field.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) SetEmployeeIdType(v int32) {
	o.EmployeeIdType = &v
}

// GetEmployeeOpenId returns the EmployeeOpenId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetEmployeeOpenId() string {
	if o == nil || IsNil(o.EmployeeOpenId) {
		var ret string
		return ret
	}
	return *o.EmployeeOpenId
}

// GetEmployeeOpenIdOk returns a tuple with the EmployeeOpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetEmployeeOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmployeeOpenId) {
		return nil, false
	}
	return o.EmployeeOpenId, true
}

// HasEmployeeOpenId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) HasEmployeeOpenId() bool {
	if o != nil && !IsNil(o.EmployeeOpenId) {
		return true
	}

	return false
}

// SetEmployeeOpenId gets a reference to the given string and assigns it to the EmployeeOpenId field.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) SetEmployeeOpenId(v string) {
	o.EmployeeOpenId = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetExtend returns the Extend field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetExtend() string {
	if o == nil || IsNil(o.Extend) {
		var ret string
		return ret
	}
	return *o.Extend
}

// GetExtendOk returns a tuple with the Extend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetExtendOk() (*string, bool) {
	if o == nil || IsNil(o.Extend) {
		return nil, false
	}
	return o.Extend, true
}

// HasExtend returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) HasExtend() bool {
	if o != nil && !IsNil(o.Extend) {
		return true
	}

	return false
}

// SetExtend gets a reference to the given string and assigns it to the Extend field.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) SetExtend(v string) {
	o.Extend = &v
}

// GetIsOffSet returns the IsOffSet field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetIsOffSet() int32 {
	if o == nil || IsNil(o.IsOffSet) {
		var ret int32
		return ret
	}
	return *o.IsOffSet
}

// GetIsOffSetOk returns a tuple with the IsOffSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetIsOffSetOk() (*int32, bool) {
	if o == nil || IsNil(o.IsOffSet) {
		return nil, false
	}
	return o.IsOffSet, true
}

// HasIsOffSet returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) HasIsOffSet() bool {
	if o != nil && !IsNil(o.IsOffSet) {
		return true
	}

	return false
}

// SetIsOffSet gets a reference to the given int32 and assigns it to the IsOffSet field.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) SetIsOffSet(v int32) {
	o.IsOffSet = &v
}

// GetOutSourceId returns the OutSourceId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetOutSourceId() string {
	if o == nil || IsNil(o.OutSourceId) {
		var ret string
		return ret
	}
	return *o.OutSourceId
}

// GetOutSourceIdOk returns a tuple with the OutSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetOutSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.OutSourceId) {
		return nil, false
	}
	return o.OutSourceId, true
}

// HasOutSourceId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) HasOutSourceId() bool {
	if o != nil && !IsNil(o.OutSourceId) {
		return true
	}

	return false
}

// SetOutSourceId gets a reference to the given string and assigns it to the OutSourceId field.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) SetOutSourceId(v string) {
	o.OutSourceId = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) SetPlatform(v string) {
	o.Platform = &v
}

// GetRelateNo returns the RelateNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetRelateNo() string {
	if o == nil || IsNil(o.RelateNo) {
		var ret string
		return ret
	}
	return *o.RelateNo
}

// GetRelateNoOk returns a tuple with the RelateNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetRelateNoOk() (*string, bool) {
	if o == nil || IsNil(o.RelateNo) {
		return nil, false
	}
	return o.RelateNo, true
}

// HasRelateNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) HasRelateNo() bool {
	if o != nil && !IsNil(o.RelateNo) {
		return true
	}

	return false
}

// SetRelateNo gets a reference to the given string and assigns it to the RelateNo field.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) SetRelateNo(v string) {
	o.RelateNo = &v
}

// GetStandardId returns the StandardId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetStandardId() string {
	if o == nil || IsNil(o.StandardId) {
		var ret string
		return ret
	}
	return *o.StandardId
}

// GetStandardIdOk returns a tuple with the StandardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) GetStandardIdOk() (*string, bool) {
	if o == nil || IsNil(o.StandardId) {
		return nil, false
	}
	return o.StandardId, true
}

// HasStandardId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) HasStandardId() bool {
	if o != nil && !IsNil(o.StandardId) {
		return true
	}

	return false
}

// SetStandardId gets a reference to the given string and assigns it to the StandardId field.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) SetStandardId(v string) {
	o.StandardId = &v
}

func (o AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.DealTime) {
		toSerialize["deal_time"] = o.DealTime
	}
	if !IsNil(o.EmployeeId) {
		toSerialize["employee_id"] = o.EmployeeId
	}
	if !IsNil(o.EmployeeIdType) {
		toSerialize["employee_id_type"] = o.EmployeeIdType
	}
	if !IsNil(o.EmployeeOpenId) {
		toSerialize["employee_open_id"] = o.EmployeeOpenId
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if !IsNil(o.Extend) {
		toSerialize["extend"] = o.Extend
	}
	if !IsNil(o.IsOffSet) {
		toSerialize["is_off_set"] = o.IsOffSet
	}
	if !IsNil(o.OutSourceId) {
		toSerialize["out_source_id"] = o.OutSourceId
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.RelateNo) {
		toSerialize["relate_no"] = o.RelateNo
	}
	if !IsNil(o.StandardId) {
		toSerialize["standard_id"] = o.StandardId
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyModel struct {
	value *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) Get() *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) Set(val *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyModel(val *AlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) *NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyModel {
	return &NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
