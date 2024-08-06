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

// checks if the AlipayMerchantItemFileUploadModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMerchantItemFileUploadModel{}

// AlipayMerchantItemFileUploadModel struct for AlipayMerchantItemFileUploadModel
type AlipayMerchantItemFileUploadModel struct {
	// 业务场景描述。 小程序订单中心场景固定为 SYNC_ORDER。
	Scene *string `json:"scene,omitempty"`
}

// NewAlipayMerchantItemFileUploadModel instantiates a new AlipayMerchantItemFileUploadModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMerchantItemFileUploadModel() *AlipayMerchantItemFileUploadModel {
	this := AlipayMerchantItemFileUploadModel{}
	return &this
}

// NewAlipayMerchantItemFileUploadModelWithDefaults instantiates a new AlipayMerchantItemFileUploadModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMerchantItemFileUploadModelWithDefaults() *AlipayMerchantItemFileUploadModel {
	this := AlipayMerchantItemFileUploadModel{}
	return &this
}

// GetScene returns the Scene field value if set, zero value otherwise.
func (o *AlipayMerchantItemFileUploadModel) GetScene() string {
	if o == nil || IsNil(o.Scene) {
		var ret string
		return ret
	}
	return *o.Scene
}

// GetSceneOk returns a tuple with the Scene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantItemFileUploadModel) GetSceneOk() (*string, bool) {
	if o == nil || IsNil(o.Scene) {
		return nil, false
	}
	return o.Scene, true
}

// HasScene returns a boolean if a field has been set.
func (o *AlipayMerchantItemFileUploadModel) HasScene() bool {
	if o != nil && !IsNil(o.Scene) {
		return true
	}

	return false
}

// SetScene gets a reference to the given string and assigns it to the Scene field.
func (o *AlipayMerchantItemFileUploadModel) SetScene(v string) {
	o.Scene = &v
}

func (o AlipayMerchantItemFileUploadModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMerchantItemFileUploadModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Scene) {
		toSerialize["scene"] = o.Scene
	}
	return toSerialize, nil
}

type NullableAlipayMerchantItemFileUploadModel struct {
	value *AlipayMerchantItemFileUploadModel
	isSet bool
}

func (v NullableAlipayMerchantItemFileUploadModel) Get() *AlipayMerchantItemFileUploadModel {
	return v.value
}

func (v *NullableAlipayMerchantItemFileUploadModel) Set(val *AlipayMerchantItemFileUploadModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMerchantItemFileUploadModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMerchantItemFileUploadModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMerchantItemFileUploadModel(val *AlipayMerchantItemFileUploadModel) *NullableAlipayMerchantItemFileUploadModel {
	return &NullableAlipayMerchantItemFileUploadModel{value: val, isSet: true}
}

func (v NullableAlipayMerchantItemFileUploadModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMerchantItemFileUploadModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

