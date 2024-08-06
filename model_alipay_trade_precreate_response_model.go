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

// checks if the AlipayTradePrecreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayTradePrecreateResponseModel{}

// AlipayTradePrecreateResponseModel struct for AlipayTradePrecreateResponseModel
type AlipayTradePrecreateResponseModel struct {
	// 商户的订单号
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// 当前预下单请求生成的二维码码串，有效时间2小时，可以用二维码生成工具根据该码串值生成对应的二维码
	QrCode *string `json:"qr_code,omitempty"`
	// 当前预下单请求生成的吱口令码串，有效时间2小时，可以在支付宝app端访问对应内容
	ShareCode *string `json:"share_code,omitempty"`
}

// NewAlipayTradePrecreateResponseModel instantiates a new AlipayTradePrecreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayTradePrecreateResponseModel() *AlipayTradePrecreateResponseModel {
	this := AlipayTradePrecreateResponseModel{}
	return &this
}

// NewAlipayTradePrecreateResponseModelWithDefaults instantiates a new AlipayTradePrecreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayTradePrecreateResponseModelWithDefaults() *AlipayTradePrecreateResponseModel {
	this := AlipayTradePrecreateResponseModel{}
	return &this
}

// GetOutTradeNo returns the OutTradeNo field value if set, zero value otherwise.
func (o *AlipayTradePrecreateResponseModel) GetOutTradeNo() string {
	if o == nil || IsNil(o.OutTradeNo) {
		var ret string
		return ret
	}
	return *o.OutTradeNo
}

// GetOutTradeNoOk returns a tuple with the OutTradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradePrecreateResponseModel) GetOutTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutTradeNo) {
		return nil, false
	}
	return o.OutTradeNo, true
}

// HasOutTradeNo returns a boolean if a field has been set.
func (o *AlipayTradePrecreateResponseModel) HasOutTradeNo() bool {
	if o != nil && !IsNil(o.OutTradeNo) {
		return true
	}

	return false
}

// SetOutTradeNo gets a reference to the given string and assigns it to the OutTradeNo field.
func (o *AlipayTradePrecreateResponseModel) SetOutTradeNo(v string) {
	o.OutTradeNo = &v
}

// GetQrCode returns the QrCode field value if set, zero value otherwise.
func (o *AlipayTradePrecreateResponseModel) GetQrCode() string {
	if o == nil || IsNil(o.QrCode) {
		var ret string
		return ret
	}
	return *o.QrCode
}

// GetQrCodeOk returns a tuple with the QrCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradePrecreateResponseModel) GetQrCodeOk() (*string, bool) {
	if o == nil || IsNil(o.QrCode) {
		return nil, false
	}
	return o.QrCode, true
}

// HasQrCode returns a boolean if a field has been set.
func (o *AlipayTradePrecreateResponseModel) HasQrCode() bool {
	if o != nil && !IsNil(o.QrCode) {
		return true
	}

	return false
}

// SetQrCode gets a reference to the given string and assigns it to the QrCode field.
func (o *AlipayTradePrecreateResponseModel) SetQrCode(v string) {
	o.QrCode = &v
}

// GetShareCode returns the ShareCode field value if set, zero value otherwise.
func (o *AlipayTradePrecreateResponseModel) GetShareCode() string {
	if o == nil || IsNil(o.ShareCode) {
		var ret string
		return ret
	}
	return *o.ShareCode
}

// GetShareCodeOk returns a tuple with the ShareCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradePrecreateResponseModel) GetShareCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ShareCode) {
		return nil, false
	}
	return o.ShareCode, true
}

// HasShareCode returns a boolean if a field has been set.
func (o *AlipayTradePrecreateResponseModel) HasShareCode() bool {
	if o != nil && !IsNil(o.ShareCode) {
		return true
	}

	return false
}

// SetShareCode gets a reference to the given string and assigns it to the ShareCode field.
func (o *AlipayTradePrecreateResponseModel) SetShareCode(v string) {
	o.ShareCode = &v
}

func (o AlipayTradePrecreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayTradePrecreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OutTradeNo) {
		toSerialize["out_trade_no"] = o.OutTradeNo
	}
	if !IsNil(o.QrCode) {
		toSerialize["qr_code"] = o.QrCode
	}
	if !IsNil(o.ShareCode) {
		toSerialize["share_code"] = o.ShareCode
	}
	return toSerialize, nil
}

type NullableAlipayTradePrecreateResponseModel struct {
	value *AlipayTradePrecreateResponseModel
	isSet bool
}

func (v NullableAlipayTradePrecreateResponseModel) Get() *AlipayTradePrecreateResponseModel {
	return v.value
}

func (v *NullableAlipayTradePrecreateResponseModel) Set(val *AlipayTradePrecreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayTradePrecreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayTradePrecreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayTradePrecreateResponseModel(val *AlipayTradePrecreateResponseModel) *NullableAlipayTradePrecreateResponseModel {
	return &NullableAlipayTradePrecreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayTradePrecreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayTradePrecreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

