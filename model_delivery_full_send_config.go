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

// checks if the DeliveryFullSendConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryFullSendConfig{}

// DeliveryFullSendConfig struct for DeliveryFullSendConfig
type DeliveryFullSendConfig struct {
	DeliveryContentInfo *DeliveryContentInfo `json:"delivery_content_info,omitempty"`
	// 满足消费金额门槛，左闭区间，单位元。 说明：限制支付时的订单金额最少满足的金额门槛。 限制：value > 0 && value <= 99999。
	DeliveryFloorAmount *string `json:"delivery_floor_amount,omitempty"`
}

// NewDeliveryFullSendConfig instantiates a new DeliveryFullSendConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryFullSendConfig() *DeliveryFullSendConfig {
	this := DeliveryFullSendConfig{}
	return &this
}

// NewDeliveryFullSendConfigWithDefaults instantiates a new DeliveryFullSendConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryFullSendConfigWithDefaults() *DeliveryFullSendConfig {
	this := DeliveryFullSendConfig{}
	return &this
}

// GetDeliveryContentInfo returns the DeliveryContentInfo field value if set, zero value otherwise.
func (o *DeliveryFullSendConfig) GetDeliveryContentInfo() DeliveryContentInfo {
	if o == nil || IsNil(o.DeliveryContentInfo) {
		var ret DeliveryContentInfo
		return ret
	}
	return *o.DeliveryContentInfo
}

// GetDeliveryContentInfoOk returns a tuple with the DeliveryContentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryFullSendConfig) GetDeliveryContentInfoOk() (*DeliveryContentInfo, bool) {
	if o == nil || IsNil(o.DeliveryContentInfo) {
		return nil, false
	}
	return o.DeliveryContentInfo, true
}

// HasDeliveryContentInfo returns a boolean if a field has been set.
func (o *DeliveryFullSendConfig) HasDeliveryContentInfo() bool {
	if o != nil && !IsNil(o.DeliveryContentInfo) {
		return true
	}

	return false
}

// SetDeliveryContentInfo gets a reference to the given DeliveryContentInfo and assigns it to the DeliveryContentInfo field.
func (o *DeliveryFullSendConfig) SetDeliveryContentInfo(v DeliveryContentInfo) {
	o.DeliveryContentInfo = &v
}

// GetDeliveryFloorAmount returns the DeliveryFloorAmount field value if set, zero value otherwise.
func (o *DeliveryFullSendConfig) GetDeliveryFloorAmount() string {
	if o == nil || IsNil(o.DeliveryFloorAmount) {
		var ret string
		return ret
	}
	return *o.DeliveryFloorAmount
}

// GetDeliveryFloorAmountOk returns a tuple with the DeliveryFloorAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryFullSendConfig) GetDeliveryFloorAmountOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryFloorAmount) {
		return nil, false
	}
	return o.DeliveryFloorAmount, true
}

// HasDeliveryFloorAmount returns a boolean if a field has been set.
func (o *DeliveryFullSendConfig) HasDeliveryFloorAmount() bool {
	if o != nil && !IsNil(o.DeliveryFloorAmount) {
		return true
	}

	return false
}

// SetDeliveryFloorAmount gets a reference to the given string and assigns it to the DeliveryFloorAmount field.
func (o *DeliveryFullSendConfig) SetDeliveryFloorAmount(v string) {
	o.DeliveryFloorAmount = &v
}

func (o DeliveryFullSendConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryFullSendConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeliveryContentInfo) {
		toSerialize["delivery_content_info"] = o.DeliveryContentInfo
	}
	if !IsNil(o.DeliveryFloorAmount) {
		toSerialize["delivery_floor_amount"] = o.DeliveryFloorAmount
	}
	return toSerialize, nil
}

type NullableDeliveryFullSendConfig struct {
	value *DeliveryFullSendConfig
	isSet bool
}

func (v NullableDeliveryFullSendConfig) Get() *DeliveryFullSendConfig {
	return v.value
}

func (v *NullableDeliveryFullSendConfig) Set(val *DeliveryFullSendConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryFullSendConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryFullSendConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryFullSendConfig(val *DeliveryFullSendConfig) *NullableDeliveryFullSendConfig {
	return &NullableDeliveryFullSendConfig{value: val, isSet: true}
}

func (v NullableDeliveryFullSendConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryFullSendConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

