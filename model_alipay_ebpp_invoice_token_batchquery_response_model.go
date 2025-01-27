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

// checks if the AlipayEbppInvoiceTokenBatchqueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceTokenBatchqueryResponseModel{}

// AlipayEbppInvoiceTokenBatchqueryResponseModel struct for AlipayEbppInvoiceTokenBatchqueryResponseModel
type AlipayEbppInvoiceTokenBatchqueryResponseModel struct {
	// 发票要素列表
	InvoiceElementList []InvoiceElementModel `json:"invoice_element_list,omitempty"`
	// 支付宝用户id
	OpenId *string `json:"open_id,omitempty"`
	// 支付宝用户id
	UserId *string `json:"user_id,omitempty"`
}

// NewAlipayEbppInvoiceTokenBatchqueryResponseModel instantiates a new AlipayEbppInvoiceTokenBatchqueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceTokenBatchqueryResponseModel() *AlipayEbppInvoiceTokenBatchqueryResponseModel {
	this := AlipayEbppInvoiceTokenBatchqueryResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceTokenBatchqueryResponseModelWithDefaults instantiates a new AlipayEbppInvoiceTokenBatchqueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceTokenBatchqueryResponseModelWithDefaults() *AlipayEbppInvoiceTokenBatchqueryResponseModel {
	this := AlipayEbppInvoiceTokenBatchqueryResponseModel{}
	return &this
}

// GetInvoiceElementList returns the InvoiceElementList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceTokenBatchqueryResponseModel) GetInvoiceElementList() []InvoiceElementModel {
	if o == nil || IsNil(o.InvoiceElementList) {
		var ret []InvoiceElementModel
		return ret
	}
	return o.InvoiceElementList
}

// GetInvoiceElementListOk returns a tuple with the InvoiceElementList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceTokenBatchqueryResponseModel) GetInvoiceElementListOk() ([]InvoiceElementModel, bool) {
	if o == nil || IsNil(o.InvoiceElementList) {
		return nil, false
	}
	return o.InvoiceElementList, true
}

// HasInvoiceElementList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceTokenBatchqueryResponseModel) HasInvoiceElementList() bool {
	if o != nil && !IsNil(o.InvoiceElementList) {
		return true
	}

	return false
}

// SetInvoiceElementList gets a reference to the given []InvoiceElementModel and assigns it to the InvoiceElementList field.
func (o *AlipayEbppInvoiceTokenBatchqueryResponseModel) SetInvoiceElementList(v []InvoiceElementModel) {
	o.InvoiceElementList = v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceTokenBatchqueryResponseModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceTokenBatchqueryResponseModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceTokenBatchqueryResponseModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AlipayEbppInvoiceTokenBatchqueryResponseModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceTokenBatchqueryResponseModel) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceTokenBatchqueryResponseModel) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceTokenBatchqueryResponseModel) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AlipayEbppInvoiceTokenBatchqueryResponseModel) SetUserId(v string) {
	o.UserId = &v
}

func (o AlipayEbppInvoiceTokenBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceTokenBatchqueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceElementList) {
		toSerialize["invoice_element_list"] = o.InvoiceElementList
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceTokenBatchqueryResponseModel struct {
	value *AlipayEbppInvoiceTokenBatchqueryResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceTokenBatchqueryResponseModel) Get() *AlipayEbppInvoiceTokenBatchqueryResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceTokenBatchqueryResponseModel) Set(val *AlipayEbppInvoiceTokenBatchqueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceTokenBatchqueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceTokenBatchqueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceTokenBatchqueryResponseModel(val *AlipayEbppInvoiceTokenBatchqueryResponseModel) *NullableAlipayEbppInvoiceTokenBatchqueryResponseModel {
	return &NullableAlipayEbppInvoiceTokenBatchqueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceTokenBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceTokenBatchqueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
