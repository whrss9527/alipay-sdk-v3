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

// checks if the AlipayMerchantImageUploadResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMerchantImageUploadResponseModel{}

// AlipayMerchantImageUploadResponseModel struct for AlipayMerchantImageUploadResponseModel
type AlipayMerchantImageUploadResponseModel struct {
	// 图片在文件存储平台的标识
	ImageId *string `json:"image_id,omitempty"`
}

// NewAlipayMerchantImageUploadResponseModel instantiates a new AlipayMerchantImageUploadResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMerchantImageUploadResponseModel() *AlipayMerchantImageUploadResponseModel {
	this := AlipayMerchantImageUploadResponseModel{}
	return &this
}

// NewAlipayMerchantImageUploadResponseModelWithDefaults instantiates a new AlipayMerchantImageUploadResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMerchantImageUploadResponseModelWithDefaults() *AlipayMerchantImageUploadResponseModel {
	this := AlipayMerchantImageUploadResponseModel{}
	return &this
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *AlipayMerchantImageUploadResponseModel) GetImageId() string {
	if o == nil || IsNil(o.ImageId) {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantImageUploadResponseModel) GetImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *AlipayMerchantImageUploadResponseModel) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *AlipayMerchantImageUploadResponseModel) SetImageId(v string) {
	o.ImageId = &v
}

func (o AlipayMerchantImageUploadResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMerchantImageUploadResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageId) {
		toSerialize["image_id"] = o.ImageId
	}
	return toSerialize, nil
}

type NullableAlipayMerchantImageUploadResponseModel struct {
	value *AlipayMerchantImageUploadResponseModel
	isSet bool
}

func (v NullableAlipayMerchantImageUploadResponseModel) Get() *AlipayMerchantImageUploadResponseModel {
	return v.value
}

func (v *NullableAlipayMerchantImageUploadResponseModel) Set(val *AlipayMerchantImageUploadResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMerchantImageUploadResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMerchantImageUploadResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMerchantImageUploadResponseModel(val *AlipayMerchantImageUploadResponseModel) *NullableAlipayMerchantImageUploadResponseModel {
	return &NullableAlipayMerchantImageUploadResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMerchantImageUploadResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMerchantImageUploadResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
