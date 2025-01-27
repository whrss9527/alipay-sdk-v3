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

// checks if the EcOrderInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EcOrderInfo{}

// EcOrderInfo struct for EcOrderInfo
type EcOrderInfo struct {
	OrderInfo *EcOrderItem `json:"order_info,omitempty"`
	// 子订单详情列表
	SubOrderList []EcOrderItem `json:"sub_order_list,omitempty"`
}

// NewEcOrderInfo instantiates a new EcOrderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEcOrderInfo() *EcOrderInfo {
	this := EcOrderInfo{}
	return &this
}

// NewEcOrderInfoWithDefaults instantiates a new EcOrderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEcOrderInfoWithDefaults() *EcOrderInfo {
	this := EcOrderInfo{}
	return &this
}

// GetOrderInfo returns the OrderInfo field value if set, zero value otherwise.
func (o *EcOrderInfo) GetOrderInfo() EcOrderItem {
	if o == nil || IsNil(o.OrderInfo) {
		var ret EcOrderItem
		return ret
	}
	return *o.OrderInfo
}

// GetOrderInfoOk returns a tuple with the OrderInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcOrderInfo) GetOrderInfoOk() (*EcOrderItem, bool) {
	if o == nil || IsNil(o.OrderInfo) {
		return nil, false
	}
	return o.OrderInfo, true
}

// HasOrderInfo returns a boolean if a field has been set.
func (o *EcOrderInfo) HasOrderInfo() bool {
	if o != nil && !IsNil(o.OrderInfo) {
		return true
	}

	return false
}

// SetOrderInfo gets a reference to the given EcOrderItem and assigns it to the OrderInfo field.
func (o *EcOrderInfo) SetOrderInfo(v EcOrderItem) {
	o.OrderInfo = &v
}

// GetSubOrderList returns the SubOrderList field value if set, zero value otherwise.
func (o *EcOrderInfo) GetSubOrderList() []EcOrderItem {
	if o == nil || IsNil(o.SubOrderList) {
		var ret []EcOrderItem
		return ret
	}
	return o.SubOrderList
}

// GetSubOrderListOk returns a tuple with the SubOrderList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcOrderInfo) GetSubOrderListOk() ([]EcOrderItem, bool) {
	if o == nil || IsNil(o.SubOrderList) {
		return nil, false
	}
	return o.SubOrderList, true
}

// HasSubOrderList returns a boolean if a field has been set.
func (o *EcOrderInfo) HasSubOrderList() bool {
	if o != nil && !IsNil(o.SubOrderList) {
		return true
	}

	return false
}

// SetSubOrderList gets a reference to the given []EcOrderItem and assigns it to the SubOrderList field.
func (o *EcOrderInfo) SetSubOrderList(v []EcOrderItem) {
	o.SubOrderList = v
}

func (o EcOrderInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EcOrderInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderInfo) {
		toSerialize["order_info"] = o.OrderInfo
	}
	if !IsNil(o.SubOrderList) {
		toSerialize["sub_order_list"] = o.SubOrderList
	}
	return toSerialize, nil
}

type NullableEcOrderInfo struct {
	value *EcOrderInfo
	isSet bool
}

func (v NullableEcOrderInfo) Get() *EcOrderInfo {
	return v.value
}

func (v *NullableEcOrderInfo) Set(val *EcOrderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEcOrderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEcOrderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcOrderInfo(val *EcOrderInfo) *NullableEcOrderInfo {
	return &NullableEcOrderInfo{value: val, isSet: true}
}

func (v NullableEcOrderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcOrderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
