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

// checks if the AlipayOpenPublicUserDataBatchqueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicUserDataBatchqueryResponseModel{}

// AlipayOpenPublicUserDataBatchqueryResponseModel struct for AlipayOpenPublicUserDataBatchqueryResponseModel
type AlipayOpenPublicUserDataBatchqueryResponseModel struct {
	// 用户分析数据
	DataList []UserAnalysisData `json:"data_list,omitempty"`
}

// NewAlipayOpenPublicUserDataBatchqueryResponseModel instantiates a new AlipayOpenPublicUserDataBatchqueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicUserDataBatchqueryResponseModel() *AlipayOpenPublicUserDataBatchqueryResponseModel {
	this := AlipayOpenPublicUserDataBatchqueryResponseModel{}
	return &this
}

// NewAlipayOpenPublicUserDataBatchqueryResponseModelWithDefaults instantiates a new AlipayOpenPublicUserDataBatchqueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicUserDataBatchqueryResponseModelWithDefaults() *AlipayOpenPublicUserDataBatchqueryResponseModel {
	this := AlipayOpenPublicUserDataBatchqueryResponseModel{}
	return &this
}

// GetDataList returns the DataList field value if set, zero value otherwise.
func (o *AlipayOpenPublicUserDataBatchqueryResponseModel) GetDataList() []UserAnalysisData {
	if o == nil || IsNil(o.DataList) {
		var ret []UserAnalysisData
		return ret
	}
	return o.DataList
}

// GetDataListOk returns a tuple with the DataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicUserDataBatchqueryResponseModel) GetDataListOk() ([]UserAnalysisData, bool) {
	if o == nil || IsNil(o.DataList) {
		return nil, false
	}
	return o.DataList, true
}

// HasDataList returns a boolean if a field has been set.
func (o *AlipayOpenPublicUserDataBatchqueryResponseModel) HasDataList() bool {
	if o != nil && !IsNil(o.DataList) {
		return true
	}

	return false
}

// SetDataList gets a reference to the given []UserAnalysisData and assigns it to the DataList field.
func (o *AlipayOpenPublicUserDataBatchqueryResponseModel) SetDataList(v []UserAnalysisData) {
	o.DataList = v
}

func (o AlipayOpenPublicUserDataBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicUserDataBatchqueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataList) {
		toSerialize["data_list"] = o.DataList
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicUserDataBatchqueryResponseModel struct {
	value *AlipayOpenPublicUserDataBatchqueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenPublicUserDataBatchqueryResponseModel) Get() *AlipayOpenPublicUserDataBatchqueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenPublicUserDataBatchqueryResponseModel) Set(val *AlipayOpenPublicUserDataBatchqueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicUserDataBatchqueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicUserDataBatchqueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicUserDataBatchqueryResponseModel(val *AlipayOpenPublicUserDataBatchqueryResponseModel) *NullableAlipayOpenPublicUserDataBatchqueryResponseModel {
	return &NullableAlipayOpenPublicUserDataBatchqueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicUserDataBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicUserDataBatchqueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
