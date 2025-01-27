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

// checks if the VoucherPackageSalesInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherPackageSalesInfo{}

// VoucherPackageSalesInfo struct for VoucherPackageSalesInfo
type VoucherPackageSalesInfo struct {
	// 券包售卖预算，单位是份数
	Budget *int32 `json:"budget,omitempty"`
	// 券包购买支付渠道 pcredit：花呗 creditCard：信用卡 credit_group：花呗与信用卡 为空则不限渠道
	PayChannel *string `json:"pay_channel,omitempty"`
	// 券包购买链接
	PurchaseUrl *string `json:"purchase_url,omitempty"`
	// 券包售卖期限内最大购买次数
	SaleCountLimitInPeriod *int32 `json:"sale_count_limit_in_period,omitempty"`
	// 券包购买期限类型 NO：不限制  ALL：售卖时间内  DAY：天  WEEK：周  MONTH：月
	SalePeriodLimit *string `json:"sale_period_limit,omitempty"`
	// 券包售卖价格，单位是元
	SalePrice *string `json:"sale_price,omitempty"`
}

// NewVoucherPackageSalesInfo instantiates a new VoucherPackageSalesInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherPackageSalesInfo() *VoucherPackageSalesInfo {
	this := VoucherPackageSalesInfo{}
	return &this
}

// NewVoucherPackageSalesInfoWithDefaults instantiates a new VoucherPackageSalesInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherPackageSalesInfoWithDefaults() *VoucherPackageSalesInfo {
	this := VoucherPackageSalesInfo{}
	return &this
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *VoucherPackageSalesInfo) GetBudget() int32 {
	if o == nil || IsNil(o.Budget) {
		var ret int32
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherPackageSalesInfo) GetBudgetOk() (*int32, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *VoucherPackageSalesInfo) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given int32 and assigns it to the Budget field.
func (o *VoucherPackageSalesInfo) SetBudget(v int32) {
	o.Budget = &v
}

// GetPayChannel returns the PayChannel field value if set, zero value otherwise.
func (o *VoucherPackageSalesInfo) GetPayChannel() string {
	if o == nil || IsNil(o.PayChannel) {
		var ret string
		return ret
	}
	return *o.PayChannel
}

// GetPayChannelOk returns a tuple with the PayChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherPackageSalesInfo) GetPayChannelOk() (*string, bool) {
	if o == nil || IsNil(o.PayChannel) {
		return nil, false
	}
	return o.PayChannel, true
}

// HasPayChannel returns a boolean if a field has been set.
func (o *VoucherPackageSalesInfo) HasPayChannel() bool {
	if o != nil && !IsNil(o.PayChannel) {
		return true
	}

	return false
}

// SetPayChannel gets a reference to the given string and assigns it to the PayChannel field.
func (o *VoucherPackageSalesInfo) SetPayChannel(v string) {
	o.PayChannel = &v
}

// GetPurchaseUrl returns the PurchaseUrl field value if set, zero value otherwise.
func (o *VoucherPackageSalesInfo) GetPurchaseUrl() string {
	if o == nil || IsNil(o.PurchaseUrl) {
		var ret string
		return ret
	}
	return *o.PurchaseUrl
}

// GetPurchaseUrlOk returns a tuple with the PurchaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherPackageSalesInfo) GetPurchaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseUrl) {
		return nil, false
	}
	return o.PurchaseUrl, true
}

// HasPurchaseUrl returns a boolean if a field has been set.
func (o *VoucherPackageSalesInfo) HasPurchaseUrl() bool {
	if o != nil && !IsNil(o.PurchaseUrl) {
		return true
	}

	return false
}

// SetPurchaseUrl gets a reference to the given string and assigns it to the PurchaseUrl field.
func (o *VoucherPackageSalesInfo) SetPurchaseUrl(v string) {
	o.PurchaseUrl = &v
}

// GetSaleCountLimitInPeriod returns the SaleCountLimitInPeriod field value if set, zero value otherwise.
func (o *VoucherPackageSalesInfo) GetSaleCountLimitInPeriod() int32 {
	if o == nil || IsNil(o.SaleCountLimitInPeriod) {
		var ret int32
		return ret
	}
	return *o.SaleCountLimitInPeriod
}

