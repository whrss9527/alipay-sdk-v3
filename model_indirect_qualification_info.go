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

// checks if the IndirectQualificationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndirectQualificationInfo{}

// IndirectQualificationInfo struct for IndirectQualificationInfo
type IndirectQualificationInfo struct {
	// 行业经营许可证资质照片，一个行业类目下最多上传6张资质照片（使用图片上传接口）
	ImageList []string `json:"image_list,omitempty"`
	// 行业类目编号，支付宝商家行业二级类目code
	MccCode *string `json:"mcc_code,omitempty"`
}

// NewIndirectQualificationInfo instantiates a new IndirectQualificationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndirectQualificationInfo() *IndirectQualificationInfo {
	this := IndirectQualificationInfo{}
	return &this
}

// NewIndirectQualificationInfoWithDefaults instantiates a new IndirectQualificationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndirectQualificationInfoWithDefaults() *IndirectQualificationInfo {
	this := IndirectQualificationInfo{}
	return &this
}

// GetImageList returns the ImageList field value if set, zero value otherwise.
func (o *IndirectQualificationInfo) GetImageList() []string {
	if o == nil || IsNil(o.ImageList) {
		var ret []string
		return ret
	}
	return o.ImageList
}

// GetImageListOk returns a tuple with the ImageList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndirectQualificationInfo) GetImageListOk() ([]string, bool) {
	if o == nil || IsNil(o.ImageList) {
		return nil, false
	}
	return o.ImageList, true
}

// HasImageList returns a boolean if a field has been set.
func (o *IndirectQualificationInfo) HasImageList() bool {
	if o != nil && !IsNil(o.ImageList) {
		return true
	}

	return false
}

// SetImageList gets a reference to the given []string and assigns it to the ImageList field.
func (o *IndirectQualificationInfo) SetImageList(v []string) {
	o.ImageList = v
}

// GetMccCode returns the MccCode field value if set, zero value otherwise.
func (o *IndirectQualificationInfo) GetMccCode() string {
	if o == nil || IsNil(o.MccCode) {
		var ret string
		return ret
	}
	return *o.MccCode
}

// GetMccCodeOk returns a tuple with the MccCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndirectQualificationInfo) GetMccCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MccCode) {
		return nil, false
	}
	return o.MccCode, true
}

// HasMccCode returns a boolean if a field has been set.
func (o *IndirectQualificationInfo) HasMccCode() bool {
	if o != nil && !IsNil(o.MccCode) {
		return true
	}

	return false
}

// SetMccCode gets a reference to the given string and assigns it to the MccCode field.
func (o *IndirectQualificationInfo) SetMccCode(v string) {
	o.MccCode = &v
}

func (o IndirectQualificationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndirectQualificationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageList) {
		toSerialize["image_list"] = o.ImageList
	}
	if !IsNil(o.MccCode) {
		toSerialize["mcc_code"] = o.MccCode
	}
	return toSerialize, nil
}

type NullableIndirectQualificationInfo struct {
	value *IndirectQualificationInfo
	isSet bool
}

func (v NullableIndirectQualificationInfo) Get() *IndirectQualificationInfo {
	return v.value
}

func (v *NullableIndirectQualificationInfo) Set(val *IndirectQualificationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIndirectQualificationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIndirectQualificationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndirectQualificationInfo(val *IndirectQualificationInfo) *NullableIndirectQualificationInfo {
	return &NullableIndirectQualificationInfo{value: val, isSet: true}
}

func (v NullableIndirectQualificationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndirectQualificationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
