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

// checks if the AlipayMarketingActivityVoucherpackageBatchqueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityVoucherpackageBatchqueryResponseModel{}

// AlipayMarketingActivityVoucherpackageBatchqueryResponseModel struct for AlipayMarketingActivityVoucherpackageBatchqueryResponseModel
type AlipayMarketingActivityVoucherpackageBatchqueryResponseModel struct {
	// 分页参数，当前所在的页码
	PageNum *int32 `json:"page_num,omitempty"`
	// 分页参数，每页记录数，最大不可超过30
	PageSize *int32 `json:"page_size,omitempty"`
	// 券包总数量
	TotalSize *int32 `json:"total_size,omitempty"`
	// 券包信息
	VoucherPackageInfo []VoucherPackageInfo `json:"voucher_package_info,omitempty"`
}

// NewAlipayMarketingActivityVoucherpackageBatchqueryResponseModel instantiates a new AlipayMarketingActivityVoucherpackageBatchqueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityVoucherpackageBatchqueryResponseModel() *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel {
	this := AlipayMarketingActivityVoucherpackageBatchqueryResponseModel{}
	return &this
}

// NewAlipayMarketingActivityVoucherpackageBatchqueryResponseModelWithDefaults instantiates a new AlipayMarketingActivityVoucherpackageBatchqueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityVoucherpackageBatchqueryResponseModelWithDefaults() *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel {
	this := AlipayMarketingActivityVoucherpackageBatchqueryResponseModel{}
	return &this
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) GetPageNum() int32 {
	if o == nil || IsNil(o.PageNum) {
		var ret int32
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) GetPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int32 and assigns it to the PageNum field.
func (o *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) SetPageNum(v int32) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) GetTotalSize() int32 {
	if o == nil || IsNil(o.TotalSize) {
		var ret int32
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) GetTotalSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int32 and assigns it to the TotalSize field.
func (o *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) SetTotalSize(v int32) {
	o.TotalSize = &v
}

// GetVoucherPackageInfo returns the VoucherPackageInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) GetVoucherPackageInfo() []VoucherPackageInfo {
	if o == nil || IsNil(o.VoucherPackageInfo) {
		var ret []VoucherPackageInfo
		return ret
	}
	return o.VoucherPackageInfo
}

// GetVoucherPackageInfoOk returns a tuple with the VoucherPackageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) GetVoucherPackageInfoOk() ([]VoucherPackageInfo, bool) {
	if o == nil || IsNil(o.VoucherPackageInfo) {
		return nil, false
	}
	return o.VoucherPackageInfo, true
}

// HasVoucherPackageInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) HasVoucherPackageInfo() bool {
	if o != nil && !IsNil(o.VoucherPackageInfo) {
		return true
	}

	return false
}

// SetVoucherPackageInfo gets a reference to the given []VoucherPackageInfo and assigns it to the VoucherPackageInfo field.
func (o *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) SetVoucherPackageInfo(v []VoucherPackageInfo) {
	o.VoucherPackageInfo = v
}

func (o AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.TotalSize) {
		toSerialize["total_size"] = o.TotalSize
	}
	if !IsNil(o.VoucherPackageInfo) {
		toSerialize["voucher_package_info"] = o.VoucherPackageInfo
	}
	return toSerialize, nil
}

type NullableAlipayMarketingActivityVoucherpackageBatchqueryResponseModel struct {
	value *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel
	isSet bool
}

func (v NullableAlipayMarketingActivityVoucherpackageBatchqueryResponseModel) Get() *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityVoucherpackageBatchqueryResponseModel) Set(val *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityVoucherpackageBatchqueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityVoucherpackageBatchqueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityVoucherpackageBatchqueryResponseModel(val *AlipayMarketingActivityVoucherpackageBatchqueryResponseModel) *NullableAlipayMarketingActivityVoucherpackageBatchqueryResponseModel {
	return &NullableAlipayMarketingActivityVoucherpackageBatchqueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityVoucherpackageBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityVoucherpackageBatchqueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

