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

// checks if the AlipayMerchantTradecomplainBatchqueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMerchantTradecomplainBatchqueryResponseModel{}

// AlipayMerchantTradecomplainBatchqueryResponseModel struct for AlipayMerchantTradecomplainBatchqueryResponseModel
type AlipayMerchantTradecomplainBatchqueryResponseModel struct {
	// 当前页码数
	PageNum *int32 `json:"page_num,omitempty"`
	// 每页记录数
	PageSize *int32 `json:"page_size,omitempty"`
	// 总条数
	TotalNum *int32 `json:"total_num,omitempty"`
	// 总页码数
	TotalPageNum *int32 `json:"total_page_num,omitempty"`
	// 交易纠纷工单列表信息
	TradeComplainInfos []TradeComplainQueryResponse `json:"trade_complain_infos,omitempty"`
}

// NewAlipayMerchantTradecomplainBatchqueryResponseModel instantiates a new AlipayMerchantTradecomplainBatchqueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMerchantTradecomplainBatchqueryResponseModel() *AlipayMerchantTradecomplainBatchqueryResponseModel {
	this := AlipayMerchantTradecomplainBatchqueryResponseModel{}
	return &this
}

// NewAlipayMerchantTradecomplainBatchqueryResponseModelWithDefaults instantiates a new AlipayMerchantTradecomplainBatchqueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMerchantTradecomplainBatchqueryResponseModelWithDefaults() *AlipayMerchantTradecomplainBatchqueryResponseModel {
	this := AlipayMerchantTradecomplainBatchqueryResponseModel{}
	return &this
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) GetPageNum() int32 {
	if o == nil || IsNil(o.PageNum) {
		var ret int32
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) GetPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int32 and assigns it to the PageNum field.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) SetPageNum(v int32) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalNum returns the TotalNum field value if set, zero value otherwise.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) GetTotalNum() int32 {
	if o == nil || IsNil(o.TotalNum) {
		var ret int32
		return ret
	}
	return *o.TotalNum
}

// GetTotalNumOk returns a tuple with the TotalNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) GetTotalNumOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalNum) {
		return nil, false
	}
	return o.TotalNum, true
}

// HasTotalNum returns a boolean if a field has been set.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) HasTotalNum() bool {
	if o != nil && !IsNil(o.TotalNum) {
		return true
	}

	return false
}

// SetTotalNum gets a reference to the given int32 and assigns it to the TotalNum field.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) SetTotalNum(v int32) {
	o.TotalNum = &v
}

// GetTotalPageNum returns the TotalPageNum field value if set, zero value otherwise.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) GetTotalPageNum() int32 {
	if o == nil || IsNil(o.TotalPageNum) {
		var ret int32
		return ret
	}
	return *o.TotalPageNum
}

// GetTotalPageNumOk returns a tuple with the TotalPageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) GetTotalPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPageNum) {
		return nil, false
	}
	return o.TotalPageNum, true
}

// HasTotalPageNum returns a boolean if a field has been set.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) HasTotalPageNum() bool {
	if o != nil && !IsNil(o.TotalPageNum) {
		return true
	}

	return false
}

// SetTotalPageNum gets a reference to the given int32 and assigns it to the TotalPageNum field.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) SetTotalPageNum(v int32) {
	o.TotalPageNum = &v
}

// GetTradeComplainInfos returns the TradeComplainInfos field value if set, zero value otherwise.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) GetTradeComplainInfos() []TradeComplainQueryResponse {
	if o == nil || IsNil(o.TradeComplainInfos) {
		var ret []TradeComplainQueryResponse
		return ret
	}
	return o.TradeComplainInfos
}

// GetTradeComplainInfosOk returns a tuple with the TradeComplainInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) GetTradeComplainInfosOk() ([]TradeComplainQueryResponse, bool) {
	if o == nil || IsNil(o.TradeComplainInfos) {
		return nil, false
	}
	return o.TradeComplainInfos, true
}

// HasTradeComplainInfos returns a boolean if a field has been set.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) HasTradeComplainInfos() bool {
	if o != nil && !IsNil(o.TradeComplainInfos) {
		return true
	}

	return false
}

// SetTradeComplainInfos gets a reference to the given []TradeComplainQueryResponse and assigns it to the TradeComplainInfos field.
func (o *AlipayMerchantTradecomplainBatchqueryResponseModel) SetTradeComplainInfos(v []TradeComplainQueryResponse) {
	o.TradeComplainInfos = v
}

func (o AlipayMerchantTradecomplainBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMerchantTradecomplainBatchqueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.TotalNum) {
		toSerialize["total_num"] = o.TotalNum
	}
	if !IsNil(o.TotalPageNum) {
		toSerialize["total_page_num"] = o.TotalPageNum
	}
	if !IsNil(o.TradeComplainInfos) {
		toSerialize["trade_complain_infos"] = o.TradeComplainInfos
	}
	return toSerialize, nil
}

type NullableAlipayMerchantTradecomplainBatchqueryResponseModel struct {
	value *AlipayMerchantTradecomplainBatchqueryResponseModel
	isSet bool
}

func (v NullableAlipayMerchantTradecomplainBatchqueryResponseModel) Get() *AlipayMerchantTradecomplainBatchqueryResponseModel {
	return v.value
}

func (v *NullableAlipayMerchantTradecomplainBatchqueryResponseModel) Set(val *AlipayMerchantTradecomplainBatchqueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMerchantTradecomplainBatchqueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMerchantTradecomplainBatchqueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMerchantTradecomplainBatchqueryResponseModel(val *AlipayMerchantTradecomplainBatchqueryResponseModel) *NullableAlipayMerchantTradecomplainBatchqueryResponseModel {
	return &NullableAlipayMerchantTradecomplainBatchqueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMerchantTradecomplainBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMerchantTradecomplainBatchqueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

