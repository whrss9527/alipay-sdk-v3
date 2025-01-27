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

// checks if the ZhimaCreditPayafteruseCreditbizorderFinishModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZhimaCreditPayafteruseCreditbizorderFinishModel{}

// ZhimaCreditPayafteruseCreditbizorderFinishModel struct for ZhimaCreditPayafteruseCreditbizorderFinishModel
type ZhimaCreditPayafteruseCreditbizorderFinishModel struct {
	// 信用服务订单号
	CreditBizOrderId *string `json:"credit_biz_order_id,omitempty"`
	// 字符串类型，用户此订单是否守约。 传\"true\"时，用户在芝麻信用-守约记录中，该笔订单是已守约状态； 传\"false\"时，用户在芝麻信用-守约记录中，该笔订单是已取消状态。
	IsFulfilled *string `json:"is_fulfilled,omitempty"`
	// 商户外部请求号
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// 商户对本次操作的附言描述
	Remark *string `json:"remark,omitempty"`
}

// NewZhimaCreditPayafteruseCreditbizorderFinishModel instantiates a new ZhimaCreditPayafteruseCreditbizorderFinishModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZhimaCreditPayafteruseCreditbizorderFinishModel() *ZhimaCreditPayafteruseCreditbizorderFinishModel {
	this := ZhimaCreditPayafteruseCreditbizorderFinishModel{}
	return &this
}

// NewZhimaCreditPayafteruseCreditbizorderFinishModelWithDefaults instantiates a new ZhimaCreditPayafteruseCreditbizorderFinishModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZhimaCreditPayafteruseCreditbizorderFinishModelWithDefaults() *ZhimaCreditPayafteruseCreditbizorderFinishModel {
	this := ZhimaCreditPayafteruseCreditbizorderFinishModel{}
	return &this
}

// GetCreditBizOrderId returns the CreditBizOrderId field value if set, zero value otherwise.
func (o *ZhimaCreditPayafteruseCreditbizorderFinishModel) GetCreditBizOrderId() string {
	if o == nil || IsNil(o.CreditBizOrderId) {
		var ret string
		return ret
	}
	return *o.CreditBizOrderId
}

// GetCreditBizOrderIdOk returns a tuple with the CreditBizOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPayafteruseCreditbizorderFinishModel) GetCreditBizOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreditBizOrderId) {
		return nil, false
	}
	return o.CreditBizOrderId, true
}

// HasCreditBizOrderId returns a boolean if a field has been set.
func (o *ZhimaCreditPayafteruseCreditbizorderFinishModel) HasCreditBizOrderId() bool {
	if o != nil && !IsNil(o.CreditBizOrderId) {
		return true
	}

	return false
}

// SetCreditBizOrderId gets a reference to the given string and assigns it to the CreditBizOrderId field.
func (o *ZhimaCreditPayafteruseCreditbizorderFinishModel) SetCreditBizOrderId(v string) {
	o.CreditBizOrderId = &v
}

// GetIsFulfilled returns the IsFulfilled field value if set, zero value otherwise.
func (o *ZhimaCreditPayafteruseCreditbizorderFinishModel) GetIsFulfilled() string {
	if o == nil || IsNil(o.IsFulfilled) {
		var ret string
		return ret
	}
	return *o.IsFulfilled
}

// GetIsFulfilledOk returns a tuple with the IsFulfilled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPayafteruseCreditbizorderFinishModel) GetIsFulfilledOk() (*string, bool) {
	if o == nil || IsNil(o.IsFulfilled) {
		return nil, false
	}
	return o.IsFulfilled, true
}

// HasIsFulfilled returns a boolean if a field has been set.
func (o *ZhimaCreditPayafteruseCreditbizorderFinishModel) HasIsFulfilled() bool {
	if o != nil && !IsNil(o.IsFulfilled) {
		return true
	}

	return false
}

// SetIsFulfilled gets a reference to the given string and assigns it to the IsFulfilled field.
func (o *ZhimaCreditPayafteruseCreditbizorderFinishModel) SetIsFulfilled(v string) {
	o.IsFulfilled = &v
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *ZhimaCreditPayafteruseCreditbizorderFinishModel) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPayafteruseCreditbizorderFinishModel) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *ZhimaCreditPayafteruseCreditbizorderFinishModel) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *ZhimaCreditPayafteruseCreditbizorderFinishModel) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *ZhimaCreditPayafteruseCreditbizorderFinishModel) GetRemark() string {
	if o == nil || IsNil(o.Remark) {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPayafteruseCreditbizorderFinishModel) GetRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.Remark) {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *ZhimaCreditPayafteruseCreditbizorderFinishModel) HasRemark() bool {
	if o != nil && !IsNil(o.Remark) {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *ZhimaCreditPayafteruseCreditbizorderFinishModel) SetRemark(v string) {
	o.Remark = &v
}

func (o ZhimaCreditPayafteruseCreditbizorderFinishModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZhimaCreditPayafteruseCreditbizorderFinishModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreditBizOrderId) {
		toSerialize["credit_biz_order_id"] = o.CreditBizOrderId
	}
	if !IsNil(o.IsFulfilled) {
		toSerialize["is_fulfilled"] = o.IsFulfilled
	}
	if !IsNil(o.OutRequestNo) {
		toSerialize["out_request_no"] = o.OutRequestNo
	}
	if !IsNil(o.Remark) {
		toSerialize["remark"] = o.Remark
	}
	return toSerialize, nil
}

type NullableZhimaCreditPayafteruseCreditbizorderFinishModel struct {
	value *ZhimaCreditPayafteruseCreditbizorderFinishModel
	isSet bool
}

func (v NullableZhimaCreditPayafteruseCreditbizorderFinishModel) Get() *ZhimaCreditPayafteruseCreditbizorderFinishModel {
	return v.value
}

func (v *NullableZhimaCreditPayafteruseCreditbizorderFinishModel) Set(val *ZhimaCreditPayafteruseCreditbizorderFinishModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaCreditPayafteruseCreditbizorderFinishModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaCreditPayafteruseCreditbizorderFinishModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaCreditPayafteruseCreditbizorderFinishModel(val *ZhimaCreditPayafteruseCreditbizorderFinishModel) *NullableZhimaCreditPayafteruseCreditbizorderFinishModel {
	return &NullableZhimaCreditPayafteruseCreditbizorderFinishModel{value: val, isSet: true}
}

func (v NullableZhimaCreditPayafteruseCreditbizorderFinishModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaCreditPayafteruseCreditbizorderFinishModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
