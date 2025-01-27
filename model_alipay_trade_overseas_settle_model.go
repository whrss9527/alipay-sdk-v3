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

// checks if the AlipayTradeOverseasSettleModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayTradeOverseasSettleModel{}

// AlipayTradeOverseasSettleModel struct for AlipayTradeOverseasSettleModel
type AlipayTradeOverseasSettleModel struct {
	// 跨境结算的人民币金额，单位为元；等于交易金额 (实际为实收金额)，加上净补差金额，减去已退款金额，减去净收费金额，再减去净分账金额，
	Amount       *string               `json:"amount,omitempty"`
	ExtendParams *OverseasExtendParams `json:"extend_params,omitempty"`
	// 跨境结算的外币币种
	ForeignSettleCurrency *string `json:"foreign_settle_currency,omitempty"`
	// 外部请求号，开发者自行生成并保证唯一性，作为业务幂等性控制
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// 支付宝交易号
	TradeNo *string `json:"trade_no,omitempty"`
}

// NewAlipayTradeOverseasSettleModel instantiates a new AlipayTradeOverseasSettleModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayTradeOverseasSettleModel() *AlipayTradeOverseasSettleModel {
	this := AlipayTradeOverseasSettleModel{}
	return &this
}

// NewAlipayTradeOverseasSettleModelWithDefaults instantiates a new AlipayTradeOverseasSettleModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayTradeOverseasSettleModelWithDefaults() *AlipayTradeOverseasSettleModel {
	this := AlipayTradeOverseasSettleModel{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AlipayTradeOverseasSettleModel) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeOverseasSettleModel) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AlipayTradeOverseasSettleModel) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AlipayTradeOverseasSettleModel) SetAmount(v string) {
	o.Amount = &v
}

// GetExtendParams returns the ExtendParams field value if set, zero value otherwise.
func (o *AlipayTradeOverseasSettleModel) GetExtendParams() OverseasExtendParams {
	if o == nil || IsNil(o.ExtendParams) {
		var ret OverseasExtendParams
		return ret
	}
	return *o.ExtendParams
}

// GetExtendParamsOk returns a tuple with the ExtendParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeOverseasSettleModel) GetExtendParamsOk() (*OverseasExtendParams, bool) {
	if o == nil || IsNil(o.ExtendParams) {
		return nil, false
	}
	return o.ExtendParams, true
}

// HasExtendParams returns a boolean if a field has been set.
func (o *AlipayTradeOverseasSettleModel) HasExtendParams() bool {
	if o != nil && !IsNil(o.ExtendParams) {
		return true
	}

	return false
}

// SetExtendParams gets a reference to the given OverseasExtendParams and assigns it to the ExtendParams field.
func (o *AlipayTradeOverseasSettleModel) SetExtendParams(v OverseasExtendParams) {
	o.ExtendParams = &v
}

// GetForeignSettleCurrency returns the ForeignSettleCurrency field value if set, zero value otherwise.
func (o *AlipayTradeOverseasSettleModel) GetForeignSettleCurrency() string {
	if o == nil || IsNil(o.ForeignSettleCurrency) {
		var ret string
		return ret
	}
	return *o.ForeignSettleCurrency
}

// GetForeignSettleCurrencyOk returns a tuple with the ForeignSettleCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeOverseasSettleModel) GetForeignSettleCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.ForeignSettleCurrency) {
		return nil, false
	}
	return o.ForeignSettleCurrency, true
}

// HasForeignSettleCurrency returns a boolean if a field has been set.
func (o *AlipayTradeOverseasSettleModel) HasForeignSettleCurrency() bool {
	if o != nil && !IsNil(o.ForeignSettleCurrency) {
		return true
	}

	return false
}

// SetForeignSettleCurrency gets a reference to the given string and assigns it to the ForeignSettleCurrency field.
func (o *AlipayTradeOverseasSettleModel) SetForeignSettleCurrency(v string) {
	o.ForeignSettleCurrency = &v
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *AlipayTradeOverseasSettleModel) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeOverseasSettleModel) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *AlipayTradeOverseasSettleModel) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *AlipayTradeOverseasSettleModel) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

// GetTradeNo returns the TradeNo field value if set, zero value otherwise.
func (o *AlipayTradeOverseasSettleModel) GetTradeNo() string {
	if o == nil || IsNil(o.TradeNo) {
		var ret string
		return ret
	}
	return *o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeOverseasSettleModel) GetTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.TradeNo) {
		return nil, false
	}
	return o.TradeNo, true
}

// HasTradeNo returns a boolean if a field has been set.
func (o *AlipayTradeOverseasSettleModel) HasTradeNo() bool {
	if o != nil && !IsNil(o.TradeNo) {
		return true
	}

	return false
}

// SetTradeNo gets a reference to the given string and assigns it to the TradeNo field.
func (o *AlipayTradeOverseasSettleModel) SetTradeNo(v string) {
	o.TradeNo = &v
}

func (o AlipayTradeOverseasSettleModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayTradeOverseasSettleModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.ExtendParams) {
		toSerialize["extend_params"] = o.ExtendParams
	}
	if !IsNil(o.ForeignSettleCurrency) {
		toSerialize["foreign_settle_currency"] = o.ForeignSettleCurrency
	}
	if !IsNil(o.OutRequestNo) {
		toSerialize["out_request_no"] = o.OutRequestNo
	}
	if !IsNil(o.TradeNo) {
		toSerialize["trade_no"] = o.TradeNo
	}
	return toSerialize, nil
}

type NullableAlipayTradeOverseasSettleModel struct {
	value *AlipayTradeOverseasSettleModel
	isSet bool
}

func (v NullableAlipayTradeOverseasSettleModel) Get() *AlipayTradeOverseasSettleModel {
	return v.value
}

func (v *NullableAlipayTradeOverseasSettleModel) Set(val *AlipayTradeOverseasSettleModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayTradeOverseasSettleModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayTradeOverseasSettleModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayTradeOverseasSettleModel(val *AlipayTradeOverseasSettleModel) *NullableAlipayTradeOverseasSettleModel {
	return &NullableAlipayTradeOverseasSettleModel{value: val, isSet: true}
}

func (v NullableAlipayTradeOverseasSettleModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayTradeOverseasSettleModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
