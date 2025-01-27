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

// checks if the AlipayBossFncInvoiceApplyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayBossFncInvoiceApplyResponseModel{}

// AlipayBossFncInvoiceApplyResponseModel struct for AlipayBossFncInvoiceApplyResponseModel
type AlipayBossFncInvoiceApplyResponseModel struct {
	// 开票申请ID，唯一性ID
	ResultSet *string `json:"result_set,omitempty"`
}

// NewAlipayBossFncInvoiceApplyResponseModel instantiates a new AlipayBossFncInvoiceApplyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayBossFncInvoiceApplyResponseModel() *AlipayBossFncInvoiceApplyResponseModel {
	this := AlipayBossFncInvoiceApplyResponseModel{}
	return &this
}

// NewAlipayBossFncInvoiceApplyResponseModelWithDefaults instantiates a new AlipayBossFncInvoiceApplyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayBossFncInvoiceApplyResponseModelWithDefaults() *AlipayBossFncInvoiceApplyResponseModel {
	this := AlipayBossFncInvoiceApplyResponseModel{}
	return &this
}

// GetResultSet returns the ResultSet field value if set, zero value otherwise.
func (o *AlipayBossFncInvoiceApplyResponseModel) GetResultSet() string {
	if o == nil || IsNil(o.ResultSet) {
		var ret string
		return ret
	}
	return *o.ResultSet
}

// GetResultSetOk returns a tuple with the ResultSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayBossFncInvoiceApplyResponseModel) GetResultSetOk() (*string, bool) {
	if o == nil || IsNil(o.ResultSet) {
		return nil, false
	}
	return o.ResultSet, true
}

// HasResultSet returns a boolean if a field has been set.
func (o *AlipayBossFncInvoiceApplyResponseModel) HasResultSet() bool {
	if o != nil && !IsNil(o.ResultSet) {
		return true
	}

	return false
}

// SetResultSet gets a reference to the given string and assigns it to the ResultSet field.
func (o *AlipayBossFncInvoiceApplyResponseModel) SetResultSet(v string) {
	o.ResultSet = &v
}

func (o AlipayBossFncInvoiceApplyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayBossFncInvoiceApplyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultSet) {
		toSerialize["result_set"] = o.ResultSet
	}
	return toSerialize, nil
}

type NullableAlipayBossFncInvoiceApplyResponseModel struct {
	value *AlipayBossFncInvoiceApplyResponseModel
	isSet bool
}

func (v NullableAlipayBossFncInvoiceApplyResponseModel) Get() *AlipayBossFncInvoiceApplyResponseModel {
	return v.value
}

func (v *NullableAlipayBossFncInvoiceApplyResponseModel) Set(val *AlipayBossFncInvoiceApplyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayBossFncInvoiceApplyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayBossFncInvoiceApplyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayBossFncInvoiceApplyResponseModel(val *AlipayBossFncInvoiceApplyResponseModel) *NullableAlipayBossFncInvoiceApplyResponseModel {
	return &NullableAlipayBossFncInvoiceApplyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayBossFncInvoiceApplyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayBossFncInvoiceApplyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
