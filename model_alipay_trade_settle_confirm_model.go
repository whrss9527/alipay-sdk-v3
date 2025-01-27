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

// checks if the AlipayTradeSettleConfirmModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayTradeSettleConfirmModel{}

// AlipayTradeSettleConfirmModel struct for AlipayTradeSettleConfirmModel
type AlipayTradeSettleConfirmModel struct {
	ExtendParams *SettleConfirmExtendParams `json:"extend_params,omitempty"`
	// 确认结算请求流水号，开发者自行生成并保证唯一性，作为业务幂等性控制
	OutRequestNo *string     `json:"out_request_no,omitempty"`
	SettleInfo   *SettleInfo `json:"settle_info,omitempty"`
	// 支付宝交易号
	TradeNo *string `json:"trade_no,omitempty"`
}

// NewAlipayTradeSettleConfirmModel instantiates a new AlipayTradeSettleConfirmModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayTradeSettleConfirmModel() *AlipayTradeSettleConfirmModel {
	this := AlipayTradeSettleConfirmModel{}
	return &this
}

// NewAlipayTradeSettleConfirmModelWithDefaults instantiates a new AlipayTradeSettleConfirmModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayTradeSettleConfirmModelWithDefaults() *AlipayTradeSettleConfirmModel {
	this := AlipayTradeSettleConfirmModel{}
	return &this
}

// GetExtendParams returns the ExtendParams field value if set, zero value otherwise.
func (o *AlipayTradeSettleConfirmModel) GetExtendParams() SettleConfirmExtendParams {
	if o == nil || IsNil(o.ExtendParams) {
		var ret SettleConfirmExtendParams
		return ret
	}
	return *o.ExtendParams
}

// GetExtendParamsOk returns a tuple with the ExtendParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeSettleConfirmModel) GetExtendParamsOk() (*SettleConfirmExtendParams, bool) {
	if o == nil || IsNil(o.ExtendParams) {
		return nil, false
	}
	return o.ExtendParams, true
}

// HasExtendParams returns a boolean if a field has been set.
func (o *AlipayTradeSettleConfirmModel) HasExtendParams() bool {
	if o != nil && !IsNil(o.ExtendParams) {
		return true
	}

	return false
}

// SetExtendParams gets a reference to the given SettleConfirmExtendParams and assigns it to the ExtendParams field.
func (o *AlipayTradeSettleConfirmModel) SetExtendParams(v SettleConfirmExtendParams) {
	o.ExtendParams = &v
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *AlipayTradeSettleConfirmModel) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeSettleConfirmModel) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *AlipayTradeSettleConfirmModel) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *AlipayTradeSettleConfirmModel) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

// GetSettleInfo returns the SettleInfo field value if set, zero value otherwise.
func (o *AlipayTradeSettleConfirmModel) GetSettleInfo() SettleInfo {
	if o == nil || IsNil(o.SettleInfo) {
		var ret SettleInfo
		return ret
	}
	return *o.SettleInfo
}

// GetSettleInfoOk returns a tuple with the SettleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeSettleConfirmModel) GetSettleInfoOk() (*SettleInfo, bool) {
	if o == nil || IsNil(o.SettleInfo) {
		return nil, false
	}
	return o.SettleInfo, true
}

// HasSettleInfo returns a boolean if a field has been set.
func (o *AlipayTradeSettleConfirmModel) HasSettleInfo() bool {
	if o != nil && !IsNil(o.SettleInfo) {
		return true
	}

	return false
}

// SetSettleInfo gets a reference to the given SettleInfo and assigns it to the SettleInfo field.
func (o *AlipayTradeSettleConfirmModel) SetSettleInfo(v SettleInfo) {
	o.SettleInfo = &v
}

// GetTradeNo returns the TradeNo field value if set, zero value otherwise.
func (o *AlipayTradeSettleConfirmModel) GetTradeNo() string {
	if o == nil || IsNil(o.TradeNo) {
		var ret string
		return ret
	}
	return *o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeSettleConfirmModel) GetTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.TradeNo) {
		return nil, false
	}
	return o.TradeNo, true
}

// HasTradeNo returns a boolean if a field has been set.
func (o *AlipayTradeSettleConfirmModel) HasTradeNo() bool {
	if o != nil && !IsNil(o.TradeNo) {
		return true
	}

	return false
}

// SetTradeNo gets a reference to the given string and assigns it to the TradeNo field.
func (o *AlipayTradeSettleConfirmModel) SetTradeNo(v string) {
	o.TradeNo = &v
}

func (o AlipayTradeSettleConfirmModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayTradeSettleConfirmModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtendParams) {
		toSerialize["extend_params"] = o.ExtendParams
	}
	if !IsNil(o.OutRequestNo) {
		toSerialize["out_request_no"] = o.OutRequestNo
	}
	if !IsNil(o.SettleInfo) {
		toSerialize["settle_info"] = o.SettleInfo
	}
	if !IsNil(o.TradeNo) {
		toSerialize["trade_no"] = o.TradeNo
	}
	return toSerialize, nil
}

type NullableAlipayTradeSettleConfirmModel struct {
	value *AlipayTradeSettleConfirmModel
	isSet bool
}

func (v NullableAlipayTradeSettleConfirmModel) Get() *AlipayTradeSettleConfirmModel {
	return v.value
}

func (v *NullableAlipayTradeSettleConfirmModel) Set(val *AlipayTradeSettleConfirmModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayTradeSettleConfirmModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayTradeSettleConfirmModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayTradeSettleConfirmModel(val *AlipayTradeSettleConfirmModel) *NullableAlipayTradeSettleConfirmModel {
	return &NullableAlipayTradeSettleConfirmModel{value: val, isSet: true}
}

func (v NullableAlipayTradeSettleConfirmModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayTradeSettleConfirmModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
