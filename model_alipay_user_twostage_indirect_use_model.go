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

// checks if the AlipayUserTwostageIndirectUseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayUserTwostageIndirectUseModel{}

// AlipayUserTwostageIndirectUseModel struct for AlipayUserTwostageIndirectUseModel
type AlipayUserTwostageIndirectUseModel struct {
	// 商户扫描用户的付款码值。18~24位，25~30开头，例如28开头的18位的数字；或人脸支付的ftoken等。
	DynamicId *string `json:"dynamic_id,omitempty"`
	// 进件信息中，SMID对应的银行机构的PID信息，一般为2088开头的16位数字。
	OrgPid *string `json:"org_pid,omitempty"`
	// 进件信息中，二级商户ID（ sub_merchant_id)信息，一般为2088开头的16位数字。
	PaySmid *string `json:"pay_smid,omitempty"`
	// 外部业务号，用于标识这笔解码请求，对同一个码的重复解码请求，sence_no必须与上一次保持一致，每次请求的sence_no必须不一样，如alipay.user.twostage.common.use接口配合alipay.trade.pay（统一收单交易支付接口）一并使用时，alipay.trade.pay接口的extend_params属性中必须设置DYNAMIC_TOKEN_OUT_BIZ_NO，且值必须与sence_no保持一致
	SenceNo *string `json:"sence_no,omitempty"`
	// 进件信息中，SMID对应渠道的PID信息，一般为2088开头的16位数字。
	SourcePid *string `json:"source_pid,omitempty"`
}

// NewAlipayUserTwostageIndirectUseModel instantiates a new AlipayUserTwostageIndirectUseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayUserTwostageIndirectUseModel() *AlipayUserTwostageIndirectUseModel {
	this := AlipayUserTwostageIndirectUseModel{}
	return &this
}

// NewAlipayUserTwostageIndirectUseModelWithDefaults instantiates a new AlipayUserTwostageIndirectUseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayUserTwostageIndirectUseModelWithDefaults() *AlipayUserTwostageIndirectUseModel {
	this := AlipayUserTwostageIndirectUseModel{}
	return &this
}

// GetDynamicId returns the DynamicId field value if set, zero value otherwise.
func (o *AlipayUserTwostageIndirectUseModel) GetDynamicId() string {
	if o == nil || IsNil(o.DynamicId) {
		var ret string
		return ret
	}
	return *o.DynamicId
}

// GetDynamicIdOk returns a tuple with the DynamicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserTwostageIndirectUseModel) GetDynamicIdOk() (*string, bool) {
	if o == nil || IsNil(o.DynamicId) {
		return nil, false
	}
	return o.DynamicId, true
}

// HasDynamicId returns a boolean if a field has been set.
func (o *AlipayUserTwostageIndirectUseModel) HasDynamicId() bool {
	if o != nil && !IsNil(o.DynamicId) {
		return true
	}

	return false
}

// SetDynamicId gets a reference to the given string and assigns it to the DynamicId field.
func (o *AlipayUserTwostageIndirectUseModel) SetDynamicId(v string) {
	o.DynamicId = &v
}

// GetOrgPid returns the OrgPid field value if set, zero value otherwise.
func (o *AlipayUserTwostageIndirectUseModel) GetOrgPid() string {
	if o == nil || IsNil(o.OrgPid) {
		var ret string
		return ret
	}
	return *o.OrgPid
}

// GetOrgPidOk returns a tuple with the OrgPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserTwostageIndirectUseModel) GetOrgPidOk() (*string, bool) {
	if o == nil || IsNil(o.OrgPid) {
		return nil, false
	}
	return o.OrgPid, true
}

// HasOrgPid returns a boolean if a field has been set.
func (o *AlipayUserTwostageIndirectUseModel) HasOrgPid() bool {
	if o != nil && !IsNil(o.OrgPid) {
		return true
	}

	return false
}

// SetOrgPid gets a reference to the given string and assigns it to the OrgPid field.
func (o *AlipayUserTwostageIndirectUseModel) SetOrgPid(v string) {
	o.OrgPid = &v
}

// GetPaySmid returns the PaySmid field value if set, zero value otherwise.
func (o *AlipayUserTwostageIndirectUseModel) GetPaySmid() string {
	if o == nil || IsNil(o.PaySmid) {
		var ret string
		return ret
	}
	return *o.PaySmid
}

