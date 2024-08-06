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

// checks if the GavintestNewonline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GavintestNewonline{}

// GavintestNewonline struct for GavintestNewonline
type GavintestNewonline struct {
	// 1
	Dem []string `json:"dem,omitempty"`
	String *GavintestNewLeveaOne `json:"string,omitempty"`
}

// NewGavintestNewonline instantiates a new GavintestNewonline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGavintestNewonline() *GavintestNewonline {
	this := GavintestNewonline{}
	return &this
}

// NewGavintestNewonlineWithDefaults instantiates a new GavintestNewonline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGavintestNewonlineWithDefaults() *GavintestNewonline {
	this := GavintestNewonline{}
	return &this
}

// GetDem returns the Dem field value if set, zero value otherwise.
func (o *GavintestNewonline) GetDem() []string {
	if o == nil || IsNil(o.Dem) {
		var ret []string
		return ret
	}
	return o.Dem
}

// GetDemOk returns a tuple with the Dem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GavintestNewonline) GetDemOk() ([]string, bool) {
	if o == nil || IsNil(o.Dem) {
		return nil, false
	}
	return o.Dem, true
}

// HasDem returns a boolean if a field has been set.
func (o *GavintestNewonline) HasDem() bool {
	if o != nil && !IsNil(o.Dem) {
		return true
	}

	return false
}

// SetDem gets a reference to the given []string and assigns it to the Dem field.
func (o *GavintestNewonline) SetDem(v []string) {
	o.Dem = v
}

// GetString returns the String field value if set, zero value otherwise.
func (o *GavintestNewonline) GetString() GavintestNewLeveaOne {
	if o == nil || IsNil(o.String) {
		var ret GavintestNewLeveaOne
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GavintestNewonline) GetStringOk() (*GavintestNewLeveaOne, bool) {
	if o == nil || IsNil(o.String) {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *GavintestNewonline) HasString() bool {
	if o != nil && !IsNil(o.String) {
		return true
	}

	return false
}

// SetString gets a reference to the given GavintestNewLeveaOne and assigns it to the String field.
func (o *GavintestNewonline) SetString(v GavintestNewLeveaOne) {
	o.String = &v
}

func (o GavintestNewonline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GavintestNewonline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dem) {
		toSerialize["dem"] = o.Dem
	}
	if !IsNil(o.String) {
		toSerialize["string"] = o.String
	}
	return toSerialize, nil
}

type NullableGavintestNewonline struct {
	value *GavintestNewonline
	isSet bool
}

func (v NullableGavintestNewonline) Get() *GavintestNewonline {
	return v.value
}

func (v *NullableGavintestNewonline) Set(val *GavintestNewonline) {
	v.value = val
	v.isSet = true
}

func (v NullableGavintestNewonline) IsSet() bool {
	return v.isSet
}

func (v *NullableGavintestNewonline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGavintestNewonline(val *GavintestNewonline) *NullableGavintestNewonline {
	return &NullableGavintestNewonline{value: val, isSet: true}
}

func (v NullableGavintestNewonline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGavintestNewonline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

