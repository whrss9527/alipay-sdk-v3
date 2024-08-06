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

// checks if the AlipayOpenSearchBoxBatchqueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSearchBoxBatchqueryModel{}

// AlipayOpenSearchBoxBatchqueryModel struct for AlipayOpenSearchBoxBatchqueryModel
type AlipayOpenSearchBoxBatchqueryModel struct {
	// 商户id，代运营模式下传入。代运营模式，需要服务商已获得商家\"运营支付宝小程序\"授权。
	MerchantId *string `json:"merchant_id,omitempty"`
	// 分页查询的当前页号,从1开始
	PageNumber *int32 `json:"page_number,omitempty"`
	// 每页查询的数量，默认10，不得超过50
	PageSize *int32 `json:"page_size,omitempty"`
}

// NewAlipayOpenSearchBoxBatchqueryModel instantiates a new AlipayOpenSearchBoxBatchqueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSearchBoxBatchqueryModel() *AlipayOpenSearchBoxBatchqueryModel {
	this := AlipayOpenSearchBoxBatchqueryModel{}
	return &this
}

// NewAlipayOpenSearchBoxBatchqueryModelWithDefaults instantiates a new AlipayOpenSearchBoxBatchqueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSearchBoxBatchqueryModelWithDefaults() *AlipayOpenSearchBoxBatchqueryModel {
	this := AlipayOpenSearchBoxBatchqueryModel{}
	return &this
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxBatchqueryModel) GetMerchantId() string {
	if o == nil || IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxBatchqueryModel) GetMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxBatchqueryModel) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *AlipayOpenSearchBoxBatchqueryModel) SetMerchantId(v string) {
	o.MerchantId = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxBatchqueryModel) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxBatchqueryModel) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxBatchqueryModel) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *AlipayOpenSearchBoxBatchqueryModel) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxBatchqueryModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxBatchqueryModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxBatchqueryModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayOpenSearchBoxBatchqueryModel) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o AlipayOpenSearchBoxBatchqueryModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSearchBoxBatchqueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantId) {
		toSerialize["merchant_id"] = o.MerchantId
	}
	if !IsNil(o.PageNumber) {
		toSerialize["page_number"] = o.PageNumber
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	return toSerialize, nil
}

type NullableAlipayOpenSearchBoxBatchqueryModel struct {
	value *AlipayOpenSearchBoxBatchqueryModel
	isSet bool
}

func (v NullableAlipayOpenSearchBoxBatchqueryModel) Get() *AlipayOpenSearchBoxBatchqueryModel {
	return v.value
}

func (v *NullableAlipayOpenSearchBoxBatchqueryModel) Set(val *AlipayOpenSearchBoxBatchqueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchBoxBatchqueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchBoxBatchqueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchBoxBatchqueryModel(val *AlipayOpenSearchBoxBatchqueryModel) *NullableAlipayOpenSearchBoxBatchqueryModel {
	return &NullableAlipayOpenSearchBoxBatchqueryModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchBoxBatchqueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchBoxBatchqueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

