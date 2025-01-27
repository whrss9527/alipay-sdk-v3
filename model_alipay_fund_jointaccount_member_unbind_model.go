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

// checks if the AlipayFundJointaccountMemberUnbindModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundJointaccountMemberUnbindModel{}

// AlipayFundJointaccountMemberUnbindModel struct for AlipayFundJointaccountMemberUnbindModel
type AlipayFundJointaccountMemberUnbindModel struct {
	// 账本id
	AccountId *string `json:"account_id,omitempty"`
	// 授权协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 场景码
	BizScene *string `json:"biz_scene,omitempty"`
	// 成员账号： identity_type是ALIPAY_USER_ID填支付宝会员ID（2088开头）； 是ALIPAY_LOGON_ID 填支付宝登录号
	Identity *string `json:"identity,omitempty"`
	// 账号类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID 2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式
	IdentityType *string `json:"identity_type,omitempty"`
	// 姓名，账号类型为ALIPAY_LOGON_ID时必填
	Name *string `json:"name,omitempty"`
	// 产品码
	ProductCode *string `json:"product_code,omitempty"`
	// 成员当前状态： 邀请中（PROCESSING）、正常（NORMAL）
	Status *string `json:"status,omitempty"`
}

// NewAlipayFundJointaccountMemberUnbindModel instantiates a new AlipayFundJointaccountMemberUnbindModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundJointaccountMemberUnbindModel() *AlipayFundJointaccountMemberUnbindModel {
	this := AlipayFundJointaccountMemberUnbindModel{}
	return &this
}

// NewAlipayFundJointaccountMemberUnbindModelWithDefaults instantiates a new AlipayFundJointaccountMemberUnbindModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundJointaccountMemberUnbindModelWithDefaults() *AlipayFundJointaccountMemberUnbindModel {
	this := AlipayFundJointaccountMemberUnbindModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberUnbindModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberUnbindModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberUnbindModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayFundJointaccountMemberUnbindModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberUnbindModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberUnbindModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberUnbindModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayFundJointaccountMemberUnbindModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetBizScene returns the BizScene field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberUnbindModel) GetBizScene() string {
	if o == nil || IsNil(o.BizScene) {
		var ret string
		return ret
	}
	return *o.BizScene
}

// GetBizSceneOk returns a tuple with the BizScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberUnbindModel) GetBizSceneOk() (*string, bool) {
	if o == nil || IsNil(o.BizScene) {
		return nil, false
	}
	return o.BizScene, true
}

// HasBizScene returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberUnbindModel) HasBizScene() bool {
	if o != nil && !IsNil(o.BizScene) {
		return true
	}

	return false
}

// SetBizScene gets a reference to the given string and assigns it to the BizScene field.
func (o *AlipayFundJointaccountMemberUnbindModel) SetBizScene(v string) {
	o.BizScene = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberUnbindModel) GetIdentity() string {
	if o == nil || IsNil(o.Identity) {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberUnbindModel) GetIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberUnbindModel) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *AlipayFundJointaccountMemberUnbindModel) SetIdentity(v string) {
	o.Identity = &v
}

// GetIdentityType returns the IdentityType field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberUnbindModel) GetIdentityType() string {
	if o == nil || IsNil(o.IdentityType) {
		var ret string
		return ret
	}
	return *o.IdentityType
}

// GetIdentityTypeOk returns a tuple with the IdentityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberUnbindModel) GetIdentityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityType) {
		return nil, false
	}
	return o.IdentityType, true
}

// HasIdentityType returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberUnbindModel) HasIdentityType() bool {
	if o != nil && !IsNil(o.IdentityType) {
		return true
	}

	return false
}

// SetIdentityType gets a reference to the given string and assigns it to the IdentityType field.
func (o *AlipayFundJointaccountMemberUnbindModel) SetIdentityType(v string) {
	o.IdentityType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberUnbindModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberUnbindModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberUnbindModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlipayFundJointaccountMemberUnbindModel) SetName(v string) {
	o.Name = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberUnbindModel) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberUnbindModel) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberUnbindModel) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *AlipayFundJointaccountMemberUnbindModel) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayFundJointaccountMemberUnbindModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundJointaccountMemberUnbindModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayFundJointaccountMemberUnbindModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayFundJointaccountMemberUnbindModel) SetStatus(v string) {
	o.Status = &v
}

func (o AlipayFundJointaccountMemberUnbindModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundJointaccountMemberUnbindModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.BizScene) {
		toSerialize["biz_scene"] = o.BizScene
	}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !IsNil(o.IdentityType) {
		toSerialize["identity_type"] = o.IdentityType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProductCode) {
		toSerialize["product_code"] = o.ProductCode
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAlipayFundJointaccountMemberUnbindModel struct {
	value *AlipayFundJointaccountMemberUnbindModel
	isSet bool
}

func (v NullableAlipayFundJointaccountMemberUnbindModel) Get() *AlipayFundJointaccountMemberUnbindModel {
	return v.value
}

func (v *NullableAlipayFundJointaccountMemberUnbindModel) Set(val *AlipayFundJointaccountMemberUnbindModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundJointaccountMemberUnbindModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundJointaccountMemberUnbindModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundJointaccountMemberUnbindModel(val *AlipayFundJointaccountMemberUnbindModel) *NullableAlipayFundJointaccountMemberUnbindModel {
	return &NullableAlipayFundJointaccountMemberUnbindModel{value: val, isSet: true}
}

func (v NullableAlipayFundJointaccountMemberUnbindModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundJointaccountMemberUnbindModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
