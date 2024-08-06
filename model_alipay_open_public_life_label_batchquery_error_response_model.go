/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel{}

// AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel struct for AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel
type AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel struct {
	// 错误码
	Code string `json:"code"`
	// 解决方案链接
	Links *string `json:"links,omitempty"`
	// 错误描述
	Message string `json:"message"`
}

type _AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel

// NewAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel instantiates a new AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel(code string, message string) *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel {
	this := AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAlipayOpenPublicLifeLabelBatchqueryErrorResponseModelWithDefaults instantiates a new AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicLifeLabelBatchqueryErrorResponseModelWithDefaults() *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel {
	this := AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel{}
	return &this
}

// GetCode returns the Code field value
func (o *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) SetCode(v string) {
	o.Code = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) GetLinks() string {
	if o == nil || IsNil(o.Links) {
		var ret string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) GetLinksOk() (*string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given string and assigns it to the Links field.
func (o *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) SetLinks(v string) {
	o.Links = &v
}

// GetMessage returns the Message field value
func (o *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) SetMessage(v string) {
	o.Message = v
}

func (o AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) UnmarshalJSON(data []byte) (err error) {
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
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel := _AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel)

	if err != nil {
		return err
	}

	*o = AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel(varAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel)

	return err
}

type NullableAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel struct {
	value *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel
	isSet bool
}

func (v NullableAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) Get() *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel {
	return v.value
}

func (v *NullableAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) Set(val *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel(val *AlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) *NullableAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel {
	return &NullableAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicLifeLabelBatchqueryErrorResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

