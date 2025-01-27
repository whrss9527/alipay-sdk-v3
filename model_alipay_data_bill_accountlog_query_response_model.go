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

// checks if the AlipayDataBillAccountlogQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayDataBillAccountlogQueryResponseModel{}

// AlipayDataBillAccountlogQueryResponseModel struct for AlipayDataBillAccountlogQueryResponseModel
type AlipayDataBillAccountlogQueryResponseModel struct {
	// 账务明细返回结果。 当查询为空的时候，不返回，有结果的时候才会返回。
	DetailList []AccountLogItemResult `json:"detail_list,omitempty"`
	// 分页号，从1开始
	PageNo *string `json:"page_no,omitempty"`
	// 分页大小1000-2000
	PageSize *string `json:"page_size,omitempty"`
	// 账务明细总数。返回满足查询条件的明细的数量
	TotalSize *string `json:"total_size,omitempty"`
}

// NewAlipayDataBillAccountlogQueryResponseModel instantiates a new AlipayDataBillAccountlogQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayDataBillAccountlogQueryResponseModel() *AlipayDataBillAccountlogQueryResponseModel {
	this := AlipayDataBillAccountlogQueryResponseModel{}
	return &this
}

// NewAlipayDataBillAccountlogQueryResponseModelWithDefaults instantiates a new AlipayDataBillAccountlogQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayDataBillAccountlogQueryResponseModelWithDefaults() *AlipayDataBillAccountlogQueryResponseModel {
	this := AlipayDataBillAccountlogQueryResponseModel{}
	return &this
}

// GetDetailList returns the DetailList field value if set, zero value otherwise.
func (o *AlipayDataBillAccountlogQueryResponseModel) GetDetailList() []AccountLogItemResult {
	if o == nil || IsNil(o.DetailList) {
		var ret []AccountLogItemResult
		return ret
	}
	return o.DetailList
}

// GetDetailListOk returns a tuple with the DetailList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayDataBillAccountlogQueryResponseModel) GetDetailListOk() ([]AccountLogItemResult, bool) {
	if o == nil || IsNil(o.DetailList) {
		return nil, false
	}
	return o.DetailList, true
}

// HasDetailList returns a boolean if a field has been set.
func (o *AlipayDataBillAccountlogQueryResponseModel) HasDetailList() bool {
	if o != nil && !IsNil(o.DetailList) {
		return true
	}

	return false
}

// SetDetailList gets a reference to the given []AccountLogItemResult and assigns it to the DetailList field.
func (o *AlipayDataBillAccountlogQueryResponseModel) SetDetailList(v []AccountLogItemResult) {
	o.DetailList = v
}

// GetPageNo returns the PageNo field value if set, zero value otherwise.
func (o *AlipayDataBillAccountlogQueryResponseModel) GetPageNo() string {
	if o == nil || IsNil(o.PageNo) {
		var ret string
		return ret
	}
	return *o.PageNo
}

// GetPageNoOk returns a tuple with the PageNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayDataBillAccountlogQueryResponseModel) GetPageNoOk() (*string, bool) {
	if o == nil || IsNil(o.PageNo) {
		return nil, false
	}
	return o.PageNo, true
}

// HasPageNo returns a boolean if a field has been set.
func (o *AlipayDataBillAccountlogQueryResponseModel) HasPageNo() bool {
	if o != nil && !IsNil(o.PageNo) {
		return true
	}

	return false
}

// SetPageNo gets a reference to the given string and assigns it to the PageNo field.
func (o *AlipayDataBillAccountlogQueryResponseModel) SetPageNo(v string) {
	o.PageNo = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayDataBillAccountlogQueryResponseModel) GetPageSize() string {
	if o == nil || IsNil(o.PageSize) {
		var ret string
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayDataBillAccountlogQueryResponseModel) GetPageSizeOk() (*string, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayDataBillAccountlogQueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given string and assigns it to the PageSize field.
func (o *AlipayDataBillAccountlogQueryResponseModel) SetPageSize(v string) {
	o.PageSize = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *AlipayDataBillAccountlogQueryResponseModel) GetTotalSize() string {
	if o == nil || IsNil(o.TotalSize) {
		var ret string
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayDataBillAccountlogQueryResponseModel) GetTotalSizeOk() (*string, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *AlipayDataBillAccountlogQueryResponseModel) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given string and assigns it to the TotalSize field.
func (o *AlipayDataBillAccountlogQueryResponseModel) SetTotalSize(v string) {
	o.TotalSize = &v
}

func (o AlipayDataBillAccountlogQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayDataBillAccountlogQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DetailList) {
		toSerialize["detail_list"] = o.DetailList
	}
	if !IsNil(o.PageNo) {
		toSerialize["page_no"] = o.PageNo
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.TotalSize) {
		toSerialize["total_size"] = o.TotalSize
	}
	return toSerialize, nil
}

type NullableAlipayDataBillAccountlogQueryResponseModel struct {
	value *AlipayDataBillAccountlogQueryResponseModel
	isSet bool
}

func (v NullableAlipayDataBillAccountlogQueryResponseModel) Get() *AlipayDataBillAccountlogQueryResponseModel {
	return v.value
}

func (v *NullableAlipayDataBillAccountlogQueryResponseModel) Set(val *AlipayDataBillAccountlogQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayDataBillAccountlogQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayDataBillAccountlogQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayDataBillAccountlogQueryResponseModel(val *AlipayDataBillAccountlogQueryResponseModel) *NullableAlipayDataBillAccountlogQueryResponseModel {
	return &NullableAlipayDataBillAccountlogQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayDataBillAccountlogQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayDataBillAccountlogQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
