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

// checks if the IndustryQualificationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndustryQualificationInfo{}

// IndustryQualificationInfo struct for IndustryQualificationInfo
type IndustryQualificationInfo struct {
	// 商户行业资质图片。其值为通过ant.merchant.expand.indirect.image.upload上传图片得到的image_id
	IndustryQualificationImage *string `json:"industry_qualification_image,omitempty"`
	// <a href=\"https://gw.alipayobjects.com/os/bmw-prod/7aa3a36b-2bc2-4d57-815f-08edd55ef67e.xlsx\">商户行业资质类型，具体选值参见文档</a>
	IndustryQualificationType *string `json:"industry_qualification_type,omitempty"`
}

// NewIndustryQualificationInfo instantiates a new IndustryQualificationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndustryQualificationInfo() *IndustryQualificationInfo {
	this := IndustryQualificationInfo{}
	return &this
}

// NewIndustryQualificationInfoWithDefaults instantiates a new IndustryQualificationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndustryQualificationInfoWithDefaults() *IndustryQualificationInfo {
	this := IndustryQualificationInfo{}
	return &this
}

// GetIndustryQualificationImage returns the IndustryQualificationImage field value if set, zero value otherwise.
func (o *IndustryQualificationInfo) GetIndustryQualificationImage() string {
	if o == nil || IsNil(o.IndustryQualificationImage) {
		var ret string
		return ret
	}
	return *o.IndustryQualificationImage
}

// GetIndustryQualificationImageOk returns a tuple with the IndustryQualificationImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndustryQualificationInfo) GetIndustryQualificationImageOk() (*string, bool) {
	if o == nil || IsNil(o.IndustryQualificationImage) {
		return nil, false
	}
	return o.IndustryQualificationImage, true
}

// HasIndustryQualificationImage returns a boolean if a field has been set.
func (o *IndustryQualificationInfo) HasIndustryQualificationImage() bool {
	if o != nil && !IsNil(o.IndustryQualificationImage) {
		return true
	}

	return false
}

// SetIndustryQualificationImage gets a reference to the given string and assigns it to the IndustryQualificationImage field.
func (o *IndustryQualificationInfo) SetIndustryQualificationImage(v string) {
	o.IndustryQualificationImage = &v
}

// GetIndustryQualificationType returns the IndustryQualificationType field value if set, zero value otherwise.
func (o *IndustryQualificationInfo) GetIndustryQualificationType() string {
	if o == nil || IsNil(o.IndustryQualificationType) {
		var ret string
		return ret
	}
	return *o.IndustryQualificationType
}

// GetIndustryQualificationTypeOk returns a tuple with the IndustryQualificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndustryQualificationInfo) GetIndustryQualificationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IndustryQualificationType) {
		return nil, false
	}
	return o.IndustryQualificationType, true
}

// HasIndustryQualificationType returns a boolean if a field has been set.
func (o *IndustryQualificationInfo) HasIndustryQualificationType() bool {
	if o != nil && !IsNil(o.IndustryQualificationType) {
		return true
	}

	return false
}

// SetIndustryQualificationType gets a reference to the given string and assigns it to the IndustryQualificationType field.
func (o *IndustryQualificationInfo) SetIndustryQualificationType(v string) {
	o.IndustryQualificationType = &v
}

func (o IndustryQualificationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndustryQualificationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IndustryQualificationImage) {
		toSerialize["industry_qualification_image"] = o.IndustryQualificationImage
	}
	if !IsNil(o.IndustryQualificationType) {
		toSerialize["industry_qualification_type"] = o.IndustryQualificationType
	}
	return toSerialize, nil
}

type NullableIndustryQualificationInfo struct {
	value *IndustryQualificationInfo
	isSet bool
}

func (v NullableIndustryQualificationInfo) Get() *IndustryQualificationInfo {
	return v.value
}

func (v *NullableIndustryQualificationInfo) Set(val *IndustryQualificationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIndustryQualificationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIndustryQualificationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndustryQualificationInfo(val *IndustryQualificationInfo) *NullableIndustryQualificationInfo {
	return &NullableIndustryQualificationInfo{value: val, isSet: true}
}

func (v NullableIndustryQualificationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndustryQualificationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
