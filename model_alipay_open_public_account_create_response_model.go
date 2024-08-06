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

// checks if the AlipayOpenPublicAccountCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicAccountCreateResponseModel{}

// AlipayOpenPublicAccountCreateResponseModel struct for AlipayOpenPublicAccountCreateResponseModel
type AlipayOpenPublicAccountCreateResponseModel struct {
	// 协议号，商户会员在支付宝服务窗账号中的唯一标识
	AgreementId *string `json:"agreement_id,omitempty"`
}

// NewAlipayOpenPublicAccountCreateResponseModel instantiates a new AlipayOpenPublicAccountCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicAccountCreateResponseModel() *AlipayOpenPublicAccountCreateResponseModel {
	this := AlipayOpenPublicAccountCreateResponseModel{}
	return &this
}

// NewAlipayOpenPublicAccountCreateResponseModelWithDefaults instantiates a new AlipayOpenPublicAccountCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicAccountCreateResponseModelWithDefaults() *AlipayOpenPublicAccountCreateResponseModel {
	this := AlipayOpenPublicAccountCreateResponseModel{}
	return &this
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise.
func (o *AlipayOpenPublicAccountCreateResponseModel) GetAgreementId() string {
	if o == nil || IsNil(o.AgreementId) {
		var ret string
		return ret
	}
	return *o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicAccountCreateResponseModel) GetAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementId) {
		return nil, false
	}
	return o.AgreementId, true
}

// HasAgreementId returns a boolean if a field has been set.
func (o *AlipayOpenPublicAccountCreateResponseModel) HasAgreementId() bool {
	if o != nil && !IsNil(o.AgreementId) {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given string and assigns it to the AgreementId field.
func (o *AlipayOpenPublicAccountCreateResponseModel) SetAgreementId(v string) {
	o.AgreementId = &v
}

func (o AlipayOpenPublicAccountCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicAccountCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementId) {
		toSerialize["agreement_id"] = o.AgreementId
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicAccountCreateResponseModel struct {
	value *AlipayOpenPublicAccountCreateResponseModel
	isSet bool
}

func (v NullableAlipayOpenPublicAccountCreateResponseModel) Get() *AlipayOpenPublicAccountCreateResponseModel {
	return v.value
}

func (v *NullableAlipayOpenPublicAccountCreateResponseModel) Set(val *AlipayOpenPublicAccountCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicAccountCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicAccountCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicAccountCreateResponseModel(val *AlipayOpenPublicAccountCreateResponseModel) *NullableAlipayOpenPublicAccountCreateResponseModel {
	return &NullableAlipayOpenPublicAccountCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicAccountCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicAccountCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

