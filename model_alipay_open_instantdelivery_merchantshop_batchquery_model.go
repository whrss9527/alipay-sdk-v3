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

// checks if the AlipayOpenInstantdeliveryMerchantshopBatchqueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenInstantdeliveryMerchantshopBatchqueryModel{}

// AlipayOpenInstantdeliveryMerchantshopBatchqueryModel struct for AlipayOpenInstantdeliveryMerchantshopBatchqueryModel
type AlipayOpenInstantdeliveryMerchantshopBatchqueryModel struct {
	// 当前页码
	PageNo *int32 `json:"page_no,omitempty"`
	// 分页数量, 最大50。
	PageSize *int32 `json:"page_size,omitempty"`
	// 门店名称
	ShopName *string `json:"shop_name,omitempty"`
	// 商家门店编码。
	ShopNo *string `json:"shop_no,omitempty"`
}

// NewAlipayOpenInstantdeliveryMerchantshopBatchqueryModel instantiates a new AlipayOpenInstantdeliveryMerchantshopBatchqueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenInstantdeliveryMerchantshopBatchqueryModel() *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel {
	this := AlipayOpenInstantdeliveryMerchantshopBatchqueryModel{}
	return &this
}

// NewAlipayOpenInstantdeliveryMerchantshopBatchqueryModelWithDefaults instantiates a new AlipayOpenInstantdeliveryMerchantshopBatchqueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenInstantdeliveryMerchantshopBatchqueryModelWithDefaults() *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel {
	this := AlipayOpenInstantdeliveryMerchantshopBatchqueryModel{}
	return &this
}

// GetPageNo returns the PageNo field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) GetPageNo() int32 {
	if o == nil || IsNil(o.PageNo) {
		var ret int32
		return ret
	}
	return *o.PageNo
}

// GetPageNoOk returns a tuple with the PageNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) GetPageNoOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNo) {
		return nil, false
	}
	return o.PageNo, true
}

// HasPageNo returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) HasPageNo() bool {
	if o != nil && !IsNil(o.PageNo) {
		return true
	}

	return false
}

// SetPageNo gets a reference to the given int32 and assigns it to the PageNo field.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) SetPageNo(v int32) {
	o.PageNo = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetShopName returns the ShopName field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) GetShopName() string {
	if o == nil || IsNil(o.ShopName) {
		var ret string
		return ret
	}
	return *o.ShopName
}

// GetShopNameOk returns a tuple with the ShopName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) GetShopNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShopName) {
		return nil, false
	}
	return o.ShopName, true
}

// HasShopName returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) HasShopName() bool {
	if o != nil && !IsNil(o.ShopName) {
		return true
	}

	return false
}

// SetShopName gets a reference to the given string and assigns it to the ShopName field.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) SetShopName(v string) {
	o.ShopName = &v
}

// GetShopNo returns the ShopNo field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) GetShopNo() string {
	if o == nil || IsNil(o.ShopNo) {
		var ret string
		return ret
	}
	return *o.ShopNo
}

// GetShopNoOk returns a tuple with the ShopNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) GetShopNoOk() (*string, bool) {
	if o == nil || IsNil(o.ShopNo) {
		return nil, false
	}
	return o.ShopNo, true
}

// HasShopNo returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) HasShopNo() bool {
	if o != nil && !IsNil(o.ShopNo) {
		return true
	}

	return false
}

// SetShopNo gets a reference to the given string and assigns it to the ShopNo field.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) SetShopNo(v string) {
	o.ShopNo = &v
}

func (o AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageNo) {
		toSerialize["page_no"] = o.PageNo
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.ShopName) {
		toSerialize["shop_name"] = o.ShopName
	}
	if !IsNil(o.ShopNo) {
		toSerialize["shop_no"] = o.ShopNo
	}
	return toSerialize, nil
}

type NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryModel struct {
	value *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel
	isSet bool
}

func (v NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryModel) Get() *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel {
	return v.value
}

func (v *NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryModel) Set(val *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenInstantdeliveryMerchantshopBatchqueryModel(val *AlipayOpenInstantdeliveryMerchantshopBatchqueryModel) *NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryModel {
	return &NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryModel{value: val, isSet: true}
}

func (v NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

