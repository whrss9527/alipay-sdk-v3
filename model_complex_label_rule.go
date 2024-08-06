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

// checks if the ComplexLabelRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComplexLabelRule{}

// ComplexLabelRule struct for ComplexLabelRule
type ComplexLabelRule struct {
	// 标签id
	LabelId *string `json:"label_id,omitempty"`
	// 标签取值，当有多个取值时用英文\",\"分隔（比如使用in操作符时）；不允许传入下划线\"_\"、竖线\"|\"或者空格\" \"
	LabelValue *string `json:"label_value,omitempty"`
	// 目前支持EQ（等于）、NEQ（不等于）、LT（小于），GT（大于）、LTEQ（小于等于）、GTEQ（大于等于）、LIKE（匹配）、BETWEEN（范围）、IN（包含）、NOTIN（不包含）操作
	Operator *string `json:"operator,omitempty"`
}

// NewComplexLabelRule instantiates a new ComplexLabelRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComplexLabelRule() *ComplexLabelRule {
	this := ComplexLabelRule{}
	return &this
}

// NewComplexLabelRuleWithDefaults instantiates a new ComplexLabelRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComplexLabelRuleWithDefaults() *ComplexLabelRule {
	this := ComplexLabelRule{}
	return &this
}

// GetLabelId returns the LabelId field value if set, zero value otherwise.
func (o *ComplexLabelRule) GetLabelId() string {
	if o == nil || IsNil(o.LabelId) {
		var ret string
		return ret
	}
	return *o.LabelId
}

// GetLabelIdOk returns a tuple with the LabelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplexLabelRule) GetLabelIdOk() (*string, bool) {
	if o == nil || IsNil(o.LabelId) {
		return nil, false
	}
	return o.LabelId, true
}

// HasLabelId returns a boolean if a field has been set.
func (o *ComplexLabelRule) HasLabelId() bool {
	if o != nil && !IsNil(o.LabelId) {
		return true
	}

	return false
}

// SetLabelId gets a reference to the given string and assigns it to the LabelId field.
func (o *ComplexLabelRule) SetLabelId(v string) {
	o.LabelId = &v
}

// GetLabelValue returns the LabelValue field value if set, zero value otherwise.
func (o *ComplexLabelRule) GetLabelValue() string {
	if o == nil || IsNil(o.LabelValue) {
		var ret string
		return ret
	}
	return *o.LabelValue
}

// GetLabelValueOk returns a tuple with the LabelValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplexLabelRule) GetLabelValueOk() (*string, bool) {
	if o == nil || IsNil(o.LabelValue) {
		return nil, false
	}
	return o.LabelValue, true
}

// HasLabelValue returns a boolean if a field has been set.
func (o *ComplexLabelRule) HasLabelValue() bool {
	if o != nil && !IsNil(o.LabelValue) {
		return true
	}

	return false
}

// SetLabelValue gets a reference to the given string and assigns it to the LabelValue field.
func (o *ComplexLabelRule) SetLabelValue(v string) {
	o.LabelValue = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ComplexLabelRule) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComplexLabelRule) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ComplexLabelRule) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *ComplexLabelRule) SetOperator(v string) {
	o.Operator = &v
}

func (o ComplexLabelRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComplexLabelRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LabelId) {
		toSerialize["label_id"] = o.LabelId
	}
	if !IsNil(o.LabelValue) {
		toSerialize["label_value"] = o.LabelValue
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	return toSerialize, nil
}

type NullableComplexLabelRule struct {
	value *ComplexLabelRule
	isSet bool
}

func (v NullableComplexLabelRule) Get() *ComplexLabelRule {
	return v.value
}

func (v *NullableComplexLabelRule) Set(val *ComplexLabelRule) {
	v.value = val
	v.isSet = true
}

func (v NullableComplexLabelRule) IsSet() bool {
	return v.isSet
}

func (v *NullableComplexLabelRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComplexLabelRule(val *ComplexLabelRule) *NullableComplexLabelRule {
	return &NullableComplexLabelRule{value: val, isSet: true}
}

func (v NullableComplexLabelRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComplexLabelRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

