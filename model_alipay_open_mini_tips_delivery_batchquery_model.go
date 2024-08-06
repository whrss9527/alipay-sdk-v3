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

// checks if the AlipayOpenMiniTipsDeliveryBatchqueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniTipsDeliveryBatchqueryModel{}

// AlipayOpenMiniTipsDeliveryBatchqueryModel struct for AlipayOpenMiniTipsDeliveryBatchqueryModel
type AlipayOpenMiniTipsDeliveryBatchqueryModel struct {
	// 每页记录条数，最小1，最大20
	PageNumber *string `json:"page_number,omitempty"`
	// 查询的页数，从1开始，最大1000
	PageSize *string `json:"page_size,omitempty"`
}

// NewAlipayOpenMiniTipsDeliveryBatchqueryModel instantiates a new AlipayOpenMiniTipsDeliveryBatchqueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniTipsDeliveryBatchqueryModel() *AlipayOpenMiniTipsDeliveryBatchqueryModel {
	this := AlipayOpenMiniTipsDeliveryBatchqueryModel{}
	return &this
}

// NewAlipayOpenMiniTipsDeliveryBatchqueryModelWithDefaults instantiates a new AlipayOpenMiniTipsDeliveryBatchqueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniTipsDeliveryBatchqueryModelWithDefaults() *AlipayOpenMiniTipsDeliveryBatchqueryModel {
	this := AlipayOpenMiniTipsDeliveryBatchqueryModel{}
	return &this
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *AlipayOpenMiniTipsDeliveryBatchqueryModel) GetPageNumber() string {
	if o == nil || IsNil(o.PageNumber) {
		var ret string
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTipsDeliveryBatchqueryModel) GetPageNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *AlipayOpenMiniTipsDeliveryBatchqueryModel) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given string and assigns it to the PageNumber field.
func (o *AlipayOpenMiniTipsDeliveryBatchqueryModel) SetPageNumber(v string) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayOpenMiniTipsDeliveryBatchqueryModel) GetPageSize() string {
	if o == nil || IsNil(o.PageSize) {
		var ret string
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTipsDeliveryBatchqueryModel) GetPageSizeOk() (*string, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayOpenMiniTipsDeliveryBatchqueryModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given string and assigns it to the PageSize field.
func (o *AlipayOpenMiniTipsDeliveryBatchqueryModel) SetPageSize(v string) {
	o.PageSize = &v
}

func (o AlipayOpenMiniTipsDeliveryBatchqueryModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniTipsDeliveryBatchqueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageNumber) {
		toSerialize["page_number"] = o.PageNumber
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniTipsDeliveryBatchqueryModel struct {
	value *AlipayOpenMiniTipsDeliveryBatchqueryModel
	isSet bool
}

func (v NullableAlipayOpenMiniTipsDeliveryBatchqueryModel) Get() *AlipayOpenMiniTipsDeliveryBatchqueryModel {
	return v.value
}

func (v *NullableAlipayOpenMiniTipsDeliveryBatchqueryModel) Set(val *AlipayOpenMiniTipsDeliveryBatchqueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniTipsDeliveryBatchqueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniTipsDeliveryBatchqueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniTipsDeliveryBatchqueryModel(val *AlipayOpenMiniTipsDeliveryBatchqueryModel) *NullableAlipayOpenMiniTipsDeliveryBatchqueryModel {
	return &NullableAlipayOpenMiniTipsDeliveryBatchqueryModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniTipsDeliveryBatchqueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniTipsDeliveryBatchqueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

