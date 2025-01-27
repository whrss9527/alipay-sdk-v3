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

// checks if the AlipayFundJointaccountMemberBindModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundJointaccountMemberBindModel{}

// AlipayFundJointaccountMemberBindModel struct for AlipayFundJointaccountMemberBindModel
type AlipayFundJointaccountMemberBindModel struct {
	// 账本ID
	AccountId    *string               `json:"account_id,omitempty"`
	AccountQuota *JointAccountQuotaDTO `json:"account_quota,omitempty"`
	// 授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 业务场景，联系支付宝分配
	BizScene       *string                        `json:"biz_scene,omitempty"`
	BusinessParams *InviteMemberBusinessParamsDTO `json:"business_params,omitempty"`
	// 成员账号： identity_type是ALIPAY_USER_ID填支付宝会员ID（2088开头）； 是ALIPAY_LOGON_ID 填支付宝登录号； 是ALIPAY_OPEN_ID 填支付宝openId
	Identity *string `json:"identity,omitempty"`
	// 账号类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID  2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式 3、ALIPAY_OPEN_ID：支付宝openId
	IdentityType         *string                  `json:"identity_type,omitempty"`
	IdentityVerifiedInfo *IdentityVerifiedInfoDTO `json:"identity_verified_info,omitempty"`
	// 成员姓名，账号类型为ALIPAY_LOGON_ID时必填
	Name *string `json:"name,omitempty"`
	// 产品码
	ProductCode *string `json:"product_code,omitempty"`
}

// NewAlipayFundJointaccountMemberBindModel instantiates a new AlipayFundJointaccountMemberBindModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundJointaccountMemberBindModel() *AlipayFundJointaccountMemberBindModel {
	this := AlipayFundJointaccountMemberBindModel{}
	return &this
}

// NewAlipayFundJointaccountMemberBindModelWithDefaults instantiates a new AlipayFundJointaccountMemberBindModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundJointaccountMemberBindModelWithDefaults() *AlipayFundJointaccountMemberBindModel {
	this := AlipayFundJointaccountMemberBindModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberBindModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberBindModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberBindModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayFundJointaccountMemberBindModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAccountQuota returns the AccountQuota field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberBindModel) GetAccountQuota() JointAccountQuotaDTO {
	if o == nil || IsNil(o.AccountQuota) {
		var ret JointAccountQuotaDTO
		return ret
	}
	return *o.AccountQuota
}

// GetAccountQuotaOk returns a tuple with the AccountQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberBindModel) GetAccountQuotaOk() (*JointAccountQuotaDTO, bool) {
	if o == nil || IsNil(o.AccountQuota) {
		return nil, false
	}
	return o.AccountQuota, true
}

// HasAccountQuota returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberBindModel) HasAccountQuota() bool {
	if o != nil && !IsNil(o.AccountQuota) {
		return true
	}

	return false
}

// SetAccountQuota gets a reference to the given JointAccountQuotaDTO and assigns it to the AccountQuota field.
func (o *AlipayFundJointaccountMemberBindModel) SetAccountQuota(v JointAccountQuotaDTO) {
	o.AccountQuota = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberBindModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberBindModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberBindModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayFundJointaccountMemberBindModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetBizScene returns the BizScene field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberBindModel) GetBizScene() string {
	if o == nil || IsNil(o.BizScene) {
		var ret string
		return ret
	}
	return *o.BizScene
}

// GetBizSceneOk returns a tuple with the BizScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberBindModel) GetBizSceneOk() (*string, bool) {
	if o == nil || IsNil(o.BizScene) {
		return nil, false
	}
	return o.BizScene, true
}

// HasBizScene returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberBindModel) HasBizScene() bool {
	if o != nil && !IsNil(o.BizScene) {
		return true
	}

	return false
}

// SetBizScene gets a reference to the given string and assigns it to the BizScene field.
func (o *AlipayFundJointaccountMemberBindModel) SetBizScene(v string) {
	o.BizScene = &v
}

// GetBusinessParams returns the BusinessParams field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberBindModel) GetBusinessParams() InviteMemberBusinessParamsDTO {
	if o == nil || IsNil(o.BusinessParams) {
		var ret InviteMemberBusinessParamsDTO
		return ret
	}
	return *o.BusinessParams
}

// GetBusinessParamsOk returns a tuple with the BusinessParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberBindModel) GetBusinessParamsOk() (*InviteMemberBusinessParamsDTO, bool) {
	if o == nil || IsNil(o.BusinessParams) {
		return nil, false
	}
	return o.BusinessParams, true
}

