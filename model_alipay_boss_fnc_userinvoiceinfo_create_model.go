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

// checks if the AlipayBossFncUserinvoiceinfoCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayBossFncUserinvoiceinfoCreateModel{}

// AlipayBossFncUserinvoiceinfoCreateModel struct for AlipayBossFncUserinvoiceinfoCreateModel
type AlipayBossFncUserinvoiceinfoCreateModel struct {
	// 是否接受电子票 true:是，false:否
	AcceptElectronic *bool `json:"accept_electronic,omitempty"`
	// 公司注册地址，公司营业执照上登记的住址，一般纳税人必须填写，小规模纳税人无需填写
	Address *string `json:"address,omitempty"`
	// 开票资料pid寻址方式 true:pid寻址mid,优先以mid维度保存开票资料,若无mid则以pid维度保存开票资料 false:不寻址mid,以pid维度保存开票资料;非主站商户体系都是以pid（ipRoleId）维度保存开票资料，该值直接填为false
	Addressing *bool `json:"addressing,omitempty"`
	// 是否自动申请开票 true:是，false:否
	Auto *bool `json:"auto,omitempty"`
	// 银行账户，公司银行账号  一般纳税人必须填写，小规模纳税人无需填写
	BankAccount *string `json:"bank_account,omitempty"`
	// 开户行，办理银行开户手续的营业网点  一般纳税人必须填写，小规模纳税人无需填写
	BankName *string `json:"bank_name,omitempty"`
	// 营业执照地址，营业执照上传oss上的地址
	BusinessLicenceUrl *string `json:"business_licence_url,omitempty"`
	// 是否暂停开票，商户添加的开票资料就是不暂停开票，如果后续要暂停开票，只能去结算中台处理。
	Hold *bool `json:"hold,omitempty"`
	// 商户的pid（ipRoleId）
	IpRoleId *string `json:"ip_role_id,omitempty"`
	// 银行开户许可证附件，银行开户许可证上传oss上的地址
	OpenAccountPermitUrl *string `json:"open_account_permit_url,omitempty"`
	// 当前操作人
	Operator *string `json:"operator,omitempty"`
	// 当前操作人类型,01:商户 02:管理员(小二)
	OperatorType *string `json:"operator_type,omitempty"`
	// 其它资质证明地址，其他资质证明附件上传oss上的地址
	OtherQualificationUrl *string `json:"other_qualification_url,omitempty"`
	// 外部只允许添加商户开票资料，该值填写为false
	Ou *bool `json:"ou,omitempty"`
	// 纳税人识别号，税务登记证上的号码。一般纳税人必须填写，小规模纳税人如果是商户也必须填写，个人无需填写
	TaxNo *string `json:"tax_no,omitempty"`
	// 纳税人资格开始时间 （格式：时间戳）  一般纳税人必须填写，小规模纳税人无需填写
	TaxPayerQualiValid *string `json:"tax_payer_quali_valid,omitempty"`
	// 纳税人资格种类，01：一般纳税人；02：小规模纳税人；03：国际商户
	TaxPayerQualification *string `json:"tax_payer_qualification,omitempty"`
	// 一般纳税人资格证书地址，一般纳税人资格证书上传oss上的地址
	TaxQualificationUrl *string `json:"tax_qualification_url,omitempty"`
	// 税务登记证地址，税务登记证上传oss上的地址
	TaxRegCertUrl *string `json:"tax_reg_cert_url,omitempty"`
	// 公司注册电话（手机号和座机均可）  一般纳税人必须填写，小规模纳税人无需填写
	Telephone *string `json:"telephone,omitempty"`
	// 发票抬头，票面信息上的抬头信息
	Title *string `json:"title,omitempty"`
	// 收件人列表,若不修改此项可以不填  选择非电子票时，邮寄信息必填，且保证邮寄信息的ipRoleId和开票资料的ipRoleId相等
	UserMailInfoOrderList []UserMailInfoOrder `json:"user_mail_info_order_list,omitempty"`
}

// NewAlipayBossFncUserinvoiceinfoCreateModel instantiates a new AlipayBossFncUserinvoiceinfoCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayBossFncUserinvoiceinfoCreateModel() *AlipayBossFncUserinvoiceinfoCreateModel {
	this := AlipayBossFncUserinvoiceinfoCreateModel{}
	return &this
}

