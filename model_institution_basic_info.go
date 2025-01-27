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

// checks if the InstitutionBasicInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstitutionBasicInfo{}

// InstitutionBasicInfo struct for InstitutionBasicInfo
type InstitutionBasicInfo struct {
	// 费控咨询模式
	ConsultMode *string `json:"consult_mode,omitempty"`
	// 制度是否启用
	Effective *string `json:"effective,omitempty"`
	// 制度生效结束时间
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// 制度生效起始时间
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// 制度描述
	InstitutionDesc *string `json:"institution_desc,omitempty"`
	// 制度id
	InstitutionId *string `json:"institution_id,omitempty"`
	// 制度名称
	InstitutionName *string `json:"institution_name,omitempty"`
}

// NewInstitutionBasicInfo instantiates a new InstitutionBasicInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionBasicInfo() *InstitutionBasicInfo {
	this := InstitutionBasicInfo{}
	return &this
}

// NewInstitutionBasicInfoWithDefaults instantiates a new InstitutionBasicInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionBasicInfoWithDefaults() *InstitutionBasicInfo {
	this := InstitutionBasicInfo{}
	return &this
}

// GetConsultMode returns the ConsultMode field value if set, zero value otherwise.
func (o *InstitutionBasicInfo) GetConsultMode() string {
	if o == nil || IsNil(o.ConsultMode) {
		var ret string
		return ret
	}
	return *o.ConsultMode
}

// GetConsultModeOk returns a tuple with the ConsultMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionBasicInfo) GetConsultModeOk() (*string, bool) {
	if o == nil || IsNil(o.ConsultMode) {
		return nil, false
	}
	return o.ConsultMode, true
}

// HasConsultMode returns a boolean if a field has been set.
func (o *InstitutionBasicInfo) HasConsultMode() bool {
	if o != nil && !IsNil(o.ConsultMode) {
		return true
	}

	return false
}

// SetConsultMode gets a reference to the given string and assigns it to the ConsultMode field.
func (o *InstitutionBasicInfo) SetConsultMode(v string) {
	o.ConsultMode = &v
}

// GetEffective returns the Effective field value if set, zero value otherwise.
func (o *InstitutionBasicInfo) GetEffective() string {
	if o == nil || IsNil(o.Effective) {
		var ret string
		return ret
	}
	return *o.Effective
}

// GetEffectiveOk returns a tuple with the Effective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionBasicInfo) GetEffectiveOk() (*string, bool) {
	if o == nil || IsNil(o.Effective) {
		return nil, false
	}
	return o.Effective, true
}

// HasEffective returns a boolean if a field has been set.
func (o *InstitutionBasicInfo) HasEffective() bool {
	if o != nil && !IsNil(o.Effective) {
		return true
	}

	return false
}

// SetEffective gets a reference to the given string and assigns it to the Effective field.
func (o *InstitutionBasicInfo) SetEffective(v string) {
	o.Effective = &v
}

// GetEffectiveEndDate returns the EffectiveEndDate field value if set, zero value otherwise.
func (o *InstitutionBasicInfo) GetEffectiveEndDate() string {
	if o == nil || IsNil(o.EffectiveEndDate) {
		var ret string
		return ret
	}
	return *o.EffectiveEndDate
}

// GetEffectiveEndDateOk returns a tuple with the EffectiveEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionBasicInfo) GetEffectiveEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveEndDate) {
		return nil, false
	}
	return o.EffectiveEndDate, true
}

// HasEffectiveEndDate returns a boolean if a field has been set.
func (o *InstitutionBasicInfo) HasEffectiveEndDate() bool {
	if o != nil && !IsNil(o.EffectiveEndDate) {
		return true
	}

	return false
}

// SetEffectiveEndDate gets a reference to the given string and assigns it to the EffectiveEndDate field.
func (o *InstitutionBasicInfo) SetEffectiveEndDate(v string) {
	o.EffectiveEndDate = &v
}

// GetEffectiveStartDate returns the EffectiveStartDate field value if set, zero value otherwise.
func (o *InstitutionBasicInfo) GetEffectiveStartDate() string {
	if o == nil || IsNil(o.EffectiveStartDate) {
		var ret string
		return ret
	}
	return *o.EffectiveStartDate
}

// GetEffectiveStartDateOk returns a tuple with the EffectiveStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionBasicInfo) GetEffectiveStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveStartDate) {
		return nil, false
	}
	return o.EffectiveStartDate, true
}

