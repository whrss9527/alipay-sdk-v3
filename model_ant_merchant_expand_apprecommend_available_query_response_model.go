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

// checks if the AntMerchantExpandApprecommendAvailableQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AntMerchantExpandApprecommendAvailableQueryResponseModel{}

// AntMerchantExpandApprecommendAvailableQueryResponseModel struct for AntMerchantExpandApprecommendAvailableQueryResponseModel
type AntMerchantExpandApprecommendAvailableQueryResponseModel struct {
	// 总页数
	List []string `json:"list,omitempty"`
	// 当前页码
	PageNumber *int32 `json:"page_number,omitempty"`
	// 单页行数
	PageSize *int32 `json:"page_size,omitempty"`
	// 总页数
	TotalPages *int32 `json:"total_pages,omitempty"`
	// 总行数
	TotalSize *int32 `json:"total_size,omitempty"`
}

// NewAntMerchantExpandApprecommendAvailableQueryResponseModel instantiates a new AntMerchantExpandApprecommendAvailableQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntMerchantExpandApprecommendAvailableQueryResponseModel() *AntMerchantExpandApprecommendAvailableQueryResponseModel {
	this := AntMerchantExpandApprecommendAvailableQueryResponseModel{}
	return &this
}

// NewAntMerchantExpandApprecommendAvailableQueryResponseModelWithDefaults instantiates a new AntMerchantExpandApprecommendAvailableQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntMerchantExpandApprecommendAvailableQueryResponseModelWithDefaults() *AntMerchantExpandApprecommendAvailableQueryResponseModel {
	this := AntMerchantExpandApprecommendAvailableQueryResponseModel{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) GetList() []string {
	if o == nil || IsNil(o.List) {
		var ret []string
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) GetListOk() ([]string, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []string and assigns it to the List field.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) SetList(v []string) {
	o.List = v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) GetTotalSize() int32 {
	if o == nil || IsNil(o.TotalSize) {
		var ret int32
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) GetTotalSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int32 and assigns it to the TotalSize field.
func (o *AntMerchantExpandApprecommendAvailableQueryResponseModel) SetTotalSize(v int32) {
	o.TotalSize = &v
}

func (o AntMerchantExpandApprecommendAvailableQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AntMerchantExpandApprecommendAvailableQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	if !IsNil(o.PageNumber) {
		toSerialize["page_number"] = o.PageNumber
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.TotalPages) {
		toSerialize["total_pages"] = o.TotalPages
	}
	if !IsNil(o.TotalSize) {
		toSerialize["total_size"] = o.TotalSize
	}
	return toSerialize, nil
}

type NullableAntMerchantExpandApprecommendAvailableQueryResponseModel struct {
	value *AntMerchantExpandApprecommendAvailableQueryResponseModel
	isSet bool
}

func (v NullableAntMerchantExpandApprecommendAvailableQueryResponseModel) Get() *AntMerchantExpandApprecommendAvailableQueryResponseModel {
	return v.value
}

func (v *NullableAntMerchantExpandApprecommendAvailableQueryResponseModel) Set(val *AntMerchantExpandApprecommendAvailableQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandApprecommendAvailableQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandApprecommendAvailableQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandApprecommendAvailableQueryResponseModel(val *AntMerchantExpandApprecommendAvailableQueryResponseModel) *NullableAntMerchantExpandApprecommendAvailableQueryResponseModel {
	return &NullableAntMerchantExpandApprecommendAvailableQueryResponseModel{value: val, isSet: true}
}

func (v NullableAntMerchantExpandApprecommendAvailableQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandApprecommendAvailableQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
