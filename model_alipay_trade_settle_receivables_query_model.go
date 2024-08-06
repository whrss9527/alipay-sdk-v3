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

// checks if the AlipayTradeSettleReceivablesQueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayTradeSettleReceivablesQueryModel{}

// AlipayTradeSettleReceivablesQueryModel struct for AlipayTradeSettleReceivablesQueryModel
type AlipayTradeSettleReceivablesQueryModel struct {
	// 收单产品码，商家和支付宝签约的产品码
	BizProduct *string `json:"biz_product,omitempty"`
	// 扩展参数
	ExtendParams *string `json:"extend_params,omitempty"`
	MerchantInfo *SettleEntity `json:"merchant_info,omitempty"`
	// 外部请求号，32个字符以内，可包含字母、数字、下划线。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// 当trade_no不为空时，该字段不生效。 查询历史日期，格式为 yyyyMMdd ，取值范围为昨日起至往前30日内； 不传入时，查询实时待结算余额返回； 传入过去某一天日期，查询对应日期的日终待结算余额返回（注意：日常场景下，昨日日终待结算余额只可在当天 02:00 后查询，在当天 02:00 前查询返回查询错误；大促场景下昨日日终可查时间会适当延后）； 传入过去某一天非近30天内，返回参数错误。
	QueryHisDate *string `json:"query_his_date,omitempty"`
	// 支付宝交易号，当该笔交易为直付通账期模式，查询该笔交易待确认结算金额时必传
	TradeNo *string `json:"trade_no,omitempty"`
}

// NewAlipayTradeSettleReceivablesQueryModel instantiates a new AlipayTradeSettleReceivablesQueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayTradeSettleReceivablesQueryModel() *AlipayTradeSettleReceivablesQueryModel {
	this := AlipayTradeSettleReceivablesQueryModel{}
	return &this
}

// NewAlipayTradeSettleReceivablesQueryModelWithDefaults instantiates a new AlipayTradeSettleReceivablesQueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayTradeSettleReceivablesQueryModelWithDefaults() *AlipayTradeSettleReceivablesQueryModel {
	this := AlipayTradeSettleReceivablesQueryModel{}
	return &this
}

// GetBizProduct returns the BizProduct field value if set, zero value otherwise.
func (o *AlipayTradeSettleReceivablesQueryModel) GetBizProduct() string {
	if o == nil || IsNil(o.BizProduct) {
		var ret string
		return ret
	}
	return *o.BizProduct
}

// GetBizProductOk returns a tuple with the BizProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeSettleReceivablesQueryModel) GetBizProductOk() (*string, bool) {
	if o == nil || IsNil(o.BizProduct) {
		return nil, false
	}
	return o.BizProduct, true
}

// HasBizProduct returns a boolean if a field has been set.
func (o *AlipayTradeSettleReceivablesQueryModel) HasBizProduct() bool {
	if o != nil && !IsNil(o.BizProduct) {
		return true
	}

	return false
}

// SetBizProduct gets a reference to the given string and assigns it to the BizProduct field.
func (o *AlipayTradeSettleReceivablesQueryModel) SetBizProduct(v string) {
	o.BizProduct = &v
}

// GetExtendParams returns the ExtendParams field value if set, zero value otherwise.
func (o *AlipayTradeSettleReceivablesQueryModel) GetExtendParams() string {
	if o == nil || IsNil(o.ExtendParams) {
		var ret string
		return ret
	}
	return *o.ExtendParams
}

// GetExtendParamsOk returns a tuple with the ExtendParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeSettleReceivablesQueryModel) GetExtendParamsOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendParams) {
		return nil, false
	}
	return o.ExtendParams, true
}

// HasExtendParams returns a boolean if a field has been set.
func (o *AlipayTradeSettleReceivablesQueryModel) HasExtendParams() bool {
	if o != nil && !IsNil(o.ExtendParams) {
		return true
	}

	return false
}

// SetExtendParams gets a reference to the given string and assigns it to the ExtendParams field.
func (o *AlipayTradeSettleReceivablesQueryModel) SetExtendParams(v string) {
	o.ExtendParams = &v
}

// GetMerchantInfo returns the MerchantInfo field value if set, zero value otherwise.
func (o *AlipayTradeSettleReceivablesQueryModel) GetMerchantInfo() SettleEntity {
	if o == nil || IsNil(o.MerchantInfo) {
		var ret SettleEntity
		return ret
	}
	return *o.MerchantInfo
}

