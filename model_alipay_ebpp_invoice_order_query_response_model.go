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

// checks if the AlipayEbppInvoiceOrderQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceOrderQueryResponseModel{}

// AlipayEbppInvoiceOrderQueryResponseModel struct for AlipayEbppInvoiceOrderQueryResponseModel
type AlipayEbppInvoiceOrderQueryResponseModel struct {
	InvoiceElementModel *InvoiceElementModel `json:"invoice_element_model,omitempty"`
	// 开票申请时传入订单号（支持主单号、子单号），不限是否为支付宝体内交易单号
	OrderNo *string `json:"order_no,omitempty"`
}

// NewAlipayEbppInvoiceOrderQueryResponseModel instantiates a new AlipayEbppInvoiceOrderQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceOrderQueryResponseModel() *AlipayEbppInvoiceOrderQueryResponseModel {
	this := AlipayEbppInvoiceOrderQueryResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceOrderQueryResponseModelWithDefaults instantiates a new AlipayEbppInvoiceOrderQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceOrderQueryResponseModelWithDefaults() *AlipayEbppInvoiceOrderQueryResponseModel {
	this := AlipayEbppInvoiceOrderQueryResponseModel{}
	return &this
}

// GetInvoiceElementModel returns the InvoiceElementModel field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceOrderQueryResponseModel) GetInvoiceElementModel() InvoiceElementModel {
	if o == nil || IsNil(o.InvoiceElementModel) {
		var ret InvoiceElementModel
		return ret
	}
	return *o.InvoiceElementModel
}

// GetInvoiceElementModelOk returns a tuple with the InvoiceElementModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceOrderQueryResponseModel) GetInvoiceElementModelOk() (*InvoiceElementModel, bool) {
	if o == nil || IsNil(o.InvoiceElementModel) {
		return nil, false
	}
	return o.InvoiceElementModel, true
}

// HasInvoiceElementModel returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceOrderQueryResponseModel) HasInvoiceElementModel() bool {
	if o != nil && !IsNil(o.InvoiceElementModel) {
		return true
	}

	return false
}

// SetInvoiceElementModel gets a reference to the given InvoiceElementModel and assigns it to the InvoiceElementModel field.
func (o *AlipayEbppInvoiceOrderQueryResponseModel) SetInvoiceElementModel(v InvoiceElementModel) {
	o.InvoiceElementModel = &v
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceOrderQueryResponseModel) GetOrderNo() string {
	if o == nil || IsNil(o.OrderNo) {
		var ret string
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceOrderQueryResponseModel) GetOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceOrderQueryResponseModel) HasOrderNo() bool {
	if o != nil && !IsNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given string and assigns it to the OrderNo field.
func (o *AlipayEbppInvoiceOrderQueryResponseModel) SetOrderNo(v string) {
	o.OrderNo = &v
}

func (o AlipayEbppInvoiceOrderQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceOrderQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceElementModel) {
		toSerialize["invoice_element_model"] = o.InvoiceElementModel
	}
	if !IsNil(o.OrderNo) {
		toSerialize["order_no"] = o.OrderNo
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceOrderQueryResponseModel struct {
	value *AlipayEbppInvoiceOrderQueryResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceOrderQueryResponseModel) Get() *AlipayEbppInvoiceOrderQueryResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceOrderQueryResponseModel) Set(val *AlipayEbppInvoiceOrderQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceOrderQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceOrderQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceOrderQueryResponseModel(val *AlipayEbppInvoiceOrderQueryResponseModel) *NullableAlipayEbppInvoiceOrderQueryResponseModel {
	return &NullableAlipayEbppInvoiceOrderQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceOrderQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceOrderQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
