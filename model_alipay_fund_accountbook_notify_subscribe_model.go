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

// checks if the AlipayFundAccountbookNotifySubscribeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundAccountbookNotifySubscribeModel{}

// AlipayFundAccountbookNotifySubscribeModel struct for AlipayFundAccountbookNotifySubscribeModel
type AlipayFundAccountbookNotifySubscribeModel struct {
	// 记账本ID
	AccountBookId *string `json:"account_book_id,omitempty"`
	// 协议号。 若是基于协议的记账本，则agreement_no必传； 若是自创建的记账本，则agreement_no不传；
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 场景码。固定为DEFAULT
	BizScene *string `json:"biz_scene,omitempty"`
	// 产品码。固定为SATF_FUND_BOOK
	ProductCode *string `json:"product_code,omitempty"`
}

// NewAlipayFundAccountbookNotifySubscribeModel instantiates a new AlipayFundAccountbookNotifySubscribeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundAccountbookNotifySubscribeModel() *AlipayFundAccountbookNotifySubscribeModel {
	this := AlipayFundAccountbookNotifySubscribeModel{}
	return &this
}

// NewAlipayFundAccountbookNotifySubscribeModelWithDefaults instantiates a new AlipayFundAccountbookNotifySubscribeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundAccountbookNotifySubscribeModelWithDefaults() *AlipayFundAccountbookNotifySubscribeModel {
	this := AlipayFundAccountbookNotifySubscribeModel{}
	return &this
}

// GetAccountBookId returns the AccountBookId field value if set, zero value otherwise.
func (o *AlipayFundAccountbookNotifySubscribeModel) GetAccountBookId() string {
	if o == nil || IsNil(o.AccountBookId) {
		var ret string
		return ret
	}
	return *o.AccountBookId
}

// GetAccountBookIdOk returns a tuple with the AccountBookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAccountbookNotifySubscribeModel) GetAccountBookIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountBookId) {
		return nil, false
	}
	return o.AccountBookId, true
}

// HasAccountBookId returns a boolean if a field has been set.
func (o *AlipayFundAccountbookNotifySubscribeModel) HasAccountBookId() bool {
	if o != nil && !IsNil(o.AccountBookId) {
		return true
	}

	return false
}

// SetAccountBookId gets a reference to the given string and assigns it to the AccountBookId field.
func (o *AlipayFundAccountbookNotifySubscribeModel) SetAccountBookId(v string) {
	o.AccountBookId = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayFundAccountbookNotifySubscribeModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAccountbookNotifySubscribeModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayFundAccountbookNotifySubscribeModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayFundAccountbookNotifySubscribeModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetBizScene returns the BizScene field value if set, zero value otherwise.
func (o *AlipayFundAccountbookNotifySubscribeModel) GetBizScene() string {
	if o == nil || IsNil(o.BizScene) {
		var ret string
		return ret
	}
	return *o.BizScene
}

// GetBizSceneOk returns a tuple with the BizScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAccountbookNotifySubscribeModel) GetBizSceneOk() (*string, bool) {
	if o == nil || IsNil(o.BizScene) {
		return nil, false
	}
	return o.BizScene, true
}

// HasBizScene returns a boolean if a field has been set.
func (o *AlipayFundAccountbookNotifySubscribeModel) HasBizScene() bool {
	if o != nil && !IsNil(o.BizScene) {
		return true
	}

	return false
}

// SetBizScene gets a reference to the given string and assigns it to the BizScene field.
func (o *AlipayFundAccountbookNotifySubscribeModel) SetBizScene(v string) {
	o.BizScene = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *AlipayFundAccountbookNotifySubscribeModel) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAccountbookNotifySubscribeModel) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *AlipayFundAccountbookNotifySubscribeModel) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *AlipayFundAccountbookNotifySubscribeModel) SetProductCode(v string) {
	o.ProductCode = &v
}

func (o AlipayFundAccountbookNotifySubscribeModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundAccountbookNotifySubscribeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountBookId) {
		toSerialize["account_book_id"] = o.AccountBookId
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.BizScene) {
		toSerialize["biz_scene"] = o.BizScene
	}
	if !IsNil(o.ProductCode) {
		toSerialize["product_code"] = o.ProductCode
	}
	return toSerialize, nil
}

type NullableAlipayFundAccountbookNotifySubscribeModel struct {
	value *AlipayFundAccountbookNotifySubscribeModel
	isSet bool
}

func (v NullableAlipayFundAccountbookNotifySubscribeModel) Get() *AlipayFundAccountbookNotifySubscribeModel {
	return v.value
}

func (v *NullableAlipayFundAccountbookNotifySubscribeModel) Set(val *AlipayFundAccountbookNotifySubscribeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundAccountbookNotifySubscribeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundAccountbookNotifySubscribeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundAccountbookNotifySubscribeModel(val *AlipayFundAccountbookNotifySubscribeModel) *NullableAlipayFundAccountbookNotifySubscribeModel {
	return &NullableAlipayFundAccountbookNotifySubscribeModel{value: val, isSet: true}
}

func (v NullableAlipayFundAccountbookNotifySubscribeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundAccountbookNotifySubscribeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
