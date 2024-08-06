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

// checks if the ChargeItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeItems{}

// ChargeItems struct for ChargeItems
type ChargeItems struct {
	// 缴费项是否必选  如果缴费项是多选模式，此参数生效。 “Y”表示必填，“N”或空表示非必填。
	ItemMandatory *string `json:"item_mandatory,omitempty"`
	// 缴费项最大可选数  如果缴费项是多选模式，此参数生效，范围是1-9，如果为空，则最大项默认为9
	ItemMaximum *int32 `json:"item_maximum,omitempty"`
	// 缴费项名称
	ItemName *string `json:"item_name,omitempty"`
	// 缴费项金额
	ItemPrice *string `json:"item_price,omitempty"`
	// 缴费项序号，如果缴费项是多选模式，此项为必填，建议从1开始的连续数字，  用户支付成功后，通过passback_params参数带回已选择的缴费项。例如:orderNo=uoo234234&isvOrderNo=24werwe&items=1-2|2-1|3-5  1-2|2-1|3-5 表示：缴费项序列号-缴费项数|缴费项序列号-缴费项数
	ItemSerialNumber *int32 `json:"item_serial_number,omitempty"`
}

// NewChargeItems instantiates a new ChargeItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeItems() *ChargeItems {
	this := ChargeItems{}
	return &this
}

// NewChargeItemsWithDefaults instantiates a new ChargeItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeItemsWithDefaults() *ChargeItems {
	this := ChargeItems{}
	return &this
}

// GetItemMandatory returns the ItemMandatory field value if set, zero value otherwise.
func (o *ChargeItems) GetItemMandatory() string {
	if o == nil || IsNil(o.ItemMandatory) {
		var ret string
		return ret
	}
	return *o.ItemMandatory
}

// GetItemMandatoryOk returns a tuple with the ItemMandatory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeItems) GetItemMandatoryOk() (*string, bool) {
	if o == nil || IsNil(o.ItemMandatory) {
		return nil, false
	}
	return o.ItemMandatory, true
}

// HasItemMandatory returns a boolean if a field has been set.
func (o *ChargeItems) HasItemMandatory() bool {
	if o != nil && !IsNil(o.ItemMandatory) {
		return true
	}

	return false
}

// SetItemMandatory gets a reference to the given string and assigns it to the ItemMandatory field.
func (o *ChargeItems) SetItemMandatory(v string) {
	o.ItemMandatory = &v
}

// GetItemMaximum returns the ItemMaximum field value if set, zero value otherwise.
func (o *ChargeItems) GetItemMaximum() int32 {
	if o == nil || IsNil(o.ItemMaximum) {
		var ret int32
		return ret
	}
	return *o.ItemMaximum
}

// GetItemMaximumOk returns a tuple with the ItemMaximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeItems) GetItemMaximumOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemMaximum) {
		return nil, false
	}
	return o.ItemMaximum, true
}

// HasItemMaximum returns a boolean if a field has been set.
func (o *ChargeItems) HasItemMaximum() bool {
	if o != nil && !IsNil(o.ItemMaximum) {
		return true
	}

	return false
}

// SetItemMaximum gets a reference to the given int32 and assigns it to the ItemMaximum field.
func (o *ChargeItems) SetItemMaximum(v int32) {
	o.ItemMaximum = &v
}

// GetItemName returns the ItemName field value if set, zero value otherwise.
func (o *ChargeItems) GetItemName() string {
	if o == nil || IsNil(o.ItemName) {
		var ret string
		return ret
	}
	return *o.ItemName
}

// GetItemNameOk returns a tuple with the ItemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeItems) GetItemNameOk() (*string, bool) {
	if o == nil || IsNil(o.ItemName) {
		return nil, false
	}
	return o.ItemName, true
}

// HasItemName returns a boolean if a field has been set.
func (o *ChargeItems) HasItemName() bool {
	if o != nil && !IsNil(o.ItemName) {
		return true
	}

	return false
}

// SetItemName gets a reference to the given string and assigns it to the ItemName field.
func (o *ChargeItems) SetItemName(v string) {
	o.ItemName = &v
}

// GetItemPrice returns the ItemPrice field value if set, zero value otherwise.
func (o *ChargeItems) GetItemPrice() string {
	if o == nil || IsNil(o.ItemPrice) {
		var ret string
		return ret
	}
	return *o.ItemPrice
}

// GetItemPriceOk returns a tuple with the ItemPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeItems) GetItemPriceOk() (*string, bool) {
	if o == nil || IsNil(o.ItemPrice) {
		return nil, false
	}
	return o.ItemPrice, true
}

// HasItemPrice returns a boolean if a field has been set.
func (o *ChargeItems) HasItemPrice() bool {
	if o != nil && !IsNil(o.ItemPrice) {
		return true
	}

	return false
}

// SetItemPrice gets a reference to the given string and assigns it to the ItemPrice field.
func (o *ChargeItems) SetItemPrice(v string) {
	o.ItemPrice = &v
}

// GetItemSerialNumber returns the ItemSerialNumber field value if set, zero value otherwise.
func (o *ChargeItems) GetItemSerialNumber() int32 {
	if o == nil || IsNil(o.ItemSerialNumber) {
		var ret int32
		return ret
	}
	return *o.ItemSerialNumber
}

// GetItemSerialNumberOk returns a tuple with the ItemSerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeItems) GetItemSerialNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemSerialNumber) {
		return nil, false
	}
	return o.ItemSerialNumber, true
}

// HasItemSerialNumber returns a boolean if a field has been set.
func (o *ChargeItems) HasItemSerialNumber() bool {
	if o != nil && !IsNil(o.ItemSerialNumber) {
		return true
	}

	return false
}

// SetItemSerialNumber gets a reference to the given int32 and assigns it to the ItemSerialNumber field.
func (o *ChargeItems) SetItemSerialNumber(v int32) {
	o.ItemSerialNumber = &v
}

func (o ChargeItems) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemMandatory) {
		toSerialize["item_mandatory"] = o.ItemMandatory
	}
	if !IsNil(o.ItemMaximum) {
		toSerialize["item_maximum"] = o.ItemMaximum
	}
	if !IsNil(o.ItemName) {
		toSerialize["item_name"] = o.ItemName
	}
	if !IsNil(o.ItemPrice) {
		toSerialize["item_price"] = o.ItemPrice
	}
	if !IsNil(o.ItemSerialNumber) {
		toSerialize["item_serial_number"] = o.ItemSerialNumber
	}
	return toSerialize, nil
}

type NullableChargeItems struct {
	value *ChargeItems
	isSet bool
}

func (v NullableChargeItems) Get() *ChargeItems {
	return v.value
}

func (v *NullableChargeItems) Set(val *ChargeItems) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeItems) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeItems(val *ChargeItems) *NullableChargeItems {
	return &NullableChargeItems{value: val, isSet: true}
}

func (v NullableChargeItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

