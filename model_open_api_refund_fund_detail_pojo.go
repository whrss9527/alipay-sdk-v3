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

// checks if the OpenApiRefundFundDetailPojo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenApiRefundFundDetailPojo{}

// OpenApiRefundFundDetailPojo struct for OpenApiRefundFundDetailPojo
type OpenApiRefundFundDetailPojo struct {
	// 退款资金明细  详：  若type为paySerialNo则funds为  [{\"paySerialNo\":\"支付流水1\",\"refundFee\":\"退款金额1\"},{\"paySerialNo\":\"支付流水2\",\"refundFee\":\"退款金额2\"}]
	Funds []string `json:"funds,omitempty"`
	// 收入方账户  为空则原路退回
	TransIn *string `json:"trans_in,omitempty"`
	// 收入方账户类型
	TransInType *string `json:"trans_in_type,omitempty"`
	// 描述资金明细类型  详：若type为paySerialNo  则funds参数中体现的都为对应支付流水的退款明细
	Type *string `json:"type,omitempty"`
}

// NewOpenApiRefundFundDetailPojo instantiates a new OpenApiRefundFundDetailPojo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenApiRefundFundDetailPojo() *OpenApiRefundFundDetailPojo {
	this := OpenApiRefundFundDetailPojo{}
	return &this
}

// NewOpenApiRefundFundDetailPojoWithDefaults instantiates a new OpenApiRefundFundDetailPojo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenApiRefundFundDetailPojoWithDefaults() *OpenApiRefundFundDetailPojo {
	this := OpenApiRefundFundDetailPojo{}
	return &this
}

// GetFunds returns the Funds field value if set, zero value otherwise.
func (o *OpenApiRefundFundDetailPojo) GetFunds() []string {
	if o == nil || IsNil(o.Funds) {
		var ret []string
		return ret
	}
	return o.Funds
}

// GetFundsOk returns a tuple with the Funds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenApiRefundFundDetailPojo) GetFundsOk() ([]string, bool) {
	if o == nil || IsNil(o.Funds) {
		return nil, false
	}
	return o.Funds, true
}

// HasFunds returns a boolean if a field has been set.
func (o *OpenApiRefundFundDetailPojo) HasFunds() bool {
	if o != nil && !IsNil(o.Funds) {
		return true
	}

	return false
}

// SetFunds gets a reference to the given []string and assigns it to the Funds field.
func (o *OpenApiRefundFundDetailPojo) SetFunds(v []string) {
	o.Funds = v
}

// GetTransIn returns the TransIn field value if set, zero value otherwise.
func (o *OpenApiRefundFundDetailPojo) GetTransIn() string {
	if o == nil || IsNil(o.TransIn) {
		var ret string
		return ret
	}
	return *o.TransIn
}

// GetTransInOk returns a tuple with the TransIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenApiRefundFundDetailPojo) GetTransInOk() (*string, bool) {
	if o == nil || IsNil(o.TransIn) {
		return nil, false
	}
	return o.TransIn, true
}

// HasTransIn returns a boolean if a field has been set.
func (o *OpenApiRefundFundDetailPojo) HasTransIn() bool {
	if o != nil && !IsNil(o.TransIn) {
		return true
	}

	return false
}

// SetTransIn gets a reference to the given string and assigns it to the TransIn field.
func (o *OpenApiRefundFundDetailPojo) SetTransIn(v string) {
	o.TransIn = &v
}

// GetTransInType returns the TransInType field value if set, zero value otherwise.
func (o *OpenApiRefundFundDetailPojo) GetTransInType() string {
	if o == nil || IsNil(o.TransInType) {
		var ret string
		return ret
	}
	return *o.TransInType
}

// GetTransInTypeOk returns a tuple with the TransInType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenApiRefundFundDetailPojo) GetTransInTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransInType) {
		return nil, false
	}
	return o.TransInType, true
}

// HasTransInType returns a boolean if a field has been set.
func (o *OpenApiRefundFundDetailPojo) HasTransInType() bool {
	if o != nil && !IsNil(o.TransInType) {
		return true
	}

	return false
}

// SetTransInType gets a reference to the given string and assigns it to the TransInType field.
func (o *OpenApiRefundFundDetailPojo) SetTransInType(v string) {
	o.TransInType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OpenApiRefundFundDetailPojo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenApiRefundFundDetailPojo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OpenApiRefundFundDetailPojo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OpenApiRefundFundDetailPojo) SetType(v string) {
	o.Type = &v
}

func (o OpenApiRefundFundDetailPojo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenApiRefundFundDetailPojo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Funds) {
		toSerialize["funds"] = o.Funds
	}
	if !IsNil(o.TransIn) {
		toSerialize["trans_in"] = o.TransIn
	}
	if !IsNil(o.TransInType) {
		toSerialize["trans_in_type"] = o.TransInType
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableOpenApiRefundFundDetailPojo struct {
	value *OpenApiRefundFundDetailPojo
	isSet bool
}

func (v NullableOpenApiRefundFundDetailPojo) Get() *OpenApiRefundFundDetailPojo {
	return v.value
}

func (v *NullableOpenApiRefundFundDetailPojo) Set(val *OpenApiRefundFundDetailPojo) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenApiRefundFundDetailPojo) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenApiRefundFundDetailPojo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenApiRefundFundDetailPojo(val *OpenApiRefundFundDetailPojo) *NullableOpenApiRefundFundDetailPojo {
	return &NullableOpenApiRefundFundDetailPojo{value: val, isSet: true}
}

func (v NullableOpenApiRefundFundDetailPojo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenApiRefundFundDetailPojo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
