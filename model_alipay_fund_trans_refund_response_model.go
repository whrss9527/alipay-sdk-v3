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

// checks if the AlipayFundTransRefundResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundTransRefundResponseModel{}

// AlipayFundTransRefundResponseModel struct for AlipayFundTransRefundResponseModel
type AlipayFundTransRefundResponseModel struct {
	// 发红包时支付宝返回的支付宝订单号order_id。
	OrderId *string `json:"order_id,omitempty"`
	// 标识一次资金退回请求，一笔资金退回失败后重新提交，要采用原来的资金退回单号。总退款金额不能超过用户实际支付金额。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// 本次退款的金额，单位为元，支持两位小数
	RefundAmount *string `json:"refund_amount,omitempty"`
	// 退款资金退回
	RefundDate *string `json:"refund_date,omitempty"`
	// 退款的支付宝系统内部单据id
	RefundOrderId *string `json:"refund_order_id,omitempty"`
	// SUCCESS：退款成功
	Status *string `json:"status,omitempty"`
}

// NewAlipayFundTransRefundResponseModel instantiates a new AlipayFundTransRefundResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundTransRefundResponseModel() *AlipayFundTransRefundResponseModel {
	this := AlipayFundTransRefundResponseModel{}
	return &this
}

// NewAlipayFundTransRefundResponseModelWithDefaults instantiates a new AlipayFundTransRefundResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundTransRefundResponseModelWithDefaults() *AlipayFundTransRefundResponseModel {
	this := AlipayFundTransRefundResponseModel{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *AlipayFundTransRefundResponseModel) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransRefundResponseModel) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *AlipayFundTransRefundResponseModel) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *AlipayFundTransRefundResponseModel) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *AlipayFundTransRefundResponseModel) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransRefundResponseModel) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *AlipayFundTransRefundResponseModel) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *AlipayFundTransRefundResponseModel) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

// GetRefundAmount returns the RefundAmount field value if set, zero value otherwise.
func (o *AlipayFundTransRefundResponseModel) GetRefundAmount() string {
	if o == nil || IsNil(o.RefundAmount) {
		var ret string
		return ret
	}
	return *o.RefundAmount
}

// GetRefundAmountOk returns a tuple with the RefundAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransRefundResponseModel) GetRefundAmountOk() (*string, bool) {
	if o == nil || IsNil(o.RefundAmount) {
		return nil, false
	}
	return o.RefundAmount, true
}

// HasRefundAmount returns a boolean if a field has been set.
func (o *AlipayFundTransRefundResponseModel) HasRefundAmount() bool {
	if o != nil && !IsNil(o.RefundAmount) {
		return true
	}

	return false
}

// SetRefundAmount gets a reference to the given string and assigns it to the RefundAmount field.
func (o *AlipayFundTransRefundResponseModel) SetRefundAmount(v string) {
	o.RefundAmount = &v
}

// GetRefundDate returns the RefundDate field value if set, zero value otherwise.
func (o *AlipayFundTransRefundResponseModel) GetRefundDate() string {
	if o == nil || IsNil(o.RefundDate) {
		var ret string
		return ret
	}
	return *o.RefundDate
}

// GetRefundDateOk returns a tuple with the RefundDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransRefundResponseModel) GetRefundDateOk() (*string, bool) {
	if o == nil || IsNil(o.RefundDate) {
		return nil, false
	}
	return o.RefundDate, true
}

// HasRefundDate returns a boolean if a field has been set.
func (o *AlipayFundTransRefundResponseModel) HasRefundDate() bool {
	if o != nil && !IsNil(o.RefundDate) {
		return true
	}

	return false
}

// SetRefundDate gets a reference to the given string and assigns it to the RefundDate field.
func (o *AlipayFundTransRefundResponseModel) SetRefundDate(v string) {
	o.RefundDate = &v
}

// GetRefundOrderId returns the RefundOrderId field value if set, zero value otherwise.
func (o *AlipayFundTransRefundResponseModel) GetRefundOrderId() string {
	if o == nil || IsNil(o.RefundOrderId) {
		var ret string
		return ret
	}
	return *o.RefundOrderId
}

// GetRefundOrderIdOk returns a tuple with the RefundOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransRefundResponseModel) GetRefundOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefundOrderId) {
		return nil, false
	}
	return o.RefundOrderId, true
}

// HasRefundOrderId returns a boolean if a field has been set.
func (o *AlipayFundTransRefundResponseModel) HasRefundOrderId() bool {
	if o != nil && !IsNil(o.RefundOrderId) {
		return true
	}

	return false
}

// SetRefundOrderId gets a reference to the given string and assigns it to the RefundOrderId field.
func (o *AlipayFundTransRefundResponseModel) SetRefundOrderId(v string) {
	o.RefundOrderId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayFundTransRefundResponseModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransRefundResponseModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayFundTransRefundResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayFundTransRefundResponseModel) SetStatus(v string) {
	o.Status = &v
}

func (o AlipayFundTransRefundResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundTransRefundResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.OutRequestNo) {
		toSerialize["out_request_no"] = o.OutRequestNo
	}
	if !IsNil(o.RefundAmount) {
		toSerialize["refund_amount"] = o.RefundAmount
	}
	if !IsNil(o.RefundDate) {
		toSerialize["refund_date"] = o.RefundDate
	}
	if !IsNil(o.RefundOrderId) {
		toSerialize["refund_order_id"] = o.RefundOrderId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAlipayFundTransRefundResponseModel struct {
	value *AlipayFundTransRefundResponseModel
	isSet bool
}

func (v NullableAlipayFundTransRefundResponseModel) Get() *AlipayFundTransRefundResponseModel {
	return v.value
}

func (v *NullableAlipayFundTransRefundResponseModel) Set(val *AlipayFundTransRefundResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundTransRefundResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundTransRefundResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundTransRefundResponseModel(val *AlipayFundTransRefundResponseModel) *NullableAlipayFundTransRefundResponseModel {
	return &NullableAlipayFundTransRefundResponseModel{value: val, isSet: true}
}

func (v NullableAlipayFundTransRefundResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundTransRefundResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
