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

// checks if the VoucherAvailableGeographyShopInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherAvailableGeographyShopInfo{}

// VoucherAvailableGeographyShopInfo struct for VoucherAvailableGeographyShopInfo
type VoucherAvailableGeographyShopInfo struct {
	AvailableGeographyAllShop *VoucherAvailableGeographyAllShopInfo `json:"available_geography_all_shop,omitempty"`
	// 代运营商业关系门店列表，列表中的门店id是调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id 接口参数是列表类型。
	AvailableRealShopIds []string `json:"available_real_shop_ids,omitempty"`
	// 券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id 接口参数是列表类型。
	AvailableShopIds []string `json:"available_shop_ids,omitempty"`
}

// NewVoucherAvailableGeographyShopInfo instantiates a new VoucherAvailableGeographyShopInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherAvailableGeographyShopInfo() *VoucherAvailableGeographyShopInfo {
	this := VoucherAvailableGeographyShopInfo{}
	return &this
}

// NewVoucherAvailableGeographyShopInfoWithDefaults instantiates a new VoucherAvailableGeographyShopInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherAvailableGeographyShopInfoWithDefaults() *VoucherAvailableGeographyShopInfo {
	this := VoucherAvailableGeographyShopInfo{}
	return &this
}

// GetAvailableGeographyAllShop returns the AvailableGeographyAllShop field value if set, zero value otherwise.
func (o *VoucherAvailableGeographyShopInfo) GetAvailableGeographyAllShop() VoucherAvailableGeographyAllShopInfo {
	if o == nil || IsNil(o.AvailableGeographyAllShop) {
		var ret VoucherAvailableGeographyAllShopInfo
		return ret
	}
	return *o.AvailableGeographyAllShop
}

// GetAvailableGeographyAllShopOk returns a tuple with the AvailableGeographyAllShop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableGeographyShopInfo) GetAvailableGeographyAllShopOk() (*VoucherAvailableGeographyAllShopInfo, bool) {
	if o == nil || IsNil(o.AvailableGeographyAllShop) {
		return nil, false
	}
	return o.AvailableGeographyAllShop, true
}

// HasAvailableGeographyAllShop returns a boolean if a field has been set.
func (o *VoucherAvailableGeographyShopInfo) HasAvailableGeographyAllShop() bool {
	if o != nil && !IsNil(o.AvailableGeographyAllShop) {
		return true
	}

	return false
}

// SetAvailableGeographyAllShop gets a reference to the given VoucherAvailableGeographyAllShopInfo and assigns it to the AvailableGeographyAllShop field.
func (o *VoucherAvailableGeographyShopInfo) SetAvailableGeographyAllShop(v VoucherAvailableGeographyAllShopInfo) {
	o.AvailableGeographyAllShop = &v
}

// GetAvailableRealShopIds returns the AvailableRealShopIds field value if set, zero value otherwise.
func (o *VoucherAvailableGeographyShopInfo) GetAvailableRealShopIds() []string {
	if o == nil || IsNil(o.AvailableRealShopIds) {
		var ret []string
		return ret
	}
	return o.AvailableRealShopIds
}

// GetAvailableRealShopIdsOk returns a tuple with the AvailableRealShopIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableGeographyShopInfo) GetAvailableRealShopIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailableRealShopIds) {
		return nil, false
	}
	return o.AvailableRealShopIds, true
}

// HasAvailableRealShopIds returns a boolean if a field has been set.
func (o *VoucherAvailableGeographyShopInfo) HasAvailableRealShopIds() bool {
	if o != nil && !IsNil(o.AvailableRealShopIds) {
		return true
	}

	return false
}

// SetAvailableRealShopIds gets a reference to the given []string and assigns it to the AvailableRealShopIds field.
func (o *VoucherAvailableGeographyShopInfo) SetAvailableRealShopIds(v []string) {
	o.AvailableRealShopIds = v
}

// GetAvailableShopIds returns the AvailableShopIds field value if set, zero value otherwise.
func (o *VoucherAvailableGeographyShopInfo) GetAvailableShopIds() []string {
	if o == nil || IsNil(o.AvailableShopIds) {
		var ret []string
		return ret
	}
	return o.AvailableShopIds
}

// GetAvailableShopIdsOk returns a tuple with the AvailableShopIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableGeographyShopInfo) GetAvailableShopIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailableShopIds) {
		return nil, false
	}
	return o.AvailableShopIds, true
}

// HasAvailableShopIds returns a boolean if a field has been set.
func (o *VoucherAvailableGeographyShopInfo) HasAvailableShopIds() bool {
	if o != nil && !IsNil(o.AvailableShopIds) {
		return true
	}

	return false
}

// SetAvailableShopIds gets a reference to the given []string and assigns it to the AvailableShopIds field.
func (o *VoucherAvailableGeographyShopInfo) SetAvailableShopIds(v []string) {
	o.AvailableShopIds = v
}

func (o VoucherAvailableGeographyShopInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherAvailableGeographyShopInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailableGeographyAllShop) {
		toSerialize["available_geography_all_shop"] = o.AvailableGeographyAllShop
	}
	if !IsNil(o.AvailableRealShopIds) {
		toSerialize["available_real_shop_ids"] = o.AvailableRealShopIds
	}
	if !IsNil(o.AvailableShopIds) {
		toSerialize["available_shop_ids"] = o.AvailableShopIds
	}
	return toSerialize, nil
}

type NullableVoucherAvailableGeographyShopInfo struct {
	value *VoucherAvailableGeographyShopInfo
	isSet bool
}

func (v NullableVoucherAvailableGeographyShopInfo) Get() *VoucherAvailableGeographyShopInfo {
	return v.value
}

func (v *NullableVoucherAvailableGeographyShopInfo) Set(val *VoucherAvailableGeographyShopInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherAvailableGeographyShopInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherAvailableGeographyShopInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherAvailableGeographyShopInfo(val *VoucherAvailableGeographyShopInfo) *NullableVoucherAvailableGeographyShopInfo {
	return &NullableVoucherAvailableGeographyShopInfo{value: val, isSet: true}
}

func (v NullableVoucherAvailableGeographyShopInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherAvailableGeographyShopInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