// NewAlipayBossFncUserinvoiceinfoCreateModelWithDefaults instantiates a new AlipayBossFncUserinvoiceinfoCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayBossFncUserinvoiceinfoCreateModelWithDefaults() *AlipayBossFncUserinvoiceinfoCreateModel {
	this := AlipayBossFncUserinvoiceinfoCreateModel{}
	return &this
}

// GetAcceptElectronic returns the AcceptElectronic field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetAcceptElectronic() bool {
	if o == nil || IsNil(o.AcceptElectronic) {
		var ret bool
		return ret
	}
	return *o.AcceptElectronic
}

// GetAcceptElectronicOk returns a tuple with the AcceptElectronic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetAcceptElectronicOk() (*bool, bool) {
	if o == nil || IsNil(o.AcceptElectronic) {
		return nil, false
	}
	return o.AcceptElectronic, true
}

// HasAcceptElectronic returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasAcceptElectronic() bool {
	if o != nil && !IsNil(o.AcceptElectronic) {
		return true
	}

	return false
}

// SetAcceptElectronic gets a reference to the given bool and assigns it to the AcceptElectronic field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetAcceptElectronic(v bool) {
	o.AcceptElectronic = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetAddress(v string) {
	o.Address = &v
}

// GetAddressing returns the Addressing field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetAddressing() bool {
	if o == nil || IsNil(o.Addressing) {
		var ret bool
		return ret
	}
	return *o.Addressing
}

// GetAddressingOk returns a tuple with the Addressing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetAddressingOk() (*bool, bool) {
	if o == nil || IsNil(o.Addressing) {
		return nil, false
	}
	return o.Addressing, true
}

// HasAddressing returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasAddressing() bool {
	if o != nil && !IsNil(o.Addressing) {
		return true
	}

	return false
}

// SetAddressing gets a reference to the given bool and assigns it to the Addressing field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetAddressing(v bool) {
	o.Addressing = &v
}

// GetAuto returns the Auto field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetAuto() bool {
	if o == nil || IsNil(o.Auto) {
		var ret bool
		return ret
	}
	return *o.Auto
}

// GetAutoOk returns a tuple with the Auto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetAutoOk() (*bool, bool) {
	if o == nil || IsNil(o.Auto) {
		return nil, false
	}
	return o.Auto, true
}

// HasAuto returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasAuto() bool {
	if o != nil && !IsNil(o.Auto) {
		return true
	}

	return false
}

// SetAuto gets a reference to the given bool and assigns it to the Auto field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetAuto(v bool) {
	o.Auto = &v
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetBankAccount() string {
	if o == nil || IsNil(o.BankAccount) {
		var ret string
		return ret
	}
	return *o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetBankAccountOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccount) {
		return nil, false
	}
	return o.BankAccount, true
}

// HasBankAccount returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasBankAccount() bool {
	if o != nil && !IsNil(o.BankAccount) {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given string and assigns it to the BankAccount field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetBankAccount(v string) {
	o.BankAccount = &v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetBankName() string {
	if o == nil || IsNil(o.BankName) {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.BankName) {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasBankName() bool {
	if o != nil && !IsNil(o.BankName) {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetBankName(v string) {
	o.BankName = &v
}

// GetBusinessLicenceUrl returns the BusinessLicenceUrl field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetBusinessLicenceUrl() string {
	if o == nil || IsNil(o.BusinessLicenceUrl) {
		var ret string
		return ret
	}
	return *o.BusinessLicenceUrl
}

// GetBusinessLicenceUrlOk returns a tuple with the BusinessLicenceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetBusinessLicenceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessLicenceUrl) {
		return nil, false
	}
	return o.BusinessLicenceUrl, true
}

// HasBusinessLicenceUrl returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasBusinessLicenceUrl() bool {
	if o != nil && !IsNil(o.BusinessLicenceUrl) {
		return true
	}

	return false
}

// SetBusinessLicenceUrl gets a reference to the given string and assigns it to the BusinessLicenceUrl field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetBusinessLicenceUrl(v string) {
	o.BusinessLicenceUrl = &v
}

// GetHold returns the Hold field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetHold() bool {
	if o == nil || IsNil(o.Hold) {
		var ret bool
		return ret
	}
	return *o.Hold
}

// GetHoldOk returns a tuple with the Hold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetHoldOk() (*bool, bool) {
	if o == nil || IsNil(o.Hold) {
		return nil, false
	}
	return o.Hold, true
}

// HasHold returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasHold() bool {
	if o != nil && !IsNil(o.Hold) {
		return true
	}

	return false
}

