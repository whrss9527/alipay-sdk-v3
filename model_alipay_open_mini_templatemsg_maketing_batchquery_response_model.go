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

// checks if the AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel{}

// AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel struct for AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel
type AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel struct {
	// 投放详情列表
	MarketingDeliveryDetailList []MarketingDeliveryDetail `json:"marketing_delivery_detail_list,omitempty"`
	// 投放详情总条数
	Total *int32 `json:"total,omitempty"`
}

// NewAlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel instantiates a new AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel() *AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel {
	this := AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel{}
	return &this
}

// NewAlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModelWithDefaults instantiates a new AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModelWithDefaults() *AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel {
	this := AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel{}
	return &this
}

// GetMarketingDeliveryDetailList returns the MarketingDeliveryDetailList field value if set, zero value otherwise.
func (o *AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) GetMarketingDeliveryDetailList() []MarketingDeliveryDetail {
	if o == nil || IsNil(o.MarketingDeliveryDetailList) {
		var ret []MarketingDeliveryDetail
		return ret
	}
	return o.MarketingDeliveryDetailList
}

// GetMarketingDeliveryDetailListOk returns a tuple with the MarketingDeliveryDetailList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) GetMarketingDeliveryDetailListOk() ([]MarketingDeliveryDetail, bool) {
	if o == nil || IsNil(o.MarketingDeliveryDetailList) {
		return nil, false
	}
	return o.MarketingDeliveryDetailList, true
}

// HasMarketingDeliveryDetailList returns a boolean if a field has been set.
func (o *AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) HasMarketingDeliveryDetailList() bool {
	if o != nil && !IsNil(o.MarketingDeliveryDetailList) {
		return true
	}

	return false
}

// SetMarketingDeliveryDetailList gets a reference to the given []MarketingDeliveryDetail and assigns it to the MarketingDeliveryDetailList field.
func (o *AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) SetMarketingDeliveryDetailList(v []MarketingDeliveryDetail) {
	o.MarketingDeliveryDetailList = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) SetTotal(v int32) {
	o.Total = &v
}

func (o AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketingDeliveryDetailList) {
		toSerialize["marketing_delivery_detail_list"] = o.MarketingDeliveryDetailList
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel struct {
	value *AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) Get() *AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) Set(val *AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel(val *AlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) *NullableAlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel {
	return &NullableAlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniTemplatemsgMaketingBatchqueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
