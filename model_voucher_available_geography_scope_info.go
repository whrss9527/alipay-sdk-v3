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

// checks if the VoucherAvailableGeographyScopeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherAvailableGeographyScopeInfo{}

// VoucherAvailableGeographyScopeInfo struct for VoucherAvailableGeographyScopeInfo
type VoucherAvailableGeographyScopeInfo struct {
	AvailableGeographyCityInfo *VoucherAvailableGeographyCityInfo `json:"available_geography_city_info,omitempty"`
	// 券可用地理位置类型。
	AvailableGeographyScopeType *string                            `json:"available_geography_scope_type,omitempty"`
	AvailableGeographyShopInfo  *VoucherAvailableGeographyShopInfo `json:"available_geography_shop_info,omitempty"`
}

// NewVoucherAvailableGeographyScopeInfo instantiates a new VoucherAvailableGeographyScopeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherAvailableGeographyScopeInfo() *VoucherAvailableGeographyScopeInfo {
	this := VoucherAvailableGeographyScopeInfo{}
	return &this
}

// NewVoucherAvailableGeographyScopeInfoWithDefaults instantiates a new VoucherAvailableGeographyScopeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherAvailableGeographyScopeInfoWithDefaults() *VoucherAvailableGeographyScopeInfo {
	this := VoucherAvailableGeographyScopeInfo{}
	return &this
}

// GetAvailableGeographyCityInfo returns the AvailableGeographyCityInfo field value if set, zero value otherwise.
func (o *VoucherAvailableGeographyScopeInfo) GetAvailableGeographyCityInfo() VoucherAvailableGeographyCityInfo {
	if o == nil || IsNil(o.AvailableGeographyCityInfo) {
		var ret VoucherAvailableGeographyCityInfo
		return ret
	}
	return *o.AvailableGeographyCityInfo
}

// GetAvailableGeographyCityInfoOk returns a tuple with the AvailableGeographyCityInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableGeographyScopeInfo) GetAvailableGeographyCityInfoOk() (*VoucherAvailableGeographyCityInfo, bool) {
	if o == nil || IsNil(o.AvailableGeographyCityInfo) {
		return nil, false
	}
	return o.AvailableGeographyCityInfo, true
}

// HasAvailableGeographyCityInfo returns a boolean if a field has been set.
func (o *VoucherAvailableGeographyScopeInfo) HasAvailableGeographyCityInfo() bool {
	if o != nil && !IsNil(o.AvailableGeographyCityInfo) {
		return true
	}

	return false
}

// SetAvailableGeographyCityInfo gets a reference to the given VoucherAvailableGeographyCityInfo and assigns it to the AvailableGeographyCityInfo field.
func (o *VoucherAvailableGeographyScopeInfo) SetAvailableGeographyCityInfo(v VoucherAvailableGeographyCityInfo) {
	o.AvailableGeographyCityInfo = &v
}

// GetAvailableGeographyScopeType returns the AvailableGeographyScopeType field value if set, zero value otherwise.
func (o *VoucherAvailableGeographyScopeInfo) GetAvailableGeographyScopeType() string {
	if o == nil || IsNil(o.AvailableGeographyScopeType) {
		var ret string
		return ret
	}
	return *o.AvailableGeographyScopeType
}

// GetAvailableGeographyScopeTypeOk returns a tuple with the AvailableGeographyScopeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableGeographyScopeInfo) GetAvailableGeographyScopeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AvailableGeographyScopeType) {
		return nil, false
	}
	return o.AvailableGeographyScopeType, true
}

// HasAvailableGeographyScopeType returns a boolean if a field has been set.
func (o *VoucherAvailableGeographyScopeInfo) HasAvailableGeographyScopeType() bool {
	if o != nil && !IsNil(o.AvailableGeographyScopeType) {
		return true
	}

	return false
}

// SetAvailableGeographyScopeType gets a reference to the given string and assigns it to the AvailableGeographyScopeType field.
func (o *VoucherAvailableGeographyScopeInfo) SetAvailableGeographyScopeType(v string) {
	o.AvailableGeographyScopeType = &v
}

// GetAvailableGeographyShopInfo returns the AvailableGeographyShopInfo field value if set, zero value otherwise.
func (o *VoucherAvailableGeographyScopeInfo) GetAvailableGeographyShopInfo() VoucherAvailableGeographyShopInfo {
	if o == nil || IsNil(o.AvailableGeographyShopInfo) {
		var ret VoucherAvailableGeographyShopInfo
		return ret
	}
	return *o.AvailableGeographyShopInfo
}

// GetAvailableGeographyShopInfoOk returns a tuple with the AvailableGeographyShopInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableGeographyScopeInfo) GetAvailableGeographyShopInfoOk() (*VoucherAvailableGeographyShopInfo, bool) {
	if o == nil || IsNil(o.AvailableGeographyShopInfo) {
		return nil, false
	}
	return o.AvailableGeographyShopInfo, true
}

// HasAvailableGeographyShopInfo returns a boolean if a field has been set.
func (o *VoucherAvailableGeographyScopeInfo) HasAvailableGeographyShopInfo() bool {
	if o != nil && !IsNil(o.AvailableGeographyShopInfo) {
		return true
	}

	return false
}

// SetAvailableGeographyShopInfo gets a reference to the given VoucherAvailableGeographyShopInfo and assigns it to the AvailableGeographyShopInfo field.
func (o *VoucherAvailableGeographyScopeInfo) SetAvailableGeographyShopInfo(v VoucherAvailableGeographyShopInfo) {
	o.AvailableGeographyShopInfo = &v
}

func (o VoucherAvailableGeographyScopeInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherAvailableGeographyScopeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailableGeographyCityInfo) {
		toSerialize["available_geography_city_info"] = o.AvailableGeographyCityInfo
	}
	if !IsNil(o.AvailableGeographyScopeType) {
		toSerialize["available_geography_scope_type"] = o.AvailableGeographyScopeType
	}
	if !IsNil(o.AvailableGeographyShopInfo) {
		toSerialize["available_geography_shop_info"] = o.AvailableGeographyShopInfo
	}
	return toSerialize, nil
}

type NullableVoucherAvailableGeographyScopeInfo struct {
	value *VoucherAvailableGeographyScopeInfo
	isSet bool
}

func (v NullableVoucherAvailableGeographyScopeInfo) Get() *VoucherAvailableGeographyScopeInfo {
	return v.value
}

func (v *NullableVoucherAvailableGeographyScopeInfo) Set(val *VoucherAvailableGeographyScopeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherAvailableGeographyScopeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherAvailableGeographyScopeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherAvailableGeographyScopeInfo(val *VoucherAvailableGeographyScopeInfo) *NullableVoucherAvailableGeographyScopeInfo {
	return &NullableVoucherAvailableGeographyScopeInfo{value: val, isSet: true}
}

func (v NullableVoucherAvailableGeographyScopeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherAvailableGeographyScopeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
