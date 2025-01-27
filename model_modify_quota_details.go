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

// checks if the ModifyQuotaDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyQuotaDetails{}

// ModifyQuotaDetails struct for ModifyQuotaDetails
type ModifyQuotaDetails struct {
	// 具体额度，单位：元。如金额为空，表示删除已有的额度设置，无已有额度设置则忽略。
	QuotaAmount *string `json:"quota_amount,omitempty"`
	// 额度维度 MONTH/DAY/SINGLE 分别代表月、日、单次
	QuotaDimension *string `json:"quota_dimension,omitempty"`
	// PAYER/PAYEE 额度管控的角色，收or付款方，目前只支持付款方
	Role *string `json:"role,omitempty"`
}

// NewModifyQuotaDetails instantiates a new ModifyQuotaDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyQuotaDetails() *ModifyQuotaDetails {
	this := ModifyQuotaDetails{}
	return &this
}

// NewModifyQuotaDetailsWithDefaults instantiates a new ModifyQuotaDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyQuotaDetailsWithDefaults() *ModifyQuotaDetails {
	this := ModifyQuotaDetails{}
	return &this
}

// GetQuotaAmount returns the QuotaAmount field value if set, zero value otherwise.
func (o *ModifyQuotaDetails) GetQuotaAmount() string {
	if o == nil || IsNil(o.QuotaAmount) {
		var ret string
		return ret
	}
	return *o.QuotaAmount
}

// GetQuotaAmountOk returns a tuple with the QuotaAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyQuotaDetails) GetQuotaAmountOk() (*string, bool) {
	if o == nil || IsNil(o.QuotaAmount) {
		return nil, false
	}
	return o.QuotaAmount, true
}

// HasQuotaAmount returns a boolean if a field has been set.
func (o *ModifyQuotaDetails) HasQuotaAmount() bool {
	if o != nil && !IsNil(o.QuotaAmount) {
		return true
	}

	return false
}

// SetQuotaAmount gets a reference to the given string and assigns it to the QuotaAmount field.
func (o *ModifyQuotaDetails) SetQuotaAmount(v string) {
	o.QuotaAmount = &v
}

// GetQuotaDimension returns the QuotaDimension field value if set, zero value otherwise.
func (o *ModifyQuotaDetails) GetQuotaDimension() string {
	if o == nil || IsNil(o.QuotaDimension) {
		var ret string
		return ret
	}
	return *o.QuotaDimension
}

// GetQuotaDimensionOk returns a tuple with the QuotaDimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyQuotaDetails) GetQuotaDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.QuotaDimension) {
		return nil, false
	}
	return o.QuotaDimension, true
}

// HasQuotaDimension returns a boolean if a field has been set.
func (o *ModifyQuotaDetails) HasQuotaDimension() bool {
	if o != nil && !IsNil(o.QuotaDimension) {
		return true
	}

	return false
}

// SetQuotaDimension gets a reference to the given string and assigns it to the QuotaDimension field.
func (o *ModifyQuotaDetails) SetQuotaDimension(v string) {
	o.QuotaDimension = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ModifyQuotaDetails) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyQuotaDetails) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ModifyQuotaDetails) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ModifyQuotaDetails) SetRole(v string) {
	o.Role = &v
}

func (o ModifyQuotaDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyQuotaDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QuotaAmount) {
		toSerialize["quota_amount"] = o.QuotaAmount
	}
	if !IsNil(o.QuotaDimension) {
		toSerialize["quota_dimension"] = o.QuotaDimension
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableModifyQuotaDetails struct {
	value *ModifyQuotaDetails
	isSet bool
}

func (v NullableModifyQuotaDetails) Get() *ModifyQuotaDetails {
	return v.value
}

func (v *NullableModifyQuotaDetails) Set(val *ModifyQuotaDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyQuotaDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyQuotaDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyQuotaDetails(val *ModifyQuotaDetails) *NullableModifyQuotaDetails {
	return &NullableModifyQuotaDetails{value: val, isSet: true}
}

func (v NullableModifyQuotaDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyQuotaDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
