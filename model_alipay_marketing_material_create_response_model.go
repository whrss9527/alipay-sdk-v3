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

// checks if the AlipayMarketingMaterialCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingMaterialCreateResponseModel{}

// AlipayMarketingMaterialCreateResponseModel struct for AlipayMarketingMaterialCreateResponseModel
type AlipayMarketingMaterialCreateResponseModel struct {
	// 素材ID。
	MaterialId *string `json:"material_id,omitempty"`
}

// NewAlipayMarketingMaterialCreateResponseModel instantiates a new AlipayMarketingMaterialCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingMaterialCreateResponseModel() *AlipayMarketingMaterialCreateResponseModel {
	this := AlipayMarketingMaterialCreateResponseModel{}
	return &this
}

// NewAlipayMarketingMaterialCreateResponseModelWithDefaults instantiates a new AlipayMarketingMaterialCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingMaterialCreateResponseModelWithDefaults() *AlipayMarketingMaterialCreateResponseModel {
	this := AlipayMarketingMaterialCreateResponseModel{}
	return &this
}

// GetMaterialId returns the MaterialId field value if set, zero value otherwise.
func (o *AlipayMarketingMaterialCreateResponseModel) GetMaterialId() string {
	if o == nil || IsNil(o.MaterialId) {
		var ret string
		return ret
	}
	return *o.MaterialId
}

// GetMaterialIdOk returns a tuple with the MaterialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingMaterialCreateResponseModel) GetMaterialIdOk() (*string, bool) {
	if o == nil || IsNil(o.MaterialId) {
		return nil, false
	}
	return o.MaterialId, true
}

// HasMaterialId returns a boolean if a field has been set.
func (o *AlipayMarketingMaterialCreateResponseModel) HasMaterialId() bool {
	if o != nil && !IsNil(o.MaterialId) {
		return true
	}

	return false
}

// SetMaterialId gets a reference to the given string and assigns it to the MaterialId field.
func (o *AlipayMarketingMaterialCreateResponseModel) SetMaterialId(v string) {
	o.MaterialId = &v
}

func (o AlipayMarketingMaterialCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingMaterialCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaterialId) {
		toSerialize["material_id"] = o.MaterialId
	}
	return toSerialize, nil
}

type NullableAlipayMarketingMaterialCreateResponseModel struct {
	value *AlipayMarketingMaterialCreateResponseModel
	isSet bool
}

func (v NullableAlipayMarketingMaterialCreateResponseModel) Get() *AlipayMarketingMaterialCreateResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingMaterialCreateResponseModel) Set(val *AlipayMarketingMaterialCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingMaterialCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingMaterialCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingMaterialCreateResponseModel(val *AlipayMarketingMaterialCreateResponseModel) *NullableAlipayMarketingMaterialCreateResponseModel {
	return &NullableAlipayMarketingMaterialCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingMaterialCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingMaterialCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

