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

// checks if the AlipayOpenSpIsvRelationQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSpIsvRelationQueryResponseModel{}

// AlipayOpenSpIsvRelationQueryResponseModel struct for AlipayOpenSpIsvRelationQueryResponseModel
type AlipayOpenSpIsvRelationQueryResponseModel struct {
	// 当前页码
	CurrentPage *int32 `json:"current_page,omitempty"`
	// 每页数量
	PageSize *int32 `json:"page_size,omitempty"`
	// 推广服务商信息列表
	PromotionRelations []PromotionRelationDTO `json:"promotion_relations,omitempty"`
	// 总记录数
	TotalSize *int32 `json:"total_size,omitempty"`
}

// NewAlipayOpenSpIsvRelationQueryResponseModel instantiates a new AlipayOpenSpIsvRelationQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSpIsvRelationQueryResponseModel() *AlipayOpenSpIsvRelationQueryResponseModel {
	this := AlipayOpenSpIsvRelationQueryResponseModel{}
	return &this
}

// NewAlipayOpenSpIsvRelationQueryResponseModelWithDefaults instantiates a new AlipayOpenSpIsvRelationQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSpIsvRelationQueryResponseModelWithDefaults() *AlipayOpenSpIsvRelationQueryResponseModel {
	this := AlipayOpenSpIsvRelationQueryResponseModel{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *AlipayOpenSpIsvRelationQueryResponseModel) GetCurrentPage() int32 {
	if o == nil || IsNil(o.CurrentPage) {
		var ret int32
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpIsvRelationQueryResponseModel) GetCurrentPageOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentPage) {
		return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *AlipayOpenSpIsvRelationQueryResponseModel) HasCurrentPage() bool {
	if o != nil && !IsNil(o.CurrentPage) {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given int32 and assigns it to the CurrentPage field.
func (o *AlipayOpenSpIsvRelationQueryResponseModel) SetCurrentPage(v int32) {
	o.CurrentPage = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayOpenSpIsvRelationQueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpIsvRelationQueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayOpenSpIsvRelationQueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayOpenSpIsvRelationQueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPromotionRelations returns the PromotionRelations field value if set, zero value otherwise.
func (o *AlipayOpenSpIsvRelationQueryResponseModel) GetPromotionRelations() []PromotionRelationDTO {
	if o == nil || IsNil(o.PromotionRelations) {
		var ret []PromotionRelationDTO
		return ret
	}
	return o.PromotionRelations
}

// GetPromotionRelationsOk returns a tuple with the PromotionRelations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpIsvRelationQueryResponseModel) GetPromotionRelationsOk() ([]PromotionRelationDTO, bool) {
	if o == nil || IsNil(o.PromotionRelations) {
		return nil, false
	}
	return o.PromotionRelations, true
}

// HasPromotionRelations returns a boolean if a field has been set.
func (o *AlipayOpenSpIsvRelationQueryResponseModel) HasPromotionRelations() bool {
	if o != nil && !IsNil(o.PromotionRelations) {
		return true
	}

	return false
}

// SetPromotionRelations gets a reference to the given []PromotionRelationDTO and assigns it to the PromotionRelations field.
func (o *AlipayOpenSpIsvRelationQueryResponseModel) SetPromotionRelations(v []PromotionRelationDTO) {
	o.PromotionRelations = v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *AlipayOpenSpIsvRelationQueryResponseModel) GetTotalSize() int32 {
	if o == nil || IsNil(o.TotalSize) {
		var ret int32
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpIsvRelationQueryResponseModel) GetTotalSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *AlipayOpenSpIsvRelationQueryResponseModel) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int32 and assigns it to the TotalSize field.
func (o *AlipayOpenSpIsvRelationQueryResponseModel) SetTotalSize(v int32) {
	o.TotalSize = &v
}

func (o AlipayOpenSpIsvRelationQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSpIsvRelationQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentPage) {
		toSerialize["current_page"] = o.CurrentPage
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.PromotionRelations) {
		toSerialize["promotion_relations"] = o.PromotionRelations
	}
	if !IsNil(o.TotalSize) {
		toSerialize["total_size"] = o.TotalSize
	}
	return toSerialize, nil
}

type NullableAlipayOpenSpIsvRelationQueryResponseModel struct {
	value *AlipayOpenSpIsvRelationQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenSpIsvRelationQueryResponseModel) Get() *AlipayOpenSpIsvRelationQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenSpIsvRelationQueryResponseModel) Set(val *AlipayOpenSpIsvRelationQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSpIsvRelationQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSpIsvRelationQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSpIsvRelationQueryResponseModel(val *AlipayOpenSpIsvRelationQueryResponseModel) *NullableAlipayOpenSpIsvRelationQueryResponseModel {
	return &NullableAlipayOpenSpIsvRelationQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSpIsvRelationQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSpIsvRelationQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

