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

// checks if the AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel{}

// AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel struct for AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel
type AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel struct {
	// 企业共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// 授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 开票规则ID
	InvoiceRuleId *string `json:"invoice_rule_id,omitempty"`
	// 开票规则名称
	InvoiceRuleName *string `json:"invoice_rule_name,omitempty"`
	// 发票抬头
	InvoiceTitleId *string `json:"invoice_title_id,omitempty"`
	// 收件人地址
	ReceiveAddress *string `json:"receive_address,omitempty"`
	// 收件人姓名
	ReceiveName *string `json:"receive_name,omitempty"`
	// 收件人手机号
	ReceivePhone *string `json:"receive_phone,omitempty"`
	// 销方类型
	SellerType *string `json:"seller_type,omitempty"`
}

// NewAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel instantiates a new AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel() *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel {
	this := AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel{}
	return &this
}

// NewAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModelWithDefaults instantiates a new AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModelWithDefaults() *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel {
	this := AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetInvoiceRuleId returns the InvoiceRuleId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetInvoiceRuleId() string {
	if o == nil || IsNil(o.InvoiceRuleId) {
		var ret string
		return ret
	}
	return *o.InvoiceRuleId
}

// GetInvoiceRuleIdOk returns a tuple with the InvoiceRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetInvoiceRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceRuleId) {
		return nil, false
	}
	return o.InvoiceRuleId, true
}

// HasInvoiceRuleId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) HasInvoiceRuleId() bool {
	if o != nil && !IsNil(o.InvoiceRuleId) {
		return true
	}

	return false
}

// SetInvoiceRuleId gets a reference to the given string and assigns it to the InvoiceRuleId field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) SetInvoiceRuleId(v string) {
	o.InvoiceRuleId = &v
}

// GetInvoiceRuleName returns the InvoiceRuleName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetInvoiceRuleName() string {
	if o == nil || IsNil(o.InvoiceRuleName) {
		var ret string
		return ret
	}
	return *o.InvoiceRuleName
}

// GetInvoiceRuleNameOk returns a tuple with the InvoiceRuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetInvoiceRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceRuleName) {
		return nil, false
	}
	return o.InvoiceRuleName, true
}

// HasInvoiceRuleName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) HasInvoiceRuleName() bool {
	if o != nil && !IsNil(o.InvoiceRuleName) {
		return true
	}

	return false
}

// SetInvoiceRuleName gets a reference to the given string and assigns it to the InvoiceRuleName field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) SetInvoiceRuleName(v string) {
	o.InvoiceRuleName = &v
}

// GetInvoiceTitleId returns the InvoiceTitleId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetInvoiceTitleId() string {
	if o == nil || IsNil(o.InvoiceTitleId) {
		var ret string
		return ret
	}
	return *o.InvoiceTitleId
}

// GetInvoiceTitleIdOk returns a tuple with the InvoiceTitleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetInvoiceTitleIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceTitleId) {
		return nil, false
	}
	return o.InvoiceTitleId, true
}

// HasInvoiceTitleId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) HasInvoiceTitleId() bool {
	if o != nil && !IsNil(o.InvoiceTitleId) {
		return true
	}

	return false
}

// SetInvoiceTitleId gets a reference to the given string and assigns it to the InvoiceTitleId field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) SetInvoiceTitleId(v string) {
	o.InvoiceTitleId = &v
}

// GetReceiveAddress returns the ReceiveAddress field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetReceiveAddress() string {
	if o == nil || IsNil(o.ReceiveAddress) {
		var ret string
		return ret
	}
	return *o.ReceiveAddress
}

// GetReceiveAddressOk returns a tuple with the ReceiveAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetReceiveAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiveAddress) {
		return nil, false
	}
	return o.ReceiveAddress, true
}

// HasReceiveAddress returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) HasReceiveAddress() bool {
	if o != nil && !IsNil(o.ReceiveAddress) {
		return true
	}

	return false
}

// SetReceiveAddress gets a reference to the given string and assigns it to the ReceiveAddress field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) SetReceiveAddress(v string) {
	o.ReceiveAddress = &v
}

// GetReceiveName returns the ReceiveName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetReceiveName() string {
	if o == nil || IsNil(o.ReceiveName) {
		var ret string
		return ret
	}
	return *o.ReceiveName
}

// GetReceiveNameOk returns a tuple with the ReceiveName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetReceiveNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiveName) {
		return nil, false
	}
	return o.ReceiveName, true
}

// HasReceiveName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) HasReceiveName() bool {
	if o != nil && !IsNil(o.ReceiveName) {
		return true
	}

	return false
}

// SetReceiveName gets a reference to the given string and assigns it to the ReceiveName field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) SetReceiveName(v string) {
	o.ReceiveName = &v
}

// GetReceivePhone returns the ReceivePhone field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetReceivePhone() string {
	if o == nil || IsNil(o.ReceivePhone) {
		var ret string
		return ret
	}
	return *o.ReceivePhone
}

// GetReceivePhoneOk returns a tuple with the ReceivePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetReceivePhoneOk() (*string, bool) {
	if o == nil || IsNil(o.ReceivePhone) {
		return nil, false
	}
	return o.ReceivePhone, true
}

// HasReceivePhone returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) HasReceivePhone() bool {
	if o != nil && !IsNil(o.ReceivePhone) {
		return true
	}

	return false
}

// SetReceivePhone gets a reference to the given string and assigns it to the ReceivePhone field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) SetReceivePhone(v string) {
	o.ReceivePhone = &v
}

// GetSellerType returns the SellerType field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetSellerType() string {
	if o == nil || IsNil(o.SellerType) {
		var ret string
		return ret
	}
	return *o.SellerType
}

// GetSellerTypeOk returns a tuple with the SellerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) GetSellerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SellerType) {
		return nil, false
	}
	return o.SellerType, true
}

// HasSellerType returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) HasSellerType() bool {
	if o != nil && !IsNil(o.SellerType) {
		return true
	}

	return false
}

// SetSellerType gets a reference to the given string and assigns it to the SellerType field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) SetSellerType(v string) {
	o.SellerType = &v
}

func (o AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.InvoiceRuleId) {
		toSerialize["invoice_rule_id"] = o.InvoiceRuleId
	}
	if !IsNil(o.InvoiceRuleName) {
		toSerialize["invoice_rule_name"] = o.InvoiceRuleName
	}
	if !IsNil(o.InvoiceTitleId) {
		toSerialize["invoice_title_id"] = o.InvoiceTitleId
	}
	if !IsNil(o.ReceiveAddress) {
		toSerialize["receive_address"] = o.ReceiveAddress
	}
	if !IsNil(o.ReceiveName) {
		toSerialize["receive_name"] = o.ReceiveName
	}
	if !IsNil(o.ReceivePhone) {
		toSerialize["receive_phone"] = o.ReceivePhone
	}
	if !IsNil(o.SellerType) {
		toSerialize["seller_type"] = o.SellerType
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel struct {
	value *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) Get() *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) Set(val *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel(val *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) *NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel {
	return &NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
