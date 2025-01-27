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

// checks if the DeliveryDisplayInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryDisplayInfo{}

// DeliveryDisplayInfo struct for DeliveryDisplayInfo
type DeliveryDisplayInfo struct {
	// 副标题。
	MainTitle *string `json:"main_title,omitempty"`
	// 副标题
	SubTitle *string `json:"sub_title,omitempty"`
}

// NewDeliveryDisplayInfo instantiates a new DeliveryDisplayInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryDisplayInfo() *DeliveryDisplayInfo {
	this := DeliveryDisplayInfo{}
	return &this
}

// NewDeliveryDisplayInfoWithDefaults instantiates a new DeliveryDisplayInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryDisplayInfoWithDefaults() *DeliveryDisplayInfo {
	this := DeliveryDisplayInfo{}
	return &this
}

// GetMainTitle returns the MainTitle field value if set, zero value otherwise.
func (o *DeliveryDisplayInfo) GetMainTitle() string {
	if o == nil || IsNil(o.MainTitle) {
		var ret string
		return ret
	}
	return *o.MainTitle
}

// GetMainTitleOk returns a tuple with the MainTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryDisplayInfo) GetMainTitleOk() (*string, bool) {
	if o == nil || IsNil(o.MainTitle) {
		return nil, false
	}
	return o.MainTitle, true
}

// HasMainTitle returns a boolean if a field has been set.
func (o *DeliveryDisplayInfo) HasMainTitle() bool {
	if o != nil && !IsNil(o.MainTitle) {
		return true
	}

	return false
}

// SetMainTitle gets a reference to the given string and assigns it to the MainTitle field.
func (o *DeliveryDisplayInfo) SetMainTitle(v string) {
	o.MainTitle = &v
}

// GetSubTitle returns the SubTitle field value if set, zero value otherwise.
func (o *DeliveryDisplayInfo) GetSubTitle() string {
	if o == nil || IsNil(o.SubTitle) {
		var ret string
		return ret
	}
	return *o.SubTitle
}

// GetSubTitleOk returns a tuple with the SubTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryDisplayInfo) GetSubTitleOk() (*string, bool) {
	if o == nil || IsNil(o.SubTitle) {
		return nil, false
	}
	return o.SubTitle, true
}

// HasSubTitle returns a boolean if a field has been set.
func (o *DeliveryDisplayInfo) HasSubTitle() bool {
	if o != nil && !IsNil(o.SubTitle) {
		return true
	}

	return false
}

// SetSubTitle gets a reference to the given string and assigns it to the SubTitle field.
func (o *DeliveryDisplayInfo) SetSubTitle(v string) {
	o.SubTitle = &v
}

func (o DeliveryDisplayInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryDisplayInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MainTitle) {
		toSerialize["main_title"] = o.MainTitle
	}
	if !IsNil(o.SubTitle) {
		toSerialize["sub_title"] = o.SubTitle
	}
	return toSerialize, nil
}

type NullableDeliveryDisplayInfo struct {
	value *DeliveryDisplayInfo
	isSet bool
}

func (v NullableDeliveryDisplayInfo) Get() *DeliveryDisplayInfo {
	return v.value
}

func (v *NullableDeliveryDisplayInfo) Set(val *DeliveryDisplayInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryDisplayInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryDisplayInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryDisplayInfo(val *DeliveryDisplayInfo) *NullableDeliveryDisplayInfo {
	return &NullableDeliveryDisplayInfo{value: val, isSet: true}
}

func (v NullableDeliveryDisplayInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryDisplayInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
