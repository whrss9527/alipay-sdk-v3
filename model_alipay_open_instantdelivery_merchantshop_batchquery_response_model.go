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

// checks if the AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel{}

// AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel struct for AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel
type AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel struct {
	// 当前页码
	PageNo *int32 `json:"page_no,omitempty"`
	// 分页数量, 最大50。
	PageSize *int32 `json:"page_size,omitempty"`
	// 门店列表数据。
	Records []MerchantShopDTO `json:"records,omitempty"`
	// 总数量
	TotalNo *int32 `json:"total_no,omitempty"`
}

// NewAlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel instantiates a new AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel() *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel {
	this := AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel{}
	return &this
}

// NewAlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModelWithDefaults instantiates a new AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModelWithDefaults() *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel {
	this := AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel{}
	return &this
}

// GetPageNo returns the PageNo field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) GetPageNo() int32 {
	if o == nil || IsNil(o.PageNo) {
		var ret int32
		return ret
	}
	return *o.PageNo
}

// GetPageNoOk returns a tuple with the PageNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) GetPageNoOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNo) {
		return nil, false
	}
	return o.PageNo, true
}

// HasPageNo returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) HasPageNo() bool {
	if o != nil && !IsNil(o.PageNo) {
		return true
	}

	return false
}

// SetPageNo gets a reference to the given int32 and assigns it to the PageNo field.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) SetPageNo(v int32) {
	o.PageNo = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) GetRecords() []MerchantShopDTO {
	if o == nil || IsNil(o.Records) {
		var ret []MerchantShopDTO
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) GetRecordsOk() ([]MerchantShopDTO, bool) {
	if o == nil || IsNil(o.Records) {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) HasRecords() bool {
	if o != nil && !IsNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []MerchantShopDTO and assigns it to the Records field.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) SetRecords(v []MerchantShopDTO) {
	o.Records = v
}

// GetTotalNo returns the TotalNo field value if set, zero value otherwise.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) GetTotalNo() int32 {
	if o == nil || IsNil(o.TotalNo) {
		var ret int32
		return ret
	}
	return *o.TotalNo
}

// GetTotalNoOk returns a tuple with the TotalNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) GetTotalNoOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalNo) {
		return nil, false
	}
	return o.TotalNo, true
}

// HasTotalNo returns a boolean if a field has been set.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) HasTotalNo() bool {
	if o != nil && !IsNil(o.TotalNo) {
		return true
	}

	return false
}

// SetTotalNo gets a reference to the given int32 and assigns it to the TotalNo field.
func (o *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) SetTotalNo(v int32) {
	o.TotalNo = &v
}

func (o AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageNo) {
		toSerialize["page_no"] = o.PageNo
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.Records) {
		toSerialize["records"] = o.Records
	}
	if !IsNil(o.TotalNo) {
		toSerialize["total_no"] = o.TotalNo
	}
	return toSerialize, nil
}

type NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel struct {
	value *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) Get() *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) Set(val *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel(val *AlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) *NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel {
	return &NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenInstantdeliveryMerchantshopBatchqueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
