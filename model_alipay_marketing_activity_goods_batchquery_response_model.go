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

// checks if the AlipayMarketingActivityGoodsBatchqueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityGoodsBatchqueryResponseModel{}

// AlipayMarketingActivityGoodsBatchqueryResponseModel struct for AlipayMarketingActivityGoodsBatchqueryResponseModel
type AlipayMarketingActivityGoodsBatchqueryResponseModel struct {
	// 活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// 小程序商品信息
	AppItemInfos []AppItemInfo `json:"app_item_infos,omitempty"`
	// 商品编码列表
	GoodsInfos []ActivityGoodsInfo `json:"goods_infos,omitempty"`
	// 分页查询页码。
	PageNum *int32 `json:"page_num,omitempty"`
	// 分页查询单页数据条数。
	PageSize *int32 `json:"page_size,omitempty"`
	// 商品编码总数量
	TotalSize *int32 `json:"total_size,omitempty"`
}

// NewAlipayMarketingActivityGoodsBatchqueryResponseModel instantiates a new AlipayMarketingActivityGoodsBatchqueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityGoodsBatchqueryResponseModel() *AlipayMarketingActivityGoodsBatchqueryResponseModel {
	this := AlipayMarketingActivityGoodsBatchqueryResponseModel{}
	return &this
}

// NewAlipayMarketingActivityGoodsBatchqueryResponseModelWithDefaults instantiates a new AlipayMarketingActivityGoodsBatchqueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityGoodsBatchqueryResponseModelWithDefaults() *AlipayMarketingActivityGoodsBatchqueryResponseModel {
	this := AlipayMarketingActivityGoodsBatchqueryResponseModel{}
	return &this
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId) {
		var ret string
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) GetActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given string and assigns it to the ActivityId field.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) SetActivityId(v string) {
	o.ActivityId = &v
}

// GetAppItemInfos returns the AppItemInfos field value if set, zero value otherwise.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) GetAppItemInfos() []AppItemInfo {
	if o == nil || IsNil(o.AppItemInfos) {
		var ret []AppItemInfo
		return ret
	}
	return o.AppItemInfos
}

// GetAppItemInfosOk returns a tuple with the AppItemInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) GetAppItemInfosOk() ([]AppItemInfo, bool) {
	if o == nil || IsNil(o.AppItemInfos) {
		return nil, false
	}
	return o.AppItemInfos, true
}

// HasAppItemInfos returns a boolean if a field has been set.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) HasAppItemInfos() bool {
	if o != nil && !IsNil(o.AppItemInfos) {
		return true
	}

	return false
}

// SetAppItemInfos gets a reference to the given []AppItemInfo and assigns it to the AppItemInfos field.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) SetAppItemInfos(v []AppItemInfo) {
	o.AppItemInfos = v
}

// GetGoodsInfos returns the GoodsInfos field value if set, zero value otherwise.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) GetGoodsInfos() []ActivityGoodsInfo {
	if o == nil || IsNil(o.GoodsInfos) {
		var ret []ActivityGoodsInfo
		return ret
	}
	return o.GoodsInfos
}

// GetGoodsInfosOk returns a tuple with the GoodsInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) GetGoodsInfosOk() ([]ActivityGoodsInfo, bool) {
	if o == nil || IsNil(o.GoodsInfos) {
		return nil, false
	}
	return o.GoodsInfos, true
}

// HasGoodsInfos returns a boolean if a field has been set.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) HasGoodsInfos() bool {
	if o != nil && !IsNil(o.GoodsInfos) {
		return true
	}

	return false
}

// SetGoodsInfos gets a reference to the given []ActivityGoodsInfo and assigns it to the GoodsInfos field.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) SetGoodsInfos(v []ActivityGoodsInfo) {
	o.GoodsInfos = v
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) GetPageNum() int32 {
	if o == nil || IsNil(o.PageNum) {
		var ret int32
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) GetPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int32 and assigns it to the PageNum field.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) SetPageNum(v int32) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) GetTotalSize() int32 {
	if o == nil || IsNil(o.TotalSize) {
		var ret int32
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) GetTotalSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int32 and assigns it to the TotalSize field.
func (o *AlipayMarketingActivityGoodsBatchqueryResponseModel) SetTotalSize(v int32) {
	o.TotalSize = &v
}

func (o AlipayMarketingActivityGoodsBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityGoodsBatchqueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
	}
	if !IsNil(o.AppItemInfos) {
		toSerialize["app_item_infos"] = o.AppItemInfos
	}
	if !IsNil(o.GoodsInfos) {
		toSerialize["goods_infos"] = o.GoodsInfos
	}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.TotalSize) {
		toSerialize["total_size"] = o.TotalSize
	}
	return toSerialize, nil
}

type NullableAlipayMarketingActivityGoodsBatchqueryResponseModel struct {
	value *AlipayMarketingActivityGoodsBatchqueryResponseModel
	isSet bool
}

func (v NullableAlipayMarketingActivityGoodsBatchqueryResponseModel) Get() *AlipayMarketingActivityGoodsBatchqueryResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityGoodsBatchqueryResponseModel) Set(val *AlipayMarketingActivityGoodsBatchqueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityGoodsBatchqueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityGoodsBatchqueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityGoodsBatchqueryResponseModel(val *AlipayMarketingActivityGoodsBatchqueryResponseModel) *NullableAlipayMarketingActivityGoodsBatchqueryResponseModel {
	return &NullableAlipayMarketingActivityGoodsBatchqueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityGoodsBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityGoodsBatchqueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
