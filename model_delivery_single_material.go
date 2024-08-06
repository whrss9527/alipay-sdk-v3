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

// checks if the DeliverySingleMaterial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliverySingleMaterial{}

// DeliverySingleMaterial struct for DeliverySingleMaterial
type DeliverySingleMaterial struct {
	// 投放计划图片素材。通过接口<a href =\"https://opendocs.alipay.com/open/049d64?pathHash=a411214d&scene=928cd58dc06f4466b9ad657776327a06\">alipay.marketing.material.image.upload</a>上传图片返回的resource_id。
	DeliveryImage *string `json:"delivery_image,omitempty"`
}

// NewDeliverySingleMaterial instantiates a new DeliverySingleMaterial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliverySingleMaterial() *DeliverySingleMaterial {
	this := DeliverySingleMaterial{}
	return &this
}

// NewDeliverySingleMaterialWithDefaults instantiates a new DeliverySingleMaterial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliverySingleMaterialWithDefaults() *DeliverySingleMaterial {
	this := DeliverySingleMaterial{}
	return &this
}

// GetDeliveryImage returns the DeliveryImage field value if set, zero value otherwise.
func (o *DeliverySingleMaterial) GetDeliveryImage() string {
	if o == nil || IsNil(o.DeliveryImage) {
		var ret string
		return ret
	}
	return *o.DeliveryImage
}

// GetDeliveryImageOk returns a tuple with the DeliveryImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverySingleMaterial) GetDeliveryImageOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryImage) {
		return nil, false
	}
	return o.DeliveryImage, true
}

// HasDeliveryImage returns a boolean if a field has been set.
func (o *DeliverySingleMaterial) HasDeliveryImage() bool {
	if o != nil && !IsNil(o.DeliveryImage) {
		return true
	}

	return false
}

// SetDeliveryImage gets a reference to the given string and assigns it to the DeliveryImage field.
func (o *DeliverySingleMaterial) SetDeliveryImage(v string) {
	o.DeliveryImage = &v
}

func (o DeliverySingleMaterial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliverySingleMaterial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeliveryImage) {
		toSerialize["delivery_image"] = o.DeliveryImage
	}
	return toSerialize, nil
}

type NullableDeliverySingleMaterial struct {
	value *DeliverySingleMaterial
	isSet bool
}

func (v NullableDeliverySingleMaterial) Get() *DeliverySingleMaterial {
	return v.value
}

func (v *NullableDeliverySingleMaterial) Set(val *DeliverySingleMaterial) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliverySingleMaterial) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliverySingleMaterial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliverySingleMaterial(val *DeliverySingleMaterial) *NullableDeliverySingleMaterial {
	return &NullableDeliverySingleMaterial{value: val, isSet: true}
}

func (v NullableDeliverySingleMaterial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliverySingleMaterial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

