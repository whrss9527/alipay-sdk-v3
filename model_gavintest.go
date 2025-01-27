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

// checks if the Gavintest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Gavintest{}

// Gavintest struct for Gavintest
type Gavintest struct {
	// 测试
	Newid *int32 `json:"newid,omitempty"`
}

// NewGavintest instantiates a new Gavintest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGavintest() *Gavintest {
	this := Gavintest{}
	return &this
}

// NewGavintestWithDefaults instantiates a new Gavintest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGavintestWithDefaults() *Gavintest {
	this := Gavintest{}
	return &this
}

// GetNewid returns the Newid field value if set, zero value otherwise.
func (o *Gavintest) GetNewid() int32 {
	if o == nil || IsNil(o.Newid) {
		var ret int32
		return ret
	}
	return *o.Newid
}

// GetNewidOk returns a tuple with the Newid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gavintest) GetNewidOk() (*int32, bool) {
	if o == nil || IsNil(o.Newid) {
		return nil, false
	}
	return o.Newid, true
}

// HasNewid returns a boolean if a field has been set.
func (o *Gavintest) HasNewid() bool {
	if o != nil && !IsNil(o.Newid) {
		return true
	}

	return false
}

// SetNewid gets a reference to the given int32 and assigns it to the Newid field.
func (o *Gavintest) SetNewid(v int32) {
	o.Newid = &v
}

func (o Gavintest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Gavintest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Newid) {
		toSerialize["newid"] = o.Newid
	}
	return toSerialize, nil
}

type NullableGavintest struct {
	value *Gavintest
	isSet bool
}

func (v NullableGavintest) Get() *Gavintest {
	return v.value
}

func (v *NullableGavintest) Set(val *Gavintest) {
	v.value = val
	v.isSet = true
}

func (v NullableGavintest) IsSet() bool {
	return v.isSet
}

func (v *NullableGavintest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGavintest(val *Gavintest) *NullableGavintest {
	return &NullableGavintest{value: val, isSet: true}
}

func (v NullableGavintest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGavintest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
