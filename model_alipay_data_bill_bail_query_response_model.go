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

// checks if the AlipayDataBillBailQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayDataBillBailQueryResponseModel{}

// AlipayDataBillBailQueryResponseModel struct for AlipayDataBillBailQueryResponseModel
type AlipayDataBillBailQueryResponseModel struct {
	// 保证金明细列表，最多返回5000条
	DetailList []BailDetailResult `json:"detail_list,omitempty"`
}

// NewAlipayDataBillBailQueryResponseModel instantiates a new AlipayDataBillBailQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayDataBillBailQueryResponseModel() *AlipayDataBillBailQueryResponseModel {
	this := AlipayDataBillBailQueryResponseModel{}
	return &this
}

// NewAlipayDataBillBailQueryResponseModelWithDefaults instantiates a new AlipayDataBillBailQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayDataBillBailQueryResponseModelWithDefaults() *AlipayDataBillBailQueryResponseModel {
	this := AlipayDataBillBailQueryResponseModel{}
	return &this
}

// GetDetailList returns the DetailList field value if set, zero value otherwise.
func (o *AlipayDataBillBailQueryResponseModel) GetDetailList() []BailDetailResult {
	if o == nil || IsNil(o.DetailList) {
		var ret []BailDetailResult
		return ret
	}
	return o.DetailList
}

// GetDetailListOk returns a tuple with the DetailList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayDataBillBailQueryResponseModel) GetDetailListOk() ([]BailDetailResult, bool) {
	if o == nil || IsNil(o.DetailList) {
		return nil, false
	}
	return o.DetailList, true
}

// HasDetailList returns a boolean if a field has been set.
func (o *AlipayDataBillBailQueryResponseModel) HasDetailList() bool {
	if o != nil && !IsNil(o.DetailList) {
		return true
	}

	return false
}

// SetDetailList gets a reference to the given []BailDetailResult and assigns it to the DetailList field.
func (o *AlipayDataBillBailQueryResponseModel) SetDetailList(v []BailDetailResult) {
	o.DetailList = v
}

func (o AlipayDataBillBailQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayDataBillBailQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DetailList) {
		toSerialize["detail_list"] = o.DetailList
	}
	return toSerialize, nil
}

type NullableAlipayDataBillBailQueryResponseModel struct {
	value *AlipayDataBillBailQueryResponseModel
	isSet bool
}

func (v NullableAlipayDataBillBailQueryResponseModel) Get() *AlipayDataBillBailQueryResponseModel {
	return v.value
}

func (v *NullableAlipayDataBillBailQueryResponseModel) Set(val *AlipayDataBillBailQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayDataBillBailQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayDataBillBailQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayDataBillBailQueryResponseModel(val *AlipayDataBillBailQueryResponseModel) *NullableAlipayDataBillBailQueryResponseModel {
	return &NullableAlipayDataBillBailQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayDataBillBailQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayDataBillBailQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

