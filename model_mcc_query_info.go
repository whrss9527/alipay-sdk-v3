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

// checks if the MccQueryInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MccQueryInfo{}

// MccQueryInfo struct for MccQueryInfo
type MccQueryInfo struct {
	// 是否特殊行业
	IsSpecial *bool `json:"is_special,omitempty"`
	// 一级类目code
	MccLevel1 *string `json:"mcc_level_1,omitempty"`
	// 商户一级类目名称
	MccLevel1Name *string `json:"mcc_level_1_name,omitempty"`
	// 二级类目code
	MccLevel2 *string `json:"mcc_level_2,omitempty"`
	// 二级类目名称
	MccLevel2Name *string `json:"mcc_level_2_name,omitempty"`
	// 特殊行业需要上传的资质
	MccRequirements *string `json:"mcc_requirements,omitempty"`
	// 是否需要特殊资质
	SpecialQualRequired *bool `json:"special_qual_required,omitempty"`
}

// NewMccQueryInfo instantiates a new MccQueryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMccQueryInfo() *MccQueryInfo {
	this := MccQueryInfo{}
	return &this
}

// NewMccQueryInfoWithDefaults instantiates a new MccQueryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMccQueryInfoWithDefaults() *MccQueryInfo {
	this := MccQueryInfo{}
	return &this
}

// GetIsSpecial returns the IsSpecial field value if set, zero value otherwise.
func (o *MccQueryInfo) GetIsSpecial() bool {
	if o == nil || IsNil(o.IsSpecial) {
		var ret bool
		return ret
	}
	return *o.IsSpecial
}

// GetIsSpecialOk returns a tuple with the IsSpecial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MccQueryInfo) GetIsSpecialOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSpecial) {
		return nil, false
	}
	return o.IsSpecial, true
}

// HasIsSpecial returns a boolean if a field has been set.
func (o *MccQueryInfo) HasIsSpecial() bool {
	if o != nil && !IsNil(o.IsSpecial) {
		return true
	}

	return false
}

// SetIsSpecial gets a reference to the given bool and assigns it to the IsSpecial field.
func (o *MccQueryInfo) SetIsSpecial(v bool) {
	o.IsSpecial = &v
}

// GetMccLevel1 returns the MccLevel1 field value if set, zero value otherwise.
func (o *MccQueryInfo) GetMccLevel1() string {
	if o == nil || IsNil(o.MccLevel1) {
		var ret string
		return ret
	}
	return *o.MccLevel1
}

// GetMccLevel1Ok returns a tuple with the MccLevel1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MccQueryInfo) GetMccLevel1Ok() (*string, bool) {
	if o == nil || IsNil(o.MccLevel1) {
		return nil, false
	}
	return o.MccLevel1, true
}

// HasMccLevel1 returns a boolean if a field has been set.
func (o *MccQueryInfo) HasMccLevel1() bool {
	if o != nil && !IsNil(o.MccLevel1) {
		return true
	}

	return false
}

// SetMccLevel1 gets a reference to the given string and assigns it to the MccLevel1 field.
func (o *MccQueryInfo) SetMccLevel1(v string) {
	o.MccLevel1 = &v
}

// GetMccLevel1Name returns the MccLevel1Name field value if set, zero value otherwise.
func (o *MccQueryInfo) GetMccLevel1Name() string {
	if o == nil || IsNil(o.MccLevel1Name) {
		var ret string
		return ret
	}
	return *o.MccLevel1Name
}

// GetMccLevel1NameOk returns a tuple with the MccLevel1Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MccQueryInfo) GetMccLevel1NameOk() (*string, bool) {
	if o == nil || IsNil(o.MccLevel1Name) {
		return nil, false
	}
	return o.MccLevel1Name, true
}

// HasMccLevel1Name returns a boolean if a field has been set.
func (o *MccQueryInfo) HasMccLevel1Name() bool {
	if o != nil && !IsNil(o.MccLevel1Name) {
		return true
	}

	return false
}

// SetMccLevel1Name gets a reference to the given string and assigns it to the MccLevel1Name field.
func (o *MccQueryInfo) SetMccLevel1Name(v string) {
	o.MccLevel1Name = &v
}

// GetMccLevel2 returns the MccLevel2 field value if set, zero value otherwise.
func (o *MccQueryInfo) GetMccLevel2() string {
	if o == nil || IsNil(o.MccLevel2) {
		var ret string
		return ret
	}
	return *o.MccLevel2
}

// GetMccLevel2Ok returns a tuple with the MccLevel2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MccQueryInfo) GetMccLevel2Ok() (*string, bool) {
	if o == nil || IsNil(o.MccLevel2) {
		return nil, false
	}
	return o.MccLevel2, true
}

