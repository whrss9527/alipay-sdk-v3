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

// checks if the AlipayMarketingCardBenefitModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingCardBenefitModifyModel{}

// AlipayMarketingCardBenefitModifyModel struct for AlipayMarketingCardBenefitModifyModel
type AlipayMarketingCardBenefitModifyModel struct {
	// 权益ID，通过 <a  href=\"https://opendocs.alipay.com/apis/api_5/alipay.marketing.card.benefit.create\">alipay.marketing.card.benefit.create</a>(会员卡模板外部权益创建)接口创建获取。
	BenefitId *string `json:"benefit_id,omitempty"`
	McardTemplateBenefit *McardTemplateBenefit `json:"mcard_template_benefit,omitempty"`
}

// NewAlipayMarketingCardBenefitModifyModel instantiates a new AlipayMarketingCardBenefitModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingCardBenefitModifyModel() *AlipayMarketingCardBenefitModifyModel {
	this := AlipayMarketingCardBenefitModifyModel{}
	return &this
}

// NewAlipayMarketingCardBenefitModifyModelWithDefaults instantiates a new AlipayMarketingCardBenefitModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingCardBenefitModifyModelWithDefaults() *AlipayMarketingCardBenefitModifyModel {
	this := AlipayMarketingCardBenefitModifyModel{}
	return &this
}

// GetBenefitId returns the BenefitId field value if set, zero value otherwise.
func (o *AlipayMarketingCardBenefitModifyModel) GetBenefitId() string {
	if o == nil || IsNil(o.BenefitId) {
		var ret string
		return ret
	}
	return *o.BenefitId
}

// GetBenefitIdOk returns a tuple with the BenefitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardBenefitModifyModel) GetBenefitIdOk() (*string, bool) {
	if o == nil || IsNil(o.BenefitId) {
		return nil, false
	}
	return o.BenefitId, true
}

// HasBenefitId returns a boolean if a field has been set.
func (o *AlipayMarketingCardBenefitModifyModel) HasBenefitId() bool {
	if o != nil && !IsNil(o.BenefitId) {
		return true
	}

	return false
}

// SetBenefitId gets a reference to the given string and assigns it to the BenefitId field.
func (o *AlipayMarketingCardBenefitModifyModel) SetBenefitId(v string) {
	o.BenefitId = &v
}

// GetMcardTemplateBenefit returns the McardTemplateBenefit field value if set, zero value otherwise.
func (o *AlipayMarketingCardBenefitModifyModel) GetMcardTemplateBenefit() McardTemplateBenefit {
	if o == nil || IsNil(o.McardTemplateBenefit) {
		var ret McardTemplateBenefit
		return ret
	}
	return *o.McardTemplateBenefit
}

// GetMcardTemplateBenefitOk returns a tuple with the McardTemplateBenefit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardBenefitModifyModel) GetMcardTemplateBenefitOk() (*McardTemplateBenefit, bool) {
	if o == nil || IsNil(o.McardTemplateBenefit) {
		return nil, false
	}
	return o.McardTemplateBenefit, true
}

// HasMcardTemplateBenefit returns a boolean if a field has been set.
func (o *AlipayMarketingCardBenefitModifyModel) HasMcardTemplateBenefit() bool {
	if o != nil && !IsNil(o.McardTemplateBenefit) {
		return true
	}

	return false
}

// SetMcardTemplateBenefit gets a reference to the given McardTemplateBenefit and assigns it to the McardTemplateBenefit field.
func (o *AlipayMarketingCardBenefitModifyModel) SetMcardTemplateBenefit(v McardTemplateBenefit) {
	o.McardTemplateBenefit = &v
}

func (o AlipayMarketingCardBenefitModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingCardBenefitModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BenefitId) {
		toSerialize["benefit_id"] = o.BenefitId
	}
	if !IsNil(o.McardTemplateBenefit) {
		toSerialize["mcard_template_benefit"] = o.McardTemplateBenefit
	}
	return toSerialize, nil
}

type NullableAlipayMarketingCardBenefitModifyModel struct {
	value *AlipayMarketingCardBenefitModifyModel
	isSet bool
}

func (v NullableAlipayMarketingCardBenefitModifyModel) Get() *AlipayMarketingCardBenefitModifyModel {
	return v.value
}

func (v *NullableAlipayMarketingCardBenefitModifyModel) Set(val *AlipayMarketingCardBenefitModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardBenefitModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardBenefitModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardBenefitModifyModel(val *AlipayMarketingCardBenefitModifyModel) *NullableAlipayMarketingCardBenefitModifyModel {
	return &NullableAlipayMarketingCardBenefitModifyModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardBenefitModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardBenefitModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

