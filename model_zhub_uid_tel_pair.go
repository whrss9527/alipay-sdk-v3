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

// checks if the ZhubUidTelPair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZhubUidTelPair{}

// ZhubUidTelPair struct for ZhubUidTelPair
type ZhubUidTelPair struct {
	// 支付宝用户open_id
	OpenId *string `json:"open_id,omitempty"`
	// 手机号
	Phone *string `json:"phone,omitempty"`
	// 支付宝uid
	UserId *string `json:"user_id,omitempty"`
}

// NewZhubUidTelPair instantiates a new ZhubUidTelPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZhubUidTelPair() *ZhubUidTelPair {
	this := ZhubUidTelPair{}
	return &this
}

// NewZhubUidTelPairWithDefaults instantiates a new ZhubUidTelPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZhubUidTelPairWithDefaults() *ZhubUidTelPair {
	this := ZhubUidTelPair{}
	return &this
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *ZhubUidTelPair) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhubUidTelPair) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *ZhubUidTelPair) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *ZhubUidTelPair) SetOpenId(v string) {
	o.OpenId = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ZhubUidTelPair) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhubUidTelPair) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *ZhubUidTelPair) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ZhubUidTelPair) SetPhone(v string) {
	o.Phone = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ZhubUidTelPair) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhubUidTelPair) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ZhubUidTelPair) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ZhubUidTelPair) SetUserId(v string) {
	o.UserId = &v
}

func (o ZhubUidTelPair) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZhubUidTelPair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableZhubUidTelPair struct {
	value *ZhubUidTelPair
	isSet bool
}

func (v NullableZhubUidTelPair) Get() *ZhubUidTelPair {
	return v.value
}

func (v *NullableZhubUidTelPair) Set(val *ZhubUidTelPair) {
	v.value = val
	v.isSet = true
}

func (v NullableZhubUidTelPair) IsSet() bool {
	return v.isSet
}

func (v *NullableZhubUidTelPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhubUidTelPair(val *ZhubUidTelPair) *NullableZhubUidTelPair {
	return &NullableZhubUidTelPair{value: val, isSet: true}
}

func (v NullableZhubUidTelPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhubUidTelPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
