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

// checks if the AlipayCommerceEcEnterpriseAddressModifyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceEcEnterpriseAddressModifyResponseModel{}

// AlipayCommerceEcEnterpriseAddressModifyResponseModel struct for AlipayCommerceEcEnterpriseAddressModifyResponseModel
type AlipayCommerceEcEnterpriseAddressModifyResponseModel struct {
	// 地址id
	AddressId *string `json:"address_id,omitempty"`
}

// NewAlipayCommerceEcEnterpriseAddressModifyResponseModel instantiates a new AlipayCommerceEcEnterpriseAddressModifyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceEcEnterpriseAddressModifyResponseModel() *AlipayCommerceEcEnterpriseAddressModifyResponseModel {
	this := AlipayCommerceEcEnterpriseAddressModifyResponseModel{}
	return &this
}

// NewAlipayCommerceEcEnterpriseAddressModifyResponseModelWithDefaults instantiates a new AlipayCommerceEcEnterpriseAddressModifyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceEcEnterpriseAddressModifyResponseModelWithDefaults() *AlipayCommerceEcEnterpriseAddressModifyResponseModel {
	this := AlipayCommerceEcEnterpriseAddressModifyResponseModel{}
	return &this
}

// GetAddressId returns the AddressId field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseAddressModifyResponseModel) GetAddressId() string {
	if o == nil || IsNil(o.AddressId) {
		var ret string
		return ret
	}
	return *o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseAddressModifyResponseModel) GetAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.AddressId) {
		return nil, false
	}
	return o.AddressId, true
}

// HasAddressId returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseAddressModifyResponseModel) HasAddressId() bool {
	if o != nil && !IsNil(o.AddressId) {
		return true
	}

	return false
}

// SetAddressId gets a reference to the given string and assigns it to the AddressId field.
func (o *AlipayCommerceEcEnterpriseAddressModifyResponseModel) SetAddressId(v string) {
	o.AddressId = &v
}

func (o AlipayCommerceEcEnterpriseAddressModifyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceEcEnterpriseAddressModifyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressId) {
		toSerialize["address_id"] = o.AddressId
	}
	return toSerialize, nil
}

type NullableAlipayCommerceEcEnterpriseAddressModifyResponseModel struct {
	value *AlipayCommerceEcEnterpriseAddressModifyResponseModel
	isSet bool
}

func (v NullableAlipayCommerceEcEnterpriseAddressModifyResponseModel) Get() *AlipayCommerceEcEnterpriseAddressModifyResponseModel {
	return v.value
}

func (v *NullableAlipayCommerceEcEnterpriseAddressModifyResponseModel) Set(val *AlipayCommerceEcEnterpriseAddressModifyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcEnterpriseAddressModifyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcEnterpriseAddressModifyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcEnterpriseAddressModifyResponseModel(val *AlipayCommerceEcEnterpriseAddressModifyResponseModel) *NullableAlipayCommerceEcEnterpriseAddressModifyResponseModel {
	return &NullableAlipayCommerceEcEnterpriseAddressModifyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcEnterpriseAddressModifyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcEnterpriseAddressModifyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