// SetHold gets a reference to the given bool and assigns it to the Hold field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetHold(v bool) {
	o.Hold = &v
}

// GetIpRoleId returns the IpRoleId field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetIpRoleId() string {
	if o == nil || IsNil(o.IpRoleId) {
		var ret string
		return ret
	}
	return *o.IpRoleId
}

// GetIpRoleIdOk returns a tuple with the IpRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetIpRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.IpRoleId) {
		return nil, false
	}
	return o.IpRoleId, true
}

// HasIpRoleId returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasIpRoleId() bool {
	if o != nil && !IsNil(o.IpRoleId) {
		return true
	}

	return false
}

// SetIpRoleId gets a reference to the given string and assigns it to the IpRoleId field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetIpRoleId(v string) {
	o.IpRoleId = &v
}

// GetOpenAccountPermitUrl returns the OpenAccountPermitUrl field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetOpenAccountPermitUrl() string {
	if o == nil || IsNil(o.OpenAccountPermitUrl) {
		var ret string
		return ret
	}
	return *o.OpenAccountPermitUrl
}

// GetOpenAccountPermitUrlOk returns a tuple with the OpenAccountPermitUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetOpenAccountPermitUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OpenAccountPermitUrl) {
		return nil, false
	}
	return o.OpenAccountPermitUrl, true
}

// HasOpenAccountPermitUrl returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasOpenAccountPermitUrl() bool {
	if o != nil && !IsNil(o.OpenAccountPermitUrl) {
		return true
	}

	return false
}

// SetOpenAccountPermitUrl gets a reference to the given string and assigns it to the OpenAccountPermitUrl field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetOpenAccountPermitUrl(v string) {
	o.OpenAccountPermitUrl = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetOperator(v string) {
	o.Operator = &v
}

// GetOperatorType returns the OperatorType field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetOperatorType() string {
	if o == nil || IsNil(o.OperatorType) {
		var ret string
		return ret
	}
	return *o.OperatorType
}

// GetOperatorTypeOk returns a tuple with the OperatorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetOperatorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OperatorType) {
		return nil, false
	}
	return o.OperatorType, true
}

// HasOperatorType returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasOperatorType() bool {
	if o != nil && !IsNil(o.OperatorType) {
		return true
	}

	return false
}

// SetOperatorType gets a reference to the given string and assigns it to the OperatorType field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetOperatorType(v string) {
	o.OperatorType = &v
}

// GetOtherQualificationUrl returns the OtherQualificationUrl field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetOtherQualificationUrl() string {
	if o == nil || IsNil(o.OtherQualificationUrl) {
		var ret string
		return ret
	}
	return *o.OtherQualificationUrl
}

// GetOtherQualificationUrlOk returns a tuple with the OtherQualificationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetOtherQualificationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OtherQualificationUrl) {
		return nil, false
	}
	return o.OtherQualificationUrl, true
}

// HasOtherQualificationUrl returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasOtherQualificationUrl() bool {
	if o != nil && !IsNil(o.OtherQualificationUrl) {
		return true
	}

	return false
}

// SetOtherQualificationUrl gets a reference to the given string and assigns it to the OtherQualificationUrl field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetOtherQualificationUrl(v string) {
	o.OtherQualificationUrl = &v
}

// GetOu returns the Ou field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetOu() bool {
	if o == nil || IsNil(o.Ou) {
		var ret bool
		return ret
	}
	return *o.Ou
}

// GetOuOk returns a tuple with the Ou field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetOuOk() (*bool, bool) {
	if o == nil || IsNil(o.Ou) {
		return nil, false
	}
	return o.Ou, true
}

// HasOu returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasOu() bool {
	if o != nil && !IsNil(o.Ou) {
		return true
	}

	return false
}

// SetOu gets a reference to the given bool and assigns it to the Ou field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetOu(v bool) {
	o.Ou = &v
}

// GetTaxNo returns the TaxNo field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetTaxNo() string {
	if o == nil || IsNil(o.TaxNo) {
		var ret string
		return ret
	}
	return *o.TaxNo
}

// GetTaxNoOk returns a tuple with the TaxNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetTaxNoOk() (*string, bool) {
	if o == nil || IsNil(o.TaxNo) {
		return nil, false
	}
	return o.TaxNo, true
}

// HasTaxNo returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasTaxNo() bool {
	if o != nil && !IsNil(o.TaxNo) {
		return true
	}

	return false
}

