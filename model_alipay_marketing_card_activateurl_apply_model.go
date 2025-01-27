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

// checks if the AlipayMarketingCardActivateurlApplyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingCardActivateurlApplyModel{}

// AlipayMarketingCardActivateurlApplyModel struct for AlipayMarketingCardActivateurlApplyModel
type AlipayMarketingCardActivateurlApplyModel struct {
	// 会员卡开卡表单提交后回调地址。要求必须是单纯的服务端接收回调。 说明： 1.该地址不可带参数，如需回传参数，可设置out_string入参。
	Callback *string `json:"callback,omitempty"`
	// 需要关注的生活号AppId。若需要在领卡页面展示“关注生活号”提示，需开通生活号并绑定会员卡。生活号快速接入详见：https://doc.open.alipay.com/docs/doc.htm?treeId=193&articleId=105933&docType=1
	FollowAppId *string `json:"follow_app_id,omitempty"`
	// 该值为商家拉起开卡组件的传递的 out_string  值。通常可用于区分不同业务场景，禁止将该字段作为开卡流程的必要字段，否则会导致会员推广场景下的入会失败
	OutString *string `json:"out_string,omitempty"`
	// 会员卡模板id。使用会员卡模板创建接口(alipay.marketing.card.template.create)返回的结果
	TemplateId *string `json:"template_id,omitempty"`
}

// NewAlipayMarketingCardActivateurlApplyModel instantiates a new AlipayMarketingCardActivateurlApplyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingCardActivateurlApplyModel() *AlipayMarketingCardActivateurlApplyModel {
	this := AlipayMarketingCardActivateurlApplyModel{}
	return &this
}

// NewAlipayMarketingCardActivateurlApplyModelWithDefaults instantiates a new AlipayMarketingCardActivateurlApplyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingCardActivateurlApplyModelWithDefaults() *AlipayMarketingCardActivateurlApplyModel {
	this := AlipayMarketingCardActivateurlApplyModel{}
	return &this
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *AlipayMarketingCardActivateurlApplyModel) GetCallback() string {
	if o == nil || IsNil(o.Callback) {
		var ret string
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardActivateurlApplyModel) GetCallbackOk() (*string, bool) {
	if o == nil || IsNil(o.Callback) {
		return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *AlipayMarketingCardActivateurlApplyModel) HasCallback() bool {
	if o != nil && !IsNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given string and assigns it to the Callback field.
func (o *AlipayMarketingCardActivateurlApplyModel) SetCallback(v string) {
	o.Callback = &v
}

// GetFollowAppId returns the FollowAppId field value if set, zero value otherwise.
func (o *AlipayMarketingCardActivateurlApplyModel) GetFollowAppId() string {
	if o == nil || IsNil(o.FollowAppId) {
		var ret string
		return ret
	}
	return *o.FollowAppId
}

// GetFollowAppIdOk returns a tuple with the FollowAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardActivateurlApplyModel) GetFollowAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.FollowAppId) {
		return nil, false
	}
	return o.FollowAppId, true
}

// HasFollowAppId returns a boolean if a field has been set.
func (o *AlipayMarketingCardActivateurlApplyModel) HasFollowAppId() bool {
	if o != nil && !IsNil(o.FollowAppId) {
		return true
	}

	return false
}

// SetFollowAppId gets a reference to the given string and assigns it to the FollowAppId field.
func (o *AlipayMarketingCardActivateurlApplyModel) SetFollowAppId(v string) {
	o.FollowAppId = &v
}

// GetOutString returns the OutString field value if set, zero value otherwise.
func (o *AlipayMarketingCardActivateurlApplyModel) GetOutString() string {
	if o == nil || IsNil(o.OutString) {
		var ret string
		return ret
	}
	return *o.OutString
}

// GetOutStringOk returns a tuple with the OutString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardActivateurlApplyModel) GetOutStringOk() (*string, bool) {
	if o == nil || IsNil(o.OutString) {
		return nil, false
	}
	return o.OutString, true
}

// HasOutString returns a boolean if a field has been set.
func (o *AlipayMarketingCardActivateurlApplyModel) HasOutString() bool {
	if o != nil && !IsNil(o.OutString) {
		return true
	}

	return false
}

// SetOutString gets a reference to the given string and assigns it to the OutString field.
func (o *AlipayMarketingCardActivateurlApplyModel) SetOutString(v string) {
	o.OutString = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *AlipayMarketingCardActivateurlApplyModel) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardActivateurlApplyModel) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *AlipayMarketingCardActivateurlApplyModel) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *AlipayMarketingCardActivateurlApplyModel) SetTemplateId(v string) {
	o.TemplateId = &v
}

func (o AlipayMarketingCardActivateurlApplyModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingCardActivateurlApplyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	if !IsNil(o.FollowAppId) {
		toSerialize["follow_app_id"] = o.FollowAppId
	}
	if !IsNil(o.OutString) {
		toSerialize["out_string"] = o.OutString
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	return toSerialize, nil
}

type NullableAlipayMarketingCardActivateurlApplyModel struct {
	value *AlipayMarketingCardActivateurlApplyModel
	isSet bool
}

func (v NullableAlipayMarketingCardActivateurlApplyModel) Get() *AlipayMarketingCardActivateurlApplyModel {
	return v.value
}

func (v *NullableAlipayMarketingCardActivateurlApplyModel) Set(val *AlipayMarketingCardActivateurlApplyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardActivateurlApplyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardActivateurlApplyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardActivateurlApplyModel(val *AlipayMarketingCardActivateurlApplyModel) *NullableAlipayMarketingCardActivateurlApplyModel {
	return &NullableAlipayMarketingCardActivateurlApplyModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardActivateurlApplyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardActivateurlApplyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
