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

// checks if the AlipaySecurityProdSssQueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipaySecurityProdSssQueryModel{}

// AlipaySecurityProdSssQueryModel struct for AlipaySecurityProdSssQueryModel
type AlipaySecurityProdSssQueryModel struct {
	// stst
	Tesst []PromiseDetail `json:"tesst,omitempty"`
	Xxx *JinyouTestFive `json:"xxx,omitempty"`
}

// NewAlipaySecurityProdSssQueryModel instantiates a new AlipaySecurityProdSssQueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipaySecurityProdSssQueryModel() *AlipaySecurityProdSssQueryModel {
	this := AlipaySecurityProdSssQueryModel{}
	return &this
}

// NewAlipaySecurityProdSssQueryModelWithDefaults instantiates a new AlipaySecurityProdSssQueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipaySecurityProdSssQueryModelWithDefaults() *AlipaySecurityProdSssQueryModel {
	this := AlipaySecurityProdSssQueryModel{}
	return &this
}

// GetTesst returns the Tesst field value if set, zero value otherwise.
func (o *AlipaySecurityProdSssQueryModel) GetTesst() []PromiseDetail {
	if o == nil || IsNil(o.Tesst) {
		var ret []PromiseDetail
		return ret
	}
	return o.Tesst
}

// GetTesstOk returns a tuple with the Tesst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipaySecurityProdSssQueryModel) GetTesstOk() ([]PromiseDetail, bool) {
	if o == nil || IsNil(o.Tesst) {
		return nil, false
	}
	return o.Tesst, true
}

// HasTesst returns a boolean if a field has been set.
func (o *AlipaySecurityProdSssQueryModel) HasTesst() bool {
	if o != nil && !IsNil(o.Tesst) {
		return true
	}

	return false
}

// SetTesst gets a reference to the given []PromiseDetail and assigns it to the Tesst field.
func (o *AlipaySecurityProdSssQueryModel) SetTesst(v []PromiseDetail) {
	o.Tesst = v
}

// GetXxx returns the Xxx field value if set, zero value otherwise.
func (o *AlipaySecurityProdSssQueryModel) GetXxx() JinyouTestFive {
	if o == nil || IsNil(o.Xxx) {
		var ret JinyouTestFive
		return ret
	}
	return *o.Xxx
}

// GetXxxOk returns a tuple with the Xxx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipaySecurityProdSssQueryModel) GetXxxOk() (*JinyouTestFive, bool) {
	if o == nil || IsNil(o.Xxx) {
		return nil, false
	}
	return o.Xxx, true
}

// HasXxx returns a boolean if a field has been set.
func (o *AlipaySecurityProdSssQueryModel) HasXxx() bool {
	if o != nil && !IsNil(o.Xxx) {
		return true
	}

	return false
}

// SetXxx gets a reference to the given JinyouTestFive and assigns it to the Xxx field.
func (o *AlipaySecurityProdSssQueryModel) SetXxx(v JinyouTestFive) {
	o.Xxx = &v
}

func (o AlipaySecurityProdSssQueryModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipaySecurityProdSssQueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tesst) {
		toSerialize["tesst"] = o.Tesst
	}
	if !IsNil(o.Xxx) {
		toSerialize["xxx"] = o.Xxx
	}
	return toSerialize, nil
}

type NullableAlipaySecurityProdSssQueryModel struct {
	value *AlipaySecurityProdSssQueryModel
	isSet bool
}

func (v NullableAlipaySecurityProdSssQueryModel) Get() *AlipaySecurityProdSssQueryModel {
	return v.value
}

func (v *NullableAlipaySecurityProdSssQueryModel) Set(val *AlipaySecurityProdSssQueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipaySecurityProdSssQueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipaySecurityProdSssQueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipaySecurityProdSssQueryModel(val *AlipaySecurityProdSssQueryModel) *NullableAlipaySecurityProdSssQueryModel {
	return &NullableAlipaySecurityProdSssQueryModel{value: val, isSet: true}
}

func (v NullableAlipaySecurityProdSssQueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipaySecurityProdSssQueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