// SetTaxNo gets a reference to the given string and assigns it to the TaxNo field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetTaxNo(v string) {
	o.TaxNo = &v
}

// GetTaxPayerQualiValid returns the TaxPayerQualiValid field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetTaxPayerQualiValid() string {
	if o == nil || IsNil(o.TaxPayerQualiValid) {
		var ret string
		return ret
	}
	return *o.TaxPayerQualiValid
}

// GetTaxPayerQualiValidOk returns a tuple with the TaxPayerQualiValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetTaxPayerQualiValidOk() (*string, bool) {
	if o == nil || IsNil(o.TaxPayerQualiValid) {
		return nil, false
	}
	return o.TaxPayerQualiValid, true
}

// HasTaxPayerQualiValid returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasTaxPayerQualiValid() bool {
	if o != nil && !IsNil(o.TaxPayerQualiValid) {
		return true
	}

	return false
}

// SetTaxPayerQualiValid gets a reference to the given string and assigns it to the TaxPayerQualiValid field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetTaxPayerQualiValid(v string) {
	o.TaxPayerQualiValid = &v
}

// GetTaxPayerQualification returns the TaxPayerQualification field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetTaxPayerQualification() string {
	if o == nil || IsNil(o.TaxPayerQualification) {
		var ret string
		return ret
	}
	return *o.TaxPayerQualification
}

// GetTaxPayerQualificationOk returns a tuple with the TaxPayerQualification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetTaxPayerQualificationOk() (*string, bool) {
	if o == nil || IsNil(o.TaxPayerQualification) {
		return nil, false
	}
	return o.TaxPayerQualification, true
}

// HasTaxPayerQualification returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasTaxPayerQualification() bool {
	if o != nil && !IsNil(o.TaxPayerQualification) {
		return true
	}

	return false
}

// SetTaxPayerQualification gets a reference to the given string and assigns it to the TaxPayerQualification field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetTaxPayerQualification(v string) {
	o.TaxPayerQualification = &v
}

// GetTaxQualificationUrl returns the TaxQualificationUrl field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetTaxQualificationUrl() string {
	if o == nil || IsNil(o.TaxQualificationUrl) {
		var ret string
		return ret
	}
	return *o.TaxQualificationUrl
}

// GetTaxQualificationUrlOk returns a tuple with the TaxQualificationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetTaxQualificationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TaxQualificationUrl) {
		return nil, false
	}
	return o.TaxQualificationUrl, true
}

// HasTaxQualificationUrl returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasTaxQualificationUrl() bool {
	if o != nil && !IsNil(o.TaxQualificationUrl) {
		return true
	}

	return false
}

// SetTaxQualificationUrl gets a reference to the given string and assigns it to the TaxQualificationUrl field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetTaxQualificationUrl(v string) {
	o.TaxQualificationUrl = &v
}

// GetTaxRegCertUrl returns the TaxRegCertUrl field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetTaxRegCertUrl() string {
	if o == nil || IsNil(o.TaxRegCertUrl) {
		var ret string
		return ret
	}
	return *o.TaxRegCertUrl
}

// GetTaxRegCertUrlOk returns a tuple with the TaxRegCertUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetTaxRegCertUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TaxRegCertUrl) {
		return nil, false
	}
	return o.TaxRegCertUrl, true
}

// HasTaxRegCertUrl returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasTaxRegCertUrl() bool {
	if o != nil && !IsNil(o.TaxRegCertUrl) {
		return true
	}

	return false
}

// SetTaxRegCertUrl gets a reference to the given string and assigns it to the TaxRegCertUrl field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetTaxRegCertUrl(v string) {
	o.TaxRegCertUrl = &v
}

// GetTelephone returns the Telephone field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetTelephone() string {
	if o == nil || IsNil(o.Telephone) {
		var ret string
		return ret
	}
	return *o.Telephone
}

// GetTelephoneOk returns a tuple with the Telephone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetTelephoneOk() (*string, bool) {
	if o == nil || IsNil(o.Telephone) {
		return nil, false
	}
	return o.Telephone, true
}

// HasTelephone returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasTelephone() bool {
	if o != nil && !IsNil(o.Telephone) {
		return true
	}

	return false
}

// SetTelephone gets a reference to the given string and assigns it to the Telephone field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetTelephone(v string) {
	o.Telephone = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetTitle(v string) {
	o.Title = &v
}

