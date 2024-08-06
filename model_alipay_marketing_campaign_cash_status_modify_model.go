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

// checks if the AlipayMarketingCampaignCashStatusModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingCampaignCashStatusModifyModel{}

// AlipayMarketingCampaignCashStatusModifyModel struct for AlipayMarketingCampaignCashStatusModifyModel
type AlipayMarketingCampaignCashStatusModifyModel struct {
	// 金活动修改后的状态。支持修改为： *PAUSE：活动暂停。 *READY：活动开始。 *CLOSED：活动结束。
	CampStatus *string `json:"camp_status,omitempty"`
	// 现金活动号，通过<a href=\"https://opendocs.alipay.com/apis/api_5/alipay.marketing.campaign.cash.create\">alipay.marketing.campaign.cash.create</a>(创建现金活动)接口创建现金活动获取。
	CrowdNo *string `json:"crowd_no,omitempty"`
}

// NewAlipayMarketingCampaignCashStatusModifyModel instantiates a new AlipayMarketingCampaignCashStatusModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingCampaignCashStatusModifyModel() *AlipayMarketingCampaignCashStatusModifyModel {
	this := AlipayMarketingCampaignCashStatusModifyModel{}
	return &this
}

// NewAlipayMarketingCampaignCashStatusModifyModelWithDefaults instantiates a new AlipayMarketingCampaignCashStatusModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingCampaignCashStatusModifyModelWithDefaults() *AlipayMarketingCampaignCashStatusModifyModel {
	this := AlipayMarketingCampaignCashStatusModifyModel{}
	return &this
}

// GetCampStatus returns the CampStatus field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashStatusModifyModel) GetCampStatus() string {
	if o == nil || IsNil(o.CampStatus) {
		var ret string
		return ret
	}
	return *o.CampStatus
}

// GetCampStatusOk returns a tuple with the CampStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashStatusModifyModel) GetCampStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CampStatus) {
		return nil, false
	}
	return o.CampStatus, true
}

// HasCampStatus returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashStatusModifyModel) HasCampStatus() bool {
	if o != nil && !IsNil(o.CampStatus) {
		return true
	}

	return false
}

// SetCampStatus gets a reference to the given string and assigns it to the CampStatus field.
func (o *AlipayMarketingCampaignCashStatusModifyModel) SetCampStatus(v string) {
	o.CampStatus = &v
}

// GetCrowdNo returns the CrowdNo field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashStatusModifyModel) GetCrowdNo() string {
	if o == nil || IsNil(o.CrowdNo) {
		var ret string
		return ret
	}
	return *o.CrowdNo
}

// GetCrowdNoOk returns a tuple with the CrowdNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashStatusModifyModel) GetCrowdNoOk() (*string, bool) {
	if o == nil || IsNil(o.CrowdNo) {
		return nil, false
	}
	return o.CrowdNo, true
}

// HasCrowdNo returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashStatusModifyModel) HasCrowdNo() bool {
	if o != nil && !IsNil(o.CrowdNo) {
		return true
	}

	return false
}

// SetCrowdNo gets a reference to the given string and assigns it to the CrowdNo field.
func (o *AlipayMarketingCampaignCashStatusModifyModel) SetCrowdNo(v string) {
	o.CrowdNo = &v
}

func (o AlipayMarketingCampaignCashStatusModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingCampaignCashStatusModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampStatus) {
		toSerialize["camp_status"] = o.CampStatus
	}
	if !IsNil(o.CrowdNo) {
		toSerialize["crowd_no"] = o.CrowdNo
	}
	return toSerialize, nil
}

type NullableAlipayMarketingCampaignCashStatusModifyModel struct {
	value *AlipayMarketingCampaignCashStatusModifyModel
	isSet bool
}

func (v NullableAlipayMarketingCampaignCashStatusModifyModel) Get() *AlipayMarketingCampaignCashStatusModifyModel {
	return v.value
}

func (v *NullableAlipayMarketingCampaignCashStatusModifyModel) Set(val *AlipayMarketingCampaignCashStatusModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCampaignCashStatusModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCampaignCashStatusModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCampaignCashStatusModifyModel(val *AlipayMarketingCampaignCashStatusModifyModel) *NullableAlipayMarketingCampaignCashStatusModifyModel {
	return &NullableAlipayMarketingCampaignCashStatusModifyModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingCampaignCashStatusModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCampaignCashStatusModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

