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

// checks if the OcrNormalScanInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OcrNormalScanInfo{}

// OcrNormalScanInfo struct for OcrNormalScanInfo
type OcrNormalScanInfo struct {
	// 发票代码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// 开票时间
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// 金额（元）
	Price *string `json:"price,omitempty"`
	// 明细事由
	Remark *string `json:"remark,omitempty"`
}

// NewOcrNormalScanInfo instantiates a new OcrNormalScanInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOcrNormalScanInfo() *OcrNormalScanInfo {
	this := OcrNormalScanInfo{}
	return &this
}

// NewOcrNormalScanInfoWithDefaults instantiates a new OcrNormalScanInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOcrNormalScanInfoWithDefaults() *OcrNormalScanInfo {
	this := OcrNormalScanInfo{}
	return &this
}

// GetInvoiceCode returns the InvoiceCode field value if set, zero value otherwise.
func (o *OcrNormalScanInfo) GetInvoiceCode() string {
	if o == nil || IsNil(o.InvoiceCode) {
		var ret string
		return ret
	}
	return *o.InvoiceCode
}

// GetInvoiceCodeOk returns a tuple with the InvoiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcrNormalScanInfo) GetInvoiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceCode) {
		return nil, false
	}
	return o.InvoiceCode, true
}

// HasInvoiceCode returns a boolean if a field has been set.
func (o *OcrNormalScanInfo) HasInvoiceCode() bool {
	if o != nil && !IsNil(o.InvoiceCode) {
		return true
	}

	return false
}

// SetInvoiceCode gets a reference to the given string and assigns it to the InvoiceCode field.
func (o *OcrNormalScanInfo) SetInvoiceCode(v string) {
	o.InvoiceCode = &v
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise.
func (o *OcrNormalScanInfo) GetInvoiceDate() string {
	if o == nil || IsNil(o.InvoiceDate) {
		var ret string
		return ret
	}
	return *o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcrNormalScanInfo) GetInvoiceDateOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDate) {
		return nil, false
	}
	return o.InvoiceDate, true
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *OcrNormalScanInfo) HasInvoiceDate() bool {
	if o != nil && !IsNil(o.InvoiceDate) {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given string and assigns it to the InvoiceDate field.
func (o *OcrNormalScanInfo) SetInvoiceDate(v string) {
	o.InvoiceDate = &v
}

// GetInvoiceNo returns the InvoiceNo field value if set, zero value otherwise.
func (o *OcrNormalScanInfo) GetInvoiceNo() string {
	if o == nil || IsNil(o.InvoiceNo) {
		var ret string
		return ret
	}
	return *o.InvoiceNo
}

// GetInvoiceNoOk returns a tuple with the InvoiceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcrNormalScanInfo) GetInvoiceNoOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNo) {
		return nil, false
	}
	return o.InvoiceNo, true
}

// HasInvoiceNo returns a boolean if a field has been set.
func (o *OcrNormalScanInfo) HasInvoiceNo() bool {
	if o != nil && !IsNil(o.InvoiceNo) {
		return true
	}

	return false
}

// SetInvoiceNo gets a reference to the given string and assigns it to the InvoiceNo field.
func (o *OcrNormalScanInfo) SetInvoiceNo(v string) {
	o.InvoiceNo = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OcrNormalScanInfo) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcrNormalScanInfo) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OcrNormalScanInfo) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *OcrNormalScanInfo) SetPrice(v string) {
	o.Price = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *OcrNormalScanInfo) GetRemark() string {
	if o == nil || IsNil(o.Remark) {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcrNormalScanInfo) GetRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.Remark) {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *OcrNormalScanInfo) HasRemark() bool {
	if o != nil && !IsNil(o.Remark) {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *OcrNormalScanInfo) SetRemark(v string) {
	o.Remark = &v
}

func (o OcrNormalScanInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OcrNormalScanInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceCode) {
		toSerialize["invoice_code"] = o.InvoiceCode
	}
	if !IsNil(o.InvoiceDate) {
		toSerialize["invoice_date"] = o.InvoiceDate
	}
	if !IsNil(o.InvoiceNo) {
		toSerialize["invoice_no"] = o.InvoiceNo
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Remark) {
		toSerialize["remark"] = o.Remark
	}
	return toSerialize, nil
}

type NullableOcrNormalScanInfo struct {
	value *OcrNormalScanInfo
	isSet bool
}

func (v NullableOcrNormalScanInfo) Get() *OcrNormalScanInfo {
	return v.value
}

func (v *NullableOcrNormalScanInfo) Set(val *OcrNormalScanInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOcrNormalScanInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOcrNormalScanInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOcrNormalScanInfo(val *OcrNormalScanInfo) *NullableOcrNormalScanInfo {
	return &NullableOcrNormalScanInfo{value: val, isSet: true}
}

func (v NullableOcrNormalScanInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOcrNormalScanInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
