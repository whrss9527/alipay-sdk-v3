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

// checks if the AlipayOpenAppApiQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenAppApiQueryResponseModel{}

// AlipayOpenAppApiQueryResponseModel struct for AlipayOpenAppApiQueryResponseModel
type AlipayOpenAppApiQueryResponseModel struct {
	// 应用可申请的接口出参敏感字段列表
	Apis []AuthApiDTO `json:"apis,omitempty"`
}

// NewAlipayOpenAppApiQueryResponseModel instantiates a new AlipayOpenAppApiQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenAppApiQueryResponseModel() *AlipayOpenAppApiQueryResponseModel {
	this := AlipayOpenAppApiQueryResponseModel{}
	return &this
}

// NewAlipayOpenAppApiQueryResponseModelWithDefaults instantiates a new AlipayOpenAppApiQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenAppApiQueryResponseModelWithDefaults() *AlipayOpenAppApiQueryResponseModel {
	this := AlipayOpenAppApiQueryResponseModel{}
	return &this
}

// GetApis returns the Apis field value if set, zero value otherwise.
func (o *AlipayOpenAppApiQueryResponseModel) GetApis() []AuthApiDTO {
	if o == nil || IsNil(o.Apis) {
		var ret []AuthApiDTO
		return ret
	}
	return o.Apis
}

// GetApisOk returns a tuple with the Apis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppApiQueryResponseModel) GetApisOk() ([]AuthApiDTO, bool) {
	if o == nil || IsNil(o.Apis) {
		return nil, false
	}
	return o.Apis, true
}

// HasApis returns a boolean if a field has been set.
func (o *AlipayOpenAppApiQueryResponseModel) HasApis() bool {
	if o != nil && !IsNil(o.Apis) {
		return true
	}

	return false
}

// SetApis gets a reference to the given []AuthApiDTO and assigns it to the Apis field.
func (o *AlipayOpenAppApiQueryResponseModel) SetApis(v []AuthApiDTO) {
	o.Apis = v
}

func (o AlipayOpenAppApiQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenAppApiQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apis) {
		toSerialize["apis"] = o.Apis
	}
	return toSerialize, nil
}

type NullableAlipayOpenAppApiQueryResponseModel struct {
	value *AlipayOpenAppApiQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenAppApiQueryResponseModel) Get() *AlipayOpenAppApiQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenAppApiQueryResponseModel) Set(val *AlipayOpenAppApiQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAppApiQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAppApiQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAppApiQueryResponseModel(val *AlipayOpenAppApiQueryResponseModel) *NullableAlipayOpenAppApiQueryResponseModel {
	return &NullableAlipayOpenAppApiQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenAppApiQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAppApiQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
