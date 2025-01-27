/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AlipayFundEnterprisepayGroupModifyErrorResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundEnterprisepayGroupModifyErrorResponseModel{}

// AlipayFundEnterprisepayGroupModifyErrorResponseModel struct for AlipayFundEnterprisepayGroupModifyErrorResponseModel
type AlipayFundEnterprisepayGroupModifyErrorResponseModel struct {
	// 错误码
	Code string `json:"code"`
	// 解决方案链接
	Links *string `json:"links,omitempty"`
	// 错误描述
	Message string `json:"message"`
}

type _AlipayFundEnterprisepayGroupModifyErrorResponseModel AlipayFundEnterprisepayGroupModifyErrorResponseModel

// NewAlipayFundEnterprisepayGroupModifyErrorResponseModel instantiates a new AlipayFundEnterprisepayGroupModifyErrorResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundEnterprisepayGroupModifyErrorResponseModel(code string, message string) *AlipayFundEnterprisepayGroupModifyErrorResponseModel {
	this := AlipayFundEnterprisepayGroupModifyErrorResponseModel{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAlipayFundEnterprisepayGroupModifyErrorResponseModelWithDefaults instantiates a new AlipayFundEnterprisepayGroupModifyErrorResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundEnterprisepayGroupModifyErrorResponseModelWithDefaults() *AlipayFundEnterprisepayGroupModifyErrorResponseModel {
	this := AlipayFundEnterprisepayGroupModifyErrorResponseModel{}
	return &this
}

// GetCode returns the Code field value
func (o *AlipayFundEnterprisepayGroupModifyErrorResponseModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepayGroupModifyErrorResponseModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AlipayFundEnterprisepayGroupModifyErrorResponseModel) SetCode(v string) {
	o.Code = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepayGroupModifyErrorResponseModel) GetLinks() string {
	if o == nil || IsNil(o.Links) {
		var ret string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepayGroupModifyErrorResponseModel) GetLinksOk() (*string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepayGroupModifyErrorResponseModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given string and assigns it to the Links field.
func (o *AlipayFundEnterprisepayGroupModifyErrorResponseModel) SetLinks(v string) {
	o.Links = &v
}

// GetMessage returns the Message field value
func (o *AlipayFundEnterprisepayGroupModifyErrorResponseModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepayGroupModifyErrorResponseModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AlipayFundEnterprisepayGroupModifyErrorResponseModel) SetMessage(v string) {
	o.Message = v
}

func (o AlipayFundEnterprisepayGroupModifyErrorResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundEnterprisepayGroupModifyErrorResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *AlipayFundEnterprisepayGroupModifyErrorResponseModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAlipayFundEnterprisepayGroupModifyErrorResponseModel := _AlipayFundEnterprisepayGroupModifyErrorResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlipayFundEnterprisepayGroupModifyErrorResponseModel)

	if err != nil {
		return err
	}

	*o = AlipayFundEnterprisepayGroupModifyErrorResponseModel(varAlipayFundEnterprisepayGroupModifyErrorResponseModel)

	return err
}

type NullableAlipayFundEnterprisepayGroupModifyErrorResponseModel struct {
	value *AlipayFundEnterprisepayGroupModifyErrorResponseModel
	isSet bool
}

func (v NullableAlipayFundEnterprisepayGroupModifyErrorResponseModel) Get() *AlipayFundEnterprisepayGroupModifyErrorResponseModel {
	return v.value
}

func (v *NullableAlipayFundEnterprisepayGroupModifyErrorResponseModel) Set(val *AlipayFundEnterprisepayGroupModifyErrorResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundEnterprisepayGroupModifyErrorResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundEnterprisepayGroupModifyErrorResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundEnterprisepayGroupModifyErrorResponseModel(val *AlipayFundEnterprisepayGroupModifyErrorResponseModel) *NullableAlipayFundEnterprisepayGroupModifyErrorResponseModel {
	return &NullableAlipayFundEnterprisepayGroupModifyErrorResponseModel{value: val, isSet: true}
}

func (v NullableAlipayFundEnterprisepayGroupModifyErrorResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundEnterprisepayGroupModifyErrorResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
