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

// checks if the AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel{}

// AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel struct for AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel
type AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel struct {
	// 是否成功
	Success *bool `json:"success,omitempty"`
}

// NewAlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel instantiates a new AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel() *AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel {
	this := AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceExpensecontrolQuotaModifyResponseModelWithDefaults instantiates a new AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceExpensecontrolQuotaModifyResponseModelWithDefaults() *AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel {
	this := AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel) SetSuccess(v bool) {
	o.Success = &v
}

func (o AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel struct {
	value *AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel) Get() *AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel) Set(val *AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel(val *AlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel) *NullableAlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel {
	return &NullableAlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpensecontrolQuotaModifyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