// HasEffectiveStartDate returns a boolean if a field has been set.
func (o *InstitutionBasicInfo) HasEffectiveStartDate() bool {
	if o != nil && !IsNil(o.EffectiveStartDate) {
		return true
	}

	return false
}

// SetEffectiveStartDate gets a reference to the given string and assigns it to the EffectiveStartDate field.
func (o *InstitutionBasicInfo) SetEffectiveStartDate(v string) {
	o.EffectiveStartDate = &v
}

// GetInstitutionDesc returns the InstitutionDesc field value if set, zero value otherwise.
func (o *InstitutionBasicInfo) GetInstitutionDesc() string {
	if o == nil || IsNil(o.InstitutionDesc) {
		var ret string
		return ret
	}
	return *o.InstitutionDesc
}

// GetInstitutionDescOk returns a tuple with the InstitutionDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionBasicInfo) GetInstitutionDescOk() (*string, bool) {
	if o == nil || IsNil(o.InstitutionDesc) {
		return nil, false
	}
	return o.InstitutionDesc, true
}

// HasInstitutionDesc returns a boolean if a field has been set.
func (o *InstitutionBasicInfo) HasInstitutionDesc() bool {
	if o != nil && !IsNil(o.InstitutionDesc) {
		return true
	}

	return false
}

// SetInstitutionDesc gets a reference to the given string and assigns it to the InstitutionDesc field.
func (o *InstitutionBasicInfo) SetInstitutionDesc(v string) {
	o.InstitutionDesc = &v
}

// GetInstitutionId returns the InstitutionId field value if set, zero value otherwise.
func (o *InstitutionBasicInfo) GetInstitutionId() string {
	if o == nil || IsNil(o.InstitutionId) {
		var ret string
		return ret
	}
	return *o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionBasicInfo) GetInstitutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstitutionId) {
		return nil, false
	}
	return o.InstitutionId, true
}

// HasInstitutionId returns a boolean if a field has been set.
func (o *InstitutionBasicInfo) HasInstitutionId() bool {
	if o != nil && !IsNil(o.InstitutionId) {
		return true
	}

	return false
}

// SetInstitutionId gets a reference to the given string and assigns it to the InstitutionId field.
func (o *InstitutionBasicInfo) SetInstitutionId(v string) {
	o.InstitutionId = &v
}

// GetInstitutionName returns the InstitutionName field value if set, zero value otherwise.
func (o *InstitutionBasicInfo) GetInstitutionName() string {
	if o == nil || IsNil(o.InstitutionName) {
		var ret string
		return ret
	}
	return *o.InstitutionName
}

// GetInstitutionNameOk returns a tuple with the InstitutionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionBasicInfo) GetInstitutionNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstitutionName) {
		return nil, false
	}
	return o.InstitutionName, true
}

// HasInstitutionName returns a boolean if a field has been set.
func (o *InstitutionBasicInfo) HasInstitutionName() bool {
	if o != nil && !IsNil(o.InstitutionName) {
		return true
	}

	return false
}

// SetInstitutionName gets a reference to the given string and assigns it to the InstitutionName field.
func (o *InstitutionBasicInfo) SetInstitutionName(v string) {
	o.InstitutionName = &v
}

func (o InstitutionBasicInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstitutionBasicInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsultMode) {
		toSerialize["consult_mode"] = o.ConsultMode
	}
	if !IsNil(o.Effective) {
		toSerialize["effective"] = o.Effective
	}
	if !IsNil(o.EffectiveEndDate) {
		toSerialize["effective_end_date"] = o.EffectiveEndDate
	}
	if !IsNil(o.EffectiveStartDate) {
		toSerialize["effective_start_date"] = o.EffectiveStartDate
	}
	if !IsNil(o.InstitutionDesc) {
		toSerialize["institution_desc"] = o.InstitutionDesc
	}
	if !IsNil(o.InstitutionId) {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if !IsNil(o.InstitutionName) {
		toSerialize["institution_name"] = o.InstitutionName
	}
	return toSerialize, nil
}

type NullableInstitutionBasicInfo struct {
	value *InstitutionBasicInfo
	isSet bool
}

func (v NullableInstitutionBasicInfo) Get() *InstitutionBasicInfo {
	return v.value
}

func (v *NullableInstitutionBasicInfo) Set(val *InstitutionBasicInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionBasicInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionBasicInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionBasicInfo(val *InstitutionBasicInfo) *NullableInstitutionBasicInfo {
	return &NullableInstitutionBasicInfo{value: val, isSet: true}
}

func (v NullableInstitutionBasicInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionBasicInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
