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

// checks if the AlipayMerchantIndirectAuthorderQuerystatusModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMerchantIndirectAuthorderQuerystatusModel{}

// AlipayMerchantIndirectAuthorderQuerystatusModel struct for AlipayMerchantIndirectAuthorderQuerystatusModel
type AlipayMerchantIndirectAuthorderQuerystatusModel struct {
	// 商家认证申请单号，参数二选一
	OrderNo *string `json:"order_no,omitempty"`
	// 外部业务号，参数二选一，业务自定义，保证唯一
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 1. 收单机构调用API提交申请单时，可选择是否指定单个服务商范围。非收单机构无需填写此字段。 2. 此字段填写单个服务商pid信息：填写（即：单服务商提交认证方式），查询申请单返回的认证二维码qr_code和填写服务商对应，仅能认证填写服务商下的商户；不填写（即：全服务商提交认证方式），查询申请单返回的认证二维码qr_code和收单机构对应，可认证收单机构下全部商户。
	Source *string `json:"source,omitempty"`
}

// NewAlipayMerchantIndirectAuthorderQuerystatusModel instantiates a new AlipayMerchantIndirectAuthorderQuerystatusModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMerchantIndirectAuthorderQuerystatusModel() *AlipayMerchantIndirectAuthorderQuerystatusModel {
	this := AlipayMerchantIndirectAuthorderQuerystatusModel{}
	return &this
}

// NewAlipayMerchantIndirectAuthorderQuerystatusModelWithDefaults instantiates a new AlipayMerchantIndirectAuthorderQuerystatusModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMerchantIndirectAuthorderQuerystatusModelWithDefaults() *AlipayMerchantIndirectAuthorderQuerystatusModel {
	this := AlipayMerchantIndirectAuthorderQuerystatusModel{}
	return &this
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *AlipayMerchantIndirectAuthorderQuerystatusModel) GetOrderNo() string {
	if o == nil || IsNil(o.OrderNo) {
		var ret string
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantIndirectAuthorderQuerystatusModel) GetOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *AlipayMerchantIndirectAuthorderQuerystatusModel) HasOrderNo() bool {
	if o != nil && !IsNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given string and assigns it to the OrderNo field.
func (o *AlipayMerchantIndirectAuthorderQuerystatusModel) SetOrderNo(v string) {
	o.OrderNo = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayMerchantIndirectAuthorderQuerystatusModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantIndirectAuthorderQuerystatusModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayMerchantIndirectAuthorderQuerystatusModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayMerchantIndirectAuthorderQuerystatusModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AlipayMerchantIndirectAuthorderQuerystatusModel) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantIndirectAuthorderQuerystatusModel) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AlipayMerchantIndirectAuthorderQuerystatusModel) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *AlipayMerchantIndirectAuthorderQuerystatusModel) SetSource(v string) {
	o.Source = &v
}

func (o AlipayMerchantIndirectAuthorderQuerystatusModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMerchantIndirectAuthorderQuerystatusModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderNo) {
		toSerialize["order_no"] = o.OrderNo
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableAlipayMerchantIndirectAuthorderQuerystatusModel struct {
	value *AlipayMerchantIndirectAuthorderQuerystatusModel
	isSet bool
}

func (v NullableAlipayMerchantIndirectAuthorderQuerystatusModel) Get() *AlipayMerchantIndirectAuthorderQuerystatusModel {
	return v.value
}

func (v *NullableAlipayMerchantIndirectAuthorderQuerystatusModel) Set(val *AlipayMerchantIndirectAuthorderQuerystatusModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMerchantIndirectAuthorderQuerystatusModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMerchantIndirectAuthorderQuerystatusModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMerchantIndirectAuthorderQuerystatusModel(val *AlipayMerchantIndirectAuthorderQuerystatusModel) *NullableAlipayMerchantIndirectAuthorderQuerystatusModel {
	return &NullableAlipayMerchantIndirectAuthorderQuerystatusModel{value: val, isSet: true}
}

func (v NullableAlipayMerchantIndirectAuthorderQuerystatusModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMerchantIndirectAuthorderQuerystatusModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
