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

// checks if the OrderVehicleInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderVehicleInfo{}

// OrderVehicleInfo struct for OrderVehicleInfo
type OrderVehicleInfo struct {
	// 交通工具牌照号，如车牌号等
	LicensePlateNo *string `json:"license_plate_no,omitempty"`
	// 备注信息
	Memo *string `json:"memo,omitempty"`
	// 班次
	ShiftNo *string `json:"shift_no,omitempty"`
}

// NewOrderVehicleInfo instantiates a new OrderVehicleInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderVehicleInfo() *OrderVehicleInfo {
	this := OrderVehicleInfo{}
	return &this
}

// NewOrderVehicleInfoWithDefaults instantiates a new OrderVehicleInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderVehicleInfoWithDefaults() *OrderVehicleInfo {
	this := OrderVehicleInfo{}
	return &this
}

// GetLicensePlateNo returns the LicensePlateNo field value if set, zero value otherwise.
func (o *OrderVehicleInfo) GetLicensePlateNo() string {
	if o == nil || IsNil(o.LicensePlateNo) {
		var ret string
		return ret
	}
	return *o.LicensePlateNo
}

// GetLicensePlateNoOk returns a tuple with the LicensePlateNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderVehicleInfo) GetLicensePlateNoOk() (*string, bool) {
	if o == nil || IsNil(o.LicensePlateNo) {
		return nil, false
	}
	return o.LicensePlateNo, true
}

// HasLicensePlateNo returns a boolean if a field has been set.
func (o *OrderVehicleInfo) HasLicensePlateNo() bool {
	if o != nil && !IsNil(o.LicensePlateNo) {
		return true
	}

	return false
}

// SetLicensePlateNo gets a reference to the given string and assigns it to the LicensePlateNo field.
func (o *OrderVehicleInfo) SetLicensePlateNo(v string) {
	o.LicensePlateNo = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *OrderVehicleInfo) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderVehicleInfo) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *OrderVehicleInfo) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *OrderVehicleInfo) SetMemo(v string) {
	o.Memo = &v
}

// GetShiftNo returns the ShiftNo field value if set, zero value otherwise.
func (o *OrderVehicleInfo) GetShiftNo() string {
	if o == nil || IsNil(o.ShiftNo) {
		var ret string
		return ret
	}
	return *o.ShiftNo
}

// GetShiftNoOk returns a tuple with the ShiftNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderVehicleInfo) GetShiftNoOk() (*string, bool) {
	if o == nil || IsNil(o.ShiftNo) {
		return nil, false
	}
	return o.ShiftNo, true
}

// HasShiftNo returns a boolean if a field has been set.
func (o *OrderVehicleInfo) HasShiftNo() bool {
	if o != nil && !IsNil(o.ShiftNo) {
		return true
	}

	return false
}

// SetShiftNo gets a reference to the given string and assigns it to the ShiftNo field.
func (o *OrderVehicleInfo) SetShiftNo(v string) {
	o.ShiftNo = &v
}

func (o OrderVehicleInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderVehicleInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LicensePlateNo) {
		toSerialize["license_plate_no"] = o.LicensePlateNo
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.ShiftNo) {
		toSerialize["shift_no"] = o.ShiftNo
	}
	return toSerialize, nil
}

type NullableOrderVehicleInfo struct {
	value *OrderVehicleInfo
	isSet bool
}

func (v NullableOrderVehicleInfo) Get() *OrderVehicleInfo {
	return v.value
}

func (v *NullableOrderVehicleInfo) Set(val *OrderVehicleInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderVehicleInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderVehicleInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderVehicleInfo(val *OrderVehicleInfo) *NullableOrderVehicleInfo {
	return &NullableOrderVehicleInfo{value: val, isSet: true}
}

func (v NullableOrderVehicleInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderVehicleInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

