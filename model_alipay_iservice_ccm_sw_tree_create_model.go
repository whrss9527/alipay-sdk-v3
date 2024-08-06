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

// checks if the AlipayIserviceCcmSwTreeCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmSwTreeCreateModel{}

// AlipayIserviceCcmSwTreeCreateModel struct for AlipayIserviceCcmSwTreeCreateModel
type AlipayIserviceCcmSwTreeCreateModel struct {
	// 子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// 类目名称
	Name *string `json:"name,omitempty"`
}

// NewAlipayIserviceCcmSwTreeCreateModel instantiates a new AlipayIserviceCcmSwTreeCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmSwTreeCreateModel() *AlipayIserviceCcmSwTreeCreateModel {
	this := AlipayIserviceCcmSwTreeCreateModel{}
	return &this
}

// NewAlipayIserviceCcmSwTreeCreateModelWithDefaults instantiates a new AlipayIserviceCcmSwTreeCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmSwTreeCreateModelWithDefaults() *AlipayIserviceCcmSwTreeCreateModel {
	this := AlipayIserviceCcmSwTreeCreateModel{}
	return &this
}

// GetCcsInstanceId returns the CcsInstanceId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwTreeCreateModel) GetCcsInstanceId() string {
	if o == nil || IsNil(o.CcsInstanceId) {
		var ret string
		return ret
	}
	return *o.CcsInstanceId
}

// GetCcsInstanceIdOk returns a tuple with the CcsInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwTreeCreateModel) GetCcsInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.CcsInstanceId) {
		return nil, false
	}
	return o.CcsInstanceId, true
}

// HasCcsInstanceId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwTreeCreateModel) HasCcsInstanceId() bool {
	if o != nil && !IsNil(o.CcsInstanceId) {
		return true
	}

	return false
}

// SetCcsInstanceId gets a reference to the given string and assigns it to the CcsInstanceId field.
func (o *AlipayIserviceCcmSwTreeCreateModel) SetCcsInstanceId(v string) {
	o.CcsInstanceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlipayIserviceCcmSwTreeCreateModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmSwTreeCreateModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlipayIserviceCcmSwTreeCreateModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlipayIserviceCcmSwTreeCreateModel) SetName(v string) {
	o.Name = &v
}

func (o AlipayIserviceCcmSwTreeCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmSwTreeCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CcsInstanceId) {
		toSerialize["ccs_instance_id"] = o.CcsInstanceId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmSwTreeCreateModel struct {
	value *AlipayIserviceCcmSwTreeCreateModel
	isSet bool
}

func (v NullableAlipayIserviceCcmSwTreeCreateModel) Get() *AlipayIserviceCcmSwTreeCreateModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmSwTreeCreateModel) Set(val *AlipayIserviceCcmSwTreeCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmSwTreeCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmSwTreeCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmSwTreeCreateModel(val *AlipayIserviceCcmSwTreeCreateModel) *NullableAlipayIserviceCcmSwTreeCreateModel {
	return &NullableAlipayIserviceCcmSwTreeCreateModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmSwTreeCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmSwTreeCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

