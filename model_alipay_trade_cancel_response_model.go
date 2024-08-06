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

// checks if the AlipayTradeCancelResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayTradeCancelResponseModel{}

// AlipayTradeCancelResponseModel struct for AlipayTradeCancelResponseModel
type AlipayTradeCancelResponseModel struct {
	// 本次撤销触发的交易动作,接口调用成功且交易存在时返回。可能的返回值： close：交易未支付，触发关闭交易动作，无退款； refund：交易已支付，触发交易退款动作； 未返回：未查询到交易，或接口调用失败；
	Action *string `json:"action,omitempty"`
	// 当撤销产生了退款时，返回退款时间；  只在银行间联交易场景下返回该信息；
	GmtRefundPay *string `json:"gmt_refund_pay,omitempty"`
	// 商户订单号
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// 当撤销产生了退款时，返回的退款清算编号，用于清算对账使用；  只在银行间联交易场景下返回该信息；
	RefundSettlementId *string `json:"refund_settlement_id,omitempty"`
	// 是否需要重试
	RetryFlag *string `json:"retry_flag,omitempty"`
	// 支付宝交易号; 当发生交易关闭或交易退款时返回；
	TradeNo *string `json:"trade_no,omitempty"`
}

// NewAlipayTradeCancelResponseModel instantiates a new AlipayTradeCancelResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayTradeCancelResponseModel() *AlipayTradeCancelResponseModel {
	this := AlipayTradeCancelResponseModel{}
	return &this
}

// NewAlipayTradeCancelResponseModelWithDefaults instantiates a new AlipayTradeCancelResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayTradeCancelResponseModelWithDefaults() *AlipayTradeCancelResponseModel {
	this := AlipayTradeCancelResponseModel{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AlipayTradeCancelResponseModel) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeCancelResponseModel) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AlipayTradeCancelResponseModel) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *AlipayTradeCancelResponseModel) SetAction(v string) {
	o.Action = &v
}

// GetGmtRefundPay returns the GmtRefundPay field value if set, zero value otherwise.
func (o *AlipayTradeCancelResponseModel) GetGmtRefundPay() string {
	if o == nil || IsNil(o.GmtRefundPay) {
		var ret string
		return ret
	}
	return *o.GmtRefundPay
}

// GetGmtRefundPayOk returns a tuple with the GmtRefundPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeCancelResponseModel) GetGmtRefundPayOk() (*string, bool) {
	if o == nil || IsNil(o.GmtRefundPay) {
		return nil, false
	}
	return o.GmtRefundPay, true
}

// HasGmtRefundPay returns a boolean if a field has been set.
func (o *AlipayTradeCancelResponseModel) HasGmtRefundPay() bool {
	if o != nil && !IsNil(o.GmtRefundPay) {
		return true
	}

	return false
}

// SetGmtRefundPay gets a reference to the given string and assigns it to the GmtRefundPay field.
func (o *AlipayTradeCancelResponseModel) SetGmtRefundPay(v string) {
	o.GmtRefundPay = &v
}

// GetOutTradeNo returns the OutTradeNo field value if set, zero value otherwise.
func (o *AlipayTradeCancelResponseModel) GetOutTradeNo() string {
	if o == nil || IsNil(o.OutTradeNo) {
		var ret string
		return ret
	}
	return *o.OutTradeNo
}

// GetOutTradeNoOk returns a tuple with the OutTradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeCancelResponseModel) GetOutTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutTradeNo) {
		return nil, false
	}
	return o.OutTradeNo, true
}

// HasOutTradeNo returns a boolean if a field has been set.
func (o *AlipayTradeCancelResponseModel) HasOutTradeNo() bool {
	if o != nil && !IsNil(o.OutTradeNo) {
		return true
	}

	return false
}

// SetOutTradeNo gets a reference to the given string and assigns it to the OutTradeNo field.
func (o *AlipayTradeCancelResponseModel) SetOutTradeNo(v string) {
	o.OutTradeNo = &v
}

