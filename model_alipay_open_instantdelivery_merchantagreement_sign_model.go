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

// checks if the AlipayOpenInstantdeliveryMerchantagreementSignModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenInstantdeliveryMerchantagreementSignModel{}

// AlipayOpenInstantdeliveryMerchantagreementSignModel struct for AlipayOpenInstantdeliveryMerchantagreementSignModel
type AlipayOpenInstantdeliveryMerchantagreementSignModel struct {
	// 蚂蚁统一会员ID
	OpenId *string `json:"open_id,omitempty"`
	// 蚂蚁统一会员ID
	UserId *string `json:"user_id,omitempty"`
}

// NewAlipayOpenInstantdeliveryMerchantagreementSignModel instantiates a new AlipayOpenInstantdeliveryMerchantagreementSignModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenInstantdeliveryMerchantagreementSignModel() *AlipayOpenInstantdeliveryMerchantagreementSignModel {
	this := AlipayOpenInstantdeliveryMerchantagreementSignModel{}
	return &this
}

// NewAlipayOpenInstantdeliveryMerchantagreementSignModelWithDefaults instantiates a new AlipayOpenInstantdeliveryMerchantagreementSignModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenInstantdeliveryMerchantagreementSignModelWithDefaults() *AlipayOpenInstantdeliveryMerchantagreementSignModel {
	this := AlipayOpenInstantdeliveryMerchantagreementSignModel{}
	return &this
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantagreementSignModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantagreementSignModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantagreementSignModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AlipayOpenInstantdeliveryMerchantagreementSignModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantagreementSignModel) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantagreementSignModel) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantagreementSignModel) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AlipayOpenInstantdeliveryMerchantagreementSignModel) SetUserId(v string) {
	o.UserId = &v
}

func (o AlipayOpenInstantdeliveryMerchantagreementSignModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenInstantdeliveryMerchantagreementSignModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableAlipayOpenInstantdeliveryMerchantagreementSignModel struct {
	value *AlipayOpenInstantdeliveryMerchantagreementSignModel
	isSet bool
}

func (v NullableAlipayOpenInstantdeliveryMerchantagreementSignModel) Get() *AlipayOpenInstantdeliveryMerchantagreementSignModel {
	return v.value
}

func (v *NullableAlipayOpenInstantdeliveryMerchantagreementSignModel) Set(val *AlipayOpenInstantdeliveryMerchantagreementSignModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenInstantdeliveryMerchantagreementSignModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenInstantdeliveryMerchantagreementSignModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenInstantdeliveryMerchantagreementSignModel(val *AlipayOpenInstantdeliveryMerchantagreementSignModel) *NullableAlipayOpenInstantdeliveryMerchantagreementSignModel {
	return &NullableAlipayOpenInstantdeliveryMerchantagreementSignModel{value: val, isSet: true}
}

func (v NullableAlipayOpenInstantdeliveryMerchantagreementSignModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenInstantdeliveryMerchantagreementSignModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