// HasBusinessParams returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberBindModel) HasBusinessParams() bool {
	if o != nil && !IsNil(o.BusinessParams) {
		return true
	}

	return false
}

// SetBusinessParams gets a reference to the given InviteMemberBusinessParamsDTO and assigns it to the BusinessParams field.
func (o *AlipayFundJointaccountMemberBindModel) SetBusinessParams(v InviteMemberBusinessParamsDTO) {
	o.BusinessParams = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberBindModel) GetIdentity() string {
	if o == nil || IsNil(o.Identity) {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberBindModel) GetIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberBindModel) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *AlipayFundJointaccountMemberBindModel) SetIdentity(v string) {
	o.Identity = &v
}

// GetIdentityType returns the IdentityType field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberBindModel) GetIdentityType() string {
	if o == nil || IsNil(o.IdentityType) {
		var ret string
		return ret
	}
	return *o.IdentityType
}

// GetIdentityTypeOk returns a tuple with the IdentityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberBindModel) GetIdentityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityType) {
		return nil, false
	}
	return o.IdentityType, true
}

// HasIdentityType returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberBindModel) HasIdentityType() bool {
	if o != nil && !IsNil(o.IdentityType) {
		return true
	}

	return false
}

// SetIdentityType gets a reference to the given string and assigns it to the IdentityType field.
func (o *AlipayFundJointaccountMemberBindModel) SetIdentityType(v string) {
	o.IdentityType = &v
}

// GetIdentityVerifiedInfo returns the IdentityVerifiedInfo field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberBindModel) GetIdentityVerifiedInfo() IdentityVerifiedInfoDTO {
	if o == nil || IsNil(o.IdentityVerifiedInfo) {
		var ret IdentityVerifiedInfoDTO
		return ret
	}
	return *o.IdentityVerifiedInfo
}

// GetIdentityVerifiedInfoOk returns a tuple with the IdentityVerifiedInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberBindModel) GetIdentityVerifiedInfoOk() (*IdentityVerifiedInfoDTO, bool) {
	if o == nil || IsNil(o.IdentityVerifiedInfo) {
		return nil, false
	}
	return o.IdentityVerifiedInfo, true
}

// HasIdentityVerifiedInfo returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberBindModel) HasIdentityVerifiedInfo() bool {
	if o != nil && !IsNil(o.IdentityVerifiedInfo) {
		return true
	}

	return false
}

// SetIdentityVerifiedInfo gets a reference to the given IdentityVerifiedInfoDTO and assigns it to the IdentityVerifiedInfo field.
func (o *AlipayFundJointaccountMemberBindModel) SetIdentityVerifiedInfo(v IdentityVerifiedInfoDTO) {
	o.IdentityVerifiedInfo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberBindModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberBindModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberBindModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlipayFundJointaccountMemberBindModel) SetName(v string) {
	o.Name = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberBindModel) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberBindModel) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberBindModel) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *AlipayFundJointaccountMemberBindModel) SetProductCode(v string) {
	o.ProductCode = &v
}

func (o AlipayFundJointaccountMemberBindModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundJointaccountMemberBindModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AccountQuota) {
		toSerialize["account_quota"] = o.AccountQuota
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.BizScene) {
		toSerialize["biz_scene"] = o.BizScene
	}
	if !IsNil(o.BusinessParams) {
		toSerialize["business_params"] = o.BusinessParams
	}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !IsNil(o.IdentityType) {
		toSerialize["identity_type"] = o.IdentityType
	}
	if !IsNil(o.IdentityVerifiedInfo) {
		toSerialize["identity_verified_info"] = o.IdentityVerifiedInfo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProductCode) {
		toSerialize["product_code"] = o.ProductCode
	}
	return toSerialize, nil
}

type NullableAlipayFundJointaccountMemberBindModel struct {
	value *AlipayFundJointaccountMemberBindModel
	isSet bool
}

func (v NullableAlipayFundJointaccountMemberBindModel) Get() *AlipayFundJointaccountMemberBindModel {
	return v.value
}

func (v *NullableAlipayFundJointaccountMemberBindModel) Set(val *AlipayFundJointaccountMemberBindModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundJointaccountMemberBindModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundJointaccountMemberBindModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundJointaccountMemberBindModel(val *AlipayFundJointaccountMemberBindModel) *NullableAlipayFundJointaccountMemberBindModel {
	return &NullableAlipayFundJointaccountMemberBindModel{value: val, isSet: true}
}

func (v NullableAlipayFundJointaccountMemberBindModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundJointaccountMemberBindModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
