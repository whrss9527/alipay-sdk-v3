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

// checks if the DeliveryMerchantRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryMerchantRule{}

// DeliveryMerchantRule struct for DeliveryMerchantRule
type DeliveryMerchantRule struct {
	// 指定品牌id。 说明：如商户需选择某个品牌下维护的收款账号，请上传相关品牌id
	BrandIdList []string `json:"brand_id_list,omitempty"`
	// 曝光商户选取列表。 说明：需要传入您期望曝光的商户的商户号，传入为空时默认使用投放优惠券活动的适用范围。 限制：曝光商户号需与投放归属商户号相同，或传入有跨主体授权关系的商户号或有弱绑定关系的M3账号
	DeliveryMerchantInfos []DeliveryMerchantInfo `json:"delivery_merchant_infos,omitempty"`
	// 指定支付成功页模式。
	DeliveryMerchantMode *string `json:"delivery_merchant_mode,omitempty"`
}

// NewDeliveryMerchantRule instantiates a new DeliveryMerchantRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryMerchantRule() *DeliveryMerchantRule {
	this := DeliveryMerchantRule{}
	return &this
}

// NewDeliveryMerchantRuleWithDefaults instantiates a new DeliveryMerchantRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryMerchantRuleWithDefaults() *DeliveryMerchantRule {
	this := DeliveryMerchantRule{}
	return &this
}

// GetBrandIdList returns the BrandIdList field value if set, zero value otherwise.
func (o *DeliveryMerchantRule) GetBrandIdList() []string {
	if o == nil || IsNil(o.BrandIdList) {
		var ret []string
		return ret
	}
	return o.BrandIdList
}

// GetBrandIdListOk returns a tuple with the BrandIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryMerchantRule) GetBrandIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.BrandIdList) {
		return nil, false
	}
	return o.BrandIdList, true
}

// HasBrandIdList returns a boolean if a field has been set.
func (o *DeliveryMerchantRule) HasBrandIdList() bool {
	if o != nil && !IsNil(o.BrandIdList) {
		return true
	}

	return false
}

// SetBrandIdList gets a reference to the given []string and assigns it to the BrandIdList field.
func (o *DeliveryMerchantRule) SetBrandIdList(v []string) {
	o.BrandIdList = v
}

// GetDeliveryMerchantInfos returns the DeliveryMerchantInfos field value if set, zero value otherwise.
func (o *DeliveryMerchantRule) GetDeliveryMerchantInfos() []DeliveryMerchantInfo {
	if o == nil || IsNil(o.DeliveryMerchantInfos) {
		var ret []DeliveryMerchantInfo
		return ret
	}
	return o.DeliveryMerchantInfos
}

// GetDeliveryMerchantInfosOk returns a tuple with the DeliveryMerchantInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryMerchantRule) GetDeliveryMerchantInfosOk() ([]DeliveryMerchantInfo, bool) {
	if o == nil || IsNil(o.DeliveryMerchantInfos) {
		return nil, false
	}
	return o.DeliveryMerchantInfos, true
}

// HasDeliveryMerchantInfos returns a boolean if a field has been set.
func (o *DeliveryMerchantRule) HasDeliveryMerchantInfos() bool {
	if o != nil && !IsNil(o.DeliveryMerchantInfos) {
		return true
	}

	return false
}

// SetDeliveryMerchantInfos gets a reference to the given []DeliveryMerchantInfo and assigns it to the DeliveryMerchantInfos field.
func (o *DeliveryMerchantRule) SetDeliveryMerchantInfos(v []DeliveryMerchantInfo) {
	o.DeliveryMerchantInfos = v
}

// GetDeliveryMerchantMode returns the DeliveryMerchantMode field value if set, zero value otherwise.
func (o *DeliveryMerchantRule) GetDeliveryMerchantMode() string {
	if o == nil || IsNil(o.DeliveryMerchantMode) {
		var ret string
		return ret
	}
	return *o.DeliveryMerchantMode
}

// GetDeliveryMerchantModeOk returns a tuple with the DeliveryMerchantMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryMerchantRule) GetDeliveryMerchantModeOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryMerchantMode) {
		return nil, false
	}
	return o.DeliveryMerchantMode, true
}

// HasDeliveryMerchantMode returns a boolean if a field has been set.
func (o *DeliveryMerchantRule) HasDeliveryMerchantMode() bool {
	if o != nil && !IsNil(o.DeliveryMerchantMode) {
		return true
	}

	return false
}

// SetDeliveryMerchantMode gets a reference to the given string and assigns it to the DeliveryMerchantMode field.
func (o *DeliveryMerchantRule) SetDeliveryMerchantMode(v string) {
	o.DeliveryMerchantMode = &v
}

func (o DeliveryMerchantRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryMerchantRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandIdList) {
		toSerialize["brand_id_list"] = o.BrandIdList
	}
	if !IsNil(o.DeliveryMerchantInfos) {
		toSerialize["delivery_merchant_infos"] = o.DeliveryMerchantInfos
	}
	if !IsNil(o.DeliveryMerchantMode) {
		toSerialize["delivery_merchant_mode"] = o.DeliveryMerchantMode
	}
	return toSerialize, nil
}

type NullableDeliveryMerchantRule struct {
	value *DeliveryMerchantRule
	isSet bool
}

func (v NullableDeliveryMerchantRule) Get() *DeliveryMerchantRule {
	return v.value
}

func (v *NullableDeliveryMerchantRule) Set(val *DeliveryMerchantRule) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryMerchantRule) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryMerchantRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryMerchantRule(val *DeliveryMerchantRule) *NullableDeliveryMerchantRule {
	return &NullableDeliveryMerchantRule{value: val, isSet: true}
}

func (v NullableDeliveryMerchantRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryMerchantRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