// GetSaleCountLimitInPeriodOk returns a tuple with the SaleCountLimitInPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherPackageSalesInfo) GetSaleCountLimitInPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.SaleCountLimitInPeriod) {
		return nil, false
	}
	return o.SaleCountLimitInPeriod, true
}

// HasSaleCountLimitInPeriod returns a boolean if a field has been set.
func (o *VoucherPackageSalesInfo) HasSaleCountLimitInPeriod() bool {
	if o != nil && !IsNil(o.SaleCountLimitInPeriod) {
		return true
	}

	return false
}

// SetSaleCountLimitInPeriod gets a reference to the given int32 and assigns it to the SaleCountLimitInPeriod field.
func (o *VoucherPackageSalesInfo) SetSaleCountLimitInPeriod(v int32) {
	o.SaleCountLimitInPeriod = &v
}

// GetSalePeriodLimit returns the SalePeriodLimit field value if set, zero value otherwise.
func (o *VoucherPackageSalesInfo) GetSalePeriodLimit() string {
	if o == nil || IsNil(o.SalePeriodLimit) {
		var ret string
		return ret
	}
	return *o.SalePeriodLimit
}

// GetSalePeriodLimitOk returns a tuple with the SalePeriodLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherPackageSalesInfo) GetSalePeriodLimitOk() (*string, bool) {
	if o == nil || IsNil(o.SalePeriodLimit) {
		return nil, false
	}
	return o.SalePeriodLimit, true
}

// HasSalePeriodLimit returns a boolean if a field has been set.
func (o *VoucherPackageSalesInfo) HasSalePeriodLimit() bool {
	if o != nil && !IsNil(o.SalePeriodLimit) {
		return true
	}

	return false
}

// SetSalePeriodLimit gets a reference to the given string and assigns it to the SalePeriodLimit field.
func (o *VoucherPackageSalesInfo) SetSalePeriodLimit(v string) {
	o.SalePeriodLimit = &v
}

// GetSalePrice returns the SalePrice field value if set, zero value otherwise.
func (o *VoucherPackageSalesInfo) GetSalePrice() string {
	if o == nil || IsNil(o.SalePrice) {
		var ret string
		return ret
	}
	return *o.SalePrice
}

// GetSalePriceOk returns a tuple with the SalePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherPackageSalesInfo) GetSalePriceOk() (*string, bool) {
	if o == nil || IsNil(o.SalePrice) {
		return nil, false
	}
	return o.SalePrice, true
}

// HasSalePrice returns a boolean if a field has been set.
func (o *VoucherPackageSalesInfo) HasSalePrice() bool {
	if o != nil && !IsNil(o.SalePrice) {
		return true
	}

	return false
}

// SetSalePrice gets a reference to the given string and assigns it to the SalePrice field.
func (o *VoucherPackageSalesInfo) SetSalePrice(v string) {
	o.SalePrice = &v
}

func (o VoucherPackageSalesInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherPackageSalesInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	if !IsNil(o.PayChannel) {
		toSerialize["pay_channel"] = o.PayChannel
	}
	if !IsNil(o.PurchaseUrl) {
		toSerialize["purchase_url"] = o.PurchaseUrl
	}
	if !IsNil(o.SaleCountLimitInPeriod) {
		toSerialize["sale_count_limit_in_period"] = o.SaleCountLimitInPeriod
	}
	if !IsNil(o.SalePeriodLimit) {
		toSerialize["sale_period_limit"] = o.SalePeriodLimit
	}
	if !IsNil(o.SalePrice) {
		toSerialize["sale_price"] = o.SalePrice
	}
	return toSerialize, nil
}

type NullableVoucherPackageSalesInfo struct {
	value *VoucherPackageSalesInfo
	isSet bool
}

func (v NullableVoucherPackageSalesInfo) Get() *VoucherPackageSalesInfo {
	return v.value
}

func (v *NullableVoucherPackageSalesInfo) Set(val *VoucherPackageSalesInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherPackageSalesInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherPackageSalesInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherPackageSalesInfo(val *VoucherPackageSalesInfo) *NullableVoucherPackageSalesInfo {
	return &NullableVoucherPackageSalesInfo{value: val, isSet: true}
}

func (v NullableVoucherPackageSalesInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherPackageSalesInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
