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

// checks if the AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel{}

// AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel struct for AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel
type AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel struct {
	// 错误码
	Code string `json:"code"`
	// 解决方案链接
	Links *string `json:"links,omitempty"`
	// 错误描述
	Message string `json:"message"`
}

type _AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel

// NewAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel instantiates a new AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel(code string, message string) *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel {
	this := AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModelWithDefaults instantiates a new AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModelWithDefaults() *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel {
	this := AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel{}
	return &this
}

// GetCode returns the Code field value
func (o *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) SetCode(v string) {
	o.Code = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) GetLinks() string {
	if o == nil || IsNil(o.Links) {
		var ret string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) GetLinksOk() (*string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given string and assigns it to the Links field.
func (o *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) SetLinks(v string) {
	o.Links = &v
}

// GetMessage returns the Message field value
func (o *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) SetMessage(v string) {
	o.Message = v
}

func (o AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) UnmarshalJSON(data []byte) (err error) {
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

	varAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel := _AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel)

	if err != nil {
		return err
	}

	*o = AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel(varAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel)

	return err
}

type NullableAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel struct {
	value *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) Get() *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) Set(val *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel(val *AlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) *NullableAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel {
	return &NullableAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionModelforofflineQueryErrorResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