// GetPaySmidOk returns a tuple with the PaySmid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserTwostageIndirectUseModel) GetPaySmidOk() (*string, bool) {
	if o == nil || IsNil(o.PaySmid) {
		return nil, false
	}
	return o.PaySmid, true
}

// HasPaySmid returns a boolean if a field has been set.
func (o *AlipayUserTwostageIndirectUseModel) HasPaySmid() bool {
	if o != nil && !IsNil(o.PaySmid) {
		return true
	}

	return false
}

// SetPaySmid gets a reference to the given string and assigns it to the PaySmid field.
func (o *AlipayUserTwostageIndirectUseModel) SetPaySmid(v string) {
	o.PaySmid = &v
}

// GetSenceNo returns the SenceNo field value if set, zero value otherwise.
func (o *AlipayUserTwostageIndirectUseModel) GetSenceNo() string {
	if o == nil || IsNil(o.SenceNo) {
		var ret string
		return ret
	}
	return *o.SenceNo
}

// GetSenceNoOk returns a tuple with the SenceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserTwostageIndirectUseModel) GetSenceNoOk() (*string, bool) {
	if o == nil || IsNil(o.SenceNo) {
		return nil, false
	}
	return o.SenceNo, true
}

// HasSenceNo returns a boolean if a field has been set.
func (o *AlipayUserTwostageIndirectUseModel) HasSenceNo() bool {
	if o != nil && !IsNil(o.SenceNo) {
		return true
	}

	return false
}

// SetSenceNo gets a reference to the given string and assigns it to the SenceNo field.
func (o *AlipayUserTwostageIndirectUseModel) SetSenceNo(v string) {
	o.SenceNo = &v
}

// GetSourcePid returns the SourcePid field value if set, zero value otherwise.
func (o *AlipayUserTwostageIndirectUseModel) GetSourcePid() string {
	if o == nil || IsNil(o.SourcePid) {
		var ret string
		return ret
	}
	return *o.SourcePid
}

// GetSourcePidOk returns a tuple with the SourcePid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserTwostageIndirectUseModel) GetSourcePidOk() (*string, bool) {
	if o == nil || IsNil(o.SourcePid) {
		return nil, false
	}
	return o.SourcePid, true
}

// HasSourcePid returns a boolean if a field has been set.
func (o *AlipayUserTwostageIndirectUseModel) HasSourcePid() bool {
	if o != nil && !IsNil(o.SourcePid) {
		return true
	}

	return false
}

// SetSourcePid gets a reference to the given string and assigns it to the SourcePid field.
func (o *AlipayUserTwostageIndirectUseModel) SetSourcePid(v string) {
	o.SourcePid = &v
}

func (o AlipayUserTwostageIndirectUseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayUserTwostageIndirectUseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DynamicId) {
		toSerialize["dynamic_id"] = o.DynamicId
	}
	if !IsNil(o.OrgPid) {
		toSerialize["org_pid"] = o.OrgPid
	}
	if !IsNil(o.PaySmid) {
		toSerialize["pay_smid"] = o.PaySmid
	}
	if !IsNil(o.SenceNo) {
		toSerialize["sence_no"] = o.SenceNo
	}
	if !IsNil(o.SourcePid) {
		toSerialize["source_pid"] = o.SourcePid
	}
	return toSerialize, nil
}

type NullableAlipayUserTwostageIndirectUseModel struct {
	value *AlipayUserTwostageIndirectUseModel
	isSet bool
}

func (v NullableAlipayUserTwostageIndirectUseModel) Get() *AlipayUserTwostageIndirectUseModel {
	return v.value
}

func (v *NullableAlipayUserTwostageIndirectUseModel) Set(val *AlipayUserTwostageIndirectUseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayUserTwostageIndirectUseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayUserTwostageIndirectUseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayUserTwostageIndirectUseModel(val *AlipayUserTwostageIndirectUseModel) *NullableAlipayUserTwostageIndirectUseModel {
	return &NullableAlipayUserTwostageIndirectUseModel{value: val, isSet: true}
}

func (v NullableAlipayUserTwostageIndirectUseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayUserTwostageIndirectUseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
