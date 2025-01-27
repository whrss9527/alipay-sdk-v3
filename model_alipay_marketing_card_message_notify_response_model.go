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

// checks if the AlipayMarketingCardMessageNotifyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingCardMessageNotifyResponseModel{}

// AlipayMarketingCardMessageNotifyResponseModel struct for AlipayMarketingCardMessageNotifyResponseModel
type AlipayMarketingCardMessageNotifyResponseModel struct {
	// 二级错误处理结果（如果公用返回结果为false，则可以看这个接口判断明细原因） 如果公用返回为true，则该字段为空
	ResultCode *string `json:"result_code,omitempty"`
}

// NewAlipayMarketingCardMessageNotifyResponseModel instantiates a new AlipayMarketingCardMessageNotifyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingCardMessageNotifyResponseModel() *AlipayMarketingCardMessageNotifyResponseModel {
	this := AlipayMarketingCardMessageNotifyResponseModel{}
	return &this
}

// NewAlipayMarketingCardMessageNotifyResponseModelWithDefaults instantiates a new AlipayMarketingCardMessageNotifyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingCardMessageNotifyResponseModelWithDefaults() *AlipayMarketingCardMessageNotifyResponseModel {
	this := AlipayMarketingCardMessageNotifyResponseModel{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *AlipayMarketingCardMessageNotifyResponseModel) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardMessageNotifyResponseModel) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *AlipayMarketingCardMessageNotifyResponseModel) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *AlipayMarketingCardMessageNotifyResponseModel) SetResultCode(v string) {
	o.ResultCode = &v
}

func (o AlipayMarketingCardMessageNotifyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingCardMessageNotifyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultCode) {
		toSerialize["result_code"] = o.ResultCode
	}
	return toSerialize, nil
}

type NullableAlipayMarketingCardMessageNotifyResponseModel struct {
	value *AlipayMarketingCardMessageNotifyResponseModel
	isSet bool
}

func (v NullableAlipayMarketingCardMessageNotifyResponseModel) Get() *AlipayMarketingCardMessageNotifyResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingCardMessageNotifyResponseModel) Set(val *AlipayMarketingCardMessageNotifyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardMessageNotifyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardMessageNotifyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardMessageNotifyResponseModel(val *AlipayMarketingCardMessageNotifyResponseModel) *NullableAlipayMarketingCardMessageNotifyResponseModel {
	return &NullableAlipayMarketingCardMessageNotifyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardMessageNotifyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardMessageNotifyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
