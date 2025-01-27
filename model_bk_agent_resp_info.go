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

// checks if the BkAgentRespInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BkAgentRespInfo{}

// BkAgentRespInfo struct for BkAgentRespInfo
type BkAgentRespInfo struct {
	// 枚举值，01 银联；02 网联；03 连通等
	BindclrissrId *string `json:"bindclrissr_id,omitempty"`
	// 付款机构在清算组织登记或分配的机构代码
	BindpyeracctbkId *string `json:"bindpyeracctbk_id,omitempty"`
	// 原快捷交易流水号
	BindtrxId *string `json:"bindtrx_id,omitempty"`
	// 用户在银行付款账号的标记化处理编号
	BkpyeruserCode *string `json:"bkpyeruser_code,omitempty"`
	// 设备推测位置
	EstterLocation *string `json:"estter_location,omitempty"`
}

// NewBkAgentRespInfo instantiates a new BkAgentRespInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBkAgentRespInfo() *BkAgentRespInfo {
	this := BkAgentRespInfo{}
	return &this
}

// NewBkAgentRespInfoWithDefaults instantiates a new BkAgentRespInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBkAgentRespInfoWithDefaults() *BkAgentRespInfo {
	this := BkAgentRespInfo{}
	return &this
}

// GetBindclrissrId returns the BindclrissrId field value if set, zero value otherwise.
func (o *BkAgentRespInfo) GetBindclrissrId() string {
	if o == nil || IsNil(o.BindclrissrId) {
		var ret string
		return ret
	}
	return *o.BindclrissrId
}

// GetBindclrissrIdOk returns a tuple with the BindclrissrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BkAgentRespInfo) GetBindclrissrIdOk() (*string, bool) {
	if o == nil || IsNil(o.BindclrissrId) {
		return nil, false
	}
	return o.BindclrissrId, true
}

// HasBindclrissrId returns a boolean if a field has been set.
func (o *BkAgentRespInfo) HasBindclrissrId() bool {
	if o != nil && !IsNil(o.BindclrissrId) {
		return true
	}

	return false
}

// SetBindclrissrId gets a reference to the given string and assigns it to the BindclrissrId field.
func (o *BkAgentRespInfo) SetBindclrissrId(v string) {
	o.BindclrissrId = &v
}

// GetBindpyeracctbkId returns the BindpyeracctbkId field value if set, zero value otherwise.
func (o *BkAgentRespInfo) GetBindpyeracctbkId() string {
	if o == nil || IsNil(o.BindpyeracctbkId) {
		var ret string
		return ret
	}
	return *o.BindpyeracctbkId
}

// GetBindpyeracctbkIdOk returns a tuple with the BindpyeracctbkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BkAgentRespInfo) GetBindpyeracctbkIdOk() (*string, bool) {
	if o == nil || IsNil(o.BindpyeracctbkId) {
		return nil, false
	}
	return o.BindpyeracctbkId, true
}

// HasBindpyeracctbkId returns a boolean if a field has been set.
func (o *BkAgentRespInfo) HasBindpyeracctbkId() bool {
	if o != nil && !IsNil(o.BindpyeracctbkId) {
		return true
	}

	return false
}

// SetBindpyeracctbkId gets a reference to the given string and assigns it to the BindpyeracctbkId field.
func (o *BkAgentRespInfo) SetBindpyeracctbkId(v string) {
	o.BindpyeracctbkId = &v
}

// GetBindtrxId returns the BindtrxId field value if set, zero value otherwise.
func (o *BkAgentRespInfo) GetBindtrxId() string {
	if o == nil || IsNil(o.BindtrxId) {
		var ret string
		return ret
	}
	return *o.BindtrxId
}

// GetBindtrxIdOk returns a tuple with the BindtrxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BkAgentRespInfo) GetBindtrxIdOk() (*string, bool) {
	if o == nil || IsNil(o.BindtrxId) {
		return nil, false
	}
	return o.BindtrxId, true
}

// HasBindtrxId returns a boolean if a field has been set.
func (o *BkAgentRespInfo) HasBindtrxId() bool {
	if o != nil && !IsNil(o.BindtrxId) {
		return true
	}

	return false
}

// SetBindtrxId gets a reference to the given string and assigns it to the BindtrxId field.
func (o *BkAgentRespInfo) SetBindtrxId(v string) {
	o.BindtrxId = &v
}

// GetBkpyeruserCode returns the BkpyeruserCode field value if set, zero value otherwise.
func (o *BkAgentRespInfo) GetBkpyeruserCode() string {
	if o == nil || IsNil(o.BkpyeruserCode) {
		var ret string
		return ret
	}
	return *o.BkpyeruserCode
}

// GetBkpyeruserCodeOk returns a tuple with the BkpyeruserCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BkAgentRespInfo) GetBkpyeruserCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BkpyeruserCode) {
		return nil, false
	}
	return o.BkpyeruserCode, true
}

// HasBkpyeruserCode returns a boolean if a field has been set.
func (o *BkAgentRespInfo) HasBkpyeruserCode() bool {
	if o != nil && !IsNil(o.BkpyeruserCode) {
		return true
	}

	return false
}

// SetBkpyeruserCode gets a reference to the given string and assigns it to the BkpyeruserCode field.
func (o *BkAgentRespInfo) SetBkpyeruserCode(v string) {
	o.BkpyeruserCode = &v
}

// GetEstterLocation returns the EstterLocation field value if set, zero value otherwise.
func (o *BkAgentRespInfo) GetEstterLocation() string {
	if o == nil || IsNil(o.EstterLocation) {
		var ret string
		return ret
	}
	return *o.EstterLocation
}

// GetEstterLocationOk returns a tuple with the EstterLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BkAgentRespInfo) GetEstterLocationOk() (*string, bool) {
	if o == nil || IsNil(o.EstterLocation) {
		return nil, false
	}
	return o.EstterLocation, true
}

// HasEstterLocation returns a boolean if a field has been set.
func (o *BkAgentRespInfo) HasEstterLocation() bool {
	if o != nil && !IsNil(o.EstterLocation) {
		return true
	}

	return false
}

// SetEstterLocation gets a reference to the given string and assigns it to the EstterLocation field.
func (o *BkAgentRespInfo) SetEstterLocation(v string) {
	o.EstterLocation = &v
}

func (o BkAgentRespInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BkAgentRespInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BindclrissrId) {
		toSerialize["bindclrissr_id"] = o.BindclrissrId
	}
	if !IsNil(o.BindpyeracctbkId) {
		toSerialize["bindpyeracctbk_id"] = o.BindpyeracctbkId
	}
	if !IsNil(o.BindtrxId) {
		toSerialize["bindtrx_id"] = o.BindtrxId
	}
	if !IsNil(o.BkpyeruserCode) {
		toSerialize["bkpyeruser_code"] = o.BkpyeruserCode
	}
	if !IsNil(o.EstterLocation) {
		toSerialize["estter_location"] = o.EstterLocation
	}
	return toSerialize, nil
}

type NullableBkAgentRespInfo struct {
	value *BkAgentRespInfo
	isSet bool
}

func (v NullableBkAgentRespInfo) Get() *BkAgentRespInfo {
	return v.value
}

func (v *NullableBkAgentRespInfo) Set(val *BkAgentRespInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBkAgentRespInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBkAgentRespInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBkAgentRespInfo(val *BkAgentRespInfo) *NullableBkAgentRespInfo {
	return &NullableBkAgentRespInfo{value: val, isSet: true}
}

func (v NullableBkAgentRespInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBkAgentRespInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
