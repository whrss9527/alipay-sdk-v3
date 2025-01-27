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

// checks if the CardFrontTextDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardFrontTextDTO{}

// CardFrontTextDTO struct for CardFrontTextDTO
type CardFrontTextDTO struct {
	// 文案标签
	Label *string `json:"label,omitempty"`
	// 展示文案
	Value *string `json:"value,omitempty"`
}

// NewCardFrontTextDTO instantiates a new CardFrontTextDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardFrontTextDTO() *CardFrontTextDTO {
	this := CardFrontTextDTO{}
	return &this
}

// NewCardFrontTextDTOWithDefaults instantiates a new CardFrontTextDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardFrontTextDTOWithDefaults() *CardFrontTextDTO {
	this := CardFrontTextDTO{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CardFrontTextDTO) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardFrontTextDTO) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CardFrontTextDTO) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CardFrontTextDTO) SetLabel(v string) {
	o.Label = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CardFrontTextDTO) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardFrontTextDTO) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CardFrontTextDTO) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CardFrontTextDTO) SetValue(v string) {
	o.Value = &v
}

func (o CardFrontTextDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardFrontTextDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCardFrontTextDTO struct {
	value *CardFrontTextDTO
	isSet bool
}

func (v NullableCardFrontTextDTO) Get() *CardFrontTextDTO {
	return v.value
}

func (v *NullableCardFrontTextDTO) Set(val *CardFrontTextDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCardFrontTextDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCardFrontTextDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardFrontTextDTO(val *CardFrontTextDTO) *NullableCardFrontTextDTO {
	return &NullableCardFrontTextDTO{value: val, isSet: true}
}

func (v NullableCardFrontTextDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardFrontTextDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
