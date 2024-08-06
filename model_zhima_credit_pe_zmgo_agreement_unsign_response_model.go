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

// checks if the ZhimaCreditPeZmgoAgreementUnsignResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZhimaCreditPeZmgoAgreementUnsignResponseModel{}

// ZhimaCreditPeZmgoAgreementUnsignResponseModel struct for ZhimaCreditPeZmgoAgreementUnsignResponseModel
type ZhimaCreditPeZmgoAgreementUnsignResponseModel struct {
	// 支付宝系统中用以唯一标识用户签约记录的编号。
	AgreementId *string `json:"agreement_id,omitempty"`
	// 调用芝麻GO结算受理接口时，需要传入该值
	WithholdPlanNo *string `json:"withhold_plan_no,omitempty"`
}

// NewZhimaCreditPeZmgoAgreementUnsignResponseModel instantiates a new ZhimaCreditPeZmgoAgreementUnsignResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZhimaCreditPeZmgoAgreementUnsignResponseModel() *ZhimaCreditPeZmgoAgreementUnsignResponseModel {
	this := ZhimaCreditPeZmgoAgreementUnsignResponseModel{}
	return &this
}

// NewZhimaCreditPeZmgoAgreementUnsignResponseModelWithDefaults instantiates a new ZhimaCreditPeZmgoAgreementUnsignResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZhimaCreditPeZmgoAgreementUnsignResponseModelWithDefaults() *ZhimaCreditPeZmgoAgreementUnsignResponseModel {
	this := ZhimaCreditPeZmgoAgreementUnsignResponseModel{}
	return &this
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoAgreementUnsignResponseModel) GetAgreementId() string {
	if o == nil || IsNil(o.AgreementId) {
		var ret string
		return ret
	}
	return *o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoAgreementUnsignResponseModel) GetAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementId) {
		return nil, false
	}
	return o.AgreementId, true
}

// HasAgreementId returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoAgreementUnsignResponseModel) HasAgreementId() bool {
	if o != nil && !IsNil(o.AgreementId) {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given string and assigns it to the AgreementId field.
func (o *ZhimaCreditPeZmgoAgreementUnsignResponseModel) SetAgreementId(v string) {
	o.AgreementId = &v
}

// GetWithholdPlanNo returns the WithholdPlanNo field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoAgreementUnsignResponseModel) GetWithholdPlanNo() string {
	if o == nil || IsNil(o.WithholdPlanNo) {
		var ret string
		return ret
	}
	return *o.WithholdPlanNo
}

// GetWithholdPlanNoOk returns a tuple with the WithholdPlanNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoAgreementUnsignResponseModel) GetWithholdPlanNoOk() (*string, bool) {
	if o == nil || IsNil(o.WithholdPlanNo) {
		return nil, false
	}
	return o.WithholdPlanNo, true
}

// HasWithholdPlanNo returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoAgreementUnsignResponseModel) HasWithholdPlanNo() bool {
	if o != nil && !IsNil(o.WithholdPlanNo) {
		return true
	}

	return false
}

// SetWithholdPlanNo gets a reference to the given string and assigns it to the WithholdPlanNo field.
func (o *ZhimaCreditPeZmgoAgreementUnsignResponseModel) SetWithholdPlanNo(v string) {
	o.WithholdPlanNo = &v
}

func (o ZhimaCreditPeZmgoAgreementUnsignResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZhimaCreditPeZmgoAgreementUnsignResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementId) {
		toSerialize["agreement_id"] = o.AgreementId
	}
	if !IsNil(o.WithholdPlanNo) {
		toSerialize["withhold_plan_no"] = o.WithholdPlanNo
	}
	return toSerialize, nil
}

type NullableZhimaCreditPeZmgoAgreementUnsignResponseModel struct {
	value *ZhimaCreditPeZmgoAgreementUnsignResponseModel
	isSet bool
}

func (v NullableZhimaCreditPeZmgoAgreementUnsignResponseModel) Get() *ZhimaCreditPeZmgoAgreementUnsignResponseModel {
	return v.value
}

func (v *NullableZhimaCreditPeZmgoAgreementUnsignResponseModel) Set(val *ZhimaCreditPeZmgoAgreementUnsignResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaCreditPeZmgoAgreementUnsignResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaCreditPeZmgoAgreementUnsignResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaCreditPeZmgoAgreementUnsignResponseModel(val *ZhimaCreditPeZmgoAgreementUnsignResponseModel) *NullableZhimaCreditPeZmgoAgreementUnsignResponseModel {
	return &NullableZhimaCreditPeZmgoAgreementUnsignResponseModel{value: val, isSet: true}
}

func (v NullableZhimaCreditPeZmgoAgreementUnsignResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaCreditPeZmgoAgreementUnsignResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

