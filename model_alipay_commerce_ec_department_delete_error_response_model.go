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

// checks if the AlipayCommerceEcDepartmentDeleteErrorResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceEcDepartmentDeleteErrorResponseModel{}

// AlipayCommerceEcDepartmentDeleteErrorResponseModel struct for AlipayCommerceEcDepartmentDeleteErrorResponseModel
type AlipayCommerceEcDepartmentDeleteErrorResponseModel struct {
	// 错误码
	Code string `json:"code"`
	// 解决方案链接
	Links *string `json:"links,omitempty"`
	// 错误描述
	Message string `json:"message"`
}

type _AlipayCommerceEcDepartmentDeleteErrorResponseModel AlipayCommerceEcDepartmentDeleteErrorResponseModel

// NewAlipayCommerceEcDepartmentDeleteErrorResponseModel instantiates a new AlipayCommerceEcDepartmentDeleteErrorResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceEcDepartmentDeleteErrorResponseModel(code string, message string) *AlipayCommerceEcDepartmentDeleteErrorResponseModel {
	this := AlipayCommerceEcDepartmentDeleteErrorResponseModel{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAlipayCommerceEcDepartmentDeleteErrorResponseModelWithDefaults instantiates a new AlipayCommerceEcDepartmentDeleteErrorResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceEcDepartmentDeleteErrorResponseModelWithDefaults() *AlipayCommerceEcDepartmentDeleteErrorResponseModel {
	this := AlipayCommerceEcDepartmentDeleteErrorResponseModel{}
	return &this
}

// GetCode returns the Code field value
func (o *AlipayCommerceEcDepartmentDeleteErrorResponseModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcDepartmentDeleteErrorResponseModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AlipayCommerceEcDepartmentDeleteErrorResponseModel) SetCode(v string) {
	o.Code = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlipayCommerceEcDepartmentDeleteErrorResponseModel) GetLinks() string {
	if o == nil || IsNil(o.Links) {
		var ret string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcDepartmentDeleteErrorResponseModel) GetLinksOk() (*string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlipayCommerceEcDepartmentDeleteErrorResponseModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given string and assigns it to the Links field.
func (o *AlipayCommerceEcDepartmentDeleteErrorResponseModel) SetLinks(v string) {
	o.Links = &v
}

// GetMessage returns the Message field value
func (o *AlipayCommerceEcDepartmentDeleteErrorResponseModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcDepartmentDeleteErrorResponseModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AlipayCommerceEcDepartmentDeleteErrorResponseModel) SetMessage(v string) {
	o.Message = v
}

func (o AlipayCommerceEcDepartmentDeleteErrorResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceEcDepartmentDeleteErrorResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *AlipayCommerceEcDepartmentDeleteErrorResponseModel) UnmarshalJSON(data []byte) (err error) {
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

	varAlipayCommerceEcDepartmentDeleteErrorResponseModel := _AlipayCommerceEcDepartmentDeleteErrorResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlipayCommerceEcDepartmentDeleteErrorResponseModel)

	if err != nil {
		return err
	}

	*o = AlipayCommerceEcDepartmentDeleteErrorResponseModel(varAlipayCommerceEcDepartmentDeleteErrorResponseModel)

	return err
}

type NullableAlipayCommerceEcDepartmentDeleteErrorResponseModel struct {
	value *AlipayCommerceEcDepartmentDeleteErrorResponseModel
	isSet bool
}

func (v NullableAlipayCommerceEcDepartmentDeleteErrorResponseModel) Get() *AlipayCommerceEcDepartmentDeleteErrorResponseModel {
	return v.value
}

func (v *NullableAlipayCommerceEcDepartmentDeleteErrorResponseModel) Set(val *AlipayCommerceEcDepartmentDeleteErrorResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcDepartmentDeleteErrorResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcDepartmentDeleteErrorResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcDepartmentDeleteErrorResponseModel(val *AlipayCommerceEcDepartmentDeleteErrorResponseModel) *NullableAlipayCommerceEcDepartmentDeleteErrorResponseModel {
	return &NullableAlipayCommerceEcDepartmentDeleteErrorResponseModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcDepartmentDeleteErrorResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcDepartmentDeleteErrorResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
