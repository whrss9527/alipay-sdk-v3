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

// checks if the AlipayCommerceIotDeviceReportUploadResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceIotDeviceReportUploadResponseModel{}

// AlipayCommerceIotDeviceReportUploadResponseModel struct for AlipayCommerceIotDeviceReportUploadResponseModel
type AlipayCommerceIotDeviceReportUploadResponseModel struct {
	// 是否上传成功
	Status *bool `json:"status,omitempty"`
}

// NewAlipayCommerceIotDeviceReportUploadResponseModel instantiates a new AlipayCommerceIotDeviceReportUploadResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceIotDeviceReportUploadResponseModel() *AlipayCommerceIotDeviceReportUploadResponseModel {
	this := AlipayCommerceIotDeviceReportUploadResponseModel{}
	return &this
}

// NewAlipayCommerceIotDeviceReportUploadResponseModelWithDefaults instantiates a new AlipayCommerceIotDeviceReportUploadResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceIotDeviceReportUploadResponseModelWithDefaults() *AlipayCommerceIotDeviceReportUploadResponseModel {
	this := AlipayCommerceIotDeviceReportUploadResponseModel{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayCommerceIotDeviceReportUploadResponseModel) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceIotDeviceReportUploadResponseModel) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayCommerceIotDeviceReportUploadResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *AlipayCommerceIotDeviceReportUploadResponseModel) SetStatus(v bool) {
	o.Status = &v
}

func (o AlipayCommerceIotDeviceReportUploadResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceIotDeviceReportUploadResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAlipayCommerceIotDeviceReportUploadResponseModel struct {
	value *AlipayCommerceIotDeviceReportUploadResponseModel
	isSet bool
}

func (v NullableAlipayCommerceIotDeviceReportUploadResponseModel) Get() *AlipayCommerceIotDeviceReportUploadResponseModel {
	return v.value
}

func (v *NullableAlipayCommerceIotDeviceReportUploadResponseModel) Set(val *AlipayCommerceIotDeviceReportUploadResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceIotDeviceReportUploadResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceIotDeviceReportUploadResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceIotDeviceReportUploadResponseModel(val *AlipayCommerceIotDeviceReportUploadResponseModel) *NullableAlipayCommerceIotDeviceReportUploadResponseModel {
	return &NullableAlipayCommerceIotDeviceReportUploadResponseModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceIotDeviceReportUploadResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceIotDeviceReportUploadResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
