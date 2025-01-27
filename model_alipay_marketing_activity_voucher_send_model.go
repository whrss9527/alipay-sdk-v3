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

// checks if the AlipayMarketingActivityVoucherSendModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityVoucherSendModel{}

// AlipayMarketingActivityVoucherSendModel struct for AlipayMarketingActivityVoucherSendModel
type AlipayMarketingActivityVoucherSendModel struct {
	// 活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// 商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// 领券的支付宝用户openId
	OpenId *string `json:"open_id,omitempty"`
	// 外部业务单号，用作幂等控制。  幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。  外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 领券的支付宝user_id账号
	UserId *string `json:"user_id,omitempty"`
}

// NewAlipayMarketingActivityVoucherSendModel instantiates a new AlipayMarketingActivityVoucherSendModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityVoucherSendModel() *AlipayMarketingActivityVoucherSendModel {
	this := AlipayMarketingActivityVoucherSendModel{}
	return &this
}

// NewAlipayMarketingActivityVoucherSendModelWithDefaults instantiates a new AlipayMarketingActivityVoucherSendModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityVoucherSendModelWithDefaults() *AlipayMarketingActivityVoucherSendModel {
	this := AlipayMarketingActivityVoucherSendModel{}
	return &this
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherSendModel) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId) {
		var ret string
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherSendModel) GetActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherSendModel) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given string and assigns it to the ActivityId field.
func (o *AlipayMarketingActivityVoucherSendModel) SetActivityId(v string) {
	o.ActivityId = &v
}

// GetMerchantAccessMode returns the MerchantAccessMode field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherSendModel) GetMerchantAccessMode() string {
	if o == nil || IsNil(o.MerchantAccessMode) {
		var ret string
		return ret
	}
	return *o.MerchantAccessMode
}

// GetMerchantAccessModeOk returns a tuple with the MerchantAccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherSendModel) GetMerchantAccessModeOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantAccessMode) {
		return nil, false
	}
	return o.MerchantAccessMode, true
}

// HasMerchantAccessMode returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherSendModel) HasMerchantAccessMode() bool {
	if o != nil && !IsNil(o.MerchantAccessMode) {
		return true
	}

	return false
}

// SetMerchantAccessMode gets a reference to the given string and assigns it to the MerchantAccessMode field.
func (o *AlipayMarketingActivityVoucherSendModel) SetMerchantAccessMode(v string) {
	o.MerchantAccessMode = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherSendModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherSendModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherSendModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AlipayMarketingActivityVoucherSendModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherSendModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherSendModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherSendModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayMarketingActivityVoucherSendModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherSendModel) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherSendModel) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherSendModel) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AlipayMarketingActivityVoucherSendModel) SetUserId(v string) {
	o.UserId = &v
}

func (o AlipayMarketingActivityVoucherSendModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityVoucherSendModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
	}
	if !IsNil(o.MerchantAccessMode) {
		toSerialize["merchant_access_mode"] = o.MerchantAccessMode
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableAlipayMarketingActivityVoucherSendModel struct {
	value *AlipayMarketingActivityVoucherSendModel
	isSet bool
}

func (v NullableAlipayMarketingActivityVoucherSendModel) Get() *AlipayMarketingActivityVoucherSendModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityVoucherSendModel) Set(val *AlipayMarketingActivityVoucherSendModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityVoucherSendModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityVoucherSendModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityVoucherSendModel(val *AlipayMarketingActivityVoucherSendModel) *NullableAlipayMarketingActivityVoucherSendModel {
	return &NullableAlipayMarketingActivityVoucherSendModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityVoucherSendModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityVoucherSendModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
