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

// checks if the AlipayMarketingActivityOrdervoucherCodedepositModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityOrdervoucherCodedepositModel{}

// AlipayMarketingActivityOrdervoucherCodedepositModel struct for AlipayMarketingActivityOrdervoucherCodedepositModel
type AlipayMarketingActivityOrdervoucherCodedepositModel struct {
	// 商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// 外部业务单号，用作幂等控制。  幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。  外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 券码的数量列表。接口参数为列表类型。  限制： 目前最大上传 1000 个。  单个code最长64位。  商户上传的券code列表，code允许包含的字符有0-9、a-z、A-Z、-、_、+、=、|。
	VoucherCodes []string `json:"voucher_codes,omitempty"`
}

// NewAlipayMarketingActivityOrdervoucherCodedepositModel instantiates a new AlipayMarketingActivityOrdervoucherCodedepositModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityOrdervoucherCodedepositModel() *AlipayMarketingActivityOrdervoucherCodedepositModel {
	this := AlipayMarketingActivityOrdervoucherCodedepositModel{}
	return &this
}

// NewAlipayMarketingActivityOrdervoucherCodedepositModelWithDefaults instantiates a new AlipayMarketingActivityOrdervoucherCodedepositModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityOrdervoucherCodedepositModelWithDefaults() *AlipayMarketingActivityOrdervoucherCodedepositModel {
	this := AlipayMarketingActivityOrdervoucherCodedepositModel{}
	return &this
}

// GetMerchantAccessMode returns the MerchantAccessMode field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCodedepositModel) GetMerchantAccessMode() string {
	if o == nil || IsNil(o.MerchantAccessMode) {
		var ret string
		return ret
	}
	return *o.MerchantAccessMode
}

// GetMerchantAccessModeOk returns a tuple with the MerchantAccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCodedepositModel) GetMerchantAccessModeOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantAccessMode) {
		return nil, false
	}
	return o.MerchantAccessMode, true
}

// HasMerchantAccessMode returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCodedepositModel) HasMerchantAccessMode() bool {
	if o != nil && !IsNil(o.MerchantAccessMode) {
		return true
	}

	return false
}

// SetMerchantAccessMode gets a reference to the given string and assigns it to the MerchantAccessMode field.
func (o *AlipayMarketingActivityOrdervoucherCodedepositModel) SetMerchantAccessMode(v string) {
	o.MerchantAccessMode = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCodedepositModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCodedepositModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCodedepositModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayMarketingActivityOrdervoucherCodedepositModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetVoucherCodes returns the VoucherCodes field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCodedepositModel) GetVoucherCodes() []string {
	if o == nil || IsNil(o.VoucherCodes) {
		var ret []string
		return ret
	}
	return o.VoucherCodes
}

// GetVoucherCodesOk returns a tuple with the VoucherCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCodedepositModel) GetVoucherCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.VoucherCodes) {
		return nil, false
	}
	return o.VoucherCodes, true
}

// HasVoucherCodes returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCodedepositModel) HasVoucherCodes() bool {
	if o != nil && !IsNil(o.VoucherCodes) {
		return true
	}

	return false
}

// SetVoucherCodes gets a reference to the given []string and assigns it to the VoucherCodes field.
func (o *AlipayMarketingActivityOrdervoucherCodedepositModel) SetVoucherCodes(v []string) {
	o.VoucherCodes = v
}

func (o AlipayMarketingActivityOrdervoucherCodedepositModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityOrdervoucherCodedepositModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantAccessMode) {
		toSerialize["merchant_access_mode"] = o.MerchantAccessMode
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.VoucherCodes) {
		toSerialize["voucher_codes"] = o.VoucherCodes
	}
	return toSerialize, nil
}

type NullableAlipayMarketingActivityOrdervoucherCodedepositModel struct {
	value *AlipayMarketingActivityOrdervoucherCodedepositModel
	isSet bool
}

func (v NullableAlipayMarketingActivityOrdervoucherCodedepositModel) Get() *AlipayMarketingActivityOrdervoucherCodedepositModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityOrdervoucherCodedepositModel) Set(val *AlipayMarketingActivityOrdervoucherCodedepositModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityOrdervoucherCodedepositModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityOrdervoucherCodedepositModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityOrdervoucherCodedepositModel(val *AlipayMarketingActivityOrdervoucherCodedepositModel) *NullableAlipayMarketingActivityOrdervoucherCodedepositModel {
	return &NullableAlipayMarketingActivityOrdervoucherCodedepositModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityOrdervoucherCodedepositModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityOrdervoucherCodedepositModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
