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

// checks if the TradeFundBill type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TradeFundBill{}

// TradeFundBill struct for TradeFundBill
type TradeFundBill struct {
	// 该支付工具类型所使用的金额。单位：元。
	Amount *string `json:"amount,omitempty"`
	// 银行卡支付时的银行代码
	BankCode *string `json:"bank_code,omitempty"`
	// 交易使用的资金渠道，详见 <a href=\"https://opendocs.alipay.com/open/common/103259\">支付渠道列表</a>
	FundChannel *string `json:"fund_channel,omitempty"`
	// 渠道所使用的资金类型,目前只在资金渠道(fund_channel)是银行卡渠道(BANKCARD)的情况下才返回该信息
	FundType *string `json:"fund_type,omitempty"`
	// 渠道实际付款金额
	RealAmount *string `json:"real_amount,omitempty"`
}

// NewTradeFundBill instantiates a new TradeFundBill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradeFundBill() *TradeFundBill {
	this := TradeFundBill{}
	return &this
}

// NewTradeFundBillWithDefaults instantiates a new TradeFundBill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradeFundBillWithDefaults() *TradeFundBill {
	this := TradeFundBill{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TradeFundBill) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeFundBill) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TradeFundBill) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *TradeFundBill) SetAmount(v string) {
	o.Amount = &v
}

// GetBankCode returns the BankCode field value if set, zero value otherwise.
func (o *TradeFundBill) GetBankCode() string {
	if o == nil || IsNil(o.BankCode) {
		var ret string
		return ret
	}
	return *o.BankCode
}

// GetBankCodeOk returns a tuple with the BankCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeFundBill) GetBankCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BankCode) {
		return nil, false
	}
	return o.BankCode, true
}

// HasBankCode returns a boolean if a field has been set.
func (o *TradeFundBill) HasBankCode() bool {
	if o != nil && !IsNil(o.BankCode) {
		return true
	}

	return false
}

// SetBankCode gets a reference to the given string and assigns it to the BankCode field.
func (o *TradeFundBill) SetBankCode(v string) {
	o.BankCode = &v
}

// GetFundChannel returns the FundChannel field value if set, zero value otherwise.
func (o *TradeFundBill) GetFundChannel() string {
	if o == nil || IsNil(o.FundChannel) {
		var ret string
		return ret
	}
	return *o.FundChannel
}

// GetFundChannelOk returns a tuple with the FundChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeFundBill) GetFundChannelOk() (*string, bool) {
	if o == nil || IsNil(o.FundChannel) {
		return nil, false
	}
	return o.FundChannel, true
}

// HasFundChannel returns a boolean if a field has been set.
func (o *TradeFundBill) HasFundChannel() bool {
	if o != nil && !IsNil(o.FundChannel) {
		return true
	}

	return false
}

// SetFundChannel gets a reference to the given string and assigns it to the FundChannel field.
func (o *TradeFundBill) SetFundChannel(v string) {
	o.FundChannel = &v
}

// GetFundType returns the FundType field value if set, zero value otherwise.
func (o *TradeFundBill) GetFundType() string {
	if o == nil || IsNil(o.FundType) {
		var ret string
		return ret
	}
	return *o.FundType
}

// GetFundTypeOk returns a tuple with the FundType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeFundBill) GetFundTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FundType) {
		return nil, false
	}
	return o.FundType, true
}

// HasFundType returns a boolean if a field has been set.
func (o *TradeFundBill) HasFundType() bool {
	if o != nil && !IsNil(o.FundType) {
		return true
	}

	return false
}

// SetFundType gets a reference to the given string and assigns it to the FundType field.
func (o *TradeFundBill) SetFundType(v string) {
	o.FundType = &v
}

// GetRealAmount returns the RealAmount field value if set, zero value otherwise.
func (o *TradeFundBill) GetRealAmount() string {
	if o == nil || IsNil(o.RealAmount) {
		var ret string
		return ret
	}
	return *o.RealAmount
}

// GetRealAmountOk returns a tuple with the RealAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeFundBill) GetRealAmountOk() (*string, bool) {
	if o == nil || IsNil(o.RealAmount) {
		return nil, false
	}
	return o.RealAmount, true
}

// HasRealAmount returns a boolean if a field has been set.
func (o *TradeFundBill) HasRealAmount() bool {
	if o != nil && !IsNil(o.RealAmount) {
		return true
	}

	return false
}

// SetRealAmount gets a reference to the given string and assigns it to the RealAmount field.
func (o *TradeFundBill) SetRealAmount(v string) {
	o.RealAmount = &v
}

func (o TradeFundBill) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradeFundBill) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BankCode) {
		toSerialize["bank_code"] = o.BankCode
	}
	if !IsNil(o.FundChannel) {
		toSerialize["fund_channel"] = o.FundChannel
	}
	if !IsNil(o.FundType) {
		toSerialize["fund_type"] = o.FundType
	}
	if !IsNil(o.RealAmount) {
		toSerialize["real_amount"] = o.RealAmount
	}
	return toSerialize, nil
}

type NullableTradeFundBill struct {
	value *TradeFundBill
	isSet bool
}

func (v NullableTradeFundBill) Get() *TradeFundBill {
	return v.value
}

func (v *NullableTradeFundBill) Set(val *TradeFundBill) {
	v.value = val
	v.isSet = true
}

func (v NullableTradeFundBill) IsSet() bool {
	return v.isSet
}

func (v *NullableTradeFundBill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradeFundBill(val *TradeFundBill) *NullableTradeFundBill {
	return &NullableTradeFundBill{value: val, isSet: true}
}

func (v NullableTradeFundBill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradeFundBill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
