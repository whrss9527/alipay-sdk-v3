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

// checks if the ItemOrderInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemOrderInfo{}

// ItemOrderInfo struct for ItemOrderInfo
type ItemOrderInfo struct {
	// 扩展信息，请参见产品文档。小程序订单助手业务中，扩展参数必须传递素材id；其他业务场景参见对应的产品文档。
	ExtInfo []OrderExtInfo `json:"ext_info,omitempty"`
	// 商品ID
	ItemId *string `json:"item_id,omitempty"`
	// 商品名称
	ItemName *string `json:"item_name,omitempty"`
	// 商品数量（单位：自拟）
	Quantity *int32 `json:"quantity,omitempty"`
	// 商品 sku id
	SkuId *string `json:"sku_id,omitempty"`
	// 商品状态枚举
	Status *string `json:"status,omitempty"`
	// 商品状态描述，默认无需传入，如需使用请联系业务负责人
	StatusDesc *string `json:"status_desc,omitempty"`
	// 商品规格
	Unit *string `json:"unit,omitempty"`
	// 商品单价（单位：元）
	UnitPrice *string `json:"unit_price,omitempty"`
}

// NewItemOrderInfo instantiates a new ItemOrderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemOrderInfo() *ItemOrderInfo {
	this := ItemOrderInfo{}
	return &this
}

// NewItemOrderInfoWithDefaults instantiates a new ItemOrderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemOrderInfoWithDefaults() *ItemOrderInfo {
	this := ItemOrderInfo{}
	return &this
}

// GetExtInfo returns the ExtInfo field value if set, zero value otherwise.
func (o *ItemOrderInfo) GetExtInfo() []OrderExtInfo {
	if o == nil || IsNil(o.ExtInfo) {
		var ret []OrderExtInfo
		return ret
	}
	return o.ExtInfo
}

// GetExtInfoOk returns a tuple with the ExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOrderInfo) GetExtInfoOk() ([]OrderExtInfo, bool) {
	if o == nil || IsNil(o.ExtInfo) {
		return nil, false
	}
	return o.ExtInfo, true
}

// HasExtInfo returns a boolean if a field has been set.
func (o *ItemOrderInfo) HasExtInfo() bool {
	if o != nil && !IsNil(o.ExtInfo) {
		return true
	}

	return false
}

// SetExtInfo gets a reference to the given []OrderExtInfo and assigns it to the ExtInfo field.
func (o *ItemOrderInfo) SetExtInfo(v []OrderExtInfo) {
	o.ExtInfo = v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *ItemOrderInfo) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOrderInfo) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *ItemOrderInfo) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *ItemOrderInfo) SetItemId(v string) {
	o.ItemId = &v
}

// GetItemName returns the ItemName field value if set, zero value otherwise.
func (o *ItemOrderInfo) GetItemName() string {
	if o == nil || IsNil(o.ItemName) {
		var ret string
		return ret
	}
	return *o.ItemName
}

// GetItemNameOk returns a tuple with the ItemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOrderInfo) GetItemNameOk() (*string, bool) {
	if o == nil || IsNil(o.ItemName) {
		return nil, false
	}
	return o.ItemName, true
}

// HasItemName returns a boolean if a field has been set.
func (o *ItemOrderInfo) HasItemName() bool {
	if o != nil && !IsNil(o.ItemName) {
		return true
	}

	return false
}

// SetItemName gets a reference to the given string and assigns it to the ItemName field.
func (o *ItemOrderInfo) SetItemName(v string) {
	o.ItemName = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ItemOrderInfo) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOrderInfo) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ItemOrderInfo) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *ItemOrderInfo) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetSkuId returns the SkuId field value if set, zero value otherwise.
func (o *ItemOrderInfo) GetSkuId() string {
	if o == nil || IsNil(o.SkuId) {
		var ret string
		return ret
	}
	return *o.SkuId
}

// GetSkuIdOk returns a tuple with the SkuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOrderInfo) GetSkuIdOk() (*string, bool) {
	if o == nil || IsNil(o.SkuId) {
		return nil, false
	}
	return o.SkuId, true
}

// HasSkuId returns a boolean if a field has been set.
func (o *ItemOrderInfo) HasSkuId() bool {
	if o != nil && !IsNil(o.SkuId) {
		return true
	}

	return false
}

// SetSkuId gets a reference to the given string and assigns it to the SkuId field.
func (o *ItemOrderInfo) SetSkuId(v string) {
	o.SkuId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ItemOrderInfo) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOrderInfo) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ItemOrderInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ItemOrderInfo) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDesc returns the StatusDesc field value if set, zero value otherwise.
func (o *ItemOrderInfo) GetStatusDesc() string {
	if o == nil || IsNil(o.StatusDesc) {
		var ret string
		return ret
	}
	return *o.StatusDesc
}

// GetStatusDescOk returns a tuple with the StatusDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOrderInfo) GetStatusDescOk() (*string, bool) {
	if o == nil || IsNil(o.StatusDesc) {
		return nil, false
	}
	return o.StatusDesc, true
}

// HasStatusDesc returns a boolean if a field has been set.
func (o *ItemOrderInfo) HasStatusDesc() bool {
	if o != nil && !IsNil(o.StatusDesc) {
		return true
	}

	return false
}

// SetStatusDesc gets a reference to the given string and assigns it to the StatusDesc field.
func (o *ItemOrderInfo) SetStatusDesc(v string) {
	o.StatusDesc = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ItemOrderInfo) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOrderInfo) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ItemOrderInfo) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *ItemOrderInfo) SetUnit(v string) {
	o.Unit = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise.
func (o *ItemOrderInfo) GetUnitPrice() string {
	if o == nil || IsNil(o.UnitPrice) {
		var ret string
		return ret
	}
	return *o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOrderInfo) GetUnitPriceOk() (*string, bool) {
	if o == nil || IsNil(o.UnitPrice) {
		return nil, false
	}
	return o.UnitPrice, true
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *ItemOrderInfo) HasUnitPrice() bool {
	if o != nil && !IsNil(o.UnitPrice) {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given string and assigns it to the UnitPrice field.
func (o *ItemOrderInfo) SetUnitPrice(v string) {
	o.UnitPrice = &v
}

func (o ItemOrderInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemOrderInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtInfo) {
		toSerialize["ext_info"] = o.ExtInfo
	}
	if !IsNil(o.ItemId) {
		toSerialize["item_id"] = o.ItemId
	}
	if !IsNil(o.ItemName) {
		toSerialize["item_name"] = o.ItemName
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.SkuId) {
		toSerialize["sku_id"] = o.SkuId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusDesc) {
		toSerialize["status_desc"] = o.StatusDesc
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.UnitPrice) {
		toSerialize["unit_price"] = o.UnitPrice
	}
	return toSerialize, nil
}

type NullableItemOrderInfo struct {
	value *ItemOrderInfo
	isSet bool
}

func (v NullableItemOrderInfo) Get() *ItemOrderInfo {
	return v.value
}

func (v *NullableItemOrderInfo) Set(val *ItemOrderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableItemOrderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableItemOrderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemOrderInfo(val *ItemOrderInfo) *NullableItemOrderInfo {
	return &NullableItemOrderInfo{value: val, isSet: true}
}

func (v NullableItemOrderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemOrderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
