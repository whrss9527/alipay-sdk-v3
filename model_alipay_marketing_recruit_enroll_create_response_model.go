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

// checks if the AlipayMarketingRecruitEnrollCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingRecruitEnrollCreateResponseModel{}

// AlipayMarketingRecruitEnrollCreateResponseModel struct for AlipayMarketingRecruitEnrollCreateResponseModel
type AlipayMarketingRecruitEnrollCreateResponseModel struct {
	// 报名ID
	EnrollId *string `json:"enroll_id,omitempty"`
	// 外部操作流水号。由商家/ISV 自定义，仅支持字母、数字、下划线且需保证每次操作唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
}

// NewAlipayMarketingRecruitEnrollCreateResponseModel instantiates a new AlipayMarketingRecruitEnrollCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingRecruitEnrollCreateResponseModel() *AlipayMarketingRecruitEnrollCreateResponseModel {
	this := AlipayMarketingRecruitEnrollCreateResponseModel{}
	return &this
}

// NewAlipayMarketingRecruitEnrollCreateResponseModelWithDefaults instantiates a new AlipayMarketingRecruitEnrollCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingRecruitEnrollCreateResponseModelWithDefaults() *AlipayMarketingRecruitEnrollCreateResponseModel {
	this := AlipayMarketingRecruitEnrollCreateResponseModel{}
	return &this
}

// GetEnrollId returns the EnrollId field value if set, zero value otherwise.
func (o *AlipayMarketingRecruitEnrollCreateResponseModel) GetEnrollId() string {
	if o == nil || IsNil(o.EnrollId) {
		var ret string
		return ret
	}
	return *o.EnrollId
}

// GetEnrollIdOk returns a tuple with the EnrollId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingRecruitEnrollCreateResponseModel) GetEnrollIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnrollId) {
		return nil, false
	}
	return o.EnrollId, true
}

// HasEnrollId returns a boolean if a field has been set.
func (o *AlipayMarketingRecruitEnrollCreateResponseModel) HasEnrollId() bool {
	if o != nil && !IsNil(o.EnrollId) {
		return true
	}

	return false
}

// SetEnrollId gets a reference to the given string and assigns it to the EnrollId field.
func (o *AlipayMarketingRecruitEnrollCreateResponseModel) SetEnrollId(v string) {
	o.EnrollId = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayMarketingRecruitEnrollCreateResponseModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingRecruitEnrollCreateResponseModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayMarketingRecruitEnrollCreateResponseModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayMarketingRecruitEnrollCreateResponseModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

func (o AlipayMarketingRecruitEnrollCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingRecruitEnrollCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnrollId) {
		toSerialize["enroll_id"] = o.EnrollId
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	return toSerialize, nil
}

type NullableAlipayMarketingRecruitEnrollCreateResponseModel struct {
	value *AlipayMarketingRecruitEnrollCreateResponseModel
	isSet bool
}

func (v NullableAlipayMarketingRecruitEnrollCreateResponseModel) Get() *AlipayMarketingRecruitEnrollCreateResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingRecruitEnrollCreateResponseModel) Set(val *AlipayMarketingRecruitEnrollCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingRecruitEnrollCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingRecruitEnrollCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingRecruitEnrollCreateResponseModel(val *AlipayMarketingRecruitEnrollCreateResponseModel) *NullableAlipayMarketingRecruitEnrollCreateResponseModel {
	return &NullableAlipayMarketingRecruitEnrollCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingRecruitEnrollCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingRecruitEnrollCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

