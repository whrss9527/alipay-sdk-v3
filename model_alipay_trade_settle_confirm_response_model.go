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

// checks if the AlipayTradeSettleConfirmResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayTradeSettleConfirmResponseModel{}

// AlipayTradeSettleConfirmResponseModel struct for AlipayTradeSettleConfirmResponseModel
type AlipayTradeSettleConfirmResponseModel struct {
	// 确认结算请求流水号，开发者自行生成并保证唯一性，作为业务幂等性控制
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// 本次确认结算的实际结算金额，单位为元。
	SettleAmount *string `json:"settle_amount,omitempty"`
	// 支付宝交易号
	TradeNo *string `json:"trade_no,omitempty"`
}

// NewAlipayTradeSettleConfirmResponseModel instantiates a new AlipayTradeSettleConfirmResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayTradeSettleConfirmResponseModel() *AlipayTradeSettleConfirmResponseModel {
	this := AlipayTradeSettleConfirmResponseModel{}
	return &this
}

// NewAlipayTradeSettleConfirmResponseModelWithDefaults instantiates a new AlipayTradeSettleConfirmResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayTradeSettleConfirmResponseModelWithDefaults() *AlipayTradeSettleConfirmResponseModel {
	this := AlipayTradeSettleConfirmResponseModel{}
	return &this
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *AlipayTradeSettleConfirmResponseModel) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeSettleConfirmResponseModel) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *AlipayTradeSettleConfirmResponseModel) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *AlipayTradeSettleConfirmResponseModel) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

// GetSettleAmount returns the SettleAmount field value if set, zero value otherwise.
func (o *AlipayTradeSettleConfirmResponseModel) GetSettleAmount() string {
	if o == nil || IsNil(o.SettleAmount) {
		var ret string
		return ret
	}
	return *o.SettleAmount
}

// GetSettleAmountOk returns a tuple with the SettleAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeSettleConfirmResponseModel) GetSettleAmountOk() (*string, bool) {
	if o == nil || IsNil(o.SettleAmount) {
		return nil, false
	}
	return o.SettleAmount, true
}

// HasSettleAmount returns a boolean if a field has been set.
func (o *AlipayTradeSettleConfirmResponseModel) HasSettleAmount() bool {
	if o != nil && !IsNil(o.SettleAmount) {
		return true
	}

	return false
}

// SetSettleAmount gets a reference to the given string and assigns it to the SettleAmount field.
func (o *AlipayTradeSettleConfirmResponseModel) SetSettleAmount(v string) {
	o.SettleAmount = &v
}

// GetTradeNo returns the TradeNo field value if set, zero value otherwise.
func (o *AlipayTradeSettleConfirmResponseModel) GetTradeNo() string {
	if o == nil || IsNil(o.TradeNo) {
		var ret string
		return ret
	}
	return *o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeSettleConfirmResponseModel) GetTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.TradeNo) {
		return nil, false
	}
	return o.TradeNo, true
}

// HasTradeNo returns a boolean if a field has been set.
func (o *AlipayTradeSettleConfirmResponseModel) HasTradeNo() bool {
	if o != nil && !IsNil(o.TradeNo) {
		return true
	}

	return false
}

// SetTradeNo gets a reference to the given string and assigns it to the TradeNo field.
func (o *AlipayTradeSettleConfirmResponseModel) SetTradeNo(v string) {
	o.TradeNo = &v
}

func (o AlipayTradeSettleConfirmResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayTradeSettleConfirmResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OutRequestNo) {
		toSerialize["out_request_no"] = o.OutRequestNo
	}
	if !IsNil(o.SettleAmount) {
		toSerialize["settle_amount"] = o.SettleAmount
	}
	if !IsNil(o.TradeNo) {
		toSerialize["trade_no"] = o.TradeNo
	}
	return toSerialize, nil
}

type NullableAlipayTradeSettleConfirmResponseModel struct {
	value *AlipayTradeSettleConfirmResponseModel
	isSet bool
}

func (v NullableAlipayTradeSettleConfirmResponseModel) Get() *AlipayTradeSettleConfirmResponseModel {
	return v.value
}

func (v *NullableAlipayTradeSettleConfirmResponseModel) Set(val *AlipayTradeSettleConfirmResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayTradeSettleConfirmResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayTradeSettleConfirmResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayTradeSettleConfirmResponseModel(val *AlipayTradeSettleConfirmResponseModel) *NullableAlipayTradeSettleConfirmResponseModel {
	return &NullableAlipayTradeSettleConfirmResponseModel{value: val, isSet: true}
}

func (v NullableAlipayTradeSettleConfirmResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayTradeSettleConfirmResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
