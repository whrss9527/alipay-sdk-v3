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

// checks if the AlipayEbppInvoiceSyncSimpleSendModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceSyncSimpleSendModel{}

// AlipayEbppInvoiceSyncSimpleSendModel struct for AlipayEbppInvoiceSyncSimpleSendModel
type AlipayEbppInvoiceSyncSimpleSendModel struct {
	InvoiceInfo *InvoicePDFSynModel `json:"invoice_info,omitempty"`
	// 开票商户品牌简称，与商户入驻时的品牌简称保持一致。详情参见 <a href=\"https://opendocs.alipay.com/open/10691/welcome-to-lark\">电子发票</a>
	MShortName *string `json:"m_short_name,omitempty"`
	// 开票商户门店简称，与商户入驻时的门店简称保持一致。
	SubMShortName *string `json:"sub_m_short_name,omitempty"`
}

// NewAlipayEbppInvoiceSyncSimpleSendModel instantiates a new AlipayEbppInvoiceSyncSimpleSendModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceSyncSimpleSendModel() *AlipayEbppInvoiceSyncSimpleSendModel {
	this := AlipayEbppInvoiceSyncSimpleSendModel{}
	return &this
}

// NewAlipayEbppInvoiceSyncSimpleSendModelWithDefaults instantiates a new AlipayEbppInvoiceSyncSimpleSendModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceSyncSimpleSendModelWithDefaults() *AlipayEbppInvoiceSyncSimpleSendModel {
	this := AlipayEbppInvoiceSyncSimpleSendModel{}
	return &this
}

// GetInvoiceInfo returns the InvoiceInfo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceSyncSimpleSendModel) GetInvoiceInfo() InvoicePDFSynModel {
	if o == nil || IsNil(o.InvoiceInfo) {
		var ret InvoicePDFSynModel
		return ret
	}
	return *o.InvoiceInfo
}

// GetInvoiceInfoOk returns a tuple with the InvoiceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceSyncSimpleSendModel) GetInvoiceInfoOk() (*InvoicePDFSynModel, bool) {
	if o == nil || IsNil(o.InvoiceInfo) {
		return nil, false
	}
	return o.InvoiceInfo, true
}

// HasInvoiceInfo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceSyncSimpleSendModel) HasInvoiceInfo() bool {
	if o != nil && !IsNil(o.InvoiceInfo) {
		return true
	}

	return false
}

// SetInvoiceInfo gets a reference to the given InvoicePDFSynModel and assigns it to the InvoiceInfo field.
func (o *AlipayEbppInvoiceSyncSimpleSendModel) SetInvoiceInfo(v InvoicePDFSynModel) {
	o.InvoiceInfo = &v
}

// GetMShortName returns the MShortName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceSyncSimpleSendModel) GetMShortName() string {
	if o == nil || IsNil(o.MShortName) {
		var ret string
		return ret
	}
	return *o.MShortName
}

// GetMShortNameOk returns a tuple with the MShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceSyncSimpleSendModel) GetMShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.MShortName) {
		return nil, false
	}
	return o.MShortName, true
}

// HasMShortName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceSyncSimpleSendModel) HasMShortName() bool {
	if o != nil && !IsNil(o.MShortName) {
		return true
	}

	return false
}

// SetMShortName gets a reference to the given string and assigns it to the MShortName field.
func (o *AlipayEbppInvoiceSyncSimpleSendModel) SetMShortName(v string) {
	o.MShortName = &v
}

// GetSubMShortName returns the SubMShortName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceSyncSimpleSendModel) GetSubMShortName() string {
	if o == nil || IsNil(o.SubMShortName) {
		var ret string
		return ret
	}
	return *o.SubMShortName
}

// GetSubMShortNameOk returns a tuple with the SubMShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceSyncSimpleSendModel) GetSubMShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubMShortName) {
		return nil, false
	}
	return o.SubMShortName, true
}

// HasSubMShortName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceSyncSimpleSendModel) HasSubMShortName() bool {
	if o != nil && !IsNil(o.SubMShortName) {
		return true
	}

	return false
}

// SetSubMShortName gets a reference to the given string and assigns it to the SubMShortName field.
func (o *AlipayEbppInvoiceSyncSimpleSendModel) SetSubMShortName(v string) {
	o.SubMShortName = &v
}

func (o AlipayEbppInvoiceSyncSimpleSendModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceSyncSimpleSendModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceInfo) {
		toSerialize["invoice_info"] = o.InvoiceInfo
	}
	if !IsNil(o.MShortName) {
		toSerialize["m_short_name"] = o.MShortName
	}
	if !IsNil(o.SubMShortName) {
		toSerialize["sub_m_short_name"] = o.SubMShortName
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceSyncSimpleSendModel struct {
	value *AlipayEbppInvoiceSyncSimpleSendModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceSyncSimpleSendModel) Get() *AlipayEbppInvoiceSyncSimpleSendModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceSyncSimpleSendModel) Set(val *AlipayEbppInvoiceSyncSimpleSendModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceSyncSimpleSendModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceSyncSimpleSendModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceSyncSimpleSendModel(val *AlipayEbppInvoiceSyncSimpleSendModel) *NullableAlipayEbppInvoiceSyncSimpleSendModel {
	return &NullableAlipayEbppInvoiceSyncSimpleSendModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceSyncSimpleSendModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceSyncSimpleSendModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