// GetMerchantInfoOk returns a tuple with the MerchantInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeSettleReceivablesQueryModel) GetMerchantInfoOk() (*SettleEntity, bool) {
	if o == nil || IsNil(o.MerchantInfo) {
		return nil, false
	}
	return o.MerchantInfo, true
}

// HasMerchantInfo returns a boolean if a field has been set.
func (o *AlipayTradeSettleReceivablesQueryModel) HasMerchantInfo() bool {
	if o != nil && !IsNil(o.MerchantInfo) {
		return true
	}

	return false
}

// SetMerchantInfo gets a reference to the given SettleEntity and assigns it to the MerchantInfo field.
func (o *AlipayTradeSettleReceivablesQueryModel) SetMerchantInfo(v SettleEntity) {
	o.MerchantInfo = &v
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *AlipayTradeSettleReceivablesQueryModel) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeSettleReceivablesQueryModel) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *AlipayTradeSettleReceivablesQueryModel) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *AlipayTradeSettleReceivablesQueryModel) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

// GetQueryHisDate returns the QueryHisDate field value if set, zero value otherwise.
func (o *AlipayTradeSettleReceivablesQueryModel) GetQueryHisDate() string {
	if o == nil || IsNil(o.QueryHisDate) {
		var ret string
		return ret
	}
	return *o.QueryHisDate
}

// GetQueryHisDateOk returns a tuple with the QueryHisDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeSettleReceivablesQueryModel) GetQueryHisDateOk() (*string, bool) {
	if o == nil || IsNil(o.QueryHisDate) {
		return nil, false
	}
	return o.QueryHisDate, true
}

// HasQueryHisDate returns a boolean if a field has been set.
func (o *AlipayTradeSettleReceivablesQueryModel) HasQueryHisDate() bool {
	if o != nil && !IsNil(o.QueryHisDate) {
		return true
	}

	return false
}

// SetQueryHisDate gets a reference to the given string and assigns it to the QueryHisDate field.
func (o *AlipayTradeSettleReceivablesQueryModel) SetQueryHisDate(v string) {
	o.QueryHisDate = &v
}

// GetTradeNo returns the TradeNo field value if set, zero value otherwise.
func (o *AlipayTradeSettleReceivablesQueryModel) GetTradeNo() string {
	if o == nil || IsNil(o.TradeNo) {
		var ret string
		return ret
	}
	return *o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeSettleReceivablesQueryModel) GetTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.TradeNo) {
		return nil, false
	}
	return o.TradeNo, true
}

// HasTradeNo returns a boolean if a field has been set.
func (o *AlipayTradeSettleReceivablesQueryModel) HasTradeNo() bool {
	if o != nil && !IsNil(o.TradeNo) {
		return true
	}

	return false
}

// SetTradeNo gets a reference to the given string and assigns it to the TradeNo field.
func (o *AlipayTradeSettleReceivablesQueryModel) SetTradeNo(v string) {
	o.TradeNo = &v
}

func (o AlipayTradeSettleReceivablesQueryModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayTradeSettleReceivablesQueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizProduct) {
		toSerialize["biz_product"] = o.BizProduct
	}
	if !IsNil(o.ExtendParams) {
		toSerialize["extend_params"] = o.ExtendParams
	}
	if !IsNil(o.MerchantInfo) {
		toSerialize["merchant_info"] = o.MerchantInfo
	}
	if !IsNil(o.OutRequestNo) {
		toSerialize["out_request_no"] = o.OutRequestNo
	}
	if !IsNil(o.QueryHisDate) {
		toSerialize["query_his_date"] = o.QueryHisDate
	}
	if !IsNil(o.TradeNo) {
		toSerialize["trade_no"] = o.TradeNo
	}
	return toSerialize, nil
}

type NullableAlipayTradeSettleReceivablesQueryModel struct {
	value *AlipayTradeSettleReceivablesQueryModel
	isSet bool
}

func (v NullableAlipayTradeSettleReceivablesQueryModel) Get() *AlipayTradeSettleReceivablesQueryModel {
	return v.value
}

func (v *NullableAlipayTradeSettleReceivablesQueryModel) Set(val *AlipayTradeSettleReceivablesQueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayTradeSettleReceivablesQueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayTradeSettleReceivablesQueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayTradeSettleReceivablesQueryModel(val *AlipayTradeSettleReceivablesQueryModel) *NullableAlipayTradeSettleReceivablesQueryModel {
	return &NullableAlipayTradeSettleReceivablesQueryModel{value: val, isSet: true}
}

func (v NullableAlipayTradeSettleReceivablesQueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayTradeSettleReceivablesQueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

