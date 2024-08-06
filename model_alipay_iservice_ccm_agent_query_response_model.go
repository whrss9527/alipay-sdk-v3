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

// checks if the AlipayIserviceCcmAgentQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmAgentQueryResponseModel{}

// AlipayIserviceCcmAgentQueryResponseModel struct for AlipayIserviceCcmAgentQueryResponseModel
type AlipayIserviceCcmAgentQueryResponseModel struct {
	// 客服列表
	Agents []AgentVO `json:"agents,omitempty"`
	// 查询结果的页码，起始值为 1，默认值为 1
	PageNum *int32 `json:"page_num,omitempty"`
	// 分页查询时设置的每页记录数，最大值 100 行，默认为 10
	PageSize *int32 `json:"page_size,omitempty"`
	// 总条目数
	TotalCount *int32 `json:"total_count,omitempty"`
}

// NewAlipayIserviceCcmAgentQueryResponseModel instantiates a new AlipayIserviceCcmAgentQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmAgentQueryResponseModel() *AlipayIserviceCcmAgentQueryResponseModel {
	this := AlipayIserviceCcmAgentQueryResponseModel{}
	return &this
}

// NewAlipayIserviceCcmAgentQueryResponseModelWithDefaults instantiates a new AlipayIserviceCcmAgentQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmAgentQueryResponseModelWithDefaults() *AlipayIserviceCcmAgentQueryResponseModel {
	this := AlipayIserviceCcmAgentQueryResponseModel{}
	return &this
}

// GetAgents returns the Agents field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentQueryResponseModel) GetAgents() []AgentVO {
	if o == nil || IsNil(o.Agents) {
		var ret []AgentVO
		return ret
	}
	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentQueryResponseModel) GetAgentsOk() ([]AgentVO, bool) {
	if o == nil || IsNil(o.Agents) {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentQueryResponseModel) HasAgents() bool {
	if o != nil && !IsNil(o.Agents) {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []AgentVO and assigns it to the Agents field.
func (o *AlipayIserviceCcmAgentQueryResponseModel) SetAgents(v []AgentVO) {
	o.Agents = v
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentQueryResponseModel) GetPageNum() int32 {
	if o == nil || IsNil(o.PageNum) {
		var ret int32
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentQueryResponseModel) GetPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentQueryResponseModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int32 and assigns it to the PageNum field.
func (o *AlipayIserviceCcmAgentQueryResponseModel) SetPageNum(v int32) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentQueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentQueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentQueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayIserviceCcmAgentQueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentQueryResponseModel) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentQueryResponseModel) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentQueryResponseModel) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *AlipayIserviceCcmAgentQueryResponseModel) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o AlipayIserviceCcmAgentQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmAgentQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Agents) {
		toSerialize["agents"] = o.Agents
	}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmAgentQueryResponseModel struct {
	value *AlipayIserviceCcmAgentQueryResponseModel
	isSet bool
}

func (v NullableAlipayIserviceCcmAgentQueryResponseModel) Get() *AlipayIserviceCcmAgentQueryResponseModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmAgentQueryResponseModel) Set(val *AlipayIserviceCcmAgentQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmAgentQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmAgentQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmAgentQueryResponseModel(val *AlipayIserviceCcmAgentQueryResponseModel) *NullableAlipayIserviceCcmAgentQueryResponseModel {
	return &NullableAlipayIserviceCcmAgentQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmAgentQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmAgentQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