// GetUserMailInfoOrderList returns the UserMailInfoOrderList field value if set, zero value otherwise.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetUserMailInfoOrderList() []UserMailInfoOrder {
	if o == nil || IsNil(o.UserMailInfoOrderList) {
		var ret []UserMailInfoOrder
		return ret
	}
	return o.UserMailInfoOrderList
}

// GetUserMailInfoOrderListOk returns a tuple with the UserMailInfoOrderList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) GetUserMailInfoOrderListOk() ([]UserMailInfoOrder, bool) {
	if o == nil || IsNil(o.UserMailInfoOrderList) {
		return nil, false
	}
	return o.UserMailInfoOrderList, true
}

// HasUserMailInfoOrderList returns a boolean if a field has been set.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) HasUserMailInfoOrderList() bool {
	if o != nil && !IsNil(o.UserMailInfoOrderList) {
		return true
	}

	return false
}

// SetUserMailInfoOrderList gets a reference to the given []UserMailInfoOrder and assigns it to the UserMailInfoOrderList field.
func (o *AlipayBossFncUserinvoiceinfoCreateModel) SetUserMailInfoOrderList(v []UserMailInfoOrder) {
	o.UserMailInfoOrderList = v
}

func (o AlipayBossFncUserinvoiceinfoCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayBossFncUserinvoiceinfoCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceptElectronic) {
		toSerialize["accept_electronic"] = o.AcceptElectronic
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Addressing) {
		toSerialize["addressing"] = o.Addressing
	}
	if !IsNil(o.Auto) {
		toSerialize["auto"] = o.Auto
	}
	if !IsNil(o.BankAccount) {
		toSerialize["bank_account"] = o.BankAccount
	}
	if !IsNil(o.BankName) {
		toSerialize["bank_name"] = o.BankName
	}
	if !IsNil(o.BusinessLicenceUrl) {
		toSerialize["business_licence_url"] = o.BusinessLicenceUrl
	}
	if !IsNil(o.Hold) {
		toSerialize["hold"] = o.Hold
	}
	if !IsNil(o.IpRoleId) {
		toSerialize["ip_role_id"] = o.IpRoleId
	}
	if !IsNil(o.OpenAccountPermitUrl) {
		toSerialize["open_account_permit_url"] = o.OpenAccountPermitUrl
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.OperatorType) {
		toSerialize["operator_type"] = o.OperatorType
	}
	if !IsNil(o.OtherQualificationUrl) {
		toSerialize["other_qualification_url"] = o.OtherQualificationUrl
	}
	if !IsNil(o.Ou) {
		toSerialize["ou"] = o.Ou
	}
	if !IsNil(o.TaxNo) {
		toSerialize["tax_no"] = o.TaxNo
	}
	if !IsNil(o.TaxPayerQualiValid) {
		toSerialize["tax_payer_quali_valid"] = o.TaxPayerQualiValid
	}
	if !IsNil(o.TaxPayerQualification) {
		toSerialize["tax_payer_qualification"] = o.TaxPayerQualification
	}
	if !IsNil(o.TaxQualificationUrl) {
		toSerialize["tax_qualification_url"] = o.TaxQualificationUrl
	}
	if !IsNil(o.TaxRegCertUrl) {
		toSerialize["tax_reg_cert_url"] = o.TaxRegCertUrl
	}
	if !IsNil(o.Telephone) {
		toSerialize["telephone"] = o.Telephone
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.UserMailInfoOrderList) {
		toSerialize["user_mail_info_order_list"] = o.UserMailInfoOrderList
	}
	return toSerialize, nil
}

type NullableAlipayBossFncUserinvoiceinfoCreateModel struct {
	value *AlipayBossFncUserinvoiceinfoCreateModel
	isSet bool
}

func (v NullableAlipayBossFncUserinvoiceinfoCreateModel) Get() *AlipayBossFncUserinvoiceinfoCreateModel {
	return v.value
}

func (v *NullableAlipayBossFncUserinvoiceinfoCreateModel) Set(val *AlipayBossFncUserinvoiceinfoCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayBossFncUserinvoiceinfoCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayBossFncUserinvoiceinfoCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayBossFncUserinvoiceinfoCreateModel(val *AlipayBossFncUserinvoiceinfoCreateModel) *NullableAlipayBossFncUserinvoiceinfoCreateModel {
	return &NullableAlipayBossFncUserinvoiceinfoCreateModel{value: val, isSet: true}
}

func (v NullableAlipayBossFncUserinvoiceinfoCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayBossFncUserinvoiceinfoCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

