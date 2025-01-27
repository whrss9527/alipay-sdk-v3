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

// checks if the VoucherAvailableGeographyCityInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherAvailableGeographyCityInfo{}

// VoucherAvailableGeographyCityInfo struct for VoucherAvailableGeographyCityInfo
type VoucherAvailableGeographyCityInfo struct {
	// 是否全国。 选择全国后，无须填写city_codes字段。系统默认填充全国全部城市信息。
	AllCity *bool `json:"all_city,omitempty"`
	// 城市编码。请按照<a href=\"https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx\">表格</a>中内容填写。（<a href=\"http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm\">参考资料</a>）
	AvailableCityCodes []string `json:"available_city_codes,omitempty"`
}

// NewVoucherAvailableGeographyCityInfo instantiates a new VoucherAvailableGeographyCityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherAvailableGeographyCityInfo() *VoucherAvailableGeographyCityInfo {
	this := VoucherAvailableGeographyCityInfo{}
	return &this
}

// NewVoucherAvailableGeographyCityInfoWithDefaults instantiates a new VoucherAvailableGeographyCityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherAvailableGeographyCityInfoWithDefaults() *VoucherAvailableGeographyCityInfo {
	this := VoucherAvailableGeographyCityInfo{}
	return &this
}

// GetAllCity returns the AllCity field value if set, zero value otherwise.
func (o *VoucherAvailableGeographyCityInfo) GetAllCity() bool {
	if o == nil || IsNil(o.AllCity) {
		var ret bool
		return ret
	}
	return *o.AllCity
}

// GetAllCityOk returns a tuple with the AllCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableGeographyCityInfo) GetAllCityOk() (*bool, bool) {
	if o == nil || IsNil(o.AllCity) {
		return nil, false
	}
	return o.AllCity, true
}

// HasAllCity returns a boolean if a field has been set.
func (o *VoucherAvailableGeographyCityInfo) HasAllCity() bool {
	if o != nil && !IsNil(o.AllCity) {
		return true
	}

	return false
}

// SetAllCity gets a reference to the given bool and assigns it to the AllCity field.
func (o *VoucherAvailableGeographyCityInfo) SetAllCity(v bool) {
	o.AllCity = &v
}

// GetAvailableCityCodes returns the AvailableCityCodes field value if set, zero value otherwise.
func (o *VoucherAvailableGeographyCityInfo) GetAvailableCityCodes() []string {
	if o == nil || IsNil(o.AvailableCityCodes) {
		var ret []string
		return ret
	}
	return o.AvailableCityCodes
}

// GetAvailableCityCodesOk returns a tuple with the AvailableCityCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableGeographyCityInfo) GetAvailableCityCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailableCityCodes) {
		return nil, false
	}
	return o.AvailableCityCodes, true
}

// HasAvailableCityCodes returns a boolean if a field has been set.
func (o *VoucherAvailableGeographyCityInfo) HasAvailableCityCodes() bool {
	if o != nil && !IsNil(o.AvailableCityCodes) {
		return true
	}

	return false
}

// SetAvailableCityCodes gets a reference to the given []string and assigns it to the AvailableCityCodes field.
func (o *VoucherAvailableGeographyCityInfo) SetAvailableCityCodes(v []string) {
	o.AvailableCityCodes = v
}

func (o VoucherAvailableGeographyCityInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherAvailableGeographyCityInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllCity) {
		toSerialize["all_city"] = o.AllCity
	}
	if !IsNil(o.AvailableCityCodes) {
		toSerialize["available_city_codes"] = o.AvailableCityCodes
	}
	return toSerialize, nil
}

type NullableVoucherAvailableGeographyCityInfo struct {
	value *VoucherAvailableGeographyCityInfo
	isSet bool
}

func (v NullableVoucherAvailableGeographyCityInfo) Get() *VoucherAvailableGeographyCityInfo {
	return v.value
}

func (v *NullableVoucherAvailableGeographyCityInfo) Set(val *VoucherAvailableGeographyCityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherAvailableGeographyCityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherAvailableGeographyCityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherAvailableGeographyCityInfo(val *VoucherAvailableGeographyCityInfo) *NullableVoucherAvailableGeographyCityInfo {
	return &NullableVoucherAvailableGeographyCityInfo{value: val, isSet: true}
}

func (v NullableVoucherAvailableGeographyCityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherAvailableGeographyCityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
