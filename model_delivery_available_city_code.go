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

// checks if the DeliveryAvailableCityCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryAvailableCityCode{}

// DeliveryAvailableCityCode struct for DeliveryAvailableCityCode
type DeliveryAvailableCityCode struct {
	// 是否全国。 与city_codes二选一。只允许填true，否则不填。
	AllCity *bool `json:"all_city,omitempty"`
	// 城市编码。与all_city二选一。请按照<a href =\"https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx\">https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx</a>  表格中内容填写。 （参考资料： <a href =\"http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/\">http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/</a> ）
	CityCodes []string `json:"city_codes,omitempty"`
}

// NewDeliveryAvailableCityCode instantiates a new DeliveryAvailableCityCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryAvailableCityCode() *DeliveryAvailableCityCode {
	this := DeliveryAvailableCityCode{}
	return &this
}

// NewDeliveryAvailableCityCodeWithDefaults instantiates a new DeliveryAvailableCityCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryAvailableCityCodeWithDefaults() *DeliveryAvailableCityCode {
	this := DeliveryAvailableCityCode{}
	return &this
}

// GetAllCity returns the AllCity field value if set, zero value otherwise.
func (o *DeliveryAvailableCityCode) GetAllCity() bool {
	if o == nil || IsNil(o.AllCity) {
		var ret bool
		return ret
	}
	return *o.AllCity
}

// GetAllCityOk returns a tuple with the AllCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryAvailableCityCode) GetAllCityOk() (*bool, bool) {
	if o == nil || IsNil(o.AllCity) {
		return nil, false
	}
	return o.AllCity, true
}

// HasAllCity returns a boolean if a field has been set.
func (o *DeliveryAvailableCityCode) HasAllCity() bool {
	if o != nil && !IsNil(o.AllCity) {
		return true
	}

	return false
}

// SetAllCity gets a reference to the given bool and assigns it to the AllCity field.
func (o *DeliveryAvailableCityCode) SetAllCity(v bool) {
	o.AllCity = &v
}

// GetCityCodes returns the CityCodes field value if set, zero value otherwise.
func (o *DeliveryAvailableCityCode) GetCityCodes() []string {
	if o == nil || IsNil(o.CityCodes) {
		var ret []string
		return ret
	}
	return o.CityCodes
}

// GetCityCodesOk returns a tuple with the CityCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryAvailableCityCode) GetCityCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.CityCodes) {
		return nil, false
	}
	return o.CityCodes, true
}

// HasCityCodes returns a boolean if a field has been set.
func (o *DeliveryAvailableCityCode) HasCityCodes() bool {
	if o != nil && !IsNil(o.CityCodes) {
		return true
	}

	return false
}

// SetCityCodes gets a reference to the given []string and assigns it to the CityCodes field.
func (o *DeliveryAvailableCityCode) SetCityCodes(v []string) {
	o.CityCodes = v
}

func (o DeliveryAvailableCityCode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryAvailableCityCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllCity) {
		toSerialize["all_city"] = o.AllCity
	}
	if !IsNil(o.CityCodes) {
		toSerialize["city_codes"] = o.CityCodes
	}
	return toSerialize, nil
}

type NullableDeliveryAvailableCityCode struct {
	value *DeliveryAvailableCityCode
	isSet bool
}

func (v NullableDeliveryAvailableCityCode) Get() *DeliveryAvailableCityCode {
	return v.value
}

func (v *NullableDeliveryAvailableCityCode) Set(val *DeliveryAvailableCityCode) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryAvailableCityCode) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryAvailableCityCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryAvailableCityCode(val *DeliveryAvailableCityCode) *NullableDeliveryAvailableCityCode {
	return &NullableDeliveryAvailableCityCode{value: val, isSet: true}
}

func (v NullableDeliveryAvailableCityCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryAvailableCityCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
