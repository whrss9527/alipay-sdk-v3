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

// checks if the AlipayMarketingActivityOrdervoucherDisassociateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityOrdervoucherDisassociateResponseModel{}

// AlipayMarketingActivityOrdervoucherDisassociateResponseModel struct for AlipayMarketingActivityOrdervoucherDisassociateResponseModel
type AlipayMarketingActivityOrdervoucherDisassociateResponseModel struct {
	// 支付宝系统取消关联订单成功的时间。 格式为：yyyy-MM-dd HH:mm:ss
	DisassociateTime *string `json:"disassociate_time,omitempty"`
}

// NewAlipayMarketingActivityOrdervoucherDisassociateResponseModel instantiates a new AlipayMarketingActivityOrdervoucherDisassociateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityOrdervoucherDisassociateResponseModel() *AlipayMarketingActivityOrdervoucherDisassociateResponseModel {
	this := AlipayMarketingActivityOrdervoucherDisassociateResponseModel{}
	return &this
}

// NewAlipayMarketingActivityOrdervoucherDisassociateResponseModelWithDefaults instantiates a new AlipayMarketingActivityOrdervoucherDisassociateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityOrdervoucherDisassociateResponseModelWithDefaults() *AlipayMarketingActivityOrdervoucherDisassociateResponseModel {
	this := AlipayMarketingActivityOrdervoucherDisassociateResponseModel{}
	return &this
}

// GetDisassociateTime returns the DisassociateTime field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherDisassociateResponseModel) GetDisassociateTime() string {
	if o == nil || IsNil(o.DisassociateTime) {
		var ret string
		return ret
	}
	return *o.DisassociateTime
}

// GetDisassociateTimeOk returns a tuple with the DisassociateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherDisassociateResponseModel) GetDisassociateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DisassociateTime) {
		return nil, false
	}
	return o.DisassociateTime, true
}

// HasDisassociateTime returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherDisassociateResponseModel) HasDisassociateTime() bool {
	if o != nil && !IsNil(o.DisassociateTime) {
		return true
	}

	return false
}

// SetDisassociateTime gets a reference to the given string and assigns it to the DisassociateTime field.
func (o *AlipayMarketingActivityOrdervoucherDisassociateResponseModel) SetDisassociateTime(v string) {
	o.DisassociateTime = &v
}

func (o AlipayMarketingActivityOrdervoucherDisassociateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityOrdervoucherDisassociateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisassociateTime) {
		toSerialize["disassociate_time"] = o.DisassociateTime
	}
	return toSerialize, nil
}

type NullableAlipayMarketingActivityOrdervoucherDisassociateResponseModel struct {
	value *AlipayMarketingActivityOrdervoucherDisassociateResponseModel
	isSet bool
}

func (v NullableAlipayMarketingActivityOrdervoucherDisassociateResponseModel) Get() *AlipayMarketingActivityOrdervoucherDisassociateResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityOrdervoucherDisassociateResponseModel) Set(val *AlipayMarketingActivityOrdervoucherDisassociateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityOrdervoucherDisassociateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityOrdervoucherDisassociateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityOrdervoucherDisassociateResponseModel(val *AlipayMarketingActivityOrdervoucherDisassociateResponseModel) *NullableAlipayMarketingActivityOrdervoucherDisassociateResponseModel {
	return &NullableAlipayMarketingActivityOrdervoucherDisassociateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityOrdervoucherDisassociateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityOrdervoucherDisassociateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