// HasMccLevel2 returns a boolean if a field has been set.
func (o *MccQueryInfo) HasMccLevel2() bool {
	if o != nil && !IsNil(o.MccLevel2) {
		return true
	}

	return false
}

// SetMccLevel2 gets a reference to the given string and assigns it to the MccLevel2 field.
func (o *MccQueryInfo) SetMccLevel2(v string) {
	o.MccLevel2 = &v
}

// GetMccLevel2Name returns the MccLevel2Name field value if set, zero value otherwise.
func (o *MccQueryInfo) GetMccLevel2Name() string {
	if o == nil || IsNil(o.MccLevel2Name) {
		var ret string
		return ret
	}
	return *o.MccLevel2Name
}

// GetMccLevel2NameOk returns a tuple with the MccLevel2Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MccQueryInfo) GetMccLevel2NameOk() (*string, bool) {
	if o == nil || IsNil(o.MccLevel2Name) {
		return nil, false
	}
	return o.MccLevel2Name, true
}

// HasMccLevel2Name returns a boolean if a field has been set.
func (o *MccQueryInfo) HasMccLevel2Name() bool {
	if o != nil && !IsNil(o.MccLevel2Name) {
		return true
	}

	return false
}

// SetMccLevel2Name gets a reference to the given string and assigns it to the MccLevel2Name field.
func (o *MccQueryInfo) SetMccLevel2Name(v string) {
	o.MccLevel2Name = &v
}

// GetMccRequirements returns the MccRequirements field value if set, zero value otherwise.
func (o *MccQueryInfo) GetMccRequirements() string {
	if o == nil || IsNil(o.MccRequirements) {
		var ret string
		return ret
	}
	return *o.MccRequirements
}

// GetMccRequirementsOk returns a tuple with the MccRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MccQueryInfo) GetMccRequirementsOk() (*string, bool) {
	if o == nil || IsNil(o.MccRequirements) {
		return nil, false
	}
	return o.MccRequirements, true
}

// HasMccRequirements returns a boolean if a field has been set.
func (o *MccQueryInfo) HasMccRequirements() bool {
	if o != nil && !IsNil(o.MccRequirements) {
		return true
	}

	return false
}

// SetMccRequirements gets a reference to the given string and assigns it to the MccRequirements field.
func (o *MccQueryInfo) SetMccRequirements(v string) {
	o.MccRequirements = &v
}

// GetSpecialQualRequired returns the SpecialQualRequired field value if set, zero value otherwise.
func (o *MccQueryInfo) GetSpecialQualRequired() bool {
	if o == nil || IsNil(o.SpecialQualRequired) {
		var ret bool
		return ret
	}
	return *o.SpecialQualRequired
}

// GetSpecialQualRequiredOk returns a tuple with the SpecialQualRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MccQueryInfo) GetSpecialQualRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.SpecialQualRequired) {
		return nil, false
	}
	return o.SpecialQualRequired, true
}

// HasSpecialQualRequired returns a boolean if a field has been set.
func (o *MccQueryInfo) HasSpecialQualRequired() bool {
	if o != nil && !IsNil(o.SpecialQualRequired) {
		return true
	}

	return false
}

// SetSpecialQualRequired gets a reference to the given bool and assigns it to the SpecialQualRequired field.
func (o *MccQueryInfo) SetSpecialQualRequired(v bool) {
	o.SpecialQualRequired = &v
}

func (o MccQueryInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MccQueryInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsSpecial) {
		toSerialize["is_special"] = o.IsSpecial
	}
	if !IsNil(o.MccLevel1) {
		toSerialize["mcc_level_1"] = o.MccLevel1
	}
	if !IsNil(o.MccLevel1Name) {
		toSerialize["mcc_level_1_name"] = o.MccLevel1Name
	}
	if !IsNil(o.MccLevel2) {
		toSerialize["mcc_level_2"] = o.MccLevel2
	}
	if !IsNil(o.MccLevel2Name) {
		toSerialize["mcc_level_2_name"] = o.MccLevel2Name
	}
	if !IsNil(o.MccRequirements) {
		toSerialize["mcc_requirements"] = o.MccRequirements
	}
	if !IsNil(o.SpecialQualRequired) {
		toSerialize["special_qual_required"] = o.SpecialQualRequired
	}
	return toSerialize, nil
}

type NullableMccQueryInfo struct {
	value *MccQueryInfo
	isSet bool
}

func (v NullableMccQueryInfo) Get() *MccQueryInfo {
	return v.value
}

func (v *NullableMccQueryInfo) Set(val *MccQueryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMccQueryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMccQueryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMccQueryInfo(val *MccQueryInfo) *NullableMccQueryInfo {
	return &NullableMccQueryInfo{value: val, isSet: true}
}

func (v NullableMccQueryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMccQueryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
