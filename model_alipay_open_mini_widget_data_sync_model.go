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

// checks if the AlipayOpenMiniWidgetDataSyncModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniWidgetDataSyncModel{}

// AlipayOpenMiniWidgetDataSyncModel struct for AlipayOpenMiniWidgetDataSyncModel
type AlipayOpenMiniWidgetDataSyncModel struct {
	// 活动信息列表
	ActivityList []WidgetActivityInfo `json:"activity_list,omitempty"`
	// 数据类型，可选值：ACTIVITY(活动数据)、GOODS(商品数据)
	DataType *string `json:"data_type,omitempty"`
	// 商品信息列表
	GoodsList []WidgetGoodsInfo `json:"goods_list,omitempty"`
	// 商家小程序ID
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// 小程序背后的商家，和mini_app_id要求对应
	Pid *string `json:"pid,omitempty"`
}

// NewAlipayOpenMiniWidgetDataSyncModel instantiates a new AlipayOpenMiniWidgetDataSyncModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniWidgetDataSyncModel() *AlipayOpenMiniWidgetDataSyncModel {
	this := AlipayOpenMiniWidgetDataSyncModel{}
	return &this
}

// NewAlipayOpenMiniWidgetDataSyncModelWithDefaults instantiates a new AlipayOpenMiniWidgetDataSyncModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniWidgetDataSyncModelWithDefaults() *AlipayOpenMiniWidgetDataSyncModel {
	this := AlipayOpenMiniWidgetDataSyncModel{}
	return &this
}

// GetActivityList returns the ActivityList field value if set, zero value otherwise.
func (o *AlipayOpenMiniWidgetDataSyncModel) GetActivityList() []WidgetActivityInfo {
	if o == nil || IsNil(o.ActivityList) {
		var ret []WidgetActivityInfo
		return ret
	}
	return o.ActivityList
}

// GetActivityListOk returns a tuple with the ActivityList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniWidgetDataSyncModel) GetActivityListOk() ([]WidgetActivityInfo, bool) {
	if o == nil || IsNil(o.ActivityList) {
		return nil, false
	}
	return o.ActivityList, true
}

// HasActivityList returns a boolean if a field has been set.
func (o *AlipayOpenMiniWidgetDataSyncModel) HasActivityList() bool {
	if o != nil && !IsNil(o.ActivityList) {
		return true
	}

	return false
}

// SetActivityList gets a reference to the given []WidgetActivityInfo and assigns it to the ActivityList field.
func (o *AlipayOpenMiniWidgetDataSyncModel) SetActivityList(v []WidgetActivityInfo) {
	o.ActivityList = v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *AlipayOpenMiniWidgetDataSyncModel) GetDataType() string {
	if o == nil || IsNil(o.DataType) {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniWidgetDataSyncModel) GetDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *AlipayOpenMiniWidgetDataSyncModel) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *AlipayOpenMiniWidgetDataSyncModel) SetDataType(v string) {
	o.DataType = &v
}

// GetGoodsList returns the GoodsList field value if set, zero value otherwise.
func (o *AlipayOpenMiniWidgetDataSyncModel) GetGoodsList() []WidgetGoodsInfo {
	if o == nil || IsNil(o.GoodsList) {
		var ret []WidgetGoodsInfo
		return ret
	}
	return o.GoodsList
}

// GetGoodsListOk returns a tuple with the GoodsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniWidgetDataSyncModel) GetGoodsListOk() ([]WidgetGoodsInfo, bool) {
	if o == nil || IsNil(o.GoodsList) {
		return nil, false
	}
	return o.GoodsList, true
}

// HasGoodsList returns a boolean if a field has been set.
func (o *AlipayOpenMiniWidgetDataSyncModel) HasGoodsList() bool {
	if o != nil && !IsNil(o.GoodsList) {
		return true
	}

	return false
}

// SetGoodsList gets a reference to the given []WidgetGoodsInfo and assigns it to the GoodsList field.
func (o *AlipayOpenMiniWidgetDataSyncModel) SetGoodsList(v []WidgetGoodsInfo) {
	o.GoodsList = v
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *AlipayOpenMiniWidgetDataSyncModel) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniWidgetDataSyncModel) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *AlipayOpenMiniWidgetDataSyncModel) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *AlipayOpenMiniWidgetDataSyncModel) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *AlipayOpenMiniWidgetDataSyncModel) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniWidgetDataSyncModel) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *AlipayOpenMiniWidgetDataSyncModel) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *AlipayOpenMiniWidgetDataSyncModel) SetPid(v string) {
	o.Pid = &v
}

func (o AlipayOpenMiniWidgetDataSyncModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniWidgetDataSyncModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityList) {
		toSerialize["activity_list"] = o.ActivityList
	}
	if !IsNil(o.DataType) {
		toSerialize["data_type"] = o.DataType
	}
	if !IsNil(o.GoodsList) {
		toSerialize["goods_list"] = o.GoodsList
	}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniWidgetDataSyncModel struct {
	value *AlipayOpenMiniWidgetDataSyncModel
	isSet bool
}

func (v NullableAlipayOpenMiniWidgetDataSyncModel) Get() *AlipayOpenMiniWidgetDataSyncModel {
	return v.value
}

func (v *NullableAlipayOpenMiniWidgetDataSyncModel) Set(val *AlipayOpenMiniWidgetDataSyncModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniWidgetDataSyncModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniWidgetDataSyncModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniWidgetDataSyncModel(val *AlipayOpenMiniWidgetDataSyncModel) *NullableAlipayOpenMiniWidgetDataSyncModel {
	return &NullableAlipayOpenMiniWidgetDataSyncModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniWidgetDataSyncModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniWidgetDataSyncModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
