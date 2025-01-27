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

// checks if the AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel{}

// AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel struct for AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel
type AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel struct {
	// 错误码
	Code string `json:"code"`
	// 解决方案链接
	Links *string `json:"links,omitempty"`
	// 错误描述
	Message string `json:"message"`
}

type _AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel

// NewAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel instantiates a new AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel(code string, message string) *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel {
	this := AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModelWithDefaults instantiates a new AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModelWithDefaults() *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel {
	this := AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel{}
	return &this
}

// GetCode returns the Code field value
func (o *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) SetCode(v string) {
	o.Code = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) GetLinks() string {
	if o == nil || IsNil(o.Links) {
		var ret string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) GetLinksOk() (*string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given string and assigns it to the Links field.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) SetLinks(v string) {
	o.Links = &v
}

// GetMessage returns the Message field value
func (o *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) SetMessage(v string) {
	o.Message = v
}

func (o AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) UnmarshalJSON(data []byte) (err error) {
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

	varAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel := _AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel)

	if err != nil {
		return err
	}

	*o = AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel(varAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel)

	return err
}

type NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel struct {
	value *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel
	isSet bool
}

func (v NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) Get() *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel {
	return v.value
}

func (v *NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) Set(val *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel(val *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) *NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel {
	return &NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateErrorResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
