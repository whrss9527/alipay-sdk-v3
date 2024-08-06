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

// checks if the MiniAppAuditReasonMemo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MiniAppAuditReasonMemo{}

// MiniAppAuditReasonMemo struct for MiniAppAuditReasonMemo
type MiniAppAuditReasonMemo struct {
	// 驳回原因
	Memo *string `json:"memo,omitempty"`
	// 规则图片链接
	MemoImageList []string `json:"memo_image_list,omitempty"`
}

// NewMiniAppAuditReasonMemo instantiates a new MiniAppAuditReasonMemo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiniAppAuditReasonMemo() *MiniAppAuditReasonMemo {
	this := MiniAppAuditReasonMemo{}
	return &this
}

// NewMiniAppAuditReasonMemoWithDefaults instantiates a new MiniAppAuditReasonMemo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiniAppAuditReasonMemoWithDefaults() *MiniAppAuditReasonMemo {
	this := MiniAppAuditReasonMemo{}
	return &this
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *MiniAppAuditReasonMemo) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppAuditReasonMemo) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *MiniAppAuditReasonMemo) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *MiniAppAuditReasonMemo) SetMemo(v string) {
	o.Memo = &v
}

// GetMemoImageList returns the MemoImageList field value if set, zero value otherwise.
func (o *MiniAppAuditReasonMemo) GetMemoImageList() []string {
	if o == nil || IsNil(o.MemoImageList) {
		var ret []string
		return ret
	}
	return o.MemoImageList
}

// GetMemoImageListOk returns a tuple with the MemoImageList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniAppAuditReasonMemo) GetMemoImageListOk() ([]string, bool) {
	if o == nil || IsNil(o.MemoImageList) {
		return nil, false
	}
	return o.MemoImageList, true
}

// HasMemoImageList returns a boolean if a field has been set.
func (o *MiniAppAuditReasonMemo) HasMemoImageList() bool {
	if o != nil && !IsNil(o.MemoImageList) {
		return true
	}

	return false
}

// SetMemoImageList gets a reference to the given []string and assigns it to the MemoImageList field.
func (o *MiniAppAuditReasonMemo) SetMemoImageList(v []string) {
	o.MemoImageList = v
}

func (o MiniAppAuditReasonMemo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiniAppAuditReasonMemo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.MemoImageList) {
		toSerialize["memo_image_list"] = o.MemoImageList
	}
	return toSerialize, nil
}

type NullableMiniAppAuditReasonMemo struct {
	value *MiniAppAuditReasonMemo
	isSet bool
}

func (v NullableMiniAppAuditReasonMemo) Get() *MiniAppAuditReasonMemo {
	return v.value
}

func (v *NullableMiniAppAuditReasonMemo) Set(val *MiniAppAuditReasonMemo) {
	v.value = val
	v.isSet = true
}

func (v NullableMiniAppAuditReasonMemo) IsSet() bool {
	return v.isSet
}

func (v *NullableMiniAppAuditReasonMemo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiniAppAuditReasonMemo(val *MiniAppAuditReasonMemo) *NullableMiniAppAuditReasonMemo {
	return &NullableMiniAppAuditReasonMemo{value: val, isSet: true}
}

func (v NullableMiniAppAuditReasonMemo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiniAppAuditReasonMemo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

