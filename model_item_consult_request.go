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

// checks if the ItemConsultRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemConsultRequest{}

// ItemConsultRequest struct for ItemConsultRequest
type ItemConsultRequest struct {
	// 商品id
	ItemId *string `json:"item_id,omitempty"`
	// 商品单价，单位为元，最多2位小数
	Price *string `json:"price,omitempty"`
	// 商品数量
	Quantity *string `json:"quantity,omitempty"`
}

// NewItemConsultRequest instantiates a new ItemConsultRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemConsultRequest() *ItemConsultRequest {
	this := ItemConsultRequest{}
	return &this
}

// NewItemConsultRequestWithDefaults instantiates a new ItemConsultRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemConsultRequestWithDefaults() *ItemConsultRequest {
	this := ItemConsultRequest{}
	return &this
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *ItemConsultRequest) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemConsultRequest) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *ItemConsultRequest) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *ItemConsultRequest) SetItemId(v string) {
	o.ItemId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ItemConsultRequest) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemConsultRequest) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ItemConsultRequest) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *ItemConsultRequest) SetPrice(v string) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ItemConsultRequest) GetQuantity() string {
	if o == nil || IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemConsultRequest) GetQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ItemConsultRequest) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *ItemConsultRequest) SetQuantity(v string) {
	o.Quantity = &v
}

func (o ItemConsultRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemConsultRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemId) {
		toSerialize["item_id"] = o.ItemId
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableItemConsultRequest struct {
	value *ItemConsultRequest
	isSet bool
}

func (v NullableItemConsultRequest) Get() *ItemConsultRequest {
	return v.value
}

func (v *NullableItemConsultRequest) Set(val *ItemConsultRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableItemConsultRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableItemConsultRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemConsultRequest(val *ItemConsultRequest) *NullableItemConsultRequest {
	return &NullableItemConsultRequest{value: val, isSet: true}
}

func (v NullableItemConsultRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemConsultRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

