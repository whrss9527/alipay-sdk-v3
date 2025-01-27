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

// checks if the AlipayIserviceCcmInstanceCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmInstanceCreateResponseModel{}

// AlipayIserviceCcmInstanceCreateResponseModel struct for AlipayIserviceCcmInstanceCreateResponseModel
type AlipayIserviceCcmInstanceCreateResponseModel struct {
	// 租户实例（数据权限）ID
	Id *string `json:"id,omitempty"`
}

// NewAlipayIserviceCcmInstanceCreateResponseModel instantiates a new AlipayIserviceCcmInstanceCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmInstanceCreateResponseModel() *AlipayIserviceCcmInstanceCreateResponseModel {
	this := AlipayIserviceCcmInstanceCreateResponseModel{}
	return &this
}

// NewAlipayIserviceCcmInstanceCreateResponseModelWithDefaults instantiates a new AlipayIserviceCcmInstanceCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmInstanceCreateResponseModelWithDefaults() *AlipayIserviceCcmInstanceCreateResponseModel {
	this := AlipayIserviceCcmInstanceCreateResponseModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlipayIserviceCcmInstanceCreateResponseModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmInstanceCreateResponseModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmInstanceCreateResponseModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AlipayIserviceCcmInstanceCreateResponseModel) SetId(v string) {
	o.Id = &v
}

func (o AlipayIserviceCcmInstanceCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmInstanceCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmInstanceCreateResponseModel struct {
	value *AlipayIserviceCcmInstanceCreateResponseModel
	isSet bool
}

func (v NullableAlipayIserviceCcmInstanceCreateResponseModel) Get() *AlipayIserviceCcmInstanceCreateResponseModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmInstanceCreateResponseModel) Set(val *AlipayIserviceCcmInstanceCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmInstanceCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmInstanceCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmInstanceCreateResponseModel(val *AlipayIserviceCcmInstanceCreateResponseModel) *NullableAlipayIserviceCcmInstanceCreateResponseModel {
	return &NullableAlipayIserviceCcmInstanceCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmInstanceCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmInstanceCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