// GetRefundSettlementId returns the RefundSettlementId field value if set, zero value otherwise.
func (o *AlipayTradeCancelResponseModel) GetRefundSettlementId() string {
	if o == nil || IsNil(o.RefundSettlementId) {
		var ret string
		return ret
	}
	return *o.RefundSettlementId
}

// GetRefundSettlementIdOk returns a tuple with the RefundSettlementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeCancelResponseModel) GetRefundSettlementIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefundSettlementId) {
		return nil, false
	}
	return o.RefundSettlementId, true
}

// HasRefundSettlementId returns a boolean if a field has been set.
func (o *AlipayTradeCancelResponseModel) HasRefundSettlementId() bool {
	if o != nil && !IsNil(o.RefundSettlementId) {
		return true
	}

	return false
}

// SetRefundSettlementId gets a reference to the given string and assigns it to the RefundSettlementId field.
func (o *AlipayTradeCancelResponseModel) SetRefundSettlementId(v string) {
	o.RefundSettlementId = &v
}

// GetRetryFlag returns the RetryFlag field value if set, zero value otherwise.
func (o *AlipayTradeCancelResponseModel) GetRetryFlag() string {
	if o == nil || IsNil(o.RetryFlag) {
		var ret string
		return ret
	}
	return *o.RetryFlag
}

// GetRetryFlagOk returns a tuple with the RetryFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeCancelResponseModel) GetRetryFlagOk() (*string, bool) {
	if o == nil || IsNil(o.RetryFlag) {
		return nil, false
	}
	return o.RetryFlag, true
}

// HasRetryFlag returns a boolean if a field has been set.
func (o *AlipayTradeCancelResponseModel) HasRetryFlag() bool {
	if o != nil && !IsNil(o.RetryFlag) {
		return true
	}

	return false
}

// SetRetryFlag gets a reference to the given string and assigns it to the RetryFlag field.
func (o *AlipayTradeCancelResponseModel) SetRetryFlag(v string) {
	o.RetryFlag = &v
}

// GetTradeNo returns the TradeNo field value if set, zero value otherwise.
func (o *AlipayTradeCancelResponseModel) GetTradeNo() string {
	if o == nil || IsNil(o.TradeNo) {
		var ret string
		return ret
	}
	return *o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeCancelResponseModel) GetTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.TradeNo) {
		return nil, false
	}
	return o.TradeNo, true
}

// HasTradeNo returns a boolean if a field has been set.
func (o *AlipayTradeCancelResponseModel) HasTradeNo() bool {
	if o != nil && !IsNil(o.TradeNo) {
		return true
	}

	return false
}

// SetTradeNo gets a reference to the given string and assigns it to the TradeNo field.
func (o *AlipayTradeCancelResponseModel) SetTradeNo(v string) {
	o.TradeNo = &v
}

func (o AlipayTradeCancelResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayTradeCancelResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.GmtRefundPay) {
		toSerialize["gmt_refund_pay"] = o.GmtRefundPay
	}
	if !IsNil(o.OutTradeNo) {
		toSerialize["out_trade_no"] = o.OutTradeNo
	}
	if !IsNil(o.RefundSettlementId) {
		toSerialize["refund_settlement_id"] = o.RefundSettlementId
	}
	if !IsNil(o.RetryFlag) {
		toSerialize["retry_flag"] = o.RetryFlag
	}
	if !IsNil(o.TradeNo) {
		toSerialize["trade_no"] = o.TradeNo
	}
	return toSerialize, nil
}

type NullableAlipayTradeCancelResponseModel struct {
	value *AlipayTradeCancelResponseModel
	isSet bool
}

func (v NullableAlipayTradeCancelResponseModel) Get() *AlipayTradeCancelResponseModel {
	return v.value
}

func (v *NullableAlipayTradeCancelResponseModel) Set(val *AlipayTradeCancelResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayTradeCancelResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayTradeCancelResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayTradeCancelResponseModel(val *AlipayTradeCancelResponseModel) *NullableAlipayTradeCancelResponseModel {
	return &NullableAlipayTradeCancelResponseModel{value: val, isSet: true}
}

func (v NullableAlipayTradeCancelResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayTradeCancelResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

