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

// checks if the AlipayCommerceEcUserEnterpriseQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceEcUserEnterpriseQueryResponseModel{}

// AlipayCommerceEcUserEnterpriseQueryResponseModel struct for AlipayCommerceEcUserEnterpriseQueryResponseModel
type AlipayCommerceEcUserEnterpriseQueryResponseModel struct {
	// 用户所属企业列表
	EnterpriseInfoList []EnterpriseInfoDTO `json:"enterprise_info_list,omitempty"`
}

// NewAlipayCommerceEcUserEnterpriseQueryResponseModel instantiates a new AlipayCommerceEcUserEnterpriseQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceEcUserEnterpriseQueryResponseModel() *AlipayCommerceEcUserEnterpriseQueryResponseModel {
	this := AlipayCommerceEcUserEnterpriseQueryResponseModel{}
	return &this
}

// NewAlipayCommerceEcUserEnterpriseQueryResponseModelWithDefaults instantiates a new AlipayCommerceEcUserEnterpriseQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceEcUserEnterpriseQueryResponseModelWithDefaults() *AlipayCommerceEcUserEnterpriseQueryResponseModel {
	this := AlipayCommerceEcUserEnterpriseQueryResponseModel{}
	return &this
}

// GetEnterpriseInfoList returns the EnterpriseInfoList field value if set, zero value otherwise.
func (o *AlipayCommerceEcUserEnterpriseQueryResponseModel) GetEnterpriseInfoList() []EnterpriseInfoDTO {
	if o == nil || IsNil(o.EnterpriseInfoList) {
		var ret []EnterpriseInfoDTO
		return ret
	}
	return o.EnterpriseInfoList
}

// GetEnterpriseInfoListOk returns a tuple with the EnterpriseInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcUserEnterpriseQueryResponseModel) GetEnterpriseInfoListOk() ([]EnterpriseInfoDTO, bool) {
	if o == nil || IsNil(o.EnterpriseInfoList) {
		return nil, false
	}
	return o.EnterpriseInfoList, true
}

// HasEnterpriseInfoList returns a boolean if a field has been set.
func (o *AlipayCommerceEcUserEnterpriseQueryResponseModel) HasEnterpriseInfoList() bool {
	if o != nil && !IsNil(o.EnterpriseInfoList) {
		return true
	}

	return false
}

// SetEnterpriseInfoList gets a reference to the given []EnterpriseInfoDTO and assigns it to the EnterpriseInfoList field.
func (o *AlipayCommerceEcUserEnterpriseQueryResponseModel) SetEnterpriseInfoList(v []EnterpriseInfoDTO) {
	o.EnterpriseInfoList = v
}

func (o AlipayCommerceEcUserEnterpriseQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceEcUserEnterpriseQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnterpriseInfoList) {
		toSerialize["enterprise_info_list"] = o.EnterpriseInfoList
	}
	return toSerialize, nil
}

type NullableAlipayCommerceEcUserEnterpriseQueryResponseModel struct {
	value *AlipayCommerceEcUserEnterpriseQueryResponseModel
	isSet bool
}

func (v NullableAlipayCommerceEcUserEnterpriseQueryResponseModel) Get() *AlipayCommerceEcUserEnterpriseQueryResponseModel {
	return v.value
}

func (v *NullableAlipayCommerceEcUserEnterpriseQueryResponseModel) Set(val *AlipayCommerceEcUserEnterpriseQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcUserEnterpriseQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcUserEnterpriseQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcUserEnterpriseQueryResponseModel(val *AlipayCommerceEcUserEnterpriseQueryResponseModel) *NullableAlipayCommerceEcUserEnterpriseQueryResponseModel {
	return &NullableAlipayCommerceEcUserEnterpriseQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcUserEnterpriseQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcUserEnterpriseQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
