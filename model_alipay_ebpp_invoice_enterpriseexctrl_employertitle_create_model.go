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

// checks if the AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel{}

// AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel struct for AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel
type AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel struct {
	// 企业共同账户id
	AccountId *string `json:"account_id,omitempty"`
	// 企业地址
	Address *string `json:"address,omitempty"`
	// 授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 开户行账号
	BankAccount *string `json:"bank_account,omitempty"`
	// 开户行名称
	BankName *string `json:"bank_name,omitempty"`
	// 企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 税号
	TaxRegisterNo *string `json:"tax_register_no,omitempty"`
	// 电话
	Telephone *string `json:"telephone,omitempty"`
	// 企业抬头名称
	TitleName *string `json:"title_name,omitempty"`
}

// NewAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel instantiates a new AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel() *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel {
	this := AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel{}
	return &this
}

// NewAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModelWithDefaults instantiates a new AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModelWithDefaults() *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel {
	this := AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) SetAddress(v string) {
	o.Address = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetBankAccount() string {
	if o == nil || IsNil(o.BankAccount) {
		var ret string
		return ret
	}
	return *o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetBankAccountOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccount) {
		return nil, false
	}
	return o.BankAccount, true
}

// HasBankAccount returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) HasBankAccount() bool {
	if o != nil && !IsNil(o.BankAccount) {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given string and assigns it to the BankAccount field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) SetBankAccount(v string) {
	o.BankAccount = &v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetBankName() string {
	if o == nil || IsNil(o.BankName) {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.BankName) {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) HasBankName() bool {
	if o != nil && !IsNil(o.BankName) {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) SetBankName(v string) {
	o.BankName = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetTaxRegisterNo returns the TaxRegisterNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetTaxRegisterNo() string {
	if o == nil || IsNil(o.TaxRegisterNo) {
		var ret string
		return ret
	}
	return *o.TaxRegisterNo
}

// GetTaxRegisterNoOk returns a tuple with the TaxRegisterNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetTaxRegisterNoOk() (*string, bool) {
	if o == nil || IsNil(o.TaxRegisterNo) {
		return nil, false
	}
	return o.TaxRegisterNo, true
}

// HasTaxRegisterNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) HasTaxRegisterNo() bool {
	if o != nil && !IsNil(o.TaxRegisterNo) {
		return true
	}

	return false
}

// SetTaxRegisterNo gets a reference to the given string and assigns it to the TaxRegisterNo field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) SetTaxRegisterNo(v string) {
	o.TaxRegisterNo = &v
}

// GetTelephone returns the Telephone field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetTelephone() string {
	if o == nil || IsNil(o.Telephone) {
		var ret string
		return ret
	}
	return *o.Telephone
}

// GetTelephoneOk returns a tuple with the Telephone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetTelephoneOk() (*string, bool) {
	if o == nil || IsNil(o.Telephone) {
		return nil, false
	}
	return o.Telephone, true
}

// HasTelephone returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) HasTelephone() bool {
	if o != nil && !IsNil(o.Telephone) {
		return true
	}

	return false
}

// SetTelephone gets a reference to the given string and assigns it to the Telephone field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) SetTelephone(v string) {
	o.Telephone = &v
}

// GetTitleName returns the TitleName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetTitleName() string {
	if o == nil || IsNil(o.TitleName) {
		var ret string
		return ret
	}
	return *o.TitleName
}

// GetTitleNameOk returns a tuple with the TitleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) GetTitleNameOk() (*string, bool) {
	if o == nil || IsNil(o.TitleName) {
		return nil, false
	}
	return o.TitleName, true
}

// HasTitleName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) HasTitleName() bool {
	if o != nil && !IsNil(o.TitleName) {
		return true
	}

	return false
}

// SetTitleName gets a reference to the given string and assigns it to the TitleName field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) SetTitleName(v string) {
	o.TitleName = &v
}

func (o AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.BankAccount) {
		toSerialize["bank_account"] = o.BankAccount
	}
	if !IsNil(o.BankName) {
		toSerialize["bank_name"] = o.BankName
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if !IsNil(o.TaxRegisterNo) {
		toSerialize["tax_register_no"] = o.TaxRegisterNo
	}
	if !IsNil(o.Telephone) {
		toSerialize["telephone"] = o.Telephone
	}
	if !IsNil(o.TitleName) {
		toSerialize["title_name"] = o.TitleName
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel struct {
	value *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) Get() *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) Set(val *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel(val *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel {
	return &NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

