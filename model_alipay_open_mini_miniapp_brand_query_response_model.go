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

// checks if the AlipayOpenMiniMiniappBrandQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniMiniappBrandQueryResponseModel{}

// AlipayOpenMiniMiniappBrandQueryResponseModel struct for AlipayOpenMiniMiniappBrandQueryResponseModel
type AlipayOpenMiniMiniappBrandQueryResponseModel struct {
	MerchantBrandListResult *MerchantBrandListResult `json:"merchant_brand_list_result,omitempty"`
	MiniappBrandAuditResult *MiniappBrandAuditResult `json:"miniapp_brand_audit_result,omitempty"`
}

// NewAlipayOpenMiniMiniappBrandQueryResponseModel instantiates a new AlipayOpenMiniMiniappBrandQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniMiniappBrandQueryResponseModel() *AlipayOpenMiniMiniappBrandQueryResponseModel {
	this := AlipayOpenMiniMiniappBrandQueryResponseModel{}
	return &this
}

// NewAlipayOpenMiniMiniappBrandQueryResponseModelWithDefaults instantiates a new AlipayOpenMiniMiniappBrandQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniMiniappBrandQueryResponseModelWithDefaults() *AlipayOpenMiniMiniappBrandQueryResponseModel {
	this := AlipayOpenMiniMiniappBrandQueryResponseModel{}
	return &this
}

// GetMerchantBrandListResult returns the MerchantBrandListResult field value if set, zero value otherwise.
func (o *AlipayOpenMiniMiniappBrandQueryResponseModel) GetMerchantBrandListResult() MerchantBrandListResult {
	if o == nil || IsNil(o.MerchantBrandListResult) {
		var ret MerchantBrandListResult
		return ret
	}
	return *o.MerchantBrandListResult
}

// GetMerchantBrandListResultOk returns a tuple with the MerchantBrandListResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniMiniappBrandQueryResponseModel) GetMerchantBrandListResultOk() (*MerchantBrandListResult, bool) {
	if o == nil || IsNil(o.MerchantBrandListResult) {
		return nil, false
	}
	return o.MerchantBrandListResult, true
}

// HasMerchantBrandListResult returns a boolean if a field has been set.
func (o *AlipayOpenMiniMiniappBrandQueryResponseModel) HasMerchantBrandListResult() bool {
	if o != nil && !IsNil(o.MerchantBrandListResult) {
		return true
	}

	return false
}

// SetMerchantBrandListResult gets a reference to the given MerchantBrandListResult and assigns it to the MerchantBrandListResult field.
func (o *AlipayOpenMiniMiniappBrandQueryResponseModel) SetMerchantBrandListResult(v MerchantBrandListResult) {
	o.MerchantBrandListResult = &v
}

// GetMiniappBrandAuditResult returns the MiniappBrandAuditResult field value if set, zero value otherwise.
func (o *AlipayOpenMiniMiniappBrandQueryResponseModel) GetMiniappBrandAuditResult() MiniappBrandAuditResult {
	if o == nil || IsNil(o.MiniappBrandAuditResult) {
		var ret MiniappBrandAuditResult
		return ret
	}
	return *o.MiniappBrandAuditResult
}

// GetMiniappBrandAuditResultOk returns a tuple with the MiniappBrandAuditResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniMiniappBrandQueryResponseModel) GetMiniappBrandAuditResultOk() (*MiniappBrandAuditResult, bool) {
	if o == nil || IsNil(o.MiniappBrandAuditResult) {
		return nil, false
	}
	return o.MiniappBrandAuditResult, true
}

// HasMiniappBrandAuditResult returns a boolean if a field has been set.
func (o *AlipayOpenMiniMiniappBrandQueryResponseModel) HasMiniappBrandAuditResult() bool {
	if o != nil && !IsNil(o.MiniappBrandAuditResult) {
		return true
	}

	return false
}

// SetMiniappBrandAuditResult gets a reference to the given MiniappBrandAuditResult and assigns it to the MiniappBrandAuditResult field.
func (o *AlipayOpenMiniMiniappBrandQueryResponseModel) SetMiniappBrandAuditResult(v MiniappBrandAuditResult) {
	o.MiniappBrandAuditResult = &v
}

func (o AlipayOpenMiniMiniappBrandQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniMiniappBrandQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantBrandListResult) {
		toSerialize["merchant_brand_list_result"] = o.MerchantBrandListResult
	}
	if !IsNil(o.MiniappBrandAuditResult) {
		toSerialize["miniapp_brand_audit_result"] = o.MiniappBrandAuditResult
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniMiniappBrandQueryResponseModel struct {
	value *AlipayOpenMiniMiniappBrandQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniMiniappBrandQueryResponseModel) Get() *AlipayOpenMiniMiniappBrandQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniMiniappBrandQueryResponseModel) Set(val *AlipayOpenMiniMiniappBrandQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniMiniappBrandQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniMiniappBrandQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniMiniappBrandQueryResponseModel(val *AlipayOpenMiniMiniappBrandQueryResponseModel) *NullableAlipayOpenMiniMiniappBrandQueryResponseModel {
	return &NullableAlipayOpenMiniMiniappBrandQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniMiniappBrandQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniMiniappBrandQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

