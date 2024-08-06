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

// checks if the AlipayOpenMiniWidgetGoodsQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniWidgetGoodsQueryResponseModel{}

// AlipayOpenMiniWidgetGoodsQueryResponseModel struct for AlipayOpenMiniWidgetGoodsQueryResponseModel
type AlipayOpenMiniWidgetGoodsQueryResponseModel struct {
	// 商品信息列表
	DataList []GoodsQueryResponse `json:"data_list,omitempty"`
	// 查询第几页
	PageNum *int32 `json:"page_num,omitempty"`
	// 查询页面数量
	PageSize *int32 `json:"page_size,omitempty"`
	// 查询结果总数
	Total *int32 `json:"total,omitempty"`
}

// NewAlipayOpenMiniWidgetGoodsQueryResponseModel instantiates a new AlipayOpenMiniWidgetGoodsQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniWidgetGoodsQueryResponseModel() *AlipayOpenMiniWidgetGoodsQueryResponseModel {
	this := AlipayOpenMiniWidgetGoodsQueryResponseModel{}
	return &this
}

// NewAlipayOpenMiniWidgetGoodsQueryResponseModelWithDefaults instantiates a new AlipayOpenMiniWidgetGoodsQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniWidgetGoodsQueryResponseModelWithDefaults() *AlipayOpenMiniWidgetGoodsQueryResponseModel {
	this := AlipayOpenMiniWidgetGoodsQueryResponseModel{}
	return &this
}

// GetDataList returns the DataList field value if set, zero value otherwise.
func (o *AlipayOpenMiniWidgetGoodsQueryResponseModel) GetDataList() []GoodsQueryResponse {
	if o == nil || IsNil(o.DataList) {
		var ret []GoodsQueryResponse
		return ret
	}
	return o.DataList
}

// GetDataListOk returns a tuple with the DataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniWidgetGoodsQueryResponseModel) GetDataListOk() ([]GoodsQueryResponse, bool) {
	if o == nil || IsNil(o.DataList) {
		return nil, false
	}
	return o.DataList, true
}

// HasDataList returns a boolean if a field has been set.
func (o *AlipayOpenMiniWidgetGoodsQueryResponseModel) HasDataList() bool {
	if o != nil && !IsNil(o.DataList) {
		return true
	}

	return false
}

// SetDataList gets a reference to the given []GoodsQueryResponse and assigns it to the DataList field.
func (o *AlipayOpenMiniWidgetGoodsQueryResponseModel) SetDataList(v []GoodsQueryResponse) {
	o.DataList = v
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayOpenMiniWidgetGoodsQueryResponseModel) GetPageNum() int32 {
	if o == nil || IsNil(o.PageNum) {
		var ret int32
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniWidgetGoodsQueryResponseModel) GetPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayOpenMiniWidgetGoodsQueryResponseModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int32 and assigns it to the PageNum field.
func (o *AlipayOpenMiniWidgetGoodsQueryResponseModel) SetPageNum(v int32) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayOpenMiniWidgetGoodsQueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniWidgetGoodsQueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayOpenMiniWidgetGoodsQueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayOpenMiniWidgetGoodsQueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *AlipayOpenMiniWidgetGoodsQueryResponseModel) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniWidgetGoodsQueryResponseModel) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *AlipayOpenMiniWidgetGoodsQueryResponseModel) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *AlipayOpenMiniWidgetGoodsQueryResponseModel) SetTotal(v int32) {
	o.Total = &v
}

func (o AlipayOpenMiniWidgetGoodsQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniWidgetGoodsQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataList) {
		toSerialize["data_list"] = o.DataList
	}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniWidgetGoodsQueryResponseModel struct {
	value *AlipayOpenMiniWidgetGoodsQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniWidgetGoodsQueryResponseModel) Get() *AlipayOpenMiniWidgetGoodsQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniWidgetGoodsQueryResponseModel) Set(val *AlipayOpenMiniWidgetGoodsQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniWidgetGoodsQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniWidgetGoodsQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniWidgetGoodsQueryResponseModel(val *AlipayOpenMiniWidgetGoodsQueryResponseModel) *NullableAlipayOpenMiniWidgetGoodsQueryResponseModel {
	return &NullableAlipayOpenMiniWidgetGoodsQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniWidgetGoodsQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniWidgetGoodsQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

